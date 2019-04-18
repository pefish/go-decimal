package p_decimal

import (
	"testing"
)

func TestDecimalClass_Lt(t *testing.T) {
	if Decimal.Start(`12`).Lt(10) {
		t.Error()
	}
}

func TestDecimalClass_AbsForString(t *testing.T) {
	if Decimal.Start(`-1`).AbsForString() != `1` {
		t.Error()
	}
}

func TestDecimalClass_AddForString(t *testing.T) {
	if Decimal.Start(`-1`).AddForString(1) != `0` {
		t.Error()
	}
}

func TestDecimalClass_SubForString(t *testing.T) {
	if Decimal.Start(`-1`).SubForString(1) != `-2` {
		t.Error(Decimal.Start(`-1`).SubForString(1))
	}
}

func TestDecimalClass_DivForString(t *testing.T) {
	if Decimal.Start(`10`).DivForString(2) != `5` {
		t.Error()
	}
}

func TestDecimalClass_MultiForString(t *testing.T) {
	if Decimal.Start(`10`).MultiForString(2) != `20` {
		t.Error()
	}
}

func TestDecimalClass_TruncForString(t *testing.T) {
	if Decimal.Start(`10.5695`).TruncForString(2) != `10.56` {
		t.Error()
	}
}

func TestDecimalClass_RoundForString(t *testing.T) {
	if Decimal.Start(`10.5695`).RoundForString(2) != `10.57` {
		t.Error()
	}
}

func TestDecimalClass_RoundUpForString(t *testing.T) {
	if Decimal.Start(`10.5615`).RoundUpForString(2) != `10.57` {
		t.Error()
	}
}

func TestDecimalClass_RoundDownForString(t *testing.T) {
	if Decimal.Start(`10.5695`).RoundDownForString(2) != `10.56` {
		t.Error()
	}
}
