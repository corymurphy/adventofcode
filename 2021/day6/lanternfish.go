package main

type LanternFish struct {
	Age int
}

func NewLanternFish(initialAge int) LanternFish {
	return LanternFish{
		Age: initialAge,
	}
}

func (l *LanternFish) IncrementAge() bool {
	if l.Age == 0 {
		l.Age = 6
		return true
	} else {
		l.Age--
		return false
	}
}
