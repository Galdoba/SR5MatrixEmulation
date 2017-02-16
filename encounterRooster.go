package main

import (
	"fmt"
	"math/rand"
	//"bufio"
	"os"
	"strconv"
	"strings"
)

var actionValid bool
var order []int
var orderID []int
var actionName string
var masterIconList IconList
var targetList IconList
var iconSource Icon
var iconTarget Icon
var iconNil Icon
var turn int

type IconList struct {
	iconArray []Icon
	size      int
	isOk      bool
}

func createRooster() {
	markList = createMarkList()
	matrixActionList = createMatrixActionList()
	personaActionList = createPersonaActList()
	iconActionList = createIconActList()
//	var comm []string
	actionValid = false
	markList = createMarkList()

	fmt.Println("Start Creating Rooster")
	masterIconList = makeIconList()
	targetList = makeTargetList()
	masterIconList = addIcon(masterIconList, createGridIcon()) //Икона сети (надо расширить до остальных)
	fmt.Println("Add Icons")
	
	mainBody()
}

func mainBody () {
	var comm []string
	//go station()
	for masterIconList.iconArray[0].getIconMcm() > 0 {
		fmt.Println(iconSource.getIconName(), "start action")
		fmt.Println(masterIconList.iconArray)
		if len(masterIconList.iconArray) > 128 {
			outputRed("maximum obcjects reached")
			outputRed("stop")
			os.Exit(3)
		}
		//markList = updateMarks()
		order = makeCombatOrder()
		//fmt.Println(order, "order before pick")
		iconSource = pickIconSource(order)
		createTargetList()
		if len(targetList.iconArray) < 1 {
			outputRed("WARNING!!! NO TARGETS SPOTTED...")
			comm[1] = "HOLD"
			comm[2] = iconSource.getIconName()
		} else {
			pickIconTarget(targetList)
		}
		markList = updateMarks()
		if iconSource.isPlayer == true {
			outputRed("there are " + strconv.Itoa(len(markList)) + " marks on the list")
			for actionValid == false {
				comm = userInput()
				iconTarget = asignIconTarget(comm[2])
				actionName, actionValid := chooseMatrixAction(iconSource, iconTarget, comm)
				//actionValid = checkMarksQty(iconSource, iconTarget, actionName)
				if actionValid == true {
					comm[1] = actionName
					if len(comm) == 4 {
						outputRed(comm[0] + ">" + comm[1] + ">" + comm[2] + ">" + comm[3])
					} else {
						outputRed(comm[0] + ">" + comm[1] + ">" + comm[2])
					}
					fmt.Println(comm)
					outputRed("command accepted...")
					outputRed("performing...")

				} else {
					outputRed(comm[0] + ">" + comm[1] + ">" + comm[2])
					outputRed("command rejected...")
					_, reason := checkMarksQty(iconSource, iconTarget, actionName)
					outputRed(reason)
				}
			}
		} else {
			comm = formCommand(iconSource, iconTarget)
			confirmCommand(comm)
			actionName, _ := chooseMatrixAction(iconSource, iconTarget, comm)
			comm[1] = actionName
			outputRed(comm[0] + ">" + comm[1] + ">" + comm[2])
		}
		actionValid = false
		//outputRed(comm[0] + ">" + comm[1] + ">" + comm[2])
		//нужен утвердитель команды

		/*		confirmCommand(comm)
				actionName, _ := chooseMatrixAction(iconSource, iconTarget, comm)
				comm[1] = actionName
				outputRed(comm[0] + ">" + comm[1] + ">" + comm[2])
		*/
		//comm[0] = strings.Replace(comm[0], " ", "_", -1)
    	//comm[0] = strings.ToUpper(comm[0])
    	//comm[1] = strings.Replace(comm[1], " ", "_", -1)
    	//comm[1] = strings.ToUpper(comm[1])
    	//comm[2] = strings.Replace(comm[2], " ", "_", -1)
    	//comm[2] = strings.ToUpper(comm[2])
		if comm[1] == "MATRIX_SEARCH" {
			comm[2] = strings.Replace(comm[2], " ", "_", -1)
    		comm[2] = strings.ToUpper(comm[2])
			if len(comm) == 3 {
				comm = append(comm, "random")
				comm[3] = "random"
			}
			
			
			switch comm[2] {
			case "HOST":
				//masterIconList = addIcon(masterIconList, createHostIcon(comm[3]))
				doMatrixSimpleAction(iconSource, actionName)
				masterIconList = addIcon(masterIconList, createHostIcon(comm[3]))
				PatrolIC := createICIcon(0)
				masterIconList.iconArray = append(masterIconList.iconArray, *PatrolIC)
			
			}
		} else {
			doMatrixOppAction(iconSource, iconTarget, actionName)
		}

		//comm = nil
		checkPlay()
		masterIconList = destroyIcon(masterIconList)
		markList = updateMarks()
		//fmt.Println(iconSource.getIconName(), "end action")
	}

}

