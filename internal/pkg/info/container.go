package info

import (
	"encoding/json"
	"github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/docker"
	"strconv"
)

// ContainerStat def
type ContainerStat struct {
	ContainerID string           `json:"containerId"`
	Name        string           `json:"name"`
	Image       string           `json:"image"`
	Status      string           `json:"status"`
	Running     bool             `json:"running"`
	CPU         containerCPUStat `json:"cpu"`
	Memory      containerMemStat `json:"memory"`
}

// ContainerStatFormat def
type ContainerStatFormat struct {
	ContainerID string                 `json:"containerId"`
	Name        string                 `json:"name"`
	Image       string                 `json:"image"`
	Status      string                 `json:"status"`
	Running     string                 `json:"running"`
	CPU         containerCPUStatFormat `json:"cpu"`
	Memory      containerMemStatFormat `json:"memory"`
}

// Text return
func (ds ContainerStat) Text() ContainerStatFormat {
	cID := ds.ContainerID
	n := ds.Name
	i := ds.Image
	s := ds.Status
	r := strconv.FormatBool(ds.Running)
	cp := ds.CPU.Text()
	mem := ds.Memory.Text()

	return ContainerStatFormat{cID, n, i, s, r, cp, mem}
}

// Format return
func (ds ContainerStat) Format() ContainerStatFormat {
	cID := ds.ContainerID
	n := ds.Name
	i := ds.Image
	s := ds.Status
	r := strconv.FormatBool(ds.Running)
	cp := ds.CPU.Format()
	mem := ds.Memory.Format()

	return ContainerStatFormat{cID, n, i, s, r, cp, mem}
}

func (ds ContainerStatFormat) String() string {
	s, _ := json.Marshal(ds)
	return string(s)
}

type containerCPUStat struct {
	Percent percent  `json:"percent"`
	Times       cpuTimes `json:"times"`
}

type containerCPUStatFormat struct {
	Percent string         `json:"percent"`
	Times       cpuTimesFormat `json:"times"`
}

func (c containerCPUStat) Text() containerCPUStatFormat {
	up := c.Percent.Text()
	t := c.Times.Text()

	return containerCPUStatFormat{up, t}
}

func (c containerCPUStat) Format() containerCPUStatFormat {
	up := c.Percent.Format()
	t := c.Times.Format()

	return containerCPUStatFormat{up, t}
}

func newContainerCPUStat(per float64, time *cpu.TimesStat) containerCPUStat {
	p := percent(per)
	t := newCPUTimes(time)

	return containerCPUStat{p, t}
}

type containerMemStat struct {
	ContainerID             string `json:"containerId"`
	Cache                   byte   `json:"cache"`
	RSS                     number `json:"rss"`
	RSSHuge                 number `json:"rssHuge"`
	MappedFile              number `json:"mappedFile"`
	Pgpgin                  number `json:"pgpgin"`
	Pgpgout                 number `json:"pgpgout"`
	Pgfault                 number `json:"pgfault"`
	Pgmajfault              number `json:"pgmajfault"`
	InactiveAnon            number `json:"inactiveAnon"`
	ActiveAnon              number `json:"activeAnon"`
	InactiveFile            number `json:"inactiveFile"`
	ActiveFile              number `json:"activeFile"`
	Unevictable             byte `json:"unevictable"`
	HierarchicalMemoryLimit byte   `json:"hierarchicalMemoryLimit"`
	TotalCache              byte   `json:"totalCache"`
	TotalRSS                number `json:"totalRss"`
	TotalRSSHuge            number `json:"totalRssHuge"`
	TotalMappedFile         number `json:"totalMappedFile"`
	TotalPgpgIn             number `json:"totalPgpgin"`
	TotalPgpgOut            number `json:"totalPgpgout"`
	TotalPgFault            number `json:"totalPgfault"`
	TotalPgMajFault         number `json:"totalPgmajfault"`
	TotalInactiveAnon       number `json:"totalInactiveAnon"`
	TotalActiveAnon         number `json:"totalActiveAnon"`
	TotalInactiveFile       number `json:"totalInactiveFile"`
	TotalActiveFile         number `json:"totalActiveFile"`
	TotalUnevictable        number `json:"totalUnevictable"`
	MemUsage                byte   `json:"memUsage"`
	MemMaxUsage             byte   `json:"memMaxUsage"`
	MemLimit                byte   `json:"memoryLimit"`
	MemFailCnt              byte   `json:"memoryFailcnt"`
}

