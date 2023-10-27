package dbqueries

const (
	CreateNewPost = "INSERT INTO blogs(id,title,content) VALUES(default,$1,$2) RETURNING id"
	DeletePost    = "DELETE FROM blogs WHERE id=$1"
	UpdatePost    = "UPDATE blogs SET title=$2,content=$3 WHERE id=$1"
	GetPostByID   = "SELECT title,content FROM blogs WHERE id=$1"
	GetAllPost    = "SELECT id,title,content FROM blogs ORDER BY id"
)
