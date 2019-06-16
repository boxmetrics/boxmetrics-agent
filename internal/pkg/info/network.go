package info

import (
	"encoding/json"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"github.com/shirou/gopsutil/net"
	"time"
)

// NetworkStat def
type NetworkStat struct {
	Usage      networkIOCounter   `json:"usage"`
	Interfaces []networkInterface `json:"interfaces"`
}

// NetworkStatFormat def
type NetworkStatFormat struct {
	Usage      networkIOCounterFormat   `json:"usage"`
	Interfaces []networkInterfaceFormat `json:"interfaces"`
}

// Text formatting
func (ns NetworkStat) Text() NetworkStatFormat {
	u := ns.Usage.Text()
	var ints []networkInterfaceFormat

	for _, i := range ns.Interfaces {
		ints = append(ints, i.Text())
	}

	return NetworkStatFormat{u, ints}
}

// Format network stat
func (ns NetworkStat) Format() NetworkStatFormat {
	u := ns.Usage.Format()
	var ints []networkInterfaceFormat

	for _, i := range ns.Interfaces {
		ints = append(ints, i.Format())
	}

	return NetworkStatFormat{u, ints}
}

func (ns NetworkStatFormat) String() string {
	s, _ := json.Marshal(ns)
	return string(s)
}

type networkInterface struct {
	MTU          number              `json:"mtu"`
	Name         string              `json:"name"`
	HardwareAddr string              `json:"hardwareaddr"`
	Flags        []string            `json:"flags"`
	Addrs        []net.InterfaceAddr `json:"addrs"`
	Usage        networkIOCounter    `json:"usage"`
}

type networkInterfaceFormat struct {
	MTU          string                 `json:"mtu"`
	Name         string                 `json:"name"`
	HardwareAddr string                 `json:"hardwareaddr"`
	Flags        []string               `json:"flags"`
	Addrs        []net.InterfaceAddr    `json:"addrs"`
	Usage        networkIOCounterFormat `json:"usage"`
}

func (ni networkInterface) Text() networkInterfaceFormat {
	mtu := ni.MTU.Text()
	n := ni.Name
	ha := ni.HardwareAddr
	flags := ni.Flags
	addrs := ni.Addrs
	u := ni.Usage.Text()

	return networkInterfaceFormat{mtu, n, ha, flags, addrs, u}
}

func (ni networkInterface) Format() networkInterfaceFormat {
	mtu := ni.MTU.Format()
	n := ni.Name
	ha := ni.HardwareAddr
	flags := ni.Flags
	addrs := ni.Addrs
	u := ni.Usage.Format()

	return networkInterfaceFormat{mtu, n, ha, flags, addrs, u}
}

func newNetworkInterface(nic *net.InterfaceStat, ioBefore *net.IOCountersStat, ioAfter *net.IOCountersStat) networkInterface {
	mtu := number(nic.MTU)
	n := nic.Name
	ha := nic.HardwareAddr
	flags := nic.Flags
	addrs := nic.Addrs
	byteSentPerSec := ioAfter.BytesSent - ioBefore.BytesSent
	byteRecvPerSec := ioAfter.BytesRecv - ioBefore.BytesRecv
	u := newNetworkIOCounter(ioAfter, byteSentPerSec, byteRecvPerSec)

	return networkInterface{mtu, n, ha, flags, addrs, u}
}

type networkIOCounter struct {
	Name            string    `json:"name"`
	BytesSent       byte      `json:"bytesSent"`
	BytesRecv       byte      `json:"bytesRecv"`
	BytesSentPerSec byteSpeed `json:"bytesSentPerSec"`
	BytesRecvPerSec byteSpeed `json:"bytesRecvPerSec"`
	PacketsSent     number    `json:"packetsSent"`
	PacketsRecv     number    `json:"packetsRecv"`
	Errin           number    `json:"errin"`
	Errout          number    `json:"errout"`
	Dropin          number    `json:"dropin"`
	Dropout         number    `json:"dropout"`
	Fifoin          number    `json:"fifoin"`
	Fifoout         number    `json:"fifoout"`
}

type networkIOCounterFormat struct {
	Name            string `json:"name"`
	BytesSent       string `json:"bytesSent"`
	BytesRecv       string `json:"bytesRecv"`
	BytesSentPerSec string `json:"bytesSentPerSec"`
	BytesRecvPerSec string `json:"bytesRecvPerSec"`
	PacketsSent     string `json:"packetsSent"`
	PacketsRecv     string `json:"packetsRecv"`
	Errin           string `json:"errin"`
	Errout          string `json:"errout"`
	Dropin          string `json:"dropin"`
	Dropout         string `json:"dropout"`
	Fifoin          string `json:"fifoin"`
	Fifoout         string `json:"fifoout"`
}

