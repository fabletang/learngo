package v2

import (
	"testing"
)

var data = string(make([]byte, 500))

var bts, _ = json.Marshal(Car{name: "bmw", id: 1234, remark: data})

//func int() {
//data = string(make([]byte, 500))
//bts, _ = json.Marshal(Car{name: "bmw", id: 1234, remark: data})
//}

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
//pkg: github.com/fabletang/learngo/syncpool/v2
//cpu: Intel(R) Core(TM) i7-4770HQ CPU @ 2.20GHz
//BenchmarkUnmarshal-8             1331906               901.8 ns/op           244 B/op          9 allocs/op
//BenchmarkUnmarshalGobel-8       10392008               123.3 ns/op             0 B/op          0 allocs/op
//BenchmarkUnmarshalWithPool-8     8561763               142.7 ns/op             0 B/op          0 allocs/op
//PASS
//ok      github.com/fabletang/learngo/syncpool/v2        4.888s
