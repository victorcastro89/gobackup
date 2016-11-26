package main

import (
	"bytes"
	"strings"
	"testing"
)

// func BenchmarkConcat(b *testing.B) {
// 	b.ResetTimer()
// 	var str string
// 	for n := 0; n < b.N; n++ {
// 		str = concat1(str, "x")
// 	}
// 	b.StopTimer()

// 	if s := strings.Repeat("x", b.N); str != s {
// 		b.Errorf("unexpected result; got=%s, want=%s", str, s)
// 	}
// }

func BenchmarkBuffer(b *testing.B) {
	b.ResetTimer()
	var buffer bytes.Buffer
	var str string
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
	str = buffer.String()
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", buffer.String(), s)
	}
}

func BenchmarkAppend(b *testing.B) {
	b.ResetTimer()
	var parts []string
	var str string
	for n := 0; n < b.N; n++ {
		parts = append(parts, "x")
	}
	str = strings.Join(parts, "")
	b.StopTimer()

	if s := strings.Repeat("x", b.N); str != s {
		b.Errorf("unexpected result; got=%s, want=%s", str, s)
	}
}

func concat(str1 string, str2 string) string {
	length := len(str1) + len(str2)

	alloc := make([]byte, length)
	bl := 0
	bl += copy(alloc[bl:], str1)
	bl += copy(alloc[bl:], str2)
	return string(alloc[:length])
}

func concat1(str1 string, str2 string) string {
	var buffer bytes.Buffer
	buffer.WriteString(str2)

	return buffer.String()
}
