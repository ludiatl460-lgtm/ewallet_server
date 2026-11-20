package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Date struct {
	time.Time
}

// 1. ເພີ່ມ Method Value() ເພື່ອໃຫ້ Database Driver ສามารถບັນທຶກລົງ Postgres ໄດ້
func (d Date) Value() (driver.Value, error) {
	if d.IsZero() {
		return nil, nil
	}
	return d.Time, nil
}

// 2. ເພີ່ມ Method Scan() ເພື່ອให້ນ່ານຂໍ້ມູນອອກຈາກ Database ມາໃສ່ Struct ໄດ້
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		d.Time = time.Time{}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		d.Time = v
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

type User struct {
	Id               int64      `json:"id" db:"id"`
	LaoName          string     `json:"lao_name" db:"lao_name"`
	LaoLastName      string     `json:"lao_lastname" db:"lao_lastname"`
	EngName          string     `json:"eng_name" db:"eng_name"`
	EngLastName      string     `json:"eng_lastname" db:"eng_lastname"`
	Gender           string     `json:"gender" db:"gender"`
	Nationality      string     `json:"nationality" db:"nationality"`
	Birth            Date       `json:"birth" db:"birth"`
	BirthAddress     string     `json:"birth_address" db:"birth_address"`
	Address          string     `json:"address" db:"address"`
	ContactPerson    string     `json:"contact_person" db:"contact_person"`
	ContactPersonTel string     `json:"contact_person_tel" db:"contact_person_tel"`
	Profile          string     `json:"profile" db:"profile"`
	Email            string     `json:"email" db:"email"`
	Tel              string     `json:"tel" db:"tel"`
	Status           string     `json:"status" db:"status"`
	Role             string     `json:"role" db:"role"`
	Password         string     `json:"password" db:"password"`
	Pin              string     `json:"pin" db:"pin"`
	LastLoginAt      *time.Time `json:"last_login_at" db:"last_login_at"`
}
