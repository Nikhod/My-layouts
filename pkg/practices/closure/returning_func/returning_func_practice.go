package returning_func

func Count(num, step int) (func() int, func() int) {
	var counter = num

	next := func() int {
		counter += step
		return counter
	}

	current := func() int {
		return counter
	}
	return current, next
}

type Counter struct {
	Reset   func()
	Next    func() int
	Current func() int
}

func NewCounter(start, step int) Counter {
	var count = start

	current := func() int {
		return count
	}

	next := func() int {
		count += step
		return count
	}

	reboot := func() {
		count = 0
	}

	return Counter{
		Reset:   reboot,
		Next:    next,
		Current: current,
	}

}
