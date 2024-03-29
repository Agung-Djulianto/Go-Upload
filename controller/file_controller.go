package controller

import (
	"Go-upload/model"
	"Go-upload/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Tambahkan definisi struct ResponseFailed di sini

type FileController struct {
	FileService *service.FileService
}

func NewFileController(fileservice *service.FileService) *FileController {
	return &FileController{
		FileService: fileservice,
	}
}

func (fc *FileController) UploadFile(ctx *gin.Context) {

	var newFile model.FileRequest

	if err := ctx.ShouldBind(&newFile); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	response, err := fc.FileService.UploadFile(newFile, file, header)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusCreated, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		},
		Data: response,
	})
}

func (Fc *FileController) GetAll(ctx *gin.Context) {
	file, err := Fc.FileService.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: file,
	})
}

func (Fc *FileController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")

	file, err := Fc.FileService.GetByID(id)
	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: "file" + err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	filePath := "asset/" + file.FileName

	_, err = os.Stat(filePath)
	if os.IsNotExist(err) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusNotFound,
				Message: http.StatusText(http.StatusNotFound),
			},
			Error: err.Error(),
		})
		return
	}

	// Mengirimkan file sebagai respons
	ctx.File(filePath)
}

func (Fc *FileController) DeleteFIle(ctx *gin.Context) {
	id := ctx.Param("id")

	err := Fc.FileService.DeleteFIle(id)
	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: "File not found",
			})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: "Delete file Successfully",
	})
}
