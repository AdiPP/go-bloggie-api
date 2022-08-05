package infrastructure

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func InitResty() *resty.Client {
	rest := resty.New()
	rest.SetTimeout(5 * time.Second)
	return rest
}
