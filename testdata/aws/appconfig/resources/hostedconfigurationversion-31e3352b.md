Provides an AppConfig Hosted Configuration Version resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Freeform

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appconfig.HostedConfigurationVersion("example", {
    applicationId: aws_appconfig_application.example.id,
    configurationProfileId: aws_appconfig_configuration_profile.example.configuration_profile_id,
    description: "Example Freeform Hosted Configuration Version",
    contentType: "application/json",
    content: JSON.stringify({
        foo: "bar",
        fruit: [
            "apple",
            "pear",
            "orange",
        ],
        isThingEnabled: true,
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.appconfig.HostedConfigurationVersion("example",
    application_id=aws_appconfig_application["example"]["id"],
    configuration_profile_id=aws_appconfig_configuration_profile["example"]["configuration_profile_id"],
    description="Example Freeform Hosted Configuration Version",
    content_type="application/json",
    content=json.dumps({
        "foo": "bar",
        "fruit": [
            "apple",
            "pear",
            "orange",
        ],
        "isThingEnabled": True,
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
    var example = new Aws.AppConfig.HostedConfigurationVersion("example", new()
    {
        ApplicationId = aws_appconfig_application.Example.Id,
        ConfigurationProfileId = aws_appconfig_configuration_profile.Example.Configuration_profile_id,
        Description = "Example Freeform Hosted Configuration Version",
        ContentType = "application/json",
        Content = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["foo"] = "bar",
            ["fruit"] = new[]
            {
                "apple",
                "pear",
                "orange",
            },
            ["isThingEnabled"] = true,
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appconfig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"foo": "bar",
			"fruit": []string{
				"apple",
				"pear",
				"orange",
			},
			"isThingEnabled": true,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = appconfig.NewHostedConfigurationVersion(ctx, "example", &appconfig.HostedConfigurationVersionArgs{
			ApplicationId:          pulumi.Any(aws_appconfig_application.Example.Id),
			ConfigurationProfileId: pulumi.Any(aws_appconfig_configuration_profile.Example.Configuration_profile_id),
			Description:            pulumi.String("Example Freeform Hosted Configuration Version"),
			ContentType:            pulumi.String("application/json"),
			Content:                pulumi.String(json0),
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
import com.pulumi.aws.appconfig.HostedConfigurationVersion;
import com.pulumi.aws.appconfig.HostedConfigurationVersionArgs;
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
        var example = new HostedConfigurationVersion("example", HostedConfigurationVersionArgs.builder()        
            .applicationId(aws_appconfig_application.example().id())
            .configurationProfileId(aws_appconfig_configuration_profile.example().configuration_profile_id())
            .description("Example Freeform Hosted Configuration Version")
            .contentType("application/json")
            .content(serializeJson(
                jsonObject(
                    jsonProperty("foo", "bar"),
                    jsonProperty("fruit", jsonArray(
                        "apple", 
                        "pear", 
                        "orange"
                    )),
                    jsonProperty("isThingEnabled", true)
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appconfig:HostedConfigurationVersion
    properties:
      applicationId: ${aws_appconfig_application.example.id}
      configurationProfileId: ${aws_appconfig_configuration_profile.example.configuration_profile_id}
      description: Example Freeform Hosted Configuration Version
      contentType: application/json
      content:
        fn::toJSON:
          foo: bar
          fruit:
            - apple
            - pear
            - orange
          isThingEnabled: true
```
{{% /example %}}
{{% example %}}
### Feature Flags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appconfig.HostedConfigurationVersion("example", {
    applicationId: aws_appconfig_application.example.id,
    configurationProfileId: aws_appconfig_configuration_profile.example.configuration_profile_id,
    description: "Example Feature Flag Configuration Version",
    contentType: "application/json",
    content: JSON.stringify({
        flags: {
            foo: {
                name: "foo",
                _deprecation: {
                    status: "planned",
                },
            },
            bar: {
                name: "bar",
                attributes: {
                    someAttribute: {
                        constraints: {
                            type: "string",
                            required: true,
                        },
                    },
                    someOtherAttribute: {
                        constraints: {
                            type: "number",
                            required: true,
                        },
                    },
                },
            },
        },
        values: {
            foo: {
                enabled: "true",
            },
            bar: {
                enabled: "true",
                someAttribute: "Hello World",
                someOtherAttribute: 123,
            },
        },
        version: "1",
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.appconfig.HostedConfigurationVersion("example",
    application_id=aws_appconfig_application["example"]["id"],
    configuration_profile_id=aws_appconfig_configuration_profile["example"]["configuration_profile_id"],
    description="Example Feature Flag Configuration Version",
    content_type="application/json",
    content=json.dumps({
        "flags": {
            "foo": {
                "name": "foo",
                "_deprecation": {
                    "status": "planned",
                },
            },
            "bar": {
                "name": "bar",
                "attributes": {
                    "someAttribute": {
                        "constraints": {
                            "type": "string",
                            "required": True,
                        },
                    },
                    "someOtherAttribute": {
                        "constraints": {
                            "type": "number",
                            "required": True,
                        },
                    },
                },
            },
        },
        "values": {
            "foo": {
                "enabled": "true",
            },
            "bar": {
                "enabled": "true",
                "someAttribute": "Hello World",
                "someOtherAttribute": 123,
            },
        },
        "version": "1",
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
    var example = new Aws.AppConfig.HostedConfigurationVersion("example", new()
    {
        ApplicationId = aws_appconfig_application.Example.Id,
        ConfigurationProfileId = aws_appconfig_configuration_profile.Example.Configuration_profile_id,
        Description = "Example Feature Flag Configuration Version",
        ContentType = "application/json",
        Content = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["flags"] = new Dictionary<string, object?>
            {
                ["foo"] = new Dictionary<string, object?>
                {
                    ["name"] = "foo",
                    ["_deprecation"] = new Dictionary<string, object?>
                    {
                        ["status"] = "planned",
                    },
                },
                ["bar"] = new Dictionary<string, object?>
                {
                    ["name"] = "bar",
                    ["attributes"] = new Dictionary<string, object?>
                    {
                        ["someAttribute"] = new Dictionary<string, object?>
                        {
                            ["constraints"] = new Dictionary<string, object?>
                            {
                                ["type"] = "string",
                                ["required"] = true,
                            },
                        },
                        ["someOtherAttribute"] = new Dictionary<string, object?>
                        {
                            ["constraints"] = new Dictionary<string, object?>
                            {
                                ["type"] = "number",
                                ["required"] = true,
                            },
                        },
                    },
                },
            },
            ["values"] = new Dictionary<string, object?>
            {
                ["foo"] = new Dictionary<string, object?>
                {
                    ["enabled"] = "true",
                },
                ["bar"] = new Dictionary<string, object?>
                {
                    ["enabled"] = "true",
                    ["someAttribute"] = "Hello World",
                    ["someOtherAttribute"] = 123,
                },
            },
            ["version"] = "1",
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/appconfig"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"flags": map[string]interface{}{
				"foo": map[string]interface{}{
					"name": "foo",
					"_deprecation": map[string]interface{}{
						"status": "planned",
					},
				},
				"bar": map[string]interface{}{
					"name": "bar",
					"attributes": map[string]interface{}{
						"someAttribute": map[string]interface{}{
							"constraints": map[string]interface{}{
								"type":     "string",
								"required": true,
							},
						},
						"someOtherAttribute": map[string]interface{}{
							"constraints": map[string]interface{}{
								"type":     "number",
								"required": true,
							},
						},
					},
				},
			},
			"values": map[string]interface{}{
				"foo": map[string]interface{}{
					"enabled": "true",
				},
				"bar": map[string]interface{}{
					"enabled":            "true",
					"someAttribute":      "Hello World",
					"someOtherAttribute": 123,
				},
			},
			"version": "1",
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = appconfig.NewHostedConfigurationVersion(ctx, "example", &appconfig.HostedConfigurationVersionArgs{
			ApplicationId:          pulumi.Any(aws_appconfig_application.Example.Id),
			ConfigurationProfileId: pulumi.Any(aws_appconfig_configuration_profile.Example.Configuration_profile_id),
			Description:            pulumi.String("Example Feature Flag Configuration Version"),
			ContentType:            pulumi.String("application/json"),
			Content:                pulumi.String(json0),
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
import com.pulumi.aws.appconfig.HostedConfigurationVersion;
import com.pulumi.aws.appconfig.HostedConfigurationVersionArgs;
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
        var example = new HostedConfigurationVersion("example", HostedConfigurationVersionArgs.builder()        
            .applicationId(aws_appconfig_application.example().id())
            .configurationProfileId(aws_appconfig_configuration_profile.example().configuration_profile_id())
            .description("Example Feature Flag Configuration Version")
            .contentType("application/json")
            .content(serializeJson(
                jsonObject(
                    jsonProperty("flags", jsonObject(
                        jsonProperty("foo", jsonObject(
                            jsonProperty("name", "foo"),
                            jsonProperty("_deprecation", jsonObject(
                                jsonProperty("status", "planned")
                            ))
                        )),
                        jsonProperty("bar", jsonObject(
                            jsonProperty("name", "bar"),
                            jsonProperty("attributes", jsonObject(
                                jsonProperty("someAttribute", jsonObject(
                                    jsonProperty("constraints", jsonObject(
                                        jsonProperty("type", "string"),
                                        jsonProperty("required", true)
                                    ))
                                )),
                                jsonProperty("someOtherAttribute", jsonObject(
                                    jsonProperty("constraints", jsonObject(
                                        jsonProperty("type", "number"),
                                        jsonProperty("required", true)
                                    ))
                                ))
                            ))
                        ))
                    )),
                    jsonProperty("values", jsonObject(
                        jsonProperty("foo", jsonObject(
                            jsonProperty("enabled", "true")
                        )),
                        jsonProperty("bar", jsonObject(
                            jsonProperty("enabled", "true"),
                            jsonProperty("someAttribute", "Hello World"),
                            jsonProperty("someOtherAttribute", 123)
                        ))
                    )),
                    jsonProperty("version", "1")
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:appconfig:HostedConfigurationVersion
    properties:
      applicationId: ${aws_appconfig_application.example.id}
      configurationProfileId: ${aws_appconfig_configuration_profile.example.configuration_profile_id}
      description: Example Feature Flag Configuration Version
      contentType: application/json
      content:
        fn::toJSON:
          flags:
            foo:
              name: foo
              _deprecation:
                status: planned
            bar:
              name: bar
              attributes:
                someAttribute:
                  constraints:
                    type: string
                    required: true
                someOtherAttribute:
                  constraints:
                    type: number
                    required: true
          values:
            foo:
              enabled: 'true'
            bar:
              enabled: 'true'
              someAttribute: Hello World
              someOtherAttribute: 123
          version: '1'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import AppConfig Hosted Configuration Versions using the application ID, configuration profile ID, and version number separated by a slash (`/`). For example:

```sh
 $ pulumi import aws:appconfig/hostedConfigurationVersion:HostedConfigurationVersion example 71abcde/11xxxxx/2
```
 