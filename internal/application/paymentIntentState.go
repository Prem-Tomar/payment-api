package application

type PaymentIntentState string

const (
	PaymentIntentStateCreated              PaymentIntentState = "created"
	PaymentIntentStateRequirePaymentMethod PaymentIntentState = "requires_payment_method"
	PaymentIntentStateRequiredConfirmation PaymentIntentState = "required_confirmation"
)
