package main

import (
	"context"
	"fmt"
	"log"
	"plat/framework"
	"time"
)

func FoolControllerHandler(ctx *framework.Context) error {
	// Set 1s timeout
	durationCtx, cancel := context.WithTimeout(ctx.BaseContext(), time.Second)
	defer cancel()

	panicChan := make(chan interface{}, 1)
	finish := make(chan struct{}, 1)

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()

		// Do business

		// test context time out
		time.Sleep(10 * time.Millisecond)
		ctx.Json(200, "OK")
		finish <- struct{}{}

	}()

	select {
	case <-finish:
		fmt.Println("Finish")
	case p := <-panicChan:
		log.Println(p)
		ctx.Json(500, "panic")
	case <-durationCtx.Done():
		fmt.Println("Time Out")
		ctx.Json(500, "time out")
	}

	return nil
}
