// Package models provides db/json models
package models

import "github.com/jinzhu/gorm"

// Vehicle struct for holding ymmt
type Vehicle struct {
	dbModel gorm.Model
	Year    string `json:"year"`
	Make    string `json:"make"`
	Model   string `json:"model"`
	Trim    string `json:"trim"`
}
