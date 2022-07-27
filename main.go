package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


func main() {
	fmt.Println("Start server...")
	engine:= gin.Default()
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "hello world test",
        })
    })
    engine.Run(":3000")
	// router := gin.Default()
	// db := infrastructure.DbInit()
	// defer db.Close()
	
	// router.GET("/", func(c *gin.Context) {
	// 		todos := infrastructure.DbRead()
	// 		c.HTML(200, "index.html", gin.H{
	// 				"todos": todos,
	// 		})
	// })
	
	// router.POST("/new", func(c *gin.Context) {
	// 		text := c.PostForm("text")
	// 		rawStatus := c.PostForm("status")
	// 		id, err := strconv.Atoi(rawStatus)
	// 		if err != nil {
	// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		}
	// 		status := domain.Status(id)
	// 		rawTime := c.PostForm("deadline")
	// 		deadline, err := strconv.Atoi(rawTime)
	// 		if err != nil {
	// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		}
	// 		todo := domain.Todo{Text: text, Status: status, Deadline: deadline}
	
	// 		if err != nil {
	// 				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 				return
	// 		}
	// 		infrastructure.DbCreate(todo)
	// 		c.Redirect(302, "/")
	// })
}