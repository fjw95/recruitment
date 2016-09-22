package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type QueueStruct struct {
	Value []interface{}
	Size  int
}

func New(size int) Queue {
	var queStr QueueStruct
	queStr.Value = make([]interface{}, size)
	queStr.Size = size
	var q Queue = queStr
	return q
}

func (q QueueStruct) Push(key interface{}) {
	if q.Value[0] == nil {
		q.Value[0] = key
		return
	} else {
		if q.Value[q.Size-1] == nil {
			for i := q.Len() - 1; i >= 0; i-- {
				if q.Value[i] != nil {
					q.Value[i+1] = q.Value[i]
				}
			}

		} else {
			for i := q.Len() - 2; i >= 0; i-- {
				if q.Value[i] != nil {
					q.Value[i+1] = q.Value[i]
				}
			}
		}
	}
	q.Value[0] = key
}

func (q QueueStruct) Pop() interface{} {
	var popItem interface{}

	for i := q.Len() - 1; i >= 0; i-- {
		popItem, q.Value[i] = q.Value[i], nil
	}
	return popItem
}

func (q QueueStruct) Contains(key interface{}) bool {
	for _, v := range q.Value {
		if v == key {
			return true
		}
	}
	return false
}

func (q QueueStruct) Len() int {
	var filledSlots int

	for _, k := range q.Value {
		if k == nil {
			return filledSlots
		} else {
			filledSlots++
		}
	}

	return filledSlots
}

func (q QueueStruct) Keys() []interface{} {
	return q.Value
}
