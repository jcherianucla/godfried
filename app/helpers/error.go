package helpers

func CheckError(err error) {
	if err != nil {
		Logger.Error(err)
		panic(err)
	}
}
