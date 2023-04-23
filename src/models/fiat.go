package models

import "time"

type Fiat struct {
	Code         string    `json:"code"`
	CodeIn       string    `json:"codein"`
	CreationDate time.Time `json:"date"`
}
