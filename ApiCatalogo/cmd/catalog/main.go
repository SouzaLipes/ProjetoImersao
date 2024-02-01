package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/devfullcycle/imersaoluis/goapi/internal/database"
	"github.com/devfullcycle/imersaoluis/goapi/internal/service"
	"github.com/devfullcycle/imersaoluis/goapi/internal/webserver"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	CategoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	ProductService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(CategoryService)
	webProductHandler := webserver.NewWebProductHandler(ProductService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Get("/products/category/{category_id}", webProductHandler.GetProductByCategoryID)
	c.Post("/products", webProductHandler.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
