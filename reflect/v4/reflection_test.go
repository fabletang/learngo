package v4

import (
	"reflect"
	"testing"
)

type Car struct {
	Num  int
	MFRS string
}
type Person struct {
	Name string
	Nick string
	Car  Car
}

func Test_walk(t *testing.T) {
	tests := []struct {
		Name  string
		Input interface{}
		Want  []string
	}{
		{"Struct with one string field", struct{ Name string }{"a"}, []string{"a"}},
		{"Struct with two string field", struct {
			Name string
			Nick string
		}{"a", "aa"}, []string{"a", "aa"}},
		{"Struct with non string field", Car{1, "bmw"},
			[]string{"bmw"}},
		{"Nested Struct with string field", Person{"a", "aa", Car{1, "bmw"}},
			[]string{"a", "aa", "bmw"}},
		{"Struct Ptr with non string field", &Car{1, "bmw"},
			[]string{"bmw"}},
		{"Nested Struct Ptr with string field", &Person{"a", "aa", Car{1, "bmw"}},
			[]string{"a", "aa", "bmw"}},
		{"Slice with struct",
			[]Car{
				{666, "bmw"},
				{999, "byd"},
			},
			[]string{"bmw", "byd"}},
		{"Array with struct",
			[2]Car{
				{666, "bmw"},
				{999, "byd"},
			},
			[]string{"bmw", "byd"}},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(i string) {
				got = append(got, i)
			})
			if !reflect.DeepEqual(got, tt.Want) {
				t.Errorf("got:%v,want:%v\n", got, tt.Want)
			}
		})
	}
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Car)

		go func() {
			aChannel <- Car{33, "bmw"}
			aChannel <- Car{34, "forcus"}
			close(aChannel)
		}()

		var got []string
		want := []string{"bmw", "forcus"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
