package models

type Voucher struct {
	ID          int    `json:"id"`
	BrandID     int    `json:"brand_id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	PointCost   int    `json:"point_cost"`
	ValidFrom   string `json:"valid_from"`
	ValidUntil  string `json:"valid_until"`
	Stock       int    `json:"stock"`
}
