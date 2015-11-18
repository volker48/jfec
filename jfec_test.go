package jfec

import (
	"bytes"
	"testing"
	"testing/quick"
	"math/rand"
	"reflect"
)

func headerVals(vals []reflect.Value, r *rand.Rand) {
	n := rand.Intn(256)
	if n == 0 {
		n = 1
	}
	k := rand.Intn(n)

	if k == 0 {
		k = 1
	}

	vals[0] = reflect.ValueOf(byte(k))
	vals[1] = reflect.ValueOf(byte(n))

	var num int
	if n == 1 {
		num = 0
	} else {
		num = rand.Intn(n)
	}

	vals[2] = reflect.ValueOf(uint(num))

	inputSize := rand.Int63()

	vals[3] = reflect.ValueOf(inputSize)

}

func TestHeaderQuick(t *testing.T) {
	f := func(k, n byte, num uint, inputSize int64) bool {
		fec := NewFec(k, n)
		padding := getPadding(inputSize, k)
		header := gen_header(fec, padding, num)
		rdr := bytes.NewReader(header)
		headerN, headerK, headerPad, headerNum, err := parseHeader(rdr)
		return n == uint8(headerN) && k == uint8(headerK) && padding == headerPad && num == headerNum && err == nil
	}
	conf := &quick.Config{Values: headerVals}
	if err := quick.Check(f, conf); err != nil {
		t.Error(err)
	}
}

func TestHeaderGenParse(t *testing.T) {
	fec := NewFec(3, 10)
	padding := getPadding(100, fec.k)
	header := gen_header(fec, padding, 0)
	rdr := bytes.NewReader(header)
	n, k, pad, sh, err := parseHeader(rdr)
	if err != nil {
		t.Error("Unexpected error", err)
	}
	if n != 10 {
		t.Errorf("Expected %d for n but was %d", 10, n)
	}
	if k != 3 {
		t.Errorf("Expected %d for k, but was %d", 3, k)
	}
	if pad != padding {
		t.Errorf("Expected %d for pad but was %d", padding, pad)
	}
	if sh != 0 {
		t.Errorf("Expected %d for sh but was %d", 0, sh)
	}
}
