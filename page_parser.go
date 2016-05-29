package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ParseDiagnosticsPage(pagebody string) (signalsdata *SignalsData, err error) {
	pagebytes := bytes.NewBufferString(pagebody)

	doc, err := goquery.NewDocumentFromReader(pagebytes)
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("%v", p)
			signalsdata = nil
		}
	}()

	forward_rows := doc.Find("table.light").Slice(0, 1).Find("tr")
	fchannels_count := forward_rows.Length() - 1
	fchannels := make(Channels, fchannels_count, fchannels_count)
	forward_rows.Slice(1, forward_rows.Length()).Each(func(i int, tr *goquery.Selection) {
		var frequency, power, snr, ber float32
		var channel, modulation int

		row_str, _ := tr.Html()
		layout := `<td align="right">%d</td><td align="right">%f MHz</td><td align="right">%f dBmV</td><td align="right">%f dB</td><td align="right">%f %%</td><td align="right">%d QAM</td>`
		_, err := fmt.Sscanf(row_str, layout, &channel, &frequency, &power, &snr, &ber, &modulation)
		if err != nil {
			panic(fmt.Errorf("scanning forward rows: %v", err))
		}

		fchannels[i] = &ChannelData{
			channel:    channel,
			power:      power,
			frequency:  frequency,
			modulation: modulation,
			snr:        snr,
			ber:        ber,
		}
	})

	back_rows := doc.Find("table.light").Slice(1, 2).Find("tr")
	bchannels_count := back_rows.Length() - 1
	bchannels := make(Channels, bchannels_count, bchannels_count)
	back_rows.Slice(1, back_rows.Length()).Each(func(i int, tr *goquery.Selection) {
		var frequency, power float32
		var channel, modulation int

		row_str, _ := tr.Html()
		var layout string
		var err error
		if strings.Contains(row_str, "QPSK") {
			layout = `<td align="right">%d</td><td align="right">%f MHz</td><td align="right">%f dBmV</td><td align="right">QPSK</td>`
			_, err = fmt.Sscanf(row_str, layout, &channel, &frequency, &power)
			if err != nil {
				panic(fmt.Errorf("scanning return signal rows with QPSK: %v", err))
			}
		} else {
			layout = `<td align="right">%d</td><td align="right">%f MHz</td><td align="right">%f dBmV</td><td align="right">%d QAM</td>`
			_, err = fmt.Sscanf(row_str, layout, &channel, &frequency, &power, &modulation)
			if err != nil {
				panic(fmt.Errorf("scanning return signal rows with QAM: %v", err))
			}
		}

		bchannels[i] = &ChannelData{
			channel:    channel,
			power:      power,
			frequency:  frequency,
			modulation: modulation,
		}
	})

	return &SignalsData{
		ForwardSignals: fchannels,
		ReturnSignals:  bchannels,
	}, err
}
