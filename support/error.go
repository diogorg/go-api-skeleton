package support

func Panic(err error) {
	if err != nil {
		panic(err.Error())
	}
}
