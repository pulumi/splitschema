Provides a CodeCommit Trigger Resource.



{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testRepository = new aws.codecommit.Repository("testRepository", {repositoryName: "test"});
const testTrigger = new aws.codecommit.Trigger("testTrigger", {
    repositoryName: testRepository.repositoryName,
    triggers: [{
        name: "all",
        events: ["all"],
        destinationArn: aws_sns_topic.test.arn,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

test_repository = aws.codecommit.Repository("testRepository", repository_name="test")
test_trigger = aws.codecommit.Trigger("testTrigger",
    repository_name=test_repository.repository_name,
    triggers=[aws.codecommit.TriggerTriggerArgs(
        name="all",
        events=["all"],
        destination_arn=aws_sns_topic["test"]["arn"],
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testRepository = new Aws.CodeCommit.Repository("testRepository", new()
    {
        RepositoryName = "test",
    });

    var testTrigger = new Aws.CodeCommit.Trigger("testTrigger", new()
    {
        RepositoryName = testRepository.RepositoryName,
        Triggers = new[]
        {
            new Aws.CodeCommit.Inputs.TriggerTriggerArgs
            {
                Name = "all",
                Events = new[]
                {
                    "all",
                },
                DestinationArn = aws_sns_topic.Test.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codecommit"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testRepository, err := codecommit.NewRepository(ctx, "testRepository", &codecommit.RepositoryArgs{
			RepositoryName: pulumi.String("test"),
		})
		if err != nil {
			return err
		}
		_, err = codecommit.NewTrigger(ctx, "testTrigger", &codecommit.TriggerArgs{
			RepositoryName: testRepository.RepositoryName,
			Triggers: codecommit.TriggerTriggerArray{
				&codecommit.TriggerTriggerArgs{
					Name: pulumi.String("all"),
					Events: pulumi.StringArray{
						pulumi.String("all"),
					},
					DestinationArn: pulumi.Any(aws_sns_topic.Test.Arn),
				},
			},
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
import com.pulumi.aws.codecommit.Repository;
import com.pulumi.aws.codecommit.RepositoryArgs;
import com.pulumi.aws.codecommit.Trigger;
import com.pulumi.aws.codecommit.TriggerArgs;
import com.pulumi.aws.codecommit.inputs.TriggerTriggerArgs;
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
        var testRepository = new Repository("testRepository", RepositoryArgs.builder()        
            .repositoryName("test")
            .build());

        var testTrigger = new Trigger("testTrigger", TriggerArgs.builder()        
            .repositoryName(testRepository.repositoryName())
            .triggers(TriggerTriggerArgs.builder()
                .name("all")
                .events("all")
                .destinationArn(aws_sns_topic.test().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testRepository:
    type: aws:codecommit:Repository
    properties:
      repositoryName: test
  testTrigger:
    type: aws:codecommit:Trigger
    properties:
      repositoryName: ${testRepository.repositoryName}
      triggers:
        - name: all
          events:
            - all
          destinationArn: ${aws_sns_topic.test.arn}
```
{{% /example %}}
{{% /examples %}}