Get Lake Formation principals designated as data lake administrators and lists of principal permission entries for default create database and default create table permissions.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.lakeformation.getDataLakeSettings({
    catalogId: "14916253649",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lakeformation.get_data_lake_settings(catalog_id="14916253649")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.LakeFormation.GetDataLakeSettings.Invoke(new()
    {
        CatalogId = "14916253649",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lakeformation"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lakeformation.LookupDataLakeSettings(ctx, &lakeformation.LookupDataLakeSettingsArgs{
			CatalogId: pulumi.StringRef("14916253649"),
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
import com.pulumi.aws.lakeformation.LakeformationFunctions;
import com.pulumi.aws.lakeformation.inputs.GetDataLakeSettingsArgs;
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
        final var example = LakeformationFunctions.getDataLakeSettings(GetDataLakeSettingsArgs.builder()
            .catalogId("14916253649")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:lakeformation:getDataLakeSettings
      Arguments:
        catalogId: '14916253649'
```
{{% /example %}}
{{% /examples %}}