func checkPlay() {
	assert(masterIconList.iconArray[0].getIconMcm() > 0, "Connection Terminated. player destroyed")
}

func pickIconTarget(targetList IconList) Icon {
	assert(targetList.isOk, "No targetList")
	i := rand.Intn(len(targetList.iconArray))
	iconTarget = targetList.iconArray[i]
	return iconTarget
}

func asignIconTarget(comm2 string) Icon {
	comm2 = strings.ToUpper(comm2)
	for i:= range masterIconList.iconArray {
		if comm2 == masterIconList.iconArray[i].getIconName() {
			iconTarget = masterIconList.iconArray[i]
			return iconTarget
		}
	}
	return iconTarget
}

func createTargetList() IconList {
	var icon Icon
	for i := range masterIconList.iconArray {
		if iconSource.isPlayer != masterIconList.iconArray[i].isPlayer {
			icon = masterIconList.iconArray[i]
			targetList = addIcon(targetList, &icon)
		}
	}
	return targetList
}

func makeCombatOrder() []int {
	var icon Icon
	size := getIconlistSize()
	order = make([]int, size)
	for i := range masterIconList.iconArray {
		icon = masterIconList.iconArray[i]
		if icon.getIconType() == "Host" || icon.getIconType() == "Grid" {
			icon.setIconInitiative(-100)
		
		}
		//assert(icon.getIconInitiative() > 0, "Initiative Less than 0")
		//icon.setIconInitiative(icon.rollInitiative())
		order[i] = icon.getIconInitiative()
		masterIconList.iconArray[i] = icon

	}
		fmt.Println("Order before sorting:", order)
	bubbleSort(order)
	fmt.Println("Order after sorting:", order)
	endPass()
	return order
}

func endPass () {
	if order[0] < 1 {
		turn++
		if turn != 1 {
			hostAction()	
		}
		allRollInitiative()
		bubbleSort(order)
	}
}

func pickIconSource(order []int) Icon {
	var icon Icon
	for i := range masterIconList.iconArray {
		icon = masterIconList.iconArray[i]
		if icon.getIconInitiative() == order[0] {
			iconSource = icon
			fmt.Println("ходит икона ", iconSource.getIconID())
			fmt.Println("order", order)
			return iconSource
		}
		//fmt.Println(i, icon.getIconInitiative(), order[0])
	}
	//fmt.Println("не смогли выбрать IconSource")
	//fmt.Println(masterIconList.iconArray)
	return iconSource
}

func confirmCommand(comm []string) {
	//утверждаем iconSource
	assert(comm[0] != "", "empty iconTarget")
	for i := range masterIconList.iconArray {
		if comm[0] == masterIconList.iconArray[i].getIconName() {
			iconSource = masterIconList.iconArray[i]
		}
	}
	//утверждаем iconTarget
	assert(comm[2] != "", "empty iconTarget")
	for i := range masterIconList.iconArray {
		if comm[2] == masterIconList.iconArray[i].getIconName() {
			iconTarget = masterIconList.iconArray[i]
		}
	}
	//утверждаем action //возможно нужен будет отдельный метод выбора действия
	assert(comm[1] != "", "empty action")
	for i := range matrixActionList {
		if comm[1] == matrixActionList[i] {
			actionName = comm[1]
		}
	}
	//утверждаем actionInfo
	if len(comm) > 3 {
		if comm[3] != "" {
			fmt.Println("добaвляем условие")
		} else {
			fmt.Println("условия нет и вообще этого не должно быть")

		}
	}

}

