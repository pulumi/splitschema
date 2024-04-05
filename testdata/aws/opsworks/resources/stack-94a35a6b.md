Provides an OpsWorks stack resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const main = new aws.opsworks.Stack("main", {
    region: "us-west-1",
    serviceRoleArn: aws_iam_role.opsworks.arn,
    defaultInstanceProfileArn: aws_iam_instance_profile.opsworks.arn,
    tags: {
        Name: "foobar-stack",
    },
    customJson: `{
 "foobar": {
    "version": "1.0.0"
  }
}
`,
});
```
```python
import pulumi
import pulumi_aws as aws

main = aws.opsworks.Stack("main",
    region="us-west-1",
    service_role_arn=aws_iam_role["opsworks"]["arn"],
    default_instance_profile_arn=aws_iam_instance_profile["opsworks"]["arn"],
    tags={
        "Name": "foobar-stack",
    },
    custom_json="""{
 "foobar": {
    "version": "1.0.0"
  }
}
""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var main = new Aws.OpsWorks.Stack("main", new()
    {
        Region = "us-west-1",
        ServiceRoleArn = aws_iam_role.Opsworks.Arn,
        DefaultInstanceProfileArn = aws_iam_instance_profile.Opsworks.Arn,
        Tags = 
        {
            { "Name", "foobar-stack" },
        },
        CustomJson = @"{
 ""foobar"": {
    ""version"": ""1.0.0""
  }
}
",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewStack(ctx, "main", &opsworks.StackArgs{
			Region:                    pulumi.String("us-west-1"),
			ServiceRoleArn:            pulumi.Any(aws_iam_role.Opsworks.Arn),
			DefaultInstanceProfileArn: pulumi.Any(aws_iam_instance_profile.Opsworks.Arn),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("foobar-stack"),
			},
			CustomJson: pulumi.String("{\n \"foobar\": {\n    \"version\": \"1.0.0\"\n  }\n}\n"),
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
import com.pulumi.aws.opsworks.Stack;
import com.pulumi.aws.opsworks.StackArgs;
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
        var main = new Stack("main", StackArgs.builder()        
            .region("us-west-1")
            .serviceRoleArn(aws_iam_role.opsworks().arn())
            .defaultInstanceProfileArn(aws_iam_instance_profile.opsworks().arn())
            .tags(Map.of("Name", "foobar-stack"))
            .customJson("""
{
 "foobar": {
    "version": "1.0.0"
  }
}
            """)
            .build());

    }
}
```
```yaml
resources:
  main:
    type: aws:opsworks:Stack
    properties:
      region: us-west-1
      serviceRoleArn: ${aws_iam_role.opsworks.arn}
      defaultInstanceProfileArn: ${aws_iam_instance_profile.opsworks.arn}
      tags:
        Name: foobar-stack
      customJson: |
        {
         "foobar": {
            "version": "1.0.0"
          }
        }
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import OpsWorks stacks using the `id`. For example:

```sh
 $ pulumi import aws:opsworks/stack:Stack bar 00000000-0000-0000-0000-000000000000
```
 