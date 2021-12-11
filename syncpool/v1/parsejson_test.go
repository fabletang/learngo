package v1

import (
	"encoding/json"
	"testing"
)

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
//pkg: github.com/fabletang/learngo/syncpool/v1
//cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
//BenchmarkUnmarshal                372061              3097 ns/op             669 B/op          4 allocs/op
//BenchmarkUnmarshal-2             1592846               921.3 ns/op           247 B/op          4 allocs/op
//BenchmarkUnmarshal-4             2912947               414.0 ns/op           216 B/op          4 allocs/op
//BenchmarkUnmarshalGobel          1000000              4628 ns/op             385 B/op          2 allocs/op
//BenchmarkUnmarshalGobel-2        1000000              1434 ns/op             152 B/op          2 allocs/op
//BenchmarkUnmarshalGobel-4        1453827               828.6 ns/op           152 B/op          2 allocs/op
//BenchmarkUnmarshalWithPool       1000000              3093 ns/op             152 B/op          2 allocs/op
//BenchmarkUnmarshalWithPool-2     1000000              1642 ns/op             152 B/op          2 allocs/op
//BenchmarkUnmarshalWithPool-4     1327105               943.7 ns/op           152 B/op          2 allocs/op
//PASS
//ok      github.com/fabletang/learngo/syncpool/v1        27.396s
