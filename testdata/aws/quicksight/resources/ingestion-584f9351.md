Resource for managing an AWS QuickSight Ingestion.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.quicksight.Ingestion("example", {
    dataSetId: aws_quicksight_data_set.example.data_set_id,
    ingestionId: "example-id",
    ingestionType: "FULL_REFRESH",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.quicksight.Ingestion("example",
    data_set_id=aws_quicksight_data_set["example"]["data_set_id"],
    ingestion_id="example-id",
    ingestion_type="FULL_REFRESH")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Quicksight.Ingestion("example", new()
    {
        DataSetId = aws_quicksight_data_set.Example.Data_set_id,
        IngestionId = "example-id",
        IngestionType = "FULL_REFRESH",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/quicksight"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := quicksight.NewIngestion(ctx, "example", &quicksight.IngestionArgs{
			DataSetId:     pulumi.Any(aws_quicksight_data_set.Example.Data_set_id),
			IngestionId:   pulumi.String("example-id"),
			IngestionType: pulumi.String("FULL_REFRESH"),
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
import com.pulumi.aws.quicksight.Ingestion;
import com.pulumi.aws.quicksight.IngestionArgs;
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
        var example = new Ingestion("example", IngestionArgs.builder()        
            .dataSetId(aws_quicksight_data_set.example().data_set_id())
            .ingestionId("example-id")
            .ingestionType("FULL_REFRESH")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:quicksight:Ingestion
    properties:
      dataSetId: ${aws_quicksight_data_set.example.data_set_id}
      ingestionId: example-id
      ingestionType: FULL_REFRESH
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import QuickSight Ingestion using the AWS account ID, data set ID, and ingestion ID separated by commas (`,`). For example:

```sh
 $ pulumi import aws:quicksight/ingestion:Ingestion example 123456789012,example-dataset-id,example-ingestion-id
```
 