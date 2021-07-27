package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/sergief/terraform-provider-synology/client"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func fileItem() *schema.Resource {
	return &schema.Resource{
		CreateContext: fileCreateItem,
		ReadContext:   fileReadItem,
		UpdateContext: fileUpdateItem,
		DeleteContext: fileDeleteItem,
		Schema: map[string]*schema.Schema{
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Contents of the file",
			},
			"filename": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The filename including path",
			},
		},
	}
}

func fileCreateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(client.SynologyClient)

	content := d.Get("content").(string)
	filename := d.Get("filename").(string)

	service := FileItemService{synologyClient: client}
	service.Create(filename, []byte(content))

	fileReadItem(ctx, d, m)

	return diags
}

func fileReadItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(client.SynologyClient)

	filename := d.Get("filename").(string)

	service := FileItemService{synologyClient: client}

	content := service.Read(filename)

	d.Set("filename", filename)
	d.Set("content", string(content))
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func fileUpdateItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return fileCreateItem(ctx, d, m)
}

func fileDeleteItem(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	client := m.(client.SynologyClient)

	filename := d.Get("filename").(string)

	service := FileItemService{synologyClient: client}

	service.Delete(filename)

	return diags
}
