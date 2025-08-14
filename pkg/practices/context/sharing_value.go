package context

import (
	"context"
	"net/http"
)

const KeyUserId string = "my_user_key"

// one can use this func for taking user_id, but don't forget
// rename the func to naming starting from great register
func getUserIdByContext(ctx context.Context) (userID int) {
	return ctx.Value(KeyUserId).(int)
}

// emulation of working middleware
func UserIdWriter(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := takeUserIdByLoginAndPass("my_login", "my_hashed_pass")
		ctx := context.WithValue(r.Context(), KeyUserId, userID)
		r = r.WithContext(ctx) // перезапись контекста http запроса - важно
		next.ServeHTTP(w, r)
		// after endpoint's code
	}
}

// emulation of request to database
func takeUserIdByLoginAndPass(login, pass string) (userID int) {
	return len(login) + len(pass)
}
