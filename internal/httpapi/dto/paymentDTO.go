package dto

type CreatePaymentIntentRequest struct {
	Amount        int64             `json:"amount" binding:"gt=0"`
	Currency      string            `json:"currency" binding:"required"`
	PaymentMethod PaymentMethodType `json:"payment_method" binding:"required,oneof=card"`
	MerchantRef   string            `json:"merchant_ref" binding:"required,max=100"`
	CustomerRef   string            `json:"customer_ref" binding:"max=100"`
}

type CreatePaymentIntentResponse struct {
	Status string `json:"status"`
	ID     string `json:"id"`
}

type PaymentMethodType string

const (
	PaymentMethodCard PaymentMethodType = "card"
)

type GetPaymentIntentResponse struct {
	ID string `json:"id"`
}
