package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/samarec1812/currency-rates/models"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

// --code=USD --date=2022-10-08
type Options struct {
	Code string
	Date string
}

func (opt Options) Validate() error {
	if opt.Code == "" {
		return fmt.Errorf("no valid code")
	}

	if opt.Date == "" {
		return fmt.Errorf("no valid date")
	}

	if _, err := time.Parse(time.DateOnly, opt.Date); err != nil {
		return fmt.Errorf("no valid date")
	}

	return nil
}

func ParseFlags() (*Options, error) {
	var opts Options

	flag.StringVar(&opts.Code, "code", "USD", "currency code. by default - USD")
	flag.StringVar(&opts.Date, "date", "", "currency exchange rate date. by default - stdout")

	flag.Parse()

	return &opts, nil
}

func makeReq(date string) ([]byte, error) {

	URL := fmt.Sprintf("https://www.cbr-xml-daily.ru/archive/%s/daily_json.js", date) ///  + params.Encode())
	res, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return resBody, nil
}

func DateProcessing(date string) string {
	return strings.ReplaceAll(date, "-", "/")
}

func main() {
	opts, err := ParseFlags()
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "can not parse flags:", err)
		os.Exit(1)
	}

	if err = opts.Validate(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "validation error:", err)
		os.Exit(1)
	}

	bytes, err := makeReq(DateProcessing(opts.Date))
	if err != nil {
		log.Fatalf(err.Error())
	}

	var curse models.Currency
	err = json.Unmarshal(bytes, &curse)
	if err != nil {
		log.Fatalf(err.Error())
	}

	//fmt.Println(curse.Valute)
	//
	fmt.Println(Show(curse.Valute, opts.Code))
}

func Show(s interface{}, code string) (res string) {
	//a := reflect.ValueOf(s)
	numField := reflect.ValueOf(s).NumField()
	for i := 0; i < numField; i++ {
		if reflect.TypeOf(s).Field(i).Tag.Get("json") == code {
			//USD (Доллар США): 61,2475
			searchStruct := reflect.ValueOf(s).Field(i)

			res = fmt.Sprintf("%s (%s): %f", searchStruct.FieldByName("CharCode"), searchStruct.FieldByName("Name"), searchStruct.FieldByName("Value").Float())
			break
		}
	}
	return
}
