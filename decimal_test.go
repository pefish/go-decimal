package p_decimal

import (
	"testing"
)

func TestDecimalClass_Lt(t *testing.T) {
	if Decimal.Start(`12`).Lt(10) {
		t.Error()
	}
}
