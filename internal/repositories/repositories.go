package repositories

import (
	"Nikcase/pkg/models"
	"gorm.io/gorm"
	"log"
	"time"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

// Получаем структуру user по логину
func (r *Repository) GetUser(login string) (user *models.Users, err error) {
	sqlQuery := `select * from users
where users.login = ?;`
	err = r.db.Raw(sqlQuery, login).Scan(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Вовращаюсь Юзера если такой логин Существует в базе данных (Не при регистрации!)
func (r *Repository) ValidateLogin(login string) (user *models.Users, err error) {
	err = r.db.Where("login = ?", login).First(&user).Error
	if err != nil {
		log.Println("Логин Не прошел Валидацию!")
		return nil, err
	}
	return user, nil
}

//____________________________________________________________________________________

// Вернёт true если Этот Логин свободен, в противном случае - false
func (r *Repository) IsLoginFree(login string) bool {
	var user models.Users
	amountOfChar := r.db.Where("login = ?", login).First(&user).RowsAffected
	if amountOfChar != 0 {
		return false
	}
	return true
}

// Регистрирует Ползователя В БД
func (r *Repository) RegistrateUserToDB(user *models.Users, hash []byte) error {
	sqlQuery := `insert into users(name, login, password, created_at, tokens)
				 values (?, ?, ?, ?, ?);`
	err := r.db.Exec(sqlQuery, user.Name, user.Login, hash, time.Now(), "No Token").Error
	if err != nil {
		return err
	}
	return nil
}

// Добавляет Пользователю Его токен
func (r *Repository) AddTokenToDB(userId int, token string) error {
	sqlQuery := `update users
set tokens = ?
where id = ?;`
	err := r.db.Exec(sqlQuery, token, userId).Error
	if err != nil {
		return err
	}
	return nil
}
