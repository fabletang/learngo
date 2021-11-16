package v1

import (
	"encoding/json"
	"testing"
)

var data = string(make([]byte, 500))
var bts, _ = json.Marshal(Car{name: "bmw", id: 1234, remark: data})

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Unmarshal(bts)
	}
}
func BenchmarkUnmarshalGobel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UnmarshalGlobal(bts)
	}
}
func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		UnmarshalWithPool(bts)
	}
}

//:!go test -bench="^BenchmarkUnmar"  -benchmem
//goos: darwin
//goarch: amd64
//pkg: github.com/fabletang/learngo/syncpool/v1
//cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
//BenchmarkUnmarshal-8             4610998               250.5 ns/op           216 B/op          4 allocs/op
//BenchmarkUnmarshalGobel-8        4205817               283.1 ns/op           152 B/op          2 allocs/op
//BenchmarkUnmarshalWithPool-8     3870793               311.1 ns/op           152 B/op          2 allocs/op
//PASS
//ok      github.com/fabletang/learngo/syncpool/v1        4.430s
