package user

type User struct {
	ID       string `gorm:"primaryKey"`
	Email    string
	Password string
}