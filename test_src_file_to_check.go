package main

import (
	"context"
)

func normalAdd(a, b int) int {

	return a + b

}

func ooook(ctx context.Context, a, b int) int {
	return a + b
}
