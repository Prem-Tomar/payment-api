package dto

type CreatePaymentIntentRequest struct {
	Amount        int64  `json:"amount" binding:"required"`
	Currency      string `json:"currency" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	MerchantRef   string `json:"merchant_ref" binding:"required"`
	CustomerRef   string `json:"customer_ref,omitempty"`
}

type CreatePaymentIntentResponse struct {
	Status string `json:"status"`
}
