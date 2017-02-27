package main

var relations []Info

type Info struct {
	sourceID int
	targetID int
	spotted bool
}

func createInfo(sourceID int, targetID int) *Info {
	var newInfo Info
	newInfo.sourceID = sourceID
	newInfo.targetID = targetID
	
	//  fmt.Println("Создали МАРКу")
	return &newInfo
}

func createRelations() []Info {
	relations := make([]Info, 0, 1)
	return relations
}

func updateInfo() []Info {
	for i := range masterIconList.iconArray {
		sourceID := masterIconList.iconArray[i].getIconID()
		//mark.setMarkSourceID(masterIconList.iconArray[i].getIconID())
		for j := range masterIconList.iconArray {
			targetID := masterIconList.iconArray[j].getIconID()
			// fmt.Println("Проверяем можно ли создать марку", i, j)
			//mark.setMarkTargetID(masterIconList.iconArray[j].getIconID())
			info := createInfo(sourceID, targetID)
			//fmt.Println("создали марку", i, j, mark)
			if checkInfoExistiense(info) == false {
				//  fmt.Println("Такой марки НЕТ - добавляем", i, j, mark)
				relations = addInfo(relations, info)
				//fmt.Println("Создали марку", mark, i, j)
			}
		}
	}
	//fmt.Println(markList)
	return relations
}

func checkInfoExistiense(info *Info) bool {
	markExists := false
	for i := range relations {
		//fmt.Println("Размер списка =", len(markList))
		//fmt.Println("Проверяем список котовых марок.(позиция в списке/координаты)", i,markList[i].getMarkSourceID(), "=",mark.getMarkSourceID(), markList[i].getMarkTargetID(), "=", mark.getMarkTargetID()  )
		if relations[i].getInfoSourceID() == info.getInfoSourceID() && relations[i].getInfoTargetID() == info.getInfoTargetID() {
			markExists = true
		}
	}
	return markExists
}

func addInfo(relations []Info, info *Info) []Info {
	relations = append(relations, *info)
	return relations
}

func clearInfo(id int) {
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

func (info *Info) setInfoSourceID(sourceID int) {
	info.sourceID = sourceID
}

func (info *Info) getInfoSourceID() int {
	return info.sourceID
}

func (info *Info) setInfoTargetID(targetID int) {
	info.targetID = targetID
}

func (info *Info) getInfoTargetID() int {
	return info.targetID
}