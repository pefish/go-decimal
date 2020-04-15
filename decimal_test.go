package go_decimal

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

func TestDecimalClass_RoundForRemainZeroString(t *testing.T) {
	if Decimal.Start(`10.5`).RoundForRemainZeroString(4) != `10.5000` {
		t.Error()
	}

	if Decimal.Start(`10.01`).RoundForRemainZeroString(1) != `10.0` {
		t.Error()
	}
}

func TestDecimalClass_GetPrecision(t *testing.T) {
	if Decimal.Start(`10.5`).GetPrecision() != 1 {
		t.Error()
	}

	if Decimal.Start(`10.5000`).GetPrecision() != 1 {
		t.Error()
	}

	if Decimal.Start(`10.0005`).GetPrecision() != 4 {
		t.Error()
	}
}

func TestDecimalClass_IsPrecision(t *testing.T) {
	if !Decimal.Start(`10.5`).IsPrecision(1) {
		t.Error()
	}

	if !Decimal.Start(`10.5000`).IsPrecision(1) {
		t.Error()
	}

	if !Decimal.Start(`10`).IsPrecision(0) {
		t.Error()
	}
}

func TestDecimalClass_ShiftedBy(t *testing.T) {
	if Decimal.Start(`10.5`).ShiftedBy(1).EndForString() != `105` {
		t.Error()
	}

	if Decimal.Start(`10.5`).ShiftedBy(5).EndForString() != `1050000` {
		t.Error()
	}

	if Decimal.Start(`10000`).ShiftedBy(18).EndForString() != `10000000000000000000000` {
		t.Error()
	}
}

func TestDecimalClass_UnShiftedBy(t *testing.T) {
	if Decimal.Start(`10.5`).UnShiftedBy(1).EndForString() != `1.05` {
		t.Error()
	}

	if Decimal.Start(`1050000`).UnShiftedBy(5).EndForString() != `10.5` {
		t.Error()
	}

	if Decimal.Start(`10500000000000000`).UnShiftedBy(15).EndForString() != `10.5` {
		t.Error()
	}

	if Decimal.Start(`1`).UnShiftedBy(0).EndForString() != `1` {
		t.Error()
	}
}

func TestDecimalClass_Eq(t *testing.T) {
	if !Decimal.Start(`-1`).Eq(-1) {
		t.Error()
	}
}

func TestDecimalClass_EndForString(t *testing.T) {
	a := `57.7800`
	if Decimal.Start(a).EndForString() != `57.78` {
		t.Error(Decimal.Start(a).EndForString())
	}
}