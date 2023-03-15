package main

import (
    "GO-CRUD/models"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    models.InitDb()

    router.Run("localhost:8080")
}
