package placeholder

import (
	"context"
	"encoding/json"
	"github.com/adiputraaa/bloggie/pkg/entities"
	"github.com/adiputraaa/bloggie/pkg/repository"
	"net/http"
)

type Client struct {
	client *http.Client
}

func NewRepository() *Client {
	return &Client{client: &http.Client{}}
}

func (pr *Client) FindAllPosts(_ context.Context) (posts entities.Posts, err error) {
	request, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)

	if err != nil {
		return
	}

	response, err := pr.client.Do(request)

	if err != nil {
		return
	}

	defer response.Body.Close()

	var (
		ps Posts
	)

	err = json.NewDecoder(response.Body).Decode(&ps)

	for _, p := range ps {
		posts = append(posts, entities.Post{
			UserId: p.UserId,
			Id:     p.Id,
			Title:  p.Title,
			Body:   p.Body,
		})
	}

	return
}

func (pr *Client) FindAllPostsByUser(ctx context.Context, param *repository.FindAllPostsByUserParam) (posts entities.Posts, err error) {
	ps, err := pr.FindAllPosts(ctx)

	if err != nil {
		return
	}

	for _, p := range ps {
		if p.UserId == param.User {
			posts = append(posts, p)
		}
	}

	return
}
