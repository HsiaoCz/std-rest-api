package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Key string

// context with golang
func main() {
	start := time.Now()
	var str = "foo"
	ctx := context.WithValue(context.Background(), str, "bar")
	userID := 10
	val, err := fetchUserData(ctx, userID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result:", val)
	fmt.Println("took:", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userID int) (int, error) {
	val := ctx.Value("foo")
	fmt.Println(val)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()

	respch := make(chan Response)

	go func() {
		val, err := fetchThirdPartyStuffWhichCanBeSlow(ctx)
		respch <- Response{
			value: val,
			err:   err,
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("fetching data from third party took to long")
		case resp := <-respch:
			return resp.value, resp.err
		}
	}
}

// fetchThirdPartyStuffWhichCanBeSlow some function
// some real function
func fetchThirdPartyStuffWhichCanBeSlow(ctx context.Context) (int, error) {
	time.Sleep(time.Millisecond * 500)
	return 666, nil
}
