package tests

import "testing"

func TestNoop(t *testing.T) {
	noopCGO()
	noopGO()
	noopCall()
	noopCallNoStack()
	noopCallSlow()
}

func TestAdd(t *testing.T) {
	var rets = []int{
		addCGO(999, 888),
		addGO(999, 888),
		addCall(999, 888),
		addCallNoStack(999, 888),
		addCallSlow(999, 888),
	}
	for i := 0; i < len(rets); i++ {
		if rets[i] != 1887 {
			t.Fatalf("bad news %d, %d != %d", i, rets[i], 1887)
		}
	}
}

func TestBigStack(t *testing.T) {
	var rets = []int{
		bigStackCGO(100000),
		bigStackGO(100000),
		bigStackCallWithStack(100000),
		bigStackCallSlow(100000),
	}
	for i := 0; i < len(rets); i++ {
		if rets[i] != 12742320 {
			t.Fatalf("bad news %d, %d != %d", i, rets[i], 12742320)
		}
	}
}

func BenchmarkNoopCGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noopCGO()
	}
}
func BenchmarkNoopCallSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noopCallSlow()
	}
}
func BenchmarkNoopCall(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noopCall()
	}
}
func BenchmarkNoopCallNoStack(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noopCallNoStack()
	}
}

func BenchmarkNoopGO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		noopGO()
	}
}

func BenchmarkAddCGO(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = addCGO(i, j)
	}
}
func BenchmarkAddCallSlow(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = addCallSlow(i, j)
	}
}
func BenchmarkAddCall(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = addCall(i, j)
	}
}
func BenchmarkAddCallNoStack(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = addCallNoStack(i, j)
	}
}
func BenchmarkAddGO(b *testing.B) {
	var j int
	for i := 0; i < b.N; i++ {
		j = addGO(i, j)
	}
}

func TestBigSysIO(t *testing.T) {
	count := bigsysio("hello.txt", 8765, 3)
	if count == 0 {
		t.Fail()
	}
}
