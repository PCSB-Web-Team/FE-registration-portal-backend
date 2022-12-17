package models

type GenerateOrder struct {
	Email string `json:"email" binding:"required"`
}

type CreatedOrder struct {
	OrderID string `json:"id"`
}

type SuccessfulPayment struct {
	Email               string `json:"email" binding:"required"`
	Razorpay_payment_id string `json:"razorpay_payment_id" binding:"required"`
	Razorpay_order_id   string `json:"razorpay_order_id"`
	Razorpay_signature  string `json:"razorpay_signature"`
}
