package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const defaultUrlActiveTimeInHours = time.Hour * 24

var (
	ErrConvertingExpirationDateToRFC3339 = errors.New("could not convert expiration date to RFC3339")
	ErrConvertingCreatedDateToRFC3339    = errors.New("could not convert created date to RFC3339")

	alphabet = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	base = int64(len(alphabet))
)

type Url struct {
	Original       string
	Short          string
	ExpirationDate string
	DateCreated    string
}

func NewUrl(original string, expirationDate string) (*Url, error) {
	// TODO: Add UTC check how
	createdDate := time.Now().UTC()

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

	_, err := time.Parse(time.RFC3339Nano, createdDate.Format(time.RFC3339Nano))
	if err != nil {
		return nil, ErrConvertingCreatedDateToRFC3339
	}

	return &Url{
		Original:       original,
		Short:          Shorten(original),
		DateCreated:    createdDate.Format(time.RFC3339Nano),
		ExpirationDate: expDate.Format(time.RFC3339Nano),
	}, nil
}

func Shorten(origUrl string) string {
	hash := sha256.Sum256([]byte(origUrl))
	hexStr := hex.EncodeToString([]byte(fmt.Sprintf("%s", hash[:5])))
	num, _ := strconv.ParseInt(hexStr, 16, 64)
	encoded := decimalToBase62(num)
	return encoded
}

func decimalToBase62(strDecimal int64) string {
	encoded := ""
	for strDecimal > 0 {
		r := strDecimal % base
		strDecimal /= base
		encoded = fmt.Sprintf("%s%s", alphabet[r], encoded)
	}

	return encoded
}