func doMatrixOppAction(iconSource Icon, iconTarget Icon, actionName string) { //должно быть еще название действия и механизмы выбора
	fmt.Println(masterIconList)
	fmt.Println("формируем дайспул")
	//строитель дайспула пойдет в отдельную функцию
	dicePoolSrc := iconSource.getIconDeviceRating() * 2
	dicePoolTrgt := iconTarget.getIconDeviceRating() * 2
	limit := iconSource.getIconDeviceRating()
	//
	netHits, _, _ := opposedTest(dicePoolSrc, dicePoolTrgt, limit)
	//распределение эффектов пойдет в отдельную функцию
	actionEffect(&actionName, &iconSource, &iconTarget, &netHits)

	iconSource.setIconInitiative(iconSource.getIconInitiative() - 10)
	renewIconSource(iconSource)
	renewIconTarget(iconTarget)
	targetList.iconArray = nil //зачищаем список целей в конце действия
	fmt.Println(masterIconList)
}

func doMatrixSimpleAction(iconSource Icon, actionName string) { //должно быть еще название действия и механизмы выбора
	fmt.Println(masterIconList)
	fmt.Println("формируем дайспул")
	//строитель дайспула пойдет в отдельную функцию
	dicePoolSrc := iconSource.getIconDeviceRating() * 2
	threshold := 1
	limit := iconSource.getIconDeviceRating()
	netHits, _, _ := simpleTest(dicePoolSrc, limit, threshold)
	//распределение эффектов пойдет в отдельную функцию
	actionEffect(&actionName, &iconSource, &iconTarget, &netHits)

	iconSource.setIconInitiative(iconSource.getIconInitiative() - 10)
	renewIconSource(iconSource)
	renewIconTarget(iconTarget)
	targetList.iconArray = nil //зачищаем список целей в конце действия
	fmt.Println(masterIconList)
}



func renewIconSource(iconSource Icon) {
	for i := range masterIconList.iconArray {
		if iconSource.getIconID() == masterIconList.iconArray[i].getIconID() {
			masterIconList.iconArray[i] = iconSource
			//fmt.Println(masterIconList.iconArray)
		}
	}
	resetIcon(iconSource)
}

func renewIconTarget(iconTarget Icon) {
	for i := range masterIconList.iconArray {
		if iconTarget.getIconID() == masterIconList.iconArray[i].getIconID() {
			masterIconList.iconArray[i] = iconTarget
			//fmt.Println(masterIconList.iconArray)
		}
	}
	resetIcon(iconTarget)
}

func resetIcon(icon Icon) Icon {
	icon.setIconID(-2)
	icon.setIconDeviceRating(-1)
	icon.setIconInitiative(-1)
	icon.setIconMcm(-1)
	return icon
}

func makeIconList() IconList {
	masterIconList.iconArray = make([]Icon, 0, 1)
	masterIconList.isOk = true
	masterIconList = addIcon(masterIconList, createPersona())
	return masterIconList
}

func makeTargetList() IconList {
	targetList.iconArray = make([]Icon, 0, 1)
	targetList.isOk = true
	return targetList
}

func addIcon(masterIconList IconList, newIcon *Icon) IconList {
	assert(masterIconList.isOk, "No Icon List")
	masterIconList.iconArray = append(masterIconList.iconArray, *newIcon)
	return masterIconList
}

func getIconlistSize() int {
	size := len(masterIconList.iconArray)
	return size
}

func bubbleSort(order []int) {
	// n is the number of items in our list
	n := len(order)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if order[i-1] < order[i] {

				// swap values using Go's tuple assignment
				order[i], order[i-1] = order[i-1], order[i]
				swapped = true
			}
		}
	}
}

func allRollInitiative() []int {
	order = nil
	for i := range masterIconList.iconArray {
		var icon Icon
		//order = nil
		icon = masterIconList.iconArray[i]
		icon.setIconInitiative(icon.rollInitiative())
		if icon.getIconType() == "Grid" || icon.getIconType() == "Host"{
			icon.setIconInitiative(-100)
		}
		masterIconList.iconArray[i].setIconInitiative(icon.getIconInitiative())
		fmt.Println(order, "/", i, "/", len(masterIconList.iconArray), "/", masterIconList.iconArray)
		order = append(order, masterIconList.iconArray[i].getIconInitiative())
	}
	return order
}


