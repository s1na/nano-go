package nano

import (
	"errors"

	"github.com/shopspring/decimal"
)

// We use a 128 bit integer to represent account balances, this is too large
// to present to the user so we defined a set of SI prefixes to make the
// numbers more accessible and avoid confusion.
// The reference wallet uses Mxrb as a divider.
// Gxrb = 1000000000000000000000000000000000, 10^33
// Mxrb = 1000000000000000000000000000000, 10^30
// kxrb = 1000000000000000000000000000, 10^27
// xrb = 1000000000000000000000000, 10^24
// mxrb = 1000000000000000000000, 10^21
// uxrb = 1000000000000000000, 10^18
//
// 1 Mxrb used to be also called 1 Mrai
//
// 1 xrb is 10^24 raw
//
// 1 raw is the smallest possible division
var (
	ratios = map[string]decimal.Decimal{
		"Gxrb": decimal.New(1, 33),
		"Mxrb": decimal.New(1, 30),
		"Mrai": decimal.New(1, 30),
		"kxrb": decimal.New(1, 27),
		"krai": decimal.New(1, 27),
		"xrb":  decimal.New(1, 24),
		"mxrb": decimal.New(1, 21),
		"uxrb": decimal.New(1, 18),
		"raw":  decimal.New(1, 0),
	}
)

func Convert(value, from, to string) (string, error) {
	fr, exists := ratios[from]
	if !exists {
		return "", errors.New("From ratio is invalid")
	}

	tr, exists := ratios[to]
	if !exists {
		return "", errors.New("To ratio is invalid")
	}

	v, err := decimal.NewFromString(value)
	if err != nil {
		return "", err
	}

	return v.Mul(fr.Div(tr)).String(), nil
}

func MraiToRaw(value string) (string, error) {
	v, err := decimal.NewFromString(value)
	if err != nil {
		return "", err
	}

	return v.Mul(ratios["Mxrb"]).String(), nil
}

func MraiFromRaw(value string) (string, error) {
	v, err := decimal.NewFromString(value)
	if err != nil {
		return "", err
	}

	return v.Div(ratios["Mxrb"]).String(), nil
}

func KraiToRaw(value string) (string, error) {
	v, err := decimal.NewFromString(value)
	if err != nil {
		return "", err
	}

	return v.Mul(ratios["kxrb"]).String(), nil
}

func KraiFromRaw(value string) (string, error) {
	v, err := decimal.NewFromString(value)
	if err != nil {
		return "", err
	}

	return v.Div(ratios["kxrb"]).String(), nil
}

func RaiFromRaw(value string) (string, error) {
	v, err := decimal.NewFromString(value)
	if err != nil {
		return "", err
	}

	return v.Mul(ratios["xrb"]).String(), nil
}

func RaiToRaw(value string) (string, error) {
	v, err := decimal.NewFromString(value)
	if err != nil {
		return "", err
	}

	return v.Div(ratios["xrb"]).String(), nil
}
