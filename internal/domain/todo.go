package domain

import "time"

type Todo struct {
	ID          uint `gorm:"primaryKey" json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool `json:"completed"`
	CompletedAt *time.Time `json:"completed_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TodoRepository interface {
	Create(todo *Todo) error
	FindAll() ([]Todo, error)
	FindByID(id uint) (*Todo, error)
	Update(todo *Todo) error
	Delete(id uint) error
}