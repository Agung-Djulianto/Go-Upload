package service

import (
	"Go-upload/helper"
	"Go-upload/model"
	"Go-upload/repository"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type FileService struct {
	FileRepository *repository.FileRepository
}

func NewFileService(FileRepository *repository.FileRepository) *FileService {
	return &FileService{
		FileRepository: FileRepository,
	}
}

func (fs *FileService) UploadFile(request model.FileRequest, file multipart.File, header *multipart.FileHeader) (model.FileResponse, error) {
	// Validasi jenis file (contoh: hanya izinkan file gambar)
	if !helper.IsValidImageFile(header) {
		return model.FileResponse{}, errors.New("Invalid file type. Only image files are allowed")
	}

	// Simpan file ke dalam sistem
	filename := header.Filename
	filePath := filepath.Join("asset", filename)
	err := saveFile(file, filePath)
	if err != nil {
		return model.FileResponse{}, err
	}

	// Buat entri file di database
	id := helper.GenerateID()
	fileData := model.File{
		ID:       id,
		FileName: filename,
	}
	createdFile, err := fs.FileRepository.Create(fileData)
	if err != nil {
		return model.FileResponse{}, err
	}

	return model.FileResponse{
		ID:       createdFile.ID,
		FileName: createdFile.FileName,
	}, nil
}

func saveFile(file multipart.File, filePath string) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	return err
}

func (Fs *FileService) GetAll() ([]model.FileResponse, error) {
	fileResponse := make([]model.FileResponse, 0)

	response, err := Fs.FileRepository.GetAll()

	if err != nil {
		return []model.FileResponse{}, err
	}

	for _, value := range response {
		fileResponse = append(fileResponse, model.FileResponse{
			ID:        value.ID,
			FileName:  value.FileName,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		})
	}

	return fileResponse, err

}

func (Fs *FileService) GetByID(id string) (model.FileResponse, error) {
	response, err := Fs.FileRepository.GetByID(id)

	if err != nil {
		if err == model.ErrorNotFound {
			return model.FileResponse{}, model.ErrorNotFound
		}
		return model.FileResponse{}, err
	}

	return model.FileResponse{
		ID:        response.ID,
		FileName:  response.FileName,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
	}, nil
}

func (Fs *FileService) DeleteFIle(id string) error {
	file, err := Fs.FileRepository.GetByID(id)

	if err != nil {
		return err
	}

	filePath := filepath.Join("asset", file.FileName)
	if err := os.Remove(filePath); err != nil {
		return err
	}

	if err := Fs.FileRepository.DeleteFile(id); err != nil {
		return err
	}

	return nil
}
