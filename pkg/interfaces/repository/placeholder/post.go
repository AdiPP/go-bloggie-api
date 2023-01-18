package placeholder

type Post struct {
	UserId int64  `json:"userId,omitempty"`
	Id     int64  `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

type Posts []Post
