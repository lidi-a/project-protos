package task

type Task struct {
	ID     string `gorm:"primarykey"`
	Title   string 
	IsDone bool
	UserID string `gorm:"user_id"`
}