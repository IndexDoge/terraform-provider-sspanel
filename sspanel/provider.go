package sspanel

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
	"time"
)

func init() {
	schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			DataSourcesMap: map[string]*schema.Resource{
				"sspanel_nodes": dataSourceNodes(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"sspanel_node": resourceNode(),
			},
			Schema: map[string]*schema.Schema{
				"url": {
					Type:        schema.TypeString,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("SSPANEL_BASE_URL", nil),
				},
				"token": {
					Type:        schema.TypeString,
					Sensitive:   true,
					Optional:    true,
					DefaultFunc: schema.EnvDefaultFunc("SSPANEL_TOKEN", nil),
				},
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

type apiClient struct {
	http *resty.Client
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (rClient interface{}, diags diag.Diagnostics) {
		url := d.Get("url").(string)
		token := d.Get("token").(string)

		if !strings.HasSuffix(url, "/admin/api/") {
			url = fmt.Sprintf("%s/admin/api/", url)
		}

		client := resty.New().
			EnableTrace().
			SetTimeout(10*time.Second).
			SetBaseURL(url).
			SetHeader("User-Agent", fmt.Sprintf("terraform-provider-sspanel/%s", version)).
			SetHeader("Authorization", fmt.Sprintf("Bearer %s", token))

		data, err := client.R().Get("/ping")
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to ping server",
				Detail:   fmt.Sprintf("Unable to reach the server: %s", err.Error()),
			})

			return nil, diags
		}

		if data.StatusCode() != 200 {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create client",
				Detail:   "Unable to authenticate user for authenticated admin",
			})

			return nil, diags
		}

		return &apiClient{
			http: client,
		}, nil
	}
}
