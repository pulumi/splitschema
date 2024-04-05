Use this data source to get the ID of a registered AMI for use in other
resources.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getAmi({
    executableUsers: ["self"],
    filters: [
        {
            name: "name",
            values: ["myami-*"],
        },
        {
            name: "root-device-type",
            values: ["ebs"],
        },
        {
            name: "virtualization-type",
            values: ["hvm"],
        },
    ],
    mostRecent: true,
    nameRegex: "^myami-\\d{3}",
    owners: ["self"],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_ami(executable_users=["self"],
    filters=[
        aws.ec2.GetAmiFilterArgs(
            name="name",
            values=["myami-*"],
        ),
        aws.ec2.GetAmiFilterArgs(
            name="root-device-type",
            values=["ebs"],
        ),
        aws.ec2.GetAmiFilterArgs(
            name="virtualization-type",
            values=["hvm"],
        ),
    ],
    most_recent=True,
    name_regex="^myami-\\d{3}",
    owners=["self"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetAmi.Invoke(new()
    {
        ExecutableUsers = new[]
        {
            "self",
        },
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "myami-*",
                },
            },
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "root-device-type",
                Values = new[]
                {
                    "ebs",
                },
            },
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "virtualization-type",
                Values = new[]
                {
                    "hvm",
                },
            },
        },
        MostRecent = true,
        NameRegex = "^myami-\\d{3}",
        Owners = new[]
        {
            "self",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			ExecutableUsers: []string{
				"self",
			},
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"myami-*",
					},
				},
				{
					Name: "root-device-type",
					Values: []string{
						"ebs",
					},
				},
				{
					Name: "virtualization-type",
					Values: []string{
						"hvm",
					},
				},
			},
			MostRecent: pulumi.BoolRef(true),
			NameRegex:  pulumi.StringRef("^myami-\\d{3}"),
			Owners: []string{
				"self",
			},
		}, nil)
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
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
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
        final var example = Ec2Functions.getAmi(GetAmiArgs.builder()
            .executableUsers("self")
            .filters(            
                GetAmiFilterArgs.builder()
                    .name("name")
                    .values("myami-*")
                    .build(),
                GetAmiFilterArgs.builder()
                    .name("root-device-type")
                    .values("ebs")
                    .build(),
                GetAmiFilterArgs.builder()
                    .name("virtualization-type")
                    .values("hvm")
                    .build())
            .mostRecent(true)
            .nameRegex("^myami-\\d{3}")
            .owners("self")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        executableUsers:
          - self
        filters:
          - name: name
            values:
              - myami-*
          - name: root-device-type
            values:
              - ebs
          - name: virtualization-type
            values:
              - hvm
        mostRecent: true
        nameRegex: ^myami-\d{3}
        owners:
          - self
```
{{% /example %}}
{{% /examples %}}