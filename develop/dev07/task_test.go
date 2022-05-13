package main

import (
	"testing"
	"time"
)

func Test_or(t *testing.T) {
	type args struct {
		channels []<-chan interface{}
	}
	tests := []struct {
		name     string
		args     args
		wantLess time.Duration
	}{
		{
			name: "ok",
			args: args{channels: []<-chan interface{}{
				sig(1 * time.Second),
				sig(3 * time.Second),
				sig(4 * time.Second),
			}},
			wantLess: 4300 * time.Millisecond,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			<-or(tt.args.channels...)

			tDuration := time.Since(start)
			if tDuration > tt.wantLess {
				t.Errorf("or() time duration = %v, want %v", tDuration, tt.wantLess)
			}

		})

	}
}
