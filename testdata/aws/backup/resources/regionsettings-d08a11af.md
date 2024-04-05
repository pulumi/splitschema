Provides an AWS Backup Region Settings resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.backup.RegionSettings("test", {
    resourceTypeManagementPreference: {
        DynamoDB: true,
        EFS: true,
    },
    resourceTypeOptInPreference: {
        Aurora: true,
        DocumentDB: true,
        DynamoDB: true,
        EBS: true,
        EC2: true,
        EFS: true,
        FSx: true,
        Neptune: true,
        RDS: true,
        "Storage Gateway": true,
        VirtualMachine: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.backup.RegionSettings("test",
    resource_type_management_preference={
        "DynamoDB": True,
        "EFS": True,
    },
    resource_type_opt_in_preference={
        "Aurora": True,
        "DocumentDB": True,
        "DynamoDB": True,
        "EBS": True,
        "EC2": True,
        "EFS": True,
        "FSx": True,
        "Neptune": True,
        "RDS": True,
        "Storage Gateway": True,
        "VirtualMachine": True,
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Backup.RegionSettings("test", new()
    {
        ResourceTypeManagementPreference = 
        {
            { "DynamoDB", true },
            { "EFS", true },
        },
        ResourceTypeOptInPreference = 
        {
            { "Aurora", true },
            { "DocumentDB", true },
            { "DynamoDB", true },
            { "EBS", true },
            { "EC2", true },
            { "EFS", true },
            { "FSx", true },
            { "Neptune", true },
            { "RDS", true },
            { "Storage Gateway", true },
            { "VirtualMachine", true },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/backup"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := backup.NewRegionSettings(ctx, "test", &backup.RegionSettingsArgs{
			ResourceTypeManagementPreference: pulumi.BoolMap{
				"DynamoDB": pulumi.Bool(true),
				"EFS":      pulumi.Bool(true),
			},
			ResourceTypeOptInPreference: pulumi.BoolMap{
				"Aurora":          pulumi.Bool(true),
				"DocumentDB":      pulumi.Bool(true),
				"DynamoDB":        pulumi.Bool(true),
				"EBS":             pulumi.Bool(true),
				"EC2":             pulumi.Bool(true),
				"EFS":             pulumi.Bool(true),
				"FSx":             pulumi.Bool(true),
				"Neptune":         pulumi.Bool(true),
				"RDS":             pulumi.Bool(true),
				"Storage Gateway": pulumi.Bool(true),
				"VirtualMachine":  pulumi.Bool(true),
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
import com.pulumi.aws.backup.RegionSettings;
import com.pulumi.aws.backup.RegionSettingsArgs;
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
        var test = new RegionSettings("test", RegionSettingsArgs.builder()        
            .resourceTypeManagementPreference(Map.ofEntries(
                Map.entry("DynamoDB", true),
                Map.entry("EFS", true)
            ))
            .resourceTypeOptInPreference(Map.ofEntries(
                Map.entry("Aurora", true),
                Map.entry("DocumentDB", true),
                Map.entry("DynamoDB", true),
                Map.entry("EBS", true),
                Map.entry("EC2", true),
                Map.entry("EFS", true),
                Map.entry("FSx", true),
                Map.entry("Neptune", true),
                Map.entry("RDS", true),
                Map.entry("Storage Gateway", true),
                Map.entry("VirtualMachine", true)
            ))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:backup:RegionSettings
    properties:
      resourceTypeManagementPreference:
        DynamoDB: true
        EFS: true
      resourceTypeOptInPreference:
        Aurora: true
        DocumentDB: true
        DynamoDB: true
        EBS: true
        EC2: true
        EFS: true
        FSx: true
        Neptune: true
        RDS: true
        Storage Gateway: true
        VirtualMachine: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Backup Region Settings using the `region`. For example:

```sh
 $ pulumi import aws:backup/regionSettings:RegionSettings test us-west-2
```
 