package repository

import "strconv"

type NullableID struct {
	Value string
}

type TodoItem struct {
	ID     *NullableID // shenanigans to allow for null ID...  some other mechanism would likely be more appropriate for a "real" store.
	Text   string
	Done   bool
	UserID string
}

// TodoRepository stuff...
// ... PoCing a "repository" pattern integrated with DataLoader pattern, and demonstrating that persistence model != graphql schema model
// (hence there's a TodoItem struct here that is the repository's type for TODOs...  which is separate from the graph.model.Todo)
type TodoRepository interface {
	GetTodos() (todos []*TodoItem, err error)

	GetTodosById(ids []string) (todos []*TodoItem, errors []error)

	Upsert(items []*TodoItem) (todos []*TodoItem, errors []error)

	Delete(ids []string) (errors []error)
}

type Error string

func (e Error) Error() string { return string(e) }

const (
	ErrNotFound = Error("could not find the item you were looking for")
)

// InMemoryTodoRepository - for in-mem hackery...
type InMemoryTodoRepository struct {
	items  map[string]*TodoItem
	nextID int // Better to use GUID but for faking things out and experimenting, it's easier to type ints instead of copy-paste guids..  :-)
}

func NewInMemoryTodoRepository() TodoRepository {
	var result TodoRepository
	result = &InMemoryTodoRepository{
		items:  map[string]*TodoItem{},
		nextID: 1,
	}
	return result
}

func (r *InMemoryTodoRepository) GetTodos() (todos []*TodoItem, err error) {
	// TODO: This needs PoC "paging" support ASAP...

	// NOTE: I need to chase down what the idiom is for avoiding all this copying of value types..
	// Should this be a slicke of pointers to TodoItem instead, and just accept that the slice itself will be copied all over the place
	// (because we can't take the address of a slice?), etc etc...
	todos = make([]*TodoItem, len(r.items))
	i := 0
	for _, v := range r.items {
		todos[i] = v
		i++
	}
	return todos, err
}

func (r *InMemoryTodoRepository) GetTodosById(ids []string) (todos []*TodoItem, errors []error) {
	todos = make([]*TodoItem, len(ids))
	errors = make([]error, len(ids))

	for i, id := range ids {
		existing, ok := r.items[id]
		if ok {
			todos[i] = existing
		} else {
			errors[i] = ErrNotFound
		}
	}

	return todos, errors
}

func (r *InMemoryTodoRepository) Upsert(items []*TodoItem) (todos []*TodoItem, errors []error) {
	errors = make([]error, len(items))

	for _, item := range items {
		if item.ID == nil {
			item.ID = &NullableID{Value: strconv.Itoa(r.nextID)}
			r.nextID++
		}

		r.items[item.ID.Value] = item

		// TODO: if there were actually possibility of errors here, go ahead and populate the corresponding error
	}

	return items, errors
}

func (r *InMemoryTodoRepository) Delete(ids []string) (errors []error) {
	errors = make([]error, len(ids))

	for _, id := range ids {
		delete(r.items, id)

		// TODO: if there were actually possibility of errors here, go ahead and populate the corresponding error
	}

	return errors
}
