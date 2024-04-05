Provides an [IAM service-linked role](https://docs.aws.amazon.com/IAM/latest/UserGuide/using-service-linked-roles.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const elasticbeanstalk = new aws.iam.ServiceLinkedRole("elasticbeanstalk", {awsServiceName: "elasticbeanstalk.amazonaws.com"});
```
```python
import pulumi
import pulumi_aws as aws

elasticbeanstalk = aws.iam.ServiceLinkedRole("elasticbeanstalk", aws_service_name="elasticbeanstalk.amazonaws.com")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var elasticbeanstalk = new Aws.Iam.ServiceLinkedRole("elasticbeanstalk", new()
    {
        AwsServiceName = "elasticbeanstalk.amazonaws.com",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewServiceLinkedRole(ctx, "elasticbeanstalk", &iam.ServiceLinkedRoleArgs{
			AwsServiceName: pulumi.String("elasticbeanstalk.amazonaws.com"),
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
import com.pulumi.aws.iam.ServiceLinkedRole;
import com.pulumi.aws.iam.ServiceLinkedRoleArgs;
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
        var elasticbeanstalk = new ServiceLinkedRole("elasticbeanstalk", ServiceLinkedRoleArgs.builder()        
            .awsServiceName("elasticbeanstalk.amazonaws.com")
            .build());

    }
}
```
```yaml
resources:
  elasticbeanstalk:
    type: aws:iam:ServiceLinkedRole
    properties:
      awsServiceName: elasticbeanstalk.amazonaws.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM service-linked roles using role ARN. For example:

```sh
 $ pulumi import aws:iam/serviceLinkedRole:ServiceLinkedRole elasticbeanstalk arn:aws:iam::123456789012:role/aws-service-role/elasticbeanstalk.amazonaws.com/AWSServiceRoleForElasticBeanstalk
```
 