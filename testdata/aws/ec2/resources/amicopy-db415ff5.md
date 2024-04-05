The "AMI copy" resource allows duplication of an Amazon Machine Image (AMI),
including cross-region copies.

If the source AMI has associated EBS snapshots, those will also be duplicated
along with the AMI.

This is useful for taking a single AMI provisioned in one region and making
it available in another for a multi-region deployment.

Copying an AMI can take several minutes. The creation of this resource will
block until the new AMI is available for use on new instances.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ec2.AmiCopy("example", {
    description: "A copy of ami-xxxxxxxx",
    sourceAmiId: "ami-xxxxxxxx",
    sourceAmiRegion: "us-west-1",
    tags: {
        Name: "HelloWorld",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.AmiCopy("example",
    description="A copy of ami-xxxxxxxx",
    source_ami_id="ami-xxxxxxxx",
    source_ami_region="us-west-1",
    tags={
        "Name": "HelloWorld",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ec2.AmiCopy("example", new()
    {
        Description = "A copy of ami-xxxxxxxx",
        SourceAmiId = "ami-xxxxxxxx",
        SourceAmiRegion = "us-west-1",
        Tags = 
        {
            { "Name", "HelloWorld" },
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
		_, err := ec2.NewAmiCopy(ctx, "example", &ec2.AmiCopyArgs{
			Description:     pulumi.String("A copy of ami-xxxxxxxx"),
			SourceAmiId:     pulumi.String("ami-xxxxxxxx"),
			SourceAmiRegion: pulumi.String("us-west-1"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
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
import com.pulumi.aws.ec2.AmiCopy;
import com.pulumi.aws.ec2.AmiCopyArgs;
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
        var example = new AmiCopy("example", AmiCopyArgs.builder()        
            .description("A copy of ami-xxxxxxxx")
            .sourceAmiId("ami-xxxxxxxx")
            .sourceAmiRegion("us-west-1")
            .tags(Map.of("Name", "HelloWorld"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ec2:AmiCopy
    properties:
      description: A copy of ami-xxxxxxxx
      sourceAmiId: ami-xxxxxxxx
      sourceAmiRegion: us-west-1
      tags:
        Name: HelloWorld
```
{{% /example %}}
{{% /examples %}}