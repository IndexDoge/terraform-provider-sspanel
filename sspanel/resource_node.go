package sspanel

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
	"strings"
)

func resourceNode() *schema.Resource {
	return &schema.Resource{
		Description: "SSPanel node",

		CreateContext: resourceNodeCreate,
		ReadContext:   resourceNodeRead,
		UpdateContext: resourceNodeUpdate,
		DeleteContext: resourceNodeDelete,

		Schema: map[string]*schema.Schema{
			"id": {
				Description: "Node id",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"name": {
				Description: "Node name",
				Type:        schema.TypeString,
				Required:    true,
			},
			"group": {
				Description: "Node group",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
			},
			"server": {
				Description: "Node server details",
				Type:        schema.TypeString,
				Required:    true,
			},
			"mu_only": {
				Description: "Node for multiuser",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     -1,
			},
			"rate": {
				Description: "Traffic rate",
				Type:        schema.TypeFloat,
				Optional:    true,
				Default:     1.0,
			},
			"info": {
				Description: "Node remark for user",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"type": {
				Description: "Is node enabled",
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
			},
			"speedlimit": {
				Description: "Node speedlimit",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
			},
			"sort": {
				Description: "Node type",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"node_ip": {
				Description: "Node real ip",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"status": {
				Description: "Node status text",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
			},
			"class": {
				Description: "Node class requirement",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
			},
			"bandwidth_limit": {
				Description: "Node bandwidth limit",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
			},
			"bandwidth_reset_day": {
				Description: "The day that bandwidth limit reset",
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
			},
		},
	}
}

func resourceNodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	// use the meta value to retrieve your client from the provider configure method
	body := map[string]interface{}{
		"name":                    d.Get("name"),
		"server":                  d.Get("server"),
		"mu_only":                 d.Get("mu_only"),
		"rate":                    d.Get("rate"),
		"info":                    d.Get("info"),
		"type":                    d.Get("type"),
		"sort":                    d.Get("sort"),
		"node_speedlimit":         d.Get("speedlimit"),
		"group":                   d.Get("group"),
		"status":                  d.Get("status"),
		"class":                   d.Get("class"),
		"node_bandwidth_limit":    d.Get("bandwidth_limit"),
		"bandwidthlimit_resetday": d.Get("bandwidth_reset_day"),
	}

	client := meta.(*apiClient)
	res, err := client.http.R().
		SetBody(body).
		Post("node")

	if err != nil {
		return diag.FromErr(err)
	}

	responseJson := make(map[string]interface{}, 0)

	err = json.NewDecoder(strings.NewReader(res.String())).Decode(&responseJson)
	if err != nil {
		return diag.FromErr(err)
	}

	if responseJson["ret"].(float64) != 1 {
		return diag.Errorf("Create node failure: %s", responseJson["msg"].(string))
	}
	tflog.Trace(ctx, "created a resource")

	d.SetId(strconv.FormatInt(int64(responseJson["node_id"].(float64)), 10))

	return
}

func resourceNodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)

	res, err := client.http.R().
		Get(fmt.Sprintf("node/%s", d.Id()))

	if err != nil {
		return diag.FromErr(err)
	}

	responseJson := make(map[string]interface{}, 0)

	err = json.NewDecoder(strings.NewReader(res.String())).Decode(&responseJson)

	if responseJson["ret"].(float64) != 1 {
		return diag.Errorf("Read node info failure: %s", responseJson["msg"].(string))
	}

	node := responseJson["node"].(map[string]interface{})

	// Set Node Id
	d.SetId(strconv.FormatInt(int64(node["id"].(float64)), 10))

	d.Set("server", node["server"])
	d.Set("name", node["name"])
	d.Set("status", node["status"])
	
	return nil
}

func resourceNodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)

	return diag.Errorf("not implemented")
}

func resourceNodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	res, err := client.http.R().
		SetQueryParam("id", d.Id()).
		Delete("node")

	if err != nil {
		return diag.FromErr(err)
	}

	responseJson := make(map[string]interface{}, 0)

	err = json.NewDecoder(strings.NewReader(res.String())).Decode(&responseJson)

	if responseJson["ret"].(float64) != 1 {
		return diag.Errorf("Delete node failure: %s", responseJson["msg"].(string))
	}

	return nil
}
