package main

import (
	"fmt"
	"math/rand"
	//"bufio"
	//"os"
    
)

var order []int
var orderID []int
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
	matrixActionList = createMatrixActionList()
	var comm []string
	fmt.Println("Start Creating Rooster")
	masterIconList = makeIconList()
    targetList = makeTargetList()
	newIcon := createIcon(1)
	masterIconList = addIcon(masterIconList, newIcon)
	masterIconList = addIcon(masterIconList, createIcon(2))
	masterIconList = addIcon(masterIconList, createIcon(2))
	fmt.Println("Add Icons")
	fmt.Println(masterIconList.iconArray)
	fmt.Println(len(masterIconList.iconArray))
	for masterIconList.iconArray[0].getIconMcm() > 0 {
		makeCombatOrder()
		fmt.Println(order)
		pickIconSource(order)
        createTargetList()
         assert(len(targetList.iconArray)>0, "No Targets") // опционально - потом если целей нет будет действие ожидания
        pickIconTarget(targetList)
		if iconSource.isPlayer == true {
			comm = userInput()
			//fmt.Println(command)
		} else {
			comm = formCommand(iconSource, iconTarget)	
			//outputRed(command)
		}
		outputRed(comm[0] + ">" + comm[1] + ">" + comm[2])
		//нужен утвердитель команды
		doMatrixAction(iconSource, iconTarget)
        checkPlay()
        masterIconList = destroyIcon(masterIconList)
		//fmt.Println(masterIconList.iconArray, "в конце хода")
        if len(masterIconList.iconArray) < 3 {
            masterIconList = addIcon(masterIconList, createIcon(6))
        }
	}

}

func checkPlay() {
    assert(masterIconList.iconArray[0].getIconMcm() > 0, "Connection Terminated. player destroyed" )

    
}

func pickIconTarget(targetList IconList) Icon {
    assert(targetList.isOk, "No targetList")
    fmt.Println("targetList=", targetList.iconArray)
    i := rand.Intn(len(targetList.iconArray))
    iconTarget = targetList.iconArray[i]
    //fmt.Println("iconTarget = ", iconTarget)
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
		//assert(icon.getIconInitiative() > 0, "Initiative Less than 0")
		//icon.setIconInitiative(icon.rollInitiative())
		order[i] = icon.getIconInitiative()
		masterIconList.iconArray[i] = icon

	}
	fmt.Println("Order before sorting:", order)
	bubbleSort(order)
	fmt.Println("Order after sorting:", order)
	if order[0] < 0 {
        turn++
        fmt.Println("Highest ini < 0. Rerolling:", order)
		fmt.Println("Start turn", turn)
        fmt.Println("************")
        fmt.Println("************")
        fmt.Println("************")
		allRollInitiative()
		fmt.Println("Order before sorting:", order)
		bubbleSort(order)
		fmt.Println("Order after sorting:", order)
	}
	return order
}

func pickIconSource([]int) Icon {
	var icon Icon
	for i := range masterIconList.iconArray {
		icon = masterIconList.iconArray[i]
		if icon.getIconInitiative() == order[0] {
			iconSource = icon
			fmt.Println("ходит икона ", iconSource.getIconID())
			return iconSource
		}
		fmt.Println(i, icon.getIconInitiative(), order[0])
	}
	fmt.Println("не смогли выбрать IconSource")
	fmt.Println(masterIconList.iconArray)
	return iconSource
}

func doMatrixAction(iconSource Icon, iconTarget Icon) { //должно быть еще название действия и механизмы выбора
    //строитель дайспула пойдет в отдельную функцию
    dicePoolSrc := iconSource.getIconDeviceRating() * 2
    dicePoolTrgt := iconTarget.getIconDeviceRating() * 2
    limit := iconSource.getIconDeviceRating()
    //
    netHits,_,_ := opposedTest(dicePoolSrc, dicePoolTrgt, limit)
    //распределение эффектов пойдет в отдельную функцию
    if netHits > 0 {
        iconTarget.setIconMcm(iconTarget.getIconMcm() - netHits)
       // fmt.Println("should hit")
       // fmt.Println(masterIconList.iconArray)
       // fmt.Println(iconTarget)
    }
    iconSource.setIconInitiative(iconSource.getIconInitiative() - 10)
	renewIconSource(iconSource)
    renewIconTarget(iconTarget)
    targetList.iconArray = nil //зачищаем список целей в конце действия
}

func destroyIcon(masterIconList IconList) IconList {
    //fmt.Println( "Destroy Icons:", masterIconList.iconArray)
    for i := range masterIconList.iconArray {
        if masterIconList.iconArray[i].getIconMcm() < 1 {
           //toDelete := masterIconList.iconArray[i].getIconID()
           result := []Icon{}
		   fmt.Println(masterIconList.iconArray[i].getIconName(), "destroyed")
           result = append(result, masterIconList.iconArray[0:i]...)
           result = append(result, masterIconList.iconArray[i+1:]...)
           masterIconList.iconArray = result
           break
           
        }
    }
    return masterIconList
}

func renewIconSource(iconSource Icon) {
    for i := range masterIconList.iconArray {
		if iconSource.getIconID() == masterIconList.iconArray[i].getIconID() {
			masterIconList.iconArray[i] = iconSource
            fmt.Println(masterIconList.iconArray)
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
	for i := range masterIconList.iconArray {
		var icon Icon
		icon = masterIconList.iconArray[i]
		icon.setIconInitiative(icon.rollInitiative())
		masterIconList.iconArray[i].setIconInitiative(icon.getIconInitiative())
		order[i] = masterIconList.iconArray[i].getIconInitiative()
	}
	return order
}
