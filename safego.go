package safego

import (
	"context"
	"log"
)

// Go executes the provided function fn in a goroutine with safety features.
// It provides the following safety mechanisms:
//
//	`Panic recovery` to prevent application crashes
//	`Parameters`:
//	    'fn': The function to be executed in the goroutine
//
// Any panics during execution will be recovered and logged.
func Go(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
			}
		}()
		fn()
	}()
}

// GoContext executes the provided function fn in a goroutine with safety features.
// It provides the following safety mechanisms:
//
//	`Panic recovery` to prevent application crashes
//	`Context` checking for proper cancellation support
//	`Parameters`:
//	    'ctx': Context for cancellation control
//	    'fn': The function to be executed in the goroutine
//
// The function will not execute if the context is already canceled.
// Any panics during execution will be recovered and logged.
func GoContext(ctx context.Context, fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Recovered from panic: %v", r)
			}
		}()

		// If the context is already canceled, exit without executing fn.
		select {
		case <-ctx.Done():
			log.Printf("Context canceled: %v", ctx.Err())
			return
		default:
			fn()
		}
	}()
}
