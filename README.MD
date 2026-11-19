<!-- TABLE campaign -->
CREATE TABLE campaign (
    id BIGSERIAL PRIMARY KEY,
    promotion_title VARCHAR(50) NOT NULL,
    description VARCHAR(50),
    start_at DATE NOT NULL,
    end_at DATE NOT NULL
);

<!-- CREATE TABLE province -->
CREATE TABLE province (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL    
);

<!-- CREATE TABLE district -->
CREATE TABLE district (
    id BIGSERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL, 
    province_id BIGINT NOT NULL,     
    CONSTRAINT fk_province
        FOREIGN KEY (province_id) 
        REFERENCES province(id) 
        ON DELETE CASCADE
);

<!-- CREATE TABLE user -->
-- 1. ສ້າງ Type ຂອງ Enum ສຳລັບ gender ແລະ status/role (ຖ້າຈຳເປັນ)
CREATE TYPE user_gender AS ENUM ('male', 'female');
CREATE TYPE user_status AS ENUM ('active', 'closed', 'pending', 'suspended');
CREATE TYPE user_role AS ENUM ('admin', 'customer');

<!-- CREATE TABLE users -->
-- 2. ສ້າງຕາຕະລາງ users (
    id BIGSERIAL PRIMARY KEY,
    lao_name VARCHAR(30) NOT NULL,
    lao_lastname VARCHAR(30),
    eng_name VARCHAR(30) NOT NULL,
    eng_lastname VARCHAR(30),
    gender user_gender,
    nationality VARCHAR(100) NOT NULL,
    birth DATE NOT NULL,
    birth_address VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    contact_person VARCHAR(50) NOT NULL,
    contact_person_tel VARCHAR(13) NOT NULL,
    profile VARCHAR(100),
    email VARCHAR(100) UNIQUE NOT NULL,
    tel VARCHAR(13) UNIQUE NOT NULL,
    status user_status DEFAULT 'pending',
    role user_role DEFAULT 'customer',
    password VARCHAR(255) NOT NULL,
    pin VARCHAR(255) NOT NULL,
    last_login_at TIMESTAMP
);


<!-- CREATE TABLE device -->
CREATE TABLE device (
    id BIGSERIAL PRIMARY KEY,                     
    device_uuid VARCHAR(15) UNIQUE NOT NULL,      
    device_agent VARCHAR(100) NOT NULL,         
    push_token VARCHAR(50),                       
    is_trusted BOOLEAN DEFAULT TRUE               
);

<!-- CREATE TABLE user_device -->
-- 1. (ທາງເລືອກ) ຖ້າຕ້ອງການສ້າງ Type ສຳລັບ status enum ໄວ້ລ່ວງໜ້າ
CREATE TYPE device_status AS ENUM ('active', 'revoke');

-- 2. ສ້າງຕາຕະລາງກາງ user_device
CREATE TABLE user_device (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    device_id BIGINT NOT NULL,
    registered_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status device_status DEFAULT 'active',
    last_used_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_main_device BOOLEAN DEFAULT FALSE,
    CONSTRAINT fk_user_device_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_user_device_device FOREIGN KEY (device_id) REFERENCES device(id) ON DELETE CASCADE,
    CONSTRAINT uk_user_device_unique UNIQUE (user_id, device_id)
);


<!-- CREATE TABLE device_log -->
CREATE TABLE device_log (
    id BIGSERIAL PRIMARY KEY,                          
    user_id BIGINT NOT NULL,                           
    session_token_hash VARCHAR(50) NOT NULL,           
    ip_address VARCHAR(15),                            
    geo_location VARCHAR(30)                           
);
ALTER TABLE device_log
ADD CONSTRAINT fk_device_log_user 
    FOREIGN KEY (user_id) 
    REFERENCES "users"(id) 
    ON DELETE CASCADE;

CREATE INDEX idx_device_log_user_id ON device_log(user_id);
CREATE INDEX idx_device_log_token_hash ON device_log(session_token_hash);




