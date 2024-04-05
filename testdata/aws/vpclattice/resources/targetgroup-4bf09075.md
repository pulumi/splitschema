Resource for managing an AWS VPC Lattice Target Group.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpclattice.TargetGroup("example", {
    type: "INSTANCE",
    config: {
        vpcIdentifier: aws_vpc.example.id,
        port: 443,
        protocol: "HTTPS",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.TargetGroup("example",
    type="INSTANCE",
    config=aws.vpclattice.TargetGroupConfigArgs(
        vpc_identifier=aws_vpc["example"]["id"],
        port=443,
        protocol="HTTPS",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VpcLattice.TargetGroup("example", new()
    {
        Type = "INSTANCE",
        Config = new Aws.VpcLattice.Inputs.TargetGroupConfigArgs
        {
            VpcIdentifier = aws_vpc.Example.Id,
            Port = 443,
            Protocol = "HTTPS",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.NewTargetGroup(ctx, "example", &vpclattice.TargetGroupArgs{
			Type: pulumi.String("INSTANCE"),
			Config: &vpclattice.TargetGroupConfigArgs{
				VpcIdentifier: pulumi.Any(aws_vpc.Example.Id),
				Port:          pulumi.Int(443),
				Protocol:      pulumi.String("HTTPS"),
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
import com.pulumi.aws.vpclattice.TargetGroup;
import com.pulumi.aws.vpclattice.TargetGroupArgs;
import com.pulumi.aws.vpclattice.inputs.TargetGroupConfigArgs;
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
        var example = new TargetGroup("example", TargetGroupArgs.builder()        
            .type("INSTANCE")
            .config(TargetGroupConfigArgs.builder()
                .vpcIdentifier(aws_vpc.example().id())
                .port(443)
                .protocol("HTTPS")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:vpclattice:TargetGroup
    properties:
      type: INSTANCE
      config:
        vpcIdentifier: ${aws_vpc.example.id}
        port: 443
        protocol: HTTPS
```
{{% /example %}}
{{% example %}}
### Basic usage with Health check

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpclattice.TargetGroup("example", {
    type: "IP",
    config: {
        vpcIdentifier: aws_vpc.example.id,
        ipAddressType: "IPV4",
        port: 443,
        protocol: "HTTPS",
        protocolVersion: "HTTP1",
        healthCheck: {
            enabled: true,
            healthCheckIntervalSeconds: 20,
            healthCheckTimeoutSeconds: 10,
            healthyThresholdCount: 7,
            unhealthyThresholdCount: 3,
            matcher: {
                value: "200-299",
            },
            path: "/instance",
            port: 80,
            protocol: "HTTP",
            protocolVersion: "HTTP1",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.TargetGroup("example",
    type="IP",
    config=aws.vpclattice.TargetGroupConfigArgs(
        vpc_identifier=aws_vpc["example"]["id"],
        ip_address_type="IPV4",
        port=443,
        protocol="HTTPS",
        protocol_version="HTTP1",
        health_check=aws.vpclattice.TargetGroupConfigHealthCheckArgs(
            enabled=True,
            health_check_interval_seconds=20,
            health_check_timeout_seconds=10,
            healthy_threshold_count=7,
            unhealthy_threshold_count=3,
            matcher=aws.vpclattice.TargetGroupConfigHealthCheckMatcherArgs(
                value="200-299",
            ),
            path="/instance",
            port=80,
            protocol="HTTP",
            protocol_version="HTTP1",
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
    var example = new Aws.VpcLattice.TargetGroup("example", new()
    {
        Type = "IP",
        Config = new Aws.VpcLattice.Inputs.TargetGroupConfigArgs
        {
            VpcIdentifier = aws_vpc.Example.Id,
            IpAddressType = "IPV4",
            Port = 443,
            Protocol = "HTTPS",
            ProtocolVersion = "HTTP1",
            HealthCheck = new Aws.VpcLattice.Inputs.TargetGroupConfigHealthCheckArgs
            {
                Enabled = true,
                HealthCheckIntervalSeconds = 20,
                HealthCheckTimeoutSeconds = 10,
                HealthyThresholdCount = 7,
                UnhealthyThresholdCount = 3,
                Matcher = new Aws.VpcLattice.Inputs.TargetGroupConfigHealthCheckMatcherArgs
                {
                    Value = "200-299",
                },
                Path = "/instance",
                Port = 80,
                Protocol = "HTTP",
                ProtocolVersion = "HTTP1",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.NewTargetGroup(ctx, "example", &vpclattice.TargetGroupArgs{
			Type: pulumi.String("IP"),
			Config: &vpclattice.TargetGroupConfigArgs{
				VpcIdentifier:   pulumi.Any(aws_vpc.Example.Id),
				IpAddressType:   pulumi.String("IPV4"),
				Port:            pulumi.Int(443),
				Protocol:        pulumi.String("HTTPS"),
				ProtocolVersion: pulumi.String("HTTP1"),
				HealthCheck: &vpclattice.TargetGroupConfigHealthCheckArgs{
					Enabled:                    pulumi.Bool(true),
					HealthCheckIntervalSeconds: pulumi.Int(20),
					HealthCheckTimeoutSeconds:  pulumi.Int(10),
					HealthyThresholdCount:      pulumi.Int(7),
					UnhealthyThresholdCount:    pulumi.Int(3),
					Matcher: &vpclattice.TargetGroupConfigHealthCheckMatcherArgs{
						Value: pulumi.String("200-299"),
					},
					Path:            pulumi.String("/instance"),
					Port:            pulumi.Int(80),
					Protocol:        pulumi.String("HTTP"),
					ProtocolVersion: pulumi.String("HTTP1"),
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
import com.pulumi.aws.vpclattice.TargetGroup;
import com.pulumi.aws.vpclattice.TargetGroupArgs;
import com.pulumi.aws.vpclattice.inputs.TargetGroupConfigArgs;
import com.pulumi.aws.vpclattice.inputs.TargetGroupConfigHealthCheckArgs;
import com.pulumi.aws.vpclattice.inputs.TargetGroupConfigHealthCheckMatcherArgs;
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
        var example = new TargetGroup("example", TargetGroupArgs.builder()        
            .type("IP")
            .config(TargetGroupConfigArgs.builder()
                .vpcIdentifier(aws_vpc.example().id())
                .ipAddressType("IPV4")
                .port(443)
                .protocol("HTTPS")
                .protocolVersion("HTTP1")
                .healthCheck(TargetGroupConfigHealthCheckArgs.builder()
                    .enabled(true)
                    .healthCheckIntervalSeconds(20)
                    .healthCheckTimeoutSeconds(10)
                    .healthyThresholdCount(7)
                    .unhealthyThresholdCount(3)
                    .matcher(TargetGroupConfigHealthCheckMatcherArgs.builder()
                        .value("200-299")
                        .build())
                    .path("/instance")
                    .port(80)
                    .protocol("HTTP")
                    .protocolVersion("HTTP1")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:vpclattice:TargetGroup
    properties:
      type: IP
      config:
        vpcIdentifier: ${aws_vpc.example.id}
        ipAddressType: IPV4
        port: 443
        protocol: HTTPS
        protocolVersion: HTTP1
        healthCheck:
          enabled: true
          healthCheckIntervalSeconds: 20
          healthCheckTimeoutSeconds: 10
          healthyThresholdCount: 7
          unhealthyThresholdCount: 3
          matcher:
            value: 200-299
          path: /instance
          port: 80
          protocol: HTTP
          protocolVersion: HTTP1
```
{{% /example %}}
{{% example %}}
### ALB

If the type is ALB, `health_check` block is not supported.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpclattice.TargetGroup("example", {
    type: "ALB",
    config: {
        vpcIdentifier: aws_vpc.example.id,
        port: 443,
        protocol: "HTTPS",
        protocolVersion: "HTTP1",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.TargetGroup("example",
    type="ALB",
    config=aws.vpclattice.TargetGroupConfigArgs(
        vpc_identifier=aws_vpc["example"]["id"],
        port=443,
        protocol="HTTPS",
        protocol_version="HTTP1",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VpcLattice.TargetGroup("example", new()
    {
        Type = "ALB",
        Config = new Aws.VpcLattice.Inputs.TargetGroupConfigArgs
        {
            VpcIdentifier = aws_vpc.Example.Id,
            Port = 443,
            Protocol = "HTTPS",
            ProtocolVersion = "HTTP1",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.NewTargetGroup(ctx, "example", &vpclattice.TargetGroupArgs{
			Type: pulumi.String("ALB"),
			Config: &vpclattice.TargetGroupConfigArgs{
				VpcIdentifier:   pulumi.Any(aws_vpc.Example.Id),
				Port:            pulumi.Int(443),
				Protocol:        pulumi.String("HTTPS"),
				ProtocolVersion: pulumi.String("HTTP1"),
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
import com.pulumi.aws.vpclattice.TargetGroup;
import com.pulumi.aws.vpclattice.TargetGroupArgs;
import com.pulumi.aws.vpclattice.inputs.TargetGroupConfigArgs;
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
        var example = new TargetGroup("example", TargetGroupArgs.builder()        
            .type("ALB")
            .config(TargetGroupConfigArgs.builder()
                .vpcIdentifier(aws_vpc.example().id())
                .port(443)
                .protocol("HTTPS")
                .protocolVersion("HTTP1")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:vpclattice:TargetGroup
    properties:
      type: ALB
      config:
        vpcIdentifier: ${aws_vpc.example.id}
        port: 443
        protocol: HTTPS
        protocolVersion: HTTP1
```
{{% /example %}}
{{% example %}}
### Lambda

If the type is Lambda, `config` block is not supported.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.vpclattice.TargetGroup("example", {type: "LAMBDA"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.vpclattice.TargetGroup("example", type="LAMBDA")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.VpcLattice.TargetGroup("example", new()
    {
        Type = "LAMBDA",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := vpclattice.NewTargetGroup(ctx, "example", &vpclattice.TargetGroupArgs{
			Type: pulumi.String("LAMBDA"),
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
import com.pulumi.aws.vpclattice.TargetGroup;
import com.pulumi.aws.vpclattice.TargetGroupArgs;
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
        var example = new TargetGroup("example", TargetGroupArgs.builder()        
            .type("LAMBDA")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:vpclattice:TargetGroup
    properties:
      type: LAMBDA
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Lattice Target Group using the `id`. For example:

```sh
 $ pulumi import aws:vpclattice/targetGroup:TargetGroup example tg-0c11d4dc16ed96bdb
```
 