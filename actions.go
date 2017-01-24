package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

var matrixActionList []string

func createMatrixActionList() []string {
    matrixActionList := make ([]string, 0)
    matrixActionList = append(matrixActionList, "HOLD" )
    matrixActionList = append(matrixActionList, "DATA_SPIKE" )
    matrixActionList = append(matrixActionList, "HACK" )
    fmt.Println(matrixActionList)
    return matrixActionList
}


