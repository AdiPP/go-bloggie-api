package infrastructure

import "github.com/go-resty/resty/v2"

func InitResty() *resty.Client {
	rest := resty.New()
	return rest
}
