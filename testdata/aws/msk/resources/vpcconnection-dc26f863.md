Resource for managing an AWS Managed Streaming for Kafka VPC Connection.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.msk.VpcConnection("test", {
    authentication: "SASL_IAM",
    targetClusterArn: "aws_msk_cluster.arn",
    vpcId: aws_vpc.test.id,
    clientSubnets: aws_subnet.test.map(__item => __item.id),
    securityGroups: [aws_security_group.test.id],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.msk.VpcConnection("test",
    authentication="SASL_IAM",
    target_cluster_arn="aws_msk_cluster.arn",
    vpc_id=aws_vpc["test"]["id"],
    client_subnets=[__item["id"] for __item in aws_subnet["test"]],
    security_groups=[aws_security_group["test"]["id"]])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Msk.VpcConnection("test", new()
    {
        Authentication = "SASL_IAM",
        TargetClusterArn = "aws_msk_cluster.arn",
        VpcId = aws_vpc.Test.Id,
        ClientSubnets = aws_subnet.Test.Select(__item => __item.Id).ToList(),
        SecurityGroups = new[]
        {
            aws_security_group.Test.Id,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/msk"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
var splat0 []interface{}
for _, val0 := range aws_subnet.Test {
splat0 = append(splat0, val0.Id)
}
_, err := msk.NewVpcConnection(ctx, "test", &msk.VpcConnectionArgs{
Authentication: pulumi.String("SASL_IAM"),
TargetClusterArn: pulumi.String("aws_msk_cluster.arn"),
VpcId: pulumi.Any(aws_vpc.Test.Id),
ClientSubnets: toPulumiArray(splat0),
SecurityGroups: pulumi.StringArray{
aws_security_group.Test.Id,
},
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
import com.pulumi.aws.msk.VpcConnection;
import com.pulumi.aws.msk.VpcConnectionArgs;
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
        var test = new VpcConnection("test", VpcConnectionArgs.builder()        
            .authentication("SASL_IAM")
            .targetClusterArn("aws_msk_cluster.arn")
            .vpcId(aws_vpc.test().id())
            .clientSubnets(aws_subnet.test().stream().map(element -> element.id()).collect(toList()))
            .securityGroups(aws_security_group.test().id())
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MSK configurations using the configuration ARN. For example:

```sh
 $ pulumi import aws:msk/vpcConnection:VpcConnection example arn:aws:kafka:eu-west-2:123456789012:vpc-connection/123456789012/example/38173259-79cd-4ee8-87f3-682ea6023f48-2
```
 