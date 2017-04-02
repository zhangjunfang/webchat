package err

func CheckError(err error, message string) {
	if err != nil {
		panic(err)
	}
}
