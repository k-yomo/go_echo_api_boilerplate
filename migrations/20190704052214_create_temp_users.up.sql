CREATE TABLE temp_users (
  id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'ID',
  phone_number VARCHAR(20) NOT NULL UNIQUE COMMENT '電話番号',
  auth_code VARCHAR(255) NOT NULL COMMENT '認証コード',
  auth_key VARCHAR(255) NOT NULL COMMENT '認証キー',
  created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE auth_code_auth_key_idx (auth_code, auth_key)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '仮登録ユーザー';
