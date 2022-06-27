package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Renos-id/go-starter-template/lib/response"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"github.com/sirupsen/logrus"
)

var (
	v *validator.Validate
)

// HTTP middleware setting a value on the request context
func RequestValidation[K interface{}](data K) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logrus.Error("Error Create ROom")
			resp := response.CommonResponse{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&data)
			v = validator.New()
			en := en.New()
			uni := ut.New(en, en)
			trans, _ := uni.GetTranslator("id")
			_ = en_translations.RegisterDefaultTranslations(v, trans)
			if err != nil {
				resp = response.WriteError(500, "Failed decode Request in Request Validation", err)
				resp.ToJSON(w)
				return
			}
			err = v.Struct(data)
			if err != nil {
				errs := translateError(err, trans)
				resp = response.WriteError(422, "Validation Failed", errs)
				resp.ToJSON(w)
				return
			}
			ctx := context.WithValue(r.Context(), "body", data)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func translateError(err error, trans ut.Translator) (errs response.ValidationErrors) {
	var errors response.ValidationErrors
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Sprintf(e.Translate(trans))
		errs = append(errors, response.ValidationError{
			Field:   strings.ToLower(e.Field()),
			Message: strings.Replace(translatedErr, "_", " ", 1),
		})
	}
	return errs
}
