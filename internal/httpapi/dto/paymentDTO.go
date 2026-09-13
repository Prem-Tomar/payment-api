package dto

type CreatePaymentIntentRequest struct {
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	PaymentMethod string `json:"payment_method"`
	MerchantRef   string `json:"merchant_ref,omitempty"`
	CustomerRef   string `json:"customer_ref,omitempty"`
}
