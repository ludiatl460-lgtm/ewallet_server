package models

import "time"

type Kyc struct {
	Id             int64       `json:"id" db:"id"`
	KycNo          string      `json:"kyc_no" db:"kyc_no"`
	UserId         int64       `json:"user_id" db:"user_id"`
	IssueDate      *string     `json:"issue_date" db:"issue_date"`
	ExpiryDate     *string     `json:"expiry_date" db:"expiry_date"`
	Status         string      `json:"status" db:"status"`
	IsActive       bool        `json:"is_active" db:"is_active"`
	RejectedReason *string     `json:"rejected_reason" db:"rejected_reason"`
	ManageBy       *string     `json:"manage_by" db:"manage_by"`
	ManageAt       *time.Time  `json:"manage_at" db:"manage_at"`
	Address        *string     `json:"address" db:"address"`
	DocumentType   string      `json:"document_type" db:"document_type"`
	DocumentNo     string      `json:"document_no" db:"document_no"`
	DocumentIssue  string      `json:"document_issue" db:"document_issue"`
	DocumentExpiry string      `json:"document_expiry" db:"document_expiry"`
	Details        []KycDetail `json:"details" db:"-"`
}