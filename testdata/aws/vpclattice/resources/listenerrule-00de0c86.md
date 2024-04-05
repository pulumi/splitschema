Resource for managing an AWS VPC Lattice Listener Rule.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.vpclattice.ListenerRule("test", {
    listenerIdentifier: aws_vpclattice_listener.example.listener_id,
    serviceIdentifier: aws_vpclattice_service.example.id,
    priority: 20,
    match: {
        httpMatch: {
            headerMatches: [{
                name: "example-header",
                caseSensitive: false,
                match: {
                    exact: "example-contains",
                },
            }],
            pathMatch: {
                caseSensitive: true,
                match: {
                    prefix: "/example-path",
                },
            },
        },
    },
    action: {
        forward: {
            targetGroups: [
                {
                    targetGroupIdentifier: aws_vpclattice_target_group.example.id,
                    weight: 1,
                },
                {
                    targetGroupIdentifier: aws_vpclattice_target_group.example2.id,
                    weight: 2,
                },
            ],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.vpclattice.ListenerRule("test",
    listener_identifier=aws_vpclattice_listener["example"]["listener_id"],
    service_identifier=aws_vpclattice_service["example"]["id"],
    priority=20,
    match=aws.vpclattice.ListenerRuleMatchArgs(
        http_match=aws.vpclattice.ListenerRuleMatchHttpMatchArgs(
            header_matches=[aws.vpclattice.ListenerRuleMatchHttpMatchHeaderMatchArgs(
                name="example-header",
                case_sensitive=False,
                match=aws.vpclattice.ListenerRuleMatchHttpMatchHeaderMatchMatchArgs(
                    exact="example-contains",
                ),
            )],
            path_match=aws.vpclattice.ListenerRuleMatchHttpMatchPathMatchArgs(
                case_sensitive=True,
                match=aws.vpclattice.ListenerRuleMatchHttpMatchPathMatchMatchArgs(
                    prefix="/example-path",
                ),
            ),
        ),
    ),
    action=aws.vpclattice.ListenerRuleActionArgs(
        forward=aws.vpclattice.ListenerRuleActionForwardArgs(
            target_groups=[
                aws.vpclattice.ListenerRuleActionForwardTargetGroupArgs(
                    target_group_identifier=aws_vpclattice_target_group["example"]["id"],
                    weight=1,
                ),
                aws.vpclattice.ListenerRuleActionForwardTargetGroupArgs(
                    target_group_identifier=aws_vpclattice_target_group["example2"]["id"],
                    weight=2,
                ),
            ],
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
    var test = new Aws.VpcLattice.ListenerRule("test", new()
    {
        ListenerIdentifier = aws_vpclattice_listener.Example.Listener_id,
        ServiceIdentifier = aws_vpclattice_service.Example.Id,
        Priority = 20,
        Match = new Aws.VpcLattice.Inputs.ListenerRuleMatchArgs
        {
            HttpMatch = new Aws.VpcLattice.Inputs.ListenerRuleMatchHttpMatchArgs
            {
                HeaderMatches = new[]
                {
                    new Aws.VpcLattice.Inputs.ListenerRuleMatchHttpMatchHeaderMatchArgs
                    {
                        Name = "example-header",
                        CaseSensitive = false,
                        Match = new Aws.VpcLattice.Inputs.ListenerRuleMatchHttpMatchHeaderMatchMatchArgs
                        {
                            Exact = "example-contains",
                        },
                    },
                },
                PathMatch = new Aws.VpcLattice.Inputs.ListenerRuleMatchHttpMatchPathMatchArgs
                {
                    CaseSensitive = true,
                    Match = new Aws.VpcLattice.Inputs.ListenerRuleMatchHttpMatchPathMatchMatchArgs
                    {
                        Prefix = "/example-path",
                    },
                },
            },
        },
        Action = new Aws.VpcLattice.Inputs.ListenerRuleActionArgs
        {
            Forward = new Aws.VpcLattice.Inputs.ListenerRuleActionForwardArgs
            {
                TargetGroups = new[]
                {
                    new Aws.VpcLattice.Inputs.ListenerRuleActionForwardTargetGroupArgs
                    {
                        TargetGroupIdentifier = aws_vpclattice_target_group.Example.Id,
                        Weight = 1,
                    },
                    new Aws.VpcLattice.Inputs.ListenerRuleActionForwardTargetGroupArgs
                    {
                        TargetGroupIdentifier = aws_vpclattice_target_group.Example2.Id,
                        Weight = 2,
                    },
                },
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
		_, err := vpclattice.NewListenerRule(ctx, "test", &vpclattice.ListenerRuleArgs{
			ListenerIdentifier: pulumi.Any(aws_vpclattice_listener.Example.Listener_id),
			ServiceIdentifier:  pulumi.Any(aws_vpclattice_service.Example.Id),
			Priority:           pulumi.Int(20),
			Match: &vpclattice.ListenerRuleMatchArgs{
				HttpMatch: &vpclattice.ListenerRuleMatchHttpMatchArgs{
					HeaderMatches: vpclattice.ListenerRuleMatchHttpMatchHeaderMatchArray{
						&vpclattice.ListenerRuleMatchHttpMatchHeaderMatchArgs{
							Name:          pulumi.String("example-header"),
							CaseSensitive: pulumi.Bool(false),
							Match: &vpclattice.ListenerRuleMatchHttpMatchHeaderMatchMatchArgs{
								Exact: pulumi.String("example-contains"),
							},
						},
					},
					PathMatch: &vpclattice.ListenerRuleMatchHttpMatchPathMatchArgs{
						CaseSensitive: pulumi.Bool(true),
						Match: &vpclattice.ListenerRuleMatchHttpMatchPathMatchMatchArgs{
							Prefix: pulumi.String("/example-path"),
						},
					},
				},
			},
			Action: &vpclattice.ListenerRuleActionArgs{
				Forward: &vpclattice.ListenerRuleActionForwardArgs{
					TargetGroups: vpclattice.ListenerRuleActionForwardTargetGroupArray{
						&vpclattice.ListenerRuleActionForwardTargetGroupArgs{
							TargetGroupIdentifier: pulumi.Any(aws_vpclattice_target_group.Example.Id),
							Weight:                pulumi.Int(1),
						},
						&vpclattice.ListenerRuleActionForwardTargetGroupArgs{
							TargetGroupIdentifier: pulumi.Any(aws_vpclattice_target_group.Example2.Id),
							Weight:                pulumi.Int(2),
						},
					},
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
import com.pulumi.aws.vpclattice.ListenerRule;
import com.pulumi.aws.vpclattice.ListenerRuleArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleMatchArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleMatchHttpMatchArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleMatchHttpMatchPathMatchArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleMatchHttpMatchPathMatchMatchArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleActionArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleActionForwardArgs;
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
        var test = new ListenerRule("test", ListenerRuleArgs.builder()        
            .listenerIdentifier(aws_vpclattice_listener.example().listener_id())
            .serviceIdentifier(aws_vpclattice_service.example().id())
            .priority(20)
            .match(ListenerRuleMatchArgs.builder()
                .httpMatch(ListenerRuleMatchHttpMatchArgs.builder()
                    .headerMatches(ListenerRuleMatchHttpMatchHeaderMatchArgs.builder()
                        .name("example-header")
                        .caseSensitive(false)
                        .match(ListenerRuleMatchHttpMatchHeaderMatchMatchArgs.builder()
                            .exact("example-contains")
                            .build())
                        .build())
                    .pathMatch(ListenerRuleMatchHttpMatchPathMatchArgs.builder()
                        .caseSensitive(true)
                        .match(ListenerRuleMatchHttpMatchPathMatchMatchArgs.builder()
                            .prefix("/example-path")
                            .build())
                        .build())
                    .build())
                .build())
            .action(ListenerRuleActionArgs.builder()
                .forward(ListenerRuleActionForwardArgs.builder()
                    .targetGroups(                    
                        ListenerRuleActionForwardTargetGroupArgs.builder()
                            .targetGroupIdentifier(aws_vpclattice_target_group.example().id())
                            .weight(1)
                            .build(),
                        ListenerRuleActionForwardTargetGroupArgs.builder()
                            .targetGroupIdentifier(aws_vpclattice_target_group.example2().id())
                            .weight(2)
                            .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:vpclattice:ListenerRule
    properties:
      listenerIdentifier: ${aws_vpclattice_listener.example.listener_id}
      serviceIdentifier: ${aws_vpclattice_service.example.id}
      priority: 20
      match:
        httpMatch:
          headerMatches:
            - name: example-header
              caseSensitive: false
              match:
                exact: example-contains
          pathMatch:
            caseSensitive: true
            match:
              prefix: /example-path
      action:
        forward:
          targetGroups:
            - targetGroupIdentifier: ${aws_vpclattice_target_group.example.id}
              weight: 1
            - targetGroupIdentifier: ${aws_vpclattice_target_group.example2.id}
              weight: 2
```
{{% /example %}}
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.vpclattice.ListenerRule("test", {
    listenerIdentifier: aws_vpclattice_listener.example.listener_id,
    serviceIdentifier: aws_vpclattice_service.example.id,
    priority: 10,
    match: {
        httpMatch: {
            pathMatch: {
                caseSensitive: false,
                match: {
                    exact: "/example-path",
                },
            },
        },
    },
    action: {
        fixedResponse: {
            statusCode: 404,
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.vpclattice.ListenerRule("test",
    listener_identifier=aws_vpclattice_listener["example"]["listener_id"],
    service_identifier=aws_vpclattice_service["example"]["id"],
    priority=10,
    match=aws.vpclattice.ListenerRuleMatchArgs(
        http_match=aws.vpclattice.ListenerRuleMatchHttpMatchArgs(
            path_match=aws.vpclattice.ListenerRuleMatchHttpMatchPathMatchArgs(
                case_sensitive=False,
                match=aws.vpclattice.ListenerRuleMatchHttpMatchPathMatchMatchArgs(
                    exact="/example-path",
                ),
            ),
        ),
    ),
    action=aws.vpclattice.ListenerRuleActionArgs(
        fixed_response=aws.vpclattice.ListenerRuleActionFixedResponseArgs(
            status_code=404,
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
    var test = new Aws.VpcLattice.ListenerRule("test", new()
    {
        ListenerIdentifier = aws_vpclattice_listener.Example.Listener_id,
        ServiceIdentifier = aws_vpclattice_service.Example.Id,
        Priority = 10,
        Match = new Aws.VpcLattice.Inputs.ListenerRuleMatchArgs
        {
            HttpMatch = new Aws.VpcLattice.Inputs.ListenerRuleMatchHttpMatchArgs
            {
                PathMatch = new Aws.VpcLattice.Inputs.ListenerRuleMatchHttpMatchPathMatchArgs
                {
                    CaseSensitive = false,
                    Match = new Aws.VpcLattice.Inputs.ListenerRuleMatchHttpMatchPathMatchMatchArgs
                    {
                        Exact = "/example-path",
                    },
                },
            },
        },
        Action = new Aws.VpcLattice.Inputs.ListenerRuleActionArgs
        {
            FixedResponse = new Aws.VpcLattice.Inputs.ListenerRuleActionFixedResponseArgs
            {
                StatusCode = 404,
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
		_, err := vpclattice.NewListenerRule(ctx, "test", &vpclattice.ListenerRuleArgs{
			ListenerIdentifier: pulumi.Any(aws_vpclattice_listener.Example.Listener_id),
			ServiceIdentifier:  pulumi.Any(aws_vpclattice_service.Example.Id),
			Priority:           pulumi.Int(10),
			Match: &vpclattice.ListenerRuleMatchArgs{
				HttpMatch: &vpclattice.ListenerRuleMatchHttpMatchArgs{
					PathMatch: &vpclattice.ListenerRuleMatchHttpMatchPathMatchArgs{
						CaseSensitive: pulumi.Bool(false),
						Match: &vpclattice.ListenerRuleMatchHttpMatchPathMatchMatchArgs{
							Exact: pulumi.String("/example-path"),
						},
					},
				},
			},
			Action: &vpclattice.ListenerRuleActionArgs{
				FixedResponse: &vpclattice.ListenerRuleActionFixedResponseArgs{
					StatusCode: pulumi.Int(404),
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
import com.pulumi.aws.vpclattice.ListenerRule;
import com.pulumi.aws.vpclattice.ListenerRuleArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleMatchArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleMatchHttpMatchArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleMatchHttpMatchPathMatchArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleMatchHttpMatchPathMatchMatchArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleActionArgs;
import com.pulumi.aws.vpclattice.inputs.ListenerRuleActionFixedResponseArgs;
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
        var test = new ListenerRule("test", ListenerRuleArgs.builder()        
            .listenerIdentifier(aws_vpclattice_listener.example().listener_id())
            .serviceIdentifier(aws_vpclattice_service.example().id())
            .priority(10)
            .match(ListenerRuleMatchArgs.builder()
                .httpMatch(ListenerRuleMatchHttpMatchArgs.builder()
                    .pathMatch(ListenerRuleMatchHttpMatchPathMatchArgs.builder()
                        .caseSensitive(false)
                        .match(ListenerRuleMatchHttpMatchPathMatchMatchArgs.builder()
                            .exact("/example-path")
                            .build())
                        .build())
                    .build())
                .build())
            .action(ListenerRuleActionArgs.builder()
                .fixedResponse(ListenerRuleActionFixedResponseArgs.builder()
                    .statusCode(404)
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:vpclattice:ListenerRule
    properties:
      listenerIdentifier: ${aws_vpclattice_listener.example.listener_id}
      serviceIdentifier: ${aws_vpclattice_service.example.id}
      priority: 10
      match:
        httpMatch:
          pathMatch:
            caseSensitive: false
            match:
              exact: /example-path
      action:
        fixedResponse:
          statusCode: 404
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Lattice Listener Rule using the `id`. For example:

```sh
 $ pulumi import aws:vpclattice/listenerRule:ListenerRule example service123/listener456/rule789
```
 