package blogs

import (
	"go_task_antino/db"
	"go_task_antino/dbqueries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdatePost(c *gin.Context) {
	id := c.Query("id")
	title := c.Query("title")
	content := c.Query("content")

	_, e := db.Driver.Exec(dbqueries.UpdatePost, id, title, content)
	if e != nil {
		c.JSON(http.StatusInternalServerError, e.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "post updated",
	})
}
