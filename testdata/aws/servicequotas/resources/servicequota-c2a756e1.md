Manages an individual Service Quota.

> **NOTE:** Global quotas apply to all AWS regions, but can only be accessed in `us-east-1` in the Commercial partition or `us-gov-west-1` in the GovCloud partition. In other regions, the AWS API will return the error `The request failed because the specified service does not exist.`

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicequotas.ServiceQuota("example", {
    quotaCode: "L-F678F1CE",
    serviceCode: "vpc",
    value: 75,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicequotas.ServiceQuota("example",
    quota_code="L-F678F1CE",
    service_code="vpc",
    value=75)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceQuotas.ServiceQuota("example", new()
    {
        QuotaCode = "L-F678F1CE",
        ServiceCode = "vpc",
        Value = 75,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicequotas"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicequotas.NewServiceQuota(ctx, "example", &servicequotas.ServiceQuotaArgs{
			QuotaCode:   pulumi.String("L-F678F1CE"),
			ServiceCode: pulumi.String("vpc"),
			Value:       pulumi.Float64(75),
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
import com.pulumi.aws.servicequotas.ServiceQuota;
import com.pulumi.aws.servicequotas.ServiceQuotaArgs;
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
        var example = new ServiceQuota("example", ServiceQuotaArgs.builder()        
            .quotaCode("L-F678F1CE")
            .serviceCode("vpc")
            .value(75)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicequotas:ServiceQuota
    properties:
      quotaCode: L-F678F1CE
      serviceCode: vpc
      value: 75
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicequotas_service_quota` using the service code and quota code, separated by a front slash (`/`). For example:

~> __NOTE:__ This resource does not require explicit import and will assume management of an existing service quota on Pulumi resource creation.

```sh
 $ pulumi import aws:servicequotas/serviceQuota:ServiceQuota example vpc/L-F678F1CE
```
 