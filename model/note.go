package model

type Note struct {
	Id    int      `gorm:"autoIncrement;not null;uniqueIndex;primaryKey" json:"id"`
	Title string   `gorm:"not null" json:"title"`
	Tag   []string `json:"tag"`
	Note  string   `gorm:"not null" json:"note"`
}
