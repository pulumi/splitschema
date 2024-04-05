Provides a resource for managing a replication set in AWS Systems Manager Incident Manager.

> **NOTE:** Deleting a replication set also deletes all Incident Manager related data including response plans, incident records, contacts and escalation plans.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

Create a replication set.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const replicationSetName = new aws.ssmincidents.ReplicationSet("replicationSetName", {
    regions: [{
        name: "us-west-2",
    }],
    tags: {
        exampleTag: "exampleValue",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

replication_set_name = aws.ssmincidents.ReplicationSet("replicationSetName",
    regions=[aws.ssmincidents.ReplicationSetRegionArgs(
        name="us-west-2",
    )],
    tags={
        "exampleTag": "exampleValue",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var replicationSetName = new Aws.SsmIncidents.ReplicationSet("replicationSetName", new()
    {
        Regions = new[]
        {
            new Aws.SsmIncidents.Inputs.ReplicationSetRegionArgs
            {
                Name = "us-west-2",
            },
        },
        Tags = 
        {
            { "exampleTag", "exampleValue" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssmincidents.NewReplicationSet(ctx, "replicationSetName", &ssmincidents.ReplicationSetArgs{
			Regions: ssmincidents.ReplicationSetRegionArray{
				&ssmincidents.ReplicationSetRegionArgs{
					Name: pulumi.String("us-west-2"),
				},
			},
			Tags: pulumi.StringMap{
				"exampleTag": pulumi.String("exampleValue"),
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
import com.pulumi.aws.ssmincidents.ReplicationSet;
import com.pulumi.aws.ssmincidents.ReplicationSetArgs;
import com.pulumi.aws.ssmincidents.inputs.ReplicationSetRegionArgs;
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
        var replicationSetName = new ReplicationSet("replicationSetName", ReplicationSetArgs.builder()        
            .regions(ReplicationSetRegionArgs.builder()
                .name("us-west-2")
                .build())
            .tags(Map.of("exampleTag", "exampleValue"))
            .build());

    }
}
```
```yaml
resources:
  replicationSetName:
    type: aws:ssmincidents:ReplicationSet
    properties:
      regions:
        - name: us-west-2
      tags:
        exampleTag: exampleValue
```

Add a Region to a replication set. (You can add only one Region at a time.)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const replicationSetName = new aws.ssmincidents.ReplicationSet("replicationSetName", {regions: [
    {
        name: "us-west-2",
    },
    {
        name: "ap-southeast-2",
    },
]});
```
```python
import pulumi
import pulumi_aws as aws

replication_set_name = aws.ssmincidents.ReplicationSet("replicationSetName", regions=[
    aws.ssmincidents.ReplicationSetRegionArgs(
        name="us-west-2",
    ),
    aws.ssmincidents.ReplicationSetRegionArgs(
        name="ap-southeast-2",
    ),
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var replicationSetName = new Aws.SsmIncidents.ReplicationSet("replicationSetName", new()
    {
        Regions = new[]
        {
            new Aws.SsmIncidents.Inputs.ReplicationSetRegionArgs
            {
                Name = "us-west-2",
            },
            new Aws.SsmIncidents.Inputs.ReplicationSetRegionArgs
            {
                Name = "ap-southeast-2",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssmincidents.NewReplicationSet(ctx, "replicationSetName", &ssmincidents.ReplicationSetArgs{
			Regions: ssmincidents.ReplicationSetRegionArray{
				&ssmincidents.ReplicationSetRegionArgs{
					Name: pulumi.String("us-west-2"),
				},
				&ssmincidents.ReplicationSetRegionArgs{
					Name: pulumi.String("ap-southeast-2"),
				},
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
import com.pulumi.aws.ssmincidents.ReplicationSet;
import com.pulumi.aws.ssmincidents.ReplicationSetArgs;
import com.pulumi.aws.ssmincidents.inputs.ReplicationSetRegionArgs;
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
        var replicationSetName = new ReplicationSet("replicationSetName", ReplicationSetArgs.builder()        
            .regions(            
                ReplicationSetRegionArgs.builder()
                    .name("us-west-2")
                    .build(),
                ReplicationSetRegionArgs.builder()
                    .name("ap-southeast-2")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  replicationSetName:
    type: aws:ssmincidents:ReplicationSet
    properties:
      regions:
        - name: us-west-2
        - name: ap-southeast-2
```

Delete a Region from a replication set. (You can delete only one Region at a time.)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const replicationSetName = new aws.ssmincidents.ReplicationSet("replicationSetName", {regions: [{
    name: "us-west-2",
}]});
```
```python
import pulumi
import pulumi_aws as aws

replication_set_name = aws.ssmincidents.ReplicationSet("replicationSetName", regions=[aws.ssmincidents.ReplicationSetRegionArgs(
    name="us-west-2",
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var replicationSetName = new Aws.SsmIncidents.ReplicationSet("replicationSetName", new()
    {
        Regions = new[]
        {
            new Aws.SsmIncidents.Inputs.ReplicationSetRegionArgs
            {
                Name = "us-west-2",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ssmincidents.NewReplicationSet(ctx, "replicationSetName", &ssmincidents.ReplicationSetArgs{
			Regions: ssmincidents.ReplicationSetRegionArray{
				&ssmincidents.ReplicationSetRegionArgs{
					Name: pulumi.String("us-west-2"),
				},
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
import com.pulumi.aws.ssmincidents.ReplicationSet;
import com.pulumi.aws.ssmincidents.ReplicationSetArgs;
import com.pulumi.aws.ssmincidents.inputs.ReplicationSetRegionArgs;
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
        var replicationSetName = new ReplicationSet("replicationSetName", ReplicationSetArgs.builder()        
            .regions(ReplicationSetRegionArgs.builder()
                .name("us-west-2")
                .build())
            .build());

    }
}
```
```yaml
resources:
  replicationSetName:
    type: aws:ssmincidents:ReplicationSet
    properties:
      regions:
        - name: us-west-2
```
{{% /example %}}
{{% /examples %}}
## Basic Usage with an AWS Customer Managed Key

Create a replication set with an AWS Key Management Service (AWS KMS) customer manager key:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {});
const replicationSetName = new aws.ssmincidents.ReplicationSet("replicationSetName", {
    regions: [{
        name: "us-west-2",
        kmsKeyArn: exampleKey.arn,
    }],
    tags: {
        exampleTag: "exampleValue",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey")
replication_set_name = aws.ssmincidents.ReplicationSet("replicationSetName",
    regions=[aws.ssmincidents.ReplicationSetRegionArgs(
        name="us-west-2",
        kms_key_arn=example_key.arn,
    )],
    tags={
        "exampleTag": "exampleValue",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey");

    var replicationSetName = new Aws.SsmIncidents.ReplicationSet("replicationSetName", new()
    {
        Regions = new[]
        {
            new Aws.SsmIncidents.Inputs.ReplicationSetRegionArgs
            {
                Name = "us-west-2",
                KmsKeyArn = exampleKey.Arn,
            },
        },
        Tags = 
        {
            { "exampleTag", "exampleValue" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssmincidents"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", nil)
		if err != nil {
			return err
		}
		_, err = ssmincidents.NewReplicationSet(ctx, "replicationSetName", &ssmincidents.ReplicationSetArgs{
			Regions: ssmincidents.ReplicationSetRegionArray{
				&ssmincidents.ReplicationSetRegionArgs{
					Name:      pulumi.String("us-west-2"),
					KmsKeyArn: exampleKey.Arn,
				},
			},
			Tags: pulumi.StringMap{
				"exampleTag": pulumi.String("exampleValue"),
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
import com.pulumi.aws.kms.Key;
import com.pulumi.aws.ssmincidents.ReplicationSet;
import com.pulumi.aws.ssmincidents.ReplicationSetArgs;
import com.pulumi.aws.ssmincidents.inputs.ReplicationSetRegionArgs;
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
        var exampleKey = new Key("exampleKey");

        var replicationSetName = new ReplicationSet("replicationSetName", ReplicationSetArgs.builder()        
            .regions(ReplicationSetRegionArgs.builder()
                .name("us-west-2")
                .kmsKeyArn(exampleKey.arn())
                .build())
            .tags(Map.of("exampleTag", "exampleValue"))
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
  replicationSetName:
    type: aws:ssmincidents:ReplicationSet
    properties:
      regions:
        - name: us-west-2
          kmsKeyArn: ${exampleKey.arn}
      tags:
        exampleTag: exampleValue
```


## Import

Using `pulumi import`, import an Incident Manager replication. For example:

```sh
 $ pulumi import aws:ssmincidents/replicationSet:ReplicationSet replicationSetName import
```
 