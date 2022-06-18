package modelHelper

import "time"

type DateAuditModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type Header struct {
	Status       string `json:"status"`
	ResponseCode string `json:"response_code"`
	Message      string `json:"message"`
}
