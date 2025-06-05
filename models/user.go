package models

import (
	"net/http"
)

type User struct {
	ID        int      `db:"id" json:"id"`
	ClientID  int      `db:"client_id" json:"clientID"`
	FirstName string   `db:"first_name" json:"firstName"`
	LastName  string   `db:"last_name" json:"lastName"`
	Metadata  Metadata `db:"metadata" json:"metadata"`
}

func (u *User) Bind(r *http.Request) error {
	return nil
}
