package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"restapigo/models"
)

func CreateVoucherHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var voucher models.Voucher
		if err := json.NewDecoder(r.Body).Decode(&voucher); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validateVoucher(voucher); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		query := `
            INSERT INTO vouchers (brand_id, code, name, description, point_cost, valid_from, valid_until, stock)
            VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

		result, err := db.Exec(
			query,
			voucher.BrandID,
			voucher.Code,
			voucher.Name,
			voucher.Description,
			voucher.PointCost,
			voucher.ValidFrom,
			voucher.ValidUntil,
			voucher.Stock,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		lastId, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		voucher.ID = int(lastId)
		json.NewEncoder(w).Encode(voucher)
	}
}

func GetVoucherHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		voucherID := r.URL.Query().Get("id")
		if voucherID == "" {
			http.Error(w, "voucher ID is required", http.StatusBadRequest)
			return
		}

		var voucher models.Voucher
		query := `
            SELECT id, brand_id, code, name, description, point_cost, valid_from, valid_until, stock
            FROM vouchers
            WHERE id = ?`

		err := db.QueryRow(query, voucherID).Scan(
			&voucher.ID,
			&voucher.BrandID,
			&voucher.Code,
			&voucher.Name,
			&voucher.Description,
			&voucher.PointCost,
			&voucher.ValidFrom,
			&voucher.ValidUntil,
			&voucher.Stock,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(voucher)
	}
}

func GetVouchersByBrandHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		brandID := r.URL.Query().Get("id")
		if brandID == "" {
			http.Error(w, "brand ID is required", http.StatusBadRequest)
			return
		}

		rows, err := db.Query(`
            SELECT id, brand_id, code, name, description, point_cost, valid_from, valid_until, stock
            FROM vouchers
            WHERE brand_id = ?`, brandID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		vouchers := []models.Voucher{}
		for rows.Next() {
			var v models.Voucher
			err := rows.Scan(
				&v.ID,
				&v.BrandID,
				&v.Code,
				&v.Name,
				&v.Description,
				&v.PointCost,
				&v.ValidFrom,
				&v.ValidUntil,
				&v.Stock,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			vouchers = append(vouchers, v)
		}

		json.NewEncoder(w).Encode(vouchers)
	}
}

func validateVoucher(voucher models.Voucher) error {
	if voucher.Code == "" {
		return errors.New("voucher code is required")
	}
	if voucher.Name == "" {
		return errors.New("voucher name is required")
	}
	if voucher.PointCost < 0 {
		return errors.New("point_cost must be non-negative")
	}
	if voucher.ValidFrom == "" || voucher.ValidUntil == "" {
		return errors.New("valid_from and valid_until are required")
	}
	return nil
}
