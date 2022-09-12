package main

import (
	"fmt"

	"strings"

	"github.com/bbedward/go-opencl/opencl"
)

const (
	deviceType = opencl.DeviceTypeAll

	dataSize = 128

	programCode = `
kernel void kern(global float* out)
{
	size_t i = get_global_id(0);
	out[i] = i;
}
`
)

func printHeader(name string) {
	fmt.Println(strings.ToUpper(name))
	for _ = range name {
		fmt.Print("=")
	}
	fmt.Println()
}
