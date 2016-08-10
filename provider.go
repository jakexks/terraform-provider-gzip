package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	
	//"compress/gzip"
	"compress/flate"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"compressionlevel": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Default: "DefaultCompression",
				Description: "The amount of compression to use: NoCompression, BestSpeed, BestCompression or DefaultCompression",
			},
		},
		ResourcesMap:  map[string]*schema.Resource{
			"gzip_me": resourceGzipme(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc: configurefunc,
	}
}

func configurefunc (d *schema.ResourceData)(interface{}, error){
	levels := map[string]int{
		"NoCompression": flate.NoCompression,
		"BestSpeed": flate.BestSpeed,
		"BestCompression": flate.BestCompression,
		"DefaultCompression": flate.DefaultCompression,
	}
	
	return &GZipper{
		CompressionLevel: levels[d.Get("compressionlevel").(string)],
	}, nil
}

type GZipper struct {
	CompressionLevel int
}