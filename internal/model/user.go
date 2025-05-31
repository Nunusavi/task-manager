package model

type User struct {
	ID	   int    `db:"id" json:"id"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password,omitempty"`
	CreatedAt string `db:"created_at" json::created_at,omitempty`
}