package repository

import (
	"Go-upload/model"
	"errors"

	"gorm.io/gorm"
)

type FileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) *FileRepository {
	return &FileRepository{
		db: db,
	}
}

func (fr *FileRepository) Create(newFile model.File) (model.File, error) {

	tx := fr.db.Save(&newFile)

	if tx.Error != nil {
		return model.File{}, tx.Error
	}
	return newFile, nil
}

func (Fr *FileRepository) GetAll() ([]model.File, error) {
	var file []model.File

	tx := Fr.db.Find(&file)

	return file, tx.Error
}

func (Fr *FileRepository) GetByID(id string) (model.File, error) {
	file := model.File{}

	tx := Fr.db.First(&file, "id = ?", id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return model.File{}, model.ErrorNotFound
		}
		return model.File{}, tx.Error
	}

	return file, nil

}

func (Fr *FileRepository) DeleteFile(id string) error {

	file := model.File{}
	if err := Fr.db.Where("id = ?", id).First(&file).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
		return err
	}

	if err := Fr.db.Delete(&file).Error; err != nil {
		return err
	}
	return nil

}
