package main_test

import (
	"testing"
    "encoding/hex"
    "encoding/base64"
)

func BenchmarkTohex(b *testing.B) {
	for i := 0; i < b.N; i++ {
        s := "Package hex implements hexadecimal encoding and decoding."
        hex.EncodeToString([]byte(s))
	}
}

func BenchmarkTo64(b *testing.B) {
    for i := 0; i < b.N; i++ {
        s := "Package hex implements hexadecimal encoding and decoding."
        //hex.EncodeToString([]byte(s))
        base64.StdEncoding.EncodeToString([]byte(s))
    }
}

func TestTohex(t *testing.T) {
    //Tohex()
}
