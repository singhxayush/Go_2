package main

import (
	"fmt"
	"go-fiber-psql/models"
	"go-fiber-psql/storage"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Repository contains a reference to DB and all CRUD methods
type Repository struct {
	DB *gorm.DB
}

func (r *Repository) setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/get_books", r.getAllBooks)
	api.Get("/get_book/:id", r.GetBook)
	api.Post("/create_books", r.createBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
}

// Book stores the structure of a single book
type Book struct {
	Author   string `json:"author"`
	Title    string `json:"title"`
	Publiser string `json:"publisher"`
}

func (r *Repository) createBook(c *fiber.Ctx) error {
	book := Book{}
	err := c.BodyParser(&book)
	if err != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"message": "Request Failed"})
		return err
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "Couldn't create Book"})
		return err
	}

	c.Status(http.StatusOK).JSON(fiber.Map{"message": "Book has been added"})

	return nil
}

func (r *Repository) getAllBooks(c *fiber.Ctx) error {
	bookModels := &[]models.Books{}
	err := r.DB.Find(bookModels).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "could not get the books"})
		return err
	}

	c.Status(http.StatusOK).JSON(fiber.Map{"message": "books fetched succfully", "data": bookModels})

	return nil
}

// DeleteBook deletes a book by its given ID in the request params
func (r *Repository) DeleteBook(c *fiber.Ctx) error {
	// Get the book ID from the request parameters
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "ID cannot be empty"})
	}

	// Check if the book exists
	book := &models.Books{}
	if err := r.DB.First(book, id).Error; err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"message": "book not found"})
	}

	// Delete the book
	if err := r.DB.Delete(book, id).Error; err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "could not delete book"})
	}

	// Respond with success message
	return c.Status(http.StatusOK).JSON(fiber.Map{"message": "book deleted successfully"})
}

// GetBook is to get a specific book by it's ID
func (r *Repository) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	bookModel := &models.Books{}
	if id == "" {
		c.Status(http.StatusInternalServerError).JSON(fiber.Map{"message": "id cannot be empty"})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(bookModel).Error

	if err != nil {
		c.Status(http.StatusBadRequest).JSON(fiber.Map{"message": "could not get the book"})
		return err
	}

	c.Status(http.StatusOK).JSON(fiber.Map{"message": "book id fetched successfully", "data": bookModel})

	return nil
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLM"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("DB Connection Failed", err)
	}

	err = models.MigrateBooks(db)

	r := Repository{
		DB: db,
	}

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		ServerHeader:      "Go-PSQL Todo App",
	})

	app.Use(logger.New())

	r.setupRoutes(app)

	// port := os.Getenv("PORT")

	log.Fatal(app.Listen(":8080"))
}
