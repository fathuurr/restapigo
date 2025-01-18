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

func TestCreateRedemptionHandler(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	handler := handlers.CreateRedemptionHandler(db)

	t.Run("Valid Request", func(t *testing.T) {
		body := `{
			"customer_id": 1,
			"voucher_items": [
				{"voucher_id": 1, "quantity": 2}
			]
		}`
		req := httptest.NewRequest("POST", "/transaction/redemption", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rr.Code)
		}

		var redemption models.RedemptionResponse
		json.Unmarshal(rr.Body.Bytes(), &redemption)

		if redemption.CustomerID != 1 {
			t.Errorf("Expected customer_id 1, got %v", redemption.CustomerID)
		}
	})
}
