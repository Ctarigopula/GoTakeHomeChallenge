package payloads

import (
	"net/http"
)

type MessagesMarkDeleted struct {
	IDs         []int  `json:"ids"`
	DeletedWhen string `json:"deletedWhen"`
}

func (m MessagesMarkDeleted) Bind(r *http.Request) error {
	return nil
}
