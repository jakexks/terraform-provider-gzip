package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	
	"compress/gzip"
	"bytes"
	"encoding/hex"
	"encoding/base64"
	"io"
	"crypto/sha256"
)

func resourceGzipme() *schema.Resource {
    	return &schema.Resource{
		Create: createGzipme,
		Read:   readGzipme,
		//Update: updateGzipme,
		Delete: deleteGzipme,
		Schema: map[string]*schema.Schema{
			"input": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"output": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func createGzipme(d *schema.ResourceData, gzipper interface{}) error {
	data_in := d.Get("input").(string)
	gzbuffer := bytes.Buffer{}
	gzw := gzip.NewWriter(&gzbuffer)
	if _, err := gzw.Write([]byte(data_in)); err != nil {
	    return err
	}
	gzw.Close()
	
	gzr := bytes.NewReader(gzbuffer.Bytes())
	b64buffer := bytes.Buffer{}
	b64w := base64.NewEncoder(base64.StdEncoding, &b64buffer)
	if _, err := io.Copy(b64w, gzr); err != nil {
	    return err
	}
	b64w.Close()
	d.SetId(hash(string(b64buffer.Bytes())))
	return d.Set("output", string(b64buffer.Bytes()))
}

func readGzipme(d *schema.ResourceData, gzipper interface{}) error {
	data_in := d.Get("input").(string)
	gzbuffer := bytes.Buffer{}
	gzw := gzip.NewWriter(&gzbuffer)
	if _, err := gzw.Write([]byte(data_in)); err != nil {
	    return err
	}
	gzw.Close()
	
	gzr := bytes.NewReader(gzbuffer.Bytes())
	b64buffer := bytes.Buffer{}
	b64w := base64.NewEncoder(base64.StdEncoding, &b64buffer)
	if _, err := io.Copy(b64w, gzr); err != nil {
	    return err
	}
	b64w.Close()
	
	d.SetId(hash(string(b64buffer.Bytes())))
	return d.Set("output", string(b64buffer.Bytes()))
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
	
	d.SetId(hash(string(b64buffer.Bytes())))
	return d.Set("output", string(b64buffer.Bytes()))
}

func deleteGzipme(d *schema.ResourceData, gzipper interface{}) error {
	return nil
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}