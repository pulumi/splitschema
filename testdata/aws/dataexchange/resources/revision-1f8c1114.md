Provides a resource to manage AWS Data Exchange Revisions.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.dataexchange.Revision("example", {dataSetId: aws_dataexchange_data_set.example.id});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.dataexchange.Revision("example", data_set_id=aws_dataexchange_data_set["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.DataExchange.Revision("example", new()
    {
        DataSetId = aws_dataexchange_data_set.Example.Id,
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
		_, err := dataexchange.NewRevision(ctx, "example", &dataexchange.RevisionArgs{
			DataSetId: pulumi.Any(aws_dataexchange_data_set.Example.Id),
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
import com.pulumi.aws.dataexchange.Revision;
import com.pulumi.aws.dataexchange.RevisionArgs;
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
        var example = new Revision("example", RevisionArgs.builder()        
            .dataSetId(aws_dataexchange_data_set.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:dataexchange:Revision
    properties:
      dataSetId: ${aws_dataexchange_data_set.example.id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DataExchange Revisions using their `data-set-id:revision-id`. For example:

```sh
 $ pulumi import aws:dataexchange/revision:Revision example 4fa784c7-ccb4-4dbf-ba4f-02198320daa1:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
```
 