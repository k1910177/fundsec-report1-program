package main

import (
	"crypto/rand"
	"testing"

	sbox "github.com/k1910177/funsec-report1/sbox"
	ttable "github.com/k1910177/funsec-report1/ttable"
)

func BenchmarkAES128SBox(b *testing.B) {
	nb, nk, nr := 4, 4, 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		src := make([]byte, nb*4)
		key := make([]byte, nk*4)
		rand.Read(src)
		rand.Read(key)

		sbox.Encrypt(src, key, nr)
	}
}

func BenchmarkAES128TTable(b *testing.B) {
	nb, nk, nr := 4, 4, 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		src := make([]byte, nb*4)
		key := make([]byte, nk*4)
		rand.Read(src)
		rand.Read(key)

		ttable.Encrypt(src, key, nr)
	}
}

func BenchmarkAES192SBox(b *testing.B) {
	nb, nk, nr := 4, 6, 12
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		src := make([]byte, nb*4)
		key := make([]byte, nk*4)
		rand.Read(src)
		rand.Read(key)

		sbox.Encrypt(src, key, nr)
	}
}

func BenchmarkAES192TTable(b *testing.B) {
	nb, nk, nr := 4, 6, 12
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		src := make([]byte, nb*4)
		key := make([]byte, nk*4)
		rand.Read(src)
		rand.Read(key)

		ttable.Encrypt(src, key, nr)
	}
}

func BenchmarkAES256SBox(b *testing.B) {
	nb, nk, nr := 4, 8, 14
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		src := make([]byte, nb*4)
		key := make([]byte, nk*4)
		rand.Read(src)
		rand.Read(key)

		sbox.Encrypt(src, key, nr)
	}
}

func BenchmarkAES256TTable(b *testing.B) {
	nb, nk, nr := 4, 8, 14
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		src := make([]byte, nb*4)
		key := make([]byte, nk*4)
		rand.Read(src)
		rand.Read(key)

		ttable.Encrypt(src, key, nr)
	}
}
