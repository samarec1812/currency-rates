package models

import "time"

// Currency autogenerate code
type Currency struct {
	Date   time.Time `json:"Date"`
	Valute struct {
		Aud struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"AUD"`
		Azn struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"AZN"`
		Gbp struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"GBP"`
		Amd struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"AMD"`
		Byn struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"BYN"`
		Bgn struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"BGN"`
		Brl struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"BRL"`
		Huf struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"HUF"`
		Vnd struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"VND"`
		Hkd struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"HKD"`
		Gel struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"GEL"`
		Dkk struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"DKK"`
		Aed struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"AED"`
		Usd struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"USD"`
		Eur struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"EUR"`
		Egp struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"EGP"`
		Inr struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"INR"`
		Idr struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"IDR"`
		Kzt struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"KZT"`
		Cad struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"CAD"`
		Qar struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"QAR"`
		Kgs struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"KGS"`
		Cny struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"CNY"`
		Mdl struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"MDL"`
		Nzd struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"NZD"`
		Nok struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"NOK"`
		Pln struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"PLN"`
		Ron struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"RON"`
		Xdr struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"XDR"`
		Sgd struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"SGD"`
		Tjs struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"TJS"`
		Thb struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"THB"`
		Try struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"TRY"`
		Tmt struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"TMT"`
		Uzs struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"UZS"`
		Uah struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"UAH"`
		Czk struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"CZK"`
		Sek struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"SEK"`
		Chf struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"CHF"`
		Rsd struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"RSD"`
		Zar struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"ZAR"`
		Krw struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"KRW"`
		Jpy struct {
			CharCode string  `json:"CharCode"`
			Name     string  `json:"Name"`
			Value    float64 `json:"Value"`
		} `json:"JPY"`
	} `json:"Valute"`
}
