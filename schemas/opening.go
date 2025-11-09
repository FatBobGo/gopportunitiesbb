package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Role string
	Company string
	Location string
	Remote bool
	Link string
	Salary int64
}

// OpeningResponse defines the API response structure
type OpeningResponse struct {
    ID        uint       `json:"id" example:"1"`
    CreatedAt time.Time  `json:"created_at" example:"2025-11-09T21:02:00Z"`
    UpdatedAt time.Time  `json:"updated_at" example:"2025-11-09T21:02:00Z"`
    DeletedAt *time.Time `json:"deleted_at" example:"null"`
    Role      string     `json:"role" example:"Software Engineer"`
    Company   string     `json:"company" example:"Tech Corp"`
    Location  string     `json:"location" example:"New York"`
    Remote    bool       `json:"remote" example:"true"`
    Link      string     `json:"link" example:"https://example.com/job"`
    Salary    int64      `json:"salary" example:"100000"`
}

