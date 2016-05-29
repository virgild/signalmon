package main

type Channel interface {
	Number() int
	Frequency() float32
	Power() float32
	Modulation() int
	SNR() float32
	BER() float32
}

type ChannelData struct {
	channel    int     `json:"channel"`
	frequency  float32 `json:"frequency"`
	power      float32 `json:"power"`
	snr        float32 `json:"snr"`
	ber        float32 `json:"ber"`
	modulation int     `json:"modulation"`
}

func (c *ChannelData) Number() int {
	return c.channel
}

func (c *ChannelData) Frequency() float32 {
	return c.frequency
}

func (c *ChannelData) Power() float32 {
	return c.power
}

func (c *ChannelData) Modulation() int {
	return c.modulation
}

func (c *ChannelData) SNR() float32 {
	return c.snr
}

func (c *ChannelData) BER() float32 {
	return c.ber
}

type Channels []Channel

func (c Channels) Len() int {
	return len(c)
}

func (c Channels) Less(i, j int) bool {
	return c[i].Number() < c[j].Number()
}

func (c Channels) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type SignalsData struct {
	ForwardSignals Channels
	ReturnSignals  Channels
}
