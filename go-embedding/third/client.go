package third

import (
	"context"
	"fmt"
	"github.com/shrimalmadhur/playground/go-embedding/first/types"
)

type Client struct {
	 client types.Client
}

func (c Client) MethodA(ctx context.Context, message string) {
	fmt.Println(fmt.Sprintf("client impl message MethodA: %s", message))
}

func (c Client) MethodB(ctx context.Context, message string) {
	fmt.Println(fmt.Sprintf("client impl message MethodB: %s", message))
}

func NewClient() Client {
	return Client{}
}