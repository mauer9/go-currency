package transport

import (
	"errors"
	"regexp"
)

func (r *SaveCurrenciesRequest) Validate() error {
	dateFormat := regexp.MustCompile(`^\d{2}\.\d{2}\.\d{4}$`)

	if !dateFormat.MatchString(r.Date) {
		return errors.New("incorrect date format")
	}

	return nil
}

func (r *GetCurrencyByDateAndCodeRequest) Validate() error {
	dateFormat := regexp.MustCompile(`^\d{2}\.\d{2}\.\d{4}$`)

	if !dateFormat.MatchString(r.Date) {
		return errors.New("incorrect date format")
	}

	if r.Code == "" {
		return errors.New("incorrect request")
	}

	return nil
}
