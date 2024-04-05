Resource for managing an AWS Lightsail Distribution.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

Below is a basic example with a bucket as an origin.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testBucket = new aws.lightsail.Bucket("testBucket", {bundleId: "small_1_0"});
const testDistribution = new aws.lightsail.Distribution("testDistribution", {
    bundleId: "small_1_0",
    origin: {
        name: testBucket.name,
        regionName: testBucket.region,
    },
    defaultCacheBehavior: {
        behavior: "cache",
    },
    cacheBehaviorSettings: {
        allowedHttpMethods: "GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE",
        cachedHttpMethods: "GET,HEAD",
        defaultTtl: 86400,
        maximumTtl: 31536000,
        minimumTtl: 0,
        forwardedCookies: {
            option: "none",
        },
        forwardedHeaders: {
            option: "default",
        },
        forwardedQueryStrings: {
            option: false,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test_bucket = aws.lightsail.Bucket("testBucket", bundle_id="small_1_0")
test_distribution = aws.lightsail.Distribution("testDistribution",
    bundle_id="small_1_0",
    origin=aws.lightsail.DistributionOriginArgs(
        name=test_bucket.name,
        region_name=test_bucket.region,
    ),
    default_cache_behavior=aws.lightsail.DistributionDefaultCacheBehaviorArgs(
        behavior="cache",
    ),
    cache_behavior_settings=aws.lightsail.DistributionCacheBehaviorSettingsArgs(
        allowed_http_methods="GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE",
        cached_http_methods="GET,HEAD",
        default_ttl=86400,
        maximum_ttl=31536000,
        minimum_ttl=0,
        forwarded_cookies=aws.lightsail.DistributionCacheBehaviorSettingsForwardedCookiesArgs(
            option="none",
        ),
        forwarded_headers=aws.lightsail.DistributionCacheBehaviorSettingsForwardedHeadersArgs(
            option="default",
        ),
        forwarded_query_strings=aws.lightsail.DistributionCacheBehaviorSettingsForwardedQueryStringsArgs(
            option=False,
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testBucket = new Aws.LightSail.Bucket("testBucket", new()
    {
        BundleId = "small_1_0",
    });

    var testDistribution = new Aws.LightSail.Distribution("testDistribution", new()
    {
        BundleId = "small_1_0",
        Origin = new Aws.LightSail.Inputs.DistributionOriginArgs
        {
            Name = testBucket.Name,
            RegionName = testBucket.Region,
        },
        DefaultCacheBehavior = new Aws.LightSail.Inputs.DistributionDefaultCacheBehaviorArgs
        {
            Behavior = "cache",
        },
        CacheBehaviorSettings = new Aws.LightSail.Inputs.DistributionCacheBehaviorSettingsArgs
        {
            AllowedHttpMethods = "GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE",
            CachedHttpMethods = "GET,HEAD",
            DefaultTtl = 86400,
            MaximumTtl = 31536000,
            MinimumTtl = 0,
            ForwardedCookies = new Aws.LightSail.Inputs.DistributionCacheBehaviorSettingsForwardedCookiesArgs
            {
                Option = "none",
            },
            ForwardedHeaders = new Aws.LightSail.Inputs.DistributionCacheBehaviorSettingsForwardedHeadersArgs
            {
                Option = "default",
            },
            ForwardedQueryStrings = new Aws.LightSail.Inputs.DistributionCacheBehaviorSettingsForwardedQueryStringsArgs
            {
                Option = false,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testBucket, err := lightsail.NewBucket(ctx, "testBucket", &lightsail.BucketArgs{
			BundleId: pulumi.String("small_1_0"),
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewDistribution(ctx, "testDistribution", &lightsail.DistributionArgs{
			BundleId: pulumi.String("small_1_0"),
			Origin: &lightsail.DistributionOriginArgs{
				Name:       testBucket.Name,
				RegionName: testBucket.Region,
			},
			DefaultCacheBehavior: &lightsail.DistributionDefaultCacheBehaviorArgs{
				Behavior: pulumi.String("cache"),
			},
			CacheBehaviorSettings: &lightsail.DistributionCacheBehaviorSettingsArgs{
				AllowedHttpMethods: pulumi.String("GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE"),
				CachedHttpMethods:  pulumi.String("GET,HEAD"),
				DefaultTtl:         pulumi.Int(86400),
				MaximumTtl:         pulumi.Int(31536000),
				MinimumTtl:         pulumi.Int(0),
				ForwardedCookies: &lightsail.DistributionCacheBehaviorSettingsForwardedCookiesArgs{
					Option: pulumi.String("none"),
				},
				ForwardedHeaders: &lightsail.DistributionCacheBehaviorSettingsForwardedHeadersArgs{
					Option: pulumi.String("default"),
				},
				ForwardedQueryStrings: &lightsail.DistributionCacheBehaviorSettingsForwardedQueryStringsArgs{
					Option: pulumi.Bool(false),
				},
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
import com.pulumi.aws.lightsail.Bucket;
import com.pulumi.aws.lightsail.BucketArgs;
import com.pulumi.aws.lightsail.Distribution;
import com.pulumi.aws.lightsail.DistributionArgs;
import com.pulumi.aws.lightsail.inputs.DistributionOriginArgs;
import com.pulumi.aws.lightsail.inputs.DistributionDefaultCacheBehaviorArgs;
import com.pulumi.aws.lightsail.inputs.DistributionCacheBehaviorSettingsArgs;
import com.pulumi.aws.lightsail.inputs.DistributionCacheBehaviorSettingsForwardedCookiesArgs;
import com.pulumi.aws.lightsail.inputs.DistributionCacheBehaviorSettingsForwardedHeadersArgs;
import com.pulumi.aws.lightsail.inputs.DistributionCacheBehaviorSettingsForwardedQueryStringsArgs;
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
        var testBucket = new Bucket("testBucket", BucketArgs.builder()        
            .bundleId("small_1_0")
            .build());

        var testDistribution = new Distribution("testDistribution", DistributionArgs.builder()        
            .bundleId("small_1_0")
            .origin(DistributionOriginArgs.builder()
                .name(testBucket.name())
                .regionName(testBucket.region())
                .build())
            .defaultCacheBehavior(DistributionDefaultCacheBehaviorArgs.builder()
                .behavior("cache")
                .build())
            .cacheBehaviorSettings(DistributionCacheBehaviorSettingsArgs.builder()
                .allowedHttpMethods("GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE")
                .cachedHttpMethods("GET,HEAD")
                .defaultTtl(86400)
                .maximumTtl(31536000)
                .minimumTtl(0)
                .forwardedCookies(DistributionCacheBehaviorSettingsForwardedCookiesArgs.builder()
                    .option("none")
                    .build())
                .forwardedHeaders(DistributionCacheBehaviorSettingsForwardedHeadersArgs.builder()
                    .option("default")
                    .build())
                .forwardedQueryStrings(DistributionCacheBehaviorSettingsForwardedQueryStringsArgs.builder()
                    .option(false)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testBucket:
    type: aws:lightsail:Bucket
    properties:
      bundleId: small_1_0
  testDistribution:
    type: aws:lightsail:Distribution
    properties:
      bundleId: small_1_0
      origin:
        name: ${testBucket.name}
        regionName: ${testBucket.region}
      defaultCacheBehavior:
        behavior: cache
      cacheBehaviorSettings:
        allowedHttpMethods: GET,HEAD,OPTIONS,PUT,PATCH,POST,DELETE
        cachedHttpMethods: GET,HEAD
        defaultTtl: 86400
        maximumTtl: 3.1536e+07
        minimumTtl: 0
        forwardedCookies:
          option: none
        forwardedHeaders:
          option: default
        forwardedQueryStrings:
          option: false
```
{{% /example %}}
{{% example %}}
### instance origin example

Below is an example of an instance as the origin.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const available = aws.getAvailabilityZones({
    state: "available",
    filters: [{
        name: "opt-in-status",
        values: ["opt-in-not-required"],
    }],
});
const testStaticIp = new aws.lightsail.StaticIp("testStaticIp", {});
const testInstance = new aws.lightsail.Instance("testInstance", {
    availabilityZone: available.then(available => available.names?.[0]),
    blueprintId: "amazon_linux_2",
    bundleId: "micro_1_0",
});
const testStaticIpAttachment = new aws.lightsail.StaticIpAttachment("testStaticIpAttachment", {
    staticIpName: testStaticIp.name,
    instanceName: testInstance.name,
});
const testDistribution = new aws.lightsail.Distribution("testDistribution", {
    bundleId: "small_1_0",
    origin: {
        name: testInstance.name,
        regionName: available.then(available => available.id),
    },
    defaultCacheBehavior: {
        behavior: "cache",
    },
}, {
    dependsOn: [testStaticIpAttachment],
});
```
```python
import pulumi
import pulumi_aws as aws

available = aws.get_availability_zones(state="available",
    filters=[aws.GetAvailabilityZonesFilterArgs(
        name="opt-in-status",
        values=["opt-in-not-required"],
    )])
test_static_ip = aws.lightsail.StaticIp("testStaticIp")
test_instance = aws.lightsail.Instance("testInstance",
    availability_zone=available.names[0],
    blueprint_id="amazon_linux_2",
    bundle_id="micro_1_0")
test_static_ip_attachment = aws.lightsail.StaticIpAttachment("testStaticIpAttachment",
    static_ip_name=test_static_ip.name,
    instance_name=test_instance.name)
test_distribution = aws.lightsail.Distribution("testDistribution",
    bundle_id="small_1_0",
    origin=aws.lightsail.DistributionOriginArgs(
        name=test_instance.name,
        region_name=available.id,
    ),
    default_cache_behavior=aws.lightsail.DistributionDefaultCacheBehaviorArgs(
        behavior="cache",
    ),
    opts=pulumi.ResourceOptions(depends_on=[test_static_ip_attachment]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var available = Aws.GetAvailabilityZones.Invoke(new()
    {
        State = "available",
        Filters = new[]
        {
            new Aws.Inputs.GetAvailabilityZonesFilterInputArgs
            {
                Name = "opt-in-status",
                Values = new[]
                {
                    "opt-in-not-required",
                },
            },
        },
    });

    var testStaticIp = new Aws.LightSail.StaticIp("testStaticIp");

    var testInstance = new Aws.LightSail.Instance("testInstance", new()
    {
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
        BlueprintId = "amazon_linux_2",
        BundleId = "micro_1_0",
    });

    var testStaticIpAttachment = new Aws.LightSail.StaticIpAttachment("testStaticIpAttachment", new()
    {
        StaticIpName = testStaticIp.Name,
        InstanceName = testInstance.Name,
    });

    var testDistribution = new Aws.LightSail.Distribution("testDistribution", new()
    {
        BundleId = "small_1_0",
        Origin = new Aws.LightSail.Inputs.DistributionOriginArgs
        {
            Name = testInstance.Name,
            RegionName = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Id),
        },
        DefaultCacheBehavior = new Aws.LightSail.Inputs.DistributionDefaultCacheBehaviorArgs
        {
            Behavior = "cache",
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            testStaticIpAttachment,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		available, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			State: pulumi.StringRef("available"),
			Filters: []aws.GetAvailabilityZonesFilter{
				{
					Name: "opt-in-status",
					Values: []string{
						"opt-in-not-required",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		testStaticIp, err := lightsail.NewStaticIp(ctx, "testStaticIp", nil)
		if err != nil {
			return err
		}
		testInstance, err := lightsail.NewInstance(ctx, "testInstance", &lightsail.InstanceArgs{
			AvailabilityZone: *pulumi.String(available.Names[0]),
			BlueprintId:      pulumi.String("amazon_linux_2"),
			BundleId:         pulumi.String("micro_1_0"),
		})
		if err != nil {
			return err
		}
		testStaticIpAttachment, err := lightsail.NewStaticIpAttachment(ctx, "testStaticIpAttachment", &lightsail.StaticIpAttachmentArgs{
			StaticIpName: testStaticIp.Name,
			InstanceName: testInstance.Name,
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewDistribution(ctx, "testDistribution", &lightsail.DistributionArgs{
			BundleId: pulumi.String("small_1_0"),
			Origin: &lightsail.DistributionOriginArgs{
				Name:       testInstance.Name,
				RegionName: *pulumi.String(available.Id),
			},
			DefaultCacheBehavior: &lightsail.DistributionDefaultCacheBehaviorArgs{
				Behavior: pulumi.String("cache"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			testStaticIpAttachment,
		}))
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
import com.pulumi.aws.lightsail.StaticIp;
import com.pulumi.aws.lightsail.Instance;
import com.pulumi.aws.lightsail.InstanceArgs;
import com.pulumi.aws.lightsail.StaticIpAttachment;
import com.pulumi.aws.lightsail.StaticIpAttachmentArgs;
import com.pulumi.aws.lightsail.Distribution;
import com.pulumi.aws.lightsail.DistributionArgs;
import com.pulumi.aws.lightsail.inputs.DistributionOriginArgs;
import com.pulumi.aws.lightsail.inputs.DistributionDefaultCacheBehaviorArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        final var available = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
            .state("available")
            .filters(GetAvailabilityZonesFilterArgs.builder()
                .name("opt-in-status")
                .values("opt-in-not-required")
                .build())
            .build());

        var testStaticIp = new StaticIp("testStaticIp");

        var testInstance = new Instance("testInstance", InstanceArgs.builder()        
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]))
            .blueprintId("amazon_linux_2")
            .bundleId("micro_1_0")
            .build());

        var testStaticIpAttachment = new StaticIpAttachment("testStaticIpAttachment", StaticIpAttachmentArgs.builder()        
            .staticIpName(testStaticIp.name())
            .instanceName(testInstance.name())
            .build());

        var testDistribution = new Distribution("testDistribution", DistributionArgs.builder()        
            .bundleId("small_1_0")
            .origin(DistributionOriginArgs.builder()
                .name(testInstance.name())
                .regionName(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.id()))
                .build())
            .defaultCacheBehavior(DistributionDefaultCacheBehaviorArgs.builder()
                .behavior("cache")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(testStaticIpAttachment)
                .build());

    }
}
```
```yaml
resources:
  testStaticIpAttachment:
    type: aws:lightsail:StaticIpAttachment
    properties:
      staticIpName: ${testStaticIp.name}
      instanceName: ${testInstance.name}
  testStaticIp:
    type: aws:lightsail:StaticIp
  testInstance:
    type: aws:lightsail:Instance
    properties:
      availabilityZone: ${available.names[0]}
      blueprintId: amazon_linux_2
      bundleId: micro_1_0
  testDistribution:
    type: aws:lightsail:Distribution
    properties:
      bundleId: small_1_0
      origin:
        name: ${testInstance.name}
        regionName: ${available.id}
      defaultCacheBehavior:
        behavior: cache
    options:
      dependson:
        - ${testStaticIpAttachment}
variables:
  available:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments:
        state: available
        filters:
          - name: opt-in-status
            values:
              - opt-in-not-required
```
{{% /example %}}
{{% example %}}
### lb origin example

Below is an example with a load balancer as an origin

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const available = aws.getAvailabilityZones({
    state: "available",
    filters: [{
        name: "opt-in-status",
        values: ["opt-in-not-required"],
    }],
});
const testLb = new aws.lightsail.Lb("testLb", {
    healthCheckPath: "/",
    instancePort: 80,
    tags: {
        foo: "bar",
    },
});
const testInstance = new aws.lightsail.Instance("testInstance", {
    availabilityZone: available.then(available => available.names?.[0]),
    blueprintId: "amazon_linux_2",
    bundleId: "nano_1_0",
});
const testLbAttachment = new aws.lightsail.LbAttachment("testLbAttachment", {
    lbName: testLb.name,
    instanceName: testInstance.name,
});
const testDistribution = new aws.lightsail.Distribution("testDistribution", {
    bundleId: "small_1_0",
    origin: {
        name: testLb.name,
        regionName: available.then(available => available.id),
    },
    defaultCacheBehavior: {
        behavior: "cache",
    },
}, {
    dependsOn: [testLbAttachment],
});
```
```python
import pulumi
import pulumi_aws as aws

available = aws.get_availability_zones(state="available",
    filters=[aws.GetAvailabilityZonesFilterArgs(
        name="opt-in-status",
        values=["opt-in-not-required"],
    )])
test_lb = aws.lightsail.Lb("testLb",
    health_check_path="/",
    instance_port=80,
    tags={
        "foo": "bar",
    })
test_instance = aws.lightsail.Instance("testInstance",
    availability_zone=available.names[0],
    blueprint_id="amazon_linux_2",
    bundle_id="nano_1_0")
test_lb_attachment = aws.lightsail.LbAttachment("testLbAttachment",
    lb_name=test_lb.name,
    instance_name=test_instance.name)
test_distribution = aws.lightsail.Distribution("testDistribution",
    bundle_id="small_1_0",
    origin=aws.lightsail.DistributionOriginArgs(
        name=test_lb.name,
        region_name=available.id,
    ),
    default_cache_behavior=aws.lightsail.DistributionDefaultCacheBehaviorArgs(
        behavior="cache",
    ),
    opts=pulumi.ResourceOptions(depends_on=[test_lb_attachment]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var available = Aws.GetAvailabilityZones.Invoke(new()
    {
        State = "available",
        Filters = new[]
        {
            new Aws.Inputs.GetAvailabilityZonesFilterInputArgs
            {
                Name = "opt-in-status",
                Values = new[]
                {
                    "opt-in-not-required",
                },
            },
        },
    });

    var testLb = new Aws.LightSail.Lb("testLb", new()
    {
        HealthCheckPath = "/",
        InstancePort = 80,
        Tags = 
        {
            { "foo", "bar" },
        },
    });

    var testInstance = new Aws.LightSail.Instance("testInstance", new()
    {
        AvailabilityZone = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Names[0]),
        BlueprintId = "amazon_linux_2",
        BundleId = "nano_1_0",
    });

    var testLbAttachment = new Aws.LightSail.LbAttachment("testLbAttachment", new()
    {
        LbName = testLb.Name,
        InstanceName = testInstance.Name,
    });

    var testDistribution = new Aws.LightSail.Distribution("testDistribution", new()
    {
        BundleId = "small_1_0",
        Origin = new Aws.LightSail.Inputs.DistributionOriginArgs
        {
            Name = testLb.Name,
            RegionName = available.Apply(getAvailabilityZonesResult => getAvailabilityZonesResult.Id),
        },
        DefaultCacheBehavior = new Aws.LightSail.Inputs.DistributionDefaultCacheBehaviorArgs
        {
            Behavior = "cache",
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            testLbAttachment,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		available, err := aws.GetAvailabilityZones(ctx, &aws.GetAvailabilityZonesArgs{
			State: pulumi.StringRef("available"),
			Filters: []aws.GetAvailabilityZonesFilter{
				{
					Name: "opt-in-status",
					Values: []string{
						"opt-in-not-required",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		testLb, err := lightsail.NewLb(ctx, "testLb", &lightsail.LbArgs{
			HealthCheckPath: pulumi.String("/"),
			InstancePort:    pulumi.Int(80),
			Tags: pulumi.StringMap{
				"foo": pulumi.String("bar"),
			},
		})
		if err != nil {
			return err
		}
		testInstance, err := lightsail.NewInstance(ctx, "testInstance", &lightsail.InstanceArgs{
			AvailabilityZone: *pulumi.String(available.Names[0]),
			BlueprintId:      pulumi.String("amazon_linux_2"),
			BundleId:         pulumi.String("nano_1_0"),
		})
		if err != nil {
			return err
		}
		testLbAttachment, err := lightsail.NewLbAttachment(ctx, "testLbAttachment", &lightsail.LbAttachmentArgs{
			LbName:       testLb.Name,
			InstanceName: testInstance.Name,
		})
		if err != nil {
			return err
		}
		_, err = lightsail.NewDistribution(ctx, "testDistribution", &lightsail.DistributionArgs{
			BundleId: pulumi.String("small_1_0"),
			Origin: &lightsail.DistributionOriginArgs{
				Name:       testLb.Name,
				RegionName: *pulumi.String(available.Id),
			},
			DefaultCacheBehavior: &lightsail.DistributionDefaultCacheBehaviorArgs{
				Behavior: pulumi.String("cache"),
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			testLbAttachment,
		}))
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetAvailabilityZonesArgs;
import com.pulumi.aws.lightsail.Lb;
import com.pulumi.aws.lightsail.LbArgs;
import com.pulumi.aws.lightsail.Instance;
import com.pulumi.aws.lightsail.InstanceArgs;
import com.pulumi.aws.lightsail.LbAttachment;
import com.pulumi.aws.lightsail.LbAttachmentArgs;
import com.pulumi.aws.lightsail.Distribution;
import com.pulumi.aws.lightsail.DistributionArgs;
import com.pulumi.aws.lightsail.inputs.DistributionOriginArgs;
import com.pulumi.aws.lightsail.inputs.DistributionDefaultCacheBehaviorArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        final var available = AwsFunctions.getAvailabilityZones(GetAvailabilityZonesArgs.builder()
            .state("available")
            .filters(GetAvailabilityZonesFilterArgs.builder()
                .name("opt-in-status")
                .values("opt-in-not-required")
                .build())
            .build());

        var testLb = new Lb("testLb", LbArgs.builder()        
            .healthCheckPath("/")
            .instancePort("80")
            .tags(Map.of("foo", "bar"))
            .build());

        var testInstance = new Instance("testInstance", InstanceArgs.builder()        
            .availabilityZone(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.names()[0]))
            .blueprintId("amazon_linux_2")
            .bundleId("nano_1_0")
            .build());

        var testLbAttachment = new LbAttachment("testLbAttachment", LbAttachmentArgs.builder()        
            .lbName(testLb.name())
            .instanceName(testInstance.name())
            .build());

        var testDistribution = new Distribution("testDistribution", DistributionArgs.builder()        
            .bundleId("small_1_0")
            .origin(DistributionOriginArgs.builder()
                .name(testLb.name())
                .regionName(available.applyValue(getAvailabilityZonesResult -> getAvailabilityZonesResult.id()))
                .build())
            .defaultCacheBehavior(DistributionDefaultCacheBehaviorArgs.builder()
                .behavior("cache")
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(testLbAttachment)
                .build());

    }
}
```
```yaml
resources:
  testLb:
    type: aws:lightsail:Lb
    properties:
      healthCheckPath: /
      instancePort: '80'
      tags:
        foo: bar
  testInstance:
    type: aws:lightsail:Instance
    properties:
      availabilityZone: ${available.names[0]}
      blueprintId: amazon_linux_2
      bundleId: nano_1_0
  testLbAttachment:
    type: aws:lightsail:LbAttachment
    properties:
      lbName: ${testLb.name}
      instanceName: ${testInstance.name}
  testDistribution:
    type: aws:lightsail:Distribution
    properties:
      bundleId: small_1_0
      origin:
        name: ${testLb.name}
        regionName: ${available.id}
      defaultCacheBehavior:
        behavior: cache
    options:
      dependson:
        - ${testLbAttachment}
variables:
  available:
    fn::invoke:
      Function: aws:getAvailabilityZones
      Arguments:
        state: available
        filters:
          - name: opt-in-status
            values:
              - opt-in-not-required
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lightsail Distribution using the `id`. For example:

```sh
 $ pulumi import aws:lightsail/distribution:Distribution example rft-8012925589
```
 