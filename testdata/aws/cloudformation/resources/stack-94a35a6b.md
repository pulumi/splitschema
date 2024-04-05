Provides a CloudFormation Stack resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const network = new aws.cloudformation.Stack("network", {
    parameters: {
        VPCCidr: "10.0.0.0/16",
    },
    templateBody: JSON.stringify({
        Parameters: {
            VPCCidr: {
                Type: "String",
                Default: "10.0.0.0/16",
                Description: "Enter the CIDR block for the VPC. Default is 10.0.0.0/16.",
            },
        },
        Resources: {
            myVpc: {
                Type: "AWS::EC2::VPC",
                Properties: {
                    CidrBlock: {
                        Ref: "VPCCidr",
                    },
                    Tags: [{
                        Key: "Name",
                        Value: "Primary_CF_VPC",
                    }],
                },
            },
        },
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

network = aws.cloudformation.Stack("network",
    parameters={
        "VPCCidr": "10.0.0.0/16",
    },
    template_body=json.dumps({
        "Parameters": {
            "VPCCidr": {
                "Type": "String",
                "Default": "10.0.0.0/16",
                "Description": "Enter the CIDR block for the VPC. Default is 10.0.0.0/16.",
            },
        },
        "Resources": {
            "myVpc": {
                "Type": "AWS::EC2::VPC",
                "Properties": {
                    "CidrBlock": {
                        "Ref": "VPCCidr",
                    },
                    "Tags": [{
                        "Key": "Name",
                        "Value": "Primary_CF_VPC",
                    }],
                },
            },
        },
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
    var network = new Aws.CloudFormation.Stack("network", new()
    {
        Parameters = 
        {
            { "VPCCidr", "10.0.0.0/16" },
        },
        TemplateBody = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Parameters"] = new Dictionary<string, object?>
            {
                ["VPCCidr"] = new Dictionary<string, object?>
                {
                    ["Type"] = "String",
                    ["Default"] = "10.0.0.0/16",
                    ["Description"] = "Enter the CIDR block for the VPC. Default is 10.0.0.0/16.",
                },
            },
            ["Resources"] = new Dictionary<string, object?>
            {
                ["myVpc"] = new Dictionary<string, object?>
                {
                    ["Type"] = "AWS::EC2::VPC",
                    ["Properties"] = new Dictionary<string, object?>
                    {
                        ["CidrBlock"] = new Dictionary<string, object?>
                        {
                            ["Ref"] = "VPCCidr",
                        },
                        ["Tags"] = new[]
                        {
                            new Dictionary<string, object?>
                            {
                                ["Key"] = "Name",
                                ["Value"] = "Primary_CF_VPC",
                            },
                        },
                    },
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Parameters": map[string]interface{}{
				"VPCCidr": map[string]interface{}{
					"Type":        "String",
					"Default":     "10.0.0.0/16",
					"Description": "Enter the CIDR block for the VPC. Default is 10.0.0.0/16.",
				},
			},
			"Resources": map[string]interface{}{
				"myVpc": map[string]interface{}{
					"Type": "AWS::EC2::VPC",
					"Properties": map[string]interface{}{
						"CidrBlock": map[string]interface{}{
							"Ref": "VPCCidr",
						},
						"Tags": []map[string]interface{}{
							map[string]interface{}{
								"Key":   "Name",
								"Value": "Primary_CF_VPC",
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = cloudformation.NewStack(ctx, "network", &cloudformation.StackArgs{
			Parameters: pulumi.StringMap{
				"VPCCidr": pulumi.String("10.0.0.0/16"),
			},
			TemplateBody: pulumi.String(json0),
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
import com.pulumi.aws.cloudformation.Stack;
import com.pulumi.aws.cloudformation.StackArgs;
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
        var network = new Stack("network", StackArgs.builder()        
            .parameters(Map.of("VPCCidr", "10.0.0.0/16"))
            .templateBody(serializeJson(
                jsonObject(
                    jsonProperty("Parameters", jsonObject(
                        jsonProperty("VPCCidr", jsonObject(
                            jsonProperty("Type", "String"),
                            jsonProperty("Default", "10.0.0.0/16"),
                            jsonProperty("Description", "Enter the CIDR block for the VPC. Default is 10.0.0.0/16.")
                        ))
                    )),
                    jsonProperty("Resources", jsonObject(
                        jsonProperty("myVpc", jsonObject(
                            jsonProperty("Type", "AWS::EC2::VPC"),
                            jsonProperty("Properties", jsonObject(
                                jsonProperty("CidrBlock", jsonObject(
                                    jsonProperty("Ref", "VPCCidr")
                                )),
                                jsonProperty("Tags", jsonArray(jsonObject(
                                    jsonProperty("Key", "Name"),
                                    jsonProperty("Value", "Primary_CF_VPC")
                                )))
                            ))
                        ))
                    ))
                )))
            .build());

    }
}
```
```yaml
resources:
  network:
    type: aws:cloudformation:Stack
    properties:
      parameters:
        VPCCidr: 10.0.0.0/16
      templateBody:
        fn::toJSON:
          Parameters:
            VPCCidr:
              Type: String
              Default: 10.0.0.0/16
              Description: Enter the CIDR block for the VPC. Default is 10.0.0.0/16.
          Resources:
            myVpc:
              Type: AWS::EC2::VPC
              Properties:
                CidrBlock:
                  Ref: VPCCidr
                Tags:
                  - Key: Name
                    Value: Primary_CF_VPC
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cloudformation Stacks using the `name`. For example:

```sh
 $ pulumi import aws:cloudformation/stack:Stack stack networking-stack
```
 