{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.cloudfront.getOriginRequestPolicy({
    name: "example-policy",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudfront.get_origin_request_policy(name="example-policy")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.CloudFront.GetOriginRequestPolicy.Invoke(new()
    {
        Name = "example-policy",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.LookupOriginRequestPolicy(ctx, &cloudfront.LookupOriginRequestPolicyArgs{
			Name: pulumi.StringRef("example-policy"),
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
import com.pulumi.aws.cloudfront.CloudfrontFunctions;
import com.pulumi.aws.cloudfront.inputs.GetOriginRequestPolicyArgs;
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
        final var example = CloudfrontFunctions.getOriginRequestPolicy(GetOriginRequestPolicyArgs.builder()
            .name("example-policy")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:cloudfront:getOriginRequestPolicy
      Arguments:
        name: example-policy
```
{{% /example %}}
{{% example %}}
### AWS-Managed Policies

AWS managed origin request policy names are prefixed with `Managed-`:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const uaReferer = aws.cloudfront.getOriginRequestPolicy({
    name: "Managed-UserAgentRefererHeaders",
});
```
```python
import pulumi
import pulumi_aws as aws

ua_referer = aws.cloudfront.get_origin_request_policy(name="Managed-UserAgentRefererHeaders")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var uaReferer = Aws.CloudFront.GetOriginRequestPolicy.Invoke(new()
    {
        Name = "Managed-UserAgentRefererHeaders",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.LookupOriginRequestPolicy(ctx, &cloudfront.LookupOriginRequestPolicyArgs{
			Name: pulumi.StringRef("Managed-UserAgentRefererHeaders"),
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
import com.pulumi.aws.cloudfront.CloudfrontFunctions;
import com.pulumi.aws.cloudfront.inputs.GetOriginRequestPolicyArgs;
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
        final var uaReferer = CloudfrontFunctions.getOriginRequestPolicy(GetOriginRequestPolicyArgs.builder()
            .name("Managed-UserAgentRefererHeaders")
            .build());

    }
}
```
```yaml
variables:
  uaReferer:
    fn::invoke:
      Function: aws:cloudfront:getOriginRequestPolicy
      Arguments:
        name: Managed-UserAgentRefererHeaders
```
{{% /example %}}
{{% /examples %}}