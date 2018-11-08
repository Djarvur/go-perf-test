package language_test

import "testing"

type testTypeInterface interface {
	Inc()
}

type testTypeStruct struct {
	n int64
}

func (s *testTypeStruct) Inc() {
	s.n++
}

var testStruct = testTypeStruct{}
var testInterface testTypeInterface = &testStruct
var testInterfaceEmpty interface{} = &testStruct

func BenchmarkStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testStruct.Inc()
	}
}

func BenchmarkInterface(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testInterface.Inc()
	}
}

func BenchmarkInterfaceRuntime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		testInterfaceEmpty.(testTypeInterface).Inc()
	}
}
