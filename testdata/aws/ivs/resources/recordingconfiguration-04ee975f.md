Resource for managing an AWS IVS (Interactive Video) Recording Configuration.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.ivs.RecordingConfiguration("example", {destinationConfiguration: {
    s3: {
        bucketName: "ivs-stream-archive",
    },
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ivs.RecordingConfiguration("example", destination_configuration=aws.ivs.RecordingConfigurationDestinationConfigurationArgs(
    s3=aws.ivs.RecordingConfigurationDestinationConfigurationS3Args(
        bucket_name="ivs-stream-archive",
    ),
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Ivs.RecordingConfiguration("example", new()
    {
        DestinationConfiguration = new Aws.Ivs.Inputs.RecordingConfigurationDestinationConfigurationArgs
        {
            S3 = new Aws.Ivs.Inputs.RecordingConfigurationDestinationConfigurationS3Args
            {
                BucketName = "ivs-stream-archive",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ivs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ivs.NewRecordingConfiguration(ctx, "example", &ivs.RecordingConfigurationArgs{
			DestinationConfiguration: &ivs.RecordingConfigurationDestinationConfigurationArgs{
				S3: &ivs.RecordingConfigurationDestinationConfigurationS3Args{
					BucketName: pulumi.String("ivs-stream-archive"),
				},
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
import com.pulumi.aws.ivs.RecordingConfiguration;
import com.pulumi.aws.ivs.RecordingConfigurationArgs;
import com.pulumi.aws.ivs.inputs.RecordingConfigurationDestinationConfigurationArgs;
import com.pulumi.aws.ivs.inputs.RecordingConfigurationDestinationConfigurationS3Args;
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
        var example = new RecordingConfiguration("example", RecordingConfigurationArgs.builder()        
            .destinationConfiguration(RecordingConfigurationDestinationConfigurationArgs.builder()
                .s3(RecordingConfigurationDestinationConfigurationS3Args.builder()
                    .bucketName("ivs-stream-archive")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:ivs:RecordingConfiguration
    properties:
      destinationConfiguration:
        s3:
          bucketName: ivs-stream-archive
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IVS (Interactive Video) Recording Configuration using the ARN. For example:

```sh
 $ pulumi import aws:ivs/recordingConfiguration:RecordingConfiguration example arn:aws:ivs:us-west-2:326937407773:recording-configuration/KAk1sHBl2L47
```
 