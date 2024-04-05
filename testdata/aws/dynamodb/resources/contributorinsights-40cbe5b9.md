Provides a DynamoDB contributor insights resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.dynamodb.ContributorInsights("test", {tableName: "ExampleTableName"});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.dynamodb.ContributorInsights("test", table_name="ExampleTableName")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.DynamoDB.ContributorInsights("test", new()
    {
        TableName = "ExampleTableName",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dynamodb.NewContributorInsights(ctx, "test", &dynamodb.ContributorInsightsArgs{
			TableName: pulumi.String("ExampleTableName"),
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
import com.pulumi.aws.dynamodb.ContributorInsights;
import com.pulumi.aws.dynamodb.ContributorInsightsArgs;
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
        var test = new ContributorInsights("test", ContributorInsightsArgs.builder()        
            .tableName("ExampleTableName")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:dynamodb:ContributorInsights
    properties:
      tableName: ExampleTableName
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_dynamodb_contributor_insights` using the format `name:table_name/index:index_name`, followed by the account number. For example:

```sh
 $ pulumi import aws:dynamodb/contributorInsights:ContributorInsights test name:ExampleTableName/index:ExampleIndexName/123456789012
```
 