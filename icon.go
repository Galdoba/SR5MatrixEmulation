package main

import (
	"fmt"
	
)

var newID int

type Icon struct {
	id              int
    deviceRating    int
	initiative      int
	mcm             int //Matrix Condition Monitor
    maxMCM          int
    isPlayer        bool
}

func createIcon(deviceRating int) Icon {
	var newIcon Icon
	newIcon.id = newID
    newID++
	newIcon.deviceRating = deviceRating
    newIcon.initiative = -1
    newIcon.maxMCM = 8 + newIcon.deviceRating/2
    newIcon.mcm = newIcon.maxMCM
    //newIcon.isPlayer = false
    fmt.Println("Created new Icon:", newIcon)
	return newIcon
}

func (icon *Icon) setIconID (id int) {
    icon.id = id
}

func (icon *Icon) getIconID () int {
    return icon.id
}

func (icon *Icon) setIconDeviceRating (deviceRating int) {
    icon.deviceRating = deviceRating
}

func (icon *Icon) getIconDeviceRating () int {
    return icon.deviceRating
}

func (icon *Icon) setIconInitiative (initiative int) {
    icon.initiative = initiative
}

func (icon *Icon) getIconInitiative () int {
    return icon.initiative
}

func (icon *Icon) setIconMaxMCM (maxMCM int) {
    icon.maxMCM = maxMCM
}

func (icon *Icon) getIconMaxMCM () int {
    return icon.maxMCM
}

func (icon *Icon) setIconMcm (mcm int) {
    icon.mcm = mcm
}

func (icon *Icon) getIconMcm () int {
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

func createPersona() Icon {
	var newIcon Icon
	newIcon.id = newID
    newID++
	newIcon.deviceRating = 6
    newIcon.initiative = -1
    newIcon.maxMCM = 8 + newIcon.deviceRating/2
    newIcon.mcm = newIcon.maxMCM
    newIcon.isPlayer = true
    fmt.Println("Created new Icon:", newIcon)
	return newIcon
}

func (icon *Icon) rollInitiative () int{
    fmt.Println("Icon №", icon.getIconID(), "rolls for initiative..." )
    init := icon.getIconDeviceRating() * 2 + xd6Test(4)
    return init
}