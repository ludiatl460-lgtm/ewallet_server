<!-- system_audit_log -->
CREATE TYPE audit_method AS ENUM ('POST', 'PUT');

CREATE TABLE system_audit_log (
    id BIGSERIAL PRIMARY KEY,
    admin_id BIGINT,                      -- ລະຫັດແອັດມິນ (ອ້າງອີງຈາກ user.id ແຕ່ບໍ່ມີ FK Constraint)
    action VARCHAR(30),                   -- ການດໍາເນີນການ (ເຊັ່ນ: CREATE, UPDATE, DELETE)
    method audit_method,                  -- ວິທີ HTTP Method (POST ຫຼື PUT)
    target_table VARCHAR(50),             -- ຊື່ຕາຕະລາງທີ່ມີການຈັດການ
    target_table_row_id BIGINT,           -- ລະຫັດ ID ຂອງແຖວໃນຕາຕະລາງນັ້ນ
    old_value JSONB,                       -- ຂໍ້ມູນເກົ່າກ່ອນແກ້ໄຂ (ເກັບເປັນ JSON)
    new_value JSONB,                       -- ຂໍ້ມູນໃໝ່ຫຼັງແກ້ໄຂ (ເກັບເປັນ JSON)
    ip_address VARCHAR(15),               -- ເລກ IP ຂອງຜູ້ເຮັດລາຍການ
    agent TEXT,                           -- ຂໍ້ມູນ User Agent ຂອງ Client
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP -- ເວລາທີ່ສ້າງ Log
);



<!-- OTP_LOG -->
-- ສ້າງຕາຕະລາງ otp_log
CREATE TABLE IF NOT EXISTS otp_log (
    id              BIGSERIAL PRIMARY KEY,
    
    user_id         BIGINT       NOT NULL,                       
    
    otp             VARCHAR(255)  NOT NULL,                       
    
    start_at        TIMESTAMPTZ  NOT NULL DEFAULT NOW(),         
    end_at          TIMESTAMPTZ  NOT NULL,                       
    
    is_confirm      BOOLEAN      NOT NULL DEFAULT FALSE,         
    
    incorrect_count INT          NOT NULL DEFAULT 0,             
    
    ip_address      VARCHAR(45),                                 

    CONSTRAINT chk_otp_log_time      CHECK (end_at > start_at),
    CONSTRAINT chk_otp_log_incorrect CHECK (incorrect_count >= 0)
);

-- ສ້າງ Index ເພື່ອຄົ້ນຫາ "OTP ທີ່ຍັງໃຊ້ໄດ້ຂອງ user" ໄດ້ໄວທີ່ສຸດ
-- (ແກ້ໄຂ: ເອົາ type ອອກ)
CREATE INDEX IF NOT EXISTS idx_otp_log_user_active
    ON otp_log (user_id, end_at DESC)
    WHERE is_confirm = FALSE;

-- ສ້າງ Index ສຳລັບການລຶບ OTP ທີ່ໝົດອາຍຸແລ້ວ
CREATE INDEX IF NOT EXISTS idx_otp_log_expires_at 
    ON otp_log(end_at);



<!-- QR_CODE TABLE -->
-- 1. ສ້າງ Enum Type ສຳລັບປະເພດ QR (Static ຫຼື Dynamic)
CREATE TYPE qr_type_enum AS ENUM ('static', 'dynamic');

CREATE TABLE IF NOT EXISTS qr_code (
    id              BIGSERIAL PRIMARY KEY,
    
    account_id      BIGINT NOT NULL,                 
    
    qr_type         qr_type_enum NOT NULL,           
    
    qr_payload      VARCHAR(255) UNIQUE NOT NULL,    
    
    amount          FLOAT,                           
    
    expires_at      TIMESTAMPTZ,                
    CONSTRAINT chk_qr_logic CHECK (
        (qr_type = 'static') OR 
        (qr_type = 'dynamic' AND amount > 0 AND expires_at IS NOT NULL)
    )
);

CREATE INDEX IF NOT EXISTS idx_qr_code_account_id ON qr_code(account_id);

CREATE INDEX IF NOT EXISTS idx_qr_code_payload ON qr_code(qr_payload);



<!-- NOTIFICATION_TABLE -->
-- 1. ສ້າງ Enum ສຳລັບ Channel ແລະ Status
CREATE TYPE notification_channel AS ENUM ('SMS', 'email');
CREATE TYPE notification_status AS ENUM ('queued', 'sent', 'delivered', 'failed');

CREATE TABLE IF NOT EXISTS notification (
    id          BIGSERIAL PRIMARY KEY,
    
    user_id     BIGINT NOT NULL,                 
    
    channel     notification_channel NOT NULL,   
    
    title       VARCHAR(50),                     
    message     TEXT NOT NULL,                            
    
    status      notification_status NOT NULL DEFAULT 'queued', 
    
    is_read     BOOLEAN NOT NULL DEFAULT FALSE,  
    read_at     TIMESTAMPTZ,                     
    
    sent_at     TIMESTAMPTZ                      

    -- CONSTRAINT chk_read_time CHECK (is_read = FALSE OR read_at IS NOT NULL)
);

-- 3. ສ້າງ Index ເພື່ອໃຫ້ການດຶງແຈ້ງເຕືອນຂອງຜູ້ໃຊ້ໄວຂຶ້ນ
CREATE INDEX IF NOT EXISTS idx_notification_user_id ON notification(user_id);
CREATE INDEX IF NOT EXISTS idx_notification_is_read ON notification(is_read);




<!-- ACCOUNT SCRIPTS -->
CREATE TYPE account_status_enum AS ENUM ('active', 'frozen', 'closed');

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

-- 3. ສ້າງ Index ເພື່ອໃຫ້ການຄົ້ນຫາໄວຂຶ້ນ
-- (account_number ມີ UNIQUE ແລ້ວ ຈຶ່ງບໍ່ຕ້ອງສ້າງ Index ຊ້ຳ)
CREATE INDEX IF NOT EXISTS idx_account_user_id ON account(user_id);
CREATE INDEX IF NOT EXISTS idx_account_status ON account(status);