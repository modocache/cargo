package queues

type Queue struct {
	elements []interface{}
}

func NewQueue() *Queue {
	return &Queue{elements: make([]interface{}, 0)}
}

func (queue *Queue) IsEmpty() bool {
	return len(queue.elements) == 0
}

func (queue *Queue) Push(element interface{}) {
	queue.elements = append(queue.elements, element)
}

func (queue *Queue) Pop() interface{} {
	element := queue.elements[0]
	queue.elements = queue.elements[1:]
	return element
}
