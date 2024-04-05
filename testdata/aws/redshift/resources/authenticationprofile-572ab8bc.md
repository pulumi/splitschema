Creates a Redshift authentication profile

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.AuthenticationProfile("example", {
    authenticationProfileName: "example",
    authenticationProfileContent: JSON.stringify({
        AllowDBUserOverride: "1",
        Client_ID: "ExampleClientID",
        App_ID: "example",
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.redshift.AuthenticationProfile("example",
    authentication_profile_name="example",
    authentication_profile_content=json.dumps({
        "AllowDBUserOverride": "1",
        "Client_ID": "ExampleClientID",
        "App_ID": "example",
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
    var example = new Aws.RedShift.AuthenticationProfile("example", new()
    {
        AuthenticationProfileName = "example",
        AuthenticationProfileContent = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["AllowDBUserOverride"] = "1",
            ["Client_ID"] = "ExampleClientID",
            ["App_ID"] = "example",
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"AllowDBUserOverride": "1",
			"Client_ID":           "ExampleClientID",
			"App_ID":              "example",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = redshift.NewAuthenticationProfile(ctx, "example", &redshift.AuthenticationProfileArgs{
			AuthenticationProfileName:    pulumi.String("example"),
			AuthenticationProfileContent: pulumi.String(json0),
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
import com.pulumi.aws.redshift.AuthenticationProfile;
import com.pulumi.aws.redshift.AuthenticationProfileArgs;
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
        var example = new AuthenticationProfile("example", AuthenticationProfileArgs.builder()        
            .authenticationProfileName("example")
            .authenticationProfileContent(serializeJson(
                jsonObject(
                    jsonProperty("AllowDBUserOverride", "1"),
                    jsonProperty("Client_ID", "ExampleClientID"),
                    jsonProperty("App_ID", "example")
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:AuthenticationProfile
    properties:
      authenticationProfileName: example
      authenticationProfileContent:
        fn::toJSON:
          AllowDBUserOverride: '1'
          Client_ID: ExampleClientID
          App_ID: example
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Authentication by `authentication_profile_name`. For example:

```sh
 $ pulumi import aws:redshift/authenticationProfile:AuthenticationProfile test example
```
 