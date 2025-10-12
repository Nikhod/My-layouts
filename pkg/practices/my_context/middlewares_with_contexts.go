package my_context

import (
	"context"
	"net/http"
	"time"
)

const MyKey = "qwerty"

func GetSetOfContext(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		valueContext := context.WithValue(r.Context(), MyKey, "value_from_middleware")
		timeout, cancelFunc := context.WithTimeout(valueContext, time.Second*3)
		defer cancelFunc()
		r = r.WithContext(timeout)

		next.ServeHTTP(w, r)
	}
}
