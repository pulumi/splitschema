> **NOTE:** The AWS Region specified by a provider must always be one of the Regions specified for the replication set.

Use this data source to manage a replication set in AWS Systems Manager Incident Manager.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ssmincidents.getReplicationSet({});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ssmincidents.get_replication_set()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.SsmIncidents.GetReplicationSet.Invoke();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssmincidents.LookupReplicationSet(ctx, nil, nil)
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
import com.pulumi.aws.ssmincidents.SsmincidentsFunctions;
import com.pulumi.aws.ssmincidents.inputs.GetReplicationSetArgs;
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
        final var example = SsmincidentsFunctions.getReplicationSet();

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:ssmincidents:getReplicationSet
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}