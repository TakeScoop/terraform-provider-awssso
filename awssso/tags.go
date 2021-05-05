package aws

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func tagsSchemaComputed() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeMap,
		Optional: true,
		Computed: true,
		Elem:     &schema.Schema{Type: schema.TypeString},
	}
}
