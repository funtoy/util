package util

import (
	"fmt"
	"testing"
)

func TestNewSeat(t *testing.T) {
	seat := NewSeat(2)
	pos, ok := seat.Find("Bill")
	if !ok {
		t.Fail()
	}
	if pos != 0 {
		t.Fail()
	}
	pos, ok = seat.Find("Tom")
	if !ok {
		t.Fail()
	}
	if pos != 1 {
		t.Fail()
	}

	pos, ok = seat.Find("Bing")
	if ok {
		t.Fail()
	}
	seat.Leave("Tom")
	pos, ok = seat.Find("Bing")
	if !ok {
		t.Fail()
	}
	if pos != 1 {
		t.Fail()
	}
	fmt.Println(seat)

	key, ok := seat.WhoIs(1)
	if !ok {
		t.Fail()
	}
	if key != "Bing" {
		t.Fail()
	}
	pos, ok = seat.WhereIs("Bill")
	if !ok {
		t.Fail()
	}
	if pos != 0 {
		t.Fail()
	}
	_, ok = seat.WhoIs(2)
	if ok {
		t.Fail()
	}

}