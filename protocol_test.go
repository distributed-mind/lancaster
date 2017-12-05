// protocol_test.go
package main

import (
	"testing"
)

func TestNakRegions_Init(t *testing.T) {
	r := nakRegions(make([]nakRegion, 0, 1))
	if len(r) != 0 {
		t.Fatal("len(r) != 0")
	}
}

func TestNakRegions_Clear(t *testing.T) {
	r := nakRegions(make([]nakRegion, 0, 1))
	r.Clear(10)
	if len(r) != 1 {
		t.Fatal("len(r) != 1")
	}
	if r[0].start != 0 {
		t.Fatal("r[0].start != 0")
	}
	if r[0].endEx != 10 {
		t.Fatal("r[0].endEx != 10")
	}
}

// [].ack(?, ?) => []
func TestNakRegions_Ack1(t *testing.T) {
	r := nakRegions(make([]nakRegion, 0, 1))
	r.Ack(0, 10)
	if len(r) != 0 {
		t.Fatal("len(r) != 0")
	}
}

// [(0, 10)].ack(0, 10) => []
func TestNakRegions_Ack2(t *testing.T) {
	r := nakRegions(make([]nakRegion, 0, 1))
	r.Clear(10)
	r.Ack(0, 10)
	if len(r) != 0 {
		t.Fatal("len(r) != 0")
	}
}

// [(0, 10)].ack(0,  5) => [(5, 10)]
func TestNakRegions_Ack3(t *testing.T) {
	r := nakRegions(make([]nakRegion, 0, 1))
	r.Clear(10)
	r.Ack(0, 5)
	if len(r) != 1 {
		t.Fatal("len(r) != 1")
	}
	if r[0].start != 5 {
		t.Fatal("r[0].start != 5")
	}
	if r[0].endEx != 10 {
		t.Fatal("r[0].endEx != 10")
	}
}

// [(0, 10)].ack(5, 10) => [(0,  5)]
func TestNakRegions_Ack4(t *testing.T) {
	r := nakRegions(make([]nakRegion, 0, 1))
	r.Clear(10)
	r.Ack(5, 10)
	if len(r) != 1 {
		t.Fatal("len(r) != 1")
	}
	if r[0].start != 0 {
		t.Fatal("r[0].start != 0")
	}
	if r[0].endEx != 5 {
		t.Fatal("r[0].endEx != 5")
	}
}

// [(0, 10)].ack(2,  5) => [(0,  2), (5, 10)]
func TestNakRegions_Ack5(t *testing.T) {
	r := nakRegions(make([]nakRegion, 0, 1))
	r.Clear(10)
	r.Ack(2, 5)
	if len(r) != 2 {
		t.Fatal("len(r) != 2")
	}
	if r[0].start != 0 {
		t.Fatal("r[0].start != 0")
	}
	if r[0].endEx != 2 {
		t.Fatal("r[0].endEx != 2")
	}
	if r[1].start != 5 {
		t.Fatal("r[1].start != 5")
	}
	if r[1].endEx != 10 {
		t.Fatal("r[1].endEx != 10")
	}
}
