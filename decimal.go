package go_decimal

import (
	"fmt"
	"github.com/pefish/go-decimal/lib"
	go_format "github.com/pefish/go-format"
	"github.com/pkg/errors"
	"math/big"
	"strings"
)

type DecimalClass struct {
	result decimal.Decimal
}

var Decimal = DecimalClass{}

func (d *DecimalClass) MustEq(a interface{}) bool {
	b, err := d.Eq(a)
	if err != nil {
		panic(err)
	}
	return b
}

// =
func (d *DecimalClass) Eq(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return d.result.Equal(deci), nil
}

func (d *DecimalClass) MustNeq(a interface{}) bool {
	b, err := d.Neq(a)
	if err != nil {
		panic(err)
	}
	return b
}

// !=
func (d *DecimalClass) Neq(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return !d.result.Equal(deci), nil
}

func (d *DecimalClass) MustLt(a interface{}) bool {
	b, err := d.Lt(a)
	if err != nil {
		panic(err)
	}
	return b
}

// <
func (d *DecimalClass) Lt(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return d.result.LessThan(deci), nil
}

func (d *DecimalClass) MustLte(a interface{}) bool {
	b, err := d.Lte(a)
	if err != nil {
		panic(err)
	}
	return b
}

// <=
func (d *DecimalClass) Lte(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return d.result.LessThanOrEqual(deci), nil
}

func (d *DecimalClass) MustGt(a interface{}) bool {
	b, err := d.Gt(a)
	if err != nil {
		panic(err)
	}
	return b
}

// >
func (d *DecimalClass) Gt(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return d.result.GreaterThan(deci), nil
}

func (d *DecimalClass) MustGte(a interface{}) bool {
	b, err := d.Gte(a)
	if err != nil {
		panic(err)
	}
	return b
}

// >=
func (d *DecimalClass) Gte(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return d.result.GreaterThanOrEqual(deci), nil
}

func (d *DecimalClass) MustStart(a interface{}) *DecimalClass {
	deciClass, err := d.Start(a)
	if err != nil {
		panic(err)
	}
	return deciClass
}

// 开始计算。小数后面有0的话会自动去除
func (d *DecimalClass) Start(a interface{}) (*DecimalClass, error) {
	decimalInstanceNew := DecimalClass{}
	deci, err := decimalInstanceNew.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	decimalInstanceNew.result = deci
	return &decimalInstanceNew, nil
}

// |-1| = 1
func (d *DecimalClass) AbsForString() string {
	return d.Abs().result.String()
}

// |-1| = 1
func (d *DecimalClass) Abs() *DecimalClass {
	d.result = d.result.Abs()
	return d
}

func (d *DecimalClass) MustAddForString(a interface{}) string {
	result, err := d.AddForString(a)
	if err != nil {
		panic(err)
	}
	return result
}

// +
func (d *DecimalClass) AddForString(a interface{}) (string, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return "", err
	}
	deciClass, err := d.Add(deci)
	if err != nil {
		return "", err
	}
	return deciClass.result.String(), nil
}

func (d *DecimalClass) MustAdd(a interface{}) *DecimalClass {
	result, err := d.Add(a)
	if err != nil {
		panic(err)
	}
	return result
}

// +
func (d *DecimalClass) Add(a interface{}) (*DecimalClass, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Add(deci)
	return d, nil
}

func (d *DecimalClass) MustSubForString(a interface{}) string {
	result, err := d.SubForString(a)
	if err != nil {
		panic(err)
	}
	return result
}

// -
func (d *DecimalClass) SubForString(a interface{}) (string, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return "", err
	}
	deciClass, err := d.Sub(deci)
	if err != nil {
		return "", err
	}
	return deciClass.result.String(), nil
}

func (d *DecimalClass) MustSub(a interface{}) *DecimalClass {
	result, err := d.Sub(a)
	if err != nil {
		panic(err)
	}
	return result
}

// -
func (d *DecimalClass) Sub(a interface{}) (*DecimalClass, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Sub(deci)
	return d, nil
}

func (d *DecimalClass) MustDivForString(a interface{}) string {
	result, err := d.DivForString(a)
	if err != nil {
		panic(err)
	}
	return result
}

// /
func (d *DecimalClass) DivForString(a interface{}) (string, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return "", err
	}
	deciClass, err := d.Div(deci)
	if err != nil {
		return "", err
	}
	return deciClass.result.String(), nil
}

func (d *DecimalClass) MustDiv(a interface{}) *DecimalClass {
	result, err := d.Div(a)
	if err != nil {
		panic(err)
	}
	return result
}

// /
func (d *DecimalClass) Div(a interface{}) (*DecimalClass, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Div(deci)
	return d, nil
}

func (d *DecimalClass) MustShiftedBy(a interface{}) *DecimalClass {
	result, err := d.ShiftedBy(a)
	if err != nil {
		panic(err)
	}
	return result
}

