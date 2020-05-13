// Package bjorklund is used to generate evenly distributed sequences of 0s and 1s.
// Based on the algorithm proposed in "The Theory of Rep-Rate Pattern Generation in the SNS Timing System" by E. Bjorklund.
//
package bjorklund

// Bjorklund passes variables needed to compute an Euclidian
// pattern given a number of pulses and slots.
type Bjorklund struct {
	Slots      int
	Pulses     int
	counts     []int
	remainders []int
	Pattern    Pattern
}

func CreatePattern(pulses, slots, rotation int) (Pattern, error) {
	b, err := NewBjorklund(pulses, slots)
	if err != nil {
		return nil, err
	}
	pattern := b.Compute()
	return pattern.Rotate(rotation), nil
}

// NewBjorklund constructs a new Bjorklund struct to be
// computed later given the number of available slots in
// the cycle and the expected number of pulses (ones)  in it.
func NewBjorklund(pulses, slots int) (*Bjorklund, error) {
	if pulses > slots {
		return nil,  PulsesGreaterThanSlotsErr{Pulses: pulses, Slots: slots}
	}
	return &Bjorklund{Slots: slots, Pulses: pulses}, nil
}

// Compute creates a new pattern based on bjorklund`s
// number of pulses and number of slots. Returns an evenly
// distributed sequence of length `b.Slots` with `b.Pulses ones.`
func (b *Bjorklund) Compute() Pattern {
	level := b.computeParameters()

	b.buildPattern(level)

	i := b.Pattern.computeStartingOne()
	return b.Pattern.Rotate(i)
}

func (b *Bjorklund) computeParameters() int {
	b.remainders = append(b.remainders, b.Pulses)
	divisor := b.Slots - b.Pulses

	level := 0
	for {
		b.counts = append(b.counts, divisor/b.remainders[level])
		b.remainders = append(b.remainders, divisor%b.remainders[level])
		divisor = b.remainders[level]
		level++
		if b.remainders[level] <= 1 {
			break
		}
	}
	b.counts = append(b.counts, divisor)
	return level
}

func (b *Bjorklund) buildPattern(level int) {
	switch level {
	case -1:
		b.Pattern = append(b.Pattern, 0)
	case -2:
		b.Pattern = append(b.Pattern, 1)
	default:
		for i := 0; i < b.counts[level]; i++ {
			b.buildPattern(level - 1)
		}
		if b.remainders[level] != 0 {
			b.buildPattern(level - 2)
		}
	}
}
