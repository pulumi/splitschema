Provides an Amazon Connect Contact Flow Module resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

This resource embeds or references Contact Flows Modules specified in Amazon Connect Contact Flow Language. For more information see
[Amazon Connect Flow language](https://docs.aws.amazon.com/connect/latest/adminguide/flow-language.html)

!> **WARN:** Contact Flow Modules exported from the Console [See Contact Flow import/export which is the same for Contact Flow Modules](https://docs.aws.amazon.com/connect/latest/adminguide/contact-flow-import-export.html) are not in the Amazon Connect Contact Flow Language and can not be used with this resource. Instead, the recommendation is to use the AWS CLI [`describe-contact-flow-module`](https://docs.aws.amazon.com/cli/latest/reference/connect/describe-contact-flow-module.html).
See example below which uses `jq` to extract the `Content` attribute and saves it to a local file.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.ContactFlowModule("example", {
    instanceId: "aaaaaaaa-bbbb-cccc-dddd-111111111111",
    description: "Example Contact Flow Module Description",
    content: JSON.stringify({
        Version: "2019-10-30",
        StartAction: "12345678-1234-1234-1234-123456789012",
        Actions: [
            {
                Identifier: "12345678-1234-1234-1234-123456789012",
                Parameters: {
                    Text: "Hello contact flow module",
                },
                Transitions: {
                    NextAction: "abcdef-abcd-abcd-abcd-abcdefghijkl",
                    Errors: [],
                    Conditions: [],
                },
                Type: "MessageParticipant",
            },
            {
                Identifier: "abcdef-abcd-abcd-abcd-abcdefghijkl",
                Type: "DisconnectParticipant",
                Parameters: {},
                Transitions: {},
            },
        ],
        Settings: {
            InputParameters: [],
            OutputParameters: [],
            Transitions: [
                {
                    DisplayName: "Success",
                    ReferenceName: "Success",
                    Description: "",
                },
                {
                    DisplayName: "Error",
                    ReferenceName: "Error",
                    Description: "",
                },
            ],
        },
    }),
    tags: {
        Name: "Example Contact Flow Module",
        Application: "Example",
        Method: "Create",
    },
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.connect.ContactFlowModule("example",
    instance_id="aaaaaaaa-bbbb-cccc-dddd-111111111111",
    description="Example Contact Flow Module Description",
    content=json.dumps({
        "Version": "2019-10-30",
        "StartAction": "12345678-1234-1234-1234-123456789012",
        "Actions": [
            {
                "Identifier": "12345678-1234-1234-1234-123456789012",
                "Parameters": {
                    "Text": "Hello contact flow module",
                },
                "Transitions": {
                    "NextAction": "abcdef-abcd-abcd-abcd-abcdefghijkl",
                    "Errors": [],
                    "Conditions": [],
                },
                "Type": "MessageParticipant",
            },
            {
                "Identifier": "abcdef-abcd-abcd-abcd-abcdefghijkl",
                "Type": "DisconnectParticipant",
                "Parameters": {},
                "Transitions": {},
            },
        ],
        "Settings": {
            "InputParameters": [],
            "OutputParameters": [],
            "Transitions": [
                {
                    "DisplayName": "Success",
                    "ReferenceName": "Success",
                    "Description": "",
                },
                {
                    "DisplayName": "Error",
                    "ReferenceName": "Error",
                    "Description": "",
                },
            ],
        },
    }),
    tags={
        "Name": "Example Contact Flow Module",
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
    var example = new Aws.Connect.ContactFlowModule("example", new()
    {
        InstanceId = "aaaaaaaa-bbbb-cccc-dddd-111111111111",
        Description = "Example Contact Flow Module Description",
        Content = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Version"] = "2019-10-30",
            ["StartAction"] = "12345678-1234-1234-1234-123456789012",
            ["Actions"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Identifier"] = "12345678-1234-1234-1234-123456789012",
                    ["Parameters"] = new Dictionary<string, object?>
                    {
                        ["Text"] = "Hello contact flow module",
                    },
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
                    ["Type"] = "MessageParticipant",
                },
                new Dictionary<string, object?>
                {
                    ["Identifier"] = "abcdef-abcd-abcd-abcd-abcdefghijkl",
                    ["Type"] = "DisconnectParticipant",
                    ["Parameters"] = new Dictionary<string, object?>
                    {
                    },
                    ["Transitions"] = new Dictionary<string, object?>
                    {
                    },
                },
            },
            ["Settings"] = new Dictionary<string, object?>
            {
                ["InputParameters"] = new[]
                {
                },
                ["OutputParameters"] = new[]
                {
                },
                ["Transitions"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["DisplayName"] = "Success",
                        ["ReferenceName"] = "Success",
                        ["Description"] = "",
                    },
                    new Dictionary<string, object?>
                    {
                        ["DisplayName"] = "Error",
                        ["ReferenceName"] = "Error",
                        ["Description"] = "",
                    },
                },
            },
        }),
        Tags = 
        {
            { "Name", "Example Contact Flow Module" },
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
					"Parameters": map[string]interface{}{
						"Text": "Hello contact flow module",
					},
					"Transitions": map[string]interface{}{
						"NextAction": "abcdef-abcd-abcd-abcd-abcdefghijkl",
						"Errors":     []interface{}{},
						"Conditions": []interface{}{},
					},
					"Type": "MessageParticipant",
				},
				map[string]interface{}{
					"Identifier":  "abcdef-abcd-abcd-abcd-abcdefghijkl",
					"Type":        "DisconnectParticipant",
					"Parameters":  nil,
					"Transitions": nil,
				},
			},
			"Settings": map[string]interface{}{
				"InputParameters":  []interface{}{},
				"OutputParameters": []interface{}{},
				"Transitions": []map[string]interface{}{
					map[string]interface{}{
						"DisplayName":   "Success",
						"ReferenceName": "Success",
						"Description":   "",
					},
					map[string]interface{}{
						"DisplayName":   "Error",
						"ReferenceName": "Error",
						"Description":   "",
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = connect.NewContactFlowModule(ctx, "example", &connect.ContactFlowModuleArgs{
			InstanceId:  pulumi.String("aaaaaaaa-bbbb-cccc-dddd-111111111111"),
			Description: pulumi.String("Example Contact Flow Module Description"),
			Content:     pulumi.String(json0),
			Tags: pulumi.StringMap{
				"Name":        pulumi.String("Example Contact Flow Module"),
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
import com.pulumi.aws.connect.ContactFlowModule;
import com.pulumi.aws.connect.ContactFlowModuleArgs;
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
        var example = new ContactFlowModule("example", ContactFlowModuleArgs.builder()        
            .instanceId("aaaaaaaa-bbbb-cccc-dddd-111111111111")
            .description("Example Contact Flow Module Description")
            .content(serializeJson(
                jsonObject(
                    jsonProperty("Version", "2019-10-30"),
                    jsonProperty("StartAction", "12345678-1234-1234-1234-123456789012"),
                    jsonProperty("Actions", jsonArray(
                        jsonObject(
                            jsonProperty("Identifier", "12345678-1234-1234-1234-123456789012"),
                            jsonProperty("Parameters", jsonObject(
                                jsonProperty("Text", "Hello contact flow module")
                            )),
                            jsonProperty("Transitions", jsonObject(
                                jsonProperty("NextAction", "abcdef-abcd-abcd-abcd-abcdefghijkl"),
                                jsonProperty("Errors", jsonArray(
                                )),
                                jsonProperty("Conditions", jsonArray(
                                ))
                            )),
                            jsonProperty("Type", "MessageParticipant")
                        ), 
                        jsonObject(
                            jsonProperty("Identifier", "abcdef-abcd-abcd-abcd-abcdefghijkl"),
                            jsonProperty("Type", "DisconnectParticipant"),
                            jsonProperty("Parameters", jsonObject(

                            )),
                            jsonProperty("Transitions", jsonObject(

                            ))
                        )
                    )),
                    jsonProperty("Settings", jsonObject(
                        jsonProperty("InputParameters", jsonArray(
                        )),
                        jsonProperty("OutputParameters", jsonArray(
                        )),
                        jsonProperty("Transitions", jsonArray(
                            jsonObject(
                                jsonProperty("DisplayName", "Success"),
                                jsonProperty("ReferenceName", "Success"),
                                jsonProperty("Description", "")
                            ), 
                            jsonObject(
                                jsonProperty("DisplayName", "Error"),
                                jsonProperty("ReferenceName", "Error"),
                                jsonProperty("Description", "")
                            )
                        ))
                    ))
                )))
            .tags(Map.ofEntries(
                Map.entry("Name", "Example Contact Flow Module"),
                Map.entry("Application", "Example"),
                Map.entry("Method", "Create")
            ))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:ContactFlowModule
    properties:
      instanceId: aaaaaaaa-bbbb-cccc-dddd-111111111111
      description: Example Contact Flow Module Description
      content:
        fn::toJSON:
          Version: 2019-10-30
          StartAction: 12345678-1234-1234-1234-123456789012
          Actions:
            - Identifier: 12345678-1234-1234-1234-123456789012
              Parameters:
                Text: Hello contact flow module
              Transitions:
                NextAction: abcdef-abcd-abcd-abcd-abcdefghijkl
                Errors: []
                Conditions: []
              Type: MessageParticipant
            - Identifier: abcdef-abcd-abcd-abcd-abcdefghijkl
              Type: DisconnectParticipant
              Parameters: {}
              Transitions: {}
          Settings:
            InputParameters: []
            OutputParameters: []
            Transitions:
              - DisplayName: Success
                ReferenceName: Success
                Description:
              - DisplayName: Error
                ReferenceName: Error
                Description:
      tags:
        Name: Example Contact Flow Module
        Application: Example
        Method: Create
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect Contact Flow Modules using the `instance_id` and `contact_flow_module_id` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:connect/contactFlowModule:ContactFlowModule example f1288a1f-6193-445a-b47e-af739b2:c1d4e5f6-1b3c-1b3c-1b3c-c1d4e5f6c1d4e5
```
 