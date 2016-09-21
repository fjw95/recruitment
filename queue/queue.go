package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type QueueStruct struct {
	key []interface{}
}

var (
	arrayInterface []interface{}
)

func New(size int) Queue {
	var q Queue = QueueStruct{make([]interface{}, size)}
	return q
}


func (q QueueStruct)Push(key interface{}) {
	arrayInterface = append(arrayInterface, key)
}

func (q QueueStruct) Pop() interface{} {
	lastIndex := len(arrayInterface) - 1
	popValue := arrayInterface[lastIndex]
	arrayInterface = arrayInterface[:lastIndex-1]
	return popValue
}

func (q QueueStruct) Contains(key interface{}) bool {
	for _, v := range arrayInterface {
		if v == key {
			return true
		}
	}
	return false
}

func (q QueueStruct) Len() int {
	return len(arrayInterface)
}

func (q QueueStruct) Keys() []interface{} {
	return q.key
}
