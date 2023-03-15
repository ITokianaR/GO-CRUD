package main

import (
    "GO-CRUD/models"
    "GO-CRUD/controllers"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.InitDb()

    //routes list
    router.POST("/post", controllers.CreatePost)

    router.Run("localhost:8080")
}
