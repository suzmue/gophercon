package bar

func startOfBar() {
	GatherIntel()
}

func name1() {
	hi := speak["name1"]
	_ = hi
	startOfBar()
}

func name2() {
	hi := speak["name2"]
	_ = hi
	name1()
}

func name3() {
	hi := speak["name3"]
	_ = hi
	name2()
}

func name4() {
	hi := speak["name4"]
	_ = hi
	name3()
}

func name5() {
	hi := speak["name5"]
	_ = hi
	name4()
}

func endOfBar() {
	name5()
}

func enterBar() {
	endOfBar()
}
