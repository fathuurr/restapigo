-- +goose Up
CREATE TABLE redemption_items (
      id INT AUTO_INCREMENT PRIMARY KEY,
      redemption_id INT NOT NULL,
      voucher_id INT NOT NULL,
      quantity INT NOT NULL,
      points_per_unit INT NOT NULL,
      total_points INT NOT NULL,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      CONSTRAINT fk_redemption_items_redemption FOREIGN KEY (redemption_id) REFERENCES redemptions(id),
      CONSTRAINT fk_redemption_items_voucher FOREIGN KEY (voucher_id) REFERENCES vouchers(id),
      CONSTRAINT quantity_positive CHECK (quantity > 0)
);

-- Create indexes in separate statements
CREATE INDEX idx_redemption_items_redemption_id ON redemption_items(redemption_id);
CREATE INDEX idx_redemption_items_voucher_id ON redemption_items(voucher_id);

-- +goose Down
DROP TABLE redemption_items;
