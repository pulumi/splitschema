Provides a resource to manage AWS Device Farm Uploads.

> **NOTE:** AWS currently has limited regional support for Device Farm (e.g., `us-west-2`). See [AWS Device Farm endpoints and quotas](https://docs.aws.amazon.com/general/latest/gr/devicefarm.html) for information on supported regions.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleProject = new aws.devicefarm.Project("exampleProject", {});
const exampleUpload = new aws.devicefarm.Upload("exampleUpload", {
    projectArn: exampleProject.arn,
    type: "APPIUM_JAVA_TESTNG_TEST_SPEC",
});
```
```python
import pulumi
import pulumi_aws as aws

example_project = aws.devicefarm.Project("exampleProject")
example_upload = aws.devicefarm.Upload("exampleUpload",
    project_arn=example_project.arn,
    type="APPIUM_JAVA_TESTNG_TEST_SPEC")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleProject = new Aws.DeviceFarm.Project("exampleProject");

    var exampleUpload = new Aws.DeviceFarm.Upload("exampleUpload", new()
    {
        ProjectArn = exampleProject.Arn,
        Type = "APPIUM_JAVA_TESTNG_TEST_SPEC",
    });

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
		exampleProject, err := devicefarm.NewProject(ctx, "exampleProject", nil)
		if err != nil {
			return err
		}
		_, err = devicefarm.NewUpload(ctx, "exampleUpload", &devicefarm.UploadArgs{
			ProjectArn: exampleProject.Arn,
			Type:       pulumi.String("APPIUM_JAVA_TESTNG_TEST_SPEC"),
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
import com.pulumi.aws.devicefarm.Project;
import com.pulumi.aws.devicefarm.Upload;
import com.pulumi.aws.devicefarm.UploadArgs;
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
        var exampleProject = new Project("exampleProject");

        var exampleUpload = new Upload("exampleUpload", UploadArgs.builder()        
            .projectArn(exampleProject.arn())
            .type("APPIUM_JAVA_TESTNG_TEST_SPEC")
            .build());

    }
}
```
```yaml
resources:
  exampleProject:
    type: aws:devicefarm:Project
  exampleUpload:
    type: aws:devicefarm:Upload
    properties:
      projectArn: ${exampleProject.arn}
      type: APPIUM_JAVA_TESTNG_TEST_SPEC
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import DeviceFarm Uploads using their ARN. For example:

```sh
 $ pulumi import aws:devicefarm/upload:Upload example arn:aws:devicefarm:us-west-2:123456789012:upload:4fa784c7-ccb4-4dbf-ba4f-02198320daa1
```
 