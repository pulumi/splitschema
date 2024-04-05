Manages a Service Catalog Tag Option.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.servicecatalog.TagOption("example", {
    key: "nyckel",
    value: "värde",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.servicecatalog.TagOption("example",
    key="nyckel",
    value="värde")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.ServiceCatalog.TagOption("example", new()
    {
        Key = "nyckel",
        Value = "värde",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/servicecatalog"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicecatalog.NewTagOption(ctx, "example", &servicecatalog.TagOptionArgs{
			Key:   pulumi.String("nyckel"),
			Value: pulumi.String("värde"),
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
import com.pulumi.aws.servicecatalog.TagOption;
import com.pulumi.aws.servicecatalog.TagOptionArgs;
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
        var example = new TagOption("example", TagOptionArgs.builder()        
            .key("nyckel")
            .value("värde")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:servicecatalog:TagOption
    properties:
      key: nyckel
      value: värde
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_servicecatalog_tag_option` using the tag option ID. For example:

```sh
 $ pulumi import aws:servicecatalog/tagOption:TagOption example tag-pjtvagohlyo3m
```
 