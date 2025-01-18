-- +goose Up
CREATE TABLE vouchers (
      id INT AUTO_INCREMENT PRIMARY KEY,
      brand_id INT NOT NULL,
      code VARCHAR(50) NOT NULL UNIQUE,
      name VARCHAR(255) NOT NULL,
      description TEXT,
      point_cost INT NOT NULL,
      valid_from TIMESTAMP NULL,
      valid_until TIMESTAMP NULL,
      stock INT DEFAULT 0,
      created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
      CONSTRAINT fk_voucher_brand FOREIGN KEY (brand_id) REFERENCES brands(id),
      CONSTRAINT point_cost_positive CHECK (point_cost >= 0)
);

CREATE INDEX idx_vouchers_brand_id ON vouchers(brand_id);

-- +goose Down
DROP TABLE vouchers;
