package main

import "testing"

func Test_unpackString(t *testing.T) {
	tests := []struct {
		name    string
		str     string
		want    string
		wantErr bool
	}{
		{
			name:    "a4bc2d5e => aaaabccddddde",
			str:     "a4bc2d5e",
			want:    "aaaabccddddde",
			wantErr: false,
		},
		{
			name:    "abcd => abcd",
			str:     "abcd",
			want:    "abcd",
			wantErr: false,
		},
		{
			name:    "45 => empty",
			str:     "45",
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty => empty",
			str:     "",
			want:    "",
			wantErr: false,
		},
		{
			name:    "with spaces",
			str:     "a4bc2 d5e",
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpackString(tt.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("unpackString() = %v, want %v", got, tt.want)
			}
		})
	}
}
