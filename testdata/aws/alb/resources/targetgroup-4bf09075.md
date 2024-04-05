Provides a Target Group resource for use with Load Balancer resources.

> **Note:** `aws.alb.TargetGroup` is known as `aws.lb.TargetGroup`. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}
### Instance Target Group

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
const test = new aws.lb.TargetGroup("test", {
    port: 80,
    protocol: "HTTP",
    vpcId: main.id,
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
test = aws.lb.TargetGroup("test",
    port=80,
    protocol="HTTP",
    vpc_id=main.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var test = new Aws.LB.TargetGroup("test", new()
    {
        Port = 80,
        Protocol = "HTTP",
        VpcId = main.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = lb.NewTargetGroup(ctx, "test", &lb.TargetGroupArgs{
			Port:     pulumi.Int(80),
			Protocol: pulumi.String("HTTP"),
			VpcId:    main.ID(),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.lb.TargetGroup;
import com.pulumi.aws.lb.TargetGroupArgs;
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
        var main = new Vpc("main", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var test = new TargetGroup("test", TargetGroupArgs.builder()        
            .port(80)
            .protocol("HTTP")
            .vpcId(main.id())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lb:TargetGroup
    properties:
      port: 80
      protocol: HTTP
      vpcId: ${main.id}
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
```
{{% /example %}}
{{% example %}}
### IP Target Group

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.ec2.Vpc("main", {cidrBlock: "10.0.0.0/16"});
const ip_example = new aws.lb.TargetGroup("ip-example", {
    port: 80,
    protocol: "HTTP",
    targetType: "ip",
    vpcId: main.id,
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.ec2.Vpc("main", cidr_block="10.0.0.0/16")
ip_example = aws.lb.TargetGroup("ip-example",
    port=80,
    protocol="HTTP",
    target_type="ip",
    vpc_id=main.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.Ec2.Vpc("main", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var ip_example = new Aws.LB.TargetGroup("ip-example", new()
    {
        Port = 80,
        Protocol = "HTTP",
        TargetType = "ip",
        VpcId = main.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		main, err := ec2.NewVpc(ctx, "main", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		_, err = lb.NewTargetGroup(ctx, "ip-example", &lb.TargetGroupArgs{
			Port:       pulumi.Int(80),
			Protocol:   pulumi.String("HTTP"),
			TargetType: pulumi.String("ip"),
			VpcId:      main.ID(),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.lb.TargetGroup;
import com.pulumi.aws.lb.TargetGroupArgs;
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
        var main = new Vpc("main", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var ip_example = new TargetGroup("ip-example", TargetGroupArgs.builder()        
            .port(80)
            .protocol("HTTP")
            .targetType("ip")
            .vpcId(main.id())
            .build());

    }
}
```
```yaml
resources:
  ip-example:
    type: aws:lb:TargetGroup
    properties:
      port: 80
      protocol: HTTP
      targetType: ip
      vpcId: ${main.id}
  main:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
```
{{% /example %}}
{{% example %}}
### Lambda Target Group

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const lambda_example = new aws.lb.TargetGroup("lambda-example", {targetType: "lambda"});
```
```python
import pulumi
import pulumi_aws as aws

lambda_example = aws.lb.TargetGroup("lambda-example", target_type="lambda")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var lambda_example = new Aws.LB.TargetGroup("lambda-example", new()
    {
        TargetType = "lambda",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewTargetGroup(ctx, "lambda-example", &lb.TargetGroupArgs{
			TargetType: pulumi.String("lambda"),
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
import com.pulumi.aws.lb.TargetGroup;
import com.pulumi.aws.lb.TargetGroupArgs;
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
        var lambda_example = new TargetGroup("lambda-example", TargetGroupArgs.builder()        
            .targetType("lambda")
            .build());

    }
}
```
```yaml
resources:
  lambda-example:
    type: aws:lb:TargetGroup
    properties:
      targetType: lambda
```
{{% /example %}}
{{% example %}}
### ALB Target Group

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const alb_example = new aws.lb.TargetGroup("alb-example", {
    targetType: "alb",
    port: 80,
    protocol: "TCP",
    vpcId: aws_vpc.main.id,
});
```
```python
import pulumi
import pulumi_aws as aws

alb_example = aws.lb.TargetGroup("alb-example",
    target_type="alb",
    port=80,
    protocol="TCP",
    vpc_id=aws_vpc["main"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var alb_example = new Aws.LB.TargetGroup("alb-example", new()
    {
        TargetType = "alb",
        Port = 80,
        Protocol = "TCP",
        VpcId = aws_vpc.Main.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewTargetGroup(ctx, "alb-example", &lb.TargetGroupArgs{
			TargetType: pulumi.String("alb"),
			Port:       pulumi.Int(80),
			Protocol:   pulumi.String("TCP"),
			VpcId:      pulumi.Any(aws_vpc.Main.Id),
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
import com.pulumi.aws.lb.TargetGroup;
import com.pulumi.aws.lb.TargetGroupArgs;
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
        var alb_example = new TargetGroup("alb-example", TargetGroupArgs.builder()        
            .targetType("alb")
            .port(80)
            .protocol("TCP")
            .vpcId(aws_vpc.main().id())
            .build());

    }
}
```
```yaml
resources:
  alb-example:
    type: aws:lb:TargetGroup
    properties:
      targetType: alb
      port: 80
      protocol: TCP
      vpcId: ${aws_vpc.main.id}
```
{{% /example %}}
{{% example %}}
### Target group with unhealthy connection termination disabled

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const tcp_example = new aws.lb.TargetGroup("tcp-example", {
    port: 25,
    protocol: "TCP",
    vpcId: aws_vpc.main.id,
    targetHealthStates: [{
        enableUnhealthyConnectionTermination: false,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

tcp_example = aws.lb.TargetGroup("tcp-example",
    port=25,
    protocol="TCP",
    vpc_id=aws_vpc["main"]["id"],
    target_health_states=[aws.lb.TargetGroupTargetHealthStateArgs(
        enable_unhealthy_connection_termination=False,
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var tcp_example = new Aws.LB.TargetGroup("tcp-example", new()
    {
        Port = 25,
        Protocol = "TCP",
        VpcId = aws_vpc.Main.Id,
        TargetHealthStates = new[]
        {
            new Aws.LB.Inputs.TargetGroupTargetHealthStateArgs
            {
                EnableUnhealthyConnectionTermination = false,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewTargetGroup(ctx, "tcp-example", &lb.TargetGroupArgs{
			Port:     pulumi.Int(25),
			Protocol: pulumi.String("TCP"),
			VpcId:    pulumi.Any(aws_vpc.Main.Id),
			TargetHealthStates: lb.TargetGroupTargetHealthStateArray{
				&lb.TargetGroupTargetHealthStateArgs{
					EnableUnhealthyConnectionTermination: pulumi.Bool(false),
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
import com.pulumi.aws.lb.TargetGroup;
import com.pulumi.aws.lb.TargetGroupArgs;
import com.pulumi.aws.lb.inputs.TargetGroupTargetHealthStateArgs;
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
        var tcp_example = new TargetGroup("tcp-example", TargetGroupArgs.builder()        
            .port(25)
            .protocol("TCP")
            .vpcId(aws_vpc.main().id())
            .targetHealthStates(TargetGroupTargetHealthStateArgs.builder()
                .enableUnhealthyConnectionTermination(false)
                .build())
            .build());

    }
}
```
```yaml
resources:
  tcp-example:
    type: aws:lb:TargetGroup
    properties:
      port: 25
      protocol: TCP
      vpcId: ${aws_vpc.main.id}
      targetHealthStates:
        - enableUnhealthyConnectionTermination: false
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Target Groups using their ARN. For example:

```sh
 $ pulumi import aws:alb/targetGroup:TargetGroup app_front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:targetgroup/app-front-end/20cfe21448b66314
```
 