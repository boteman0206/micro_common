package utils

import (
	"fmt"
	"testing"
)

func TestUtils_GetIP(t *testing.T) {
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
			ip := GetIP()
			fmt.Println("ip: ", ip)
		})
	}
}
