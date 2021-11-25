package v1

import (
	"encoding/json"
	"sync"
)

//Car car struct
type Car struct {
	name   string //name
	id     int32  //num
	remark string //reduce
}

var car = new(Car)

//Unmarshal normal
func Unmarshal(bts []byte) (car Car, err error) {
	err = json.Unmarshal(bts, car)
	return
}

//UnmarshalGlobal define global para
func UnmarshalGlobal(bts []byte) (err error) {
	err = json.Unmarshal(bts, car)
	return
}

//CarPool the pool of car
var CarPool = sync.Pool{
	New: func() interface{} {
		return new(Car)
	},
}

//UnmarshalWithPool with sync.Pool
func UnmarshalWithPool(bts []byte) (Car, error) {
	car := CarPool.Get().(*Car)
	err := json.Unmarshal(bts, car)
	CarPool.Put(car)
	return *car, err
}
