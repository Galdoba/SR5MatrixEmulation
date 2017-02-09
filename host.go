package main

import "math/rand"
import "fmt"

//"strconv"

var icMasterList []string
var host Host
var ic IC

type IC struct {
	icName string
	isLoaded bool
	deviceRating   int
	attack         int
	sleaze         int
	dataProcessing int
	firewall       int
}

type Host struct {
	deviceRating   int
	attack         int
	sleaze         int
	dataProcessing int
	firewall       int
	grid           string
	icArray		   []IC
	isCreated	   bool
}

func createHost(deviceRating int) *Host {
	icMasterList = generateICMasterList()
	setSeed()
	if deviceRating > 0 && deviceRating < 13 {
		host.deviceRating = deviceRating
	} else {
		host.deviceRating = rand.Intn(12) + 1
	}
	atributeArray := []int{0, 1, 2, 3}
	for i := rand.Intn(100); i > 0; i-- {
		shuffleInt(atributeArray)
	}
	//shuffleInt(atributeArray)
	host.setHostAttack(atributeArray[0] + host.getHostRating())
	host.setHostSleaze(atributeArray[1] + host.getHostRating())
	host.setHostDataProcessing(atributeArray[2] + host.getHostRating())
	host.setHostFirewall(atributeArray[3] + host.getHostRating())
	//atributeArray = 0, 1, 2, 3

	fmt.Println("host: rating/ATT/SLZ/DTPROSS/FRWALL")
	fmt.Println(host)
	hostICList := pickICforHost(&host, icMasterList)
	host.icArray = make([]IC, 0)
	for i := range hostICList {
		newIC := createIC(hostICList, i)
		host.icArray = append(host.icArray, *newIC)
	}
	host.icArray[0].isLoaded = true
	host.isCreated = true
	return &host
}

func createIC(hostICList []string, i int) *IC {
	var localIC IC
	localIC.icName = hostICList[i]
	localIC.deviceRating = host.getHostRating()
	localIC.attack = host.getHostAttack()
	localIC.sleaze = host.getHostSleaze()
	localIC.dataProcessing = host.getHostDataProcessing()
	localIC.firewall = host.getHostFirewall()
	localIC.isLoaded = false
	return &localIC
}

func generateICMasterList() []string {
	icMasterList := make([]string, 0)
	icMasterList = append(icMasterList, "Acid_IC")
	icMasterList = append(icMasterList, "Binder_IC")
	icMasterList = append(icMasterList, "Black_IC")
	icMasterList = append(icMasterList, "Blaster_IC")
	icMasterList = append(icMasterList, "Bloodhound_IC")
	icMasterList = append(icMasterList, "Catapult_IC")
	icMasterList = append(icMasterList, "Crash_IC")
	icMasterList = append(icMasterList, "Jammer_IC")
	icMasterList = append(icMasterList, "Killer_IC")
	icMasterList = append(icMasterList, "Marker_IC")
	icMasterList = append(icMasterList, "Probe_IC")
	icMasterList = append(icMasterList, "Scramble_IC")
	icMasterList = append(icMasterList, "Shocker_IC")
	icMasterList = append(icMasterList, "Sparky_IC")
	icMasterList = append(icMasterList, "Tar_Baby_IC")
	icMasterList = append(icMasterList, "Track_IC")
	//fmt.Println(icMasterList)
	return icMasterList
}

func pickICforHost(host *Host, icMasterList []string) []string {
	activeIClist := make([]string, 0)
	activeIClist = icMasterList
	for i := rand.Intn(100); i > 0; i-- {
		shuffleString(activeIClist)
	}
	hostICList := make([]string, 0)
	hostICList = append(hostICList, "Patrol_IC")
	for i := 1; i < host.getHostRating(); i++ {
		hostICList = append(hostICList, activeIClist[i])
	}
	for i := range hostICList {
		fmt.Println("IC #", i+1, "is", hostICList[i])
	}
	return hostICList
}

func (host *Host) setHostRating(deviceRating int) {
	host.deviceRating = deviceRating
}

func (host *Host) getHostRating() int {
	return host.deviceRating
}

func (host *Host) setHostAttack(attack int) {
	host.attack = attack
}

func (host *Host) getHostAttack() int {
	return host.attack
}

func (host *Host) setHostSleaze(sleaze int) {
	host.sleaze = sleaze
}

func (host *Host) getHostSleaze() int {
	return host.sleaze
}

func (host *Host) setHostDataProcessing(dataProcessing int) {
	host.dataProcessing = dataProcessing
}

func (host *Host) getHostDataProcessing() int {
	return host.dataProcessing
}

func (host *Host) setHostFirewall(firewall int) {
	host.firewall = firewall
}

func (host *Host) getHostFirewall() int {
	return host.firewall
}

func shuffleInt(atributeArray []int) {

	for i := len(atributeArray) - 1; i > 0; i-- {
		j := rand.Intn(i)
		atributeArray[i], atributeArray[j] = atributeArray[j], atributeArray[i]
	}
}

func shuffleString(atributeArray []string) {

	for i := len(atributeArray) - 1; i > 0; i-- {
		j := rand.Intn(i)
		atributeArray[i], atributeArray[j] = atributeArray[j], atributeArray[i]
	}
}

func hostAction() {
	proceed := false
	canLoad := 0
	setSeed()
	//host.icArray[0].isLoaded = false
	if host.icArray[0].isLoaded == false {
		newIC := createICIcon(0)
		masterIconList.iconArray = append(masterIconList.iconArray, *newIC)		
	} else {
		for j := range host.icArray {
			if host.icArray[j].isLoaded == false {
				canLoad++ 
			}
		}
		for proceed == false && canLoad > 0 {	
			i := rand.Intn(host.getHostRating())
	//		fmt.Println("choose", i, "can pick", canLoad)
			if host.icArray[i].isLoaded == false {
				newIC := createICIcon(i)
				masterIconList.iconArray = append(masterIconList.iconArray, *newIC)
				host.icArray[i].isLoaded = true
				break
			}
		}
	}
}