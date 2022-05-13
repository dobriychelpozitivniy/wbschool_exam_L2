package main

import (
	"reflect"
	"testing"
)

func initStrings() []string {
	return []string{
		"aaaaaaaaaaaaaa",
		"bbbbbbbbbbbbb",
		"cccccccccccc",
		"wwwwwwwwwwwwwwwww",
		"qqqqqqqqqqqqqqq",
		"eeeeeeeeeeeeeeeeeee",
		"python",
		"python sefsfdues sbeifsbei dnjsf",
		"asiudhiwudsbqwdpython",
		"kek eadiajodjeafh b",
		"python sefsfdues sbeifsbei dnjsf",
		"asiudhiwudsbqwdpython",
		"KEK eadiajodjeafh b",
		"python sefsfdues sbeifsbei dnjsf",
		"asiudhiwudsbqwdpython",
		"kek eadiajodjeafh b",
		"python sefsfdues sbeifsbei dnjsf",
		"asiudhiwudsbqwdpython check",
		"kek eadiajodjeafh b",
		"python sefsfdues sbeifsbei dnjsf",
		"wkudhiqwudhqiwudhqiw",
		"KEK eadiajodjeafh b",
		"qdlwijoqwdjoqwuhe2189e",
		"wkudhiqwudhqiwudhqiw",
		"kek eadiajodjeafh b",
		"qdlwijoqwdjoqwuhe2189e",
		"wkudhiqwudhqiwudhqiw",
		"kek eadiajodjeafh b",
		"qdlwijoqwdjoqwuhe2189e",
		"wkudhiqwudhqiwudhqiw",
		"wkudhiqwudhqiwudhqiw",
		"wkudhiqwudhqiwudhqiw",
		"wkudhiqwudhqiwudhqiw",
		"wkudhiqwudhqiwudhqiw",
	}
}

func Test_grep(t *testing.T) {
	type args struct {
		s     []string
		f     string
		flags Flags
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "ok",
			args: args{
				s: initStrings(),
				f: "kek",
				flags: Flags{
					"A": 0,
					"B": 0,
					"C": 0,
					"c": 0,
					"F": false,
					"v": false,
					"i": false,
					"n": false,
				},
			},
			want: []string{
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
			},
		},
		{
			name: "ok -n",
			args: args{
				s: initStrings(),
				f: "kek",
				flags: Flags{
					"A": 0,
					"B": 0,
					"C": 0,
					"c": 0,
					"F": false,
					"v": false,
					"i": false,
					"n": true,
				},
			},
			want: []string{
				"10 ::: kek eadiajodjeafh b",
				"16 ::: kek eadiajodjeafh b",
				"19 ::: kek eadiajodjeafh b",
				"25 ::: kek eadiajodjeafh b",
				"28 ::: kek eadiajodjeafh b",
			},
		},
		{
			name: "ok -A 3",
			args: args{
				s: initStrings(),
				f: "kek",
				flags: Flags{
					"A": 3,
					"B": 0,
					"C": 0,
					"c": 0,
					"F": false,
					"v": false,
					"i": false,
					"n": false,
				},
			},
			want: []string{
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"qdlwijoqwdjoqwuhe2189e",
				"wkudhiqwudhqiwudhqiw",
				"wkudhiqwudhqiwudhqiw",
			},
		},
		{
			name: "ok -B 3",
			args: args{
				s: initStrings(),
				f: "kek",
				flags: Flags{
					"A": 0,
					"B": 3,
					"C": 0,
					"c": 0,
					"F": false,
					"v": false,
					"i": false,
					"n": false,
				},
			},
			want: []string{
				"python",
				"python sefsfdues sbeifsbei dnjsf",
				"asiudhiwudsbqwdpython",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
			},
		},
		{
			name: "ok -C 3",
			args: args{
				s: initStrings(),
				f: "kek",
				flags: Flags{
					"A": 0,
					"B": 0,
					"C": 3,
					"c": 0,
					"F": false,
					"v": false,
					"i": false,
					"n": false,
				},
			},
			want: []string{
				"python",
				"python sefsfdues sbeifsbei dnjsf",
				"asiudhiwudsbqwdpython",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"qdlwijoqwdjoqwuhe2189e",
				"wkudhiqwudhqiwudhqiw",
				"wkudhiqwudhqiwudhqiw",
			},
		},
		{
			name: "ok -c 2",
			args: args{
				s: initStrings(),
				f: "kek",
				flags: Flags{
					"A": 0,
					"B": 0,
					"C": 0,
					"c": 2,
					"F": false,
					"v": false,
					"i": false,
					"n": false,
				},
			},
			want: []string{
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
			},
		},
		{
			name: "ok -i",
			args: args{
				s: initStrings(),
				f: "KEK",
				flags: Flags{
					"A": 0,
					"B": 0,
					"C": 0,
					"c": 0,
					"F": false,
					"v": false,
					"i": true,
					"n": false,
				},
			},
			want: []string{
				"kek eadiajodjeafh b",
				"KEK eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"KEK eadiajodjeafh b",
				"kek eadiajodjeafh b",
				"kek eadiajodjeafh b",
			},
		},
		{
			name: "ok -v",
			args: args{
				s: initStrings(),
				f: "kek",
				flags: Flags{
					"A": 0,
					"B": 0,
					"C": 0,
					"c": 0,
					"F": false,
					"v": true,
					"i": false,
					"n": false,
				},
			},
			want: []string{
				"aaaaaaaaaaaaaa",
				"bbbbbbbbbbbbb",
				"cccccccccccc",
				"wwwwwwwwwwwwwwwww",
				"qqqqqqqqqqqqqqq",
				"eeeeeeeeeeeeeeeeeee",
				"python",
				"python sefsfdues sbeifsbei dnjsf",
				"asiudhiwudsbqwdpython",
				"python sefsfdues sbeifsbei dnjsf",
				"asiudhiwudsbqwdpython",
				"KEK eadiajodjeafh b",
				"python sefsfdues sbeifsbei dnjsf",
				"asiudhiwudsbqwdpython",
				"python sefsfdues sbeifsbei dnjsf",
				"asiudhiwudsbqwdpython check",
				"python sefsfdues sbeifsbei dnjsf",
				"wkudhiqwudhqiwudhqiw",
				"KEK eadiajodjeafh b",
				"qdlwijoqwdjoqwuhe2189e",
				"wkudhiqwudhqiwudhqiw",
				"qdlwijoqwdjoqwuhe2189e",
				"wkudhiqwudhqiwudhqiw",
				"qdlwijoqwdjoqwuhe2189e",
				"wkudhiqwudhqiwudhqiw",
				"wkudhiqwudhqiwudhqiw",
				"wkudhiqwudhqiwudhqiw",
				"wkudhiqwudhqiwudhqiw",
				"wkudhiqwudhqiwudhqiw",
			},
		},
		{
			name: "ok -F",
			args: args{
				s: initStrings(),
				f: "python",
				flags: Flags{
					"A": 0,
					"B": 0,
					"C": 0,
					"c": 0,
					"F": true,
					"v": false,
					"i": false,
					"n": false,
				},
			},
			want: []string{
				"python",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := grep(tt.args.s, tt.args.f, tt.args.flags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("grep() = %v, want %v", got, tt.want)
			}
		})
	}
}
