package tfidf

type FrequencyTable interface {
}

type Weigher struct {
	ft FrequencyTable
}

type term struct {
	term string
	tfIdf float64
}

func (t *term) String() string {
	return t.term
}

func (t *term) TfIdf() float64 {
	return t.tfIdf
}

func (w *Weigher) StringsScore(d string) []term {

	return nil
}

func New() *Weigher {
	return &Weigher{}
}
