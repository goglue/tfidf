package tfidf

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestWeigher_StringsScore(t *testing.T) {
	ds := new(documentStorageMock)
	d := `abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc
	abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc
	abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc
	abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc
	abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc abc cat cat cat`

	dc := StripSpacesLoop(d)
	w := New(ds)
	r := w.Score(dc)

	expectedAbc := 8.934030160816897
	expectedCat := 0.2763102111592855
	assert.Equal(t, expectedAbc, r["abc"])
	assert.Equal(t, expectedCat, r["cat"])
}

type documentStorageMock struct {
	mock.Mock
}

func (d *documentStorageMock) Documents() uint {
	return uint(10000000)
}

func (d *documentStorageMock) DocumentsWith(string) uint {
	return 1000
}
