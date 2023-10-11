package models

import "time"

type Library struct {
	ID   uint   `gorm:"primary_key; autoincrement" json:"library_id"`
	Name string `json:"library_name"`
}

type User struct {
	ID            uint   `gorm:"primary_key; autoincrement" json:"user_id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	ContactNumber string `json:"contact_number"`
	Role          string `json:"role"`
	LibID         uint   `json:"library_id"`
}

type BookInventory struct {
	ISBN            string `gorm:"primary_key" json:"isbn"`
	LibID           uint   `json:"library_id"`
	Title           string `json:"title"`
	Authors         string `json:"authors"`
	Publisher       string `json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     int    `json:"total_copies"`
	AvailableCopies int    `json:"available_copies"`
}

type RequestEvents struct {
	ReqID        uint      `gorm:"primary_key; autoincrement" json:"request_id"`
	BookISBN     string    `json:"book_isbn"`
	ReaderID     uint      `json:"reader_id"`
	RequestDate  time.Time `json:"request_date"`
	ApprovalDate time.Time `json:"approval_date"`
	ApproverID   uint      `json:"approver_id"`
	RequestType  string    `json:"request_type"`
}

type IssueRegistry struct {
	IssueID            uint      `gorm:"primary_key; autoincrement" json:"issue_id"`
	ISBN               string    `json:"isbn"`
	ReaderID           uint      `json:"reader_id"`
	IssueApproverID    uint      `json:"issue_approver_id"`
	IssueStatus        string    `json:"issue_status"`
	IssueDate          time.Time `json:"issue_date"`
	ExpectedReturnDate time.Time `json:"expected_return_date"`
	ReturnDate         time.Time `json:"return_date"`
	ReturnApproverID   uint      `json:"return_approver_id"`
}
