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
