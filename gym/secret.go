package gym

import "fmt"

// lockers is a slice that holds the contents of each locker
var lockers []string

// init initializes the lockers with default values and a special note in a specific locker
func init() {
	// Define the number of lockers
	const numLockers = 12234123

	// Initialize the lockers slice with a length of numLockers
	lockers = make([]string, numLockers)

	// Populate each locker with a default message
	for i := range lockers {
		lockers[i] = "This locker has some dirty socks in it."

		// Set a special note in locker 7654321
		if i == 7654321 {
			lockers[i] = `This locker contains a note! - 'Meet the monkey in the jungle.'`
		}
	}
}

// GetLockerContents returns the contents of a specific locker
func GetLockerContents(index int) (string, error) {
	if index < 0 || index >= len(lockers) {
		return "", fmt.Errorf("invalid locker index: %d", index)
	}
	return lockers[index], nil
}
