package p_decimal

import (
	"gitee.com/pefish/p-go-reflect"
	"github.com/shopspring/decimal"
)

type DecimalClass struct {
	result decimal.Decimal
}

var Decimal = DecimalClass{}

func (this *DecimalClass) Lt(a interface{}) bool {
	return this.result.LessThan(this.interfaceToDecimal(a))
}

func (this *DecimalClass) Lte(a interface{}) bool {
	return this.result.LessThanOrEqual(this.interfaceToDecimal(a))
}

func (this *DecimalClass) Gt(a interface{}) bool {
	return this.result.GreaterThan(this.interfaceToDecimal(a))
}

func (this *DecimalClass) Gte(a interface{}) bool {
	return this.result.GreaterThanOrEqual(this.interfaceToDecimal(a))
}

func (this *DecimalClass) Start(a interface{}) *DecimalClass {
	this.result = this.interfaceToDecimal(a)
	return this
}

func (this *DecimalClass) Abs() *DecimalClass {
	this.result = this.result.Abs()
	return this
}

func (this *DecimalClass) Add(a interface{}) *DecimalClass {
	this.result = this.result.Add(this.interfaceToDecimal(a))
	return this
}

func (this *DecimalClass) Sub(a interface{}) *DecimalClass {
	this.result = this.result.Sub(this.interfaceToDecimal(a))
	return this
}

func (this *DecimalClass) Div(a interface{}) *DecimalClass {
	this.result = this.result.Div(this.interfaceToDecimal(a))
	return this
}

func (this *DecimalClass) Multi(a interface{}) *DecimalClass {
	this.result = this.result.Mul(this.interfaceToDecimal(a))
	return this
}

func (this *DecimalClass) End() decimal.Decimal {
	return this.result
}

func (this *DecimalClass) EndForString() string {
	return this.result.String()
}

func (this *DecimalClass) EndWithDecimalRemain(remain int32) string {
	return this.result.StringFixed(remain)
}

func (this *DecimalClass) Trunc(precision int32) *DecimalClass {
	this.result = this.result.Truncate(precision)
	return this
}

func (this *DecimalClass) Round(precision int32) *DecimalClass {
	this.result = this.result.Round(precision)
	return this
}

func (this *DecimalClass) RoundUp(precision int32) *DecimalClass {
	if this.result.Round(precision).Equal(this.result) {
		return this
	}

	halfPrecision := decimal.New(5, -precision-1)

	this.result = this.result.Add(halfPrecision).Round(precision)
	return this
}

func (this *DecimalClass) RoundDown(precision int32) *DecimalClass {
	if this.result.Round(precision).Equal(this.result) {
		return this
	}

	halfPrecision := decimal.New(5, -precision-1)

	this.result = this.result.Sub(halfPrecision).Round(precision)
	return this
}

func (this *DecimalClass) AddForString(a interface{}, b interface{}) string {
	return this.interfaceToDecimal(a).Add(this.interfaceToDecimal(b)).String()
}

func (this *DecimalClass) SubForString(a interface{}, b interface{}) string {
	return this.interfaceToDecimal(a).Sub(this.interfaceToDecimal(b)).String()
}

func (this *DecimalClass) DivForString(a interface{}, b interface{}) string {
	return this.interfaceToDecimal(a).Div(this.interfaceToDecimal(b)).String()
}

func (this *DecimalClass) MultiForString(a interface{}, b interface{}) string {
	return this.interfaceToDecimal(a).Mul(this.interfaceToDecimal(b)).String()
}

func (this *DecimalClass) interfaceToDecimal(a interface{}) decimal.Decimal {
	if inst, ok := a.(decimal.Decimal); ok {
		return inst
	}

	decimal_, err := decimal.NewFromString(p_reflect.Reflect.ToString(a))
	if err != nil {
		panic(err)
	}
	return decimal_
}
