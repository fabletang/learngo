package v2

import (
	"testing"
)

//var data string
//var bts []byte

//func setup() {
//data = string(make([]byte, 500))
//bts, _ = json.Marshal(Car{name: "bmw", id: 1234, remark: data})
//}

var data = string(make([]byte, 500))
var bts, _ = json.Marshal(Car{name: "bmw", id: 1234, remark: data})

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go func() {
			Unmarshal(bts)
		}()
	}
}
func BenchmarkUnmarshalGobel(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go func() {
			UnmarshalGlobal(bts)
		}()
	}
}
func BenchmarkUnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go func() {
			UnmarshalWithPool(bts)
		}()
	}
}

//:!go test -bench="^BenchmarkUnmar"  -benchmem -cpu 1,2,4
//goos: darwin
//goarch: amd64
//pkg: github.com/fabletang/learngo/syncpool/v2
//cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
//BenchmarkUnmarshal                205168              5729 ns/op             683 B/op          9 allocs/op
//BenchmarkUnmarshal-2             1000000              2565 ns/op             336 B/op          9 allocs/op
//BenchmarkUnmarshal-4             1000000              1173 ns/op             244 B/op          9 allocs/op
//BenchmarkUnmarshalGobel          1000000              1887 ns/op             273 B/op          0 allocs/op
//BenchmarkUnmarshalGobel-2        2644710               498.3 ns/op            18 B/op          0 allocs/op
//BenchmarkUnmarshalGobel-4        4686962               254.2 ns/op             0 B/op          0 allocs/op
//BenchmarkUnmarshalWithPool       1000000              1048 ns/op               0 B/op          0 allocs/op
//BenchmarkUnmarshalWithPool-2     2091230               567.7 ns/op             0 B/op          0 allocs/op
//BenchmarkUnmarshalWithPool-4     4542127               255.5 ns/op             0 B/op          0 allocs/op
//PASS
//ok      github.com/fabletang/learngo/syncpool/v2        22.070s
