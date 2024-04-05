Resource for managing an AWS VPC Lattice Auth Policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleService = new aws.vpclattice.Service("exampleService", {
    authType: "AWS_IAM",
    customDomainName: "example.com",
});
const exampleAuthPolicy = new aws.vpclattice.AuthPolicy("exampleAuthPolicy", {
    resourceIdentifier: exampleService.arn,
    policy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Action: "*",
            Effect: "Allow",
            Principal: "*",
            Resource: "*",
            Condition: {
                StringNotEqualsIgnoreCase: {
                    "aws:PrincipalType": "anonymous",
                },
            },
        }],
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_service = aws.vpclattice.Service("exampleService",
    auth_type="AWS_IAM",
    custom_domain_name="example.com")
example_auth_policy = aws.vpclattice.AuthPolicy("exampleAuthPolicy",
    resource_identifier=example_service.arn,
    policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Action": "*",
            "Effect": "Allow",
            "Principal": "*",
            "Resource": "*",
            "Condition": {
                "StringNotEqualsIgnoreCase": {
                    "aws:PrincipalType": "anonymous",
                },
            },
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
    var exampleService = new Aws.VpcLattice.Service("exampleService", new()
    {
        AuthType = "AWS_IAM",
        CustomDomainName = "example.com",
    });

    var exampleAuthPolicy = new Aws.VpcLattice.AuthPolicy("exampleAuthPolicy", new()
    {
        ResourceIdentifier = exampleService.Arn,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = "*",
                    ["Effect"] = "Allow",
                    ["Principal"] = "*",
                    ["Resource"] = "*",
                    ["Condition"] = new Dictionary<string, object?>
                    {
                        ["StringNotEqualsIgnoreCase"] = new Dictionary<string, object?>
                        {
                            ["aws:PrincipalType"] = "anonymous",
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/vpclattice"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleService, err := vpclattice.NewService(ctx, "exampleService", &vpclattice.ServiceArgs{
			AuthType:         pulumi.String("AWS_IAM"),
			CustomDomainName: pulumi.String("example.com"),
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action":    "*",
					"Effect":    "Allow",
					"Principal": "*",
					"Resource":  "*",
					"Condition": map[string]interface{}{
						"StringNotEqualsIgnoreCase": map[string]interface{}{
							"aws:PrincipalType": "anonymous",
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = vpclattice.NewAuthPolicy(ctx, "exampleAuthPolicy", &vpclattice.AuthPolicyArgs{
			ResourceIdentifier: exampleService.Arn,
			Policy:             pulumi.String(json0),
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
import com.pulumi.aws.vpclattice.Service;
import com.pulumi.aws.vpclattice.ServiceArgs;
import com.pulumi.aws.vpclattice.AuthPolicy;
import com.pulumi.aws.vpclattice.AuthPolicyArgs;
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
        var exampleService = new Service("exampleService", ServiceArgs.builder()        
            .authType("AWS_IAM")
            .customDomainName("example.com")
            .build());

        var exampleAuthPolicy = new AuthPolicy("exampleAuthPolicy", AuthPolicyArgs.builder()        
            .resourceIdentifier(exampleService.arn())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "*"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", "*"),
                        jsonProperty("Resource", "*"),
                        jsonProperty("Condition", jsonObject(
                            jsonProperty("StringNotEqualsIgnoreCase", jsonObject(
                                jsonProperty("aws:PrincipalType", "anonymous")
                            ))
                        ))
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  exampleService:
    type: aws:vpclattice:Service
    properties:
      authType: AWS_IAM
      customDomainName: example.com
  exampleAuthPolicy:
    type: aws:vpclattice:AuthPolicy
    properties:
      resourceIdentifier: ${exampleService.arn}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action: '*'
              Effect: Allow
              Principal: '*'
              Resource: '*'
              Condition:
                StringNotEqualsIgnoreCase:
                  aws:PrincipalType: anonymous
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import VPC Lattice Auth Policy using the `id`. For example:

```sh
 $ pulumi import aws:vpclattice/authPolicy:AuthPolicy example abcd-12345678
```
 