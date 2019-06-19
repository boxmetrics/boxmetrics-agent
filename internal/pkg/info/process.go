package info

import (
	"encoding/json"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
	"sync"
	"time"
)

// ProcessStat def
type ProcessStat struct {
	Pid         number                  `json:"pid"`
	Name        string                  `json:"name"`
	Username    string                  `json:"username"`
	Status      string                  `json:"status"`
	Uids        numberSlice             `json:"uids"`
	Gids        numberSlice             `json:"gids"`
	Terminal    string                  `json:"terminal"`
	Cwd         string                  `json:"cwd"`
	Exe         string                  `json:"exe"`
	CmdArgs     []string                `json:"cmdArgs"`
	Ppid        number                  `json:"ppid"`
	Parent      ProcessLightStat        `json:"parent"`
	Children    []ProcessLightStat      `json:"children"`
	CPU         processCPUInfo          `json:"cpu"`
	Memory      processMemInfo          `json:"memory"`
	Network     networkIOCounter        `json:"network"`
	Disk        processDiskInfo         `json:"disk"`
	Connections []net.ConnectionStat    `json:"connections"`
	NumFDS      number                  `json:"numFds"`
	OpenFiles   []process.OpenFilesStat `json:"openFiles"`
	NumThreads  number                  `json:"numThread"`
	Threads     map[number]cpuTimes     `json:"threads"`
	Rlimits     []processLimit          `json:"limits"`
	IsRunning   bool                    `json:"isRunning"`
	Background  bool                    `json:"background"`
	Foreground  bool                    `json:"foreground"`
	CreateTime  systemtime              `json:"createTime"`
}

// ProcessStatFormat def
type ProcessStatFormat struct {
	Pid         string                    `json:"pid"`
	Name        string                    `json:"name"`
	Username    string                    `json:"username"`
	Status      string                    `json:"status"`
	Uids        []string                  `json:"uids"`
	Gids        []string                  `json:"gids"`
	Terminal    string                    `json:"terminal"`
	Cwd         string                    `json:"cwd"`
	Exe         string                    `json:"exe"`
	CmdArgs     []string                  `json:"cmdArgs"`
	Ppid        string                    `json:"ppid"`
	Parent      ProcessLightStatFormat    `json:"parent"`
	Children    []ProcessLightStatFormat  `json:"children"`
	CPU         processCPUInfoFormat      `json:"cpu"`
	Memory      processMemInfoFormat      `json:"memory"`
	Network     networkIOCounterFormat    `json:"network"`
	Disk        processDiskInfoFormat     `json:"disk"`
	Connections []net.ConnectionStat      `json:"connections"`
	NumFDS      string                    `json:"numFds"`
	OpenFiles   []process.OpenFilesStat   `json:"openFiles"`
	NumThreads  string                    `json:"numThread"`
	Threads     map[string]cpuTimesFormat `json:"threads"`
	Rlimits     []processLimitFormat      `json:"limits"`
	IsRunning   bool                      `json:"isRunning"`
	Background  bool                      `json:"background"`
	Foreground  bool                      `json:"foreground"`
	CreateTime  string                    `json:"createTime"`
}

// Text conversion
func (p ProcessStat) Text() ProcessStatFormat {
	pid := p.Pid.Text()
	name := p.Name
	username := p.Username
	status := p.Status
	uids := convertInfoSlice(p.Uids.Convert(), "text")
	gids := convertInfoSlice(p.Gids.Convert(), "text")
	term := p.Terminal
	cwd := p.Cwd
	exe := p.Exe
	cmdArgs := p.CmdArgs
	ppid := p.Ppid.Text()
	parent := p.Parent.Text()
	var children []ProcessLightStatFormat
	for _, child := range p.Children {
		children = append(children, child.Text())
	}
	cpu := p.CPU.Text()
	mem := p.Memory.Text()
	net := p.Network.Text()
	disk := p.Disk.Text()
	conns := p.Connections
	numFds := p.NumFDS.Text()
	openFile := p.OpenFiles
	numThread := p.NumThreads.Text()
	threads := make(map[string]cpuTimesFormat)
	for key, th := range p.Threads {
		threads[key.Text()] = th.Text()
	}
	var rlimits []processLimitFormat
	for _, limit := range p.Rlimits {
		rlimits = append(rlimits, limit.Text())
	}
	isR := p.IsRunning
	bck := p.Background
	fgr := p.Foreground
	ctime := p.CreateTime.Text()

	return ProcessStatFormat{pid, name, username, status, uids, gids, term, cwd, exe, cmdArgs, ppid, parent, children, cpu, mem, net, disk, conns, numFds, openFile, numThread, threads, rlimits, isR, bck, fgr, ctime}
}

