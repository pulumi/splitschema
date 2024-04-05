Provides an OpsWorks RDS DB Instance resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myInstance = new aws.opsworks.RdsDbInstance("myInstance", {
    stackId: aws_opsworks_stack.my_stack.id,
    rdsDbInstanceArn: aws_db_instance.my_instance.arn,
    dbUser: "someUser",
    dbPassword: "somePass",
});
```
```python
import pulumi
import pulumi_aws as aws

my_instance = aws.opsworks.RdsDbInstance("myInstance",
    stack_id=aws_opsworks_stack["my_stack"]["id"],
    rds_db_instance_arn=aws_db_instance["my_instance"]["arn"],
    db_user="someUser",
    db_password="somePass")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myInstance = new Aws.OpsWorks.RdsDbInstance("myInstance", new()
    {
        StackId = aws_opsworks_stack.My_stack.Id,
        RdsDbInstanceArn = aws_db_instance.My_instance.Arn,
        DbUser = "someUser",
        DbPassword = "somePass",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opsworks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := opsworks.NewRdsDbInstance(ctx, "myInstance", &opsworks.RdsDbInstanceArgs{
			StackId:          pulumi.Any(aws_opsworks_stack.My_stack.Id),
			RdsDbInstanceArn: pulumi.Any(aws_db_instance.My_instance.Arn),
			DbUser:           pulumi.String("someUser"),
			DbPassword:       pulumi.String("somePass"),
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
import com.pulumi.aws.opsworks.RdsDbInstance;
import com.pulumi.aws.opsworks.RdsDbInstanceArgs;
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
        var myInstance = new RdsDbInstance("myInstance", RdsDbInstanceArgs.builder()        
            .stackId(aws_opsworks_stack.my_stack().id())
            .rdsDbInstanceArn(aws_db_instance.my_instance().arn())
            .dbUser("someUser")
            .dbPassword("somePass")
            .build());

    }
}
```
```yaml
resources:
  myInstance:
    type: aws:opsworks:RdsDbInstance
    properties:
      stackId: ${aws_opsworks_stack.my_stack.id}
      rdsDbInstanceArn: ${aws_db_instance.my_instance.arn}
      dbUser: someUser
      dbPassword: somePass
```
{{% /example %}}
{{% /examples %}}