type containerMemStatFormat struct {
	ContainerID             string `json:"containerId"`
	Cache                   string `json:"cache"`
	RSS                     string `json:"rss"`
	RSSHuge                 string `json:"rssHuge"`
	MappedFile              string `json:"mappedFile"`
	Pgpgin                  string `json:"pgpgin"`
	Pgpgout                 string `json:"pgpgout"`
	Pgfault                 string `json:"pgfault"`
	Pgmajfault              string `json:"pgmajfault"`
	InactiveAnon            string `json:"inactiveAnon"`
	ActiveAnon              string `json:"activeAnon"`
	InactiveFile            string `json:"inactiveFile"`
	ActiveFile              string `json:"activeFile"`
	Unevictable             string `json:"unevictable"`
	HierarchicalMemoryLimit string `json:"hierarchicalMemoryLimit"`
	TotalCache              string `json:"totalCache"`
	TotalRSS                string `json:"totalRss"`
	TotalRSSHuge            string `json:"totalRssHuge"`
	TotalMappedFile         string `json:"totalMappedFile"`
	TotalPgpgIn             string `json:"totalPgpgin"`
	TotalPgpgOut            string `json:"totalPgpgout"`
	TotalPgFault            string `json:"totalPgfault"`
	TotalPgMajFault         string `json:"totalPgmajfault"`
	TotalInactiveAnon       string `json:"totalInactiveAnon"`
	TotalActiveAnon         string `json:"totalActiveAnon"`
	TotalInactiveFile       string `json:"totalInactiveFile"`
	TotalActiveFile         string `json:"totalActiveFile"`
	TotalUnevictable        string `json:"totalUnevictable"`
	MemUsage                string `json:"memUsage"`
	MemMaxUsage             string `json:"memMaxUsage"`
	MemLimit                string `json:"memoryLimit"`
	MemFailCnt              string `json:"memoryFailcnt"`
}

func (c containerMemStat) Text() containerMemStatFormat {
	cID := c.ContainerID
	cache := c.Cache.Text()
	rss := c.RSS.Text()
	rssH := c.RSSHuge.Text()
	mf := c.MappedFile.Text()
	pi := c.Pgpgin.Text()
	po := c.Pgpgout.Text()
	pf := c.Pgfault.Text()
	pjf := c.Pgmajfault.Text()
	ia := c.InactiveAnon.Text()
	aa := c.ActiveAnon.Text()
	ifile := c.InactiveFile.Text()
	afile := c.ActiveFile.Text()
	u := c.Unevictable.Text()
	hml := c.HierarchicalMemoryLimit.Text()
	tc := c.TotalCache.Text()
	trss := c.TotalRSS.Text()
	trssh := c.TotalRSSHuge.Text()
	tmf := c.TotalMappedFile.Text()
	tpi := c.TotalPgpgIn.Text()
	tpo := c.TotalPgpgOut.Text()
	tpf := c.TotalPgFault.Text()
	tpfg := c.TotalPgMajFault.Text()
	tia := c.TotalInactiveAnon.Text()
	taa := c.TotalActiveAnon.Text()
	tifile := c.TotalInactiveFile.Text()
	tafile := c.TotalActiveFile.Text()
	tu := c.TotalUnevictable.Text()
	mu := c.MemUsage.Text()
	mmu := c.MemMaxUsage.Text()
	ml := c.MemLimit.Text()
	mfc := c.MemFailCnt.Text()

	return containerMemStatFormat{cID, cache, rss, rssH, mf, pi, po, pf, pjf, ia, aa, ifile, afile, u, hml, tc, trss, trssh, tmf, tpi, tpo, tpf, tpfg, tia, taa, tifile, tafile, tu, mu, mmu, ml, mfc}
}

