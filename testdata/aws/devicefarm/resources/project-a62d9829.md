Provides a resource to manage AWS Device Farm Projects.

For more information about Device Farm Projects, see the AWS Documentation on
[Device Farm Projects][aws-get-project].

> **NOTE:** AWS currently has limited regional support for Device Farm (e.g., `us-west-2`). See [AWS Device Farm endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/devicefarm.html) for information on supported regions.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awesomeDevices = new aws.devicefarm.Project("awesomeDevices", {});
```
```python
import pulumi
import pulumi_aws as aws

awesome_devices = aws.devicefarm.Project("awesomeDevices")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awesomeDevices = new Aws.DeviceFarm.Project("awesomeDevices");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/devicefarm"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devicefarm.NewProject(ctx, "awesomeDevices", nil)
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
import com.pulumi.aws.devicefarm.Project;
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
        var awesomeDevices = new Project("awesomeDevices");

    }
}
```
```yaml
resources:
  awesomeDevices:
    type: aws:devicefarm:Project
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DeviceFarm Projects using their ARN. For example:

```sh
 $ pulumi import aws:devicefarm/project:Project example arn:aws:devicefarm:us-west-2:123456789012:project:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
```
 