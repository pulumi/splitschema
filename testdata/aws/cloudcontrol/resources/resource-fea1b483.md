Manages a Cloud Control API Resource. The configuration and lifecycle handling of these resources is proxied through Cloud Control API handlers to the backend service.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudcontrol.Resource("example", {
    typeName: "AWS::ECS::Cluster",
    desiredState: JSON.stringify({
        ClusterName: "example",
        Tags: [{
            Key: "CostCenter",
            Value: "IT",
        }],
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.cloudcontrol.Resource("example",
    type_name="AWS::ECS::Cluster",
    desired_state=json.dumps({
        "ClusterName": "example",
        "Tags": [{
            "Key": "CostCenter",
            "Value": "IT",
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
    var example = new Aws.CloudControl.Resource("example", new()
    {
        TypeName = "AWS::ECS::Cluster",
        DesiredState = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["ClusterName"] = "example",
            ["Tags"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Key"] = "CostCenter",
                    ["Value"] = "IT",
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudcontrol"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"ClusterName": "example",
			"Tags": []map[string]interface{}{
				map[string]interface{}{
					"Key":   "CostCenter",
					"Value": "IT",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = cloudcontrol.NewResource(ctx, "example", &cloudcontrol.ResourceArgs{
			TypeName:     pulumi.String("AWS::ECS::Cluster"),
			DesiredState: pulumi.String(json0),
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
import com.pulumi.aws.cloudcontrol.Resource;
import com.pulumi.aws.cloudcontrol.ResourceArgs;
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
        var example = new Resource("example", ResourceArgs.builder()        
            .typeName("AWS::ECS::Cluster")
            .desiredState(serializeJson(
                jsonObject(
                    jsonProperty("ClusterName", "example"),
                    jsonProperty("Tags", jsonArray(jsonObject(
                        jsonProperty("Key", "CostCenter"),
                        jsonProperty("Value", "IT")
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudcontrol:Resource
    properties:
      typeName: AWS::ECS::Cluster
      desiredState:
        fn::toJSON:
          ClusterName: example
          Tags:
            - Key: CostCenter
              Value: IT
```
{{% /example %}}
{{% /examples %}}