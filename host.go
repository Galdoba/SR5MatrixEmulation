package main

import "math/rand"
import "fmt"

//"strconv"

var icMasterList []string

type Host struct {
	deviceRating   int
	attack         int
	sleaze         int
	dataProcessing int
	firewall       int
	grid           string
}

func createHost(deviceRating int) {
	var host Host
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
	pickICforHost(&host, icMasterList)
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

func pickICforHost(host *Host, icMasterList []string) {
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
	//fmt.Println(hostICList)

}

/*tatic void pickICForHost() {
	String[] pickIC = new String[icList.size()];
	activeICList.clear();
	int hRtng = getHostRating();
	activeICList.add("Patrol IC");
	//loadedICList.add("Patrol IC");
	System.out.println("Patrol IC");
	for (int a=1; a < hRtng; a++) {
		if (icList.size() > 0) {
			int index = new Random().nextInt(icList.size());
			pickIC[index] = icList.get(index);
			activeICList.add(pickIC[index]);
			icList.remove(index);
			System.out.println(pickIC[index]);
		}
	}
}*/
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
