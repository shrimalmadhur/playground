package client

import (
	"context"
	"fmt"
	cl "github.com/shrimalmadhur/playground/go-embedding/third"
)

type Client struct {
	cl.Client
}

func (s Client) MethodA(ctx context.Context, message string) {
	fmt.Println(fmt.Sprintf("Override MethodA message: %s", message))
}

func NewClient() Client {
	sdkClient := cl.NewClient()
	return Client {
		sdkClient,
	}
}
