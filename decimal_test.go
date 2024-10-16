package go_decimal

import (
	"math/big"
	"testing"

	go_test_ "github.com/pefish/go-test"
)

func TestDecimalClass_Lt(t *testing.T) {
	if Decimal.MustStart(`12`).MustLt(10) {
		t.Error()
	}
}

func TestDecimalClass_AbsForString(t *testing.T) {
	if Decimal.MustStart(`-1`).AbsForString() != `1` {
		t.Error()
	}
}

func TestDecimalClass_AddForString(t *testing.T) {
	if Decimal.MustStart(`-1`).MustAddForString("-2.407990391483160877") != `-3.407990391483160877` {
		t.Error()
	}

	a := 1
	if Decimal.MustStart(&a).MustAddForString(2) != "3" {
		t.Error()
	}

	if Decimal.MustStart("0x12").MustAddForString(2) != "20" {
		t.Error()
	}

	if Decimal.MustStart("0o12").MustAddForString(2) != "12" {
		t.Error()
	}

	if Decimal.MustStart("0b11").MustAddForString(2) != "5" {
		t.Error()
	}
}

func TestDecimalClass_SubForString(t *testing.T) {
	if Decimal.MustStart(`-1`).MustSubForString(1) != `-2` {
		t.Error(Decimal.MustStart(`-1`).MustSubForString(1))
	}
}

func TestDecimalClass_DivForString(t *testing.T) {
	if Decimal.MustStart(`10`).MustDivForString(2) != `5` {
		t.Error()
	}
}

func TestDecimalClass_MultiForString(t *testing.T) {
	if Decimal.MustStart(`10`).MustMultiForString(2) != `20` {
		t.Error()
	}
}

func TestDecimalClass_TruncForString(t *testing.T) {
	if Decimal.MustStart(`10.5695`).TruncForString(2) != `10.56` {
		t.Error()
	}
}

func TestDecimalClass_RoundForString(t *testing.T) {
	if Decimal.MustStart(`10.5695`).RoundForString(2) != `10.57` {
		t.Error()
	}
}

func TestDecimalClass_RoundUpForString(t *testing.T) {
	if Decimal.MustStart(`10.5615`).RoundUpForString(2) != `10.57` {
		t.Error()
	}
}

func TestDecimalClass_RoundDownForString(t *testing.T) {
	if Decimal.MustStart(`10.5695`).RoundDownForString(2) != `10.56` {
		t.Error()
	}

	if Decimal.MustStart(`10.0000`).RoundDownForString(2) != `10` {
		t.Error()
	}
}

func TestDecimalClass_RoundForRemainZeroString(t *testing.T) {
	if Decimal.MustStart(`10.5`).RoundForRemainZeroString(4) != `10.5000` {
		t.Error()
	}

	if Decimal.MustStart(`10.01`).RoundForRemainZeroString(1) != `10.0` {
		t.Error()
	}
}

func TestDecimalClass_GetPrecision(t *testing.T) {
	if Decimal.MustStart(`10.5`).GetPrecision() != 1 {
		t.Error()
	}

	if Decimal.MustStart(`10.5000`).GetPrecision() != 1 {
		t.Error()
	}

	if Decimal.MustStart(`10.0005`).GetPrecision() != 4 {
		t.Error()
	}
}

func TestDecimalClass_IsPrecision(t *testing.T) {
	if !Decimal.MustStart(`10.5`).IsPrecision(1) {
		t.Error()
	}

	if !Decimal.MustStart(`10.5000`).IsPrecision(1) {
		t.Error()
	}

	if !Decimal.MustStart(`10`).IsPrecision(0) {
		t.Error()
	}
}

func TestDecimalClass_ShiftedBy(t *testing.T) {
	if Decimal.MustStart(`10.5`).MustShiftedBy(1).EndForString() != `105` {
		t.Error()
	}

	if Decimal.MustStart(`10.5`).MustShiftedBy(5).EndForString() != `1050000` {
		t.Error()
	}

	if Decimal.MustStart(`10000`).MustShiftedBy(18).EndForString() != `10000000000000000000000` {
		t.Error()
	}
}

func TestDecimalClass_UnShiftedBy(t *testing.T) {
	if Decimal.MustStart(`10.5`).MustUnShiftedBy(1).EndForString() != `1.05` {
		t.Error()
	}

	if Decimal.MustStart(`1050000`).MustUnShiftedBy(5).EndForString() != `10.5` {
		t.Error()
	}

	if Decimal.MustStart(`10500000000000000`).MustUnShiftedBy(15).EndForString() != `10.5` {
		t.Error()
	}

	if Decimal.MustStart(`1`).MustUnShiftedBy(0).EndForString() != `1` {
		t.Error()
	}
}

func TestDecimalClass_Eq(t *testing.T) {
	go_test_.Equal(t, true, Decimal.MustStart(`-1`).MustEq(-1))
}

func TestDecimalClass_EndForString(t *testing.T) {
	go_test_.Equal(t, "57.78", Decimal.MustStart("57.7800").EndForString())
}

func TestDecimalClass_Pow(t *testing.T) {
	go_test_.Equal(t, "144", Decimal.MustStart(12).MustPow(2).EndForString())
}

func TestDecimalClass_Mod(t *testing.T) {
	go_test_.Equal(t, "2", Decimal.MustStart(12).MustMod(10).EndForString())
}

func TestDecimalClass_Sqrt(t *testing.T) {
	result, err := Decimal.MustStart(4).Sqrt()
	go_test_.Equal(t, nil, err)
	go_test_.Equal(t, "2", result.EndForString())

	_, err = Decimal.MustStart(-4).Sqrt()
	go_test_.Equal(t, "The result of this operation is imaginary.", err.Error())
}

func TestDecimalClass_MustEndForBigFloat(t *testing.T) {
	a := big.NewFloat(1.1)
	result := Decimal.MustStart(a).MustAdd(1.2).EndForString()
	go_test_.Equal(t, "2.3", result)

	result1 := Decimal.MustStart(a).MustAdd(1.2).MustEndForBigFloat()
	go_test_.Equal(t, "2.3", result1.String())
}
