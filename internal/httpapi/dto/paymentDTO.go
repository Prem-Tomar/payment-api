package dto

type CreatePaymentIntentRequest struct {
	Amount        int64  `json:"amount" binding:"gt=0"`
	Currency      string `json:"currency" binding:"required"`
	PaymentMethod PaymentMethodType  `json:"payment_method" binding:"required,oneof=card"`
	MerchantRef   string `json:"merchant_ref" binding:"required"`
	CustomerRef   string `json:"customer_ref,omitempty"`
}

type CreatePaymentIntentResponse struct {
	Status PaymentIntentStatus `json:"status"`
	ID     string              `json:"id"`
}

type PaymentIntentStatus string

const (
	PaymentIntentStatusPlaceHolder PaymentIntentStatus = "placeholder"
)

type PaymentMethodType string

const (
	PaymentMethodCard PaymentMethodType = "card"
)
