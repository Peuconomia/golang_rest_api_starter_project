package types

type EnvMode string

const (
	None        EnvMode = "none"
	Default             = "default"
	Staging             = "staging"
	Development         = "development"
	Production          = "production"
)

func (mode EnvMode) String() string {
	switch mode {
	case None:
		return "none"
	case Default:
		return "default"
	case Staging:
		return "staging"
	case Development:
		return "development"
	case Production:
		return "production"
	default:
		return "none"
	}
}

func EnvModeFromString(environment *string) EnvMode {

	switch *environment {
	case "none":
		return None
	case "default":
		return Default
	case "staging":
		return Staging
	case "development":
		return Development
	case "production":
		return Production
	default:
		return None
	}
}
