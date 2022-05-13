package main

import (
	"reflect"
	"testing"
)

func Test_cut(t *testing.T) {
	type args struct {
		strs  []string
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
				strs: []string{
					"aqcweqwceqwcecq",
					"asdsadqwdqwodijoi:dajoiwxqwijxdqnwxdj:xqwduoxh:bqwiudxqwd:xqdiwh",
					"qwid910238123:3127312!qwegyqw!asdvbzxjc!qwuwu",
					"qwiei912:e127c12e:gacgcag:dy383d38:10101010101",
					"qwid910238123;3127312;qwegyqw;asdvbzxjc;qwuwu",
				},
				flags: Flags{
					d: ":",
					f: map[int]struct{}{
						2: {},
					},
				},
			},
			want: []string{
				"aqcweqwceqwcecq",
				"dajoiwxqwijxdqnwxdj",
				"3127312!qwegyqw!asdvbzxjc!qwuwu",
				"e127c12e",
				"qwid910238123;3127312;qwegyqw;asdvbzxjc;qwuwu",
			},
		},
		{
			name: "ok -s",
			args: args{
				strs: []string{
					"aqcweqwceqwcecq",
					"asdsadqwdqwodijoi:dajoiwxqwijxdqnwxdj:xqwduoxh:bqwiudxqwd:xqdiwh",
					"qwid910238123:3127312!qwegyqw!asdvbzxjc!qwuwu",
					"qwiei912:e127c12e:gacgcag:dy383d38:10101010101",
					"qwid910238123;3127312;qwegyqw;asdvbzxjc;qwuwu",
				},
				flags: Flags{
					d: ":",
					f: map[int]struct{}{
						2: {},
					},
					s: true,
				},
			},
			want: []string{
				"dajoiwxqwijxdqnwxdj",
				"3127312!qwegyqw!asdvbzxjc!qwuwu",
				"e127c12e",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cut(tt.args.strs, tt.args.flags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cut() = %v, want %v", got, tt.want)
			}
		})
	}
}
