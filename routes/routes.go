package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"restapigo/handlers"
)

func SetupRoutes(db *sql.DB) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/brand", handlers.CreateBrandHandler(db)).Methods("POST")
	r.HandleFunc("/customer", handlers.CreateCustomerHandler(db)).Methods("POST")
	r.HandleFunc("/voucher", handlers.CreateVoucherHandler(db)).Methods("POST")
	r.HandleFunc("/voucher", handlers.GetVoucherHandler(db)).Methods("GET")
	r.HandleFunc("/voucher/brand", handlers.GetVouchersByBrandHandler(db)).Methods("GET")
	r.HandleFunc("/transaction/redemption", handlers.CreateRedemptionHandler(db)).Methods("POST")
	r.HandleFunc("/transaction/redemption", handlers.GetRedemptionDetailsHandler(db)).Methods("GET")

	return r
}
