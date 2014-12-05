package parse_xml

func CheckError(e error) {
	if e != nil {
		panic(e)
	}
}
