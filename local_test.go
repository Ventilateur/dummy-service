package main

import (
	"net/url"
	"testing"
)

var (
	str = "http://api.houston.dev.schibsted.io/v1/houston/configs/search?tenant=qwer&userId=&nhinha=nhino"
	u, err = url.ParseRequestURI(str)
)

func Benchmark_URLToString(b *testing.B) {
	if err != nil {
		b.Fatal(err)
	}
	for n := 0; n < b.N; n++ {
		_ = u.String()
	}
}

func Benchmark_ConcatString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = str + "abcxyz"
	}
}
