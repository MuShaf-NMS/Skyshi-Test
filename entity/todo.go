package entity

import "time"

type Todo struct {
	ID              int       `gorm:"column:todo_id;primarykey" json:"id"`
	ActivityGroupID int       `gorm:"column:activity_group_id" json:"activity_group_id"`
	Title           string    `gorm:"type:varchar(100)" json:"title"`
	Priority        string    `gorm:"type:varchar(30);default:very-high" json:"priority"`
	IsActive        bool      `gorm:"type:bool" json:"is_active"`
	Status          string    `gorm:"type:varchar(100)" json:"status"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
