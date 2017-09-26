package main

import (
	"bytes"
	"fmt"
)

// RegisterList - list of registers used by the CPU
type RegisterList [16]Register

func (rl RegisterList) String() string {
	var buffer bytes.Buffer

	for index, reg := range rl {
		buffer.WriteString(fmt.Sprintf("r%d:\t%v\n", index, reg))
	}
	return buffer.String()
}
