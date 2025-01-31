package helper

func PanicfIfError(err error) {
	if err != nil {
		panic(err)
	}
}
