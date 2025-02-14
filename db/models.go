package db

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}

type Result struct {
	ID uint `json:"id"`
}
