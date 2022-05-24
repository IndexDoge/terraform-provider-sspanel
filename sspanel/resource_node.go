package sspanel

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/IndexDoge/terraform-provider-sspanel/sspanel/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			"custom_config": {
				Description: "Node custom config",
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "{}",
			},
		},
	}
}

type createNodeResponse struct {
	types.ApiResponse
	NodeId int64 `json:"node_id"`
}

func generatePostNodeResourceData(d *schema.ResourceData) map[string]interface{} {
	body := map[string]interface{}{
		"name":                    d.Get("name"),
		"server":                  d.Get("server"),
		"node_ip":                 d.Get("node_ip"),
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

	return body
}

func resourceNodeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	// use the meta value to retrieve your client from the provider configure method
	body := generatePostNodeResourceData(d)

	client := meta.(*apiClient)
	res, err := client.http.R().
		SetBody(body).
		Post("node")

	if err != nil {
		return diag.FromErr(err)
	}

	responseJson := createNodeResponse{}

	err = json.NewDecoder(strings.NewReader(res.String())).Decode(&responseJson)
	if err != nil {
		return diag.FromErr(err)
	}

	if responseJson.StatusCode != types.StatusOk {
		return diag.Errorf("Create node failure: %s", responseJson.Msg)
	}
	tflog.Trace(ctx, "created a resource")

	d.SetId(strconv.FormatInt(responseJson.NodeId, 10))

	return resourceNodeRead(ctx, d, meta)
}

func resourceNodeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)

	res, err := client.http.R().
		Get(fmt.Sprintf("node/%s", d.Id()))

	if err != nil {
		return diag.FromErr(err)
	}

	responseJson := types.ReadNodeInfoResponse{}

	err = json.NewDecoder(strings.NewReader(res.String())).Decode(&responseJson)

	if responseJson.StatusCode != types.StatusOk {
		return diag.Errorf("Read node info failure: %s", responseJson.Msg)
	}

	node := responseJson.Node

	// Set Node Id
	d.SetId(strconv.FormatInt(node.Id, 10))

	d.Set("name", node.Name)
	d.Set("server", node.Server)
	d.Set("mu_only", node.MuType)
	d.Set("rate", node.TrafficRate)
	d.Set("info", node.Info)
	d.Set("type", node.Type)
	d.Set("speedlimit", node.SpeedLimit)
	d.Set("sort", node.Sort)
	d.Set("node_ip", node.IP)
	d.Set("status", node.Status)
	d.Set("class", node.Class)
	d.Set("bandwidth_limit", node.BandwidthLimit)
	d.Set("bandwidth_reset_day", node.BandwidthResetDay)
	d.Set("custom_config", node.CustomConfig)

	return nil
}

func resourceNodeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	// client := meta.(*apiClient)
	body := generatePostNodeResourceData(d)
	client := meta.(*apiClient)
	res, err := client.http.R().
		SetBody(body).
		Put(fmt.Sprintf("node/%s", d.Id()))
	if err != nil {
		return diag.FromErr(err)
	}

	responseJson := types.ApiResponse{}

	err = json.NewDecoder(strings.NewReader(res.String())).Decode(&responseJson)

	if responseJson.StatusCode != types.StatusOk {
		return diag.Errorf("Update node failure: %s", responseJson.Msg)
	}

	return resourceNodeRead(ctx, d, meta)
}

func resourceNodeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*apiClient)

	res, err := client.http.R().
		SetQueryParam("id", d.Id()).
		Delete("node")

	if err != nil {
		return diag.FromErr(err)
	}

	responseJson := types.ApiResponse{}

	err = json.NewDecoder(strings.NewReader(res.String())).Decode(&responseJson)

	if responseJson.StatusCode != types.StatusOk {
		return diag.Errorf("Delete node failure: %s", responseJson.Msg)
	}

	return nil
}
