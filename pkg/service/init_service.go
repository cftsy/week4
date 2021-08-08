package service

import (
	"context"
	"fmt"
	"sync"
)

func StartService(ctx context.Context, wg *sync.WaitGroup) {
	fmt.Println("service")
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