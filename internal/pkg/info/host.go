package info

import (
	"encoding/json"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"github.com/shirou/gopsutil/host"
)

// HostStat def
type HostStat struct {
	Hostname             string       `json:"hostname"`
	Uptime               second       `json:"uptime"`
	BootTime             systemtime   `json:"bootTime"`
	Procs                number       `json:"procs"`
	OS                   string       `json:"os"`
	Platform             string       `json:"platform"`
	PlatformFamily       string       `json:"platformFamily"`
	PlatformVersion      string       `json:"platformVersion"`
	KernelVersion        string       `json:"kernelVersion"`
	VirtualizationSystem string       `json:"virtualizationSystem"`
	VirtualizationRole   string       `json:"virtualizationRole"`
	HostID               string       `json:"hostid"`
	Sensors              []sensorStat `json:"sensors"`
}

// HostStatFormat def
type HostStatFormat struct {
	Hostname             string             `json:"hostname"`
	Uptime               string             `json:"uptime"`
	BootTime             string             `json:"bootTime"`
	Procs                string             `json:"procs"`
	OS                   string             `json:"os"`
	Platform             string             `json:"platform"`
	PlatformFamily       string             `json:"platformFamily"`
	PlatformVersion      string             `json:"platformVersion"`
	KernelVersion        string             `json:"kernelVersion"`
	VirtualizationSystem string             `json:"virtualizationSystem"`
	VirtualizationRole   string             `json:"virtualizationRole"`
	HostID               string             `json:"hostId"`
	Sensors              []sensorStatFormat `json:"sensors"`
}

// Text formatting
func (hs HostStat) Text() HostStatFormat {
	hn := hs.Hostname
	upt := hs.Uptime.Text()
	bt := hs.BootTime.Text()
	procs := hs.Procs.Text()
	os := hs.OS
	p := hs.Platform
	pf := hs.PlatformFamily
	pv := hs.PlatformVersion
	kv := hs.KernelVersion
	vs := hs.VirtualizationSystem
	vr := hs.VirtualizationRole
	hID := hs.HostID

	var sensors []sensorStatFormat

	for _, sensor := range hs.Sensors {
		sensors = append(sensors, sensor.Text())
	}

	return HostStatFormat{hn, upt, bt, procs, os, p, pf, pv, kv, vs, vr, hID, sensors}
}

// Format formatting
func (hs HostStat) Format() HostStatFormat {
	hn := hs.Hostname
	upt := hs.Uptime.Format()
	bt := hs.BootTime.Format()
	procs := hs.Procs.Format()
	os := hs.OS
	p := hs.Platform
	pf := hs.PlatformFamily
	pv := hs.PlatformVersion
	kv := hs.KernelVersion
	vs := hs.VirtualizationSystem
	vr := hs.VirtualizationRole
	hID := hs.HostID

	var sensors []sensorStatFormat

	for _, sensor := range hs.Sensors {
		sensors = append(sensors, sensor.Format())
	}

	return HostStatFormat{hn, upt, bt, procs, os, p, pf, pv, kv, vs, vr, hID, sensors}
}

func (hs HostStatFormat) String() string {
	s, _ := json.Marshal(hs)
	return string(s)
}

func newHostStat(info *host.InfoStat, sensors []sensorStat) HostStat {
	hn := info.Hostname
	upt := second(info.Uptime)
	bt := systemtime(info.BootTime)
	procs := number(info.Procs)
	os := info.OS
	p := info.Platform
	pf := info.PlatformFamily
	pv := info.PlatformVersion
	kv := info.KernelVersion
	vs := info.VirtualizationSystem
	vr := info.VirtualizationRole
	hID := info.HostID

	return HostStat{hn, upt, bt, procs, os, p, pf, pv, kv, vs, vr, hID, sensors}
}

type sensorStat struct {
	SensorKey   string      `json:"sensorKey"`
	Temperature temperature `json:"sensorTemperature"`
}

type sensorStatFormat struct {
	SensorKey   string `json:"sensorKey"`
	Temperature string `json:"sensorTemperature"`
}

func (ss sensorStat) Text() sensorStatFormat {
	sk := ss.SensorKey
	temp := ss.Temperature.Text()

	return sensorStatFormat{sk, temp}
}

func (ss sensorStat) Format() sensorStatFormat {
	sk := ss.SensorKey
	temp := ss.Temperature.Format()

	return sensorStatFormat{sk, temp}
}

func newSensorStat(ts *host.TemperatureStat) sensorStat {
	sk := ts.SensorKey
	temp := temperature(ts.Temperature)

	return sensorStat{sk, temp}
}

// Host information
func Host(format bool) (HostStatFormat, error) {
	var sensors []sensorStat

	sensorsTemp, err := host.SensorsTemperatures()
	if err != nil {
		return HostStatFormat{}, errors.Convert(err)
	}

	for _, sensor := range sensorsTemp {
		sensors = append(sensors, newSensorStat(&sensor))
	}

	h, err := host.Info()
	if err != nil {
		return HostStatFormat{}, errors.Convert(err)
	}

	hostS := newHostStat(h, sensors)

	if format {
		return hostS.Format(), nil
	}

	return hostS.Text(), nil
}

// Users of host
func Users() ([]host.UserStat, error) {
	usr, err := host.Users()

	return usr, errors.Convert(err)
}