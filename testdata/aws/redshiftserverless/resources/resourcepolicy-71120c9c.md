Creates a new Amazon Redshift Serverless Resource Policy.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshiftserverless.ResourcePolicy("example", {
    resourceArn: aws_redshiftserverless_snapshot.example.arn,
    policy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Effect: "Allow",
            Principal: {
                AWS: ["12345678901"],
            },
            Action: ["redshift-serverless:RestoreFromSnapshot"],
            Sid: "",
        }],
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.redshiftserverless.ResourcePolicy("example",
    resource_arn=aws_redshiftserverless_snapshot["example"]["arn"],
    policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Effect": "Allow",
            "Principal": {
                "AWS": ["12345678901"],
            },
            "Action": ["redshift-serverless:RestoreFromSnapshot"],
            "Sid": "",
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
    var example = new Aws.RedshiftServerless.ResourcePolicy("example", new()
    {
        ResourceArn = aws_redshiftserverless_snapshot.Example.Arn,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["AWS"] = new[]
                        {
                            "12345678901",
                        },
                    },
                    ["Action"] = new[]
                    {
                        "redshift-serverless:RestoreFromSnapshot",
                    },
                    ["Sid"] = "",
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshiftserverless"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"AWS": []string{
							"12345678901",
						},
					},
					"Action": []string{
						"redshift-serverless:RestoreFromSnapshot",
					},
					"Sid": "",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = redshiftserverless.NewResourcePolicy(ctx, "example", &redshiftserverless.ResourcePolicyArgs{
			ResourceArn: pulumi.Any(aws_redshiftserverless_snapshot.Example.Arn),
			Policy:      pulumi.String(json0),
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
import com.pulumi.aws.redshiftserverless.ResourcePolicy;
import com.pulumi.aws.redshiftserverless.ResourcePolicyArgs;
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
        var example = new ResourcePolicy("example", ResourcePolicyArgs.builder()        
            .resourceArn(aws_redshiftserverless_snapshot.example().arn())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", jsonArray("12345678901"))
                        )),
                        jsonProperty("Action", jsonArray("redshift-serverless:RestoreFromSnapshot")),
                        jsonProperty("Sid", "")
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshiftserverless:ResourcePolicy
    properties:
      resourceArn: ${aws_redshiftserverless_snapshot.example.arn}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Effect: Allow
              Principal:
                AWS:
                  - '12345678901'
              Action:
                - redshift-serverless:RestoreFromSnapshot
              Sid:
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Serverless Resource Policies using the `resource_arn`. For example:

```sh
 $ pulumi import aws:redshiftserverless/resourcePolicy:ResourcePolicy example example
```
 