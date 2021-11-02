package kuberclient

type Cluster struct {
	Nodes *map[string]NodeInfo
}

func (c *Cluster) SetCluster() {
	c.Nodes = getNodeInfo()
}

func (c *Cluster) SetNodeUsage() {
	for key, val := range *c.Nodes {
		usage := *getNodeUsage(key)

		val.Usage.Cpu_cores = usage.Cpu_cores
		val.Usage.Cpu_usage = usage.Cpu_usage
		val.Usage.Mem_cores = usage.Mem_cores
		val.Usage.Mem_usage = usage.Mem_usage
	}
}

func (c *Cluster) SetPodInfo() {
	// getPodList(&nodes)
	getPodList(c.Nodes)
}
