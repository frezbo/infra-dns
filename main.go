package main

import (
	"github.com/pulumi/pulumi-cloudflare/sdk/v2/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		zone, err := cloudflare.NewZone(ctx, "frezbo", &cloudflare.ZoneArgs{
			Plan: pulumi.String("free"),
			Type: pulumi.String("full"),
			Zone: pulumi.String("frezbo.dev"),
		})
		if err != nil {
			return err
		}
		ctx.Export("nameservers", zone.NameServers)
		ctx.Export("zone", zone.Zone)
		return nil
	})
}
