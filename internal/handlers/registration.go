package handlers

import (
	"Nikcase/pkg/Helper"
	"Nikcase/pkg/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handlers) Registration(w http.ResponseWriter, r *http.Request) {
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.IsValidDataForRegistration(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.Service.RegistrationUser(&user)
	if err != nil {
		ResponseBytes, err := Helper.ResponseAnswer("Логин Занят другим пользователем!")
		if err != nil {
			log.Println(err)
			return
		}
		_, err = w.Write(ResponseBytes)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	answer, err := Helper.ResponseAnswer("Вы Успешно Зарегистрировались!")
	if err != nil {
		log.Println(err)
		return
	}
	_, err = w.Write(answer)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("FINISH Registration!")
}
