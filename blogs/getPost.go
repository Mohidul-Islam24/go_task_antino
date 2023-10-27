package blogs

import (
	"go_task_antino/db"
	"go_task_antino/dbqueries"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllPost(c *gin.Context) {
	var blogs Blogs

	res, e := db.Driver.Query(dbqueries.GetAllPost)
	if e != nil {
		c.JSON(http.StatusInternalServerError, e.Error)
		return
	}

	defer res.Close()

	for res.Next() {
		var id int32
		var title, content string
		var blog Blog

		er := res.Scan(&id, &title, &content)
		if er != nil {
			c.JSON(http.StatusInternalServerError, e.Error)
			return
		}
		blog.Id = id
		blog.Title = title
		blog.Content = content

		blogs = append(blogs, blog)
	}
	c.JSON(http.StatusOK, blogs)
}

func GetPostByID(c *gin.Context) {
	var blog Blog
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id can't be empty",
		})
		return
	}
	var title, content string

	e := db.Driver.QueryRow(dbqueries.GetPostByID, id).Scan(&title, &content)
	if e != nil {
		c.JSON(http.StatusInternalServerError, e.Error)
		return
	}
	intID, _ := strconv.Atoi(id)
	blog = Blog{
		Id:      int32(intID),
		Title:   title,
		Content: content,
	}

	c.JSON(http.StatusOK, blog)
}
