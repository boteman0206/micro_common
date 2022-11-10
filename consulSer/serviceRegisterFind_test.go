package consulSer

import (
	"fmt"
	"micro_product/micro_common/utils"
	"testing"
)

func TestConsul_Register(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterServer()
		})
	}
}

func TestConsul_FindSer(t *testing.T) {
	type args struct {
		SerName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				SerName: "micro_product",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			services := GetConsulServices(tt.args.SerName)
			fmt.Println(utils.JsonToString(services))
		})
	}
}
