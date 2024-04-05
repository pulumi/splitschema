Provides a resource to manage an S3 Access Grants instance, which serves as a logical grouping for access grants.
You can have one S3 Access Grants instance per Region in your account.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3control.AccessGrantsInstance("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3control.AccessGrantsInstance("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3Control.AccessGrantsInstance("example");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3control.NewAccessGrantsInstance(ctx, "example", nil)
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
import com.pulumi.aws.s3control.AccessGrantsInstance;
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
        var example = new AccessGrantsInstance("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:s3control:AccessGrantsInstance
```
{{% /example %}}
{{% example %}}
### AWS IAM Identity Center

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.s3control.AccessGrantsInstance("example", {identityCenterArn: "arn:aws:sso:::instance/ssoins-890759e9c7bfdc1d"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.s3control.AccessGrantsInstance("example", identity_center_arn="arn:aws:sso:::instance/ssoins-890759e9c7bfdc1d")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.S3Control.AccessGrantsInstance("example", new()
    {
        IdentityCenterArn = "arn:aws:sso:::instance/ssoins-890759e9c7bfdc1d",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3control"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3control.NewAccessGrantsInstance(ctx, "example", &s3control.AccessGrantsInstanceArgs{
			IdentityCenterArn: pulumi.String("arn:aws:sso:::instance/ssoins-890759e9c7bfdc1d"),
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
import com.pulumi.aws.s3control.AccessGrantsInstance;
import com.pulumi.aws.s3control.AccessGrantsInstanceArgs;
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
        var example = new AccessGrantsInstance("example", AccessGrantsInstanceArgs.builder()        
            .identityCenterArn("arn:aws:sso:::instance/ssoins-890759e9c7bfdc1d")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:s3control:AccessGrantsInstance
    properties:
      identityCenterArn: arn:aws:sso:::instance/ssoins-890759e9c7bfdc1d
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import S3 Access Grants instances using the `account_id`. For example:

```sh
 $ pulumi import aws:s3control/accessGrantsInstance:AccessGrantsInstance example 123456789012
```
 