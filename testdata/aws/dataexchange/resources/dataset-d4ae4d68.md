Provides a resource to manage AWS Data Exchange DataSets.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dataexchange.DataSet("example", {
    assetType: "S3_SNAPSHOT",
    description: "example",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.dataexchange.DataSet("example",
    asset_type="S3_SNAPSHOT",
    description="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataExchange.DataSet("example", new()
    {
        AssetType = "S3_SNAPSHOT",
        Description = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dataexchange"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dataexchange.NewDataSet(ctx, "example", &dataexchange.DataSetArgs{
			AssetType:   pulumi.String("S3_SNAPSHOT"),
			Description: pulumi.String("example"),
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
import com.pulumi.aws.dataexchange.DataSet;
import com.pulumi.aws.dataexchange.DataSetArgs;
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
        var example = new DataSet("example", DataSetArgs.builder()        
            .assetType("S3_SNAPSHOT")
            .description("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:dataexchange:DataSet
    properties:
      assetType: S3_SNAPSHOT
      description: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DataExchange DataSets using their ARN. For example:

```sh
 $ pulumi import aws:dataexchange/dataSet:DataSet example arn:aws:dataexchange:us-west-2:123456789012:data-sets/4fa784c7-ccb4-4dbf-ba4f-02198320daa1
```
 