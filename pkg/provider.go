package pkg

import (
	"github.com/hashicorp/terraform/helper/schema"

	"github.com/piotrjaromin/terraform-provider-jenkins/pkg/resources/sampleserver"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"example_server": sampleserver.ResourceServer(),
		},
	}
}
