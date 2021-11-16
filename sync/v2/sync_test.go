package v2

import (
	"testing"
)

func TestCounter_Inc(t *testing.T) {
	counter := NewCounter()
	tests := []struct {
		name   string
		fields *Counter
		want   int
	}{
		// TODO: Add test cases.
		{"incremeting the counter 1 times leaves it at 1", counter, 1},
		//{"incremeting the counter 2 times leaves it at 2", counter, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Counter{
				value: tt.fields.value,
				mu:    tt.fields.mu,
			}
			if got := c.Inc(); got != tt.want {
				t.Errorf("Counter.Inc() = %v, want %v", got, tt.want)
			}
		})
	}
}
