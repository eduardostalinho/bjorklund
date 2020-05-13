package bjorklund

import (
	"fmt"
)

// PulsesGreaterThanSlotsErr provides an error when trying to create a pattern
// with pulses grater than slots.
type PulsesGreaterThanSlotsErr struct {
	Pulses int
	Slots  int
}

func (e PulsesGreaterThanSlotsErr) Error() string {
	return fmt.Sprintf("Number of slots (%d) must be greater than the number of pulses (%d).", e.Slots, e.Pulses)

}
