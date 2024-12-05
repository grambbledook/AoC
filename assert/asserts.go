package assert

func NoError(err error) {
	if err != nil {
		panic("No Error assertion failed: " + err.Error())
	}
}

func Assert(condition bool) {
	if !condition {
		panic("Assertion failed!")
	}
}
