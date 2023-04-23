package models

import "time"

type Stock struct {
	Stock        string    `json:"stock"`
	StockMarket  int32     `json:"market"`
	CreationDate time.Time `json:"date"`
}

const (
	B3     = 0
	NASDAQ = 1
)
