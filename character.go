package diaboro

type Character struct {
	Name  string
	Class string

	Builds  map[string]Build
	Current struct {
		Name string
		Build
	}
}
