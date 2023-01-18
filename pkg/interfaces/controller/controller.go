package controller

import (
	"encoding/json"
	"github.com/adiputraaa/bloggie/pkg/usecases/post"
	"github.com/adiputraaa/bloggie/pkg/usecases/user"
	"github.com/ggicci/httpin"
	"net/http"
)

type UserController struct {
	UserInteractor *user.Interactor
	PostInteractor *post.Interactor
}

func NewUserController(userInteractor *user.Interactor, postInteractor *post.Interactor) *UserController {
	return &UserController{UserInteractor: userInteractor, PostInteractor: postInteractor}
}

type GetUserInput struct {
	User int64 `in:"path=user"`
}

func (uc *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	input := r.Context().Value(httpin.Input).(*GetUserInput)

	getUser, err := uc.UserInteractor.GetUser(r.Context(), &user.GetUserParam{
		User: input.User,
	})

	posts, err := uc.PostInteractor.FetchPostsByUser(r.Context(), &post.FetchPostsByUserParam{
		User: getUser.Id,
	})

	result := UserWithPosts{
		User:  getUser,
		Posts: posts,
	}

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
	return
}

func (uc *UserController) FetchUsers(w http.ResponseWriter, r *http.Request) {
	users, err := uc.UserInteractor.FetchUsers(r.Context())

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
	return
}

type PostController struct {
	PostInteractor *post.Interactor
}

func NewPostController(postInteractor *post.Interactor) *PostController {
	return &PostController{PostInteractor: postInteractor}
}

func (pc *PostController) FetchPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := pc.PostInteractor.FetchPosts(r.Context())

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
	return
}
