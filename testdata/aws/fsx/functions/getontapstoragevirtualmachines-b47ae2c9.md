This resource can be useful for getting back a set of FSx ONTAP Storage Virtual Machine (SVM) IDs.

{{% examples %}}
## Example Usage
{{% example %}}

The following shows outputting all SVM IDs for a given FSx ONTAP File System.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.fsx.getOntapStorageVirtualMachines({
    filters: [{
        name: "file-system-id",
        values: ["fs-12345678"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.fsx.get_ontap_storage_virtual_machines(filters=[aws.fsx.GetOntapStorageVirtualMachinesFilterArgs(
    name="file-system-id",
    values=["fs-12345678"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Fsx.GetOntapStorageVirtualMachines.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Fsx.Inputs.GetOntapStorageVirtualMachinesFilterInputArgs
            {
                Name = "file-system-id",
                Values = new[]
                {
                    "fs-12345678",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := fsx.GetOntapStorageVirtualMachines(ctx, &fsx.GetOntapStorageVirtualMachinesArgs{
			Filters: []fsx.GetOntapStorageVirtualMachinesFilter{
				{
					Name: "file-system-id",
					Values: []string{
						"fs-12345678",
					},
				},
			},
		}, nil)
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
import com.pulumi.aws.fsx.FsxFunctions;
import com.pulumi.aws.fsx.inputs.GetOntapStorageVirtualMachinesArgs;
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
        final var example = FsxFunctions.getOntapStorageVirtualMachines(GetOntapStorageVirtualMachinesArgs.builder()
            .filters(GetOntapStorageVirtualMachinesFilterArgs.builder()
                .name("file-system-id")
                .values("fs-12345678")
                .build())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:fsx:getOntapStorageVirtualMachines
      Arguments:
        filters:
          - name: file-system-id
            values:
              - fs-12345678
```
{{% /example %}}
{{% /examples %}}