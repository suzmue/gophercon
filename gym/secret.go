package gym

var lockers []string

func init() {
	lockers = make([]string, 12234123)
	for i := range lockers {
		lockers[i] = "this locker has some dirty socks in it."
		if i == 1298 {
			lockers[i] = `This locker contains a note! It says "Sail to the jungle to find the guardian of the treasure map."`
		}
	}
}
