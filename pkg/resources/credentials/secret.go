package credentials

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/piotrjaromin/gojenkins"
	"github.com/piotrjaromin/terraform-provider-jenkins/pkg/resources/credentials/util"
)

type secretProvider struct{}

func Secret() *schema.Resource {

	manager := util.CreateCredsManager(secretProvider{})

	return &schema.Resource{
		Create: manager.ResourceServerCreate,
		Read:   manager.ResourceServerRead,
		Update: manager.ResourceServerUpdate,
		Delete: manager.ResourceServerDelete,

		Schema: map[string]*schema.Schema{
			"secret": &schema.Schema{
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
		},
	}
}

func (secretProvider) Empty() interface{} {
	return gojenkins.UsernameCredentials{}
}

func (secretProvider) FromResourceData(d *schema.ResourceData) (interface{}, error) {

	return gojenkins.StringCredentials{
		ID:          d.Get("identifier").(string),
		Scope:       d.Get("scope").(string),
		Secret:      d.Get("secret").(string),
		Description: d.Get("description").(string),
	}, nil
}
