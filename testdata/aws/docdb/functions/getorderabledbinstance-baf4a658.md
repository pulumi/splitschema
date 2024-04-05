Information about DocumentDB orderable DB instances.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.docdb.getOrderableDbInstance({
    engine: "docdb",
    engineVersion: "3.6.0",
    licenseModel: "na",
    preferredInstanceClasses: [
        "db.r5.large",
        "db.r4.large",
        "db.t3.medium",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.docdb.get_orderable_db_instance(engine="docdb",
    engine_version="3.6.0",
    license_model="na",
    preferred_instance_classes=[
        "db.r5.large",
        "db.r4.large",
        "db.t3.medium",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.DocDB.GetOrderableDbInstance.Invoke(new()
    {
        Engine = "docdb",
        EngineVersion = "3.6.0",
        LicenseModel = "na",
        PreferredInstanceClasses = new[]
        {
            "db.r5.large",
            "db.r4.large",
            "db.t3.medium",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/docdb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := docdb.GetOrderableDbInstance(ctx, &docdb.GetOrderableDbInstanceArgs{
			Engine:        pulumi.StringRef("docdb"),
			EngineVersion: pulumi.StringRef("3.6.0"),
			LicenseModel:  pulumi.StringRef("na"),
			PreferredInstanceClasses: []string{
				"db.r5.large",
				"db.r4.large",
				"db.t3.medium",
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
import com.pulumi.aws.docdb.DocdbFunctions;
import com.pulumi.aws.docdb.inputs.GetOrderableDbInstanceArgs;
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
        final var test = DocdbFunctions.getOrderableDbInstance(GetOrderableDbInstanceArgs.builder()
            .engine("docdb")
            .engineVersion("3.6.0")
            .licenseModel("na")
            .preferredInstanceClasses(            
                "db.r5.large",
                "db.r4.large",
                "db.t3.medium")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:docdb:getOrderableDbInstance
      Arguments:
        engine: docdb
        engineVersion: 3.6.0
        licenseModel: na
        preferredInstanceClasses:
          - db.r5.large
          - db.r4.large
          - db.t3.medium
```
{{% /example %}}
{{% /examples %}}