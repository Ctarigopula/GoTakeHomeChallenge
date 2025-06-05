package models

import (
	"net/http"
	"time"
)

type Message struct {
	ID             int        `db:"id" json:"id"`
	Created        *time.Time `db:"sent" json:"created"`
	DeletedAt      *time.Time `db:"deleted_at" json:"deletedAt"`
	DeletedAtOrder int        `db:"deleted_at_order" json:"deletedAtOrder"`
	UserIDs        IntList    `db:"user_ids" json:"userIDs"`
	Metadata       Metadata   `db:"metadata" json:"metadata"`
}

func (m *Message) Bind(r *http.Request) error {
	return nil
}
