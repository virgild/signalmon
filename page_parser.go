package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseDiagnosticsPage(pagebody string) (Signals, error) {
	pagebytes := bytes.NewBufferString(pagebody)

	signals := Signals{
		ForwardSignals: make(map[int]ForwardSignal),
		ReturnSignals:  make(map[int]ReturnSignal),
	}

	doc, err := goquery.NewDocumentFromReader(pagebytes)
	if err != nil {
		return Signals{}, err
	}

	forward_rows := doc.Find("table.light").Slice(0, 1).Find("tr")
	forward_rows.Slice(1, forward_rows.Length()).Each(func(i int, tr *goquery.Selection) {
		fmt.Println(tr.Html())
		cells := tr.Find("td").Map(func(i int, td *goquery.Selection) string {
			return td.Text()
		})
		channel, err := strconv.Atoi(cells[0])

		var frequency, power, snr, ber float32
		var modulation int
		if err != nil {
			panic(err)
		}

		_, err = fmt.Sscanf(cells[1], "%f MHz", &frequency)
		if err != nil {
			panic(err)
		}

		_, err = fmt.Sscanf(cells[2], "%f dBmV", &power)
		if err != nil {
			panic(err)
		}

		_, err = fmt.Sscanf(cells[3], "%f dB", &snr)
		if err != nil {
			panic(err)
		}

		_, err = fmt.Sscanf(cells[4], "%f %%", &ber)
		if err != nil {
			panic(err)
		}

		_, err = fmt.Sscanf(cells[5], "%d QAM", &modulation)
		if err != nil {
			panic(err)
		}
		signals.ForwardSignals[channel] = ForwardSignal{
			Channel:    channel,
			Frequency:  float32(frequency),
			Power:      float32(power),
			SNR:        float32(snr),
			BER:        float32(ber),
			Modulation: modulation,
		}
	})

	back_rows := doc.Find("table.light").Slice(1, 2).Find("tr")
	back_rows.Slice(1, back_rows.Length()).Each(func(i int, tr *goquery.Selection) {
		cells := tr.Find("td").Map(func(i int, td *goquery.Selection) string {
			return td.Text()
		})
		channel, err := strconv.Atoi(cells[0])
		if err != nil {
			panic(err)
		}
		var frequency, power float32
		var modulation int

		_, err = fmt.Sscanf(cells[1], "%f MHz", &frequency)
		if err != nil {
			panic(err)
		}

		_, err = fmt.Sscanf(cells[2], "%f dBmV", &power)
		if err != nil {
			panic(err)
		}

		if strings.Contains(cells[3], "QPSK") {
			modulation = 0
		} else {
			_, err = fmt.Sscanf(cells[3], "%d QAM", &modulation)
			if err != nil {
				panic(err)
			}
		}

		signals.ReturnSignals[channel] = ReturnSignal{
			Channel:    channel,
			Frequency:  float32(frequency),
			Power:      float32(power),
			Modulation: modulation,
		}
	})

	return signals, nil
}
