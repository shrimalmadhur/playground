package types

import "context"

type Client interface {
	MethodA(ctx context.Context, message string)

	MethodB(ctx context.Context, message string)
}
