Provides details about resource tagging.

{{% examples %}}
## Example Usage
{{% example %}}
### Get All Resource Tag Mappings

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.resourcegroupstaggingapi.getResources({});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.resourcegroupstaggingapi.get_resources()
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.ResourceGroupsTaggingApi.GetResources.Invoke();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/resourcegroupstaggingapi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resourcegroupstaggingapi.GetResources(ctx, nil, nil)
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
import com.pulumi.aws.resourcegroupstaggingapi.ResourcegroupstaggingapiFunctions;
import com.pulumi.aws.resourcegroupstaggingapi.inputs.GetResourcesArgs;
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
        final var test = ResourcegroupstaggingapiFunctions.getResources();

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:resourcegroupstaggingapi:getResources
      Arguments: {}
```
{{% /example %}}
{{% example %}}
### Filter By Tag Key and Value

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.resourcegroupstaggingapi.getResources({
    tagFilters: [{
        key: "tag-key",
        values: [
            "tag-value-1",
            "tag-value-2",
        ],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.resourcegroupstaggingapi.get_resources(tag_filters=[aws.resourcegroupstaggingapi.GetResourcesTagFilterArgs(
    key="tag-key",
    values=[
        "tag-value-1",
        "tag-value-2",
    ],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.ResourceGroupsTaggingApi.GetResources.Invoke(new()
    {
        TagFilters = new[]
        {
            new Aws.ResourceGroupsTaggingApi.Inputs.GetResourcesTagFilterInputArgs
            {
                Key = "tag-key",
                Values = new[]
                {
                    "tag-value-1",
                    "tag-value-2",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/resourcegroupstaggingapi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resourcegroupstaggingapi.GetResources(ctx, &resourcegroupstaggingapi.GetResourcesArgs{
			TagFilters: []resourcegroupstaggingapi.GetResourcesTagFilter{
				{
					Key: "tag-key",
					Values: []string{
						"tag-value-1",
						"tag-value-2",
					},
				},
			},
		}, nil)
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
import com.pulumi.aws.resourcegroupstaggingapi.ResourcegroupstaggingapiFunctions;
import com.pulumi.aws.resourcegroupstaggingapi.inputs.GetResourcesArgs;
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
        final var test = ResourcegroupstaggingapiFunctions.getResources(GetResourcesArgs.builder()
            .tagFilters(GetResourcesTagFilterArgs.builder()
                .key("tag-key")
                .values(                
                    "tag-value-1",
                    "tag-value-2")
                .build())
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:resourcegroupstaggingapi:getResources
      Arguments:
        tagFilters:
          - key: tag-key
            values:
              - tag-value-1
              - tag-value-2
```
{{% /example %}}
{{% example %}}
### Filter By Resource Type

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.resourcegroupstaggingapi.getResources({
    resourceTypeFilters: ["ec2:instance"],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.resourcegroupstaggingapi.get_resources(resource_type_filters=["ec2:instance"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.ResourceGroupsTaggingApi.GetResources.Invoke(new()
    {
        ResourceTypeFilters = new[]
        {
            "ec2:instance",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/resourcegroupstaggingapi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resourcegroupstaggingapi.GetResources(ctx, &resourcegroupstaggingapi.GetResourcesArgs{
			ResourceTypeFilters: []string{
				"ec2:instance",
			},
		}, nil)
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
import com.pulumi.aws.resourcegroupstaggingapi.ResourcegroupstaggingapiFunctions;
import com.pulumi.aws.resourcegroupstaggingapi.inputs.GetResourcesArgs;
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
        final var test = ResourcegroupstaggingapiFunctions.getResources(GetResourcesArgs.builder()
            .resourceTypeFilters("ec2:instance")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:resourcegroupstaggingapi:getResources
      Arguments:
        resourceTypeFilters:
          - ec2:instance
```
{{% /example %}}
{{% /examples %}}