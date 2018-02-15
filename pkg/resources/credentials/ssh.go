package credentials

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/piotrjaromin/gojenkins"
	"github.com/piotrjaromin/terraform-provider-jenkins/pkg/resources/credentials/util"
	"github.com/thoas/go-funk"
)

type sshProvider struct{}

const fileOnMasterType = "fileOnMaster"
const directValueType = "directValue"

var valueTypes = []string{fileOnMasterType, directValueType}

func SSH() *schema.Resource {

	manager := util.CreateCredsManager(sshProvider{})

	return &schema.Resource{
		Create: manager.ResourceServerCreate,
		Read:   manager.ResourceServerRead,
		Update: manager.ResourceServerUpdate,
		Delete: manager.ResourceServerDelete,

		Schema: map[string]*schema.Schema{
			"identifier": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"passphrase": &schema.Schema{
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
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"value_type": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateSSHType,
			},
		},
	}
}

func (sshProvider) Empty() interface{} {
	return gojenkins.SSHCredentials{}
}

func (sshProvider) FromResourceData(d *schema.ResourceData) (interface{}, error) {
	source, err := privateKeySource(d)

	return gojenkins.SSHCredentials{
		ID:               d.Get("identifier").(string),
		Scope:            d.Get("scope").(string),
		Username:         d.Get("username").(string),
		Passphrase:       d.Get("passphrase").(string),
		Description:      d.Get("description").(string),
		PrivateKeySource: source,
	}, err
}

func privateKeySource(d *schema.ResourceData) (interface{}, error) {
	valType := d.Get("value_type").(string)
	val := d.Get("value").(string)

	switch valType {
	case fileOnMasterType:
		return gojenkins.PrivateKeyFile{
			Value: val,
			Class: gojenkins.KeySourceOnMasterType,
		}, nil
	case directValueType:
		return gojenkins.PrivateKey{
			Value: val,
			Class: gojenkins.KeySourceDirectEntryType,
		}, nil
	}

	return struct{}{}, fmt.Errorf("Invalid value type provided: %s", valType)
}

func validateSSHType(value interface{}, arg string) ([]string, []error) {
	fmt.Println("ARGS is ", arg)

	valType := value.(string)
	if !funk.ContainsString(valueTypes, valType) {
		return []string{}, []error{fmt.Errorf("valueType must be one of %v", valueTypes)}
	}

	return []string{}, []error{}
}
