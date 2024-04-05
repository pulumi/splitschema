Manages an [AWS Elasticsearch VPC Endpoint](https://docs.aws.amazon.com/elasticsearch-service/latest/APIReference/API_CreateVpcEndpoint.html). Creates an Amazon elasticsearch Service-managed VPC endpoint.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.elasticsearch.VpcEndpoint("foo", {
    domainArn: aws_elasticsearch_domain.domain_1.arn,
    vpcOptions: {
        securityGroupIds: [
            aws_security_group.test.id,
            aws_security_group.test2.id,
        ],
        subnetIds: [
            aws_subnet.test.id,
            aws_subnet.test2.id,
        ],
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.elasticsearch.VpcEndpoint("foo",
    domain_arn=aws_elasticsearch_domain["domain_1"]["arn"],
    vpc_options=aws.elasticsearch.VpcEndpointVpcOptionsArgs(
        security_group_ids=[
            aws_security_group["test"]["id"],
            aws_security_group["test2"]["id"],
        ],
        subnet_ids=[
            aws_subnet["test"]["id"],
            aws_subnet["test2"]["id"],
        ],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.ElasticSearch.VpcEndpoint("foo", new()
    {
        DomainArn = aws_elasticsearch_domain.Domain_1.Arn,
        VpcOptions = new Aws.ElasticSearch.Inputs.VpcEndpointVpcOptionsArgs
        {
            SecurityGroupIds = new[]
            {
                aws_security_group.Test.Id,
                aws_security_group.Test2.Id,
            },
            SubnetIds = new[]
            {
                aws_subnet.Test.Id,
                aws_subnet.Test2.Id,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticsearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := elasticsearch.NewVpcEndpoint(ctx, "foo", &elasticsearch.VpcEndpointArgs{
			DomainArn: pulumi.Any(aws_elasticsearch_domain.Domain_1.Arn),
			VpcOptions: &elasticsearch.VpcEndpointVpcOptionsArgs{
				SecurityGroupIds: pulumi.StringArray{
					aws_security_group.Test.Id,
					aws_security_group.Test2.Id,
				},
				SubnetIds: pulumi.StringArray{
					aws_subnet.Test.Id,
					aws_subnet.Test2.Id,
				},
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
import com.pulumi.aws.elasticsearch.VpcEndpoint;
import com.pulumi.aws.elasticsearch.VpcEndpointArgs;
import com.pulumi.aws.elasticsearch.inputs.VpcEndpointVpcOptionsArgs;
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
        var foo = new VpcEndpoint("foo", VpcEndpointArgs.builder()        
            .domainArn(aws_elasticsearch_domain.domain_1().arn())
            .vpcOptions(VpcEndpointVpcOptionsArgs.builder()
                .securityGroupIds(                
                    aws_security_group.test().id(),
                    aws_security_group.test2().id())
                .subnetIds(                
                    aws_subnet.test().id(),
                    aws_subnet.test2().id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:elasticsearch:VpcEndpoint
    properties:
      domainArn: ${aws_elasticsearch_domain.domain_1.arn}
      vpcOptions:
        securityGroupIds:
          - ${aws_security_group.test.id}
          - ${aws_security_group.test2.id}
        subnetIds:
          - ${aws_subnet.test.id}
          - ${aws_subnet.test2.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import elasticsearch VPC endpoint connections using the `id`. For example:

```sh
 $ pulumi import aws:elasticsearch/vpcEndpoint:VpcEndpoint example endpoint-id
```
 