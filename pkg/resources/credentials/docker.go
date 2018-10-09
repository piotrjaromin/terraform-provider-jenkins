package credentials

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/piotrjaromin/gojenkins"
	"github.com/piotrjaromin/terraform-provider-jenkins/pkg/resources/credentials/util"
)

type dockerProvider struct{}

func Docker() *schema.Resource {

	manager := util.CreateCredsManager(dockerProvider{})

	return &schema.Resource{
		Create: manager.ResourceServerCreate,
		Read:   manager.ResourceServerRead,
		Update: manager.ResourceServerUpdate,
		Delete: manager.ResourceServerDelete,

		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"identifier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "_",
			},
			"jobpath": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "global",
			},
			"server_ca_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"client_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
			"client_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
	}
}

func (dockerProvider) Empty() interface{} {
	return gojenkins.DockerServerCredentials{}
}

func (dockerProvider) FromResourceData(d *schema.ResourceData) (interface{}, error) {

	return gojenkins.DockerServerCredentials{
		ID:                  d.Get("identifier").(string),
		Scope:               d.Get("scope").(string),
		Username:            d.Get("username").(string),
		Description:         d.Get("description").(string),
		ServerCaCertificate: d.Get("server_ca_certificate").(string),
		ClientCertificate:   d.Get("client_certificate").(string),
		ClientKey:           d.Get("client_key").(string),
	}, nil
}
