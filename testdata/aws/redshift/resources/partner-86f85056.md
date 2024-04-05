Creates a new Amazon Redshift Partner Integration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.Partner("example", {
    clusterIdentifier: aws_redshift_cluster.example.id,
    accountId: "1234567910",
    databaseName: aws_redshift_cluster.example.database_name,
    partnerName: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.Partner("example",
    cluster_identifier=aws_redshift_cluster["example"]["id"],
    account_id="1234567910",
    database_name=aws_redshift_cluster["example"]["database_name"],
    partner_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedShift.Partner("example", new()
    {
        ClusterIdentifier = aws_redshift_cluster.Example.Id,
        AccountId = "1234567910",
        DatabaseName = aws_redshift_cluster.Example.Database_name,
        PartnerName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewPartner(ctx, "example", &redshift.PartnerArgs{
			ClusterIdentifier: pulumi.Any(aws_redshift_cluster.Example.Id),
			AccountId:         pulumi.String("1234567910"),
			DatabaseName:      pulumi.Any(aws_redshift_cluster.Example.Database_name),
			PartnerName:       pulumi.String("example"),
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
import com.pulumi.aws.redshift.Partner;
import com.pulumi.aws.redshift.PartnerArgs;
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
        var example = new Partner("example", PartnerArgs.builder()        
            .clusterIdentifier(aws_redshift_cluster.example().id())
            .accountId(1234567910)
            .databaseName(aws_redshift_cluster.example().database_name())
            .partnerName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:Partner
    properties:
      clusterIdentifier: ${aws_redshift_cluster.example.id}
      accountId: 1.23456791e+09
      databaseName: ${aws_redshift_cluster.example.database_name}
      partnerName: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift usage limits using the `id`. For example:

```sh
 $ pulumi import aws:redshift/partner:Partner example 01234567910:cluster-example-id:example:example
```
 