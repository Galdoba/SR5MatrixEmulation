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
	iconArray    []Icon
    size         int
    isOk         bool
}

func createRooster() {
    fmt.Println("Start Creating Rooster")
    masterIconList = makeIconList()
    newIcon1 := createIcon(2)
    newIcon2 := createIcon(2)
    newIcon3 := createIcon(3)
    masterIconList = addIcon(masterIconList, newIcon1)
    masterIconList = addIcon(masterIconList, newIcon2)
    masterIconList = addIcon(masterIconList, newIcon3)
    fmt.Println("Add Icons")
    fmt.Println(masterIconList.iconArray)
    fmt.Println(len(masterIconList.iconArray))
    for masterIconList.iconArray[0].getIconMcm() > 0 {
    makeCombatOrder()
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
        icon.setIconInitiative(icon.rollInitiative())
        order[i] = icon.getIconInitiative()
        fmt.Println("Icon #:", i, icon)
        fmt.Println("IconID #:", icon.getIconID())
        masterIconList.iconArray[i] = icon
        
    }
    fmt.Println("Order before sorting:", order)
    bubbleSort(order)
	//order[0], order[2] = order[2], order[0]
    fmt.Println("Order after sorting:", order)

    for i := range masterIconList.iconArray {
        icon = masterIconList.iconArray[i]
        if icon.getIconInitiative() == order[0] {
            iconSource = icon
            fmt.Println("ходит икона", iconSource.getIconID())
            
            iconSource.setIconMcm(iconSource.getIconMcm() - 1)
            iconSource.setIconInitiative(iconSource.getIconInitiative() - 10)
            masterIconList.iconArray[i] = iconSource
        }
             fmt.Println(i, icon.getIconInitiative(), order[0])
    }


fmt.Println(masterIconList.iconArray)
    return order
}

func makeIconList() IconList {
	var mList IconList
    mList.iconArray = make([]Icon, 0, 1)
	mList.isOk = true
    //createPersona()
    mList = addIcon(mList, createPersona())
    return mList
}

func addIcon (masterIconList IconList, newIcon Icon) IconList {
    assert(masterIconList.isOk, "No Icon List")
    masterIconList.iconArray = append(masterIconList.iconArray, newIcon) 
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