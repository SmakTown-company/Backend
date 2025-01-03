package models

type EmailDetails struct {
	OrderID  string    `json:"order_id,omitempty" `
	Status   string    `json:"status,omitempty"`
	Time     string    `json:"time, omitempty"`
	Address  string    `json:"address,omitempty"`
	Products []Product `json:"products,omitempty"`
}

type Product struct {
	Name     string `json:"name,omitempty"`
	Quantity int    `json:"quantity,omitempty"`
}

type EmailData struct {
	Subject  string        `json:"subject,omitempty"`
	BodyText string        `json:"body_text,omitempty"`
	Code     string        `json:"code,omitempty"`
	Details  *EmailDetails `json:"details,omitempty"`
}
