Provides an AppSync API Association.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appsync.DomainNameApiAssociation("example", {
    apiId: aws_appsync_graphql_api.example.id,
    domainName: aws_appsync_domain_name.example.domain_name,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.appsync.DomainNameApiAssociation("example",
    api_id=aws_appsync_graphql_api["example"]["id"],
    domain_name=aws_appsync_domain_name["example"]["domain_name"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AppSync.DomainNameApiAssociation("example", new()
    {
        ApiId = aws_appsync_graphql_api.Example.Id,
        DomainName = aws_appsync_domain_name.Example.Domain_name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appsync"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appsync.NewDomainNameApiAssociation(ctx, "example", &appsync.DomainNameApiAssociationArgs{
			ApiId:      pulumi.Any(aws_appsync_graphql_api.Example.Id),
			DomainName: pulumi.Any(aws_appsync_domain_name.Example.Domain_name),
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
import com.pulumi.aws.appsync.DomainNameApiAssociation;
import com.pulumi.aws.appsync.DomainNameApiAssociationArgs;
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
        var example = new DomainNameApiAssociation("example", DomainNameApiAssociationArgs.builder()        
            .apiId(aws_appsync_graphql_api.example().id())
            .domainName(aws_appsync_domain_name.example().domain_name())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appsync:DomainNameApiAssociation
    properties:
      apiId: ${aws_appsync_graphql_api.example.id}
      domainName: ${aws_appsync_domain_name.example.domain_name}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_appsync_domain_name_api_association` using the AppSync domain name. For example:

```sh
 $ pulumi import aws:appsync/domainNameApiAssociation:DomainNameApiAssociation example example.com
```
 