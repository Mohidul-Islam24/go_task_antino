package blogs

import (
	"go_task_antino/db"
	"go_task_antino/dbqueries"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeletePost(c *gin.Context) {
	id := c.Query("id")

	_, e := db.Driver.Exec(dbqueries.DeletePost, id)
	if e != nil {
		c.JSON(http.StatusInternalServerError, e.Error)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "post deleted",
	})
}
