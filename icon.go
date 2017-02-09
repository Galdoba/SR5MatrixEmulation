package main

import (
	"fmt"
	"strconv"
	"strings"
)

var newID int

type Icon struct {
	id             int
	name           string
	realName	   string
	iconType       string
	deviceRating   int
	attack         int
	sleaze         int
	dataProcessing int
	firewall       int
	initiative     int
	mcm            int //Matrix Condition Monitor
	maxMCM         int
	isPlayer       bool
}

func createIcon(deviceRating int) *Icon {
	var newIcon Icon
	newIcon.id = newID
	newID++
	newIcon.iconType = "Icon"
	newIcon.name = newIcon.iconType + strconv.Itoa(newIcon.id)
	newIcon.deviceRating = deviceRating
	newIcon.attack = deviceRating
	newIcon.sleaze = deviceRating
	newIcon.dataProcessing = deviceRating
	newIcon.firewall = deviceRating
	newIcon.initiative = -1
	newIcon.maxMCM = 8 + newIcon.deviceRating/2
	newIcon.mcm = newIcon.maxMCM
	//newIcon.isPlayer = false
	fmt.Println("Created new Icon:", newIcon)
	return &newIcon
}

func (icon *Icon) setIconID(id int) {
	icon.id = id
}

func (icon *Icon) getIconID() int {
	return icon.id
}

func (icon *Icon) setIconName(name string) {
	icon.name = name
}

func (icon *Icon) getIconName() string {
	return icon.name
}

func (icon *Icon) setIconRealName(realName string) {
	icon.realName = realName
}

func (icon *Icon) getIconRealName() string {
	return icon.realName
}

func (icon *Icon) setIconType(iconType string) {
	icon.iconType = iconType
}

func (icon *Icon) getIconType() string {
	return icon.iconType
}

func (icon *Icon) setIconDeviceRating(deviceRating int) {
	icon.deviceRating = deviceRating
}

func (icon *Icon) getIconDeviceRating() int {
	return icon.deviceRating
}

func (icon *Icon) setIconAttack(attack int) {
	icon.attack = attack
}

func (icon *Icon) getIconAttack() int {
	return icon.attack
}

func (icon *Icon) setIconSleaze(sleaze int) {
	icon.sleaze = sleaze
}

func (icon *Icon) getIconSleaze() int {
	return icon.sleaze
}

func (icon *Icon) setIconDataProcessing(dataProcessing int) {
	icon.dataProcessing = dataProcessing
}

func (icon *Icon) getIconDataProcessing() int {
	return icon.dataProcessing
}

func (icon *Icon) setIconFirewall(firewall int) {
	icon.firewall = firewall
}

func (icon *Icon) getIconFirewall() int {
	return icon.firewall
}

func (icon *Icon) setIconInitiative(initiative int) {
	icon.initiative = initiative
}

func (icon *Icon) getIconInitiative() int {
	return icon.initiative
}

func (icon *Icon) setIconMaxMCM(maxMCM int) {
	icon.maxMCM = maxMCM
}

func (icon *Icon) getIconMaxMCM() int {
	return icon.maxMCM
}

func (icon *Icon) setIconMcm(mcm int) {
	icon.mcm = mcm
}

func (icon *Icon) getIconMcm() int {
	return icon.mcm
}

/*// SetName receives a pointer to Foo so it can modify it.
func (f *Foo) SetName(name string) {
    f.name = name
}

// Name receives a copy of Foo since it doesn't need to modify it.
func (f Foo) Name() string {
    return f.name
}

func main() {
    // Notice the Foo{}. The new(Foo) was just a syntactic sugar for &Foo{}
    // and we don't need a pointer to the Foo, so I replaced it.
    // Not relevant to the problem, though.
    p := Foo{}
    p.SetName("Abc")
    name := p.Name()
    fmt.Println(name)
}*/

func createPersona() *Icon {
	var newIcon Icon
	newIcon.id = newID
	newID++
	newIcon.iconType = "Persona"
	newIcon.name = newIcon.iconType + strconv.Itoa(newIcon.id)
	newIcon.deviceRating = 6
	newIcon.attack = 5
	newIcon.sleaze = 5
	newIcon.dataProcessing = 5
	newIcon.firewall = 5
	newIcon.initiative = 1
	newIcon.maxMCM = 8 + newIcon.deviceRating/2 + 1000
	newIcon.mcm = newIcon.maxMCM
	newIcon.isPlayer = true
	fmt.Println("Created new Icon:", newIcon)
	return &newIcon
}

func createHostIcon(comm3 string) *Icon {
	var newIcon Icon
	var data []string
	var host *Host
	data = strings.SplitN(comm3, ":",2)
	if len(data) < 2 {
		i := 0
		fmt.Println(i)
		host = createHost(i)
	} else if i, err := strconv.Atoi(data[1]); err == nil {
    	fmt.Printf("i=%d, type: %T\n", i, i)
		host = createHost(i)
	}
	newIcon.id = newID
	newID++
	newIcon.iconType = "Host"
	if comm3 != "random" {
		newIcon.name = data[0]	
	} else {
		newIcon.name = newIcon.iconType + strconv.Itoa(newIcon.id)	
	}
	newIcon.deviceRating = host.getHostRating()
	newIcon.attack = host.getHostAttack()
	newIcon.sleaze = host.getHostSleaze()
	newIcon.dataProcessing = host.getHostDataProcessing()
	newIcon.firewall = host.getHostFirewall()
	newIcon.initiative = -1
	newIcon.maxMCM = 999999
	newIcon.mcm = 999999
	newIcon.isPlayer = false
	fmt.Println("Host Icon Created")
	//fmt.Println(createICIcon(0))
	fmt.Println(newIcon)
	return &newIcon
}

func createICIcon(i int) *Icon {
	var newIcon Icon
	newIcon.id = newID
	newID++
	newIcon.iconType = "IC"
	newIcon.name = newIcon.iconType + strconv.Itoa(newIcon.id)	
	newIcon.realName = host.icArray[i].icName
	newIcon.deviceRating = host.getHostRating()
	newIcon.attack = host.getHostAttack()
	newIcon.sleaze = host.getHostSleaze()
	newIcon.dataProcessing = host.getHostDataProcessing()
	newIcon.firewall = host.getHostFirewall()
	newIcon.initiative = -1
	newIcon.maxMCM = 8 + newIcon.getIconDeviceRating()/2
	newIcon.mcm = 8 + newIcon.getIconDeviceRating()/2
	newIcon.isPlayer = false
	fmt.Println("Host Icon Created")
	fmt.Println(newIcon)
	return &newIcon
}

func createGridIcon() *Icon {
	var newIcon Icon
	newIcon.id = newID
	newID++
	newIcon.iconType = "Grid"
	newIcon.name = "GridIcon"
	newIcon.deviceRating = 3
	newIcon.dataProcessing = 3
	newIcon.firewall = 3
	newIcon.initiative = -1
	newIcon.maxMCM = 999999
	newIcon.mcm = 999999
	newIcon.isPlayer = false
	fmt.Println("Grid Icon Created")
	return &newIcon
}

func (icon *Icon) rollInitiative() int {
	fmt.Println("Icon №", icon.getIconID(), "rolls for initiative...")
	init := icon.getIconDeviceRating()*2 + xd6Test(4)
	return init
}
