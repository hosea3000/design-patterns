package abstract_factory

import (
	"reflect"
	"testing"
)

func TestFactoryA_CreateProductA(t *testing.T) {
	tests := []struct {
		name string
		want AbstractProductA
	}{
		{name:"", want: &FactoryAProductA{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FactoryA{}
			if got := f.CreateProductA(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProductA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactoryA_CreateProductB(t *testing.T) {
	tests := []struct {
		name string
		want AbstractProductB
	}{
		{name:"", want: &FactoryAProductB{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FactoryA{}
			if got := f.CreateProductB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProductB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactoryB_CreateProductA(t *testing.T) {
	tests := []struct {
		name string
		want AbstractProductA
	}{
		{name:"", want: &FactoryBProductA{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FactoryB{}
			if got := f.CreateProductA(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProductA() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactoryB_CreateProductB(t *testing.T) {
	tests := []struct {
		name string
		want AbstractProductB
	}{
		{name:"", want: &FactoryBProductB{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := FactoryB{}
			if got := f.CreateProductB(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProductB() = %v, want %v", got, tt.want)
			}
		})
	}
}
