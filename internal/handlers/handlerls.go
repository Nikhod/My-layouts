package handlers

import (
	"Nikcase/internal/services"
	"Nikcase/pkg/Helper"
	"Nikcase/pkg/cache"
	"Nikcase/pkg/models"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var CACHED cache.CachedLimits
var DB *gorm.DB

type Handlers struct {
	Service *services.Service
}

func NewHandlers(service *services.Service) *Handlers {
	return &Handlers{Service: service}
}

func (h *Handlers) GetToken(w http.ResponseWriter, r *http.Request) {
	log.Println("Start!")
	var user models.Users
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		return
	}

	token, err := h.Service.GenerateToken(user.Login, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		defer log.Println(err)

		answer, err := Helper.ResponseAnswer("Такого пользователя НЕТ в Базе Данных!")
		if err != nil {
			log.Println(err)
			return
		}

		_, err = w.Write(answer)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	sendToken := models.SendToken{
		Date:  time.Now(),
		Token: token,
	}
	answer, err := json.Marshal(sendToken)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(answer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("ERROR Write Answer: ", err)
		return
	}
	log.Println("Finish!")
}

func (h *Handlers) CreateLimit(w http.ResponseWriter, r *http.Request) {
	var limit models.AmountLimits
	err := json.NewDecoder(r.Body).Decode(&limit)
	if err != nil {
		log.Println(err)
		return
	}

	CACHED.SetLimit()

}

func SetLimitToDB() error {
	DB.Create()
}
