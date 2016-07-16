package tfidf

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestWeigher_StringsScore(t *testing.T) {
	ds := new(documentStorageMock)
	d := `this is another document` // 4 words

	dc := StripSpacesLoop(d)
	w := New(ds)
	r := w.Score(dc)

	expectedAnother := 0.0752574989159953
	assert.Equal(t, expectedAnother, r["another"])
}

type documentStorageMock struct {
	mock.Mock
}

func (d *documentStorageMock) Documents() uint {
	return uint(2)
}

var dw = map[string]uint{
	"this":     2,
	"is":       2,
	"another":  1,
	"document": 0,
}

func (d *documentStorageMock) DocumentsWith(t string) uint {
	return dw[t]
}
