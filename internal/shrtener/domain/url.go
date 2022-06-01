package domain

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const defaultUrlActiveTimeInHours time.Duration = time.Hour * 24

var (
	ErrConvertingExpirationDateToRFC3339 = errors.New("could not convert expiration date to RFC3339")
	ErrConvertingCreatedDateToRFC3339    = errors.New("could not convert created date to RFC3339")

	alphabet = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	base int64 = 62
)

type Url struct {
	Original       string
	Short          string
	ExpirationDate time.Time
	DateCreated    time.Time
}

func NewUrl(original string, expirationDate string) (*Url, error) {
	// TODO: Add UTC check how
	createdDate := time.Now()

	var expDate time.Time
	if expirationDate == "" {
		expDate = createdDate.Add(defaultUrlActiveTimeInHours)
	} else {
		d, err := time.Parse(time.RFC3339Nano, expirationDate)
		if err != nil {
			return nil, ErrConvertingExpirationDateToRFC3339
		}
		expDate = d
	}

	return &Url{
		Original:       original,
		Short:          Shorten(original),
		DateCreated:    createdDate,
		ExpirationDate: expDate,
	}, nil
}

func Shorten(origUrl string) string {
	hexStr := hex.EncodeToString([]byte(origUrl))
	num, _ := strconv.ParseInt(hexStr, 16, 64)
	encoded := decimalToBas62(num)
	return encoded
}

func decimalToBas62(strDecimal int64) string {
	encoded := ""
	for strDecimal > 0 {
		r := strDecimal % base
		strDecimal /= base
		encoded = fmt.Sprintf("%s%s", string(alphabet[r]), encoded)
	}

	return encoded
}
