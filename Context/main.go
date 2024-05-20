package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "test", "value")
	userId := 10
	val, err := fetchUserData(ctx, userId)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("UserId: ", val, " Took: ", time.Since(start))
}

type Response struct {
	value int
	err   error
}

func fetchUserData(ctx context.Context, userId int) (int, error) {
	value := ctx.Value("test")
	fmt.Println("value:", value)
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*200)
	defer cancel()
	resch := make(chan Response)

	go func() {
		u, err := fetchThirdPartyStuffCanBeSlow()
		resch <- Response{
			value: u,
			err:   err,
		}
	}()
	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("Fetch took too long!")
		case data := <-resch:
			return data.value, data.err

		}
	}
}

func fetchThirdPartyStuffCanBeSlow() (int, error) {
	time.Sleep(time.Millisecond * 100)
	return 66, nil
}
