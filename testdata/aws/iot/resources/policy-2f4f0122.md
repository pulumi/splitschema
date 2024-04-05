Provides an IoT policy.

> **NOTE on policy versions:** Updating this resource creates a new, default policy version. If updating the resource would exceed the maximum number of versions (5), the oldest non-default version of the policy is deleted before the new policy version is created.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pubsub = new aws.iot.Policy("pubsub", {policy: JSON.stringify({
    Version: "2012-10-17",
    Statement: [{
        Action: ["iot:*"],
        Effect: "Allow",
        Resource: "*",
    }],
})});
```
```python
import pulumi
import json
import pulumi_aws as aws

pubsub = aws.iot.Policy("pubsub", policy=json.dumps({
    "Version": "2012-10-17",
    "Statement": [{
        "Action": ["iot:*"],
        "Effect": "Allow",
        "Resource": "*",
    }],
}))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var pubsub = new Aws.Iot.Policy("pubsub", new()
    {
        PolicyDocument = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[]
                    {
                        "iot:*",
                    },
                    ["Effect"] = "Allow",
                    ["Resource"] = "*",
                },
            },
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": []string{
						"iot:*",
					},
					"Effect":   "Allow",
					"Resource": "*",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = iot.NewPolicy(ctx, "pubsub", &iot.PolicyArgs{
			Policy: pulumi.String(json0),
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
import com.pulumi.aws.iot.Policy;
import com.pulumi.aws.iot.PolicyArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var pubsub = new Policy("pubsub", PolicyArgs.builder()        
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", jsonArray("iot:*")),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Resource", "*")
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  pubsub:
    type: aws:iot:Policy
    properties:
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action:
                - iot:*
              Effect: Allow
              Resource: '*'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IoT policies using the `name`. For example:

```sh
 $ pulumi import aws:iot/policy:Policy pubsub PubSubToAnyTopic
```
 