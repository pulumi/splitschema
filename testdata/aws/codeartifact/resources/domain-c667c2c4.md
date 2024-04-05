Provides a CodeArtifact Domain Resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.codeartifact.Domain("example", {domain: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.codeartifact.Domain("example", domain="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CodeArtifact.Domain("example", new()
    {
        DomainName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codeartifact"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := codeartifact.NewDomain(ctx, "example", &codeartifact.DomainArgs{
			Domain: pulumi.String("example"),
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
import com.pulumi.aws.codeartifact.Domain;
import com.pulumi.aws.codeartifact.DomainArgs;
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
        var example = new Domain("example", DomainArgs.builder()        
            .domain("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:codeartifact:Domain
    properties:
      domain: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CodeArtifact Domain using the CodeArtifact Domain arn. For example:

```sh
 $ pulumi import aws:codeartifact/domain:Domain example arn:aws:codeartifact:us-west-2:012345678912:domain/tf-acc-test-8593714120730241305
```
 