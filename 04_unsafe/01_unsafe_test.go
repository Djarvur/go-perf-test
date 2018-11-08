package unsafe_test

import (
	"bytes"
	"encoding/gob"
	"math/rand"
	"testing"
	"unsafe"
)

type header struct {
	data     unsafe.Pointer
	len, cap int
}

// usable only when you sure
// that you'll never accidentally
// mutate `res`
// (its size in particular)
func encodeMut(data []uint64) (res []byte) {
	sz := len(data) * 8

	dh := (*header)(unsafe.Pointer(&data))
	rh := &header{
		data: dh.data,
		len:  sz,
		cap:  sz,
	}
	res = *(*[]byte)(unsafe.Pointer(&rh))
	return
}

var Response []byte

func BenchmarkGob_WTF_Thats_An_Empty_Buffer(b *testing.B) {
	data := make([]uint64, 0)

	bb := bytes.NewBuffer(nil)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		e := gob.NewEncoder(bb)
		e.Encode(data)
		Response = bb.Bytes()
		b.SetBytes(0)
	}
}

func BenchmarkGob(b *testing.B) {
	data := make([]uint64, 128)

	for i := 0; i < 128; i++ {
		data[i] = rand.Uint64()
	}

	bb := bytes.NewBuffer(nil)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		e := gob.NewEncoder(bb)
		e.Encode(data)
		Response = bb.Bytes()
		b.SetBytes(1024)
	}
}

func BenchmarkUnsafeMut(b *testing.B) {
	data := make([]uint64, 128)

	for i := 0; i < 128; i++ {
		data[i] = rand.Uint64()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Response = encodeMut(data)
		b.SetBytes(1024)
	}
}
