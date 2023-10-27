package main

import (
	"fmt"
	"go_task_antino/blogs"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	// r.ForwardedByClientIP = true
	// r.SetTrustedProxies([]string{"127.0.0.1"})
	gin.SetMode(gin.DebugMode)

	// endpoints for CURD operation
	r.POST("/createpost", blogs.CreateNewPost)
	r.GET("/getAllPost", blogs.GetAllPost)
	r.GET("/getPostByID", blogs.GetPostByID)
	r.PUT("/updatePost", blogs.UpdatePost)
	r.DELETE("/deletePost", blogs.DeletePost)

	// starting the server
	e := r.Run(":" + "8080")
	if e != nil {
		panic(e.Error())
	}
	fmt.Println("Server is runing on port 8080")

}
