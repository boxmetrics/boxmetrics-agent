package info

// GeneralStatFormat def
type GeneralStatFormat struct {
	CPU        CPUStatFormat            `json:"cpu"`
	Disks      []DiskStatFormat         `json:"disks"`
	Containers []ContainerStatFormat    `json:"containers"`
	Host       HostStatFormat           `json:"host"`
	Memory     MemoryStatFormat         `json:"memory"`
	Network    NetworkStatFormat        `json:"network"`
	Processes  []ProcessLightStatFormat `json:"processes"`
}

// General system information
func General(format bool) (GeneralStatFormat, error) {

	cpu, _ := CPU(format)

	disk, _ := Disks(format)

	containers, _ := Containers(format)

	host, _ := Host(format)

	memory, _ := Memory(format)

	network, _ := Network(format)

	processes, _ := Processes(format)

	return GeneralStatFormat{cpu, disk, containers, host, memory, network, processes}, nil
}
