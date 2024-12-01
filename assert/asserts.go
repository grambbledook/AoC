package assert

func NoError(err error) {
	if err != nil {
		panic("No Error assertion failed: " + err.Error())
	}
}
