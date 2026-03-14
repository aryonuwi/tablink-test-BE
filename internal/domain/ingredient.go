package domain

import "time"

type Ingredient struct {
	UUID        string     `json:"uuid"`
	Name        string     `json:"name"`
	CauseAlergy bool       `json:"cause_alergy"`
	Type        int        `json:"type"`
	Status      int        `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
