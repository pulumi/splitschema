Creates and manages an AWS IoT Thing Type.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.iot.ThingType("foo", {});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.iot.ThingType("foo")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Iot.ThingType("foo");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iot.NewThingType(ctx, "foo", nil)
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
import com.pulumi.aws.iot.ThingType;
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
        var foo = new ThingType("foo");

    }
}
```
```yaml
resources:
  foo:
    type: aws:iot:ThingType
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IOT Thing Types using the name. For example:

```sh
 $ pulumi import aws:iot/thingType:ThingType example example
```
 