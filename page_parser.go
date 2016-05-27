package main

import (
	"bytes"
	"fmt"
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
		var frequency, power, snr, ber float32
		var channel, modulation int

		row_str, _ := tr.Html()
		layout := `<td align="right">%d</td><td align="right">%f MHz</td><td align="right">%f dBmV</td><td align="right">%f dB</td><td align="right">%f %%</td><td align="right">%d QAM</td>`
		_, err := fmt.Sscanf(row_str, layout, &channel, &frequency, &power, &snr, &ber, &modulation)
		if err != nil {
			panic(err)
		}

		signals.ForwardSignals[channel] = ForwardSignal{
			ChannelData: ChannelData{
				Channel:    channel,
				Frequency:  float32(frequency),
				Power:      float32(power),
				Modulation: modulation,
			},
			SNR: float32(snr),
			BER: float32(ber),
		}
	})

	back_rows := doc.Find("table.light").Slice(1, 2).Find("tr")
	back_rows.Slice(1, back_rows.Length()).Each(func(i int, tr *goquery.Selection) {
		var frequency, power float32
		var channel, modulation int

		row_str, _ := tr.Html()
		var layout string
		var err error
		if strings.Contains(row_str, "QPSK") {
			layout = `<td align="right">%d</td><td align="right">%f MHz</td><td align="right">%f dBmV</td><td align="right">QPSK</td>`
			_, err = fmt.Sscanf(row_str, layout, &channel, &frequency, &power)
		} else {
			layout = `<td align="right">%d</td><td align="right">%f MHz</td><td align="right">%f dBmV</td><td align="right">%d QAM</td>`
			_, err = fmt.Sscanf(row_str, layout, &channel, &frequency, &power, &modulation)
		}

		if err != nil {
			panic(err)
		}

		signals.ReturnSignals[channel] = ReturnSignal{
			ChannelData: ChannelData{
				Channel:    channel,
				Frequency:  float32(frequency),
				Power:      float32(power),
				Modulation: modulation,
			},
		}
	})

	return signals, nil
}
