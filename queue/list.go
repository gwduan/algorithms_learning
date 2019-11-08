package queue

type element struct {
	next  *element
	value interface{}
}

type ListQueue struct {
	head   *element
	tail   *element
	length int
}

func NewListQueue() *ListQueue {
	return &ListQueue{}
}

func (q *ListQueue) Size() int {
	return q.length
}

func (q *ListQueue) IsEmpty() bool {
	return q.length == 0
}

func (q *ListQueue) EnQueue(e interface{}) (v interface{}, err error) {
	newElement := &element{value: e}
	if q.length == 0 {
		q.tail = newElement
		q.head = newElement
	} else {
		q.tail.next = newElement
		q.tail = q.tail.next
	}
	q.length++

	return e, nil
}

func (q *ListQueue) DeQueue() (e interface{}, err error) {
	if q.length == 0 {
		return nil, ErrEmpty
	}

	headElement := q.head
	q.head = q.head.next
	q.length--
	if q.length == 0 {
		q.tail = q.head
	}

	headElement.next = nil

	return headElement.value, nil
}

func (q *ListQueue) Head() (e interface{}, err error) {
	if q.length == 0 {
		return nil, ErrEmpty
	}

	return q.head.value, nil
}

func (q *ListQueue) Tail() (e interface{}, err error) {
	if q.length == 0 {
		return nil, ErrEmpty
	}

	return q.tail.value, nil
}
