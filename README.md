# go-decimal

## Description
decimal tools

## Test

```golang
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
```