// Format conversion
func (p ProcessStat) Format() ProcessStatFormat {
	pid := p.Pid.Format()
	name := p.Name
	username := p.Username
	status := p.Status
	uids := convertInfoSlice(p.Uids.Convert(), "format")
	gids := convertInfoSlice(p.Gids.Convert(), "format")
	term := p.Terminal
	cwd := p.Cwd
	exe := p.Exe
	cmdArgs := p.CmdArgs
	ppid := p.Ppid.Format()
	parent := p.Parent.Format()
	var children []ProcessLightStatFormat
	for _, child := range p.Children {
		children = append(children, child.Format())
	}
	cpu := p.CPU.Format()
	mem := p.Memory.Format()
	net := p.Network.Format()
	disk := p.Disk.Format()
	conns := p.Connections
	numFds := p.NumFDS.Format()
	openFile := p.OpenFiles
	numThread := p.NumThreads.Format()
	threads := make(map[string]cpuTimesFormat)
	for key, th := range p.Threads {
		threads[key.Format()] = th.Format()
	}
	var rlimits []processLimitFormat
	for _, limit := range p.Rlimits {
		rlimits = append(rlimits, limit.Format())
	}
	isR := p.IsRunning
	bck := p.Background
	fgr := p.Foreground
	ctime := p.CreateTime.Format()

	return ProcessStatFormat{pid, name, username, status, uids, gids, term, cwd, exe, cmdArgs, ppid, parent, children, cpu, mem, net, disk, conns, numFds, openFile, numThread, threads, rlimits, isR, bck, fgr, ctime}
}

func (p ProcessStatFormat) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}

func newProcessStat(p *process.Process) ProcessStat {
	pid := number(p.Pid)
	name, _ := p.Name()
	username, _ := p.Username()
	status, _ := p.Status()
	uidsRaw, _ := p.Uids()
	uids := newNumberSlice32(uidsRaw)
	gidsRaw, _ := p.Gids()
	gids := newNumberSlice32(gidsRaw)
	term, _ := p.Terminal()
	cwd, _ := p.Cwd()
	exe, _ := p.Exe()
	cmdArgs, _ := p.CmdlineSlice()
	ppidRaw, _ := p.Ppid()
	ppid := number(ppidRaw)
	parentRaw, _ := p.Parent()
	parent := newProcessLightStat(parentRaw)
	childrenRaw, _ := p.Children()
	var children []ProcessLightStat
	for _, proc := range childrenRaw {
		children = append(children, newProcessLightStat(proc))
	}
	cpuPer, _ := p.CPUPercent()
	cpuTimeRaw, _ := p.Times()
	cpuTime := newCPUTimes(cpuTimeRaw)
	cpu := newProcessCPUInfo(cpuPer, cpuTime)
	memPer, _ := p.MemoryPercent()
	memInfoRaw, _ := p.MemoryInfo()
	memInfo := newProcessMemUsage(memInfoRaw)
	mem := newProcessMemInfo(memPer, memInfo)
	e := make(chan error)
	ionet := make(chan networkIOCounter)
	iodisk := make(chan processDiskInfo)
	go buildProcessNetInfo(p, ionet, e)
	go buildProcessDiskInfo(p, iodisk, e)
	net := <-ionet
	disk := <-iodisk
	conns, _ := p.Connections()
	numFdsRaw, _ := p.NumFDs()
	numFds := number(numFdsRaw)
	openFiles, _ := p.OpenFiles()
	numThreadsRaw, _ := p.NumThreads()
	numThreads := number(numThreadsRaw)
	threadsRaw, _ := p.Threads()
	threads := make(map[number]cpuTimes)
	for key, th := range threadsRaw {
		threads[number(key)] = newCPUTimes(th)
	}
	rlimitRaw, _ := p.Rlimit()
	var rlimits []processLimit
	for _, lim := range rlimitRaw {
		rlimits = append(rlimits, newProcessLimit(&lim))
	}
	isR, _ := p.IsRunning()
	bck, _ := p.Background()
	fgr, _ := p.Foreground()
	ctimeRaw, _ := p.CreateTime()
	ctime := systemtime(ctimeRaw)

	return ProcessStat{pid, name, username, status, uids, gids, term, cwd, exe, cmdArgs, ppid, parent, children, cpu, mem, net, disk, conns, numFds, openFiles, numThreads, threads, rlimits, isR, bck, fgr, ctime}
}

