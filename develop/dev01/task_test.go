package main

import (
	"testing"
)

func Test_getTime(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "OK",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTime()
			if (err != nil) != tt.wantErr {
				t.Errorf("getTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				t.Errorf("getTime() = %v, want not nil", got)
			}
		})
	}
}
