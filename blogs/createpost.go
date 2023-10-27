package blogs

import (
	"fmt"
	"go_task_antino/db"
	"go_task_antino/dbqueries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateNewPost(c *gin.Context) {
	fmt.Println("CreateNewPost")
	title := c.Query("title")
	content := c.Query("content")
	var id int32

	//db operation
	e := db.Driver.QueryRow(dbqueries.CreateNewPost, title, content).Scan(&id)
	if e != nil {
		c.JSON(http.StatusInternalServerError, e.Error)
		return
	}
	fmt.Println("new blog post with id: %d has created", id)
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}
