package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkWriteString(b *testing.B) {
	var v struct {
		Name string
	}
	v.Name = "SomeName"
	for i := 0; i < b.N; i++ {
		var tempBuffer = bytes.NewBuffer([]byte{})
		tempBuffer.WriteString(" ")
		tempBuffer.WriteString(strings.ToLower(v.Name))
		tempBuffer.WriteString(" ")
	}
}

func BenchmarkSprintf(b *testing.B) {
	var v struct {
		Name string
	}
	v.Name = "SomeName"
	for i := 0; i < b.N; i++ {
		var tempBuffer = bytes.NewBuffer([]byte{})
		tempBuffer.WriteString(fmt.Sprintf(" %s ", strings.ToLower(v.Name)))
	}
}

// 运行基准测试
// go test -bench=.
