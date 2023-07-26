package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"github.com/samarec1812/currency-rates/internal/models"
)

type App interface {
	Run() error
}

type app struct {
	options models.Options
}

func NewApp(date string, code string) App {
	return &app{
		options: models.Options{
			Date: date,
			Code: code,
		},
	}
}

func (a app) Run() error {

	bytes, err := makeReq(dateProcessing(a.options.Date))
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}

	var curse models.Currency
	err = json.Unmarshal(bytes, &curse)
	if err != nil {
		return fmt.Errorf("unmarshalling error: %w", err)
	}

	res, err := show(curse.Valute, a.options.Code)
	if err != nil {
		return fmt.Errorf("request error: %w", err)
	}
	fmt.Println(res)
	return nil
}

func dateProcessing(date string) string {
	return strings.ReplaceAll(date, "-", "/")
}

func makeReq(date string) ([]byte, error) {

	res, err := http.Get(fmt.Sprintf("https://www.cbr-xml-daily.ru/archive/%s/daily_json.js", date))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func show(s interface{}, code string) (string, error) {
	res := ""
	numField := reflect.ValueOf(s).NumField()

	for i := 0; i < numField; i++ {
		if reflect.TypeOf(s).Field(i).Tag.Get("json") == code {
			searchStruct := reflect.ValueOf(s).Field(i)

			res = fmt.Sprintf("%s (%s): %f", searchStruct.FieldByName("CharCode"), searchStruct.FieldByName("Name"), searchStruct.FieldByName("Value").Float())
			return res, nil
		}
	}

	return "", fmt.Errorf("currency not found")
}
