package repository

import (
	"examles_11_cohort/todo_server/internal/model"
)

//go:generate mockgen -package mocks -destination=../../mocks/mock_repository.go examles_11_cohort/todo_server/internal/repository Repository
type Repository interface {
	GetTodos() []model.Todo
	AddTodo(todo model.Todo) model.Todo
	UpdateTodo(updatedTodo model.Todo) (model.Todo, error)
	DeleteTodo(id string) error
}
