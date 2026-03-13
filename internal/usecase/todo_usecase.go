package usecase

import "test-go/internal/domain"

type TodoUsecase interface {
	CreateTodo(todo *domain.Todo) error
	FindAllTodos() ([]domain.Todo, error)
	FindTodoByID(id uint) (*domain.Todo, error)
	UpdateTodo(todo *domain.Todo) error
	DeleteTodo(id uint) error
}

type todoUsecase struct {
	repo domain.TodoRepository
}

func NewTodoUsecase(repo domain.TodoRepository) TodoUsecase {
	return &todoUsecase{repo}
}

func (u *todoUsecase) CreateTodo(todo *domain.Todo) error {
	return u.repo.Create(todo)
}

func (u *todoUsecase) FindAllTodos() ([]domain.Todo, error) {
	return u.repo.FindAll()
}

func (u *todoUsecase) FindTodoByID(id uint) (*domain.Todo, error) {
	return u.repo.FindByID(id)
}

func (u *todoUsecase) UpdateTodo(todo *domain.Todo) error{
	return u.repo.Update(todo)
}	

func (u *todoUsecase) DeleteTodo(id uint) error {
	return u.repo.Delete(id)
}
