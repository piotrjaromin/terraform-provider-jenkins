package credentials

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/piotrjaromin/gojenkins"
	"github.com/piotrjaromin/terraform-provider-jenkins/pkg/resources/credentials/util"
)

type usernameProvider struct{}

func Username() *schema.Resource {

	manager := util.CreateCredsManager(usernameProvider{})

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
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "_",
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

func (usernameProvider) Empty() interface{} {
	return gojenkins.UsernameCredentials{}
}

func (usernameProvider) FromResourceData(d *schema.ResourceData) (interface{}, error) {

	return gojenkins.UsernameCredentials{
		ID:          d.Get("identifier").(string),
		Scope:       d.Get("scope").(string),
		Username:    d.Get("username").(string),
		Password:    d.Get("password").(string),
		Description: d.Get("description").(string),
	}, nil
}