<!-- CREATE TABLE card  -->
CREATE TYPE card_brand_enum AS ENUM ('master', 'visa');
CREATE TYPE card_type_enum AS ENUM ('debit', 'credit');
CREATE TYPE card_status_enum AS ENUM ('blocked', 'active', 'expired');

CREATE TABLE card (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    card_token VARCHAR(100) NOT NULL,
    card_brand card_brand_enum NOT NULL,
    holder_name VARCHAR(30) NOT NULL,
    expiry TIMESTAMP NOT NULL,
    card_type card_type_enum NOT NULL,
    status card_status_enum NOT NULL DEFAULT 'active'
);

ALTER TABLE card
ADD CONSTRAINT fk_cards_user 
    FOREIGN KEY (user_id) 
    REFERENCES "users"(id) 
    ON DELETE CASCADE;

CREATE INDEX idx_cards_user_id ON card(user_id);




<!-- CREATE TABLE login_log -->
-- 2. ສ້າງ Custom Enum Type ສຳລັບສະຖານະການ login
CREATE TYPE login_status_enum AS ENUM ('success', 'failed');

CREATE TABLE login_log (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    login_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    login_status login_status_enum NOT NULL,
    ip_address VARCHAR(15)
);

ALTER TABLE login_log
ADD CONSTRAINT fk_login_log_user 
    FOREIGN KEY (user_id) 
    REFERENCES "users"(id) 
    ON DELETE CASCADE;

CREATE INDEX idx_login_log_user_id ON login_log(user_id);
CREATE INDEX idx_login_log_login_at ON login_log(login_at);


<!-- CREATE TABLE limits -->
CREATE TABLE limits (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    min_amount FLOAT NOT NULL,
    max_amount FLOAT NOT NULL,
    fee_amount FLOAT DEFAULT 0
);

-- ເພີ່ມ Foreign Key ໃຫ້ຕາຕະລາງ limits
ALTER TABLE limits
ADD CONSTRAINT fk_limit_account 
    FOREIGN KEY (account_id) 
    REFERENCES account(id) 
    ON DELETE CASCADE;


<!-- CREATE TABLE rule -->
    CREATE TABLE rule (
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    amount FLOAT,
    unit VARCHAR(20) CHECK (unit IN ('day', 'transaction', 'month'))
);

-- ເພີ່ມ Foreign Key ໃຫ້ຕາຕະລາງ rule
ALTER TABLE rule
ADD CONSTRAINT fk_rule_account 
    FOREIGN KEY (account_id) 
    REFERENCES account(id) 
    ON DELETE CASCADE;






<!-- ACCOUNT SCRIPTS -->
<!-- CREATE TABLE IF NOT EXISTS account -->
CREATE TABLE IF NOT EXISTS account (
    id                BIGSERIAL PRIMARY KEY,
    user_id           BIGINT NOT NULL,                   
    account_number    VARCHAR(15) NOT NULL UNIQUE,         
    account_type      VARCHAR(30) NOT NULL,               
    status            account_status_enum NOT NULL DEFAULT 'active',    
    opened_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
    closed_at         TIMESTAMPTZ,                       
    balance           NUMERIC(15,2) NOT NULL DEFAULT 0.00, 
    available_balance NUMERIC(15,2) NOT NULL DEFAULT 0.00, 
    hold_amount       NUMERIC(15,2) NOT NULL DEFAULT 0.00, 
    currency          VARCHAR(3) NOT NULL DEFAULT 'LAK',

    CONSTRAINT chk_account_balance CHECK (balance >= 0),
    CONSTRAINT chk_account_available CHECK (available_balance >= 0),
    CONSTRAINT chk_account_hold CHECK (hold_amount >= 0),
    CONSTRAINT chk_available_calc CHECK (available_balance <= balance)
);

CREATE INDEX IF NOT EXISTS idx_account_user_id ON account(user_id);
CREATE INDEX IF NOT EXISTS idx_account_status ON account(status);


