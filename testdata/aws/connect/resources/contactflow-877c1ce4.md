Provides an Amazon Connect Contact Flow resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

This resource embeds or references Contact Flows specified in Amazon Connect Contact Flow Language. For more information see
[Amazon Connect Flow language](https://docs.aws.amazon.com/connect/latest/adminguide/flow-language.html)

!> **WARN:** Contact Flows exported from the Console [Contact Flow import/export](https://docs.aws.amazon.com/connect/latest/adminguide/contact-flow-import-export.html) are not in the Amazon Connect Contact Flow Language and can not be used with this resource. Instead, the recommendation is to use the AWS CLI [`describe-contact-flow`](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/connect/describe-contact-flow.html).
See example below which uses `jq` to extract the `Content` attribute and saves it to a local file.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.connect.ContactFlow("test", {
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    description: "Test Contact Flow Description",
    type: "CONTACT_FLOW",
    content: JSON.stringify({
        Version: "2019-10-30",
        StartAction: "12345678-1234-1234-1234-123456789012",
        Actions: [
            {
                Identifier: "12345678-1234-1234-1234-123456789012",
                Type: "MessageParticipant",
                Transitions: {
                    NextAction: "abcdef-abcd-abcd-abcd-abcdefghijkl",
                    Errors: [],
                    Conditions: [],
                },
                Parameters: {
                    Text: "Thanks for calling the sample flow!",
                },
            },
            {
                Identifier: "abcdef-abcd-abcd-abcd-abcdefghijkl",
                Type: "DisconnectParticipant",
                Transitions: {},
                Parameters: {},
            },
        ],
    }),
    tags: {
        Name: "Test Contact Flow",
        Application: "Example",
        Method: "Create",
    },
});
```
```python
import pulumi
import json
import pulumi_aws as aws

test = aws.connect.ContactFlow("test",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    description="Test Contact Flow Description",
    type="CONTACT_FLOW",
    content=json.dumps({
        "Version": "2019-10-30",
        "StartAction": "12345678-1234-1234-1234-123456789012",
        "Actions": [
            {
                "Identifier": "12345678-1234-1234-1234-123456789012",
                "Type": "MessageParticipant",
                "Transitions": {
                    "NextAction": "abcdef-abcd-abcd-abcd-abcdefghijkl",
                    "Errors": [],
                    "Conditions": [],
                },
                "Parameters": {
                    "Text": "Thanks for calling the sample flow!",
                },
            },
            {
                "Identifier": "abcdef-abcd-abcd-abcd-abcdefghijkl",
                "Type": "DisconnectParticipant",
                "Transitions": {},
                "Parameters": {},
            },
        ],
    }),
    tags={
        "Name": "Test Contact Flow",
        "Application": "Example",
        "Method": "Create",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Connect.ContactFlow("test", new()
    {
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        Description = "Test Contact Flow Description",
        Type = "CONTACT_FLOW",
        Content = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2019-10-30",
            ["StartAction"] = "12345678-1234-1234-1234-123456789012",
            ["Actions"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Identifier"] = "12345678-1234-1234-1234-123456789012",
                    ["Type"] = "MessageParticipant",
                    ["Transitions"] = new Dictionary<string, object?>
                    {
                        ["NextAction"] = "abcdef-abcd-abcd-abcd-abcdefghijkl",
                        ["Errors"] = new[]
                        {
                        },
                        ["Conditions"] = new[]
                        {
                        },
                    },
                    ["Parameters"] = new Dictionary<string, object?>
                    {
                        ["Text"] = "Thanks for calling the sample flow!",
                    },
                },
                new Dictionary<string, object?>
                {
                    ["Identifier"] = "abcdef-abcd-abcd-abcd-abcdefghijkl",
                    ["Type"] = "DisconnectParticipant",
                    ["Transitions"] = new Dictionary<string, object?>
                    {
                    },
                    ["Parameters"] = new Dictionary<string, object?>
                    {
                    },
                },
            },
        }),
        Tags = 
        {
            { "Name", "Test Contact Flow" },
            { "Application", "Example" },
            { "Method", "Create" },
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Version":     "2019-10-30",
			"StartAction": "12345678-1234-1234-1234-123456789012",
			"Actions": []interface{}{
				map[string]interface{}{
					"Identifier": "12345678-1234-1234-1234-123456789012",
					"Type":       "MessageParticipant",
					"Transitions": map[string]interface{}{
						"NextAction": "abcdef-abcd-abcd-abcd-abcdefghijkl",
						"Errors":     []interface{}{},
						"Conditions": []interface{}{},
					},
					"Parameters": map[string]interface{}{
						"Text": "Thanks for calling the sample flow!",
					},
				},
				map[string]interface{}{
					"Identifier":  "abcdef-abcd-abcd-abcd-abcdefghijkl",
					"Type":        "DisconnectParticipant",
					"Transitions": nil,
					"Parameters":  nil,
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = connect.NewContactFlow(ctx, "test", &connect.ContactFlowArgs{
			InstanceId:  pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
			Description: pulumi.String("Test Contact Flow Description"),
			Type:        pulumi.String("CONTACT_FLOW"),
			Content:     pulumi.String(json0),
			Tags: pulumi.StringMap{
				"Name":        pulumi.String("Test Contact Flow"),
				"Application": pulumi.String("Example"),
				"Method":      pulumi.String("Create"),
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
import com.pulumi.aws.connect.ContactFlow;
import com.pulumi.aws.connect.ContactFlowArgs;
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
        var test = new ContactFlow("test", ContactFlowArgs.builder()        
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .description("Test Contact Flow Description")
            .type("CONTACT_FLOW")
            .content(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2019-10-30"),
                    jsonProperty("StartAction", "12345678-1234-1234-1234-123456789012"),
                    jsonProperty("Actions", jsonArray(
                        jsonObject(
                            jsonProperty("Identifier", "12345678-1234-1234-1234-123456789012"),
                            jsonProperty("Type", "MessageParticipant"),
                            jsonProperty("Transitions", jsonObject(
                                jsonProperty("NextAction", "abcdef-abcd-abcd-abcd-abcdefghijkl"),
                                jsonProperty("Errors", jsonArray(
                                )),
                                jsonProperty("Conditions", jsonArray(
                                ))
                            )),
                            jsonProperty("Parameters", jsonObject(
                                jsonProperty("Text", "Thanks for calling the sample flow!")
                            ))
                        ), 
                        jsonObject(
                            jsonProperty("Identifier", "abcdef-abcd-abcd-abcd-abcdefghijkl"),
                            jsonProperty("Type", "DisconnectParticipant"),
                            jsonProperty("Transitions", jsonObject(

                            )),
                            jsonProperty("Parameters", jsonObject(

                            ))
                        )
                    ))
                )))
            .tags(Map.ofEntries(
                Map.entry("Name", "Test Contact Flow"),
                Map.entry("Application", "Example"),
                Map.entry("Method", "Create")
            ))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:connect:ContactFlow
    properties:
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
      description: Test Contact Flow Description
      type: CONTACT_FLOW
      content:
        fn::toJSON:
          Version: 2019-10-30
          StartAction: 12345678-1234-1234-1234-123456789012
          Actions:
            - Identifier: 12345678-1234-1234-1234-123456789012
              Type: MessageParticipant
              Transitions:
                NextAction: abcdef-abcd-abcd-abcd-abcdefghijkl
                Errors: []
                Conditions: []
              Parameters:
                Text: Thanks for calling the sample flow!
            - Identifier: abcdef-abcd-abcd-abcd-abcdefghijkl
              Type: DisconnectParticipant
              Transitions: {}
              Parameters: {}
      tags:
        Name: Test Contact Flow
        Application: Example
        Method: Create
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect Contact Flows using the `instance_id` and `contact_flow_id` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:connect/contactFlow:ContactFlow example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
```
 