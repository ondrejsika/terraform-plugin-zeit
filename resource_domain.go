package main

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/ondrejsika/vercel-go"
)

func resourceDomainCreate(d *schema.ResourceData, m interface{}) error {
	domain := d.Get("domain").(string)
	expectedPrice := d.Get("expected_price").(int)
	removeDomainOnDestroy := d.Get("remove_domain_on_destroy").(bool)

	d.SetId(domain)
	d.Set("remove_domain_on_destroy", removeDomainOnDestroy)

	rawResp, err := vercel.NewOrigin(m.(*Config).Token, m.(*Config).ApiOrigin).RawBuyDomain(domain, expectedPrice)

	if err != nil {
		return fmt.Errorf("%s", err)
	}

	var response map[string]map[string]interface{}
	json.Unmarshal([]byte(rawResp.Body()), &response)

	if response["error"]["code"] == "not_available" {
		return fmt.Errorf("Domain %s is not available", domain)
	}

	return nil
}

func resourceDomainRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDomainUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceDomainDelete(d *schema.ResourceData, m interface{}) error {
	domain := d.Get("domain").(string)
	removeDomainOnDestroy := d.Get("remove_domain_on_destroy").(bool)

	if removeDomainOnDestroy {
		_, err := vercel.NewOrigin(m.(*Config).Token, m.(*Config).ApiOrigin).RawRemoveDomain(domain)

		if err != nil {
			return fmt.Errorf("%s", err)
		}
	}

	return nil
}

func resourceDomainImport(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	domain := d.Id()
	d.Set("domain", domain)
	d.Set("remove_domain_on_destroy", false)

	resp, err := vercel.NewOrigin(m.(*Config).Token, m.(*Config).ApiOrigin).GetDomainPrice(domain)

	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	d.Set("expected_price", resp.Price)
	return []*schema.ResourceData{d}, nil
}

func resourceDoamin() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomainCreate,
		Read:   resourceDomainRead,
		Update: resourceDomainUpdate,
		Delete: resourceDomainDelete,
		Importer: &schema.ResourceImporter{
			State: resourceDomainImport,
		},

		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"expected_price": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"remove_domain_on_destroy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
