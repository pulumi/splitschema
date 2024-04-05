Manages an EKS add-on.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.eks.Addon("example", {
    clusterName: aws_eks_cluster.example.name,
    addonName: "vpc-cni",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.eks.Addon("example",
    cluster_name=aws_eks_cluster["example"]["name"],
    addon_name="vpc-cni")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Eks.Addon("example", new()
    {
        ClusterName = aws_eks_cluster.Example.Name,
        AddonName = "vpc-cni",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eks.NewAddon(ctx, "example", &eks.AddonArgs{
			ClusterName: pulumi.Any(aws_eks_cluster.Example.Name),
			AddonName:   pulumi.String("vpc-cni"),
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
import com.pulumi.aws.eks.Addon;
import com.pulumi.aws.eks.AddonArgs;
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
        var example = new Addon("example", AddonArgs.builder()        
            .clusterName(aws_eks_cluster.example().name())
            .addonName("vpc-cni")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:eks:Addon
    properties:
      clusterName: ${aws_eks_cluster.example.name}
      addonName: vpc-cni
```
{{% /example %}}
{{% /examples %}}
## Example Update add-on usage with resolve_conflicts_on_update and PRESERVE

`resolve_conflicts_on_update` with `PRESERVE` can be used to retain the config changes applied to the add-on with kubectl while upgrading to a newer version of the add-on.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.eks.Addon("example", {
    clusterName: aws_eks_cluster.example.name,
    addonName: "coredns",
    addonVersion: "v1.10.1-eksbuild.1",
    resolveConflictsOnUpdate: "PRESERVE",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.eks.Addon("example",
    cluster_name=aws_eks_cluster["example"]["name"],
    addon_name="coredns",
    addon_version="v1.10.1-eksbuild.1",
    resolve_conflicts_on_update="PRESERVE")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Eks.Addon("example", new()
    {
        ClusterName = aws_eks_cluster.Example.Name,
        AddonName = "coredns",
        AddonVersion = "v1.10.1-eksbuild.1",
        ResolveConflictsOnUpdate = "PRESERVE",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eks.NewAddon(ctx, "example", &eks.AddonArgs{
			ClusterName:              pulumi.Any(aws_eks_cluster.Example.Name),
			AddonName:                pulumi.String("coredns"),
			AddonVersion:             pulumi.String("v1.10.1-eksbuild.1"),
			ResolveConflictsOnUpdate: pulumi.String("PRESERVE"),
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
import com.pulumi.aws.eks.Addon;
import com.pulumi.aws.eks.AddonArgs;
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
        var example = new Addon("example", AddonArgs.builder()        
            .clusterName(aws_eks_cluster.example().name())
            .addonName("coredns")
            .addonVersion("v1.10.1-eksbuild.1")
            .resolveConflictsOnUpdate("PRESERVE")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:eks:Addon
    properties:
      clusterName: ${aws_eks_cluster.example.name}
      addonName: coredns
      addonVersion: v1.10.1-eksbuild.1
      #e.g., previous version v1.9.3-eksbuild.3 and the new version is v1.10.1-eksbuild.1
      resolveConflictsOnUpdate: PRESERVE
```

## Example add-on usage with custom configuration_values

Custom add-on configuration can be passed using `configuration_values` as a single JSON string while creating or updating the add-on.

> **Note:** `configuration_values` is a single JSON string should match the valid JSON schema for each add-on with specific version.

To find the correct JSON schema for each add-on can be extracted using [describe-addon-configuration](https://docs.aws.amazon.com/cli/latest/reference/eks/describe-addon-configuration.html) call.
This below is an example for extracting the `configuration_values` schema for `coredns`.

```typescript
import * as pulumi from "@pulumi/pulumi";
```
```python
import pulumi
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;

return await Deployment.RunAsync(() => 
{
});
```
```go
package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		return nil
	})
}
```
```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
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
    }
}
```
```yaml
{}
```

Example to create a `coredns` managed addon with custom `configuration_values`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.eks.Addon("example", {
    clusterName: "mycluster",
    addonName: "coredns",
    addonVersion: "v1.10.1-eksbuild.1",
    resolveConflictsOnCreate: "OVERWRITE",
    configurationValues: JSON.stringify({
        replicaCount: 4,
        resources: {
            limits: {
                cpu: "100m",
                memory: "150Mi",
            },
            requests: {
                cpu: "100m",
                memory: "150Mi",
            },
        },
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.eks.Addon("example",
    cluster_name="mycluster",
    addon_name="coredns",
    addon_version="v1.10.1-eksbuild.1",
    resolve_conflicts_on_create="OVERWRITE",
    configuration_values=json.dumps({
        "replicaCount": 4,
        "resources": {
            "limits": {
                "cpu": "100m",
                "memory": "150Mi",
            },
            "requests": {
                "cpu": "100m",
                "memory": "150Mi",
            },
        },
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
    var example = new Aws.Eks.Addon("example", new()
    {
        ClusterName = "mycluster",
        AddonName = "coredns",
        AddonVersion = "v1.10.1-eksbuild.1",
        ResolveConflictsOnCreate = "OVERWRITE",
        ConfigurationValues = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["replicaCount"] = 4,
            ["resources"] = new Dictionary<string, object?>
            {
                ["limits"] = new Dictionary<string, object?>
                {
                    ["cpu"] = "100m",
                    ["memory"] = "150Mi",
                },
                ["requests"] = new Dictionary<string, object?>
                {
                    ["cpu"] = "100m",
                    ["memory"] = "150Mi",
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"replicaCount": 4,
			"resources": map[string]interface{}{
				"limits": map[string]interface{}{
					"cpu":    "100m",
					"memory": "150Mi",
				},
				"requests": map[string]interface{}{
					"cpu":    "100m",
					"memory": "150Mi",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = eks.NewAddon(ctx, "example", &eks.AddonArgs{
			ClusterName:              pulumi.String("mycluster"),
			AddonName:                pulumi.String("coredns"),
			AddonVersion:             pulumi.String("v1.10.1-eksbuild.1"),
			ResolveConflictsOnCreate: pulumi.String("OVERWRITE"),
			ConfigurationValues:      pulumi.String(json0),
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
import com.pulumi.aws.eks.Addon;
import com.pulumi.aws.eks.AddonArgs;
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
        var example = new Addon("example", AddonArgs.builder()        
            .clusterName("mycluster")
            .addonName("coredns")
            .addonVersion("v1.10.1-eksbuild.1")
            .resolveConflictsOnCreate("OVERWRITE")
            .configurationValues(serializeJson(
                jsonObject(
                    jsonProperty("replicaCount", 4),
                    jsonProperty("resources", jsonObject(
                        jsonProperty("limits", jsonObject(
                            jsonProperty("cpu", "100m"),
                            jsonProperty("memory", "150Mi")
                        )),
                        jsonProperty("requests", jsonObject(
                            jsonProperty("cpu", "100m"),
                            jsonProperty("memory", "150Mi")
                        ))
                    ))
                )))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:eks:Addon
    properties:
      clusterName: mycluster
      addonName: coredns
      addonVersion: v1.10.1-eksbuild.1
      resolveConflictsOnCreate: OVERWRITE
      configurationValues:
        fn::toJSON:
          replicaCount: 4
          resources:
            limits:
              cpu: 100m
              memory: 150Mi
            requests:
              cpu: 100m
              memory: 150Mi
```


## Import

Using `pulumi import`, import EKS add-on using the `cluster_name` and `addon_name` separated by a colon (`:`). For example:

```sh
 $ pulumi import aws:eks/addon:Addon my_eks_addon my_cluster_name:my_addon_name
```
 