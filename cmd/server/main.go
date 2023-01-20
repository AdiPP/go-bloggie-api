package main

import (
	"github.com/adiputraaa/bloggie/pkg/controller"
	"github.com/adiputraaa/bloggie/pkg/gateway/placeholder"
	"github.com/adiputraaa/bloggie/pkg/gateway/sqlite"
	"github.com/adiputraaa/bloggie/pkg/infrastructure"
	"github.com/adiputraaa/bloggie/pkg/usecases/post"
	"github.com/adiputraaa/bloggie/pkg/usecases/user"
	"github.com/doug-martin/goqu/v9"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln(err)
		return
	}
}

func main() {
	var (
		err error
		qb  *goqu.Database
	)

	if qb, err = infrastructure.NewQueryBuilder(); err != nil {
		log.Println(err)
		return
	}

	userRepository := sqlite.NewRepository(qb)
	postRepository := placeholder.NewRepository()

	userInteractor := user.NewInteractor(userRepository)
	postInteractor := post.NewInteractor(postRepository)

	userController := controller.NewUserController(userInteractor, postInteractor)
	postController := controller.NewPostController(postInteractor)

	router := infrastructure.NewRouter(userController, postController)

	appName := os.Getenv("APP_NAME")
	port := os.Getenv("APP_PORT")

	log.Printf("%s running on PORT :%s \n", appName, port)

	if err = http.ListenAndServe(":"+port, router); err != nil {
		log.Println(err)
		return
	}
}
