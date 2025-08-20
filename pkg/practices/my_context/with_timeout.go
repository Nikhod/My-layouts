package my_context

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

var lifetime = time.Second * 5

// запрос будет жить lifetime времени. Если все компоненты этого запроса не успеют выполнится в течение этого времени
// запрос просто перестанет существовать
func createRequestSampleWithTimeout(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), lifetime)
	defer cancel()
	var value innerService
	_ = json.NewDecoder(r.Body).Decode(&value)
	_ = serviceEmulation(&value)
	_ = saveToDB(ctx, value, &sql.DB{})
}

type innerService struct {
	firstField  string
	secondField int
	thirdField  bool
}

// emulation of service layer
func serviceEmulation(value *innerService) bool {
	if len(value.firstField) > 5 {
		if value.secondField != 0 {
			if value.thirdField == true {
				return true
			}
		}
	}
	return false
}

func saveToDB(ctx context.Context, value innerService, db *sql.DB) error {
	query := `--insert into values(first, second, third)
--values (#1, #2, #3)`
	_, _ = db.ExecContext(ctx, query, value.firstField, value.secondField, value.thirdField)
	return nil
}
