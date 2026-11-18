package models

type Dispute struct {
	Id              int64            `json:"id" db:"id"`
	DisputeRef      string           `json:"dispute_ref" db:"dispute_ref"`
	TransactionId   int64            `json:"transaction_id" db:"transaction_id"`
	DisputeType     string           `json:"dispute_type" db:"dispute_type"`
	Description     string           `json:"description" db:"description"`
	Status          string           `json:"status" db:"status"`
	ResolutionNote  *string          `json:"resolution_note" db:"resolution_note"`
	AssignedBy      *string          `json:"assigned_by" db:"assigned_by"`
	RefundAmount    *float64         `json:"refund_amount" db:"refund_amount"`
	PrincipleImages []Principle_Image `json:"principle_images" db:"-"`
}
