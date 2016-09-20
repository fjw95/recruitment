package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

/*type QueueStruct struct {
	arrayInterface []interface{}
}*/

var (
	arrayInterface []interface{}
)

func New(size int) Queue {
	var q Queue
	return q
}

/*func (q QueueStruct) new(size int) Queue {
	arrayInterface = make([]interface{}, size)
	return q.queue
}
*/
func Push(key interface{}) {
	arrayInterface = append(arrayInterface, key)
}

func Pop() interface{} {
	lastIndex := len(arrayInterface) - 1
	popValue := arrayInterface[lastIndex]
	arrayInterface = arrayInterface[:lastIndex-1]
	return popValue
}

func Contains(key interface{}) bool {
	for _, v := range arrayInterface {
		if v == key {
			return true
		}
	}
	return false
}

func Len() int {
	return len(arrayInterface)
}

func Keys() []interface{} {
	return arrayInterface[:len(arrayInterface)-1]
}
