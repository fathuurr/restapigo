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

func TestCreateVoucherHandler(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	handler := handlers.CreateVoucherHandler(db)

	t.Run("Valid Request", func(t *testing.T) {
		body := `{
			"brand_id": 1,
			"code": "DISC10",
			"name": "10% Discount",
			"point_cost": 200,
			"valid_from": "2025-01-01 00:00:00",
			"valid_until": "2025-12-31 23:59:59",
			"stock": 200
		}`
		req := httptest.NewRequest("POST", "/voucher", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rr.Code)
		}

		var voucher models.Voucher
		json.Unmarshal(rr.Body.Bytes(), &voucher)

		if voucher.Code != "DISC10" {
			t.Errorf("Expected voucher code 'DISC10', got %v", voucher.Code)
		}
	})
}
