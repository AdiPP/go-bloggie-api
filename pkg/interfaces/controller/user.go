package controller

import "github.com/adiputraaa/bloggie/pkg/entities"

type UserWithPosts struct {
	entities.User  `json:"user"`
	entities.Posts `json:"posts"`
}
