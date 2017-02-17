package main

import (
	//"fmt"
	//"strconv"
	"fmt"
	"os"
)

var markList []Mark

type Mark struct {
	sourceID int
	targetID int
	markQty  int
}

func createMark(sourceID int, targetID int) *Mark {
	var newMark Mark
	newMark.sourceID = sourceID
	newMark.targetID = targetID
	newMark.markQty = 0
	if newMark.sourceID == newMark.targetID {
		newMark.markQty = 4
	}
	//  fmt.Println("Создали МАРКу")
	return &newMark
}

func createMarkList() []Mark {
	markList := make([]Mark, 0, 1)
	return markList
}

func updateMarks() []Mark {
	for i := range masterIconList.iconArray {
		sourceID := masterIconList.iconArray[i].getIconID()
		//mark.setMarkSourceID(masterIconList.iconArray[i].getIconID())
		for j := range masterIconList.iconArray {
			targetID := masterIconList.iconArray[j].getIconID()
			// fmt.Println("Проверяем можно ли создать марку", i, j)
			//mark.setMarkTargetID(masterIconList.iconArray[j].getIconID())
			mark := createMark(sourceID, targetID)
			//fmt.Println("создали марку", i, j, mark)
			if checkMarkExistiense(mark) == false {
				//  fmt.Println("Такой марки НЕТ - добавляем", i, j, mark)
				markList = addMark(markList, mark)
				//fmt.Println("Создали марку", mark, i, j)
			}
		}
	}
	//fmt.Println(markList)
	return markList
}

func checkMarkExistiense(mark *Mark) bool {
	markExists := false
	for i := range markList {
		//fmt.Println("Размер списка =", len(markList))
		//fmt.Println("Проверяем список котовых марок.(позиция в списке/координаты)", i,markList[i].getMarkSourceID(), "=",mark.getMarkSourceID(), markList[i].getMarkTargetID(), "=", mark.getMarkTargetID()  )
		if markList[i].getMarkSourceID() == mark.getMarkSourceID() && markList[i].getMarkTargetID() == mark.getMarkTargetID() {
			markExists = true
		}
	}
	return markExists
}

func addMark(markList []Mark, mark *Mark) []Mark {
	markList = append(markList, *mark)
	return markList
}

func clearMarks(id int) {
	for i := range markList {
		if markList[i].getMarkSourceID() == id || markList[i].getMarkTargetID() == id {
			result := []Mark{}
			result = append(result, markList[0:i]...)
			result = append(result, markList[i+1:]...)
			markList = result
			clearMarks(id)
			break
		}
	}
}

func (mark *Mark) setMarkSourceID(sourceID int) {
	mark.sourceID = sourceID
}

func (mark *Mark) getMarkSourceID() int {
	return mark.sourceID
}

func (mark *Mark) setMarkTargetID(targetID int) {
	mark.targetID = targetID
}

func (mark *Mark) getMarkTargetID() int {
	return mark.targetID
}

func (mark *Mark) setMarkQty(markQty int) {
	mark.markQty = markQty
}

func (mark *Mark) getMarkQty() int {
	return mark.markQty
}

func retriveMarkQty(sourceID int, targetID int) int {
	markQty := -1
	for i := range markList {
		if markList[i].getMarkSourceID() == sourceID && markList[i].getMarkTargetID() == targetID {
			markQty = markList[i].getMarkQty()
		}
	}
	if markQty == -1 {
		fmt.Println("минус одна марка!!!")
		fmt.Println("этого не должно быть потому что не должно быть никогда")
		os.Exit(1)
	}
	return markQty
}

func retriveMarkPosition(sourceID int, targetID int) int {
	for i := range markList {
		if markList[i].getMarkSourceID() == sourceID && markList[i].getMarkTargetID() == targetID {
			return i
		}
	}
	fmt.Println("марка не найдена!!!")
	fmt.Println("этого не должно быть потому что не должно быть никогда")
	os.Exit(1)
	return 999
}

func placeMark(sourceID int, targetID int) {
	for i := range markList {
		if markList[i].getMarkSourceID() == sourceID && markList[i].getMarkTargetID() == targetID {
			markList[i].setMarkQty(markList[i].getMarkQty() + 1)
			if markList[i].getMarkQty() > 3 {
				markList[i].setMarkQty(3)
			}
		}
	}
}

func eraseMark(sourceID int, targetID int) {
	for i := range markList {
		if markList[i].getMarkSourceID() == sourceID && markList[i].getMarkTargetID() == targetID {
			markList[i].setMarkQty(markList[i].getMarkQty() - 1)
			if markList[i].getMarkQty() < 0 {
				markList[i].setMarkQty(0)
			}
		}
	}
}