<!-- ALTER TABLE account -->

ALTER TABLE account
ADD CONSTRAINT fk_account_user
FOREIGN KEY (user_id)
REFERENCES users(id)
ON UPDATE CASCADE
ON DELETE RESTRICT;






<!-- account -->
package models

import (
	"time"
)

type Account struct {
	ID               int64        `db:"id" json:"id"`
	UserID           int64        `db:"user_id" json:"user_id"`
	AccountNumber    string       `db:"account_number" json:"account_number"`
	AccountType      string       `db:"account_type" json:"account_type"`
	Status           string       `db:"status" json:"status"`
	OpenedAt         time.Time    `db:"opened_at" json:"opened_at"`
	ClosedAt         time.Time `db:"closed_at" json:"closed_at"`
	Balance          float64      `db:"balance" json:"balance"`
	AvailableBalance float64      `db:"available_balance" json:"available_balance"`
	HoldAmount       float64      `db:"hold_amount" json:"hold_amount"`
	Currency         string       `db:"currency" json:"currency"`
}






package services

import (
	"backend/config"
	"backend/models"
)

func GetAllAccounts() (list []models.Account, err error) {
	err = config.DB.Select(&list, `SELECT * FROM account ORDER BY id`)
	return list, err
}

func GetAccountByID(id int64) (*models.Account, error) {
	var a models.Account
	err := config.DB.Get(&a, `SELECT * FROM account WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

func CreateAccount(a *models.Account) (*models.Account, error) {
	query := `INSERT INTO account (user_id, account_number, account_type, status, closed_at,balance, available_balance, hold_amount, currency
	          )
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)RETURNING *`
	var created models.Account
	err := config.DB.Get(&created,query,a.UserID,a.AccountNumber,a.AccountType,a.Status,a.ClosedAt,a.Balance,a.AvailableBalance,a.HoldAmount,a.Currency,
	)
	return &created, err
}

func UpdateAccount(id int64, a *models.Account) (*models.Account, error) {
	query := `UPDATE account SET user_id=$1, account_number=$2, account_type=$3, status=$4, closed_at=$5,balance=$6, available_balance=$7, hold_amount=$8, currency=$9 WHERE id=$10 RETURNING *`
	var updated models.Account
	err := config.DB.Get(&updated,query,a.UserID,a.AccountNumber,a.AccountType,a.Status,a.ClosedAt,a.Balance,a.AvailableBalance,a.HoldAmount,a.Currency,id,
	)
	return &updated, err
}

func DeleteAccount(id int64) (*models.Account, error) {
	query := `DELETE FROM account WHERE id = $1 RETURNING *`
	var deleted models.Account
	err := config.DB.Get(&deleted, query, id)
	if err != nil {
		return nil, err
	}
	return &deleted, nil
}





package controllers

import (
	"backend/models"
	"backend/services"
	"database/sql"
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllAccounts(c *fiber.Ctx) error {
	list, err := services.GetAllAccounts()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(list)
}

func GetAccountByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	res, err := services.GetAccountByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Account not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func CreateAccount(c *fiber.Ctx) error {
	var a models.Account
	if err := c.BodyParser(&a); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := services.CreateAccount(&a)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(res)
}

func UpdateAccount(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	var a models.Account
	if err := c.BodyParser(&a); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := services.UpdateAccount(id, &a)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Account not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}

func DeleteAccount(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}

	res, err := services.DeleteAccount(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Account not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(res)
}





package routes
import(
	"backend/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupAccountRoutes(route fiber.Router){
	accounts := route.Group("/accounts")

	accounts.Get("/", controllers.GetAllAccounts)
	accounts.Get("/:id", controllers.GetAccountByID)
	accounts.Post("/", controllers.CreateAccount)
	accounts.Put("/:id", controllers.UpdateAccount)
	accounts.Delete("/:id", controllers.DeleteAccount)
}
