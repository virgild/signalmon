package main

type ChannelData struct {
	Channel    int     `json:"channel"`
	Frequency  float32 `json:"frequency"`
	Power      float32 `json:"power"`
	Modulation int     `json:"modulation"`
}

type ForwardSignal struct {
	ChannelData
	SNR float32 `json:"snr"`
	BER float32 `json:"ber"`
}

type ForwardSignals []ForwardSignal

type ReturnSignal struct {
	ChannelData
}

type ReturnSignals []ReturnSignal

type StatsData struct {
	ForwardSignals []ForwardSignal `json:"forwardsignals"`
	ReturnSignals  []ReturnSignal  `json:"returnsignals"`
}

type Signals struct {
	ForwardSignals map[int]ForwardSignal `json:"forwardsignals"`
	ReturnSignals  map[int]ReturnSignal  `json:"returnsignals"`
}
