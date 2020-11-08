package build

type Build struct {
	Name string
	ID   string
}

func Status() string {
	return ""
}

func Log(offset int) (output string, length int) {
	return
}
