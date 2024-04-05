Provides a SageMaker Endpoint resource.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const endpoint = new aws.sagemaker.Endpoint("endpoint", {
    endpointConfigName: aws_sagemaker_endpoint_configuration.ec.name,
    tags: {
        Name: "foo",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

endpoint = aws.sagemaker.Endpoint("endpoint",
    endpoint_config_name=aws_sagemaker_endpoint_configuration["ec"]["name"],
    tags={
        "Name": "foo",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var endpoint = new Aws.Sagemaker.Endpoint("endpoint", new()
    {
        EndpointConfigName = aws_sagemaker_endpoint_configuration.Ec.Name,
        Tags = 
        {
            { "Name", "foo" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sagemaker"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sagemaker.NewEndpoint(ctx, "endpoint", &sagemaker.EndpointArgs{
			EndpointConfigName: pulumi.Any(aws_sagemaker_endpoint_configuration.Ec.Name),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("foo"),
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
import com.pulumi.aws.sagemaker.Endpoint;
import com.pulumi.aws.sagemaker.EndpointArgs;
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
        var endpoint = new Endpoint("endpoint", EndpointArgs.builder()        
            .endpointConfigName(aws_sagemaker_endpoint_configuration.ec().name())
            .tags(Map.of("Name", "foo"))
            .build());

    }
}
```
```yaml
resources:
  endpoint:
    type: aws:sagemaker:Endpoint
    properties:
      endpointConfigName: ${aws_sagemaker_endpoint_configuration.ec.name}
      tags:
        Name: foo
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import endpoints using the `name`. For example:

```sh
 $ pulumi import aws:sagemaker/endpoint:Endpoint test_endpoint my-endpoint
```
 