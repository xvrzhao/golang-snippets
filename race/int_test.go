package race

import "testing"

func TestIntRace1(t *testing.T) {
	IntRace1()
}

func TestIntRace2(t *testing.T) {
	IntRace2()
}

func TestIntAtomic(t *testing.T) {
	IntAtomic()
}
