package pkg

import "errors"

var (
	ErrQueueInitialization = errors.New("queue has not been initialized")
	ErrEmptyQueue          = errors.New("queue is empty")
)

type Executable interface {
	Execute() error
}

//go:generate mockgen -destination=./mocks/queue.go -package=mocks -source=queue.go
type IQueue interface {
	Push(e Executable)
	Pull() (Executable, error)
}

type Queue struct {
	list []Executable
}

func NewQueue() *Queue {
	return &Queue{
		list: make([]Executable, 0),
	}
}

func (q *Queue) Push(e Executable) {
	q.list = append(q.list, e)
}

func (q *Queue) Pull() (Executable, error) {
	if q.list == nil {
		return nil, ErrQueueInitialization
	}

	if len(q.list) == 0 {
		return nil, ErrEmptyQueue
	}

	e := q.list[0]
	q.list = q.list[1:len(q.list)]

	return e, nil
}
