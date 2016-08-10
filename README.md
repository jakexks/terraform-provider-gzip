## Terraform GZIP resource plugin

### Usage:

```
provider "gzip" {
  compressionlevel = "BestCompression"
}

resource "gzip_me" "example_data" {
    input = "Thing that will be gzipped"
}

```

Then gzip_me.output should contain your data