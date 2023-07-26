package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/samarec1812/currency-rates/internal/app"
)

// --code=USD --date=2022-10-08
type Options struct {
	Code string
	Date string
}

func (opt Options) Validate() error {
	if opt.Code == "" {
		return errors.New("no valid code")
	}

	if opt.Date == "" {
		return errors.New("no valid date")
	}

	if _, err := time.Parse(time.DateOnly, opt.Date); err != nil {
		return errors.New("no valid date")
	}

	return nil
}

func ParseFlags() (*Options, error) {
	var opts Options

	flag.StringVar(&opts.Code, "code", "USD", "currency code. by default - USD")
	flag.StringVar(&opts.Date, "date", "", "currency exchange rate date. by default - empty string")

	flag.Parse()

	return &opts, nil
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
	a := app.NewApp(opts.Date, opts.Code)
	if err = a.Run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
