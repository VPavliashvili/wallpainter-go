package opts

type Opt struct {
	Name  string
	Value string
}

type OptParser interface {
    Parse([]string) ([]Opt, error)
}
