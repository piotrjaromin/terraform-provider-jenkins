package pkg

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/piotrjaromin/terraform-provider-jenkins/pkg/resources/credentials"
	"github.com/piotrjaromin/terraform-provider-jenkins/pkg/resources/job"
	"github.com/piotrjaromin/terraform-provider-jenkins/pkg/resources/plugins"

	"github.com/piotrjaromin/gojenkins"
)

func Provider() *schema.Provider {

	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "http://localhost:8080",
				Description: "host address of jenkins instance",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "username which should be used to loginto jenkins instance",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "password which should be used to loginto jenkins instance",
			},
			"ca_cert": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "server self-signed certificate",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"jenkins_job_xml":             job.XmlJob(),
			"jenkins_plugin":              plugins.Plugin(),
			"jenkins_username_credential": credentials.Username(),
			"jenkins_ssh_credential":      credentials.SSH(),
			"jenkins_docker_credential":   credentials.Docker(),
		},
		ConfigureFunc: configureFunc,
	}
}

func configureFunc(rd *schema.ResourceData) (interface{}, error) {

	config := JenkinsConfig{
		URL:      rd.Get("url").(string),
		Username: rd.Get("username").(string),
		Password: rd.Get("password").(string),
		CaCert:   rd.Get("ca_cert").(string),
	}

	jenkins := gojenkins.CreateJenkins(nil, config.URL, config.Username, config.Password)

	// Provide CA certificate if server is using self-signed certificate
	if len(config.CaCert) > 0 {
		jenkins.Requester.CACert = []byte(config.CaCert)
	}

	_, err := jenkins.Init()
	return jenkins, err
}

type JenkinsConfig struct {
	URL      string
	Username string
	Password string
	CaCert   string
}
