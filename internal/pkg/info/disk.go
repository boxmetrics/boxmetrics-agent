package info

import (
	"encoding/json"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"github.com/shirou/gopsutil/disk"
)

// DiskStat def
type DiskStat struct {
	Device     string       `json:"device"`
	Mountpoint string       `json:"mountpoint"`
	Fstype     string       `json:"fstype"`
	Opts       string       `json:"opts"`
	Usage      diskUsage    `json:"usage"`
	IOCounters diskCounters `json:"iocounters"`
}

// DiskStatFormat def
type DiskStatFormat struct {
	Device     string             `json:"device"`
	Mountpoint string             `json:"mountpoint"`
	Fstype     string             `json:"fstype"`
	Opts       string             `json:"opts"`
	Usage      diskUsageFormat    `json:"usage"`
	IOCounters diskCountersFormat `json:"iocounters"`
}

// Text of DiskStat
func (d DiskStat) Text() DiskStatFormat {
	u := d.Usage.Text()
	io := d.IOCounters.Text()

	return DiskStatFormat{d.Device, d.Mountpoint, d.Fstype, d.Opts, u, io}
}

// Format of DiskStat
func (d DiskStat) Format() DiskStatFormat {
	u := d.Usage.Format()
	io := d.IOCounters.Format()

	return DiskStatFormat{d.Device, d.Mountpoint, d.Fstype, d.Opts, u, io}
}

func (d DiskStatFormat) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}

func newDiskStat(d disk.PartitionStat) (DiskStat, error) {
	u, err := disk.Usage(d.Mountpoint)
	if err != nil {
		return DiskStat{}, err
	}
	usage := newDiskUsage(u)

	io, err := disk.IOCounters(d.Mountpoint)
	if err != nil {
		return DiskStat{}, err
	}
	iocounter := newDiskCounter(io[d.Mountpoint])

	return DiskStat{d.Device, d.Mountpoint, d.Fstype, d.Opts, usage, iocounter}, nil
}

type diskUsage struct {
	Path              string  `json:"path"`
	Fstype            string  `json:"fstype"`
	Total             byte    `json:"total"`
	Free              byte    `json:"free"`
	Used              byte    `json:"used"`
	UsedPercent       percent `json:"usedPercent"`
	InodesTotal       number  `json:"inodesTotal"`
	InodesUsed        number  `json:"inodesUsed"`
	InodesFree        number  `json:"inodesFree"`
	InodesUsedPercent percent `json:"inodesUsedPercent"`
}

type diskUsageFormat struct {
	Path              string `json:"path"`
	Fstype            string `json:"fstype"`
	Total             string `json:"total"`
	Free              string `json:"free"`
	Used              string `json:"used"`
	UsedPercent       string `json:"usedPercent"`
	InodesTotal       string `json:"inodesTotal"`
	InodesUsed        string `json:"inodesUsed"`
	InodesFree        string `json:"inodesFree"`
	InodesUsedPercent string `json:"inodesUsedPercent"`
}

func (d diskUsage) Text() diskUsageFormat {
	p := d.Path
	fs := d.Fstype
	t := d.Total.Text()
	f := d.Free.Text()
	u := d.Used.Text()
	up := d.UsedPercent.Text()
	it := d.InodesTotal.Text()
	iu := d.InodesUsed.Text()
	inf := d.InodesFree.Text()
	iup := d.InodesUsedPercent.Text()

	return diskUsageFormat{p, fs, t, f, u, up, it, iu, inf, iup}
}

func (d diskUsage) Format() diskUsageFormat {
	p := d.Path
	fs := d.Fstype
	t := d.Total.Format()
	f := d.Free.Format()
	u := d.Used.Format()
	up := d.UsedPercent.Format()
	it := d.InodesTotal.Format()
	iu := d.InodesUsed.Format()
	inf := d.InodesFree.Format()
	iup := d.InodesUsedPercent.Format()

	return diskUsageFormat{p, fs, t, f, u, up, it, iu, inf, iup}
}

func newDiskUsage(d *disk.UsageStat) diskUsage {
	p := d.Path
	fs := d.Fstype
	t := byte(d.Total)
	f := byte(d.Free)
	u := byte(d.Used)
	up := percent(d.UsedPercent)
	it := number(d.InodesTotal)
	iu := number(d.InodesUsed)
	inf := number(d.InodesFree)
	iup := percent(d.InodesUsedPercent)

	return diskUsage{p, fs, t, f, u, up, it, iu, inf, iup}
}

