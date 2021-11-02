package kuberclient

func getNodeInfo() *map[string]NodeInfo {
	nodes := map[string]NodeInfo{}

	cnt := getNodeCount()

	for idx := 0; idx < cnt; idx++ {
		name := getNodeName(idx)
		stat := getNodeStatus(name)
		addr := getNodeAddress(idx)
		host := getNodeHostname(idx)
		usag := getNodeUsage(name)
		allc := getNodeAllocatable(idx)
		capa := getNodeCapacity(idx)
		cndi := getNodeCondition(idx)
		sysi := getNodeSysteminfo(idx)

		var node NodeInfo
		node.Name = name
		node.Status = stat
		node.Address = addr
		node.Hostname = host

		node.Usage.Cpu_cores = usag.Cpu_cores
		node.Usage.Cpu_usage = usag.Cpu_usage
		node.Usage.Mem_cores = usag.Mem_cores
		node.Usage.Mem_usage = usag.Mem_usage

		node.Allocatable.Cpu = allc.Cpu
		node.Allocatable.Ephemeral_storage = allc.Ephemeral_storage
		node.Allocatable.Hugepages_1Gi = allc.Hugepages_1Gi
		node.Allocatable.Hugepages_2Mi = allc.Hugepages_2Mi
		node.Allocatable.Memory = allc.Memory
		node.Allocatable.Pods = allc.Pods

		node.Capacity.Cpu = capa.Cpu
		node.Capacity.Ephemeral_storage = capa.Ephemeral_storage
		node.Capacity.Hugepages_1Gi = capa.Hugepages_1Gi
		node.Capacity.Hugepages_2Mi = capa.Hugepages_2Mi
		node.Capacity.Memory = capa.Memory
		node.Capacity.Pods = capa.Pods

		node.Condition.NetworkUnavailable = cndi.NetworkUnavailable
		node.Condition.MemoryPressure = cndi.MemoryPressure
		node.Condition.DiskPressure = cndi.DiskPressure
		node.Condition.PIDPressure = cndi.PIDPressure
		node.Condition.Ready = cndi.Ready

		node.SystemInfo.Architecture = sysi.Architecture
		node.SystemInfo.BootID = sysi.BootID
		node.SystemInfo.ContainerRuntimeVersion = sysi.ContainerRuntimeVersion
		node.SystemInfo.KernelVersion = sysi.KernelVersion
		node.SystemInfo.KubeProxyVersion = sysi.KubeProxyVersion
		node.SystemInfo.KubeletVersion = sysi.KubeletVersion
		node.SystemInfo.MachineID = sysi.MachineID
		node.SystemInfo.OperatingSystem = sysi.OperatingSystem
		node.SystemInfo.OsImage = sysi.OsImage
		node.SystemInfo.SystemUUID = sysi.SystemUUID

		nodes[name] = node
	}

	getPodList(&nodes)

	return &nodes

}

// func getPodInfo() *map[string]PodInfo {
// 	return getPodList()
// }
