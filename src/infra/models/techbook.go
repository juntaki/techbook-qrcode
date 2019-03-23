package models

type Techbook struct {
	Id         int    `db:"id"`
	CurrentUrl string `db:"current_url"`
}
