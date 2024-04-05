Resource for managing an AWS CloudWatch Observability Access Manager Sink Policy.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleSink = new aws.oam.Sink("exampleSink", {});
const exampleSinkPolicy = new aws.oam.SinkPolicy("exampleSinkPolicy", {
    sinkIdentifier: exampleSink.id,
    policy: JSON.stringify({
        Version: "2012-10-17",
        Statement: [{
            Action: [
                "oam:CreateLink",
                "oam:UpdateLink",
            ],
            Effect: "Allow",
            Resource: "*",
            Principal: {
                AWS: [
                    "1111111111111",
                    "222222222222",
                ],
            },
            Condition: {
                "ForAllValues:StringEquals": {
                    "oam:ResourceTypes": [
                        "AWS::CloudWatch::Metric",
                        "AWS::Logs::LogGroup",
                    ],
                },
            },
        }],
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_sink = aws.oam.Sink("exampleSink")
example_sink_policy = aws.oam.SinkPolicy("exampleSinkPolicy",
    sink_identifier=example_sink.id,
    policy=json.dumps({
        "Version": "2012-10-17",
        "Statement": [{
            "Action": [
                "oam:CreateLink",
                "oam:UpdateLink",
            ],
            "Effect": "Allow",
            "Resource": "*",
            "Principal": {
                "AWS": [
                    "1111111111111",
                    "222222222222",
                ],
            },
            "Condition": {
                "ForAllValues:StringEquals": {
                    "oam:ResourceTypes": [
                        "AWS::CloudWatch::Metric",
                        "AWS::Logs::LogGroup",
                    ],
                },
            },
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
    var exampleSink = new Aws.Oam.Sink("exampleSink");

    var exampleSinkPolicy = new Aws.Oam.SinkPolicy("exampleSinkPolicy", new()
    {
        SinkIdentifier = exampleSink.Id,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2012-10-17",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = new[]
                    {
                        "oam:CreateLink",
                        "oam:UpdateLink",
                    },
                    ["Effect"] = "Allow",
                    ["Resource"] = "*",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["AWS"] = new[]
                        {
                            "1111111111111",
                            "222222222222",
                        },
                    },
                    ["Condition"] = new Dictionary<string, object?>
                    {
                        ["ForAllValues:StringEquals"] = new Dictionary<string, object?>
                        {
                            ["oam:ResourceTypes"] = new[]
                            {
                                "AWS::CloudWatch::Metric",
                                "AWS::Logs::LogGroup",
                            },
                        },
                    },
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/oam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleSink, err := oam.NewSink(ctx, "exampleSink", nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": []string{
						"oam:CreateLink",
						"oam:UpdateLink",
					},
					"Effect":   "Allow",
					"Resource": "*",
					"Principal": map[string]interface{}{
						"AWS": []string{
							"1111111111111",
							"222222222222",
						},
					},
					"Condition": map[string]interface{}{
						"ForAllValues:StringEquals": map[string]interface{}{
							"oam:ResourceTypes": []string{
								"AWS::CloudWatch::Metric",
								"AWS::Logs::LogGroup",
							},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = oam.NewSinkPolicy(ctx, "exampleSinkPolicy", &oam.SinkPolicyArgs{
			SinkIdentifier: exampleSink.ID(),
			Policy:         pulumi.String(json0),
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
import com.pulumi.aws.oam.Sink;
import com.pulumi.aws.oam.SinkPolicy;
import com.pulumi.aws.oam.SinkPolicyArgs;
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
        var exampleSink = new Sink("exampleSink");

        var exampleSinkPolicy = new SinkPolicy("exampleSinkPolicy", SinkPolicyArgs.builder()        
            .sinkIdentifier(exampleSink.id())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2012-10-17"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", jsonArray(
                            "oam:CreateLink", 
                            "oam:UpdateLink"
                        )),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Resource", "*"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", jsonArray(
                                "1111111111111", 
                                "222222222222"
                            ))
                        )),
                        jsonProperty("Condition", jsonObject(
                            jsonProperty("ForAllValues:StringEquals", jsonObject(
                                jsonProperty("oam:ResourceTypes", jsonArray(
                                    "AWS::CloudWatch::Metric", 
                                    "AWS::Logs::LogGroup"
                                ))
                            ))
                        ))
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  exampleSink:
    type: aws:oam:Sink
  exampleSinkPolicy:
    type: aws:oam:SinkPolicy
    properties:
      sinkIdentifier: ${exampleSink.id}
      policy:
        fn::toJSON:
          Version: 2012-10-17
          Statement:
            - Action:
                - oam:CreateLink
                - oam:UpdateLink
              Effect: Allow
              Resource: '*'
              Principal:
                AWS:
                  - '1111111111111'
                  - '222222222222'
              Condition:
                ForAllValues:StringEquals:
                  oam:ResourceTypes:
                    - AWS::CloudWatch::Metric
                    - AWS::Logs::LogGroup
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudWatch Observability Access Manager Sink Policy using the `sink_identifier`. For example:

```sh
 $ pulumi import aws:oam/sinkPolicy:SinkPolicy example arn:aws:oam:us-west-2:123456789012:sink/sink-id
```
 