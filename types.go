package main

type ForwardSignal struct {
	Channel    int     `json:"channel"`
	Frequency  float32 `json:"frequency"`
	Power      float32 `json:"power"`
	SNR        float32 `json:"snr"`
	BER        float32 `json:"ber"`
	Modulation int     `json:"modulation"`
}

type ForwardSignals []ForwardSignal

func (slice ForwardSignals) Len() int {
	return len(slice)
}

func (slice ForwardSignals) Less(a int, b int) bool {
	return slice[a].Channel < slice[b].Channel
}

func (slice ForwardSignals) Swap(a int, b int) {
	slice[a], slice[b] = slice[b], slice[a]
}

type ReturnSignal struct {
	Channel    int     `json:"channel"`
	Frequency  float32 `json:"frequency"`
	Power      float32 `json:"power"`
	Modulation int     `json:"modulation"`
}

type ReturnSignals []ReturnSignal

func (slice ReturnSignals) Len() int {
	return len(slice)
}

func (slice ReturnSignals) Less(a int, b int) bool {
	return slice[a].Channel < slice[b].Channel
}

func (slice ReturnSignals) Swap(a int, b int) {
	slice[a], slice[b] = slice[b], slice[a]
}

type StatsData struct {
	ForwardSignals []ForwardSignal `json:"forwardsignals"`
	ReturnSignals  []ReturnSignal  `json:"returnsignals"`
}

type Signals struct {
	ForwardSignals map[int]ForwardSignal `json:"forwardsignals"`
	ReturnSignals  map[int]ReturnSignal  `json:"returnsignals"`
}

type stats struct {
	StatusCode               string
	SoftwareVersion          string
	SoftwareModel            string
	Bootloader               string
	ProvisionedAddress       string
	ProvisionedTime          string
	ProvisionedConfiguration string
	Registered               string
	BPI                      string
	Tuning                   string
	Ranging                  string
	Connecting               string
	Configuring              string
	Registering              string
	CurrentState             string
	HighestStateObtained     string
}

type eventlog struct {
	Time        string
	Priority    string
	Description string
}
