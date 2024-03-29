package main

import (
	"Go-upload/controller"
	"Go-upload/database"
	"Go-upload/repository"
	"Go-upload/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.ReadDB()

	fileRepository := repository.NewFileRepository(db)
	fileService := service.NewFileService(fileRepository)
	fileController := controller.NewFileController(fileService)

	r := gin.Default()

	r.POST("/upload", fileController.UploadFile)
	r.GET("/get", fileController.GetAll)
	r.GET("/get/:id", fileController.GetByID)
	r.DELETE("/delete/:id", fileController.DeleteFIle)

	r.Run(":8088")

}
