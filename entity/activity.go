package entity

import "time"

type Activity struct {
	ID        int       `gorm:"column:activity_id;primarykey" json:"id"`
	Title     string    `gorm:"type:varchar(100)" json:"title"`
	Email     string    `gorm:"type:varchar(100)" json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Todos     []Todo    `gorm:"foreignKey:ActivityGroupID;references:ID" json:"-"`
}
