package model

import (
	yoocommon "github.com/rvinnie/yookassa-sdk-go/yookassa/common"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"time"
)

type Product struct {
	Name     string  `json:"name"`
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
	Quantity string  `json:"quantity"`
}
type Payment struct {
	ID            string                      `json:"id" db:"id"`
	UserID        string                      `json:"user_id" db:"user_id"`
	Amount        *yoocommon.Amount           `json:"amount" db:"amount"`
	Status        *yoopayment.Status          `json:"status" db:"status"`
	PaymentMethod *yoopayment.PaymentMethoder `json:"payment_method" db:"payment_method"`
	CreatedAt     *time.Time                  `json:"created_at" db:"created_at"`
	ExpiresAt     *time.Time                  `json:"expires_at" db:"expires_at"`
	ProductList   []Product                   `json:"product_list" db:"product_list"`
	URL           string                      `json:"url" db:"url"`
	Currency      string                      `json:"currency" db:"currency"`
}
