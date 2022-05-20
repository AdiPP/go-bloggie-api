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

	fmt.Printf("%s running on PORT : %s \n", os.Getenv("APP_NAME"), os.Getenv("APP_PORT"))
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), r)
}
