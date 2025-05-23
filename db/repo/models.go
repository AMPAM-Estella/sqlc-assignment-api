// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repo

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Message struct {
	ID        string           `json:"id"`
	Thread    string           `json:"thread"`
	Sender    string           `json:"sender"`
	Content   string           `json:"content"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}

type Thread struct {
	ID        string           `json:"id"`
	ThreadID  *string          `json:"thread_id"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}
