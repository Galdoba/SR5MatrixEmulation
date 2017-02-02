package main

import (
	"fmt"
	"math/rand"
	"strings"
	//"math/rand"
	//"time"
)

var matrixActionList []string
var iconActionList []string
var personaActionList []string

func createMatrixActionList() []string {
	matrixActionList := make([]string, 0)
	matrixActionList = append(matrixActionList, "HOLD")
	matrixActionList = append(matrixActionList, "DATA_SPIKE")
	matrixActionList = append(matrixActionList, "HACK")
	matrixActionList = append(matrixActionList, "BRUTE_FORCE")

	fmt.Println(matrixActionList)
	return matrixActionList
}

func checkMarksQty(iconSource Icon, iconTarget Icon, comm1 string) (bool, string) {
	actionValid := false
	sourceID := iconSource.getIconID()
	targetID := iconTarget.getIconID()
	haveMarks := 0
	i := retriveMarkPosition(sourceID, targetID)
	mark := markList[i]
	if checkMarkExistiense(&mark) == true {
		haveMarks = retriveMarkQty(sourceID, targetID)
	}
	switch comm1 {
	case "HOLD":
		if haveMarks >= 0 {
			actionValid = true
		}
	case "HACK":
		if haveMarks >= 1 {
			actionValid = true
		}
	case "DATA_SPIKE":
		if haveMarks >= 0 {
			actionValid = true
		}
	case "BRUTE_FORCE":
		if haveMarks >= 0 {
			actionValid = true
		}
	}
	reason := "a"
	if actionValid == false {
		reason = "not enough marks for this command..."
	} else {
		reason = "All Good"
	}

	return actionValid, reason
}

func createIconActList() []string {
	iconActionList := make([]string, 0)
	iconActionList = append(matrixActionList, "HOLD")
	iconActionList = append(matrixActionList, "DATA_SPIKE")
	iconActionList = append(matrixActionList, "HACK")
	return iconActionList
}

func createPersonaActList() []string {
	personaActionList := make([]string, 0)
	personaActionList = append(matrixActionList, "HOLD")
	personaActionList = append(matrixActionList, "DATA_SPIKE")
	personaActionList = append(matrixActionList, "HACK")
	personaActionList = append(matrixActionList, "BRUTE_FORCE")
	return personaActionList
}

func chooseMatrixAction(iconSource Icon, iconTarget Icon, comm []string) (string, bool) {
	actionValid = false
	sourceType := iconSource.getIconType()
	setSeed()
	for actionValid == false {
		switch sourceType {
		case "Icon":
			actionName = matrixActionList[rand.Intn(len(matrixActionList))]
			switch actionName {
			case "HOLD": //делаем проверку валидности исходя из видимости наличия марок и тд
				actionValid = true
				fmt.Println(actionName, "Всегда валидно")
			case "DATA_SPIKE":
				fmt.Println(actionName, "Всегда не валидно")
			case "HACK":
				fmt.Println(actionName, "Всегда не валидно")
			}
		case "Persona":
			actionName = personaActionList[rand.Intn(len(personaActionList))]
			if iconSource.isPlayer == true {
				actionName = comm[1]
				actionName = strings.ToUpper(actionName)
			}
			switch actionName {
			case "HOLD": //делаем проверку валидности исходя из видимости наличия марок и тд
				actionValid = true
				fmt.Println(actionName, "Всегда валидно")
			case "DATA_SPIKE":
				actionValid = true
				fmt.Println(actionName, "Всегда валидно")
			case "HACK":
				actionValid = true
				fmt.Println(actionName, "Всегда валидно")
			case "BRUTE_FORCE":
				actionValid = true
				fmt.Println(actionName, "Всегда валидно")
			}
		}
		if iconSource.isPlayer == true {
			actionValid, _ = checkMarksQty(iconSource, iconTarget, actionName)
			return actionName, actionValid
		}
	}
	return actionName, actionValid
}

func actionEffect(actionName *string, iconSource *Icon, iconTarget *Icon, netHits *int) {

	switch *actionName {
	case "DATA_SPIKE":
		applyDataSpike(iconSource, iconTarget, *netHits)
	case "BRUTE_FORCE":
		applyBruteForce(iconSource, iconTarget, *netHits)
	case "HOLD":
		//applyHold()
	}

}

func applyDataSpike(iconSource *Icon, iconTarget *Icon, netHits int) {
	if netHits > 0 {
		damage := netHits + iconSource.getIconDeviceRating()
		damage = damage + (retriveMarkQty(iconSource.getIconID(), iconTarget.getIconID()) * 2)
		damageSoak, _, _ := simpleTest((iconTarget.getIconDeviceRating() * 2), 0, 0) //забираем только успехи (пока)
		fmt.Println("damage is", damage)
		damage = damage - damageSoak
		if damage < 0 {
			damage = 0
		}
		iconTarget.setIconMcm(iconTarget.getIconMcm() - damage)
		if iconSource.isPlayer == true {
			outputRed("action success!!")
		}
	} else {
		fmt.Println(netHits)
		outputRed("Attack gets " + string(netHits) + "damage")
		iconSource.setIconMcm(iconSource.getIconMcm() + netHits)
		if iconSource.isPlayer == true {
			outputRed("action failed!!")
		}
	}
}

func applyBruteForce(iconSource *Icon, iconTarget *Icon, netHits int) {
	if netHits > 0 {
		damage := netHits / 2
		iconTarget.setIconMcm(iconTarget.getIconMcm() - damage)
		placeMark(iconSource.getIconID(), iconTarget.getIconID())
	}
}
