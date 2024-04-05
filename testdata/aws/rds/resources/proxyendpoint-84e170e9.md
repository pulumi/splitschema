Provides an RDS DB proxy endpoint resource. For additional information, see the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/rds-proxy-endpoints.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.rds.ProxyEndpoint("example", {
    dbProxyName: aws_db_proxy.test.name,
    dbProxyEndpointName: "example",
    vpcSubnetIds: aws_subnet.test.map(__item => __item.id),
    targetRole: "READ_ONLY",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.rds.ProxyEndpoint("example",
    db_proxy_name=aws_db_proxy["test"]["name"],
    db_proxy_endpoint_name="example",
    vpc_subnet_ids=[__item["id"] for __item in aws_subnet["test"]],
    target_role="READ_ONLY")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Rds.ProxyEndpoint("example", new()
    {
        DbProxyName = aws_db_proxy.Test.Name,
        DbProxyEndpointName = "example",
        VpcSubnetIds = aws_subnet.Test.Select(__item => __item.Id).ToList(),
        TargetRole = "READ_ONLY",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
var splat0 []interface{}
for _, val0 := range aws_subnet.Test {
splat0 = append(splat0, val0.Id)
}
_, err := rds.NewProxyEndpoint(ctx, "example", &rds.ProxyEndpointArgs{
DbProxyName: pulumi.Any(aws_db_proxy.Test.Name),
DbProxyEndpointName: pulumi.String("example"),
VpcSubnetIds: toPulumiArray(splat0),
TargetRole: pulumi.String("READ_ONLY"),
})
if err != nil {
return err
}
return nil
})
}
func toPulumiArray(arr []) pulumi.Array {
var pulumiArr pulumi.Array
for _, v := range arr {
pulumiArr = append(pulumiArr, pulumi.(v))
}
return pulumiArr
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.rds.ProxyEndpoint;
import com.pulumi.aws.rds.ProxyEndpointArgs;
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
        var example = new ProxyEndpoint("example", ProxyEndpointArgs.builder()        
            .dbProxyName(aws_db_proxy.test().name())
            .dbProxyEndpointName("example")
            .vpcSubnetIds(aws_subnet.test().stream().map(element -> element.id()).collect(toList()))
            .targetRole("READ_ONLY")
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DB proxy endpoints using the `DB-PROXY-NAME/DB-PROXY-ENDPOINT-NAME`. For example:

```sh
 $ pulumi import aws:rds/proxyEndpoint:ProxyEndpoint example example/example
```
 