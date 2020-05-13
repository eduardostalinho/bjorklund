package bjorklund

// Pattern is a sequence of Os and 1s representing the presence of a pulse
// in a sequence of bits.
type Pattern []int

// Rotate rotates a pattern to start in the given index.
func (p Pattern) Rotate(v int) Pattern {
	return append(p[v:], p[:v]...)
}

// Get pulses returns a slice informing the indexes of 1s in the pattern.
func (p Pattern) GetPulses() (pulses []int) {
	for i, x := range p {
		if x == 1 {
			pulses = append(pulses, i)
		}
	}
	return pulses
}

func (p Pattern) computeStartingOne() int {
	for i, x := range p {
		if x == 1 {
			return i
		}
	}
	return 0
}
