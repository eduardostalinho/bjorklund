// Bjorklund
package bjorklund_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/eduardostalinho/bjorklund"
)

func TestBjorklund(t *testing.T) {
	cases := []struct {
		pulses int
		steps  int
		want   bjorklund.Pattern
	}{
		{4, 16, bjorklund.Pattern{1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 0}},
		{5, 13, bjorklund.Pattern{1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0}},
		{3, 8, bjorklund.Pattern{1, 0, 0, 1, 0, 0, 1, 0}},
	}
	for _, c := range cases {
		testName := fmt.Sprintf("bjorklund %d on %d", c.pulses, c.steps)
		t.Run(testName, func(t *testing.T) {
			b, err := bjorklund.NewBjorklund(c.pulses, c.steps)
			if err != nil {
				t.Fatalf("unexpected error, %v", err)
			}
			pattern := b.Compute()

			if !reflect.DeepEqual(pattern, c.want) {
				t.Errorf("want %v, got %v", c.want, pattern)
			}
		})
	}
}
