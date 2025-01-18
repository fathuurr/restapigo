package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"restapigo/models"
)

func CreateCustomerHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var customer models.Customer
		if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validateCustomer(customer); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		query := `
            INSERT INTO customers (name, email, point_balance)
            VALUES (?, ?, ?)`

		result, err := db.Exec(query, customer.Name, customer.Email, customer.PointBalance)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		customer.ID = int(lastId)
		json.NewEncoder(w).Encode(customer)
	}
}

func validateCustomer(customer models.Customer) error {
	if customer.Name == "" {
		return errors.New("customer name is required")
	}
	if customer.Email == "" {
		return errors.New("customer email is required")
	}
	return nil
}
