package shared_domain

type DatabaseSettings struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Settings struct {
	Port     string
	Database DatabaseSettings
}
