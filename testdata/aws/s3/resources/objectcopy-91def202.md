Provides a resource for copying an S3 object.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.s3.ObjectCopy("test", {
    bucket: "destination_bucket",
    grants: [{
        permissions: ["READ"],
        type: "Group",
        uri: "http://acs.amazonaws.com/groups/global/AllUsers",
    }],
    key: "destination_key",
    source: "source_bucket/source_key",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.s3.ObjectCopy("test",
    bucket="destination_bucket",
    grants=[aws.s3.ObjectCopyGrantArgs(
        permissions=["READ"],
        type="Group",
        uri="http://acs.amazonaws.com/groups/global/AllUsers",
    )],
    key="destination_key",
    source="source_bucket/source_key")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.S3.ObjectCopy("test", new()
    {
        Bucket = "destination_bucket",
        Grants = new[]
        {
            new Aws.S3.Inputs.ObjectCopyGrantArgs
            {
                Permissions = new[]
                {
                    "READ",
                },
                Type = "Group",
                Uri = "http://acs.amazonaws.com/groups/global/AllUsers",
            },
        },
        Key = "destination_key",
        Source = "source_bucket/source_key",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := s3.NewObjectCopy(ctx, "test", &s3.ObjectCopyArgs{
			Bucket: pulumi.String("destination_bucket"),
			Grants: s3.ObjectCopyGrantArray{
				&s3.ObjectCopyGrantArgs{
					Permissions: pulumi.StringArray{
						pulumi.String("READ"),
					},
					Type: pulumi.String("Group"),
					Uri:  pulumi.String("http://acs.amazonaws.com/groups/global/AllUsers"),
				},
			},
			Key:    pulumi.String("destination_key"),
			Source: pulumi.String("source_bucket/source_key"),
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
import com.pulumi.aws.s3.ObjectCopy;
import com.pulumi.aws.s3.ObjectCopyArgs;
import com.pulumi.aws.s3.inputs.ObjectCopyGrantArgs;
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
        var test = new ObjectCopy("test", ObjectCopyArgs.builder()        
            .bucket("destination_bucket")
            .grants(ObjectCopyGrantArgs.builder()
                .permissions("READ")
                .type("Group")
                .uri("http://acs.amazonaws.com/groups/global/AllUsers")
                .build())
            .key("destination_key")
            .source("source_bucket/source_key")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:s3:ObjectCopy
    properties:
      bucket: destination_bucket
      grants:
        - permissions:
            - READ
          type: Group
          uri: http://acs.amazonaws.com/groups/global/AllUsers
      key: destination_key
      source: source_bucket/source_key
```
{{% /example %}}
{{% /examples %}}