func buildProcessNetInfo(p *process.Process, io chan networkIOCounter, e chan error) {
	before, err := p.NetIOCounters(false)
	if err != nil {
		io <- networkIOCounter{}
		e <- err
		return
	}
	time.Sleep(time.Second)
	after, err := p.NetIOCounters(false)
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

func buildProcessDiskInfo(p *process.Process, io chan processDiskInfo, e chan error) {
	before, err := p.IOCounters()
	if err != nil {
		io <- processDiskInfo{}
		e <- err
		return
	}
	time.Sleep(time.Second)
	after, err := p.IOCounters()
	if err != nil {
		io <- processDiskInfo{}
		e <- err
		return
	}

	io <- newProcessDiskInfo(before, after)
	e <- nil
}

type processCPUInfo struct {
	Percent percent  `json:"percent"`
	Times   cpuTimes `json:"times"`
}

type processCPUInfoFormat struct {
	Percent string         `json:"percent"`
	Times   cpuTimesFormat `json:"times"`
}

func (p processCPUInfo) Text() processCPUInfoFormat {
	per := p.Percent.Text()
	times := p.Times.Text()

	return processCPUInfoFormat{per, times}
}

func (p processCPUInfo) Format() processCPUInfoFormat {
	per := p.Percent.Format()
	times := p.Times.Format()

	return processCPUInfoFormat{per, times}
}

func newProcessCPUInfo(perCPU float64, times cpuTimes) processCPUInfo {
	per := percent(perCPU)

	return processCPUInfo{per, times}
}

type processMemInfo struct {
	Percent percent         `json:"percent"`
	Usage   processMemUsage `json:"usage"`
}

type processMemInfoFormat struct {
	Percent string                `json:"percent"`
	Usage   processMemUsageFormat `json:"usage"`
}

func (p processMemInfo) Text() processMemInfoFormat {
	per := p.Percent.Text()
	usage := p.Usage.Text()

	return processMemInfoFormat{per, usage}
}

func (p processMemInfo) Format() processMemInfoFormat {
	per := p.Percent.Format()
	usage := p.Usage.Format()

	return processMemInfoFormat{per, usage}
}

func newProcessMemInfo(memPer float32, usage processMemUsage) processMemInfo {
	per := percent(memPer)

	return processMemInfo{per, usage}
}

type processMemUsage struct {
	RSS    byte `json:"rss"`
	VMS    byte `json:"vms"`
	Data   byte `json:"data"`
	Stack  byte `json:"stack"`
	Locked byte `json:"locked"`
	Swap   byte `json:"swap"`
}

type processMemUsageFormat struct {
	RSS    string `json:"rss"`
	VMS    string `json:"vms"`
	Data   string `json:"data"`
	Stack  string `json:"stack"`
	Locked string `json:"locked"`
	Swap   string `json:"swap"`
}

func (p processMemUsage) Text() processMemUsageFormat {
	rss := p.RSS.Text()
	vms := p.VMS.Text()
	data := p.Data.Text()
	stack := p.Stack.Text()
	locked := p.Locked.Text()
	swap := p.Swap.Text()

	return processMemUsageFormat{rss, vms, data, stack, locked, swap}
}
func (p processMemUsage) Format() processMemUsageFormat {
	rss := p.RSS.Format()
	vms := p.VMS.Format()
	data := p.Data.Format()
	stack := p.Stack.Format()
	locked := p.Locked.Format()
	swap := p.Swap.Format()

	return processMemUsageFormat{rss, vms, data, stack, locked, swap}
}

func newProcessMemUsage(mem *process.MemoryInfoStat) processMemUsage {
	if mem == nil {
		return processMemUsage{}
	}
	rss := byte(mem.RSS)
	vms := byte(mem.VMS)
	data := byte(mem.Data)
	stack := byte(mem.Stack)
	locked := byte(mem.Locked)
	swap := byte(mem.Swap)

	return processMemUsage{rss, vms, data, stack, locked, swap}
}

type processDiskInfo struct {
	ReadCount        number    `json:"readCount"`
	WriteCount       number    `json:"writeCount"`
	ReadBytes        byte      `json:"readBytes"`
	WriteBytes       byte      `json:"writeBytes"`
	ReadBytesPerSec  byteSpeed `json:"readBytesPerSec"`
	WriteBytesPerSec byteSpeed `json:"writeBytesPerSec"`
}

type processDiskInfoFormat struct {
	ReadCount        string `json:"readCount"`
	WriteCount       string `json:"writeCount"`
	ReadBytes        string `json:"readBytes"`
	WriteBytes       string `json:"writeBytes"`
	ReadBytesPerSec  string `json:"readBytesPerSec"`
	WriteBytesPerSec string `json:"writeBytesPerSec"`
}

func (p processDiskInfo) Text() processDiskInfoFormat {
	rc := p.ReadCount.Text()
	wc := p.WriteCount.Text()
	rb := p.ReadBytes.Text()
	wb := p.WriteBytes.Text()
	rcps := p.ReadBytesPerSec.Text()
	wcps := p.WriteBytesPerSec.Text()

	return processDiskInfoFormat{rc, wc, rb, wb, rcps, wcps}
}

func (p processDiskInfo) Format() processDiskInfoFormat {
	rc := p.ReadCount.Format()
	wc := p.WriteCount.Format()
	rb := p.ReadBytes.Format()
	wb := p.WriteBytes.Format()
	rcps := p.ReadBytesPerSec.Format()
	wcps := p.WriteBytesPerSec.Format()

	return processDiskInfoFormat{rc, wc, rb, wb, rcps, wcps}
}

func newProcessDiskInfo(before *process.IOCountersStat, after *process.IOCountersStat) processDiskInfo {
	rc := number(after.ReadCount)
	wc := number(after.WriteCount)
	rb := byte(after.ReadBytes)
	wb := byte(after.WriteBytes)
	rcps := byteSpeed(after.ReadBytes - before.ReadBytes)
	wcps := byteSpeed(after.WriteBytes - before.WriteBytes)

	return processDiskInfo{rc, wc, rb, wb, rcps, wcps}
}

type processLimit struct {
	Resource number `json:"resource"`
	Soft     number `json:"soft"`
	Hard     number `json:"hard"`
	Used     number `json:"used"`
}

type processLimitFormat struct {
	Resource string `json:"resource"`
	Soft     string `json:"soft"`
	Hard     string `json:"hard"`
	Used     string `json:"used"`
}

func (p processLimit) Text() processLimitFormat {
	r := p.Resource.Text()
	s := p.Soft.Text()
	h := p.Hard.Text()
	u := p.Used.Text()

	return processLimitFormat{r, s, h, u}
}

func (p processLimit) Format() processLimitFormat {
	r := p.Resource.Format()
	s := p.Soft.Format()
	h := p.Hard.Format()
	u := p.Used.Format()

	return processLimitFormat{r, s, h, u}
}

func newProcessLimit(p *process.RlimitStat) processLimit {
	r := number(p.Resource)
	s := number(p.Soft)
	h := number(p.Hard)
	u := number(p.Used)

	return processLimit{r, s, h, u}
}

// ProcessLightStat def
type ProcessLightStat struct {
	Pid        number         `json:"pid"`
	Name       string         `json:"name"`
	Username   string         `json:"username"`
	Status     string         `json:"status"`
	Uids       numberSlice    `json:"uids"`
	Gids       numberSlice    `json:"gids"`
	Terminal   string         `json:"terminal"`
	Cwd        string         `json:"cwd"`
	Exe        string         `json:"exe"`
	CmdArgs    []string       `json:"cmdArgs"`
	Ppid       number         `json:"ppid"`
	CPU        processCPUInfo `json:"cpu"`
	Memory     processMemInfo `json:"memory"`
	CreateTime systemtime     `json:"createTime"`
}

// ProcessLightStatFormat def
type ProcessLightStatFormat struct {
	Pid        string               `json:"pid"`
	Name       string               `json:"name"`
	Username   string               `json:"username"`
	Status     string               `json:"status"`
	Uids       []string             `json:"uids"`
	Gids       []string             `json:"gids"`
	Terminal   string               `json:"terminal"`
	Cwd        string               `json:"cwd"`
	Exe        string               `json:"exe"`
	CmdArgs    []string             `json:"cmdArgs"`
	Ppid       string               `json:"ppid"`
	CPU        processCPUInfoFormat `json:"cpu"`
	Memory     processMemInfoFormat `json:"memory"`
	CreateTime string               `json:"createTime"`
}

// Text conversion
func (p ProcessLightStat) Text() ProcessLightStatFormat {
	pid := p.Pid.Text()
	name := p.Name
	username := p.Username
	status := p.Status
	uids := convertInfoSlice(p.Uids.Convert(), "text")
	gids := convertInfoSlice(p.Gids.Convert(), "text")
	terminal := p.Terminal
	cwd := p.Cwd
	exe := p.Exe
	cmdArgs := p.CmdArgs
	ppid := p.Pid.Text()
	cpu := p.CPU.Text()
	mem := p.Memory.Text()
	createTime := p.CreateTime.Text()

	return ProcessLightStatFormat{pid, name, username, status, uids, gids, terminal, cwd, exe, cmdArgs, ppid, cpu, mem, createTime}
}

// Format return
func (p ProcessLightStat) Format() ProcessLightStatFormat {
	pid := p.Pid.Format()
	name := p.Name
	username := p.Username
	status := p.Status
	uids := convertInfoSlice(p.Uids.Convert(), "format")
	gids := convertInfoSlice(p.Gids.Convert(), "format")
	terminal := p.Terminal
	cwd := p.Cwd
	exe := p.Exe
	cmdArgs := p.CmdArgs
	ppid := p.Pid.Format()
	cpu := p.CPU.Format()
	mem := p.Memory.Format()
	createTime := p.CreateTime.Format()

	return ProcessLightStatFormat{pid, name, username, status, uids, gids, terminal, cwd, exe, cmdArgs, ppid, cpu, mem, createTime}
}

func newProcessLightStat(p *process.Process) ProcessLightStat {
	if p == nil {
		return ProcessLightStat{}
	}

	pid := number(p.Pid)
	name, _ := p.Name()
	username, _ := p.Username()
	status, _ := p.Status()
	u, _ := p.Uids()
	uids := newNumberSlice32(u)
	g, _ := p.Gids()
	gids := newNumberSlice32(g)
	terminal, _ := p.Terminal()
	cwd, _ := p.Cwd()
	exe, _ := p.Exe()
	cmdArgs, _ := p.CmdlineSlice()
	ppidRaw, _ := p.Ppid()
	ppid := number(ppidRaw)
	cpuPer, _ := p.CPUPercent()
	cpuTimeRaw, _ := p.Times()
	cpuTime := newCPUTimes(cpuTimeRaw)
	cpu := newProcessCPUInfo(cpuPer, cpuTime)
	memPer, _ := p.MemoryPercent()
	memInfoRaw, _ := p.MemoryInfo()
	memInfo := newProcessMemUsage(memInfoRaw)
	mem := newProcessMemInfo(memPer, memInfo)
	ctime, _ := p.CreateTime()
	createTime := systemtime(ctime / 1000)

	return ProcessLightStat{pid, name, username, status, uids, gids, terminal, cwd, exe, cmdArgs, ppid, cpu, mem, createTime}
}

// Processes return all processes informations
func Processes(format bool) ([]ProcessLightStatFormat, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, errors.Convert(err)
	}

	var wg sync.WaitGroup
	var processesFormat []ProcessLightStatFormat
	for _, process := range processes {
		wg.Add(1)

		go buildProcessStat(&wg, process, format, &processesFormat)
	}
	wg.Wait()

	return processesFormat, nil
}

// Process return stat of one process
func Process(pid int32, format bool) (ProcessStatFormat, error) {
	process, err := process.NewProcess(pid)
	if err != nil {
		return ProcessStatFormat{}, errors.Convert(err)
	}

	stat := newProcessStat(process)

	if format {
		return stat.Format(), nil
	}

	return stat.Text(), nil
}

func buildProcessStat(wg *sync.WaitGroup, p *process.Process, format bool, results *[]ProcessLightStatFormat) {
	defer wg.Done()

	stat := newProcessLightStat(p)

	if format {
		*results = append(*results, stat.Format())
	} else {
		*results = append(*results, stat.Text())
	}

}
