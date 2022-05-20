package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Renos-id/go-starter-template/database"
	"github.com/Renos-id/go-starter-template/infrastructure"
	_ "github.com/lib/pq"
)

func init() {
	infrastructure.InitLoadEnv()
}

func main() {
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
	infrastructure.InitZapLogger()
	
	port := os.Getenv("APP_PORT")
	if port == "" { 
		port = "8080"
	}

	fmt.Printf("%s running on PORT : %s \n", os.Getenv("APP_NAME"), port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
