package main

import "testing"

func BenchmarkEchoWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithLoop()
	}
}

func BenchmarkEchoWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoWithJoin()
	}
}

// to run: go test -bench=.
