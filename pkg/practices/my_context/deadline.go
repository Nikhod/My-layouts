package my_context

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

// todo initialize several context in http request
// todo initialize several context in several Middlewares

func createRequestSampleWithDeadline(w http.ResponseWriter, r *http.Request) {
	//ctx, cancel := context.WithTimeout(r.Context(), lifetime)
	ctx, cancel := context.WithDeadline(r.Context(), time.Now().Add(time.Second*5))
	defer cancel()
	var value innerService
	_ = json.NewDecoder(r.Body).Decode(&value)
	_ = serviceEmulation(&value)
	_ = saveToDB(ctx, value, &sql.DB{})
}
