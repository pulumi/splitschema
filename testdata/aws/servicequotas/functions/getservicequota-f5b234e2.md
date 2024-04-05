Retrieve information about a Service Quota.

> **NOTE:** Global quotas apply to all AWS regions, but can only be accessed in `us-east-1` in the Commercial partition or `us-gov-west-1` in the GovCloud partition. In other regions, the AWS API will return the error `The request failed because the specified service does not exist.`

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byQuotaCode = aws.servicequotas.getServiceQuota({
    quotaCode: "L-F678F1CE",
    serviceCode: "vpc",
});
const byQuotaName = aws.servicequotas.getServiceQuota({
    quotaName: "VPCs per Region",
    serviceCode: "vpc",
});
```
```python
import pulumi
import pulumi_aws as aws

by_quota_code = aws.servicequotas.get_service_quota(quota_code="L-F678F1CE",
    service_code="vpc")
by_quota_name = aws.servicequotas.get_service_quota(quota_name="VPCs per Region",
    service_code="vpc")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var byQuotaCode = Aws.ServiceQuotas.GetServiceQuota.Invoke(new()
    {
        QuotaCode = "L-F678F1CE",
        ServiceCode = "vpc",
    });

    var byQuotaName = Aws.ServiceQuotas.GetServiceQuota.Invoke(new()
    {
        QuotaName = "VPCs per Region",
        ServiceCode = "vpc",
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
		_, err := servicequotas.LookupServiceQuota(ctx, &servicequotas.LookupServiceQuotaArgs{
			QuotaCode:   pulumi.StringRef("L-F678F1CE"),
			ServiceCode: "vpc",
		}, nil)
		if err != nil {
			return err
		}
		_, err = servicequotas.LookupServiceQuota(ctx, &servicequotas.LookupServiceQuotaArgs{
			QuotaName:   pulumi.StringRef("VPCs per Region"),
			ServiceCode: "vpc",
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
import com.pulumi.aws.servicequotas.ServicequotasFunctions;
import com.pulumi.aws.servicequotas.inputs.GetServiceQuotaArgs;
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
        final var byQuotaCode = ServicequotasFunctions.getServiceQuota(GetServiceQuotaArgs.builder()
            .quotaCode("L-F678F1CE")
            .serviceCode("vpc")
            .build());

        final var byQuotaName = ServicequotasFunctions.getServiceQuota(GetServiceQuotaArgs.builder()
            .quotaName("VPCs per Region")
            .serviceCode("vpc")
            .build());

    }
}
```
```yaml
variables:
  byQuotaCode:
    fn::invoke:
      Function: aws:servicequotas:getServiceQuota
      Arguments:
        quotaCode: L-F678F1CE
        serviceCode: vpc
  byQuotaName:
    fn::invoke:
      Function: aws:servicequotas:getServiceQuota
      Arguments:
        quotaName: VPCs per Region
        serviceCode: vpc
```
{{% /example %}}
{{% /examples %}}