package factory_method

import (
	"reflect"
	"testing"
)

func TestFactoryA_CreateProduct(t *testing.T) {
	tests := []struct {
		name string
		want AbstractProduct
	}{
		{name: "create productA", want: &ProductA{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FactoryA{}
			if got := f.CreateProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactoryB_CreateProduct(t *testing.T) {
	tests := []struct {
		name string
		want AbstractProduct
	}{
		{name: "create productB", want: &ProductB{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FactoryB{}
			if got := f.CreateProduct(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

