-- Таблица пользователей
CREATE TABLE users (
                       id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
                       email VARCHAR(255) UNIQUE NOT NULL,
                       login VARCHAR(100) UNIQUE NOT NULL,
                       created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                       deleted_at TIMESTAMPTZ,
                       hash VARCHAR(255) NOT NULL,
                       salt VARCHAR(255) NOT NULL
);

-- Индексы для оптимизации
CREATE INDEX idx_users_email ON users(email) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_login ON users(login) WHERE deleted_at IS NULL;

-- Триггер для автоматического обновления updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Комментарии к таблице и колонкам
COMMENT ON TABLE users IS 'Таблица для хранения данных пользователей';
COMMENT ON COLUMN users.id IS 'Уникальный идентификатор пользователя';
COMMENT ON COLUMN users.email IS 'Email пользователя (уникальный)';
COMMENT ON COLUMN users.login IS 'Логин пользователя (уникальный)';
COMMENT ON COLUMN users.created_at IS 'Дата и время создания пользователя';
COMMENT ON COLUMN users.updated_at IS 'Дата и время последнего обновления';
COMMENT ON COLUMN users.deleted_at IS 'Дата и время мягкого удаления (NULL если пользователь активен)';
COMMENT ON COLUMN users.hash IS 'Хеш пароля пользователя';
COMMENT ON COLUMN users.salt IS 'Соль для хеширования пароля';