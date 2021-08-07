package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/sergief/terraform-provider-synology/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func folderItem() *schema.Resource {
	return &schema.Resource{
		CreateContext: folderCreateItem,
		ReadContext:   folderReadItem,
		UpdateContext: folderUpdateItem,
		DeleteContext: folderDeleteItem,
		Schema: map[string]*schema.Schema{
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Path of the folder",
			},
		},
	}
}

func folderCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(client.SynologyClient)

	path := d.Get("path").(string)

	service := FolderItemService{synologyClient: client}
	err := service.Create(path)
	if err != nil {
		return diag.FromErr(err)
	}
	folderReadItem(ctx, d, m)
	return diags
}

func folderReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	path := d.Get("path").(string)

	d.Set("path", path)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func folderUpdateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return folderCreateItem(ctx, d, m)
}

func folderDeleteItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(client.SynologyClient)

	path := d.Get("path").(string)

	service := FolderItemService{synologyClient: client}

	err := service.Delete(path)
	if err != nil {
		return diag.FromErr(err)
	}
	return diags
}
