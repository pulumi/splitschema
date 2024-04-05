Resource for managing an AWS FinSpace Kx Scaling Group.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.finspace.KxScalingGroup("example", {
    environmentId: aws_finspace_kx_environment.example.id,
    availabilityZoneId: "use1-az2",
    hostType: "kx.sg.4xlarge",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.finspace.KxScalingGroup("example",
    environment_id=aws_finspace_kx_environment["example"]["id"],
    availability_zone_id="use1-az2",
    host_type="kx.sg.4xlarge")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.FinSpace.KxScalingGroup("example", new()
    {
        EnvironmentId = aws_finspace_kx_environment.Example.Id,
        AvailabilityZoneId = "use1-az2",
        HostType = "kx.sg.4xlarge",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/finspace"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := finspace.NewKxScalingGroup(ctx, "example", &finspace.KxScalingGroupArgs{
			EnvironmentId:      pulumi.Any(aws_finspace_kx_environment.Example.Id),
			AvailabilityZoneId: pulumi.String("use1-az2"),
			HostType:           pulumi.String("kx.sg.4xlarge"),
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
import com.pulumi.aws.finspace.KxScalingGroup;
import com.pulumi.aws.finspace.KxScalingGroupArgs;
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
        var example = new KxScalingGroup("example", KxScalingGroupArgs.builder()        
            .environmentId(aws_finspace_kx_environment.example().id())
            .availabilityZoneId("use1-az2")
            .hostType("kx.sg.4xlarge")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:finspace:KxScalingGroup
    properties:
      environmentId: ${aws_finspace_kx_environment.example.id}
      availabilityZoneId: use1-az2
      hostType: kx.sg.4xlarge
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import an AWS FinSpace Kx Scaling Group using the `id` (environment ID and scaling group name, comma-delimited). For example:

```sh
 $ pulumi import aws:finspace/kxScalingGroup:KxScalingGroup example n3ceo7wqxoxcti5tujqwzs,my-tf-kx-scalinggroup
```
 