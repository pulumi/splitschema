Use this data source to get the ARN of a topic in AWS Simple Notification
Service (SNS). By using this data source, you can reference SNS topics
without having to hard code the ARNs as input.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.sns.getTopic({
    name: "an_example_topic",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.sns.get_topic(name="an_example_topic")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Sns.GetTopic.Invoke(new()
    {
        Name = "an_example_topic",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sns.LookupTopic(ctx, &sns.LookupTopicArgs{
			Name: "an_example_topic",
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
import com.pulumi.aws.sns.SnsFunctions;
import com.pulumi.aws.sns.inputs.GetTopicArgs;
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
        final var example = SnsFunctions.getTopic(GetTopicArgs.builder()
            .name("an_example_topic")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:sns:getTopic
      Arguments:
        name: an_example_topic
```
{{% /example %}}
{{% /examples %}}