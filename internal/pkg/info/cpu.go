package info

import (
	"encoding/json"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"strconv"
	"time"
)

// CPUStat desc
type CPUStat struct {
	Count cpuCount  `json:"count"`
	Info  []cpuInfo `json:"info"`
}

// CPUStatFormat desc
type CPUStatFormat struct {
	Count cpuCountFormat  `json:"count"`
	Info  []cpuInfoFormat `json:"info"`
}

// Text CPUStat
func (cs CPUStat) Text() CPUStatFormat {
	c := cs.Count.Text()

	var i []cpuInfoFormat

	for _, info := range cs.Info {
		i = append(i, info.Text())
	}

	return CPUStatFormat{c, i}
}

// Format CPUStat
func (cs CPUStat) Format() CPUStatFormat {
	c := cs.Count.Format()

	var i []cpuInfoFormat

	for _, info := range cs.Info {
		i = append(i, info.Format())
	}

	return CPUStatFormat{c, i}
}

func (cs CPUStatFormat) String() string {
	s, _ := json.Marshal(cs)
	return string(s)
}

type cpuCount struct {
	Physical number `json:"physical"`
	Logical  number `json:"logical"`
}
type cpuCountFormat struct {
	Physical string `json:"physical"`
	Logical  string `json:"logical"`
}

func (c cpuCount) Text() cpuCountFormat {
	p := c.Physical.Text()
	l := c.Logical.Text()

	return cpuCountFormat{p, l}
}

func (c cpuCount) Format() cpuCountFormat {
	p := c.Physical.Format()
	l := c.Logical.Format()

	return cpuCountFormat{p, l}
}

func newCPUCount(physical int, logical int) cpuCount {
	p := number(physical)
	l := number(logical)

	return cpuCount{p, l}
}

type cpuInfo struct {
	ID      string   `json:"id"`
	Percent percent  `json:"percent"`
	Times   cpuTimes `json:"times"`
}

func (ci cpuInfo) Text() cpuInfoFormat {
	id := ci.ID
	p := ci.Percent.Text()
	t := ci.Times.Text()

	return cpuInfoFormat{id, p, t}
}

func (ci cpuInfo) Format() cpuInfoFormat {
	id := ci.ID
	p := ci.Percent.Format()
	t := ci.Times.Format()

	return cpuInfoFormat{id, p, t}
}

type cpuInfoFormat struct {
	ID      string         `json:"id"`
	Percent string         `json:"percent"`
	Times   cpuTimesFormat `json:"times"`
}

type cpuTimes struct {
	User      jiffy `json:"user"`
	System    jiffy `json:"system"`
	Idle      jiffy `json:"idle"`
	Nice      jiffy `json:"nice"`
	Iowait    jiffy `json:"iowait"`
	Irq       jiffy `json:"irq"`
	Softirq   jiffy `json:"softirq"`
	Steal     jiffy `json:"steal"`
	Guest     jiffy `json:"guest"`
	GuestNice jiffy `json:"guestNice"`
	Total     jiffy `json:"total"`
}

type cpuTimesFormat struct {
	User      string `json:"user"`
	System    string `json:"system"`
	Idle      string `json:"idle"`
	Nice      string `json:"nice"`
	Iowait    string `json:"iowait"`
	Irq       string `json:"irq"`
	Softirq   string `json:"softirq"`
	Steal     string `json:"steal"`
	Guest     string `json:"guest"`
	GuestNice string `json:"guestNice"`
	Total     string `json:"total"`
}

func (ct cpuTimes) Text() cpuTimesFormat {
	u := ct.User.Text()
	s := ct.System.Text()
	i := ct.Idle.Text()
	n := ct.Nice.Text()
	io := ct.Iowait.Text()
	irq := ct.Irq.Text()
	softirq := ct.Softirq.Text()
	st := ct.Steal.Text()
	g := ct.Guest.Text()
	gn := ct.GuestNice.Text()
	t := ct.Total.Text()

	return cpuTimesFormat{u, s, i, n, io, irq, softirq, st, g, gn, t}
}

func (ct cpuTimes) Format() cpuTimesFormat {
	u := ct.User.Format()
	s := ct.System.Format()
	i := ct.Idle.Format()
	n := ct.Nice.Format()
	io := ct.Iowait.Format()
	irq := ct.Irq.Format()
	softirq := ct.Softirq.Format()
	st := ct.Steal.Format()
	g := ct.Guest.Format()
	gn := ct.GuestNice.Format()
	t := ct.Total.Format()

	return cpuTimesFormat{u, s, i, n, io, irq, softirq, st, g, gn, t}
}

func newCPUTimes(ts *cpu.TimesStat) cpuTimes {
	if ts == nil {
		return cpuTimes{}
	}
	user := jiffy(ts.User)
	system := jiffy(ts.System)
	idle := jiffy(ts.Idle)
	nice := jiffy(ts.Nice)
	iowait := jiffy(ts.Iowait)
	irq := jiffy(ts.Irq)
	softirq := jiffy(ts.Softirq)
	steal := jiffy(ts.Steal)
	guest := jiffy(ts.Guest)
	guestnice := jiffy(ts.GuestNice)
	total := jiffy(ts.Total())

	return cpuTimes{user, system, idle, nice, iowait, irq, softirq, steal, guest, guestnice, total}
}

func getCPUInfoTotal() (cpuInfo, error) {
	totalCPUPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return cpuInfo{}, errors.Convert(err)
	}
	p := percent(totalCPUPercent[0])

	totalTimeStat, err := cpu.Times(false)
	if err != nil {
		return cpuInfo{}, errors.Convert(err)
	}
	t := newCPUTimes(&totalTimeStat[0])

	return cpuInfo{"total", p, t}, nil
}

func getCPUInfo() ([]cpuInfo, error) {
	cpus := []cpuInfo{}

	total, err := getCPUInfoTotal()
	if err != nil {
		return nil, err
	}

	cpus = append(cpus, total)

	allTimes, err := cpu.Times(true)
	if err != nil {
		return nil, errors.Convert(err)
	}

	allPercent, err := cpu.Percent(time.Second, true)
	if err != nil {
		return nil, errors.Convert(err)
	}

	for index, val := range allTimes {
		id := strconv.Itoa(index)
		percent := percent(allPercent[index])
		times := newCPUTimes(&val)

		info := cpuInfo{id, percent, times}

		cpus = append(cpus, info)

	}

	return cpus, nil
}

// CPU info
func CPU(format bool) (CPUStatFormat, error) {

	physical, err := cpu.Counts(false)
	if err != nil {
		return CPUStatFormat{}, errors.Convert(err)
	}

	logical, err := cpu.Counts(true)
	if err != nil {
		return CPUStatFormat{}, errors.Convert(err)
	}

	infos, err := getCPUInfo()
	if err != nil {
		return CPUStatFormat{}, err
	}

	count := newCPUCount(physical, logical)
	cpuStat := CPUStat{count, infos}

	if format {
		return cpuStat.Format(), nil
	}

	return cpuStat.Text(), nil
}

// CPUinfo info on cpu
func CPUinfo() ([]cpu.InfoStat, error) {
	info, err := cpu.Info()

	return info, errors.Convert(err)
}
