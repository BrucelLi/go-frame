CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON COLUMN users.id IS '用户唯一标识，自增主键';
COMMENT ON COLUMN users.username IS '用户名，最长50字符';
COMMENT ON COLUMN users.email IS '用户邮箱，唯一，最长100字符';
COMMENT ON COLUMN users.password_hash IS '用户密码的哈希值';
COMMENT ON COLUMN users.is_active IS '用户是否激活，默认为激活，true：激活；false: 未激活';
COMMENT ON COLUMN users.created_at IS '创建时间，默认为当前时间';
COMMENT ON COLUMN users.updated_at IS '更新时间，默认为当前时间';