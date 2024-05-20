package main

import (
	"bytes"
	"io"
	"testing"
)

func writeToBuffer(buf io.Writer, msg []byte) {
	buf.Write(msg)
}

func BenchmarkWriteToBuffer(b *testing.B) {
	data := []byte("Clint")
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		for j := 0; j < 100; j++ {
			writeToBuffer(buf, data)
		}
	}
}

func BenchmarkTest(b *testing.B) {
	n := 100
	b.Run("No allocation of size", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ints := []int{}
			for j := 0; j < n; j++ {
				ints = append(ints, j)
			}
		}
	})
	b.Run("Allocation of size", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ints := make([]int, n)
			for j := 0; j < n; j++ {
				ints[j] = j
			}
		}
	})

}
