package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

var matrixActionList []string
var iconActionList []string
var personaActionList []string

func createMatrixActionList() []string {
    matrixActionList := make ([]string, 0)
    matrixActionList = append(matrixActionList, "HOLD" )
    matrixActionList = append(matrixActionList, "DATA_SPIKE" )
    matrixActionList = append(matrixActionList, "HACK" )
    fmt.Println(matrixActionList)
    return matrixActionList
}




func checkMarksQty (iconSource Icon, iconTarget Icon, comm1 string) (bool, string) {
    actionValid := false
    sourceID := iconSource.getIconID()
    targetID := iconTarget.getIconID()
    haveMarks := 0
    i := retriveMarkPosition(sourceID, targetID)
    mark := markList[i]
        if checkMarkExistiense(&mark) == true {
            haveMarks = retriveMarkQty (sourceID, targetID)         
        }
    switch comm1 {
        case "HOLD": if haveMarks >= 0 {
            actionValid = true
        }
        case "HACK": if haveMarks >= 1 {
            actionValid = true
        }
        case "DATA_SPIKE": if haveMarks >=1 {
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




func createIconActList () []string {
    iconActionList := make ([]string, 0)
    iconActionList = append(matrixActionList, "HOLD" )
    //iconActionList = append(matrixActionList, "DATA_SPIKE" )
    iconActionList = append(matrixActionList, "HACK" )
    return iconActionList
}

func createPersonaActList () []string {
    personaActionList := make ([]string, 0)
    personaActionList = append(matrixActionList, "HOLD" )
    personaActionList = append(matrixActionList, "DATA_SPIKE" )
    personaActionList = append(matrixActionList, "HACK" )
    return personaActionList
}
