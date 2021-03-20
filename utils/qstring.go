package utils

import (
	"errors"
	"net/url"
)

var errNoNames = errors.New("No Names Provided")

type MarshalInterfaceTest struct {
	Names []string
}

func (u *MarshalInterfaceTest) UnmarshalQuery(v url.Values) error {
	var ok bool
	if u.Names, ok = v["names"]; ok {
		return nil
	}
	return errNoNames
}

type fairPlayRequest struct {
	ContentID          string `qstring:"r"`
	ContentType        string `qstring:"t"`
	CryptoPeriodNumber uint   `qstring:"p"`
}
