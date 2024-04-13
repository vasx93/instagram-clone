package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vasx93/instagram-clone/internal/db"
	"github.com/vasx93/instagram-clone/internal/user"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	gin.SetMode(os.Getenv("GIN_MODE"))

	router := gin.Default()

	db := db.DBConnection()

	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(userRepository)

	// setup controller and routes
	user.NewUserController(userService, router)

	PORT := os.Getenv("PORT")
	err := router.Run(":" + PORT)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server listening on port:", PORT)

}

// func getUsers(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"message": "users incoming soon!",
// 	})
// }
