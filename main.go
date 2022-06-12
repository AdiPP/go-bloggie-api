package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Renos-id/go-starter-template/database"
	"github.com/Renos-id/go-starter-template/infrastructure"

	_ "github.com/lib/pq"
)

func init() {
	infrastructure.InitLoadEnv()
}

func main() {
	// The HTTP Server
	server := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", os.Getenv("APP_PORT")), Handler: service()}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		fmt.Println("Trigger graceful shutdown")
		err := server.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()
	}()

	// Run the server
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

func service() http.Handler {
	//init DB
	if os.Getenv("DB_HOST") != "" {
		dbConn := database.Open()
		defer func() {
			err := dbConn.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	//End Init DB
	r := infrastructure.InitChiRouter()

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("%s running on PORT : %s \n", os.Getenv("APP_NAME"), port)
	return r
}
