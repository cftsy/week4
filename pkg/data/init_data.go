package data

import (
	"context"
	"fmt"
	"sync"
)

func StartData(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("data")
	wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
				wg.Done()
			}
		}()
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()
}
