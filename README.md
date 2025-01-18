# API Documentation

This documentation provides detailed instructions and examples for using the API developed for handling brands, customers, vouchers, and redemption transactions. It includes setup instructions, available endpoints, and sample request and response formats.

---

## Table of Contents
- [Setup Instructions](#setup-instructions)
- [Database Schema](#database-schema)
- [API Endpoints](#api-endpoints)
    - [Brands](#brands)
    - [Customers](#customers)
    - [Vouchers](#vouchers)
    - [Redemptions](#redemptions)
- [Sample Requests and Responses](#sample-requests-and-responses)
- [Testing](#testing)

---

## Setup Instructions

### Prerequisites
- Go 1.19 or higher
- MySQL 5.7 or higher
- [Goose](https://github.com/pressly/goose) for database migrations

### Steps
1. Clone the repository:
   ```bash
   git clone <repository_url>
   cd <repository_folder>
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the database:
   ```bash
   mysql -u root -p
   CREATE DATABASE golang;
   ```

4. Run migrations:
   ```bash
   goose -dir ./db/migrations mysql "your_user:your_password@tcp(your_host:your_port)/your_database" up
   ```

    - Replace `your_user` with your MySQL username.
    - Replace `your_password` with your MySQL password.
    - Replace `your_host` with your database host (e.g., `localhost`).
    - Replace `your_port` with your database port (e.g., `3306`).
    - Replace `your_database` with the name of your database (e.g., `golang`).

5. Start the server:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`.

---

## Database Schema

### Tables
1. **brands**: Stores brand information.
2. **vouchers**: Stores voucher information tied to brands.
3. **customers**: Stores customer details.
4. **redemptions**: Tracks redemption transactions by customers.
5. **redemption_items**: Tracks items in each redemption transaction.

---

## API Endpoints

### Brands
- **POST /brand**
    - Create a new brand.

### Customers
- **POST /customer**
    - Create a new customer.

### Vouchers
- **POST /voucher**
    - Create a new voucher.
- **GET /voucher?id={id}**
    - Retrieve voucher details by ID.
- **GET /voucher/brand?id={brand_id}**
    - Retrieve all vouchers for a brand.

### Redemptions
- **POST /transaction/redemption**
    - Create a redemption transaction.
- **GET /transaction/redemption?transactionId={id}**
    - Retrieve redemption transaction details by ID.

---

## Sample Requests and Responses

### 1. Create a Brand
**Endpoint:** `POST /brand`

**Request Body:**
```json
{
  "name": "Brand A",
  "description": "A description of Brand A."
}
```

**Response Body:**
```json
{
  "id": 1,
  "name": "Brand A",
  "description": "A description of Brand A."
}
```

---

### 2. Create a Customer
**Endpoint:** `POST /customer`

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "point_balance": 1000
}
```

**Response Body:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "point_balance": 1000
}
```

---

### 3. Create a Voucher
**Endpoint:** `POST /voucher`

**Request Body:**
```json
{
  "brand_id": 1,
  "code": "DISC10",
  "name": "10% Discount",
  "description": "A 10% discount voucher.",
  "point_cost": 500,
  "valid_from": "2025-01-01 00:00:00",
  "valid_until": "2025-12-31 23:59:59",
  "stock": 100
}
```

**Response Body:**
```json
{
  "id": 1,
  "brand_id": 1,
  "code": "DISC10",
  "name": "10% Discount",
  "description": "A 10% discount voucher.",
  "point_cost": 500,
  "valid_from": "2025-01-01 00:00:00",
  "valid_until": "2025-12-31 23:59:59",
  "stock": 100
}
```

---

### 4. Get Voucher by ID
**Endpoint:** `GET /voucher?id=1`

**Response Body:**
```json
{
  "id": 1,
  "brand_id": 1,
  "code": "DISC10",
  "name": "10% Discount",
  "description": "A 10% discount voucher.",
  "point_cost": 500,
  "valid_from": "2025-01-01 00:00:00",
  "valid_until": "2025-12-31 23:59:59",
  "stock": 100
}
```

---

### 5. Create a Redemption
**Endpoint:** `POST /transaction/redemption`

**Request Body:**
```json
{
  "customer_id": 1,
  "voucher_items": [
    {
      "voucher_id": 1,
      "quantity": 2
    }
  ]
}
```

**Response Body:**
```json
{
  "id": 1,
  "customer_id": 1,
  "total_points": 1000,
  "status": "completed",
  "items": [
    {
      "voucher_id": 1,
      "quantity": 2,
      "points_per_unit": 500,
      "total_points": 1000
    }
  ]
}
```

---

### 6. Get Redemption Details by ID
**Endpoint:** `GET /transaction/redemption?transactionId=1`

**Response Body:**
```json
{
  "id": 1,
  "customer_id": 1,
  "total_points": 1000,
  "status": "completed",
  "items": [
    {
      "voucher_id": 1,
      "quantity": 2,
      "points_per_unit": 500,
      "total_points": 1000
    }
  ]
}
```

---

## Testing

### Running Tests
To ensure the API functions as expected, unit tests have been implemented for all handlers. Follow these steps to run the tests:

#### Run All Tests
```bash
go test ./... -v
```

#### Run Tests in a Specific File
```bash
go test -v ./tests/brand_handler_test.go
```
Replace `brand_handler_test.go` with the specific test file you want to run.

#### Run a Specific Test Function
```bash
go test -v ./tests/brand_handler_test.go -run TestCreateBrandHandler
```
Replace `TestCreateBrandHandler` with the name of the test function you want to execute.

#### Folder Structure for Tests
Tests are located in the `tests` folder, and each handler has a corresponding test file:
```
project/
├── handlers/
│   ├── brand_handler.go
│   ├── customer_handler.go
│   ├── voucher_handler.go
│   ├── redemption_handler.go
├── models/
│   ├── brand.go
│   ├── customer.go
│   ├── voucher.go
│   ├── redemption.go
├── tests/
│   ├── brand_handler_test.go
│   ├── customer_handler_test.go
│   ├── voucher_handler_test.go
│   ├── redemption_handler_test.go
├── main.go
```

This structure helps in organizing tests effectively.

### Reset Database for Testing
Ensure your test database is clean before running tests. A helper function can be used to reset the database, like this:
```go
func resetTestDB(db *sql.DB) {
    db.Exec("DELETE FROM brands")
    db.Exec("DELETE FROM customers")
    db.Exec("DELETE FROM vouchers")
    db.Exec("DELETE FROM redemptions")
    db.Exec("DELETE FROM redemption_items")
    db.Exec("ALTER TABLE brands AUTO_INCREMENT = 1")
    db.Exec("ALTER TABLE customers AUTO_INCREMENT = 1")
    db.Exec("ALTER TABLE vouchers AUTO_INCREMENT = 1")
    db.Exec("ALTER TABLE redemptions AUTO_INCREMENT = 1")
    db.Exec("ALTER TABLE redemption_items AUTO_INCREMENT = 1")
}
```
Call this function before running each test to ensure a consistent testing environment.

---

## Notes
- All timestamps should be in ISO 8601 format (`YYYY-MM-DDTHH:MM:SSZ`).
- Ensure database migrations are applied before running the application.
- Use appropriate HTTP status codes for error handling.

---

