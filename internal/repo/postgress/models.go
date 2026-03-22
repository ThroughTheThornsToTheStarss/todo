package repo

type Todo struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Body      string `gorm:"type:text;not null"`
	Completed bool   `gorm:"type:boolean;default:false"`
}