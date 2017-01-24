/*
How to Call Tests outside of "dice.go"
opposedTest()  =>   opposedTest(dicePoolSrc int, dicePoolTrgt int , limit int) (return netHits int, glitch bool, critGlitch bool)

simpleTest()   =>   simpleTest(dicePoolSrc int, limit int, threshold int) (return netHits int, glitch bool, critGlitch bool)

xd6Test()      =>   xd6Test(dicePoolSrc) (return summ int)

extendedTest() => extendedTest(dicePoolSrc int, limit int, threshold int) (return netHits int, glitch bool, critGlitch bool)
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DicePool struct {
	pool     []int
	isOk     bool
	isRolled bool
}

func makeDicePool(size int) *DicePool {
	var dp DicePool
	dp.pool = make([]int, size, size*2)
	dp.isOk = true
	return &dp
}

func setSeed() {
	rand.Seed(time.Now().UnixNano())  //получаем рандомный сид от текущего времени с точностью до наносекунд
	time.Sleep(time.Millisecond * 30) //ждем 3 милисекунды чтобы сид гарантированно сменился к следующему заходу
}

func (dp *DicePool) roll() {
	assert(dp.isOk, "DicePool not initialized")
	setSeed()
	for i := range dp.pool {
		dp.pool[i] = rand.Intn(6) + 1
	}
	fmt.Println(dp.pool)
	dp.isRolled = true
}

func (dp *DicePool) successes() int {
	assert(dp.isRolled, "DicePool not Rolled")
	successes := 0
	for i := range dp.pool {
		if dp.pool[i] == 5 || dp.pool[i] == 6 {
			successes++
		}
	}
	//fmt.Println(successes)
	return successes
}

func (dp *DicePool) glitch() bool {
	assert(dp.isRolled, "DicePool not Rolled")
	glitch := false
	ones := 0
	for i := range dp.pool {
		if dp.pool[i] == 1 {
			ones++
		}
	}
	//fmt.Println(ones)
	if ones > len(dp.pool)/2 {
		glitch = true
	}
	return glitch
}

func (dp *DicePool) critGlitch() bool {
	assert(dp.isRolled, "DicePool not Rolled")
	critGlitch := false
	if dp.successes() == 0 && dp.glitch() {
		critGlitch = true
	}
	return critGlitch
}

func (dp *DicePool) summ() int {
	assert(dp.isRolled, "DicePool not Rolled")
	xd6 := 0
	total := 0
	for i := range dp.pool {
		xd6 = dp.pool[i]
		total = total + xd6
	}
	//fmt.Println(xd6)
	return total
}

func simpleTest(dicePool1 int, limit int, threshold int) (int, bool, bool) { 
	sourceIcon := makeDicePool(dicePool1)
	sourceIcon.roll()
	fmt.Println("Sucesesses =", sourceIcon.successes())
	netHits := (sourceIcon.successes() - threshold)
	glitch := sourceIcon.glitch()
	critGlitch := sourceIcon.critGlitch()
	fmt.Println("Nethits =", netHits, "Glitch =", glitch, "Critical Glitch =", critGlitch)
	return netHits, glitch, critGlitch
}

func opposedTest(dicePool1 int, dicePool2 int, limit int) (int, bool, bool) {
	suc1 := 0
	suc2 := 0
	sourceIcon := makeDicePool(dicePool1)
	sourceIcon.roll()
	targetIcon := makeDicePool(dicePool2)
	targetIcon.roll()
	suc1 = sourceIcon.successes()
	suc2 = targetIcon.successes()
	if suc1 > limit {
		fmt.Println("Succeses by Limit:", limit)
		suc1 = limit
	}
	fmt.Println("Source sucesesses =", suc1, "Target successes =", suc2)
	netHits := suc1 - suc2
	glitch := sourceIcon.glitch()
	critGlitch := sourceIcon.critGlitch()
	//fmt.Println("Nethits =", netHits, "Glitch =", glitch, "Critical Glitch =", critGlitch)
	return netHits, glitch, critGlitch
}

func xd6Test(dicePool1 int) (int) {
	sourceIcon := makeDicePool(dicePool1)
	sourceIcon.roll()
	summ := 0
	summ = sourceIcon.summ()
	return summ
}

func extendedTest(dicePool1 int, limit int, threshold int) (int, bool, bool) { 
	netHits := 0
	glitch := false
	critGlitch := false
	step := 1
	i := 0
	for i = dicePool1; i > 0; i-- {
		fmt.Println("Step =", step)
		sourceIcon := makeDicePool(i)
		sourceIcon.roll()	
		netHits = netHits + sourceIcon.successes()
		if sourceIcon.successes() == 0 {
			i++
		}
		if sourceIcon.glitch() == true {
			glitch = true
			fmt.Println("glitch")
			threshold = threshold + 2
		}
		if sourceIcon.critGlitch() == true {
			critGlitch = true
			fmt.Println("critGlitch")
			netHits = 0 // for making sure that test fail
			break
		}
		fmt.Println("Sucesesses =", sourceIcon.successes())
		if netHits >= threshold {
			break
		}
		step++
		fmt.Println("Nethits =", netHits, "Threshold =", threshold)
	}
	fmt.Println("Nethits =", netHits, "Glitch =", glitch, "Critical Glitch =", critGlitch)
	return netHits, glitch, critGlitch
}