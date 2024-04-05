Provides a CloudWatch Logs query definition resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudwatch.QueryDefinition("example", {
    logGroupNames: [
        "/aws/logGroup1",
        "/aws/logGroup2",
    ],
    queryString: `fields @timestamp, @message
| sort @timestamp desc
| limit 25

`,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudwatch.QueryDefinition("example",
    log_group_names=[
        "/aws/logGroup1",
        "/aws/logGroup2",
    ],
    query_string="""fields @timestamp, @message
| sort @timestamp desc
| limit 25

""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudWatch.QueryDefinition("example", new()
    {
        LogGroupNames = new[]
        {
            "/aws/logGroup1",
            "/aws/logGroup2",
        },
        QueryString = @"fields @timestamp, @message
| sort @timestamp desc
| limit 25

",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewQueryDefinition(ctx, "example", &cloudwatch.QueryDefinitionArgs{
			LogGroupNames: pulumi.StringArray{
				pulumi.String("/aws/logGroup1"),
				pulumi.String("/aws/logGroup2"),
			},
			QueryString: pulumi.String("fields @timestamp, @message\n| sort @timestamp desc\n| limit 25\n\n"),
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
import com.pulumi.aws.cloudwatch.QueryDefinition;
import com.pulumi.aws.cloudwatch.QueryDefinitionArgs;
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
        var example = new QueryDefinition("example", QueryDefinitionArgs.builder()        
            .logGroupNames(            
                "/aws/logGroup1",
                "/aws/logGroup2")
            .queryString("""
fields @timestamp, @message
| sort @timestamp desc
| limit 25

            """)
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudwatch:QueryDefinition
    properties:
      logGroupNames:
        - /aws/logGroup1
        - /aws/logGroup2
      queryString: |+
        fields @timestamp, @message
        | sort @timestamp desc
        | limit 25
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch query definitions using the query definition ARN. The ARN can be found on the "Edit Query" page for the query in the AWS Console. For example:

```sh
 $ pulumi import aws:cloudwatch/queryDefinition:QueryDefinition example arn:aws:logs:us-west-2:123456789012:query-definition:269951d7-6f75-496d-9d7b-6b7a5486bdbd
```
 