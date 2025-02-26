package main

import (
    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
    "net/http"
    "log"
)

// Define the User struct
type User struct {
    Email string `json:"email"`
    Password string `json:"password"`
}

// Simulated in-memory database (map) of users
var usersDatabase = map[string]User{
    "admin@gmail.com": {"admin@gmail.com", "admin"},
}

func main() {
    r := gin.Default()

	// Enable CORS and allow only http://localhost:5173
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},  // Only allow requests from localhost:5173
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Allow HTTP methods
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allow these headers
        AllowCredentials: true, // Allow credentials (cookies, authorization headers, etc.)
    }))

    // Define a POST route for login
    r.POST("/login", login)

    // Run the server on port 8080
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Unable to start server:", err)
    }
}

func login(c *gin.Context) {
    // Declare a variable of type User to bind the JSON request body to
    var user User

    // Bind the request JSON to the user struct
    if err := c.ShouldBindJSON(&user); err != nil {
        // If there is an error binding the JSON, return a bad request error
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    // Check if the user exists in the "database"
    storedUser, exists := usersDatabase[user.Email]
    if !exists {
        // If the user does not exist, return an error
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
        return
    }

    // Validate the password
    if storedUser.Password != user.Password {
        // If the password is incorrect, return an unauthorized error
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect password"})
        return
    }

    // If Email and password are correct, return a success response
    c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
