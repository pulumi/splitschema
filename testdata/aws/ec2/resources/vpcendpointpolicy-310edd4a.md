Provides a VPC Endpoint Policy resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpcEndpointService = aws.ec2.getVpcEndpointService({
    service: "dynamodb",
});
const exampleVpc = new aws.ec2.Vpc("exampleVpc", {cidrBlock: "10.0.0.0/16"});
const exampleVpcEndpoint = new aws.ec2.VpcEndpoint("exampleVpcEndpoint", {
    serviceName: exampleVpcEndpointService.then(exampleVpcEndpointService => exampleVpcEndpointService.serviceName),
    vpcId: exampleVpc.id,
});
const exampleVpcEndpointPolicy = new aws.ec2.VpcEndpointPolicy("exampleVpcEndpointPolicy", {
    vpcEndpointId: exampleVpcEndpoint.id,
    policy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Sid: "AllowAll",
            Effect: "Allow",
            Principal: {
                AWS: "*",
            },
            Action: ["dynamodb:*"],
            Resource: "*",
        }],
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_vpc_endpoint_service = aws.ec2.get_vpc_endpoint_service(service="dynamodb")
example_vpc = aws.ec2.Vpc("exampleVpc", cidr_block="10.0.0.0/16")
example_vpc_endpoint = aws.ec2.VpcEndpoint("exampleVpcEndpoint",
    service_name=example_vpc_endpoint_service.service_name,
    vpc_id=example_vpc.id)
example_vpc_endpoint_policy = aws.ec2.VpcEndpointPolicy("exampleVpcEndpointPolicy",
    vpc_endpoint_id=example_vpc_endpoint.id,
    policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Sid": "AllowAll",
            "Effect": "Allow",
            "Principal": {
                "AWS": "*",
            },
            "Action": ["dynamodb:*"],
            "Resource": "*",
        }],
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleVpcEndpointService = Aws.Ec2.GetVpcEndpointService.Invoke(new()
    {
        Service = "dynamodb",
    });

    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "10.0.0.0/16",
    });

    var exampleVpcEndpoint = new Aws.Ec2.VpcEndpoint("exampleVpcEndpoint", new()
    {
        ServiceName = exampleVpcEndpointService.Apply(getVpcEndpointServiceResult => getVpcEndpointServiceResult.ServiceName),
        VpcId = exampleVpc.Id,
    });

    var exampleVpcEndpointPolicy = new Aws.Ec2.VpcEndpointPolicy("exampleVpcEndpointPolicy", new()
    {
        VpcEndpointId = exampleVpcEndpoint.Id,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Sid"] = "AllowAll",
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["AWS"] = "*",
                    },
                    ["Action"] = new[]
                    {
                        "dynamodb:*",
                    },
                    ["Resource"] = "*",
                },
            },
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleVpcEndpointService, err := ec2.LookupVpcEndpointService(ctx, &ec2.LookupVpcEndpointServiceArgs{
			Service: pulumi.StringRef("dynamodb"),
		}, nil)
		if err != nil {
			return err
		}
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("10.0.0.0/16"),
		})
		if err != nil {
			return err
		}
		exampleVpcEndpoint, err := ec2.NewVpcEndpoint(ctx, "exampleVpcEndpoint", &ec2.VpcEndpointArgs{
			ServiceName: *pulumi.String(exampleVpcEndpointService.ServiceName),
			VpcId:       exampleVpc.ID(),
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Sid":    "AllowAll",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"AWS": "*",
					},
					"Action": []string{
						"dynamodb:*",
					},
					"Resource": "*",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = ec2.NewVpcEndpointPolicy(ctx, "exampleVpcEndpointPolicy", &ec2.VpcEndpointPolicyArgs{
			VpcEndpointId: exampleVpcEndpoint.ID(),
			Policy:        pulumi.String(json0),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetVpcEndpointServiceArgs;
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.VpcEndpoint;
import com.pulumi.aws.ec2.VpcEndpointArgs;
import com.pulumi.aws.ec2.VpcEndpointPolicy;
import com.pulumi.aws.ec2.VpcEndpointPolicyArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        final var exampleVpcEndpointService = Ec2Functions.getVpcEndpointService(GetVpcEndpointServiceArgs.builder()
            .service("dynamodb")
            .build());

        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("10.0.0.0/16")
            .build());

        var exampleVpcEndpoint = new VpcEndpoint("exampleVpcEndpoint", VpcEndpointArgs.builder()        
            .serviceName(exampleVpcEndpointService.applyValue(getVpcEndpointServiceResult -> getVpcEndpointServiceResult.serviceName()))
            .vpcId(exampleVpc.id())
            .build());

        var exampleVpcEndpointPolicy = new VpcEndpointPolicy("exampleVpcEndpointPolicy", VpcEndpointPolicyArgs.builder()        
            .vpcEndpointId(exampleVpcEndpoint.id())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Sid", "AllowAll"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", "*")
                        )),
                        jsonProperty("Action", jsonArray("dynamodb:*")),
                        jsonProperty("Resource", "*")
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 10.0.0.0/16
  exampleVpcEndpoint:
    type: aws:ec2:VpcEndpoint
    properties:
      serviceName: ${exampleVpcEndpointService.serviceName}
      vpcId: ${exampleVpc.id}
  exampleVpcEndpointPolicy:
    type: aws:ec2:VpcEndpointPolicy
    properties:
      vpcEndpointId: ${exampleVpcEndpoint.id}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Sid: AllowAll
              Effect: Allow
              Principal:
                AWS: '*'
              Action:
                - dynamodb:*
              Resource: '*'
variables:
  exampleVpcEndpointService:
    fn::invoke:
      Function: aws:ec2:getVpcEndpointService
      Arguments:
        service: dynamodb
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Endpoint Policies using the `id`. For example:

```sh
 $ pulumi import aws:ec2/vpcEndpointPolicy:VpcEndpointPolicy example vpce-3ecf2a57
```
 