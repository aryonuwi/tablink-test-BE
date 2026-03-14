package domain

import "time"

type Item struct {
	UUID      string     `json:"uuid"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Status    int        `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
