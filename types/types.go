package types

type Config struct {
	ApiVersion     string      `yaml:"apiVersion"`
	Kind           string      `yaml:"kind"`
	Clusters       []Cluster   `yaml:"clusters"`
	Users          []User      `yaml:"users"`
	Preferences    interface{} `yaml:"preferences,omitempty"`
	Contexts       []Context   `yaml:"contexts"`
	CurrentContext string      `yaml:"current-context"`
}

type Cluster struct {
	Name    string `yaml:"name"`
	Cluster struct {
		Server                   string `yaml:"server"`
		CertificateAuthorityDate string `yaml:"certificate-authority-data"`
	} `yaml:"cluster"`
}
type User struct {
	Name string `yaml:"name"`
	User struct {
		AuthProvider struct {
			Name   string `yaml:"name"`
			Config struct {
				ClientId                    string `yaml:"client-id"`
				ClientSecret                string `yaml:"client-secret"`
				IdToken                     string `yaml:"id-token"`
				IdpCertificateAuthorityData string `yaml:"idp-certificate-authority-data"`
				IdpIssuerUrl                string `yaml:"idp-issuer-url"`
				RefreshToken                string `yaml:"refresh-token"`
				ExtraScopes                 string `yaml:"extra-scopes"`
			} `yaml:"config"`
		} `yaml:"auth-provider"`
	} `yaml:"user"`
}
type Preferences struct{}
type Context struct {
	Name    string     `yaml:"name"`
	Context SubContext `yaml:"context"`
}
type SubContext struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}
