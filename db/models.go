package db

type Product struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Description string
}

type User struct {
	ID       uint   `gorm:"unique;autoIncrement"`
	Username string `gorm:"primaryKey"`
	Email    string
	Password string
}

type Result struct {
	ID uint `json:"id"`
}
