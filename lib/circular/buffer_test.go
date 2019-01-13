package circular_test

import (
	"fmt"
	"math/rand"
	"metalim/advent/2017/lib/circular"
	"testing"
	"time"

	. "github.com/logrusorgru/aurora"
)

var _t *testing.T

func verify(v, ex int) {
	if v != ex {
		_t.Fatal("got", v, "expected", ex)
	}
}

type newfun func(vals ...circular.Value) circular.Buffer

func testWith(t *testing.T, fn newfun) {
	_t = t

	l0 := fn()
	verify(l0.Length(), 0)

	l1 := fn(1)
	verify(l1.Length(), 1)
	verify(l1.Get(), 1)

	l2 := fn(1, 2)
	verify(l2.Length(), 2)
	verify(l2.Get(), 1)

	l2.InsertAfter(3)
	verify(l2.Length(), 3)
	verify(l2.Get(), 3)

	l2.Skip(1)
	verify(l2.Get(), 2)
	l2.Skip(1)
	verify(l2.Get(), 1)
	l2.Skip(1)
	verify(l2.Get(), 3)

	l2.Skip(7)
	verify(l2.Get(), 2)

	l2.Skip(-7)
	verify(l2.Get(), 3)

	l2.Skip(rand.Intn(100) - 50)
	l2.InsertAfter(l2.Get() + rand.Intn(1000))
	l2.Skip(rand.Intn(100) - 50)
	l2.InsertBefore(l2.Get() + rand.Intn(1000))

}

func TestList(t *testing.T) {
	testWith(t, circular.NewList)
}

func TestSlice(t *testing.T) {
	testWith(t, circular.NewSlice)
}

func benchmarkInsertWith(b *testing.B, fn newfun) {
	l := fn(1, 2, 3)
	fmt.Println("benchInsert", b.N)
	var stop bool
	var i int
	go func() {
		time.Sleep(time.Second)
		for !stop {
			fmt.Println(Black(i).Bold())
			time.Sleep(time.Second)
		}
	}()
	for i = 0; i < b.N; i += 2 {
		l.Skip(rand.Intn(100) - 50)
		l.InsertAfter(l.Get() + rand.Intn(1000))
		l.Skip(rand.Intn(100) - 50)
		l.InsertBefore(l.Get() + rand.Intn(1000))
	}
	stop = true
}

func benchmarkFixedWith(b *testing.B, fn newfun) {
	l := fn()
	for i := 0; i < 256; i++ {
		l.InsertAfter(i)
	}
	l.Skip(0)

	fmt.Println("benchFixed", b.N)
	for i := 0; i < b.N; i++ {
		l.Set(l.Get() + rand.Intn(256))
	}
}

func BenchmarkFixedList(b *testing.B) {
	benchmarkFixedWith(b, circular.NewList)
}

func BenchmarkFixedSlice(b *testing.B) {
	benchmarkFixedWith(b, circular.NewSlice)
}

func BenchmarkInsertList(b *testing.B) {
	benchmarkInsertWith(b, circular.NewList)
}

func BenchmarkInsertSlice(b *testing.B) { // dies with 10m timeout
	benchmarkInsertWith(b, circular.NewSlice)
}
