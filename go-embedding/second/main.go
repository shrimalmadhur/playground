package main

import (
	"context"
	"github.com/shrimalmadhur/playground/go-embedding/first/types"
	"github.com/shrimalmadhur/playground/go-embedding/second/client"
)


func main()  {
	ctx := context.Background()

	myClient := client.NewClient()

	service := types.NewService(myClient)

	service.MethodA(ctx, "hello")

	service.MethodB(ctx, "hello")
}