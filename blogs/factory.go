package blogs

type Blog struct {
	Id      int32  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Blogs []Blog
