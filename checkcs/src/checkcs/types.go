package checkcs

type Service struct {
	Project string
	Application string
	Lane string
	CIID int
}

type Check struct {
	Code int
	Status string
	Message string
}
