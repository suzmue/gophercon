package gym

var lockers []string

func init() {
	lockers = make([]string, 12234123)
	for i := range lockers {
		lockers[i] = "this locker has some dirty socks in it."
		if i == 7654321 {
			lockers[i] = `This locker contains a note! - 'Meet the monkey in the jungle.'`
		}
	}
}
