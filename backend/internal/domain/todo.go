package domain

import (
	"bytes"
	"database/sql/driver"
	"time"
)

type Todo struct {
	Id          int     `gorm:"column:id;primaryKey;autoincrement"`
	Title       string  `gorm:"column:title;not null" json:"title"`
	Description string  `gorm:"column:description" json:"description"`
	Deadline    *MyTime `gorm:"column:deadline;type:datatime" json:"deadline"`
	Priority    string  `gorm:"column:priority" json:"priority"`
	Category    string  `gorm:"column:category" json:"category"`
	Created     *MyTime `gorm:"column:created;type:datatime" json:"created"`
	Updated     *MyTime `gorm:"column:updated;type:datatime" json:"updated"`
	Status      string  `gorm:"column:status" json:"status"`
	Assignee    string  `gorm:"column:assignee" json:"assignee"`
}

type MyTime struct {
	time.Time
}

func SetTime(tt time.Time) *MyTime {
	return &MyTime{tt}
}

func (t *MyTime) UnmarshalJSON(buf []byte) error {
	s := bytes.Trim(buf, `"`)
	tt, err := time.ParseInLocation(time.DateTime, string(s), time.Local)
	*t = MyTime{tt}
	return err
}

func (t MyTime) MarshalJSON() ([]byte, error) {
	str := t.Format(time.DateTime)
	return []byte(`"` + str + `"`), nil
}

func (t *MyTime) Scan(value interface{}) error {
	t.Time = value.(time.Time)
	return nil
}

func (t MyTime) Value() (driver.Value, error) {
	return t.Time, nil
}

type Todos []Todo
