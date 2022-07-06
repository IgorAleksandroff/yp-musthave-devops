package main

import (
	"fmt"
	"runtime"
)

type gauge float64
type counter int64

func main() {

	metrics := runtime.MemStats{}
	fmt.Println(metrics)

	runtime.ReadMemStats(&metrics)
	fmt.Println(metrics)
}
