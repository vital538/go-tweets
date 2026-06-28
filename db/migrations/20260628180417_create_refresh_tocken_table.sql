-- migrate:up
CREATE TABLE IF NOT EXISTS refresh_tockens (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    refresh_tocken TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_user_id_refresh_tockens FOREIGN KEY (user_id) REFERENCES users(id)
)

-- migrate:down
DROP TABLE IF EXISTS refresh_tockens

