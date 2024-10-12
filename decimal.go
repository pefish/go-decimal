package go_decimal

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	decimal "github.com/pefish/go-decimal/lib"
	go_format "github.com/pefish/go-format"
	"github.com/pkg/errors"
)

type DecimalType struct {
	result decimal.Decimal
}

var Decimal = DecimalType{}

func (d *DecimalType) MustEq(a interface{}) bool {
	b, err := d.Eq(a)
	if err != nil {
		panic(err)
	}
	return b
}

// =
func (d *DecimalType) Eq(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return d.result.Equal(deci), nil
}

func (d *DecimalType) MustNeq(a interface{}) bool {
	b, err := d.Neq(a)
	if err != nil {
		panic(err)
	}
	return b
}

// !=
func (d *DecimalType) Neq(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return !d.result.Equal(deci), nil
}

func (d *DecimalType) MustLt(a interface{}) bool {
	b, err := d.Lt(a)
	if err != nil {
		panic(err)
	}
	return b
}

// <
func (d *DecimalType) Lt(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return d.result.LessThan(deci), nil
}

func (d *DecimalType) MustLte(a interface{}) bool {
	b, err := d.Lte(a)
	if err != nil {
		panic(err)
	}
	return b
}

// <=
func (d *DecimalType) Lte(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return d.result.LessThanOrEqual(deci), nil
}

func (d *DecimalType) MustGt(a interface{}) bool {
	b, err := d.Gt(a)
	if err != nil {
		panic(err)
	}
	return b
}

// >
func (d *DecimalType) Gt(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return d.result.GreaterThan(deci), nil
}

func (d *DecimalType) MustGte(a interface{}) bool {
	b, err := d.Gte(a)
	if err != nil {
		panic(err)
	}
	return b
}

// >=
func (d *DecimalType) Gte(a interface{}) (bool, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return false, err
	}
	return d.result.GreaterThanOrEqual(deci), nil
}

func (d *DecimalType) MustStart(a interface{}) *DecimalType {
	deciClass, err := d.Start(a)
	if err != nil {
		panic(err)
	}
	return deciClass
}

// 开始计算。小数后面有0的话会自动去除
func (d *DecimalType) Start(a interface{}) (*DecimalType, error) {
	decimalInstanceNew := DecimalType{}
	deci, err := decimalInstanceNew.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	decimalInstanceNew.result = deci
	return &decimalInstanceNew, nil
}

// |-1| = 1
func (d *DecimalType) AbsForString() string {
	return d.Abs().result.String()
}

// |-1| = 1
func (d *DecimalType) Abs() *DecimalType {
	d.result = d.result.Abs()
	return d
}

func (d *DecimalType) MustAddForString(a interface{}) string {
	result, err := d.AddForString(a)
	if err != nil {
		panic(err)
	}
	return result
}

