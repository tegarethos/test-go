package repository

import (
	"test-go/internal/domain"

	"gorm.io/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) domain.TodoRepository {
	return &todoRepository{db}
}

// Create implements [domain.TodoRepository].
func (r *todoRepository) Create(todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

// FindAll implements [domain.TodoRepository].
func (r *todoRepository) FindAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	if err := r.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// FindByID implements [domain.TodoRepository].
func (r *todoRepository) FindByID(id uint) (*domain.Todo, error) {
	var todo domain.Todo
	if err := r.db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// Update implements [domain.TodoRepository].
func (r *todoRepository) Update(todo *domain.Todo) error {
	return r.db.Save(todo).Error
}

// Delete implements [domain.TodoRepository].
func (r *todoRepository) Delete(id uint) error {
	return r.db.Delete(&domain.Todo{}, id).Error
}