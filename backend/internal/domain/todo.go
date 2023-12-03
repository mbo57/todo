package domain

import (
	"time"
)

type Todo struct {
	Id          string    `gorm:"column:id;primaryKey;autoincrement"`
	Title       string    `gorm:"column:title;not null"`
	Description string    `gorm:"column:description"`
	Deadline    time.Time `gorm:"column:deadline"`
	Priority    string    `gorm:"column:priority"`
	Category    string    `gorm:"column:category"`
	Created     time.Time `gorm:"column:created"`
	Updated     time.Time `gorm:"column:updated"`
	Status      string    `gorm:"column:status"`
	Assignee    string    `gorm:"column:assignee"`
}

type Todos []Todo
