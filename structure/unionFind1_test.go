package structure

import (
	"math/rand"
	"testing"
	"time"
)

const N = 100000
//
//// 最糙的实现判定效率
//func TestUionFind1(t *testing.T) {
//	n := N
//	s := rand.NewSource(time.Now().UnixNano())
//	r := rand.New(s)
//	u1 := NewUnionFind1(n)
//	for i := 0; i < n; i++ {
//		a := r.Intn(n)
//		b := r.Intn(n)
//		u1.Union(a, b)
//	}
//
//	for i := 0; i < n; i++ {
//		a := r.Intn(n)
//		b := r.Intn(n)
//		u1.IsConnected(a, b)
//	}
//}
//
//// 使用Parent root 的方式提升效率
//func TestUionFind2(t *testing.T) {
//	n := N
//	s := rand.NewSource(time.Now().UnixNano())
//	r := rand.New(s)
//	u2 := NewUnionFind2(n)
//	for i := 0; i < n; i++ {
//		a := r.Intn(n)
//		b := r.Intn(n)
//		u2.Union(a, b)
//	}
//
//	for i := 0; i < n; i++ {
//		a := r.Intn(n)
//		b := r.Intn(n)
//		u2.IsConnected(a, b)
//	}
//}

// optimized by size
func TestUionFind3(t *testing.T) {
	n := N*10
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	u3 := NewUnionFind3(n)
	for i := 0; i < n; i++ {
		a := r.Intn(n)
		b := r.Intn(n)
		u3.Union(a, b)
	}

	for i := 0; i < n; i++ {
		a := r.Intn(n)
		b := r.Intn(n)
		u3.IsConnected(a, b)
	}
}

// optimized by size
func TestUionFind4(t *testing.T) {
	n := N*10
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	u4 := NewUnionFind4(n)
	for i := 0; i < n; i++ {
		a := r.Intn(n)
		b := r.Intn(n)
		u4.Union(a, b)
	}

	for i := 0; i < n; i++ {
		a := r.Intn(n)
		b := r.Intn(n)
		u4.IsConnected(a, b)
	}
}

func TestUionFind5(t *testing.T) {
	n := N*10
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	u5 := NewUnionFind5(n)
	for i := 0; i < n; i++ {
		a := r.Intn(n)
		b := r.Intn(n)
		u5.Union(a, b)
	}

	for i := 0; i < n; i++ {
		a := r.Intn(n)
		b := r.Intn(n)
		u5.IsConnected(a, b)
	}
}
