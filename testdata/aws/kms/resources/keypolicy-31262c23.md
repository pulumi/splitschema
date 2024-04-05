Attaches a policy to a KMS Key.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {description: "example"});
const exampleKeyPolicy = new aws.kms.KeyPolicy("exampleKeyPolicy", {
    keyId: exampleKey.id,
    policy: JSON.stringify({
        Id: "example",
        Statement: [{
            Action: "kms:*",
            Effect: "Allow",
            Principal: {
                AWS: "*",
            },
            Resource: "*",
            Sid: "Enable IAM User Permissions",
        }],
        Version: "2012-10-17",
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey", description="example")
example_key_policy = aws.kms.KeyPolicy("exampleKeyPolicy",
    key_id=example_key.id,
    policy=json.dumps({
        "Id": "example",
        "Statement": [{
            "Action": "kms:*",
            "Effect": "Allow",
            "Principal": {
                "AWS": "*",
            },
            "Resource": "*",
            "Sid": "Enable IAM User Permissions",
        }],
        "Version": "2012-10-17",
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
    var exampleKey = new Aws.Kms.Key("exampleKey", new()
    {
        Description = "example",
    });

    var exampleKeyPolicy = new Aws.Kms.KeyPolicy("exampleKeyPolicy", new()
    {
        KeyId = exampleKey.Id,
        Policy = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["Id"] = "example",
            ["Statement"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["Action"] = "kms:*",
                    ["Effect"] = "Allow",
                    ["Principal"] = new Dictionary<string, object?>
                    {
                        ["AWS"] = "*",
                    },
                    ["Resource"] = "*",
                    ["Sid"] = "Enable IAM User Permissions",
                },
            },
            ["Version"] = "2012-10-17",
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", &kms.KeyArgs{
			Description: pulumi.String("example"),
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"Id": "example",
			"Statement": []map[string]interface{}{
				map[string]interface{}{
					"Action": "kms:*",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"AWS": "*",
					},
					"Resource": "*",
					"Sid":      "Enable IAM User Permissions",
				},
			},
			"Version": "2012-10-17",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = kms.NewKeyPolicy(ctx, "exampleKeyPolicy", &kms.KeyPolicyArgs{
			KeyId:  exampleKey.ID(),
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.kms.KeyArgs;
import com.pulumi.aws.kms.KeyPolicy;
import com.pulumi.aws.kms.KeyPolicyArgs;
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
        var exampleKey = new Key("exampleKey", KeyArgs.builder()        
            .description("example")
            .build());

        var exampleKeyPolicy = new KeyPolicy("exampleKeyPolicy", KeyPolicyArgs.builder()        
            .keyId(exampleKey.id())
            .policy(serializeJson(
                jsonObject(
                    jsonProperty("Id", "example"),
                    jsonProperty("Statement", jsonArray(jsonObject(
                        jsonProperty("Action", "kms:*"),
                        jsonProperty("Effect", "Allow"),
                        jsonProperty("Principal", jsonObject(
                            jsonProperty("AWS", "*")
                        )),
                        jsonProperty("Resource", "*"),
                        jsonProperty("Sid", "Enable IAM User Permissions")
                    ))),
                    jsonProperty("Version", "2012-10-17")
                )))
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
    properties:
      description: example
  exampleKeyPolicy:
    type: aws:kms:KeyPolicy
    properties:
      keyId: ${exampleKey.id}
      policy:
        fn::toJSON:
          Id: example
          Statement:
            - Action: kms:*
              Effect: Allow
              Principal:
                AWS: '*'
              Resource: '*'
              Sid: Enable IAM User Permissions
          Version: 2012-10-17
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import KMS Key Policies using the `key_id`. For example:

```sh
 $ pulumi import aws:kms/keyPolicy:KeyPolicy a 1234abcd-12ab-34cd-56ef-1234567890ab
```
 