package util

import (
	"github.com/hashicorp/terraform/helper/schema"
	"reflect"
)

type CredsProvider interface {
	FromResourceData(*schema.ResourceData) (interface{}, error)
	Empty() interface{}
}

type CredsManager struct {
	ResourceServerCreate func(d *schema.ResourceData, m interface{}) error
	ResourceServerRead   func(d *schema.ResourceData, m interface{}) error
	ResourceServerDelete func(d *schema.ResourceData, m interface{}) error
	ResourceServerUpdate func(d *schema.ResourceData, m interface{}) error
}

func CreateCredsManager(provider CredsProvider) CredsManager {

	id := func(cred interface{}) string {
		valStruct := reflect.ValueOf(cred).FieldByName("ID")
		return valStruct.String()
	}

	create := func(d *schema.ResourceData, m interface{}) error {

		cm, domain, jobPath := getCMAndDomain(d, m)
		cred, err := provider.FromResourceData(d)
		if err != nil {
			return err
		}

		err = cm.Add(domain, jobPath, cred)
		if err != nil {
			return err
		}

		d.SetId(id(cred))
		d.Set("cred", cred)
		return nil
	}

	read := func(d *schema.ResourceData, m interface{}) error {

		cm, domain, jobPath := getCMAndDomain(d, m)

		cred := provider.Empty()
		err := cm.GetSingle(domain, jobPath, d.Id(), &cred)
		if err != nil {
			return err
		}

		d.Set("cred", cred)
		return nil
	}

	update := func(d *schema.ResourceData, m interface{}) error {
		cm, domain, jobPath := getCMAndDomain(d, m)
		cred, err := provider.FromResourceData(d)
		if err != nil {
			return err
		}

		return cm.Update(domain, jobPath, id(cred), cred)
	}

	delete := func(d *schema.ResourceData, m interface{}) error {
		cm, domain, jobPath := getCMAndDomain(d, m)
		return cm.Delete(domain, jobPath, d.Id())
	}

	return CredsManager{
		ResourceServerCreate: create,
		ResourceServerRead:   read,
		ResourceServerDelete: delete,
		ResourceServerUpdate: update,
	}
}
