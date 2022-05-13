package main

import (
	"fmt"
	"reflect"
	"testing"
)

func initStr() []string {
	return []string{
		"dawudahiwd 5",
		"qawdasdj",
		"wdqweqweq",
		"bw",
		"bw",
		"aaaaaaa 992352343 abababab abababab",
		"aaaaaaa 662352343 abababab abababab",
		"aaaaaaa 652352343 abababab abababab",
		"awawaw 123124123123 wdqweqweq",
		"123124123123",
		"23423423523",
	}
}

func Test_sortStr(t *testing.T) {
	type args struct {
		column int
		str    []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok",
			args: args{
				column: 1,
				str:    initStr(),
			},
			want: []string{
				"123124123123",
				"23423423523",
				"aaaaaaa 652352343 abababab abababab",
				"aaaaaaa 652352343 abababab abababab",
				"aaaaaaa 652352343 abababab abababab",
				"awawaw 123124123123 wdqweqweq",
				"bw",
				"bw",
				"dawudahiwd 5",
				"qawdasdj",
				"wdqweqweq",
			},
		},
		{
			name: "third column",
			args: args{
				column: 3,
				str:    initStr(),
			},
			want: []string{
				"123124123123",
				"23423423523",
				"bw",
				"bw",
				"dawudahiwd 5",
				"qawdasdj",
				"wdqweqweq",
				"aaaaaaa 652352343 abababab abababab",
				"aaaaaaa 652352343 abababab abababab",
				"aaaaaaa 652352343 abababab abababab",
				"awawaw 123124123123 wdqweqweq",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortStr(tt.args.column, tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sortInt(t *testing.T) {
	type args struct {
		column int
		str    []string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				column: 2,
				str:    initStr(),
			},
			want: []string{
				"123124123123",
				"23423423523",
				"bw",
				"bw",
				"qawdasdj",
				"wdqweqweq",
				"dawudahiwd 5",
				"aaaaaaa 652352343 abababab abababab",
				"aaaaaaa 662352343 abababab abababab",
				"aaaaaaa 992352343 abababab abababab",
				"awawaw 123124123123 wdqweqweq",
			},
			wantErr: false,
		},
		{
			name: "sort string values",
			args: args{
				column: 1,
				str:    initStr(),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sortInt(tt.args.column, tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("sortInt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok",
			args: args{in: initStr()},
			want: []string{
				"dawudahiwd 5",
				"qawdasdj",
				"wdqweqweq",
				"bw",
				"bw",
				"aaaaaaa 992352343 abababab abababab",
				"aaaaaaa 662352343 abababab abababab",
				"aaaaaaa 652352343 abababab abababab",
				"awawaw 123124123123 wdqweqweq",
				"123124123123",
				"23423423523",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverse(tt.args.in)
		})

		fmt.Println(tt.args.in)

		if reflect.DeepEqual(tt.args.in, tt.want) {
			t.Errorf("sortInt() = %v, want %v", tt.args.in, tt.want)
		}
	}
}

func Test_deleteRepeat(t *testing.T) {
	type args struct {
		in []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok",
			args: args{in: initStr()},
			want: []string{
				"dawudahiwd 5",
				"qawdasdj",
				"wdqweqweq",
				"bw",
				"aaaaaaa 992352343 abababab abababab",
				"aaaaaaa 662352343 abababab abababab",
				"aaaaaaa 652352343 abababab abababab",
				"awawaw 123124123123 wdqweqweq",
				"123124123123",
				"23423423523",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteRepeat(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteRepeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