// +
func (d *DecimalType) AddForString(a interface{}) (string, error) {
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

func (d *DecimalType) MustAdd(a interface{}) *DecimalType {
	result, err := d.Add(a)
	if err != nil {
		panic(err)
	}
	return result
}

// +
func (d *DecimalType) Add(a interface{}) (*DecimalType, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Add(deci)
	return d, nil
}

func (d *DecimalType) MustSubForString(a interface{}) string {
	result, err := d.SubForString(a)
	if err != nil {
		panic(err)
	}
	return result
}

// -
func (d *DecimalType) SubForString(a interface{}) (string, error) {
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

func (d *DecimalType) MustSub(a interface{}) *DecimalType {
	result, err := d.Sub(a)
	if err != nil {
		panic(err)
	}
	return result
}

// -
func (d *DecimalType) Sub(a interface{}) (*DecimalType, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Sub(deci)
	return d, nil
}

func (d *DecimalType) MustDivForString(a interface{}) string {
	result, err := d.DivForString(a)
	if err != nil {
		panic(err)
	}
	return result
}

// /
func (d *DecimalType) DivForString(a interface{}) (string, error) {
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

func (d *DecimalType) MustDiv(a interface{}) *DecimalType {
	result, err := d.Div(a)
	if err != nil {
		panic(err)
	}
	return result
}

// /
func (d *DecimalType) Div(a interface{}) (*DecimalType, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Div(deci)
	return d, nil
}

func (d *DecimalType) MustShiftedBy(a interface{}) *DecimalType {
	result, err := d.ShiftedBy(a)
	if err != nil {
		panic(err)
	}
	return result
}

// * 10^x
func (d *DecimalType) ShiftedBy(a interface{}) (*DecimalType, error) {
	int32_, err := go_format.ToInt32(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Shift(int32_)
	return d, nil
}

func (d *DecimalType) MustUnShiftedBy(a interface{}) *DecimalType {
	result, err := d.UnShiftedBy(a)
	if err != nil {
		panic(err)
	}
	return result
}

// / 10^x
func (d *DecimalType) UnShiftedBy(a interface{}) (*DecimalType, error) {
	int32_, err := go_format.ToInt32(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Shift(-int32_)
	return d, nil
}

func (d *DecimalType) MustMultiForString(a interface{}) string {
	result, err := d.MultiForString(a)
	if err != nil {
		panic(err)
	}
	return result
}

// *
func (d *DecimalType) MultiForString(a interface{}) (string, error) {
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

func (d *DecimalType) MustMulti(a interface{}) *DecimalType {
	result, err := d.Multi(a)
	if err != nil {
		panic(err)
	}
	return result
}

// *
func (d *DecimalType) Multi(a interface{}) (*DecimalType, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Mul(deci)
	return d, nil
}

func (d *DecimalType) MustPow(a interface{}) *DecimalType {
	result, err := d.Pow(a)
	if err != nil {
		panic(err)
	}
	return result
}

// ^x
func (d *DecimalType) Pow(a interface{}) (*DecimalType, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Pow(deci)
	return d, nil
}

func (d *DecimalType) MustMod(a interface{}) *DecimalType {
	result, err := d.Mod(a)
	if err != nil {
		panic(err)
	}
	return result
}

// % x
func (d *DecimalType) Mod(a interface{}) (*DecimalType, error) {
	deci, err := d.interfaceToDecimal(a)
	if err != nil {
		return nil, err
	}
	d.result = d.result.Mod(deci)
	return d, nil
}

// 根号
func (d *DecimalType) Sqrt() (*DecimalType, error) {
	result, err := d.result.Sqrt()
	if err != nil {
		return nil, err
	}
	d.result = result
	return d, nil
}

func (d *DecimalType) MustSqrt() *DecimalType {
	result, err := d.Sqrt()
	if err != nil {
		panic(err)
	}
	return result
}

func (d *DecimalType) End() decimal.Decimal {
	return d.result
}

func (d *DecimalType) EndForString() string {
	return d.result.String()
}

func (d *DecimalType) MustEndForBigInt() *big.Int {
	r, err := d.EndForBigInt()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalType) EndForBigInt() (*big.Int, error) {
	result, ok := new(big.Int).SetString(d.result.String(), 10)
	if !ok {
		return nil, errors.New(fmt.Sprintf("String %s to bigInt error", d.result.String()))
	}
	return result, nil
}

func (d *DecimalType) MustEndForBigFloat() *big.Float {
	r, err := d.EndForBigFloat()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalType) EndForBigFloat() (*big.Float, error) {
	result, ok := new(big.Float).SetString(d.result.String())
	if !ok {
		return nil, errors.New(fmt.Sprintf("String %s to bigInt error", d.result.String()))
	}
	return result, nil
}

func (d *DecimalType) MustEndForUint64() uint64 {
	r, err := d.EndForUint64()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalType) EndForUint64() (uint64, error) {
	return go_format.ToUint64(d.result.String())
}

func (d *DecimalType) MustEndForInt64() int64 {
	r, err := d.EndForInt64()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalType) EndForInt64() (int64, error) {
	return go_format.ToInt64(d.result.String())
}

func (d *DecimalType) MustEndForInt() int {
	r, err := d.EndForInt()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalType) EndForInt() (int, error) {
	return go_format.ToInt(d.result.String())
}

func (d *DecimalType) MustEndForFloat64() float64 {
	r, err := d.EndForFloat64()
	if err != nil {
		panic(err)
	}
	return r
}

func (d *DecimalType) EndForFloat64() (float64, error) {
	return go_format.ToFloat64(d.result.String())
}

// 直接截取
func (d *DecimalType) TruncForString(precision int32) string {
	return d.Trunc(precision).result.String()
}

// 直接截取
func (d *DecimalType) Trunc(precision int32) *DecimalType {
	d.result = d.result.Truncate(precision)
	return d
}

// 四舍五入
func (d *DecimalType) RoundForString(precision int32) string {
	return d.Round(precision).result.String()
}

// 四舍五入（后面保留0）
func (d *DecimalType) RoundForRemainZeroString(precision int32) string {
	return d.Round(precision).result.StringRemain()
}

// 四舍五入
func (d *DecimalType) Round(precision int32) *DecimalType {
	d.result = d.result.Round(precision)
	return d
}

// 向上截取
func (d *DecimalType) RoundUpForString(precision int32) string {
	return d.RoundUp(precision).result.String()
}

// 向上截取
func (d *DecimalType) RoundUp(precision int32) *DecimalType {
	if d.result.Round(precision).Equal(d.result) {
		return d
	}

	halfPrecision := decimal.New(5, -precision-1)

	d.result = d.result.Add(halfPrecision).Round(precision)
	return d
}

// 向下截取
func (d *DecimalType) RoundDownForString(precision int32) string {
	return d.RoundDown(precision).result.String()
}

// 向下截取
func (d *DecimalType) RoundDown(precision int32) *DecimalType {
	if d.result.Round(precision).Equal(d.result) {
		return d
	}

	halfPrecision := decimal.New(5, -precision-1)

	d.result = d.result.Sub(halfPrecision).Round(precision)
	return d
}

func (d *DecimalType) interfaceToDecimal(a interface{}) (decimal.Decimal, error) {
	if inst, ok := a.(decimal.Decimal); ok {
		return inst, nil
	}
	if inst, ok := a.(*DecimalType); ok {
		return inst.result, nil
	}
	if inst, ok := a.(DecimalType); ok {
		return inst.result, nil
	}
	str := ""
	if inst, ok := a.(*big.Int); ok {
		str = inst.String()
	} else if inst, ok := a.(*big.Float); ok {
		str = inst.String()
	} else {
		str = go_format.ToString(a)
	}

	if strings.HasPrefix(str, "0x") || strings.HasPrefix(str, "0X") {
		r, err := strconv.ParseUint(str[2:], 16, 64)
		if err != nil {
			return decimal.Decimal{}, err
		}
		str = go_format.ToString(r)
	} else if strings.HasPrefix(str, "0o") || strings.HasPrefix(str, "0O") {
		r, err := strconv.ParseUint(str[2:], 8, 64)
		if err != nil {
			return decimal.Decimal{}, err
		}
		str = go_format.ToString(r)
	} else if strings.HasPrefix(str, "0b") || strings.HasPrefix(str, "0B") {
		r, err := strconv.ParseUint(str[2:], 2, 64)
		if err != nil {
			return decimal.Decimal{}, err
		}
		str = go_format.ToString(r)
	}

	decimal_, err := decimal.NewFromString(str)
	if err != nil {
		return decimal.Decimal{}, err
	}
	return decimal_, nil
}

// 判断小数的精度是不是指定精度
func (d *DecimalType) IsPrecision(precision int32) bool {
	return d.GetPrecision() == precision
}

func (d *DecimalType) GetPrecision() int32 {
	splitAmount := strings.Split(d.result.String(), `.`)
	if len(splitAmount) <= 1 {
		return 0
	} else {
		return int32(len(splitAmount[1]))
	}
}
