Provides an SNS data protection topic policy resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleTopic = new aws.sns.Topic("exampleTopic", {});
const exampleDataProtectionPolicy = new aws.sns.DataProtectionPolicy("exampleDataProtectionPolicy", {
    arn: exampleTopic.arn,
    policy: JSON.stringify({
        Description: "Example data protection policy",
        Name: "__example_data_protection_policy",
        Statement: [{
            DataDirection: "Inbound",
            DataIdentifier: ["arn:aws:dataprotection::aws:data-identifier/EmailAddress"],
            Operation: {
                Deny: {},
            },
            Principal: ["*"],
            Sid: "__deny_statement_11ba9d96",
        }],
        Version: "2021-06-01",
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_topic = aws.sns.Topic("exampleTopic")
example_data_protection_policy = aws.sns.DataProtectionPolicy("exampleDataProtectionPolicy",
    arn=example_topic.arn,
    policy=json.dumps({
        "Description": "Example data protection policy",
        "Name": "__example_data_protection_policy",
        "Statement": [{
            "DataDirection": "Inbound",
            "DataIdentifier": ["arn:aws:dataprotection::aws:data-identifier/EmailAddress"],
            "Operation": {
                "Deny": {},
            },
            "Principal": ["*"],
            "Sid": "__deny_statement_11ba9d96",
        }],
        "Version": "2021-06-01",
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
    var exampleTopic = new Aws.Sns.Topic("exampleTopic");

    var exampleDataProtectionPolicy = new Aws.Sns.DataProtectionPolicy("exampleDataProtectionPolicy", new()
    {
        Arn = exampleTopic.Arn,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Description"] = "Example data protection policy",
            ["Name"] = "__example_data_protection_policy",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["DataDirection"] = "Inbound",
                    ["DataIdentifier"] = new[]
                    {
                        "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                    },
                    ["Operation"] = new Dictionary<string, object?>
                    {
                        ["Deny"] = new Dictionary<string, object?>
                        {
                        },
                    },
                    ["Principal"] = new[]
                    {
                        "*",
                    },
                    ["Sid"] = "__deny_statement_11ba9d96",
                },
            },
            ["Version"] = "2021-06-01",
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleTopic, err := sns.NewTopic(ctx, "exampleTopic", nil)
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Description": "Example data protection policy",
			"Name":        "__example_data_protection_policy",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"DataDirection": "Inbound",
					"DataIdentifier": []string{
						"arn:aws:dataprotection::aws:data-identifier/EmailAddress",
					},
					"Operation": map[string]interface{}{
						"Deny": nil,
					},
					"Principal": []string{
						"*",
					},
					"Sid": "__deny_statement_11ba9d96",
				},
			},
			"Version": "2021-06-01",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = sns.NewDataProtectionPolicy(ctx, "exampleDataProtectionPolicy", &sns.DataProtectionPolicyArgs{
			Arn:    exampleTopic.Arn,
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
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.sns.DataProtectionPolicy;
import com.pulumi.aws.sns.DataProtectionPolicyArgs;
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
        var exampleTopic = new Topic("exampleTopic");

        var exampleDataProtectionPolicy = new DataProtectionPolicy("exampleDataProtectionPolicy", DataProtectionPolicyArgs.builder()        
            .arn(exampleTopic.arn())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Description", "Example data protection policy"),
                    jsonProperty("Name", "__example_data_protection_policy"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("DataDirection", "Inbound"),
                        jsonProperty("DataIdentifier", jsonArray("arn:aws:dataprotection::aws:data-identifier/EmailAddress")),
                        jsonProperty("Operation", jsonObject(
                            jsonProperty("Deny", jsonObject(

                            ))
                        )),
                        jsonProperty("Principal", jsonArray("*")),
                        jsonProperty("Sid", "__deny_statement_11ba9d96")
                    ))),
                    jsonProperty("Version", "2021-06-01")
                )))
            .build());

    }
}
```
```yaml
resources:
  exampleTopic:
    type: aws:sns:Topic
  exampleDataProtectionPolicy:
    type: aws:sns:DataProtectionPolicy
    properties:
      arn: ${exampleTopic.arn}
      policy:
        fn::toJSON:
          Description: Example data protection policy
          Name: __example_data_protection_policy
          Statement:
            - DataDirection: Inbound
              DataIdentifier:
                - arn:aws:dataprotection::aws:data-identifier/EmailAddress
              Operation:
                Deny: {}
              Principal:
                - '*'
              Sid: __deny_statement_11ba9d96
          Version: 2021-06-01
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SNS Data Protection Topic Policy using the topic ARN. For example:

```sh
 $ pulumi import aws:sns/dataProtectionPolicy:DataProtectionPolicy example arn:aws:sns:us-west-2:0123456789012:example
```
 