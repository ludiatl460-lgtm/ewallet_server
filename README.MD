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


