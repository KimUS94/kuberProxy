package kuberclient

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func executeCommad(cmdType string, args ...string) string {
	var res string

	cmd := exec.Command(cmdType, args...)
	output, err := cmd.Output()
	if err != nil {
		res = string("err")
	} else {
		res = string(output)
	}
	return res
}

func getNodeCount() int {
	//kubectl/get/nodes/-o=jsonpath={.items[*].metadata.name}

	cnt := 0
	output := executeCommad("kubectl", "get", "nodes", "-o=jsonpath={.items[*].metadata.name}")

	if len(output) > 0 {
		arr := strings.Split(output, " ")
		cnt = len(arr)
	}

	return cnt
}

func getNodeName(idx int) string {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].metadata.name}", idx)
	output := executeCommad("kubectl", "get", "nodes", jpath)

	return output
}

func getNodeStatus(name string) string {
	output := executeCommad("kubectl", "get", "node", name)
	temp := strings.Split(output, "\n")
	temp = strings.Fields(temp[1])

	return temp[1]
}

func getNodeAddress(idx int) string {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.addresses[0].address}", idx)
	output := executeCommad("kubectl", "get", "nodes", jpath)

	return output
}

func getNodeHostname(idx int) string {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.addresses[1].address}", idx)
	output := executeCommad("kubectl", "get", "nodes", jpath)

	return output
}

func getNodeAllocatable(idx int) *NodeAllocatable {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.cpu}", idx)
	cpu := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.ephemeral-storage}", idx)
	ephemeral_storage := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.hugepages-1Gi}", idx)
	hugepages_1Gi := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.hugepages-2Mi}", idx)
	hugepages_2Mi := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.memory}", idx)
	memory := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.allocatable.pods}", idx)
	pods := executeCommad("kubectl", "get", "nodes", jpath)

	var allocatable NodeAllocatable
	allocatable.Cpu = cpu
	allocatable.Ephemeral_storage = ephemeral_storage
	allocatable.Hugepages_1Gi = hugepages_1Gi
	allocatable.Hugepages_2Mi = hugepages_2Mi
	allocatable.Memory = memory
	allocatable.Pods = pods

	return &allocatable
}

func getNodeCapacity(idx int) *NodeCapacity {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.cpu}", idx)
	cpu := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.ephemeral-storage}", idx)
	ephemeral_storage := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.hugepages-1Gi}", idx)
	hugepages_1Gi := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.hugepages-2Mi}", idx)
	hugepages_2Mi := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.memory}", idx)
	memory := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.capacity.pods}", idx)
	pods := executeCommad("kubectl", "get", "nodes", jpath)

	var capacity NodeCapacity
	capacity.Cpu = cpu
	capacity.Ephemeral_storage = ephemeral_storage
	capacity.Hugepages_1Gi = hugepages_1Gi
	capacity.Hugepages_2Mi = hugepages_2Mi
	capacity.Memory = memory
	capacity.Pods = pods

	return &capacity
}

func getNodeCondition(idx int) *NodeCondition {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.conditions[0].status}", idx)
	net := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.conditions[1].status}", idx)
	mem := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.conditions[2].status}", idx)
	disk := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.conditions[3].status}", idx)
	pid := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.conditions[4].status}", idx)
	ready := executeCommad("kubectl", "get", "nodes", jpath)

	var condition NodeCondition
	condition.NetworkUnavailable = net
	condition.MemoryPressure = mem
	condition.DiskPressure = disk
	condition.PIDPressure = pid
	condition.Ready = ready

	return &condition
}

func getNodeSysteminfo(idx int) *NodeSysteminfo {
	jpath := fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.architecture}", idx)
	architecture := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.bootID}", idx)
	bootID := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.containerRuntimeVersion}", idx)
	containerRuntimeVersion := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.kernelVersion}", idx)
	kernelVersion := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.kubeProxyVersion}", idx)
	kubeProxyVersion := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.kubeletVersion}", idx)
	kubeletVersion := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.machineID}", idx)
	machineID := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.operatingSystem}", idx)
	operatingSystem := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.osImage}", idx)
	osImage := executeCommad("kubectl", "get", "nodes", jpath)

	jpath = fmt.Sprintf("-o=jsonpath={.items[%d].status.nodeInfo.systemUUID}", idx)
	systemUUID := executeCommad("kubectl", "get", "nodes", jpath)

	var sysInfo NodeSysteminfo
	sysInfo.Architecture = architecture
	sysInfo.BootID = bootID
	sysInfo.ContainerRuntimeVersion = containerRuntimeVersion
	sysInfo.KernelVersion = kernelVersion
	sysInfo.KubeProxyVersion = kubeProxyVersion
	sysInfo.KubeletVersion = kubeletVersion
	sysInfo.MachineID = machineID
	sysInfo.OperatingSystem = operatingSystem
	sysInfo.OsImage = osImage
	sysInfo.SystemUUID = systemUUID

	return &sysInfo
}

func getNodeUsage(name string) *nodeUsage {
	//kubectl top nodes --use-protocol-buffers

	var usage nodeUsage
	output := executeCommad("kubectl", "top", "nodes", name, "--use-protocol-buffers")

	if len(output) <= 0 {
		return &usage
	}

	if output == "err" {
		usage.Cpu_cores = "Unknown"
		usage.Cpu_usage = "Unknown"
		usage.Mem_cores = "Unknown"
		usage.Mem_usage = "Unknown"
		return &usage
	}

	arr := strings.Split(output, "\n")
	if len(arr) < 2 {
		return &usage
	}

	output = arr[1]
	arr = strings.Fields(output)
	if len(arr) < 5 {
		return nil
	}

	usage.Cpu_cores = arr[1]
	usage.Cpu_usage = arr[2]
	usage.Mem_cores = arr[3]
	usage.Mem_usage = arr[4]

	return &usage
}

func getPodList(nodes *map[string]NodeInfo) {
	// var podInfo PodInfo

	pods := map[string]PodInfo{}

	output := executeCommad("kubectl", "get", "pods", "--all-namespaces", "-o", "wide")
	if len(output) <= 0 {
		// return nil
		return
	}

	arr := strings.Split(output, "\n")
	if len(arr) <= 0 {
		// return nil
		return
	}

	// var table map[string][]PodInfo
	table := map[string][]PodInfo{}

	for i := 1; i < len(arr); i++ {
		column := strings.Fields(arr[i])
		if len(column) <= 0 {
			break

		}
		var podInfo PodInfo
		podInfo.Namespace = column[0]
		podInfo.Name = column[1]
		podInfo.Ready = column[2]
		podInfo.Status = column[3]
		podInfo.Restarts = column[4]
		podInfo.Age = column[5]
		podInfo.IP = column[6]
		podInfo.Node = column[7]
		podInfo.NominatedNode = column[8]
		podInfo.ReadinessGates = column[9]

		pods[column[1]] = podInfo

		table[podInfo.Node] = append(table[podInfo.Node], podInfo)

		fmt.Println((*nodes)[podInfo.Node])
	}

	for key, val := range table {
		node := (*nodes)[key]
		node.Pods = val
		node.PodsCnt = strconv.Itoa(len(val))

		(*nodes)[key] = node
	}
}
