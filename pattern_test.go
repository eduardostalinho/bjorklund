package bjorklund_test

import (
	"reflect"
	"testing"

	"github.com/eduardostalinho/bjorklund"
)

func TestPattern(t *testing.T) {
	t.Run("test rotate", func(t *testing.T) {
		pattern := bjorklund.Pattern{1, 0, 0, 0, 0, 0}
		want := bjorklund.Pattern{0, 0, 0, 0, 0, 1}
		got := pattern.Rotate(1)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("test get pulses", func(t *testing.T) {
		pattern := bjorklund.Pattern{1, 0, 1, 0, 0, 0}
		want := []int{0, 2}
		got := pattern.GetPulses()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}
