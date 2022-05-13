package main

import (
	"reflect"
	"testing"
)

func Test_findAnagrams(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name  string
		words []string
		want  *map[string]*[]string
	}{
		{
			name:  "ok",
			words: []string{"пяТак", "пятка", "пятка", "типок", "тяпка", "листок", "слиток", "столик", "тест"},
			want:  &map[string]*[]string{"листок": {"листок", "слиток", "столик"}, "пятак": {"пятак", "пятка", "тяпка"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findAnagrams(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
