## Terraform GZIP resource plugin

### Usage:

To install the plugin:

Edit ~/.terraformrc
```
providers {
    gzip = "terraform-provider-gzip"
}


```

Then in your terraform scripts:

```
provider "gzip" {
  compressionlevel = "BestCompression"
}

resource "gzip_me" "example_data" {
    input = "Thing that will be gzipped"
}

```

Then example_data.output should contain your data, but gzipped