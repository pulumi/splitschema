Provides an Elastic MapReduce Studio Session Mapping.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.emr.StudioSessionMapping("example", {
    studioId: aws_emr_studio.example.id,
    identityType: "USER",
    identityId: "example",
    sessionPolicyArn: aws_iam_policy.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.emr.StudioSessionMapping("example",
    studio_id=aws_emr_studio["example"]["id"],
    identity_type="USER",
    identity_id="example",
    session_policy_arn=aws_iam_policy["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Emr.StudioSessionMapping("example", new()
    {
        StudioId = aws_emr_studio.Example.Id,
        IdentityType = "USER",
        IdentityId = "example",
        SessionPolicyArn = aws_iam_policy.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.NewStudioSessionMapping(ctx, "example", &emr.StudioSessionMappingArgs{
			StudioId:         pulumi.Any(aws_emr_studio.Example.Id),
			IdentityType:     pulumi.String("USER"),
			IdentityId:       pulumi.String("example"),
			SessionPolicyArn: pulumi.Any(aws_iam_policy.Example.Arn),
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
import com.pulumi.aws.emr.StudioSessionMapping;
import com.pulumi.aws.emr.StudioSessionMappingArgs;
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
        var example = new StudioSessionMapping("example", StudioSessionMappingArgs.builder()        
            .studioId(aws_emr_studio.example().id())
            .identityType("USER")
            .identityId("example")
            .sessionPolicyArn(aws_iam_policy.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:emr:StudioSessionMapping
    properties:
      studioId: ${aws_emr_studio.example.id}
      identityType: USER
      identityId: example
      sessionPolicyArn: ${aws_iam_policy.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EMR studio session mappings using `studio-id:identity-type:identity-id`. For example:

```sh
 $ pulumi import aws:emr/studioSessionMapping:StudioSessionMapping example es-xxxxx:USER:xxxxx-xxx-xxx
```
 