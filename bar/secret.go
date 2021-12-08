package bar

func startOfBar() {
	GatherIntel()
}

func Alice() {
	hi := speak["Alice"]
	_ = hi
	startOfBar()
}

func Bob() {
	hi := speak["Bob"]
	_ = hi
	Alice()
}

func Carol() {
	hi := speak["Carol"]
	_ = hi
	Bob()
}

func Dave() {
	hi := speak["Dave"]
	_ = hi
	Carol()
}

func Eve() {
	hi := speak["Eve"]
	_ = hi
	Dave()
}

func endOfBar() {
	Eve()
}

func Bar() {
	endOfBar()
}
