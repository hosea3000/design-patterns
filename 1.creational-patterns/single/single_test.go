package single

import (
	"reflect"
	"testing"
)

func TestNewSingle(t *testing.T) {
	tests := []struct {
		name string
		want *single
	}{
		{name: "check instance is one", want: NewSingle()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingle(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingle() = %v, want %v", got, tt.want)
			}
		})
	}
}