package handlers

import (
	"context"
	"net/http"

	"github.com/supercoast/profile-api/data"
)

func (p *Profile) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")

		profile, err := data.FromJSON(r.Body)
		if err != nil {
			p.Logger.Error("Couldn't deserialize profile", "error", err)

			rw.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		err = p.Validator.Validate(profile)
		if err != nil {
			p.Logger.Error("Validation for input failed", "error", err)

			rw.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&GenericError{Message: err.Error()}, rw)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, profile)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
