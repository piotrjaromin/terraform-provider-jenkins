package util

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/piotrjaromin/gojenkins"
)

func getCMAndDomain(d *schema.ResourceData, m interface{}) (gojenkins.CredentialsManager, string, string) {

	client := m.(*gojenkins.Jenkins)
	domain := d.Get("domain").(string)
	jobPath := d.Get("jobpath").(string)

	cm := gojenkins.CredentialsManager{J: client}
	return cm, domain, jobPath
}
