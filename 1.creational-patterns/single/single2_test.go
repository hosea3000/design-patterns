package single

import (
	"reflect"
	"testing"
)

func TestNewSingle2(t *testing.T) {
	tests := []struct {
		name string
		want *single2
	}{
		{name: "check tow instance is one", want: NewSingle2()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSingle2(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSingle2() = %v, want %v", got, tt.want)
			}
		})
	}
}
