package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"restapigo/handlers"
	"restapigo/models"

	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB() *sql.DB {
	db, _ := sql.Open("mysql", "root:root@tcp(localhost:8889)/golang")
	return db
}

func TestCreateBrandHandler(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	handler := handlers.CreateBrandHandler(db)

	t.Run("Valid Request", func(t *testing.T) {
		body := `{"name": "Brand A", "description": "A description"}`
		req := httptest.NewRequest("POST", "/brand", bytes.NewBuffer([]byte(body)))
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status OK, got %v", rr.Code)
		}

		var brand models.Brand
		json.Unmarshal(rr.Body.Bytes(), &brand)

		if brand.Name != "Brand A" {
			t.Errorf("Expected brand name 'Brand A', got %v", brand.Name)
		}
	})
}
