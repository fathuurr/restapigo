-- +goose Up
-- +goose StatementBegin
CREATE TABLE redemptions (
     id INT AUTO_INCREMENT PRIMARY KEY,
     customer_id INT NOT NULL,
     total_points_used INT NOT NULL,
     status VARCHAR(50) NOT NULL DEFAULT 'pending',
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
     CONSTRAINT fk_redemption_customer FOREIGN KEY (customer_id) REFERENCES customers(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE redemptions;
-- +goose StatementEnd
