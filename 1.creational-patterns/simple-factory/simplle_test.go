package simple_factory

import (
	"reflect"
	"testing"
)

func TestFactory_CreateProduct(t *testing.T) {
	type args struct {
		productType string
	}
	tests := []struct {
		name string
		args args
		want AbstractProduct
	}{
		{name: "test productA", args: args{"ProductA"}, want: &ProductA{}},
		{name: "test productA", args: args{"ProductB"}, want: &ProductB{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Factory{}
			if got := f.CreateProduct(tt.args.productType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
