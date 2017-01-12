package main

import "fmt"

func main() {
	fmt.Println("DiceRoller")
	//Отправить в дайс 4 числа: Дайспулл 1, Дайспул 2, Лимит, Трешхолд

	for i := 1; i <= 0; i++ {
		pl1 := makeDicePool(5) // ожидаемоя запись: makeDicePool(getAtr() + getSkill() + getMod())
		fmt.Println("Roll №", i, ":")
		pl1.roll()
		fmt.Println("We got", pl1.successes(), "successes")
		fmt.Println("GLITCH is", pl1.glitch())
		fmt.Println("CriticalGlitch is", pl1.critGlitch())
		fmt.Println("Roll SUMM is", pl1.summ())
		if pl1.critGlitch() == true {
		}
	}

	/*
		SimpleTest: 9(3) [4]
		DicePool1 = 9
		threshold = 3
		Limit     = 4

	*/
	threshold := 40
	
	dicePoolSrc  := 11
	dicePoolTrgt := 9
	limit := 4
//	opposedTest(dicePool1, dicePool2, limit) 
	fmt.Println(opposedTest(dicePoolSrc, dicePoolTrgt, limit))
	fmt.Println("******************************************")
	fmt.Println(simpleTest(dicePoolSrc, limit, threshold))
	fmt.Println("******************************************")
	fmt.Println(xd6Test(dicePoolSrc))
	fmt.Println("******************************************")
	fmt.Println("******************************************")
	fmt.Println(extendedTest(dicePoolSrc, limit, threshold))
}
