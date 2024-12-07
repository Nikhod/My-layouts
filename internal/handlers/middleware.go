package handlers

import (
	"Nikcase/pkg/Helper"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func (h *Handlers) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		//код который заработает до Хэндлера
		login, pass := request.Header.Get("login"), request.Header.Get("password")
		if login == "" || pass == "" {
			answer, err := Helper.ResponseAnswer("Логин или пароль не введдены!")
			if err != nil {
				log.Println(err)
				return
			}
			_, err = writer.Write(answer)
			if err != nil {
				log.Println(err)
				return
			}
			return
		}
		user, err := h.Service.ValidatePassAndLogin(login, pass)
		if err != nil {
			log.Println(err)
			return
		}

		ctx := context.WithValue(request.Context(), "userID", user.Id)
		request = request.WithContext(ctx)
		next.ServeHTTP(writer, request)
		// код который заработает после выполнения Хендлера
	})
}
