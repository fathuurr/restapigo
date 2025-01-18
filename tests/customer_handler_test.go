package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"restapigo/handlers"
	"restapigo/models"
)

func TestCreateCustomerHandler(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	handler := handlers.CreateCustomerHandler(db)

	t.Run("Valid Request", func(t *testing.T) {
		body := `{"name": "John Doe", "email": "john.doe@example.com", "point_balance": 1000}`
		req := httptest.NewRequest("POST", "/customer", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rr.Code)
		}

		var customer models.Customer
		json.Unmarshal(rr.Body.Bytes(), &customer)

		if customer.Name != "John Doe" {
			t.Errorf("Expected customer name 'John Doe', got %v", customer.Name)
		}
	})
	
}
