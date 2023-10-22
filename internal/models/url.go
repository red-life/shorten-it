package models

type URL struct {
	Key  string `gorm:"uniqueIndex"`
	Long string `gorm:"uniqueIndex"`
}
