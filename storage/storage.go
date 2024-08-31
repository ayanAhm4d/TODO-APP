package storage

import (
	"errors"
	"todo-app/todo"
)

type Storage interface {
	Add(*todo.Todo) error
	Get(int) (*todo.Todo, error)
	Update(*todo.Todo) error
	Delete(int) error
	List() ([]*todo.Todo, error)
}

type InMemoryStorage struct {
	todos  map[int]*todo.Todo
	nextID int
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		todos:  make(map[int]*todo.Todo),
		nextID: 1,
	}
}

func (s *InMemoryStorage) Add(t *todo.Todo) error {
	t.ID = s.nextID
	s.todos[t.ID] = t
	s.nextID++
	return nil
}

func (s *InMemoryStorage) Get(id int) (*todo.Todo, error) {
	t, ok := s.todos[id]
	if !ok {
		return nil, errors.New("todo not found")
	}
	return t, nil
}

func (s *InMemoryStorage) Update(t *todo.Todo) error {
	if _, ok := s.todos[t.ID]; !ok {
		return errors.New("ID not found")
	}
	s.todos[t.ID] = t
	return nil
}

func (s *InMemoryStorage) Delete(id int) error {
	if _, ok := s.todos[id]; !ok {
		return errors.New("not found")
	}
	delete(s.todos, id)
	return nil
}

func (s *InMemoryStorage) List() ([]*todo.Todo, error) {
	todos := make([]*todo.Todo, 0, len(s.todos))
	for _, t := range s.todos {
		todos = append(todos, t)
	}
	return todos, nil
}
