package infrastructure

import (
	"github.com/adiputraaa/bloggie/pkg/controller"
	"github.com/ggicci/httpin"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func init() {
	httpin.UseGochiURLParam("path", chi.URLParam)
}

func NewRouter(userController *controller.UserController, postController *controller.PostController) *chi.Mux {
	router := chi.NewRouter()

	// A good base middleware stack
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Heartbeat("/health"))

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	router.Use(middleware.Timeout(60 * time.Second))

	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hi"))
	})

	router.With(httpin.NewInput(controller.GetUserInput{})).
		Get("/users/{user}", userController.GetUser)
	router.Get("/users", userController.FetchUsers)
	router.Get("/posts", postController.FetchPosts)

	return router
}
