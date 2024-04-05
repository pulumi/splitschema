Provides a Route53 record resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Simple routing policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const www = new aws.route53.Record("www", {
    zoneId: aws_route53_zone.primary.zone_id,
    name: "www.example.com",
    type: "A",
    ttl: 300,
    records: [aws_eip.lb.public_ip],
});
```
```python
import pulumi
import pulumi_aws as aws

www = aws.route53.Record("www",
    zone_id=aws_route53_zone["primary"]["zone_id"],
    name="www.example.com",
    type="A",
    ttl=300,
    records=[aws_eip["lb"]["public_ip"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var www = new Aws.Route53.Record("www", new()
    {
        ZoneId = aws_route53_zone.Primary.Zone_id,
        Name = "www.example.com",
        Type = "A",
        Ttl = 300,
        Records = new[]
        {
            aws_eip.Lb.Public_ip,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewRecord(ctx, "www", &route53.RecordArgs{
			ZoneId: pulumi.Any(aws_route53_zone.Primary.Zone_id),
			Name:   pulumi.String("www.example.com"),
			Type:   pulumi.String("A"),
			Ttl:    pulumi.Int(300),
			Records: pulumi.StringArray{
				aws_eip.Lb.Public_ip,
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
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
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
        var www = new Record("www", RecordArgs.builder()        
            .zoneId(aws_route53_zone.primary().zone_id())
            .name("www.example.com")
            .type("A")
            .ttl(300)
            .records(aws_eip.lb().public_ip())
            .build());

    }
}
```
```yaml
resources:
  www:
    type: aws:route53:Record
    properties:
      zoneId: ${aws_route53_zone.primary.zone_id}
      name: www.example.com
      type: A
      ttl: 300
      records:
        - ${aws_eip.lb.public_ip}
```
{{% /example %}}
{{% example %}}
### Weighted routing policy

Other routing policies are configured similarly. See [Amazon Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/routing-policy.html) for details.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const www_dev = new aws.route53.Record("www-dev", {
    zoneId: aws_route53_zone.primary.zone_id,
    name: "www",
    type: "CNAME",
    ttl: 5,
    weightedRoutingPolicies: [{
        weight: 10,
    }],
    setIdentifier: "dev",
    records: ["dev.example.com"],
});
const www_live = new aws.route53.Record("www-live", {
    zoneId: aws_route53_zone.primary.zone_id,
    name: "www",
    type: "CNAME",
    ttl: 5,
    weightedRoutingPolicies: [{
        weight: 90,
    }],
    setIdentifier: "live",
    records: ["live.example.com"],
});
```
```python
import pulumi
import pulumi_aws as aws

www_dev = aws.route53.Record("www-dev",
    zone_id=aws_route53_zone["primary"]["zone_id"],
    name="www",
    type="CNAME",
    ttl=5,
    weighted_routing_policies=[aws.route53.RecordWeightedRoutingPolicyArgs(
        weight=10,
    )],
    set_identifier="dev",
    records=["dev.example.com"])
www_live = aws.route53.Record("www-live",
    zone_id=aws_route53_zone["primary"]["zone_id"],
    name="www",
    type="CNAME",
    ttl=5,
    weighted_routing_policies=[aws.route53.RecordWeightedRoutingPolicyArgs(
        weight=90,
    )],
    set_identifier="live",
    records=["live.example.com"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var www_dev = new Aws.Route53.Record("www-dev", new()
    {
        ZoneId = aws_route53_zone.Primary.Zone_id,
        Name = "www",
        Type = "CNAME",
        Ttl = 5,
        WeightedRoutingPolicies = new[]
        {
            new Aws.Route53.Inputs.RecordWeightedRoutingPolicyArgs
            {
                Weight = 10,
            },
        },
        SetIdentifier = "dev",
        Records = new[]
        {
            "dev.example.com",
        },
    });

    var www_live = new Aws.Route53.Record("www-live", new()
    {
        ZoneId = aws_route53_zone.Primary.Zone_id,
        Name = "www",
        Type = "CNAME",
        Ttl = 5,
        WeightedRoutingPolicies = new[]
        {
            new Aws.Route53.Inputs.RecordWeightedRoutingPolicyArgs
            {
                Weight = 90,
            },
        },
        SetIdentifier = "live",
        Records = new[]
        {
            "live.example.com",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewRecord(ctx, "www-dev", &route53.RecordArgs{
			ZoneId: pulumi.Any(aws_route53_zone.Primary.Zone_id),
			Name:   pulumi.String("www"),
			Type:   pulumi.String("CNAME"),
			Ttl:    pulumi.Int(5),
			WeightedRoutingPolicies: route53.RecordWeightedRoutingPolicyArray{
				&route53.RecordWeightedRoutingPolicyArgs{
					Weight: pulumi.Int(10),
				},
			},
			SetIdentifier: pulumi.String("dev"),
			Records: pulumi.StringArray{
				pulumi.String("dev.example.com"),
			},
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "www-live", &route53.RecordArgs{
			ZoneId: pulumi.Any(aws_route53_zone.Primary.Zone_id),
			Name:   pulumi.String("www"),
			Type:   pulumi.String("CNAME"),
			Ttl:    pulumi.Int(5),
			WeightedRoutingPolicies: route53.RecordWeightedRoutingPolicyArray{
				&route53.RecordWeightedRoutingPolicyArgs{
					Weight: pulumi.Int(90),
				},
			},
			SetIdentifier: pulumi.String("live"),
			Records: pulumi.StringArray{
				pulumi.String("live.example.com"),
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
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
import com.pulumi.aws.route53.inputs.RecordWeightedRoutingPolicyArgs;
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
        var www_dev = new Record("www-dev", RecordArgs.builder()        
            .zoneId(aws_route53_zone.primary().zone_id())
            .name("www")
            .type("CNAME")
            .ttl(5)
            .weightedRoutingPolicies(RecordWeightedRoutingPolicyArgs.builder()
                .weight(10)
                .build())
            .setIdentifier("dev")
            .records("dev.example.com")
            .build());

        var www_live = new Record("www-live", RecordArgs.builder()        
            .zoneId(aws_route53_zone.primary().zone_id())
            .name("www")
            .type("CNAME")
            .ttl(5)
            .weightedRoutingPolicies(RecordWeightedRoutingPolicyArgs.builder()
                .weight(90)
                .build())
            .setIdentifier("live")
            .records("live.example.com")
            .build());

    }
}
```
```yaml
resources:
  www-dev:
    type: aws:route53:Record
    properties:
      zoneId: ${aws_route53_zone.primary.zone_id}
      name: www
      type: CNAME
      ttl: 5
      weightedRoutingPolicies:
        - weight: 10
      setIdentifier: dev
      records:
        - dev.example.com
  www-live:
    type: aws:route53:Record
    properties:
      zoneId: ${aws_route53_zone.primary.zone_id}
      name: www
      type: CNAME
      ttl: 5
      weightedRoutingPolicies:
        - weight: 90
      setIdentifier: live
      records:
        - live.example.com
```
{{% /example %}}
{{% example %}}
### Alias record

See [related part of Amazon Route 53 Developer Guide](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/resource-record-sets-choosing-alias-non-alias.html)
to understand differences between alias and non-alias records.

TTL for all alias records is [60 seconds](https://aws.amazon.com/route53/faqs/#dns_failover_do_i_need_to_adjust),
you cannot change this, therefore `ttl` has to be omitted in alias records.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.elb.LoadBalancer("main", {
    availabilityZones: ["us-east-1c"],
    listeners: [{
        instancePort: 80,
        instanceProtocol: "http",
        lbPort: 80,
        lbProtocol: "http",
    }],
});
const www = new aws.route53.Record("www", {
    zoneId: aws_route53_zone.primary.zone_id,
    name: "example.com",
    type: "A",
    aliases: [{
        name: main.dnsName,
        zoneId: main.zoneId,
        evaluateTargetHealth: true,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.elb.LoadBalancer("main",
    availability_zones=["us-east-1c"],
    listeners=[aws.elb.LoadBalancerListenerArgs(
        instance_port=80,
        instance_protocol="http",
        lb_port=80,
        lb_protocol="http",
    )])
www = aws.route53.Record("www",
    zone_id=aws_route53_zone["primary"]["zone_id"],
    name="example.com",
    type="A",
    aliases=[aws.route53.RecordAliasArgs(
        name=main.dns_name,
        zone_id=main.zone_id,
        evaluate_target_health=True,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Elb.LoadBalancer("main", new()
    {
        AvailabilityZones = new[]
        {
            "us-east-1c",
        },
        Listeners = new[]
        {
            new Aws.Elb.Inputs.LoadBalancerListenerArgs
            {
                InstancePort = 80,
                InstanceProtocol = "http",
                LbPort = 80,
                LbProtocol = "http",
            },
        },
    });

    var www = new Aws.Route53.Record("www", new()
    {
        ZoneId = aws_route53_zone.Primary.Zone_id,
        Name = "example.com",
        Type = "A",
        Aliases = new[]
        {
            new Aws.Route53.Inputs.RecordAliasArgs
            {
                Name = main.DnsName,
                ZoneId = main.ZoneId,
                EvaluateTargetHealth = true,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := elb.NewLoadBalancer(ctx, "main", &elb.LoadBalancerArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-east-1c"),
			},
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(80),
					InstanceProtocol: pulumi.String("http"),
					LbPort:           pulumi.Int(80),
					LbProtocol:       pulumi.String("http"),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "www", &route53.RecordArgs{
			ZoneId: pulumi.Any(aws_route53_zone.Primary.Zone_id),
			Name:   pulumi.String("example.com"),
			Type:   pulumi.String("A"),
			Aliases: route53.RecordAliasArray{
				&route53.RecordAliasArgs{
					Name:                 main.DnsName,
					ZoneId:               main.ZoneId,
					EvaluateTargetHealth: pulumi.Bool(true),
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
import com.pulumi.aws.elb.LoadBalancer;
import com.pulumi.aws.elb.LoadBalancerArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerListenerArgs;
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
import com.pulumi.aws.route53.inputs.RecordAliasArgs;
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
        var main = new LoadBalancer("main", LoadBalancerArgs.builder()        
            .availabilityZones("us-east-1c")
            .listeners(LoadBalancerListenerArgs.builder()
                .instancePort(80)
                .instanceProtocol("http")
                .lbPort(80)
                .lbProtocol("http")
                .build())
            .build());

        var www = new Record("www", RecordArgs.builder()        
            .zoneId(aws_route53_zone.primary().zone_id())
            .name("example.com")
            .type("A")
            .aliases(RecordAliasArgs.builder()
                .name(main.dnsName())
                .zoneId(main.zoneId())
                .evaluateTargetHealth(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:elb:LoadBalancer
    properties:
      availabilityZones:
        - us-east-1c
      listeners:
        - instancePort: 80
          instanceProtocol: http
          lbPort: 80
          lbProtocol: http
  www:
    type: aws:route53:Record
    properties:
      zoneId: ${aws_route53_zone.primary.zone_id}
      name: example.com
      type: A
      aliases:
        - name: ${main.dnsName}
          zoneId: ${main.zoneId}
          evaluateTargetHealth: true
```
{{% /example %}}
{{% example %}}
### NS and SOA Record Management

When creating Route 53 zones, the `NS` and `SOA` records for the zone are automatically created. Enabling the `allow_overwrite` argument will allow managing these records in a single deployment without the requirement for `import`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleZone = new aws.route53.Zone("exampleZone", {});
const exampleRecord = new aws.route53.Record("exampleRecord", {
    allowOverwrite: true,
    name: "test.example.com",
    ttl: 172800,
    type: "NS",
    zoneId: exampleZone.zoneId,
    records: [
        exampleZone.nameServers[0],
        exampleZone.nameServers[1],
        exampleZone.nameServers[2],
        exampleZone.nameServers[3],
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example_zone = aws.route53.Zone("exampleZone")
example_record = aws.route53.Record("exampleRecord",
    allow_overwrite=True,
    name="test.example.com",
    ttl=172800,
    type="NS",
    zone_id=example_zone.zone_id,
    records=[
        example_zone.name_servers[0],
        example_zone.name_servers[1],
        example_zone.name_servers[2],
        example_zone.name_servers[3],
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleZone = new Aws.Route53.Zone("exampleZone");

    var exampleRecord = new Aws.Route53.Record("exampleRecord", new()
    {
        AllowOverwrite = true,
        Name = "test.example.com",
        Ttl = 172800,
        Type = "NS",
        ZoneId = exampleZone.ZoneId,
        Records = new[]
        {
            exampleZone.NameServers.Apply(nameServers => nameServers[0]),
            exampleZone.NameServers.Apply(nameServers => nameServers[1]),
            exampleZone.NameServers.Apply(nameServers => nameServers[2]),
            exampleZone.NameServers.Apply(nameServers => nameServers[3]),
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleZone, err := route53.NewZone(ctx, "exampleZone", nil)
		if err != nil {
			return err
		}
		_, err = route53.NewRecord(ctx, "exampleRecord", &route53.RecordArgs{
			AllowOverwrite: pulumi.Bool(true),
			Name:           pulumi.String("test.example.com"),
			Ttl:            pulumi.Int(172800),
			Type:           pulumi.String("NS"),
			ZoneId:         exampleZone.ZoneId,
			Records: pulumi.StringArray{
				exampleZone.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[0], nil
				}).(pulumi.StringOutput),
				exampleZone.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[1], nil
				}).(pulumi.StringOutput),
				exampleZone.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[2], nil
				}).(pulumi.StringOutput),
				exampleZone.NameServers.ApplyT(func(nameServers []string) (string, error) {
					return nameServers[3], nil
				}).(pulumi.StringOutput),
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
import com.pulumi.aws.route53.Zone;
import com.pulumi.aws.route53.Record;
import com.pulumi.aws.route53.RecordArgs;
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
        var exampleZone = new Zone("exampleZone");

        var exampleRecord = new Record("exampleRecord", RecordArgs.builder()        
            .allowOverwrite(true)
            .name("test.example.com")
            .ttl(172800)
            .type("NS")
            .zoneId(exampleZone.zoneId())
            .records(            
                exampleZone.nameServers().applyValue(nameServers -> nameServers[0]),
                exampleZone.nameServers().applyValue(nameServers -> nameServers[1]),
                exampleZone.nameServers().applyValue(nameServers -> nameServers[2]),
                exampleZone.nameServers().applyValue(nameServers -> nameServers[3]))
            .build());

    }
}
```
```yaml
resources:
  exampleZone:
    type: aws:route53:Zone
  exampleRecord:
    type: aws:route53:Record
    properties:
      allowOverwrite: true
      name: test.example.com
      ttl: 172800
      type: NS
      zoneId: ${exampleZone.zoneId}
      records:
        - ${exampleZone.nameServers[0]}
        - ${exampleZone.nameServers[1]}
        - ${exampleZone.nameServers[2]}
        - ${exampleZone.nameServers[3]}
```
{{% /example %}}
{{% /examples %}}

## Import

If the record also contains a set identifier, append it:

If the record name is the empty string, it can be omitted:

__Using `pulumi import` to import__ Route53 Records using the ID of the record, record name, record type, and set identifier. For example:

Using the ID of the record, which is the zone identifier, record name, and record type, separated by underscores (`_`):

```sh
 $ pulumi import aws:route53/record:Record myrecord Z4KAPRWWNC7JR_dev.example.com_NS
```
 If the record also contains a set identifier, append it:

```sh
 $ pulumi import aws:route53/record:Record myrecord Z4KAPRWWNC7JR_dev.example.com_NS_dev
```
 