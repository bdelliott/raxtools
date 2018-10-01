package auth

import (
	"fmt"
	"github.com/rackspace/gophercloud"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type CloudConfig struct {
	Clouds map[string]Cloud `yaml:"clouds"`
}

type Cloud struct {
	Auth Auth `yaml:"auth"`

	IdentityApiVersion int      `yaml:"identity_api_version"`
	Insecure           bool     `yaml:"insecure"`
	Profile            string   `yaml:"profile"`
	Regions            []string `yaml:"regions"`
}

type Auth struct {
	URL             string `yaml:"auth_url"`
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	ProjectID       string `yaml:"project_id"`
	UserDomainName  string `yaml:"user_domain_name"`
	ProjectName     string `yaml:"project_name"`
	ProjectDomainId string `yaml:"project_domain_id"`
}

const (
	RACKSPACE          = "rackspace"
	RACKSPACE_AUTH_URL = "https://identity.api.rackspacecloud.com/v2.0/"
)

// load cloud options from the cloud yaml
func FromCloudsYaml(cloudName string) gophercloud.AuthOptions {

	path := fmt.Sprintf("%s/.config/openstack/clouds.yaml", os.Getenv("HOME"))
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// AuthOptions
	var cloudConfig CloudConfig
	err = yaml.UnmarshalStrict(buf, &cloudConfig)
	if err != nil {
		panic(err)
	}

	cloud := cloudConfig.Clouds[cloudName]
	if cloud.Auth.URL == "" && cloud.Profile == RACKSPACE {
		cloud.Auth.URL = RACKSPACE_AUTH_URL
	}

	return gophercloud.AuthOptions{
		IdentityEndpoint: cloud.Auth.URL,
		Username:         cloud.Auth.Username,
		Password:         cloud.Auth.Password,
		TenantID:         cloud.Auth.ProjectID,
	}
}
