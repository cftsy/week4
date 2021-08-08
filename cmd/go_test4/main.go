package go_test4

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"test2/pkg/data"
	"test2/pkg/service"
)

func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "cancel_func", cancel)
	ctx = context.WithValue(ctx, "wg", &wg)

	go handleSignal(cancel)
	data.StartData(ctx, &wg)
	service.StartService(ctx, &wg)
	wg.Wait()
}

func handleSignal(cancelFunc context.CancelFunc) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				cancelFunc()

			default:
				fmt.Println("rcv err signal")
			}
		}
	}()
}
