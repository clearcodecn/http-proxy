package http_proxy

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	b := make([]byte, 255)
	for i := 0; i < 255; i++ {
		b[i] = uint8(i)
	}

	original := bytes.NewBuffer(nil)

	rw := New(original, original)
	n, err := rw.Write(b)
	if err != nil {
		t.Fatal(err)
	}
	if n != len(b) {
		t.Fatal("failed")
	}
	var dst = make([]byte, 255)

	n, err = rw.Read(dst)
	if err != nil {
		t.Fatal(err)
		return
	}

	if n != len(dst) {
		t.Fatal("fail")
	}
	fmt.Println(b)
	fmt.Println(dst)
	if !bytes.Equal(b, dst) {
		t.Fatal("failed")
	}
}
