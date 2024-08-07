package main

import (
	"context"
	"fmt"
)

func normalAdd(a, b int) int {
	return a + b

}

func ooook(ctx context.Context, a, b int) int {
	fmt.Println(ctx)
	return a + b
}
