package returning_func

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// ОГРАНИЧИТЕЛЬ ЗАПРОСОВ
// Контекст:
//У тебя есть система, которая обращается к внешнему API (например, к API цен криптовалют или погодных данных). Но API имеет ограничение: не более 5 запросов в минуту.
//
//Тебе нужно реализовать компонент, который:
//
//будет возвращать функцию (замыкание),
//
//будет "разрешать" делать запрос только если не превышен лимит (5 в минуту),
//
//иначе будет отказывать (например, возвращать nil, error или специальный маркер).

func NewRateLimiterPractice() {
	validLimit := myNewLimiter()

	for i := 0; i < 5; i++ {
		if validLimit() {
			_ = sendReq()
		} else {
			log.Println("too many request")
		}
		if i == 4 {
			time.Sleep(time.Second * 65)
			continue
		}
		time.Sleep(time.Second * 10)
	}

}

const (
	limit            = 5
	lifeTimeDuration = time.Minute
)

func myNewLimiter() func() bool {
	var successfulTimestamp []time.Time

	return func() bool {
		var filtered []time.Time // количество завершившихся запросов в 1 минуту
		now := time.Now()

		for _, ts := range successfulTimestamp {
			if len(successfulTimestamp) == 0 {
				break
			}
			// в "filtered" запихиваем запросы которые входят в временной промежуток
			// если разница между нынешним временем и временем
			//успешного запроса в течение 1 минуты
			if dif := now.Sub(ts); dif <= lifeTimeDuration {
				fmt.Println("difference:", dif.Seconds(), "ts: ", ts.Second(), "now: ", now.Second())
				fmt.Println()
				filtered = append(filtered, ts)
			}
		}

		if len(filtered) < limit {
			currentReqTime := time.Now()
			filtered = append(filtered, currentReqTime)
			successfulTimestamp = filtered //  обновляем общий котел с запросами, добавив актуальные, старые - уже не нужны
			return true
		}

		return false
	}

}

// for emulation
func sendReq() error {
	log.Print("request has been sent\t")
	return nil
	q := url.Values{}
	q.Add("action", "check")
	q.Add("type_of_req", "ordinary")
	resultUrl := q.Encode()

	request, err := http.NewRequest(http.MethodGet, resultUrl, nil)
	if err != nil {
		return err
	}
	cl := http.Client{Timeout: time.Second * 15}
	response, err := cl.Do(request)
	log.Println("request is sent")

	if err != nil {
		return err
	}
	fmt.Println(response)
	return nil
}

// Итог:
// мы проходимся по срезу с временными метками, и выбираем оттуда только те, которые проходят по лимитному временному
// окну, добавляем в отдельный срез. Кол-во эл-тов в этом срезе = кол-во успешных запросов за 1 минуту (или другое
// лимитное время). Проверяем кол-во попыток (len.filtered), если все ок --> добавляем новую временную метку в срез,
// обновляем общий срез актуальными вр-ми метками
