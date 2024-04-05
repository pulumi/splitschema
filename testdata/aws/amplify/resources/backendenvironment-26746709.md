Provides an Amplify Backend Environment resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleApp = new aws.amplify.App("exampleApp", {});
const exampleBackendEnvironment = new aws.amplify.BackendEnvironment("exampleBackendEnvironment", {
    appId: exampleApp.id,
    environmentName: "example",
    deploymentArtifacts: "app-example-deployment",
    stackName: "amplify-app-example",
});
```
```python
import pulumi
import pulumi_aws as aws

example_app = aws.amplify.App("exampleApp")
example_backend_environment = aws.amplify.BackendEnvironment("exampleBackendEnvironment",
    app_id=example_app.id,
    environment_name="example",
    deployment_artifacts="app-example-deployment",
    stack_name="amplify-app-example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleApp = new Aws.Amplify.App("exampleApp");

    var exampleBackendEnvironment = new Aws.Amplify.BackendEnvironment("exampleBackendEnvironment", new()
    {
        AppId = exampleApp.Id,
        EnvironmentName = "example",
        DeploymentArtifacts = "app-example-deployment",
        StackName = "amplify-app-example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/amplify"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleApp, err := amplify.NewApp(ctx, "exampleApp", nil)
		if err != nil {
			return err
		}
		_, err = amplify.NewBackendEnvironment(ctx, "exampleBackendEnvironment", &amplify.BackendEnvironmentArgs{
			AppId:               exampleApp.ID(),
			EnvironmentName:     pulumi.String("example"),
			DeploymentArtifacts: pulumi.String("app-example-deployment"),
			StackName:           pulumi.String("amplify-app-example"),
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
import com.pulumi.aws.amplify.App;
import com.pulumi.aws.amplify.BackendEnvironment;
import com.pulumi.aws.amplify.BackendEnvironmentArgs;
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
        var exampleApp = new App("exampleApp");

        var exampleBackendEnvironment = new BackendEnvironment("exampleBackendEnvironment", BackendEnvironmentArgs.builder()        
            .appId(exampleApp.id())
            .environmentName("example")
            .deploymentArtifacts("app-example-deployment")
            .stackName("amplify-app-example")
            .build());

    }
}
```
```yaml
resources:
  exampleApp:
    type: aws:amplify:App
  exampleBackendEnvironment:
    type: aws:amplify:BackendEnvironment
    properties:
      appId: ${exampleApp.id}
      environmentName: example
      deploymentArtifacts: app-example-deployment
      stackName: amplify-app-example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amplify backend environment using `app_id` and `environment_name`. For example:

```sh
 $ pulumi import aws:amplify/backendEnvironment:BackendEnvironment example d2ypk4k47z8u6/example
```
 