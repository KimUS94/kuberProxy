package kuberclient

type NodeInfo struct {
	Name        string          `json:"name"`
	Status      string          `json:"status"`
	Address     string          `json:"address"`
	Hostname    string          `json:"hostname"`
	Usage       nodeUsage       `json:"usage"`
	Allocatable NodeAllocatable `json:"allocatable"`
	Capacity    NodeCapacity    `json:"capacity"`
	Condition   NodeCondition   `json:"condition"`
	SystemInfo  NodeSysteminfo  `json:"systemInfo"`
	Pods        []PodInfo       `json:"pods"`
	PodsCnt     string          `json:"podscnt"`
}

type PodInfo struct {
	Name           string `json:"name"`
	Namespace      string `json:"namespace"`
	Ready          string `json:"ready"`
	Status         string `json:"status"`
	Restarts       string `json:"restarts"`
	Age            string `json:"age"`
	IP             string `json:"ip"`
	Node           string `json:"Node"`
	NominatedNode  string `json:"nominatedNode"`
	ReadinessGates string `json:"readinessGates"`
}

type NodeAllocatable struct {
	Cpu               string `json:"cpu"`
	Ephemeral_storage string `json:"ephemeral_storage"`
	Hugepages_1Gi     string `json:"hugepages_1Gi"`
	Hugepages_2Mi     string `json:"hugepages_2Mi"`
	Memory            string `json:"memory"`
	Pods              string `json:"pods"`
}

type NodeCapacity struct {
	Cpu               string `json:"cpu"`
	Ephemeral_storage string `json:"ephemeral_storage"`
	Hugepages_1Gi     string `json:"hugepages_1Gi"`
	Hugepages_2Mi     string `json:"hugepages_2Mi"`
	Memory            string `json:"memory"`
	Pods              string `json:"pods"`
}

type NodeCondition struct {
	NetworkUnavailable string `json:"NetworkUnavailable"`
	MemoryPressure     string `json:"MemoryPressure"`
	DiskPressure       string `json:"DiskPressure"`
	PIDPressure        string `json:"PIDPressure"`
	Ready              string `json:"Ready"`
}

type NodeSysteminfo struct {
	Architecture            string `json:"architecture"`
	BootID                  string `json:"bootID"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KernelVersion           string `json:"kernelVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	MachineID               string `json:"machineID"`
	OperatingSystem         string `json:"operatingSystem"`
	OsImage                 string `json:"osImage"`
	SystemUUID              string `json:"systemUID"`
}

type nodeUsage struct {
	Cpu_cores string `json:"cpu_cores"`
	Cpu_usage string `json:"cpu_usage"`
	Mem_cores string `json:"mem_cores"`
	Mem_usage string `json:"mem_usage"`
}