func (nIOC networkIOCounter) Text() networkIOCounterFormat {
	n := nIOC.Name
	bs := nIOC.BytesSent.Text()
	br := nIOC.BytesRecv.Text()
	bsps := nIOC.BytesSentPerSec.Text()
	brps := nIOC.BytesRecvPerSec.Text()
	ps := nIOC.PacketsSent.Text()
	pr := nIOC.PacketsRecv.Text()
	ei := nIOC.Errin.Text()
	eo := nIOC.Errout.Text()
	di := nIOC.Dropin.Text()
	do := nIOC.Dropout.Text()
	fi := nIOC.Fifoin.Text()
	fo := nIOC.Fifoout.Text()

	return networkIOCounterFormat{n, bs, br, bsps, brps, ps, pr, ei, eo, di, do, fi, fo}
}

func (nIOC networkIOCounter) Format() networkIOCounterFormat {
	n := nIOC.Name
	bs := nIOC.BytesSent.Format()
	br := nIOC.BytesRecv.Format()
	bsps := nIOC.BytesSentPerSec.Format()
	brps := nIOC.BytesRecvPerSec.Format()
	ps := nIOC.PacketsSent.Format()
	pr := nIOC.PacketsRecv.Format()
	ei := nIOC.Errin.Format()
	eo := nIOC.Errout.Format()
	di := nIOC.Dropin.Format()
	do := nIOC.Dropout.Format()
	fi := nIOC.Fifoin.Format()
	fo := nIOC.Fifoout.Format()

	return networkIOCounterFormat{n, bs, br, bsps, brps, ps, pr, ei, eo, di, do, fi, fo}
}

func newNetworkIOCounter(io *net.IOCountersStat, byteSentPerSec uint64, byteRecvPerSec uint64) networkIOCounter {
	n := io.Name
	bs := byte(io.BytesSent)
	br := byte(io.BytesRecv)
	bsps := byteSpeed(byteSentPerSec)
	brps := byteSpeed(byteRecvPerSec)
	ps := number(io.PacketsSent)
	pr := number(io.PacketsRecv)
	ei := number(io.Errin)
	eo := number(io.Errout)
	di := number(io.Dropin)
	do := number(io.Dropout)
	fi := number(io.Fifoin)
	fo := number(io.Fifoout)

	return networkIOCounter{n, bs, br, bsps, brps, ps, pr, ei, eo, di, do, fi, fo}
}

func buildTotalIOCounter(io chan networkIOCounter, e chan error) {

	before, err := net.IOCounters(false)
	if err != nil {
		io <- networkIOCounter{}
		e <- err
		return
	}
	time.Sleep(time.Second)
	after, err := net.IOCounters(false)
	if err != nil {
		io <- networkIOCounter{}
		e <- err
		return
	}

	bsps := after[0].BytesSent - before[0].BytesSent
	brps := after[0].BytesRecv - before[0].BytesRecv

	io <- newNetworkIOCounter(&after[0], bsps, brps)
	e <- nil
}

func buildInterfaces(result chan []networkInterface, e chan error) {

	var netInts []networkInterface

	nicStat, err := net.Interfaces()
	if err != nil {
		result <- netInts
		e <- err
		return
	}

	before, err := net.IOCounters(true)
	if err != nil {
		result <- netInts
		e <- err
		return
	}

	time.Sleep(time.Second)

	after, err := net.IOCounters(true)
	if err != nil {
		result <- netInts
		e <- err
		return
	}

	for _, nic := range nicStat {
		for i := range after {
			if after[i].Name == nic.Name {

				netInts = append(netInts, newNetworkInterface(&nic, &before[i], &after[i]))
			}
		}
	}

	result <- netInts
	e <- nil
}

// Network information
func Network(format bool) (NetworkStatFormat, error) {
	e := make(chan error)
	io := make(chan networkIOCounter)
	netInts := make(chan []networkInterface)

	go buildTotalIOCounter(io, e)
	go buildInterfaces(netInts, e)

	usage := <-io
	nic := <-netInts
	err := <-e
	if err != nil {
		return NetworkStatFormat{}, errors.Convert(err)
	}

	ns := NetworkStat{usage, nic}

	if format {
		return ns.Format(), nil
	}
	return ns.Text(), nil
}

// Connections return all active connection
func Connections() ([]net.ConnectionStat, error) {
	conns, err := net.Connections("all")

	return conns, errors.Convert(err)
}
