Use this data source to get information about an AWS Serverless Application Repository application. For example, this can be used to determine the required `capabilities` for an application.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleApplication = aws.serverlessrepository.getApplication({
    applicationId: "arn:aws:serverlessrepo:us-east-1:123456789012:applications/ExampleApplication",
});
const exampleCloudFormationStack = new aws.serverlessrepository.CloudFormationStack("exampleCloudFormationStack", {
    applicationId: exampleApplication.then(exampleApplication => exampleApplication.applicationId),
    semanticVersion: exampleApplication.then(exampleApplication => exampleApplication.semanticVersion),
    capabilities: exampleApplication.then(exampleApplication => exampleApplication.requiredCapabilities),
});
```
```python
import pulumi
import pulumi_aws as aws

example_application = aws.serverlessrepository.get_application(application_id="arn:aws:serverlessrepo:us-east-1:123456789012:applications/ExampleApplication")
example_cloud_formation_stack = aws.serverlessrepository.CloudFormationStack("exampleCloudFormationStack",
    application_id=example_application.application_id,
    semantic_version=example_application.semantic_version,
    capabilities=example_application.required_capabilities)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleApplication = Aws.ServerlessRepository.GetApplication.Invoke(new()
    {
        ApplicationId = "arn:aws:serverlessrepo:us-east-1:123456789012:applications/ExampleApplication",
    });

    var exampleCloudFormationStack = new Aws.ServerlessRepository.CloudFormationStack("exampleCloudFormationStack", new()
    {
        ApplicationId = exampleApplication.Apply(getApplicationResult => getApplicationResult.ApplicationId),
        SemanticVersion = exampleApplication.Apply(getApplicationResult => getApplicationResult.SemanticVersion),
        Capabilities = exampleApplication.Apply(getApplicationResult => getApplicationResult.RequiredCapabilities),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/serverlessrepository"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleApplication, err := serverlessrepository.GetApplication(ctx, &serverlessrepository.GetApplicationArgs{
			ApplicationId: "arn:aws:serverlessrepo:us-east-1:123456789012:applications/ExampleApplication",
		}, nil)
		if err != nil {
			return err
		}
		_, err = serverlessrepository.NewCloudFormationStack(ctx, "exampleCloudFormationStack", &serverlessrepository.CloudFormationStackArgs{
			ApplicationId:   *pulumi.String(exampleApplication.ApplicationId),
			SemanticVersion: *pulumi.String(exampleApplication.SemanticVersion),
			Capabilities:    interface{}(exampleApplication.RequiredCapabilities),
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
import com.pulumi.aws.serverlessrepository.ServerlessrepositoryFunctions;
import com.pulumi.aws.serverlessrepository.inputs.GetApplicationArgs;
import com.pulumi.aws.serverlessrepository.CloudFormationStack;
import com.pulumi.aws.serverlessrepository.CloudFormationStackArgs;
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
        final var exampleApplication = ServerlessrepositoryFunctions.getApplication(GetApplicationArgs.builder()
            .applicationId("arn:aws:serverlessrepo:us-east-1:123456789012:applications/ExampleApplication")
            .build());

        var exampleCloudFormationStack = new CloudFormationStack("exampleCloudFormationStack", CloudFormationStackArgs.builder()        
            .applicationId(exampleApplication.applyValue(getApplicationResult -> getApplicationResult.applicationId()))
            .semanticVersion(exampleApplication.applyValue(getApplicationResult -> getApplicationResult.semanticVersion()))
            .capabilities(exampleApplication.applyValue(getApplicationResult -> getApplicationResult.requiredCapabilities()))
            .build());

    }
}
```
```yaml
resources:
  exampleCloudFormationStack:
    type: aws:serverlessrepository:CloudFormationStack
    properties:
      applicationId: ${exampleApplication.applicationId}
      semanticVersion: ${exampleApplication.semanticVersion}
      capabilities: ${exampleApplication.requiredCapabilities}
variables:
  exampleApplication:
    fn::invoke:
      Function: aws:serverlessrepository:getApplication
      Arguments:
        applicationId: arn:aws:serverlessrepo:us-east-1:123456789012:applications/ExampleApplication
```
{{% /example %}}
{{% /examples %}}