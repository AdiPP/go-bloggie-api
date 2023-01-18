package entities

type User struct {
	Id       int64  `json:"id,omitempty" db:"id"`
	Username string `json:"username,omitempty" db:"username"`
	Password string `json:"password,omitempty" db:"password"`
	Email    string `json:"email,omitempty" db:"email"`
}

type Users []User

type UserWithPosts struct {
	User
	Posts
}
