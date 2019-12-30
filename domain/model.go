package domain

type Variant struct {
	Name       string
	Percentage float64
}

type Experiment struct {
	Name string
	Variants []Variant
}

type User struct {
	Id string
}
