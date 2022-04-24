package main

import (
	"golang-CRUD/handler"
	"golang-CRUD/user"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&user.User{})

	userRepository := user.NewUserRepository(db)
	serviceRepository := user.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(serviceRepository)

	router := gin.Default()
	users := router.Group("/users")
	users.GET("", (*userHandler).FindAll)
	users.GET("/:ID", (*userHandler).FindByID)
	users.POST("", (*userHandler).Create)
	users.PUT("/:ID", (*userHandler).Update)
	users.DELETE("/:ID", (*userHandler).Delete)

	log.Fatal(http.ListenAndServe(":8080", router))
}
