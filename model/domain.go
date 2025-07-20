package model

import (
	"fmt"
	"strconv"
	"time"
)

type OHCL struct {
}

type Candle struct {
	Time   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	VWAP   float64
	Volume float64
	Count  int
}

type TOC struct {
	Time  time.Time
	Open  float64
	Close float64
}

// TODO return type
func Trend(tocs []TOC) {
	for _, t := range tocs {
		fmt.Printf("Time: %s, trend: %.2f\n", t.Time, t.Open-t.Close)
	}
}

func (c Candle) ToTOC() TOC {
	return TOC{
		Time:  c.Time,
		Open:  c.Open,
		Close: c.Close,
	}
}

func ParseCandles(data [][]interface{}) ([]Candle, error) {
	var candles []Candle

	for _, c := range data {
		if len(c) < 8 {
			return nil, fmt.Errorf("niepełny rekord świecy")
		}

		tsFloat, ok := c[0].(float64)
		if !ok {
			return nil, fmt.Errorf("timestamp nie jest float64")
		}
		ts := time.Unix(int64(tsFloat), 0)

		openStr, ok := c[1].(string)
		if !ok {
			return nil, fmt.Errorf("open nie jest string")
		}
		open, err := parseFloat(openStr)
		if err != nil {
			return nil, err
		}

		highStr, ok := c[2].(string)
		if !ok {
			return nil, fmt.Errorf("high nie jest string")
		}
		high, err := parseFloat(highStr)
		if err != nil {
			return nil, err
		}

		lowStr, ok := c[3].(string)
		if !ok {
			return nil, fmt.Errorf("low nie jest string")
		}
		low, err := parseFloat(lowStr)
		if err != nil {
			return nil, err
		}

		closeStr, ok := c[4].(string)
		if !ok {
			return nil, fmt.Errorf("close nie jest string")
		}
		closeVal, err := parseFloat(closeStr)
		if err != nil {
			return nil, err
		}

		vwapStr, ok := c[5].(string)
		if !ok {
			return nil, fmt.Errorf("vwap nie jest string")
		}
		vwap, err := parseFloat(vwapStr)
		if err != nil {
			return nil, err
		}

		volumeStr, ok := c[6].(string)
		if !ok {
			return nil, fmt.Errorf("volume nie jest string")
		}
		volume, err := parseFloat(volumeStr)
		if err != nil {
			return nil, err
		}

		countFloat, ok := c[7].(float64)
		if !ok {
			return nil, fmt.Errorf("count nie jest float64")
		}
		count := int(countFloat)

		candle := Candle{
			Time:   ts,
			Open:   open,
			High:   high,
			Low:    low,
			Close:  closeVal,
			VWAP:   vwap,
			Volume: volume,
			Count:  count,
		}

		candles = append(candles, candle)
	}

	return candles, nil
}

func parseFloat(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
