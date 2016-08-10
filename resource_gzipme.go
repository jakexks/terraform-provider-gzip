package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	
	"compress/gzip"
	"bytes"
	"encoding/base64"
	"io"
)

func resourceGzipme() *schema.Resource {
    	return &schema.Resource{
		Create: createGzipme,
		Read:   readGzipme,
		Update: updateGzipme,
		Delete: deleteGzipme,
		Schema: map[string]*schema.Schema{
			"input": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
			},
		},
	}
}

func createGzipme(d *schema.ResourceData, gzipper interface{}) error {
	data_in := d.Get("input").(string)
	gzbuffer := bytes.Buffer{}
	gzw := gzip.NewWriter(&gzbuffer)
	gzw.Write([]byte(data_in))
	gzw.Close()
	
	gzr := bytes.NewReader(gzbuffer.Bytes())
	b64buffer := bytes.Buffer{}
	b64w := base64.NewEncoder(base64.StdEncoding, &b64buffer)
	io.Copy(b64w, gzr)
	b64w.Close()
	
	d.Set("output", string(b64buffer.Bytes()))
	return nil
}

func readGzipme(d *schema.ResourceData, gzipper interface{}) error {
	return nil
}

func updateGzipme(d *schema.ResourceData, gzipper interface{}) error {
    data_in := d.Get("input").(string)
	gzbuffer := bytes.Buffer{}
	gzw := gzip.NewWriter(&gzbuffer)
	gzw.Write([]byte(data_in))
	gzw.Close()
	
	gzr := bytes.NewReader(gzbuffer.Bytes())
	b64buffer := bytes.Buffer{}
	b64w := base64.NewEncoder(base64.StdEncoding, &b64buffer)
	io.Copy(b64w, gzr)
	b64w.Close()
	
	d.Set("output", string(b64buffer.Bytes()))
	return nil
}

func deleteGzipme(d *schema.ResourceData, gzipper interface{}) error {
	return nil
}

