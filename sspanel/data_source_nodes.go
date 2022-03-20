package sspanel

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNodes() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Nodes read from sspanel api",

		ReadContext: dataSourceNodesRead,

		Schema: map[string]*schema.Schema{
			"nodes": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"custom_config": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"server": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"node_ip": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"info": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"sort": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"traffic_rate": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_class": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_connector": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_bandwidth": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_bandwidth_limit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_speedlimit": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"bandwidthlimit_resetday": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_heartbeat": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_group": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"mu_only": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"online": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"gfw_block": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNodesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) (diags diag.Diagnostics) {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)

	res, err := client.http.R().
		Get("nodes")

	if err != nil {
		return diag.FromErr(err)
	}

	data := make(map[string]interface{}, 0)

	err = json.NewDecoder(strings.NewReader(res.String())).Decode(&data)
	if err != nil {
		return diag.FromErr(err)
	}

	if data["ret"].(float64) != 1 {
		return diag.Errorf("Request status failure")
	}

	if err := d.Set("nodes", data["nodes"]); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
