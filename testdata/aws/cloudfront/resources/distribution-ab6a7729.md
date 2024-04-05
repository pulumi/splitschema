Creates an Amazon CloudFront web distribution.

For information about CloudFront distributions, see the [Amazon CloudFront Developer Guide](http://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/Introduction.html). For specific information about creating CloudFront web distributions, see the [POST Distribution](https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_CreateDistribution.html) page in the Amazon CloudFront API Reference.

> **NOTE:** CloudFront distributions take about 15 minutes to reach a deployed state after creation or modification. During this time, deletes to resources will be blocked. If you need to delete a distribution that is enabled and you do not want to wait, you need to use the `retain_on_delete` flag.

{{% examples %}}
## Example Usage
{{% example %}}
### S3 Origin

The example below creates a CloudFront distribution with an S3 origin.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const bucketV2 = new aws.s3.BucketV2("bucketV2", {tags: {
    Name: "My bucket",
}});
const bAcl = new aws.s3.BucketAclV2("bAcl", {
    bucket: bucketV2.id,
    acl: "private",
});
const s3OriginId = "myS3Origin";
const s3Distribution = new aws.cloudfront.Distribution("s3Distribution", {
    origins: [{
        domainName: bucketV2.bucketRegionalDomainName,
        originAccessControlId: aws_cloudfront_origin_access_control["default"].id,
        originId: s3OriginId,
    }],
    enabled: true,
    isIpv6Enabled: true,
    comment: "Some comment",
    defaultRootObject: "index.html",
    loggingConfig: {
        includeCookies: false,
        bucket: "mylogs.s3.amazonaws.com",
        prefix: "myprefix",
    },
    aliases: [
        "mysite.example.com",
        "yoursite.example.com",
    ],
    defaultCacheBehavior: {
        allowedMethods: [
            "DELETE",
            "GET",
            "HEAD",
            "OPTIONS",
            "PATCH",
            "POST",
            "PUT",
        ],
        cachedMethods: [
            "GET",
            "HEAD",
        ],
        targetOriginId: s3OriginId,
        forwardedValues: {
            queryString: false,
            cookies: {
                forward: "none",
            },
        },
        viewerProtocolPolicy: "allow-all",
        minTtl: 0,
        defaultTtl: 3600,
        maxTtl: 86400,
    },
    orderedCacheBehaviors: [
        {
            pathPattern: "/content/immutable/*",
            allowedMethods: [
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            cachedMethods: [
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            targetOriginId: s3OriginId,
            forwardedValues: {
                queryString: false,
                headers: ["Origin"],
                cookies: {
                    forward: "none",
                },
            },
            minTtl: 0,
            defaultTtl: 86400,
            maxTtl: 31536000,
            compress: true,
            viewerProtocolPolicy: "redirect-to-https",
        },
        {
            pathPattern: "/content/*",
            allowedMethods: [
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            cachedMethods: [
                "GET",
                "HEAD",
            ],
            targetOriginId: s3OriginId,
            forwardedValues: {
                queryString: false,
                cookies: {
                    forward: "none",
                },
            },
            minTtl: 0,
            defaultTtl: 3600,
            maxTtl: 86400,
            compress: true,
            viewerProtocolPolicy: "redirect-to-https",
        },
    ],
    priceClass: "PriceClass_200",
    restrictions: {
        geoRestriction: {
            restrictionType: "whitelist",
            locations: [
                "US",
                "CA",
                "GB",
                "DE",
            ],
        },
    },
    tags: {
        Environment: "production",
    },
    viewerCertificate: {
        cloudfrontDefaultCertificate: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

bucket_v2 = aws.s3.BucketV2("bucketV2", tags={
    "Name": "My bucket",
})
b_acl = aws.s3.BucketAclV2("bAcl",
    bucket=bucket_v2.id,
    acl="private")
s3_origin_id = "myS3Origin"
s3_distribution = aws.cloudfront.Distribution("s3Distribution",
    origins=[aws.cloudfront.DistributionOriginArgs(
        domain_name=bucket_v2.bucket_regional_domain_name,
        origin_access_control_id=aws_cloudfront_origin_access_control["default"]["id"],
        origin_id=s3_origin_id,
    )],
    enabled=True,
    is_ipv6_enabled=True,
    comment="Some comment",
    default_root_object="index.html",
    logging_config=aws.cloudfront.DistributionLoggingConfigArgs(
        include_cookies=False,
        bucket="mylogs.s3.amazonaws.com",
        prefix="myprefix",
    ),
    aliases=[
        "mysite.example.com",
        "yoursite.example.com",
    ],
    default_cache_behavior=aws.cloudfront.DistributionDefaultCacheBehaviorArgs(
        allowed_methods=[
            "DELETE",
            "GET",
            "HEAD",
            "OPTIONS",
            "PATCH",
            "POST",
            "PUT",
        ],
        cached_methods=[
            "GET",
            "HEAD",
        ],
        target_origin_id=s3_origin_id,
        forwarded_values=aws.cloudfront.DistributionDefaultCacheBehaviorForwardedValuesArgs(
            query_string=False,
            cookies=aws.cloudfront.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs(
                forward="none",
            ),
        ),
        viewer_protocol_policy="allow-all",
        min_ttl=0,
        default_ttl=3600,
        max_ttl=86400,
    ),
    ordered_cache_behaviors=[
        aws.cloudfront.DistributionOrderedCacheBehaviorArgs(
            path_pattern="/content/immutable/*",
            allowed_methods=[
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            cached_methods=[
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            target_origin_id=s3_origin_id,
            forwarded_values=aws.cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs(
                query_string=False,
                headers=["Origin"],
                cookies=aws.cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs(
                    forward="none",
                ),
            ),
            min_ttl=0,
            default_ttl=86400,
            max_ttl=31536000,
            compress=True,
            viewer_protocol_policy="redirect-to-https",
        ),
        aws.cloudfront.DistributionOrderedCacheBehaviorArgs(
            path_pattern="/content/*",
            allowed_methods=[
                "GET",
                "HEAD",
                "OPTIONS",
            ],
            cached_methods=[
                "GET",
                "HEAD",
            ],
            target_origin_id=s3_origin_id,
            forwarded_values=aws.cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs(
                query_string=False,
                cookies=aws.cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs(
                    forward="none",
                ),
            ),
            min_ttl=0,
            default_ttl=3600,
            max_ttl=86400,
            compress=True,
            viewer_protocol_policy="redirect-to-https",
        ),
    ],
    price_class="PriceClass_200",
    restrictions=aws.cloudfront.DistributionRestrictionsArgs(
        geo_restriction=aws.cloudfront.DistributionRestrictionsGeoRestrictionArgs(
            restriction_type="whitelist",
            locations=[
                "US",
                "CA",
                "GB",
                "DE",
            ],
        ),
    ),
    tags={
        "Environment": "production",
    },
    viewer_certificate=aws.cloudfront.DistributionViewerCertificateArgs(
        cloudfront_default_certificate=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var bucketV2 = new Aws.S3.BucketV2("bucketV2", new()
    {
        Tags = 
        {
            { "Name", "My bucket" },
        },
    });

    var bAcl = new Aws.S3.BucketAclV2("bAcl", new()
    {
        Bucket = bucketV2.Id,
        Acl = "private",
    });

    var s3OriginId = "myS3Origin";

    var s3Distribution = new Aws.CloudFront.Distribution("s3Distribution", new()
    {
        Origins = new[]
        {
            new Aws.CloudFront.Inputs.DistributionOriginArgs
            {
                DomainName = bucketV2.BucketRegionalDomainName,
                OriginAccessControlId = aws_cloudfront_origin_access_control.Default.Id,
                OriginId = s3OriginId,
            },
        },
        Enabled = true,
        IsIpv6Enabled = true,
        Comment = "Some comment",
        DefaultRootObject = "index.html",
        LoggingConfig = new Aws.CloudFront.Inputs.DistributionLoggingConfigArgs
        {
            IncludeCookies = false,
            Bucket = "mylogs.s3.amazonaws.com",
            Prefix = "myprefix",
        },
        Aliases = new[]
        {
            "mysite.example.com",
            "yoursite.example.com",
        },
        DefaultCacheBehavior = new Aws.CloudFront.Inputs.DistributionDefaultCacheBehaviorArgs
        {
            AllowedMethods = new[]
            {
                "DELETE",
                "GET",
                "HEAD",
                "OPTIONS",
                "PATCH",
                "POST",
                "PUT",
            },
            CachedMethods = new[]
            {
                "GET",
                "HEAD",
            },
            TargetOriginId = s3OriginId,
            ForwardedValues = new Aws.CloudFront.Inputs.DistributionDefaultCacheBehaviorForwardedValuesArgs
            {
                QueryString = false,
                Cookies = new Aws.CloudFront.Inputs.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs
                {
                    Forward = "none",
                },
            },
            ViewerProtocolPolicy = "allow-all",
            MinTtl = 0,
            DefaultTtl = 3600,
            MaxTtl = 86400,
        },
        OrderedCacheBehaviors = new[]
        {
            new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorArgs
            {
                PathPattern = "/content/immutable/*",
                AllowedMethods = new[]
                {
                    "GET",
                    "HEAD",
                    "OPTIONS",
                },
                CachedMethods = new[]
                {
                    "GET",
                    "HEAD",
                    "OPTIONS",
                },
                TargetOriginId = s3OriginId,
                ForwardedValues = new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorForwardedValuesArgs
                {
                    QueryString = false,
                    Headers = new[]
                    {
                        "Origin",
                    },
                    Cookies = new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs
                    {
                        Forward = "none",
                    },
                },
                MinTtl = 0,
                DefaultTtl = 86400,
                MaxTtl = 31536000,
                Compress = true,
                ViewerProtocolPolicy = "redirect-to-https",
            },
            new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorArgs
            {
                PathPattern = "/content/*",
                AllowedMethods = new[]
                {
                    "GET",
                    "HEAD",
                    "OPTIONS",
                },
                CachedMethods = new[]
                {
                    "GET",
                    "HEAD",
                },
                TargetOriginId = s3OriginId,
                ForwardedValues = new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorForwardedValuesArgs
                {
                    QueryString = false,
                    Cookies = new Aws.CloudFront.Inputs.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs
                    {
                        Forward = "none",
                    },
                },
                MinTtl = 0,
                DefaultTtl = 3600,
                MaxTtl = 86400,
                Compress = true,
                ViewerProtocolPolicy = "redirect-to-https",
            },
        },
        PriceClass = "PriceClass_200",
        Restrictions = new Aws.CloudFront.Inputs.DistributionRestrictionsArgs
        {
            GeoRestriction = new Aws.CloudFront.Inputs.DistributionRestrictionsGeoRestrictionArgs
            {
                RestrictionType = "whitelist",
                Locations = new[]
                {
                    "US",
                    "CA",
                    "GB",
                    "DE",
                },
            },
        },
        Tags = 
        {
            { "Environment", "production" },
        },
        ViewerCertificate = new Aws.CloudFront.Inputs.DistributionViewerCertificateArgs
        {
            CloudfrontDefaultCertificate = true,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		bucketV2, err := s3.NewBucketV2(ctx, "bucketV2", &s3.BucketV2Args{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("My bucket"),
			},
		})
		if err != nil {
			return err
		}
		_, err = s3.NewBucketAclV2(ctx, "bAcl", &s3.BucketAclV2Args{
			Bucket: bucketV2.ID(),
			Acl:    pulumi.String("private"),
		})
		if err != nil {
			return err
		}
		s3OriginId := "myS3Origin"
		_, err = cloudfront.NewDistribution(ctx, "s3Distribution", &cloudfront.DistributionArgs{
			Origins: cloudfront.DistributionOriginArray{
				&cloudfront.DistributionOriginArgs{
					DomainName:            bucketV2.BucketRegionalDomainName,
					OriginAccessControlId: pulumi.Any(aws_cloudfront_origin_access_control.Default.Id),
					OriginId:              pulumi.String(s3OriginId),
				},
			},
			Enabled:           pulumi.Bool(true),
			IsIpv6Enabled:     pulumi.Bool(true),
			Comment:           pulumi.String("Some comment"),
			DefaultRootObject: pulumi.String("index.html"),
			LoggingConfig: &cloudfront.DistributionLoggingConfigArgs{
				IncludeCookies: pulumi.Bool(false),
				Bucket:         pulumi.String("mylogs.s3.amazonaws.com"),
				Prefix:         pulumi.String("myprefix"),
			},
			Aliases: pulumi.StringArray{
				pulumi.String("mysite.example.com"),
				pulumi.String("yoursite.example.com"),
			},
			DefaultCacheBehavior: &cloudfront.DistributionDefaultCacheBehaviorArgs{
				AllowedMethods: pulumi.StringArray{
					pulumi.String("DELETE"),
					pulumi.String("GET"),
					pulumi.String("HEAD"),
					pulumi.String("OPTIONS"),
					pulumi.String("PATCH"),
					pulumi.String("POST"),
					pulumi.String("PUT"),
				},
				CachedMethods: pulumi.StringArray{
					pulumi.String("GET"),
					pulumi.String("HEAD"),
				},
				TargetOriginId: pulumi.String(s3OriginId),
				ForwardedValues: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesArgs{
					QueryString: pulumi.Bool(false),
					Cookies: &cloudfront.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs{
						Forward: pulumi.String("none"),
					},
				},
				ViewerProtocolPolicy: pulumi.String("allow-all"),
				MinTtl:               pulumi.Int(0),
				DefaultTtl:           pulumi.Int(3600),
				MaxTtl:               pulumi.Int(86400),
			},
			OrderedCacheBehaviors: cloudfront.DistributionOrderedCacheBehaviorArray{
				&cloudfront.DistributionOrderedCacheBehaviorArgs{
					PathPattern: pulumi.String("/content/immutable/*"),
					AllowedMethods: pulumi.StringArray{
						pulumi.String("GET"),
						pulumi.String("HEAD"),
						pulumi.String("OPTIONS"),
					},
					CachedMethods: pulumi.StringArray{
						pulumi.String("GET"),
						pulumi.String("HEAD"),
						pulumi.String("OPTIONS"),
					},
					TargetOriginId: pulumi.String(s3OriginId),
					ForwardedValues: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs{
						QueryString: pulumi.Bool(false),
						Headers: pulumi.StringArray{
							pulumi.String("Origin"),
						},
						Cookies: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs{
							Forward: pulumi.String("none"),
						},
					},
					MinTtl:               pulumi.Int(0),
					DefaultTtl:           pulumi.Int(86400),
					MaxTtl:               pulumi.Int(31536000),
					Compress:             pulumi.Bool(true),
					ViewerProtocolPolicy: pulumi.String("redirect-to-https"),
				},
				&cloudfront.DistributionOrderedCacheBehaviorArgs{
					PathPattern: pulumi.String("/content/*"),
					AllowedMethods: pulumi.StringArray{
						pulumi.String("GET"),
						pulumi.String("HEAD"),
						pulumi.String("OPTIONS"),
					},
					CachedMethods: pulumi.StringArray{
						pulumi.String("GET"),
						pulumi.String("HEAD"),
					},
					TargetOriginId: pulumi.String(s3OriginId),
					ForwardedValues: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesArgs{
						QueryString: pulumi.Bool(false),
						Cookies: &cloudfront.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs{
							Forward: pulumi.String("none"),
						},
					},
					MinTtl:               pulumi.Int(0),
					DefaultTtl:           pulumi.Int(3600),
					MaxTtl:               pulumi.Int(86400),
					Compress:             pulumi.Bool(true),
					ViewerProtocolPolicy: pulumi.String("redirect-to-https"),
				},
			},
			PriceClass: pulumi.String("PriceClass_200"),
			Restrictions: &cloudfront.DistributionRestrictionsArgs{
				GeoRestriction: &cloudfront.DistributionRestrictionsGeoRestrictionArgs{
					RestrictionType: pulumi.String("whitelist"),
					Locations: pulumi.StringArray{
						pulumi.String("US"),
						pulumi.String("CA"),
						pulumi.String("GB"),
						pulumi.String("DE"),
					},
				},
			},
			Tags: pulumi.StringMap{
				"Environment": pulumi.String("production"),
			},
			ViewerCertificate: &cloudfront.DistributionViewerCertificateArgs{
				CloudfrontDefaultCertificate: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.s3.BucketV2;
import com.pulumi.aws.s3.BucketV2Args;
import com.pulumi.aws.s3.BucketAclV2;
import com.pulumi.aws.s3.BucketAclV2Args;
import com.pulumi.aws.cloudfront.Distribution;
import com.pulumi.aws.cloudfront.DistributionArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOriginArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionLoggingConfigArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionDefaultCacheBehaviorArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionDefaultCacheBehaviorForwardedValuesArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOrderedCacheBehaviorArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOrderedCacheBehaviorForwardedValuesArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionRestrictionsArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionRestrictionsGeoRestrictionArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionViewerCertificateArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var bucketV2 = new BucketV2("bucketV2", BucketV2Args.builder()        
            .tags(Map.of("Name", "My bucket"))
            .build());

        var bAcl = new BucketAclV2("bAcl", BucketAclV2Args.builder()        
            .bucket(bucketV2.id())
            .acl("private")
            .build());

        final var s3OriginId = "myS3Origin";

        var s3Distribution = new Distribution("s3Distribution", DistributionArgs.builder()        
            .origins(DistributionOriginArgs.builder()
                .domainName(bucketV2.bucketRegionalDomainName())
                .originAccessControlId(aws_cloudfront_origin_access_control.default().id())
                .originId(s3OriginId)
                .build())
            .enabled(true)
            .isIpv6Enabled(true)
            .comment("Some comment")
            .defaultRootObject("index.html")
            .loggingConfig(DistributionLoggingConfigArgs.builder()
                .includeCookies(false)
                .bucket("mylogs.s3.amazonaws.com")
                .prefix("myprefix")
                .build())
            .aliases(            
                "mysite.example.com",
                "yoursite.example.com")
            .defaultCacheBehavior(DistributionDefaultCacheBehaviorArgs.builder()
                .allowedMethods(                
                    "DELETE",
                    "GET",
                    "HEAD",
                    "OPTIONS",
                    "PATCH",
                    "POST",
                    "PUT")
                .cachedMethods(                
                    "GET",
                    "HEAD")
                .targetOriginId(s3OriginId)
                .forwardedValues(DistributionDefaultCacheBehaviorForwardedValuesArgs.builder()
                    .queryString(false)
                    .cookies(DistributionDefaultCacheBehaviorForwardedValuesCookiesArgs.builder()
                        .forward("none")
                        .build())
                    .build())
                .viewerProtocolPolicy("allow-all")
                .minTtl(0)
                .defaultTtl(3600)
                .maxTtl(86400)
                .build())
            .orderedCacheBehaviors(            
                DistributionOrderedCacheBehaviorArgs.builder()
                    .pathPattern("/content/immutable/*")
                    .allowedMethods(                    
                        "GET",
                        "HEAD",
                        "OPTIONS")
                    .cachedMethods(                    
                        "GET",
                        "HEAD",
                        "OPTIONS")
                    .targetOriginId(s3OriginId)
                    .forwardedValues(DistributionOrderedCacheBehaviorForwardedValuesArgs.builder()
                        .queryString(false)
                        .headers("Origin")
                        .cookies(DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs.builder()
                            .forward("none")
                            .build())
                        .build())
                    .minTtl(0)
                    .defaultTtl(86400)
                    .maxTtl(31536000)
                    .compress(true)
                    .viewerProtocolPolicy("redirect-to-https")
                    .build(),
                DistributionOrderedCacheBehaviorArgs.builder()
                    .pathPattern("/content/*")
                    .allowedMethods(                    
                        "GET",
                        "HEAD",
                        "OPTIONS")
                    .cachedMethods(                    
                        "GET",
                        "HEAD")
                    .targetOriginId(s3OriginId)
                    .forwardedValues(DistributionOrderedCacheBehaviorForwardedValuesArgs.builder()
                        .queryString(false)
                        .cookies(DistributionOrderedCacheBehaviorForwardedValuesCookiesArgs.builder()
                            .forward("none")
                            .build())
                        .build())
                    .minTtl(0)
                    .defaultTtl(3600)
                    .maxTtl(86400)
                    .compress(true)
                    .viewerProtocolPolicy("redirect-to-https")
                    .build())
            .priceClass("PriceClass_200")
            .restrictions(DistributionRestrictionsArgs.builder()
                .geoRestriction(DistributionRestrictionsGeoRestrictionArgs.builder()
                    .restrictionType("whitelist")
                    .locations(                    
                        "US",
                        "CA",
                        "GB",
                        "DE")
                    .build())
                .build())
            .tags(Map.of("Environment", "production"))
            .viewerCertificate(DistributionViewerCertificateArgs.builder()
                .cloudfrontDefaultCertificate(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  bucketV2:
    type: aws:s3:BucketV2
    properties:
      tags:
        Name: My bucket
  bAcl:
    type: aws:s3:BucketAclV2
    properties:
      bucket: ${bucketV2.id}
      acl: private
  s3Distribution:
    type: aws:cloudfront:Distribution
    properties:
      origins:
        - domainName: ${bucketV2.bucketRegionalDomainName}
          originAccessControlId: ${aws_cloudfront_origin_access_control.default.id}
          originId: ${s3OriginId}
      enabled: true
      isIpv6Enabled: true
      comment: Some comment
      defaultRootObject: index.html
      loggingConfig:
        includeCookies: false
        bucket: mylogs.s3.amazonaws.com
        prefix: myprefix
      aliases:
        - mysite.example.com
        - yoursite.example.com
      defaultCacheBehavior:
        allowedMethods:
          - DELETE
          - GET
          - HEAD
          - OPTIONS
          - PATCH
          - POST
          - PUT
        cachedMethods:
          - GET
          - HEAD
        targetOriginId: ${s3OriginId}
        forwardedValues:
          queryString: false
          cookies:
            forward: none
        viewerProtocolPolicy: allow-all
        minTtl: 0
        defaultTtl: 3600
        maxTtl: 86400
      # Cache behavior with precedence 0
      orderedCacheBehaviors:
        - pathPattern: /content/immutable/*
          allowedMethods:
            - GET
            - HEAD
            - OPTIONS
          cachedMethods:
            - GET
            - HEAD
            - OPTIONS
          targetOriginId: ${s3OriginId}
          forwardedValues:
            queryString: false
            headers:
              - Origin
            cookies:
              forward: none
          minTtl: 0
          defaultTtl: 86400
          maxTtl: 3.1536e+07
          compress: true
          viewerProtocolPolicy: redirect-to-https
        - pathPattern: /content/*
          allowedMethods:
            - GET
            - HEAD
            - OPTIONS
          cachedMethods:
            - GET
            - HEAD
          targetOriginId: ${s3OriginId}
          forwardedValues:
            queryString: false
            cookies:
              forward: none
          minTtl: 0
          defaultTtl: 3600
          maxTtl: 86400
          compress: true
          viewerProtocolPolicy: redirect-to-https
      priceClass: PriceClass_200
      restrictions:
        geoRestriction:
          restrictionType: whitelist
          locations:
            - US
            - CA
            - GB
            - DE
      tags:
        Environment: production
      viewerCertificate:
        cloudfrontDefaultCertificate: true
variables:
  s3OriginId: myS3Origin
```
{{% /example %}}
{{% example %}}
### With Failover Routing

The example below creates a CloudFront distribution with an origin group for failover routing.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const s3Distribution = new aws.cloudfront.Distribution("s3Distribution", {
    originGroups: [{
        originId: "groupS3",
        failoverCriteria: {
            statusCodes: [
                403,
                404,
                500,
                502,
            ],
        },
        members: [
            {
                originId: "primaryS3",
            },
            {
                originId: "failoverS3",
            },
        ],
    }],
    origins: [
        {
            domainName: aws_s3_bucket.primary.bucket_regional_domain_name,
            originId: "primaryS3",
            s3OriginConfig: {
                originAccessIdentity: aws_cloudfront_origin_access_identity["default"].cloudfront_access_identity_path,
            },
        },
        {
            domainName: aws_s3_bucket.failover.bucket_regional_domain_name,
            originId: "failoverS3",
            s3OriginConfig: {
                originAccessIdentity: aws_cloudfront_origin_access_identity["default"].cloudfront_access_identity_path,
            },
        },
    ],
    defaultCacheBehavior: {
        targetOriginId: "groupS3",
    },
});
// ... other configuration ...
```
```python
import pulumi
import pulumi_aws as aws

s3_distribution = aws.cloudfront.Distribution("s3Distribution",
    origin_groups=[aws.cloudfront.DistributionOriginGroupArgs(
        origin_id="groupS3",
        failover_criteria=aws.cloudfront.DistributionOriginGroupFailoverCriteriaArgs(
            status_codes=[
                403,
                404,
                500,
                502,
            ],
        ),
        members=[
            aws.cloudfront.DistributionOriginGroupMemberArgs(
                origin_id="primaryS3",
            ),
            aws.cloudfront.DistributionOriginGroupMemberArgs(
                origin_id="failoverS3",
            ),
        ],
    )],
    origins=[
        aws.cloudfront.DistributionOriginArgs(
            domain_name=aws_s3_bucket["primary"]["bucket_regional_domain_name"],
            origin_id="primaryS3",
            s3_origin_config=aws.cloudfront.DistributionOriginS3OriginConfigArgs(
                origin_access_identity=aws_cloudfront_origin_access_identity["default"]["cloudfront_access_identity_path"],
            ),
        ),
        aws.cloudfront.DistributionOriginArgs(
            domain_name=aws_s3_bucket["failover"]["bucket_regional_domain_name"],
            origin_id="failoverS3",
            s3_origin_config=aws.cloudfront.DistributionOriginS3OriginConfigArgs(
                origin_access_identity=aws_cloudfront_origin_access_identity["default"]["cloudfront_access_identity_path"],
            ),
        ),
    ],
    default_cache_behavior=aws.cloudfront.DistributionDefaultCacheBehaviorArgs(
        target_origin_id="groupS3",
    ))
# ... other configuration ...
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var s3Distribution = new Aws.CloudFront.Distribution("s3Distribution", new()
    {
        OriginGroups = new[]
        {
            new Aws.CloudFront.Inputs.DistributionOriginGroupArgs
            {
                OriginId = "groupS3",
                FailoverCriteria = new Aws.CloudFront.Inputs.DistributionOriginGroupFailoverCriteriaArgs
                {
                    StatusCodes = new[]
                    {
                        403,
                        404,
                        500,
                        502,
                    },
                },
                Members = new[]
                {
                    new Aws.CloudFront.Inputs.DistributionOriginGroupMemberArgs
                    {
                        OriginId = "primaryS3",
                    },
                    new Aws.CloudFront.Inputs.DistributionOriginGroupMemberArgs
                    {
                        OriginId = "failoverS3",
                    },
                },
            },
        },
        Origins = new[]
        {
            new Aws.CloudFront.Inputs.DistributionOriginArgs
            {
                DomainName = aws_s3_bucket.Primary.Bucket_regional_domain_name,
                OriginId = "primaryS3",
                S3OriginConfig = new Aws.CloudFront.Inputs.DistributionOriginS3OriginConfigArgs
                {
                    OriginAccessIdentity = aws_cloudfront_origin_access_identity.Default.Cloudfront_access_identity_path,
                },
            },
            new Aws.CloudFront.Inputs.DistributionOriginArgs
            {
                DomainName = aws_s3_bucket.Failover.Bucket_regional_domain_name,
                OriginId = "failoverS3",
                S3OriginConfig = new Aws.CloudFront.Inputs.DistributionOriginS3OriginConfigArgs
                {
                    OriginAccessIdentity = aws_cloudfront_origin_access_identity.Default.Cloudfront_access_identity_path,
                },
            },
        },
        DefaultCacheBehavior = new Aws.CloudFront.Inputs.DistributionDefaultCacheBehaviorArgs
        {
            TargetOriginId = "groupS3",
        },
    });

    // ... other configuration ...
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.NewDistribution(ctx, "s3Distribution", &cloudfront.DistributionArgs{
			OriginGroups: cloudfront.DistributionOriginGroupArray{
				&cloudfront.DistributionOriginGroupArgs{
					OriginId: pulumi.String("groupS3"),
					FailoverCriteria: &cloudfront.DistributionOriginGroupFailoverCriteriaArgs{
						StatusCodes: pulumi.IntArray{
							pulumi.Int(403),
							pulumi.Int(404),
							pulumi.Int(500),
							pulumi.Int(502),
						},
					},
					Members: cloudfront.DistributionOriginGroupMemberArray{
						&cloudfront.DistributionOriginGroupMemberArgs{
							OriginId: pulumi.String("primaryS3"),
						},
						&cloudfront.DistributionOriginGroupMemberArgs{
							OriginId: pulumi.String("failoverS3"),
						},
					},
				},
			},
			Origins: cloudfront.DistributionOriginArray{
				&cloudfront.DistributionOriginArgs{
					DomainName: pulumi.Any(aws_s3_bucket.Primary.Bucket_regional_domain_name),
					OriginId:   pulumi.String("primaryS3"),
					S3OriginConfig: &cloudfront.DistributionOriginS3OriginConfigArgs{
						OriginAccessIdentity: pulumi.Any(aws_cloudfront_origin_access_identity.Default.Cloudfront_access_identity_path),
					},
				},
				&cloudfront.DistributionOriginArgs{
					DomainName: pulumi.Any(aws_s3_bucket.Failover.Bucket_regional_domain_name),
					OriginId:   pulumi.String("failoverS3"),
					S3OriginConfig: &cloudfront.DistributionOriginS3OriginConfigArgs{
						OriginAccessIdentity: pulumi.Any(aws_cloudfront_origin_access_identity.Default.Cloudfront_access_identity_path),
					},
				},
			},
			DefaultCacheBehavior: &cloudfront.DistributionDefaultCacheBehaviorArgs{
				TargetOriginId: pulumi.String("groupS3"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.cloudfront.Distribution;
import com.pulumi.aws.cloudfront.DistributionArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOriginGroupArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOriginGroupFailoverCriteriaArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOriginArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOriginS3OriginConfigArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionDefaultCacheBehaviorArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        var s3Distribution = new Distribution("s3Distribution", DistributionArgs.builder()        
            .originGroups(DistributionOriginGroupArgs.builder()
                .originId("groupS3")
                .failoverCriteria(DistributionOriginGroupFailoverCriteriaArgs.builder()
                    .statusCodes(                    
                        403,
                        404,
                        500,
                        502)
                    .build())
                .members(                
                    DistributionOriginGroupMemberArgs.builder()
                        .originId("primaryS3")
                        .build(),
                    DistributionOriginGroupMemberArgs.builder()
                        .originId("failoverS3")
                        .build())
                .build())
            .origins(            
                DistributionOriginArgs.builder()
                    .domainName(aws_s3_bucket.primary().bucket_regional_domain_name())
                    .originId("primaryS3")
                    .s3OriginConfig(DistributionOriginS3OriginConfigArgs.builder()
                        .originAccessIdentity(aws_cloudfront_origin_access_identity.default().cloudfront_access_identity_path())
                        .build())
                    .build(),
                DistributionOriginArgs.builder()
                    .domainName(aws_s3_bucket.failover().bucket_regional_domain_name())
                    .originId("failoverS3")
                    .s3OriginConfig(DistributionOriginS3OriginConfigArgs.builder()
                        .originAccessIdentity(aws_cloudfront_origin_access_identity.default().cloudfront_access_identity_path())
                        .build())
                    .build())
            .defaultCacheBehavior(DistributionDefaultCacheBehaviorArgs.builder()
                .targetOriginId("groupS3")
                .build())
            .build());

    }
}
```
```yaml
resources:
  s3Distribution:
    type: aws:cloudfront:Distribution
    properties:
      originGroups:
        - originId: groupS3
          failoverCriteria:
            statusCodes:
              - 403
              - 404
              - 500
              - 502
          members:
            - originId: primaryS3
            - originId: failoverS3
      origins:
        - domainName: ${aws_s3_bucket.primary.bucket_regional_domain_name}
          originId: primaryS3
          s3OriginConfig:
            originAccessIdentity: ${aws_cloudfront_origin_access_identity.default.cloudfront_access_identity_path}
        - domainName: ${aws_s3_bucket.failover.bucket_regional_domain_name}
          originId: failoverS3
          s3OriginConfig:
            originAccessIdentity: ${aws_cloudfront_origin_access_identity.default.cloudfront_access_identity_path}
      defaultCacheBehavior:
        targetOriginId: groupS3
```
{{% /example %}}
{{% example %}}
### With Managed Caching Policy

The example below creates a CloudFront distribution with an [AWS managed caching policy](https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/using-managed-cache-policies.html).

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const s3OriginId = "myS3Origin";
const s3Distribution = new aws.cloudfront.Distribution("s3Distribution", {
    origins: [{
        domainName: aws_s3_bucket.primary.bucket_regional_domain_name,
        originId: "myS3Origin",
        s3OriginConfig: {
            originAccessIdentity: aws_cloudfront_origin_access_identity["default"].cloudfront_access_identity_path,
        },
    }],
    enabled: true,
    isIpv6Enabled: true,
    comment: "Some comment",
    defaultRootObject: "index.html",
    defaultCacheBehavior: {
        cachePolicyId: "4135ea2d-6df8-44a3-9df3-4b5a84be39ad",
        allowedMethods: [
            "GET",
            "HEAD",
            "OPTIONS",
        ],
        targetOriginId: s3OriginId,
    },
    restrictions: {
        geoRestriction: {
            restrictionType: "whitelist",
            locations: [
                "US",
                "CA",
                "GB",
                "DE",
            ],
        },
    },
    viewerCertificate: {
        cloudfrontDefaultCertificate: true,
    },
});
// ... other configuration ...
```
```python
import pulumi
import pulumi_aws as aws

s3_origin_id = "myS3Origin"
s3_distribution = aws.cloudfront.Distribution("s3Distribution",
    origins=[aws.cloudfront.DistributionOriginArgs(
        domain_name=aws_s3_bucket["primary"]["bucket_regional_domain_name"],
        origin_id="myS3Origin",
        s3_origin_config=aws.cloudfront.DistributionOriginS3OriginConfigArgs(
            origin_access_identity=aws_cloudfront_origin_access_identity["default"]["cloudfront_access_identity_path"],
        ),
    )],
    enabled=True,
    is_ipv6_enabled=True,
    comment="Some comment",
    default_root_object="index.html",
    default_cache_behavior=aws.cloudfront.DistributionDefaultCacheBehaviorArgs(
        cache_policy_id="4135ea2d-6df8-44a3-9df3-4b5a84be39ad",
        allowed_methods=[
            "GET",
            "HEAD",
            "OPTIONS",
        ],
        target_origin_id=s3_origin_id,
    ),
    restrictions=aws.cloudfront.DistributionRestrictionsArgs(
        geo_restriction=aws.cloudfront.DistributionRestrictionsGeoRestrictionArgs(
            restriction_type="whitelist",
            locations=[
                "US",
                "CA",
                "GB",
                "DE",
            ],
        ),
    ),
    viewer_certificate=aws.cloudfront.DistributionViewerCertificateArgs(
        cloudfront_default_certificate=True,
    ))
# ... other configuration ...
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var s3OriginId = "myS3Origin";

    var s3Distribution = new Aws.CloudFront.Distribution("s3Distribution", new()
    {
        Origins = new[]
        {
            new Aws.CloudFront.Inputs.DistributionOriginArgs
            {
                DomainName = aws_s3_bucket.Primary.Bucket_regional_domain_name,
                OriginId = "myS3Origin",
                S3OriginConfig = new Aws.CloudFront.Inputs.DistributionOriginS3OriginConfigArgs
                {
                    OriginAccessIdentity = aws_cloudfront_origin_access_identity.Default.Cloudfront_access_identity_path,
                },
            },
        },
        Enabled = true,
        IsIpv6Enabled = true,
        Comment = "Some comment",
        DefaultRootObject = "index.html",
        DefaultCacheBehavior = new Aws.CloudFront.Inputs.DistributionDefaultCacheBehaviorArgs
        {
            CachePolicyId = "4135ea2d-6df8-44a3-9df3-4b5a84be39ad",
            AllowedMethods = new[]
            {
                "GET",
                "HEAD",
                "OPTIONS",
            },
            TargetOriginId = s3OriginId,
        },
        Restrictions = new Aws.CloudFront.Inputs.DistributionRestrictionsArgs
        {
            GeoRestriction = new Aws.CloudFront.Inputs.DistributionRestrictionsGeoRestrictionArgs
            {
                RestrictionType = "whitelist",
                Locations = new[]
                {
                    "US",
                    "CA",
                    "GB",
                    "DE",
                },
            },
        },
        ViewerCertificate = new Aws.CloudFront.Inputs.DistributionViewerCertificateArgs
        {
            CloudfrontDefaultCertificate = true,
        },
    });

    // ... other configuration ...
});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		s3OriginId := "myS3Origin"
		_, err := cloudfront.NewDistribution(ctx, "s3Distribution", &cloudfront.DistributionArgs{
			Origins: cloudfront.DistributionOriginArray{
				&cloudfront.DistributionOriginArgs{
					DomainName: pulumi.Any(aws_s3_bucket.Primary.Bucket_regional_domain_name),
					OriginId:   pulumi.String("myS3Origin"),
					S3OriginConfig: &cloudfront.DistributionOriginS3OriginConfigArgs{
						OriginAccessIdentity: pulumi.Any(aws_cloudfront_origin_access_identity.Default.Cloudfront_access_identity_path),
					},
				},
			},
			Enabled:           pulumi.Bool(true),
			IsIpv6Enabled:     pulumi.Bool(true),
			Comment:           pulumi.String("Some comment"),
			DefaultRootObject: pulumi.String("index.html"),
			DefaultCacheBehavior: &cloudfront.DistributionDefaultCacheBehaviorArgs{
				CachePolicyId: pulumi.String("4135ea2d-6df8-44a3-9df3-4b5a84be39ad"),
				AllowedMethods: pulumi.StringArray{
					pulumi.String("GET"),
					pulumi.String("HEAD"),
					pulumi.String("OPTIONS"),
				},
				TargetOriginId: pulumi.String(s3OriginId),
			},
			Restrictions: &cloudfront.DistributionRestrictionsArgs{
				GeoRestriction: &cloudfront.DistributionRestrictionsGeoRestrictionArgs{
					RestrictionType: pulumi.String("whitelist"),
					Locations: pulumi.StringArray{
						pulumi.String("US"),
						pulumi.String("CA"),
						pulumi.String("GB"),
						pulumi.String("DE"),
					},
				},
			},
			ViewerCertificate: &cloudfront.DistributionViewerCertificateArgs{
				CloudfrontDefaultCertificate: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.cloudfront.Distribution;
import com.pulumi.aws.cloudfront.DistributionArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOriginArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionOriginS3OriginConfigArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionDefaultCacheBehaviorArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionRestrictionsArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionRestrictionsGeoRestrictionArgs;
import com.pulumi.aws.cloudfront.inputs.DistributionViewerCertificateArgs;
import java.util.List;
import java.util.ArrayList;
import java.util.Map;
import java.io.File;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Pulumi.run(App::stack);
    }

    public static void stack(Context ctx) {
        final var s3OriginId = "myS3Origin";

        var s3Distribution = new Distribution("s3Distribution", DistributionArgs.builder()        
            .origins(DistributionOriginArgs.builder()
                .domainName(aws_s3_bucket.primary().bucket_regional_domain_name())
                .originId("myS3Origin")
                .s3OriginConfig(DistributionOriginS3OriginConfigArgs.builder()
                    .originAccessIdentity(aws_cloudfront_origin_access_identity.default().cloudfront_access_identity_path())
                    .build())
                .build())
            .enabled(true)
            .isIpv6Enabled(true)
            .comment("Some comment")
            .defaultRootObject("index.html")
            .defaultCacheBehavior(DistributionDefaultCacheBehaviorArgs.builder()
                .cachePolicyId("4135ea2d-6df8-44a3-9df3-4b5a84be39ad")
                .allowedMethods(                
                    "GET",
                    "HEAD",
                    "OPTIONS")
                .targetOriginId(s3OriginId)
                .build())
            .restrictions(DistributionRestrictionsArgs.builder()
                .geoRestriction(DistributionRestrictionsGeoRestrictionArgs.builder()
                    .restrictionType("whitelist")
                    .locations(                    
                        "US",
                        "CA",
                        "GB",
                        "DE")
                    .build())
                .build())
            .viewerCertificate(DistributionViewerCertificateArgs.builder()
                .cloudfrontDefaultCertificate(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  s3Distribution:
    type: aws:cloudfront:Distribution
    properties:
      origins:
        - domainName: ${aws_s3_bucket.primary.bucket_regional_domain_name}
          originId: myS3Origin
          s3OriginConfig:
            originAccessIdentity: ${aws_cloudfront_origin_access_identity.default.cloudfront_access_identity_path}
      enabled: true
      isIpv6Enabled: true
      comment: Some comment
      defaultRootObject: index.html
      defaultCacheBehavior:
        cachePolicyId: 4135ea2d-6df8-44a3-9df3-4b5a84be39ad
        allowedMethods:
          - GET
          - HEAD
          - OPTIONS
        targetOriginId: ${s3OriginId}
      restrictions:
        geoRestriction:
          restrictionType: whitelist
          locations:
            - US
            - CA
            - GB
            - DE
      viewerCertificate:
        cloudfrontDefaultCertificate: true
variables:
  s3OriginId: myS3Origin
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudFront Distributions using the `id`. For example:

```sh
 $ pulumi import aws:cloudfront/distribution:Distribution distribution E74FTE3EXAMPLE
```
 