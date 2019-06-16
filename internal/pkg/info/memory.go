package info

import (
	"encoding/json"

	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"github.com/shirou/gopsutil/mem"
)

// MemoryStat desc
type MemoryStat struct {
	Total       byte    `json:"total"`
	Available   byte    `json:"available"`
	Used        byte    `json:"used"`
	UsedPercent percent `json:"usedPercent"`
	Free        byte    `json:"free"`
}

// MemoryStatFormat desc
type MemoryStatFormat struct {
	Total       string `json:"total"`
	Available   string `json:"available"`
	Used        string `json:"used"`
	Usedpercent string `json:"usedpercent"`
	Free        string `json:"free"`
}

// Text MemoryStat values
func (ms MemoryStat) Text() MemoryStatFormat {
	t := ms.Total.Text()
	a := ms.Available.Text()
	u := ms.Used.Text()
	up := ms.UsedPercent.Text()
	f := ms.Free.Text()

	return MemoryStatFormat{t, a, u, up, f}
}

// Format MemoryStat values
func (ms MemoryStat) Format() MemoryStatFormat {
	t := ms.Total.Format()
	a := ms.Available.Format()
	u := ms.Used.Format()
	up := ms.UsedPercent.Format()
	f := ms.Free.Format()

	return MemoryStatFormat{t, a, u, up, f}
}

func (ms MemoryStatFormat) String() string {
	s, _ := json.Marshal(ms)
	return string(s)
}

// Memory return memory info
func Memory(format bool) (MemoryStatFormat, error) {
	v, err := mem.VirtualMemory()

	if err != nil {
		return MemoryStatFormat{}, errors.Convert(err)
	}

	total := byte(v.Total)
	available := byte(v.Available)
	used := byte(v.Used)
	usedPer := percent(v.UsedPercent)
	free := byte(v.Free)

	memInfo := MemoryStat{total, available, used, usedPer, free}

	if format {
		return memInfo.Format(), nil
	}

	return memInfo.Text(), nil
}