func (c containerMemStat) Format() containerMemStatFormat {
	cID := c.ContainerID
	cache := c.Cache.Format()
	rss := c.RSS.Format()
	rssH := c.RSSHuge.Format()
	mf := c.MappedFile.Format()
	pi := c.Pgpgin.Format()
	po := c.Pgpgout.Format()
	pf := c.Pgfault.Format()
	pjf := c.Pgmajfault.Format()
	ia := c.InactiveAnon.Format()
	aa := c.ActiveAnon.Format()
	ifile := c.InactiveFile.Format()
	afile := c.ActiveFile.Format()
	u := c.Unevictable.Format()
	hml := c.HierarchicalMemoryLimit.Format()
	tc := c.TotalCache.Format()
	trss := c.TotalRSS.Format()
	trssh := c.TotalRSSHuge.Format()
	tmf := c.TotalMappedFile.Format()
	tpi := c.TotalPgpgIn.Format()
	tpo := c.TotalPgpgOut.Format()
	tpf := c.TotalPgFault.Format()
	tpfg := c.TotalPgMajFault.Format()
	tia := c.TotalInactiveAnon.Format()
	taa := c.TotalActiveAnon.Format()
	tifile := c.TotalInactiveFile.Format()
	tafile := c.TotalActiveFile.Format()
	tu := c.TotalUnevictable.Format()
	mu := c.MemUsage.Format()
	mmu := c.MemMaxUsage.Format()
	ml := c.MemLimit.Format()
	mfc := c.MemFailCnt.Format()

	return containerMemStatFormat{cID, cache, rss, rssH, mf, pi, po, pf, pjf, ia, aa, ifile, afile, u, hml, tc, trss, trssh, tmf, tpi, tpo, tpf, tpfg, tia, taa, tifile, tafile, tu, mu, mmu, ml, mfc}
}

func newContainerMemStat(c *docker.CgroupMemStat) containerMemStat {
	cID := c.ContainerID
	cache := byte(c.Cache)
	rss := number(c.RSS)
	rssH := number(c.RSSHuge)
	mf := number(c.MappedFile)
	pi := number(c.Pgpgin)
	po := number(c.Pgpgout)
	pf := number(c.Pgfault)
	pjf := number(c.Pgmajfault)
	ia := number(c.InactiveAnon)
	aa := number(c.ActiveAnon)
	ifile := number(c.InactiveFile)
	afile := number(c.ActiveFile)
	u := byte(c.Unevictable)
	hml := byte(c.HierarchicalMemoryLimit)
	tc := byte(c.TotalCache)
	trss := number(c.TotalRSS)
	trssh := number(c.TotalRSSHuge)
	tmf := number(c.TotalMappedFile)
	tpi := number(c.TotalPgpgIn)
	tpo := number(c.TotalPgpgOut)
	tpf := number(c.TotalPgFault)
	tpfg := number(c.TotalPgMajFault)
	tia := number(c.TotalInactiveAnon)
	taa := number(c.TotalActiveAnon)
	tifile := number(c.TotalInactiveFile)
	tafile := number(c.TotalActiveFile)
	tu := number(c.TotalUnevictable)
	mu := byte(c.MemUsageInBytes)
	mmu := byte(c.MemMaxUsageInBytes)
	ml := byte(c.MemLimitInBytes)
	mfc := byte(c.MemFailCnt)

	return containerMemStat{cID, cache, rss, rssH, mf, pi, po, pf, pjf, ia, aa, ifile, afile, u, hml, tc, trss, trssh, tmf, tpi, tpo, tpf, tpfg, tia, taa, tifile, tafile, tu, mu, mmu, ml, mfc}
}

// Containers info
func Containers(format bool) ([]ContainerStatFormat, error) {
	var containers []ContainerStatFormat

	dock, err := docker.GetDockerStat()
	if err != nil {
		return containers, errors.Convert(err)
	}

	for _, container := range dock {
		cpuS, err := docker.CgroupCPUDocker(container.ContainerID)
		if err != nil {
			return containers, errors.Convert(err)
		}
		per, err := docker.CgroupCPUUsageDocker(container.ContainerID)
		// per := 0.00
		if err != nil {
			return containers, errors.Convert(err)
		}
		cpuStat := newContainerCPUStat(per, cpuS)

		cpuM, err := docker.CgroupMemDocker(container.ContainerID)
		if err != nil {
			return containers, errors.Convert(err)
		}
		cpuMem := newContainerMemStat(cpuM)

		dockerStat := ContainerStat{container.ContainerID, container.Name, container.Image, container.Status, container.Running, cpuStat, cpuMem}

		if format {
			containers = append(containers, dockerStat.Format())
		} else {
			containers = append(containers, dockerStat.Text())
		}
	}

	return containers, nil
}

// ContainersID return slice of container id
func ContainersID() ([]string, error) {
	containers, err := docker.GetDockerIDList()

	return containers, errors.Convert(err)
}