// * 10^x
func (d *DecimalClass) ShiftedBy(a interface{}) (*DecimalClass, error) {
	int32_, err := go_format.FormatInstance.ToInt32(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Shift(int32_)
	return d, nil
}

func (d *DecimalClass) MustUnShiftedBy(a interface{}) *DecimalClass {
	result, err := d.UnShiftedBy(a)
	if err != nil {
		panic(err)
	}
	return result
}

// / 10^x
func (d *DecimalClass) UnShiftedBy(a interface{}) (*DecimalClass, error) {
	int32_, err := go_format.FormatInstance.ToInt32(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Shift(-int32_)
	return d, nil
}

func (d *DecimalClass) MustMultiForString(a interface{}) string {
	result, err := d.MultiForString(a)
	if err != nil {
		panic(err)
	}
	return result
}

// *
func (d *DecimalClass) MultiForString(a interface{}) (string, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return "", err
	}
	deciClass, err := d.Multi(deci)
	if err != nil {
		return "", err
	}
	return deciClass.result.String(), nil
}

func (d *DecimalClass) MustMulti(a interface{}) *DecimalClass {
	result, err := d.Multi(a)
	if err != nil {
		panic(err)
	}
	return result
}

// *
func (d *DecimalClass) Multi(a interface{}) (*DecimalClass, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Mul(deci)
	return d, nil
}

func (d *DecimalClass) MustPow(a interface{}) *DecimalClass {
	result, err := d.Pow(a)
	if err != nil {
		panic(err)
	}
	return result
}

// ^x
func (d *DecimalClass) Pow(a interface{}) (*DecimalClass, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Pow(deci)
	return d, nil
}

func (d *DecimalClass) MustMod(a interface{}) *DecimalClass {
	result, err := d.Mod(a)
	if err != nil {
		panic(err)
	}
	return result
}

// % x
func (d *DecimalClass) Mod(a interface{}) (*DecimalClass, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Mod(deci)
	return d, nil
}

// 根号
func (d *DecimalClass) Sqrt() (*DecimalClass, error) {
	result, err := d.result.Sqrt()
	if err != nil {
		return nil, err
	}
	d.result = result
	return d, nil
}

func (d *DecimalClass) MustSqrt() *DecimalClass {
	result, err := d.Sqrt()
	if err != nil {
		panic(err)
	}
	return result
}

func (d *DecimalClass) End() decimal.Decimal {
	return d.result
}

func (d *DecimalClass) EndForString() string {
	return d.result.String()
}

func (d *DecimalClass) MustEndForBigInt() *big.Int {
	r, err := d.EndForBigInt()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalClass) EndForBigInt() (*big.Int, error) {
	result, ok := new(big.Int).SetString(d.result.String(), 10)
	if !ok {
		return nil, errors.New(fmt.Sprintf("string %s to bigInt error", d.result.String()))
	}
	return result, nil
}

func (d *DecimalClass) MustEndForUint64() uint64 {
	r, err := d.EndForUint64()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalClass) EndForUint64() (uint64, error) {
	return go_format.FormatInstance.ToUint64(d.result.String())
}

func (d *DecimalClass) MustEndForInt64() int64 {
	r, err := d.EndForInt64()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalClass) EndForInt64() (int64, error) {
	return go_format.FormatInstance.ToInt64(d.result.String())
}

func (d *DecimalClass) MustEndForInt() int {
	r, err := d.EndForInt()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalClass) EndForInt() (int, error) {
	return go_format.FormatInstance.ToInt(d.result.String())
}

func (d *DecimalClass) MustEndForFloat64() float64 {
	r, err := d.EndForFloat64()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalClass) EndForFloat64() (float64, error) {
	return go_format.FormatInstance.ToFloat64(d.result.String())
}

// 直接截取
func (d *DecimalClass) TruncForString(precision int32) string {
	return d.Trunc(precision).result.String()
}

// 直接截取
func (d *DecimalClass) Trunc(precision int32) *DecimalClass {
	d.result = d.result.Truncate(precision)
	return d
}

// 四舍五入
func (d *DecimalClass) RoundForString(precision int32) string {
	return d.Round(precision).result.String()
}

// 四舍五入（后面保留0）
func (d *DecimalClass) RoundForRemainZeroString(precision int32) string {
	return d.Round(precision).result.StringRemain()
}

// 四舍五入
func (d *DecimalClass) Round(precision int32) *DecimalClass {
	d.result = d.result.Round(precision)
	return d
}

// 向上截取
func (d *DecimalClass) RoundUpForString(precision int32) string {
	return d.RoundUp(precision).result.String()
}

// 向上截取
func (d *DecimalClass) RoundUp(precision int32) *DecimalClass {
	if d.result.Round(precision).Equal(d.result) {
		return d
	}

	halfPrecision := decimal.New(5, -precision-1)

	d.result = d.result.Add(halfPrecision).Round(precision)
	return d
}

// 向下截取
func (d *DecimalClass) RoundDownForString(precision int32) string {
	return d.RoundDown(precision).result.String()
}

// 向下截取
func (d *DecimalClass) RoundDown(precision int32) *DecimalClass {
	if d.result.Round(precision).Equal(d.result) {
		return d
	}

	halfPrecision := decimal.New(5, -precision-1)

	d.result = d.result.Sub(halfPrecision).Round(precision)
	return d
}

func (d *DecimalClass) mustInterfaceToDecimal(a interface{}) decimal.Decimal {
	result, err := d.interfaceToDecimal(a)
	if err != nil {
		panic(err)
	}
	return result
}

func (d *DecimalClass) interfaceToDecimal(a interface{}) (decimal.Decimal, error) {
	if inst, ok := a.(decimal.Decimal); ok {
		return inst, nil
	}
	if inst, ok := a.(*DecimalClass); ok {
		return inst.result, nil
	}
	if inst, ok := a.(DecimalClass); ok {
		return inst.result, nil
	}
	str := ""
	if inst, ok := a.(*big.Int); ok {
		str = inst.String()
	} else {
		str = go_format.FormatInstance.ToString(a)
	}

	decimal_, err := decimal.NewFromString(str)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return decimal_, nil
}

// 判断小数的精度是不是指定精度
func (d *DecimalClass) IsPrecision(precision int32) bool {
	return d.GetPrecision() == precision
}

func (d *DecimalClass) GetPrecision() int32 {
	splitAmount := strings.Split(d.result.String(), `.`)
	if len(splitAmount) <= 1 {
		return 0
	} else {
		return int32(len(splitAmount[1]))
	}
}
