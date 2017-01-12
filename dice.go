package main

import (
	"fmt"
	"math/rand"
	"time"
	//"strconv"
)

var a int
var resultArray []int

//func main() {
//roll()
//}

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
	assert(dp.isOk, "DicePool not initiated")
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

func simpleTest(threshold int) (int, bool, bool) { //Первым вхождением должна быть икона которая будет тест, вторым трешхолд (готово)
	// Конечная запись будет что-то вроде: simpleTest(src *Icon, threshold int) (int, bool, bool) { ... }
	pl1 := makeDicePool(9)
	pl1.roll()
	fmt.Println("Sucesesses =", pl1.successes())
	netHits := (pl1.successes() - threshold)
	glitch := pl1.glitch()
	critGlitch := pl1.critGlitch()
	fmt.Println("Nethits =", netHits, "Glitch =", glitch, "Critical Glitch =", critGlitch)
	return netHits, glitch, critGlitch

}

func opposedTest(threshold int) (int, bool, bool) { //ой ли???

	pl1 := makeDicePool(9)
	pl1.roll()
	fmt.Println("Sucesesses =", pl1.successes())
	netHits := (pl1.successes() - threshold)
	glitch := pl1.glitch()
	critGlitch := pl1.critGlitch()
	fmt.Println("Nethits =", netHits, "Glitch =", glitch, "Critical Glitch =", critGlitch)
	return netHits, glitch, critGlitch

}

/*func rollSuccessTest(dicePool1 int, limit int, threshold int) (int, bool, bool) {
	var source DicePool
	//var target DicePool
	source.init(dicePool1) //инициирую пул тут - хотя похорошему это надо делать в тест билдере
	//target.init(6)
	source.roll()

	successes := 0
	ones := 0
	glicth := false
	critGlitch := false
	xd6 := 0
	for _, v := range source.pool {
		if v == 5 || v == 6 {
			successes++
		} else if v == 1 {
			ones++
		}
		xd6 = xd6 + v
	}
	fmt.Println("Total succeses =", successes)
	fmt.Println("Total ones =", ones)
	if ones > dicePool1/2 {
		glicth = true
		fmt.Println("GLITCH", glicth)
		if successes == 0 {
			critGlitch = true
			fmt.Println("Critical Glitch", critGlitch)
		}
	}
	fmt.Println("Total dots =", xd6)
	//return resultArray[5]
	//выясняем удовлетворены ли мы результатом, если нет то ReRoll
	// условно да

	//for i == 1
	return successes, glicth, critGlitch
}*/
