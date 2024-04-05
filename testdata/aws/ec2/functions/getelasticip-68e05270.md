`aws.ec2.Eip` provides details about a specific Elastic IP.

{{% examples %}}
## Example Usage
{{% example %}}
### Search By Allocation ID (VPC only)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byAllocationId = aws.ec2.getElasticIp({
    id: "eipalloc-12345678",
});
```
```python
import pulumi
import pulumi_aws as aws

by_allocation_id = aws.ec2.get_elastic_ip(id="eipalloc-12345678")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var byAllocationId = Aws.Ec2.GetElasticIp.Invoke(new()
    {
        Id = "eipalloc-12345678",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
			Id: pulumi.StringRef("eipalloc-12345678"),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetElasticIpArgs;
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
        final var byAllocationId = Ec2Functions.getElasticIp(GetElasticIpArgs.builder()
            .id("eipalloc-12345678")
            .build());

    }
}
```
```yaml
variables:
  byAllocationId:
    fn::invoke:
      Function: aws:ec2:getElasticIp
      Arguments:
        id: eipalloc-12345678
```
{{% /example %}}
{{% example %}}
### Search By Filters (EC2-Classic or VPC)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byFilter = aws.ec2.getElasticIp({
    filters: [{
        name: "tag:Name",
        values: ["exampleNameTagValue"],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

by_filter = aws.ec2.get_elastic_ip(filters=[aws.ec2.GetElasticIpFilterArgs(
    name="tag:Name",
    values=["exampleNameTagValue"],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var byFilter = Aws.Ec2.GetElasticIp.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetElasticIpFilterInputArgs
            {
                Name = "tag:Name",
                Values = new[]
                {
                    "exampleNameTagValue",
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
			Filters: []ec2.GetElasticIpFilter{
				{
					Name: "tag:Name",
					Values: []string{
						"exampleNameTagValue",
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetElasticIpArgs;
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
        final var byFilter = Ec2Functions.getElasticIp(GetElasticIpArgs.builder()
            .filters(GetElasticIpFilterArgs.builder()
                .name("tag:Name")
                .values("exampleNameTagValue")
                .build())
            .build());

    }
}
```
```yaml
variables:
  byFilter:
    fn::invoke:
      Function: aws:ec2:getElasticIp
      Arguments:
        filters:
          - name: tag:Name
            values:
              - exampleNameTagValue
```
{{% /example %}}
{{% example %}}
### Search By Public IP (EC2-Classic or VPC)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byPublicIp = aws.ec2.getElasticIp({
    publicIp: "1.2.3.4",
});
```
```python
import pulumi
import pulumi_aws as aws

by_public_ip = aws.ec2.get_elastic_ip(public_ip="1.2.3.4")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var byPublicIp = Aws.Ec2.GetElasticIp.Invoke(new()
    {
        PublicIp = "1.2.3.4",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
			PublicIp: pulumi.StringRef("1.2.3.4"),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetElasticIpArgs;
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
        final var byPublicIp = Ec2Functions.getElasticIp(GetElasticIpArgs.builder()
            .publicIp("1.2.3.4")
            .build());

    }
}
```
```yaml
variables:
  byPublicIp:
    fn::invoke:
      Function: aws:ec2:getElasticIp
      Arguments:
        publicIp: 1.2.3.4
```
{{% /example %}}
{{% example %}}
### Search By Tags (EC2-Classic or VPC)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const byTags = aws.ec2.getElasticIp({
    tags: {
        Name: "exampleNameTagValue",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

by_tags = aws.ec2.get_elastic_ip(tags={
    "Name": "exampleNameTagValue",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var byTags = Aws.Ec2.GetElasticIp.Invoke(new()
    {
        Tags = 
        {
            { "Name", "exampleNameTagValue" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.GetElasticIp(ctx, &ec2.GetElasticIpArgs{
			Tags: map[string]interface{}{
				"Name": "exampleNameTagValue",
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetElasticIpArgs;
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
        final var byTags = Ec2Functions.getElasticIp(GetElasticIpArgs.builder()
            .tags(Map.of("Name", "exampleNameTagValue"))
            .build());

    }
}
```
```yaml
variables:
  byTags:
    fn::invoke:
      Function: aws:ec2:getElasticIp
      Arguments:
        tags:
          Name: exampleNameTagValue
```
{{% /example %}}
{{% /examples %}}