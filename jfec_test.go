package jfec

import (
	"bytes"
	"testing"
)

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
