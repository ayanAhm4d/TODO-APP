package todo

type Todo struct {
	ID           int
	Description  string
	IsComplelted bool
}

func NewTodo(description string) *Todo {
	return &Todo{
		Description:  description,
		IsComplelted: false,
	}
}
func (t *Todo) MarkComplete() {
	t.IsComplelted = true
}