type diskCounters struct {
	ReadCount        number      `json:"readCount"`
	MergedReadCount  number      `json:"mergedReadCount"`
	WriteCount       number      `json:"writeCount"`
	MergedWriteCount number      `json:"mergedWriteCount"`
	ReadBytes        byte        `json:"readBytes"`
	WriteBytes       byte        `json:"writeBytes"`
	ReadTime         millisecond `json:"readTime"`
	WriteTime        millisecond `json:"writeTime"`
	IopsInProgress   number      `json:"iopsInProgress"`
	IoTime           millisecond `json:"ioTime"`
	WeightedIO       number      `json:"weightedIO"`
	Name             string      `json:"name"`
	Serialnumber     string      `json:"serialnumber"`
	Label            string      `json:"label"`
}

type diskCountersFormat struct {
	ReadCount        string `json:"readCount"`
	MergedReadCount  string `json:"mergedReadCount"`
	WriteCount       string `json:"writeCount"`
	MergedWriteCount string `json:"mergedWriteCount"`
	ReadBytes        string `json:"readBytes"`
	WriteBytes       string `json:"writeBytes"`
	ReadTime         string `json:"readTime"`
	WriteTime        string `json:"writeTime"`
	IopsInProgress   string `json:"iopsInProgress"`
	IoTime           string `json:"ioTime"`
	WeightedIO       string `json:"weightedIO"`
	Name             string `json:"name"`
	Serialnumber     string `json:"serialnumber"`
	Label            string `json:"label"`
}

func (d diskCounters) Text() diskCountersFormat {
	rc := d.ReadCount.Text()
	mrc := d.MergedReadCount.Text()
	wc := d.WriteCount.Text()
	mwc := d.MergedWriteCount.Text()
	rb := d.ReadBytes.Text()
	wb := d.WriteBytes.Text()
	rt := d.ReadTime.Text()
	wt := d.WriteTime.Text()
	iip := d.IopsInProgress.Text()
	it := d.IoTime.Text()
	wi := d.WeightedIO.Text()
	n := d.Name
	sn := d.Serialnumber
	l := d.Label

	return diskCountersFormat{rc, mrc, wc, mwc, rb, wb, rt, wt, iip, it, wi, n, sn, l}
}

func (d diskCounters) Format() diskCountersFormat {
	rc := d.ReadCount.Format()
	mrc := d.MergedReadCount.Format()
	wc := d.WriteCount.Format()
	mwc := d.MergedWriteCount.Format()
	rb := d.ReadBytes.Format()
	wb := d.WriteBytes.Format()
	rt := d.ReadTime.Format()
	wt := d.WriteTime.Format()
	iip := d.IopsInProgress.Format()
	it := d.IoTime.Format()
	wi := d.WeightedIO.Format()
	n := d.Name
	sn := d.Serialnumber
	l := d.Label

	return diskCountersFormat{rc, mrc, wc, mwc, rb, wb, rt, wt, iip, it, wi, n, sn, l}
}

func newDiskCounter(d disk.IOCountersStat) diskCounters {
	rc := number(d.ReadCount)
	mrc := number(d.MergedReadCount)
	wc := number(d.WriteCount)
	mwc := number(d.MergedWriteCount)
	rb := byte(d.ReadBytes)
	wb := byte(d.WriteBytes)
	rt := millisecond(d.ReadTime)
	wt := millisecond(d.WriteTime)
	iip := number(d.IopsInProgress)
	it := millisecond(d.IoTime)
	wi := number(d.WeightedIO)
	n := d.Name
	sn := d.SerialNumber
	l := d.Label

	return diskCounters{rc, mrc, wc, mwc, rb, wb, rt, wt, iip, it, wi, n, sn, l}
}

// Disks info
func Disks(format bool) ([]DiskStatFormat, error) {
	d, err := disk.Partitions(true)

	if err != nil {
		return nil, errors.Convert(err)
	}

	var diskinfo []DiskStatFormat

	for _, disk := range d {
		ds, err := newDiskStat(disk)

		if err != nil {
			return nil, errors.Convert(err)
		}

		if format {
			diskinfo = append(diskinfo, ds.Format())
		} else {
			diskinfo = append(diskinfo, ds.Text())
		}
	}

	return diskinfo, nil
}
