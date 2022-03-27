package structural

import (
	"fmt"
	"time"
)

type Function func(float64) float64

func ProfileDecorator(fn Function) Function {
	return func(f float64) float64 {
		start := time.Now()
		result := fn(f)
		elapsed := time.Since(start)
		fmt.Printf("elapsed: %v\n", elapsed)
		return result
	}
}
