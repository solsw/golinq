package test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/enumerable"
	"github.com/solsw/golinq/enumerator"
)

func TestForEach(t *testing.T) {
	type args struct {
		en enumerator.Enumerator
		a  func(common.Elem)
	}
	var s int
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    int
	}{
		{name: "0",
			args: args{
				en: enumerable.RangeMust(1, 1000),
				a:  nil},
			wantErr: true},
		{name: "1",
			args: args{
				en: enumerable.RangeMust(1, 1000),
				a:  func(e common.Elem) { s += e.(int) }},
			want: 500500},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := enumerator.ForEach(tt.args.en, tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("ForEach() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if s != tt.want {
				t.Errorf("ForEach() = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestForEachCtx(t *testing.T) {
	var i, s int
	type args struct {
		ctx context.Context
		en  enumerator.Enumerator
		a   func(context.Context, common.Elem)
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		expectedErr error
		want        int
	}{
		{name: "1",
			args: args{
				en: enumerable.RangeMust(1, 100000),
				a: func(_ context.Context, e common.Elem) {
					i++
					s += e.(int)
				}},
			want: 5000050000},
		{name: "2",
			args: args{
				en: enumerable.RangeMust(1, 100000),
				a: func(_ context.Context, e common.Elem) {
					i++
					s += e.(int)
				}},
			wantErr:     true,
			expectedErr: context.DeadlineExceeded,
			want:        5000050000},
	}
	for _, tt := range tests {
		i = 0
		s = 0
		if tt.name == "2" {
			timeout, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
			defer cancel()
			tt.args.ctx = timeout
		}
		t.Run(tt.name, func(t *testing.T) {
			err := enumerator.ForEachCtx(tt.args.ctx, tt.args.en, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForEachCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ForEachCtx() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if s != tt.want {
				t.Errorf("ForEachCtx() = %v, want %v", s, tt.want)
			}
		})
		t.Logf("index = %d", i)
		t.Logf("sum = %d", s)
	}
}

func TestForEachConcurrent(t *testing.T) {
	type args struct {
		en enumerator.Enumerator
		a  func(common.Elem)
	}
	var s int64
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    int64
	}{
		{name: "0",
			args: args{
				en: enumerable.RangeMust(1, 10),
				a:  nil},
			wantErr: true},
		{name: "1",
			args: args{
				en: enumerable.RangeMust(1, 100000),
				a:  func(e common.Elem) { atomic.AddInt64(&s, int64(e.(int))) }},
			want: 5000050000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := enumerator.ForEachConcurrent(tt.args.en, tt.args.a); (err != nil) != tt.wantErr {
				t.Errorf("ForEachConcurrent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if s != tt.want {
				t.Errorf("ForEachConcurrent() = %v, want %v", s, tt.want)
			}
		})
	}
}

func TestForEachConcurrentCtx(t *testing.T) {
	var i, s int64
	type args struct {
		ctx context.Context
		en  enumerator.Enumerator
		a   func(context.Context, common.Elem)
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		expectedErr error
		want        int
	}{
		{name: "1",
			args: args{
				en: enumerable.RangeMust(1, 100000),
				a: func(_ context.Context, e common.Elem) {
					atomic.AddInt64(&i, int64(1))
					atomic.AddInt64(&s, int64(e.(int)))
				}},
			want: 5000050000},
		{name: "2",
			args: args{
				en: enumerable.RangeMust(1, 100000),
				a: func(_ context.Context, e common.Elem) {
					atomic.AddInt64(&i, int64(1))
					atomic.AddInt64(&s, int64(e.(int)))
				}},
			wantErr:     true,
			expectedErr: context.DeadlineExceeded,
			want:        5000050000},
	}
	for _, tt := range tests {
		i = 0
		s = 0
		if tt.name == "2" {
			timeout, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
			defer cancel()
			tt.args.ctx = timeout
		}
		t.Run(tt.name, func(t *testing.T) {
			err := enumerator.ForEachConcurrentCtx(tt.args.ctx, tt.args.en, tt.args.a)
			if (err != nil) != tt.wantErr {
				t.Errorf("ForEachConcurrentCtx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err != tt.expectedErr {
					t.Errorf("ForEachConcurrentCtx() error = %v, expectedErr %v", err, tt.expectedErr)
				}
				return
			}
			if int(s) != tt.want {
				t.Errorf("ForEachConcurrentCtx() = %v, want %v", s, tt.want)
			}
		})
		t.Logf("index = %d", i)
		t.Logf("sum = %d", s)
	}
}
