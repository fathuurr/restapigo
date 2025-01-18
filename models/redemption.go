package models

type RedemptionRequest struct {
	CustomerID   int                  `json:"customer_id"`
	VoucherItems []VoucherItemRequest `json:"voucher_items"`
}

type VoucherItemRequest struct {
	VoucherID int `json:"voucher_id"`
	Quantity  int `json:"quantity"`
}

type RedemptionResponse struct {
	ID              int              `json:"id"`
	CustomerID      int              `json:"customer_id"`
	TotalPoints     int              `json:"total_points"`
	Status          string           `json:"status"`
	RedemptionItems []RedemptionItem `json:"items"`
}

type RedemptionItem struct {
	VoucherID     int `json:"voucher_id"`
	Quantity      int `json:"quantity"`
	PointsPerUnit int `json:"points_per_unit"`
	TotalPoints   int `json:"total_points"`
}
