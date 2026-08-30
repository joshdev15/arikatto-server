package config

type ConfigurationReader interface {
	GetCors() bool
	GetPrivatePath() string
	GetPublicPath() string
}
type configurationFile struct {
	Server struct {
		Cors             bool `yaml:"cors"`
		CertificatesPath struct {
			Private string `yaml:"private"`
			Public  string `yaml:"public"`
		} `yaml:"certificates"`
	} `yaml:"server"`
}

func (cf *configurationFile) GetCors() bool {
	return cf.Server.Cors
}

func (cf *configurationFile) GetPrivatePath() string {
	return cf.Server.CertificatesPath.Private
}

func (cf *configurationFile) GetPublicPath() string {
	return cf.Server.CertificatesPath.Public
}
