# nifcloud-sdk-go-v2

nifcloud-sdk-go-v2 is the Developer Preview for the v2 of the NIFCLOUD SDK for the Go programming language.

# Getting started

## Installing

```sh
go get -u github.com/alice02/nifcloud-sdk-go-v2
```

## Hello NIFCLOUD

```go
package main

import (
        "fmt"

        "github.com/alice02/nifcloud-sdk-go-v2/nifcloud"
        "github.com/alice02/nifcloud-sdk-go-v2/nifcloud/endpoints"
        "github.com/alice02/nifcloud-sdk-go-v2/service/computing"
)

func main() {
        // Create config with credentials and region.
        cfg := nifcloud.NewConfig(
                "YOUR_ACCESS_KEY_ID",
                "YOUR_SECRET_ACCESS_KEY",
                endpoints.JpEast1RegionID,
        )

        // Create the Computing client with Config value.
        svc := computing.New(cfg)
        req := svc.DescribeInstancesRequest(nil)

        // Send the request
        resp, err := req.Send(context.TODO())
        if err != nil {
                panic(err)
        }
        fmt.Println(resp)
}
```
