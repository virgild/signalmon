package main

import (
	"sort"
	"testing"
)

func TestChannelSort(t *testing.T) {
	channels := Channels{
		&ChannelData{3, 0.0, 0.0, 0.0, 0.0, 0},
		&ChannelData{1, 0.0, 0.0, 0.0, 0.0, 0},
		&ChannelData{4, 0.0, 0.0, 0.0, 0.0, 0},
		&ChannelData{2, 0.0, 0.0, 0.0, 0.0, 0},
	}

	sort.Sort(channels)

	if channels[0].Number() != 1 {
		t.Errorf("Sort failed for ForwardSignal")
	}
}
