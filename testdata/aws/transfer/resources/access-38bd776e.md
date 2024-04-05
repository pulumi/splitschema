Provides a AWS Transfer Access resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic S3

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Access("example", {
    externalId: "S-1-1-12-1234567890-123456789-1234567890-1234",
    serverId: aws_transfer_server.example.id,
    role: aws_iam_role.example.arn,
    homeDirectory: `/${aws_s3_bucket.example.id}/`,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Access("example",
    external_id="S-1-1-12-1234567890-123456789-1234567890-1234",
    server_id=aws_transfer_server["example"]["id"],
    role=aws_iam_role["example"]["arn"],
    home_directory=f"/{aws_s3_bucket['example']['id']}/")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Access("example", new()
    {
        ExternalId = "S-1-1-12-1234567890-123456789-1234567890-1234",
        ServerId = aws_transfer_server.Example.Id,
        Role = aws_iam_role.Example.Arn,
        HomeDirectory = $"/{aws_s3_bucket.Example.Id}/",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewAccess(ctx, "example", &transfer.AccessArgs{
			ExternalId:    pulumi.String("S-1-1-12-1234567890-123456789-1234567890-1234"),
			ServerId:      pulumi.Any(aws_transfer_server.Example.Id),
			Role:          pulumi.Any(aws_iam_role.Example.Arn),
			HomeDirectory: pulumi.String(fmt.Sprintf("/%v/", aws_s3_bucket.Example.Id)),
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
import com.pulumi.aws.transfer.Access;
import com.pulumi.aws.transfer.AccessArgs;
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
        var example = new Access("example", AccessArgs.builder()        
            .externalId("S-1-1-12-1234567890-123456789-1234567890-1234")
            .serverId(aws_transfer_server.example().id())
            .role(aws_iam_role.example().arn())
            .homeDirectory(String.format("/%s/", aws_s3_bucket.example().id()))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Access
    properties:
      externalId: S-1-1-12-1234567890-123456789-1234567890-1234
      serverId: ${aws_transfer_server.example.id}
      role: ${aws_iam_role.example.arn}
      homeDirectory: /${aws_s3_bucket.example.id}/
```
{{% /example %}}
{{% example %}}
### Basic EFS

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.transfer.Access("test", {
    externalId: "S-1-1-12-1234567890-123456789-1234567890-1234",
    serverId: aws_transfer_server.test.id,
    role: aws_iam_role.test.arn,
    homeDirectory: `/${aws_efs_file_system.test.id}/`,
    posixProfile: {
        gid: 1000,
        uid: 1000,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.transfer.Access("test",
    external_id="S-1-1-12-1234567890-123456789-1234567890-1234",
    server_id=aws_transfer_server["test"]["id"],
    role=aws_iam_role["test"]["arn"],
    home_directory=f"/{aws_efs_file_system['test']['id']}/",
    posix_profile=aws.transfer.AccessPosixProfileArgs(
        gid=1000,
        uid=1000,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Transfer.Access("test", new()
    {
        ExternalId = "S-1-1-12-1234567890-123456789-1234567890-1234",
        ServerId = aws_transfer_server.Test.Id,
        Role = aws_iam_role.Test.Arn,
        HomeDirectory = $"/{aws_efs_file_system.Test.Id}/",
        PosixProfile = new Aws.Transfer.Inputs.AccessPosixProfileArgs
        {
            Gid = 1000,
            Uid = 1000,
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewAccess(ctx, "test", &transfer.AccessArgs{
			ExternalId:    pulumi.String("S-1-1-12-1234567890-123456789-1234567890-1234"),
			ServerId:      pulumi.Any(aws_transfer_server.Test.Id),
			Role:          pulumi.Any(aws_iam_role.Test.Arn),
			HomeDirectory: pulumi.String(fmt.Sprintf("/%v/", aws_efs_file_system.Test.Id)),
			PosixProfile: &transfer.AccessPosixProfileArgs{
				Gid: pulumi.Int(1000),
				Uid: pulumi.Int(1000),
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
import com.pulumi.aws.transfer.Access;
import com.pulumi.aws.transfer.AccessArgs;
import com.pulumi.aws.transfer.inputs.AccessPosixProfileArgs;
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
        var test = new Access("test", AccessArgs.builder()        
            .externalId("S-1-1-12-1234567890-123456789-1234567890-1234")
            .serverId(aws_transfer_server.test().id())
            .role(aws_iam_role.test().arn())
            .homeDirectory(String.format("/%s/", aws_efs_file_system.test().id()))
            .posixProfile(AccessPosixProfileArgs.builder()
                .gid(1000)
                .uid(1000)
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:transfer:Access
    properties:
      externalId: S-1-1-12-1234567890-123456789-1234567890-1234
      serverId: ${aws_transfer_server.test.id}
      role: ${aws_iam_role.test.arn}
      homeDirectory: /${aws_efs_file_system.test.id}/
      posixProfile:
        gid: 1000
        uid: 1000
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transfer Accesses using the `server_id` and `external_id`. For example:

```sh
 $ pulumi import aws:transfer/access:Access example s-12345678/S-1-1-12-1234567890-123456789-1234567890-1234
```
 