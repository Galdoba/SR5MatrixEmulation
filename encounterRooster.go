package main

import (
	"fmt"
)

var order []int
var orderID []int
var masterIconList IconList
var iconSource Icon
var iconTarget Icon

type IconList struct {
	iconArray []Icon
	size      int
	isOk      bool
}

func createRooster() {
	fmt.Println("Start Creating Rooster")
	masterIconList = makeIconList()
	newIcon1 := createIcon(2)
	newIcon2 := createIcon(5)
	newIcon3 := createIcon(3)
	masterIconList = addIcon(masterIconList, newIcon1)
	masterIconList = addIcon(masterIconList, newIcon2)
	masterIconList = addIcon(masterIconList, newIcon3)
	fmt.Println("Add Icons")
	fmt.Println(masterIconList.iconArray)
	fmt.Println(len(masterIconList.iconArray))
	for masterIconList.iconArray[2].getIconMcm() > 0 {
		fmt.Println(masterIconList.iconArray, "при старте хода")
		makeCombatOrder()
		fmt.Println(order)
		pickIconSource(order)
		doMatrixAction(iconSource)
		fmt.Println(masterIconList.iconArray, "в конце хода")
	}

}

func makeCombatOrder() []int {
	var icon Icon
	size := getIconlistSize()
	order = make([]int, size)
	for i := range masterIconList.iconArray {
		//for i:= 0; i < size; i++ {
		icon = masterIconList.iconArray[i]
		//assert(icon.getIconInitiative() > 0, "Initiative Less than 0")
		//icon.setIconInitiative(icon.rollInitiative())
		order[i] = icon.getIconInitiative()
		//		fmt.Println("Icon #:", i, icon)
		//		fmt.Println("IconID #:", icon.getIconID())
		masterIconList.iconArray[i] = icon

	}
	fmt.Println("Order before sorting:", order)
	bubbleSort(order)
	//order[0], order[2] = order[2], order[0]
	fmt.Println("Order after sorting:", order)
	if order[0] < 0 {
		fmt.Println("Highest ini < 0. Rerolling:", order)
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
			fmt.Println("ходит икона выбор", iconSource.getIconID())
			return iconSource
		}
		fmt.Println(i, icon.getIconInitiative(), order[0])
	}
	fmt.Println("не смогли выбрать IconSource")
	fmt.Println(masterIconList.iconArray)
	return iconSource
}

func doMatrixAction(sourceIcon Icon) { //должно быть еще название действия и механизмы выбора
	iconSource.setIconMcm(iconSource.getIconMcm() - 1)
	iconSource.setIconInitiative(iconSource.getIconInitiative() - 10)
	//masterIconList.iconArray[0] = iconSource //Почему 0?
	fmt.Println(iconSource)
	//fmt.Println("Start IconSource to Array")
	//fmt.Println("изменения в аррэе", masterIconList.iconArray)
	for i := range masterIconList.iconArray {
		if iconSource.getIconID() == masterIconList.iconArray[i].getIconID() {
			masterIconList.iconArray[i] = iconSource
			//		fmt.Println("ходит икона лоарп", iconSource.getIconID())
			//		fmt.Println("изменения в аррэе", masterIconList.iconArray)
			//		fmt.Println("End IconSource to Array")
		}
		//fmt.Println(i, icon.getIconInitiative(), order[0])
	}
	//fmt.Println("no IconSource to Array")
}

func makeIconList() IconList {
	var mList IconList
	mList.iconArray = make([]Icon, 0, 1)
	mList.isOk = true
	//createPersona()
	mList = addIcon(mList, createPersona())
	return mList
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
