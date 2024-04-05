Resource for managing QuickSight Data Source

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _default = new aws.quicksight.DataSource("default", {
    dataSourceId: "example-id",
    parameters: {
        s3: {
            manifestFileLocation: {
                bucket: "my-bucket",
                key: "path/to/manifest.json",
            },
        },
    },
    type: "S3",
});
```
```python
import pulumi
import pulumi_aws as aws

default = aws.quicksight.DataSource("default",
    data_source_id="example-id",
    parameters=aws.quicksight.DataSourceParametersArgs(
        s3=aws.quicksight.DataSourceParametersS3Args(
            manifest_file_location=aws.quicksight.DataSourceParametersS3ManifestFileLocationArgs(
                bucket="my-bucket",
                key="path/to/manifest.json",
            ),
        ),
    ),
    type="S3")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @default = new Aws.Quicksight.DataSource("default", new()
    {
        DataSourceId = "example-id",
        Parameters = new Aws.Quicksight.Inputs.DataSourceParametersArgs
        {
            S3 = new Aws.Quicksight.Inputs.DataSourceParametersS3Args
            {
                ManifestFileLocation = new Aws.Quicksight.Inputs.DataSourceParametersS3ManifestFileLocationArgs
                {
                    Bucket = "my-bucket",
                    Key = "path/to/manifest.json",
                },
            },
        },
        Type = "S3",
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
		_, err := quicksight.NewDataSource(ctx, "default", &quicksight.DataSourceArgs{
			DataSourceId: pulumi.String("example-id"),
			Parameters: &quicksight.DataSourceParametersArgs{
				S3: &quicksight.DataSourceParametersS3Args{
					ManifestFileLocation: &quicksight.DataSourceParametersS3ManifestFileLocationArgs{
						Bucket: pulumi.String("my-bucket"),
						Key:    pulumi.String("path/to/manifest.json"),
					},
				},
			},
			Type: pulumi.String("S3"),
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
import com.pulumi.aws.quicksight.DataSource;
import com.pulumi.aws.quicksight.DataSourceArgs;
import com.pulumi.aws.quicksight.inputs.DataSourceParametersArgs;
import com.pulumi.aws.quicksight.inputs.DataSourceParametersS3Args;
import com.pulumi.aws.quicksight.inputs.DataSourceParametersS3ManifestFileLocationArgs;
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
        var default_ = new DataSource("default", DataSourceArgs.builder()        
            .dataSourceId("example-id")
            .parameters(DataSourceParametersArgs.builder()
                .s3(DataSourceParametersS3Args.builder()
                    .manifestFileLocation(DataSourceParametersS3ManifestFileLocationArgs.builder()
                        .bucket("my-bucket")
                        .key("path/to/manifest.json")
                        .build())
                    .build())
                .build())
            .type("S3")
            .build());

    }
}
```
```yaml
resources:
  default:
    type: aws:quicksight:DataSource
    properties:
      dataSourceId: example-id
      parameters:
        s3:
          manifestFileLocation:
            bucket: my-bucket
            key: path/to/manifest.json
      type: S3
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import a QuickSight data source using the AWS account ID, and data source ID separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:quicksight/dataSource:DataSource example 123456789123/my-data-source-id
```
 