package models

type GenerateOrder struct {
	Email string `json:"email" binding:"required"`
}

type CreatedOrder struct {
	OrderID string `json:"id"`
}

type SuccessfulPayment struct {
	Razorpay_payment_id string `json:"razorpay_payment_id"`
	Razorpay_order_id   string `json:"razorpay_order_id"`
	Razorpay_signature  string `json:"razorpay_signature"`
}
