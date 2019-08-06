package go_decimal

import (
	"github.com/pefish/go-decimal/lib"
	"github.com/pefish/go-reflect"
	"strings"
)

type DecimalClass struct {
	result decimal.Decimal
}

var Decimal = DecimalClass{}

// =
func (this *DecimalClass) Eq(a interface{}) bool {
	return this.result.Equal(this.interfaceToDecimal(a))
}

// !=
func (this *DecimalClass) Neq(a interface{}) bool {
	return !this.result.Equal(this.interfaceToDecimal(a))
}

// <
func (this *DecimalClass) Lt(a interface{}) bool {
	return this.result.LessThan(this.interfaceToDecimal(a))
}

// <=
func (this *DecimalClass) Lte(a interface{}) bool {
	return this.result.LessThanOrEqual(this.interfaceToDecimal(a))
}

// >
func (this *DecimalClass) Gt(a interface{}) bool {
	return this.result.GreaterThan(this.interfaceToDecimal(a))
}

// >=
func (this *DecimalClass) Gte(a interface{}) bool {
	return this.result.GreaterThanOrEqual(this.interfaceToDecimal(a))
}

// 开始计算
func (this *DecimalClass) Start(a interface{}) *DecimalClass {
	decimalInstance := DecimalClass{}
	decimalInstance.result = decimalInstance.interfaceToDecimal(a)
	return &decimalInstance
}

// \-1\ = 1
func (this *DecimalClass) AbsForString() string {
	return this.Abs().result.String()
}

// \-1\ = 1
func (this *DecimalClass) Abs() *DecimalClass {
	this.result = this.result.Abs()
	return this
}

// +
func (this *DecimalClass) AddForString(a interface{}) string {
	return this.Add(this.interfaceToDecimal(a)).result.String()
}

// +
func (this *DecimalClass) Add(a interface{}) *DecimalClass {
	this.result = this.result.Add(this.interfaceToDecimal(a))
	return this
}

// -
func (this *DecimalClass) SubForString(a interface{}) string {
	return this.Sub(this.interfaceToDecimal(a)).result.String()
}

// -
func (this *DecimalClass) Sub(a interface{}) *DecimalClass {
	this.result = this.result.Sub(this.interfaceToDecimal(a))
	return this
}

// /
func (this *DecimalClass) DivForString(a interface{}) string {
	return this.Div(this.interfaceToDecimal(a)).result.String()
}

// /
func (this *DecimalClass) Div(a interface{}) *DecimalClass {
	this.result = this.result.Div(this.interfaceToDecimal(a))
	return this
}

// *
func (this *DecimalClass) MultiForString(a interface{}) string {
	return this.Multi(this.interfaceToDecimal(a)).result.String()
}

// *
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

// 直接截取
func (this *DecimalClass) TruncForString(precision int32) string {
	return this.Trunc(precision).result.String()
}

// 直接截取
func (this *DecimalClass) Trunc(precision int32) *DecimalClass {
	this.result = this.result.Truncate(precision)
	return this
}

// 四舍五入
func (this *DecimalClass) RoundForString(precision int32) string {
	return this.Round(precision).result.String()
}

// 四舍五入（后面保留0）
func (this *DecimalClass) RoundForRemainString(precision int32) string {
	return this.Round(precision).result.StringRemain()
}

// 四舍五入
func (this *DecimalClass) Round(precision int32) *DecimalClass {
	this.result = this.result.Round(precision)
	return this
}

// 向上截取
func (this *DecimalClass) RoundUpForString(precision int32) string {
	return this.RoundUp(precision).result.String()
}

// 向上截取
func (this *DecimalClass) RoundUp(precision int32) *DecimalClass {
	if this.result.Round(precision).Equal(this.result) {
		return this
	}

	halfPrecision := decimal.New(5, -precision-1)

	this.result = this.result.Add(halfPrecision).Round(precision)
	return this
}

// 向下截取
func (this *DecimalClass) RoundDownForString(precision int32) string {
	return this.RoundDown(precision).result.String()
}

// 向下截取
func (this *DecimalClass) RoundDown(precision int32) *DecimalClass {
	if this.result.Round(precision).Equal(this.result) {
		return this
	}

	halfPrecision := decimal.New(5, -precision-1)

	this.result = this.result.Sub(halfPrecision).Round(precision)
	return this
}

func (this *DecimalClass) interfaceToDecimal(a interface{}) decimal.Decimal {
	if inst, ok := a.(decimal.Decimal); ok {
		return inst
	}

	decimal_, err := decimal.NewFromString(go_reflect.Reflect.ToString(a))
	if err != nil {
		panic(err)
	}
	return decimal_
}

// 判断小数的精度是不是指定精度
func (this *DecimalClass) IsPrecision(precision int32) bool {
	splitAmount := strings.Split(this.result.String(), `.`)
	if len(splitAmount) <= 1 {
		return precision == 0
	} else {
		return len(splitAmount[1]) == int(precision)
	}
}
