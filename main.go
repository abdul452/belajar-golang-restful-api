package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/abdul452/belajar-golang-restful-api/app"
	"github.com/abdul452/belajar-golang-restful-api/controller"
	"github.com/abdul452/belajar-golang-restful-api/helper"
	"github.com/abdul452/belajar-golang-restful-api/middleware"
	"github.com/abdul452/belajar-golang-restful-api/repository"
	"github.com/abdul452/belajar-golang-restful-api/service"
	"github.com/go-playground/validator/v10"
)

func main() {
	// http router
	db := app.NewDatabase()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	// http server
	server := http.Server{
		Addr:    ":3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
