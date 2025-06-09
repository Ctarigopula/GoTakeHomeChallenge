package models

import (
	"net/http"
	"time"
)

type Message struct {
	ID             int        `gorm:"column:id" json:"id"`
	Created        *time.Time `gorm:"column:created" json:"created"`
	DeletedAt      *time.Time `gorm:"column:deleted_at" json:"deletedAt"`
	DeletedAtOrder int        `gorm:"column:deleted_at_order" json:"deletedAtOrder"`
	UserIDs        IntList    `gorm:"column:user_ids" json:"userIDs"`
	Metadata       Metadata   `gorm:"column:metadata" json:"metadata"`
}

// TableName overrides the default table name used by GORM.
// It explicitly sets the table name for the Message model to "client_messages".
func (m *Message) TableName() string {
	return "client_messages"
}

func (m *Message) Bind(r *http.Request) error {
	return nil
}
