package main

import (
	"github.com/hashicorp/terraform/helper/schema"

	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
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
	result, err := handleinput(d, gzipper.(*GZipper))
	if err != nil {
		return err
	} else {
		d.SetId(hash(result))
		return d.Set("output", result)
	}
}

func readGzipme(d *schema.ResourceData, gzipper interface{}) error {
	result, err := handleinput(d, gzipper.(*GZipper))
	if err != nil {
		return err
	} else {
		d.SetId(hash(result))
		return d.Set("output", result)
	}
}

func updateGzipme(d *schema.ResourceData, gzipper interface{}) error {
	result, err := handleinput(d, gzipper.(*GZipper))
	if err != nil {
		return err
	} else {
		d.SetId(hash(result))
		return d.Set("output", result)
	}
}

func deleteGzipme(d *schema.ResourceData, gzipper interface{}) error {
	return nil
}

func handleinput(d *schema.ResourceData, g *GZipper) (string, error) {
	data_in := d.Get("input").(string)
	gzbuffer := bytes.Buffer{}
	gzw, err := gzip.NewWriterLevel(&gzbuffer, g.CompressionLevel)
	if err != nil {
		return "", err
	}
	if _, err := gzw.Write([]byte(data_in)); err != nil {
		return "", err
	}
	gzw.Close()

	gzr := bytes.NewReader(gzbuffer.Bytes())
	b64buffer := bytes.Buffer{}
	b64w := base64.NewEncoder(base64.StdEncoding, &b64buffer)
	if _, err := io.Copy(b64w, gzr); err != nil {
		return "", err
	}
	b64w.Close()
	return string(b64buffer.Bytes()), nil
}

func hash(s string) string {
	sha := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sha[:])
}
