package enumerator

import (
	"context"
	"sync"

	"github.com/solsw/golinq/common"
	"github.com/solsw/golinq/errors"
)

// ForEach sequentially performs action 'a' on each element of the collection.
func ForEach(en Enumerator, a func(common.Elem)) error {
	if a == nil {
		return errors.NilAction
	}
	en.Reset()
	for en.MoveNext() {
		a(en.Current())
	}
	return nil
}

// ForEachCtx sequentially performs action 'a' on each element of the collection.
// 'ctx' may be used to cancel the operation in progress.
func ForEachCtx(ctx context.Context, en Enumerator, a func(context.Context, common.Elem)) error {
	if a == nil {
		return errors.NilAction
	}
	if ctx == nil {
		ctx = context.Background()
	}
	en.Reset()
	for en.MoveNext() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			a(ctx, en.Current())
		}
	}
	return nil
}

// ForEachConcurrent concurrently (using goroutines) performs action 'a' on each element of the collection.
func ForEachConcurrent(en Enumerator, a func(common.Elem)) error {
	if a == nil {
		return errors.NilAction
	}
	en.Reset()
	var wg sync.WaitGroup
	for en.MoveNext() {
		wg.Add(1)
		go func(e common.Elem) {
			defer wg.Done()
			a(e)
		}(en.Current())
	}
	wg.Wait()
	return nil
}

// ForEachConcurrentCtx concurrently (using goroutines) performs action 'a' on each element of the collection.
// 'ctx' may be used to cancel the operation in progress.
func ForEachConcurrentCtx(ctx context.Context, en Enumerator, a func(context.Context, common.Elem)) error {
	if a == nil {
		return errors.NilAction
	}
	if ctx == nil {
		ctx = context.Background()
	}
	var wg sync.WaitGroup
	en.Reset()
	for en.MoveNext() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			wg.Add(1)
			go func(e common.Elem) {
				defer wg.Done()
				a(ctx, e)
			}(en.Current())
		}
	}
	wg.Wait()
	return nil
}
