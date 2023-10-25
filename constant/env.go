package env

type Env string

const (
	Local       Env = "local"
	Development Env = "development"
	Production  Env = "production"
)

func (e Env) String() string {
	return string(e)
}
