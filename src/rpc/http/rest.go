package http

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"strconv"
	"test/src/domain"
	"time"
)

func (r *RPC) GetCurrencyRate(_ context.Context, date string) ([]domain.Currency, error) {
	url := "https://nationalbank.kz/rss/get_rates.cfm?fdate=" + date
	var crs []domain.Currency

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rates rates
	err = xml.Unmarshal(body, &rates)
	if err != nil {
		return nil, err
	}

	for _, r := range rates.Items {
		val, err := strconv.ParseFloat(r.Description, 64)
		if err != nil {
			return nil, err
		}

		crs = append(crs, domain.Currency{
			Title: r.FullName,
			Code:  r.Title,
			Value: val,
			ADate: date,
		})
	}

	return crs, nil
}
