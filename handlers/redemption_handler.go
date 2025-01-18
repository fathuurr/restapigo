package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"restapigo/models"
)

func CreateRedemptionHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req models.RedemptionRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tx, err := db.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tx.Rollback()

		var totalPoints int
		for _, item := range req.VoucherItems {
			var pointCost int
			err := tx.QueryRow("SELECT point_cost FROM vouchers WHERE id = ?", item.VoucherID).Scan(&pointCost)
			if err != nil {
				http.Error(w, "Invalid voucher ID", http.StatusBadRequest)
				return
			}
			totalPoints += pointCost * item.Quantity
		}

		result, err := tx.Exec(`
            INSERT INTO redemptions (customer_id, total_points_used, status)
            VALUES (?, ?, 'completed')`,
			req.CustomerID, totalPoints)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		redemptionID, err := result.LastInsertId()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, item := range req.VoucherItems {
			var pointCost int
			err := tx.QueryRow("SELECT point_cost FROM vouchers WHERE id = ?", item.VoucherID).Scan(&pointCost)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			_, err = tx.Exec(`
                INSERT INTO redemption_items (redemption_id, voucher_id, quantity, points_per_unit, total_points)
                VALUES (?, ?, ?, ?, ?)`,
				redemptionID, item.VoucherID, item.Quantity, pointCost, pointCost*item.Quantity)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		if err := tx.Commit(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := models.RedemptionResponse{
			ID:          int(redemptionID),
			CustomerID:  req.CustomerID,
			TotalPoints: totalPoints,
			Status:      "completed",
		}

		json.NewEncoder(w).Encode(response)
	}
}

func GetRedemptionDetailsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		transactionID := r.URL.Query().Get("transactionId")
		if transactionID == "" {
			http.Error(w, "transaction ID is required", http.StatusBadRequest)
			return
		}

		var response models.RedemptionResponse
		err := db.QueryRow(`
            SELECT id, customer_id, total_points_used, status
            FROM redemptions
            WHERE id = ?`,
			transactionID,
		).Scan(&response.ID, &response.CustomerID, &response.TotalPoints, &response.Status)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		rows, err := db.Query(`
            SELECT voucher_id, quantity, points_per_unit, total_points
            FROM redemption_items
            WHERE redemption_id = ?`,
			transactionID,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var item models.RedemptionItem
			err := rows.Scan(
				&item.VoucherID,
				&item.Quantity,
				&item.PointsPerUnit,
				&item.TotalPoints,
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			response.RedemptionItems = append(response.RedemptionItems, item)
		}

		json.NewEncoder(w).Encode(response)
	}
}
