package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"restapigo/models"
)

func CreateBrandHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var brand models.Brand
		if err := json.NewDecoder(r.Body).Decode(&brand); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validateBrand(brand); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		query := `
            INSERT INTO brands (name, description)
            VALUES (?, ?)`

		result, err := db.Exec(query, brand.Name, brand.Description)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		brand.ID = int(lastId)
		json.NewEncoder(w).Encode(brand)
	}
}

func validateBrand(brand models.Brand) error {
	if brand.Name == "" {
		return fmt.Errorf("brand name is required")
	}
	return nil
}
