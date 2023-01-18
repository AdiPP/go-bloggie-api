package entities

type Post struct {
	UserId int64  `json:"user_id,omitempty"`
	Id     int64  `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Body   string `json:"body,omitempty"`
}

type Posts []Post
