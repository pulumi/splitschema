Provides an Elastic Container Registry Replication Configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const exampleRegions = aws.getRegions({});
const exampleReplicationConfiguration = new aws.ecr.ReplicationConfiguration("exampleReplicationConfiguration", {replicationConfiguration: {
    rules: [{
        destinations: [{
            region: exampleRegions.then(exampleRegions => exampleRegions.names?.[0]),
            registryId: current.then(current => current.accountId),
        }],
    }],
}});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
example_regions = aws.get_regions()
example_replication_configuration = aws.ecr.ReplicationConfiguration("exampleReplicationConfiguration", replication_configuration=aws.ecr.ReplicationConfigurationReplicationConfigurationArgs(
    rules=[aws.ecr.ReplicationConfigurationReplicationConfigurationRuleArgs(
        destinations=[aws.ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs(
            region=example_regions.names[0],
            registry_id=current.account_id,
        )],
    )],
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var exampleRegions = Aws.GetRegions.Invoke();

    var exampleReplicationConfiguration = new Aws.Ecr.ReplicationConfiguration("exampleReplicationConfiguration", new()
    {
        ReplicationConfigurationDetails = new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationArgs
        {
            Rules = new[]
            {
                new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationRuleArgs
                {
                    Destinations = new[]
                    {
                        new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs
                        {
                            Region = exampleRegions.Apply(getRegionsResult => getRegionsResult.Names[0]),
                            RegistryId = current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleRegions, err := aws.GetRegions(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = ecr.NewReplicationConfiguration(ctx, "exampleReplicationConfiguration", &ecr.ReplicationConfigurationArgs{
			ReplicationConfiguration: &ecr.ReplicationConfigurationReplicationConfigurationArgs{
				Rules: ecr.ReplicationConfigurationReplicationConfigurationRuleArray{
					&ecr.ReplicationConfigurationReplicationConfigurationRuleArgs{
						Destinations: ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArray{
							&ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs{
								Region:     *pulumi.String(exampleRegions.Names[0]),
								RegistryId: *pulumi.String(current.AccountId),
							},
						},
					},
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetRegionsArgs;
import com.pulumi.aws.ecr.ReplicationConfiguration;
import com.pulumi.aws.ecr.ReplicationConfigurationArgs;
import com.pulumi.aws.ecr.inputs.ReplicationConfigurationReplicationConfigurationArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        final var exampleRegions = AwsFunctions.getRegions();

        var exampleReplicationConfiguration = new ReplicationConfiguration("exampleReplicationConfiguration", ReplicationConfigurationArgs.builder()        
            .replicationConfiguration(ReplicationConfigurationReplicationConfigurationArgs.builder()
                .rules(ReplicationConfigurationReplicationConfigurationRuleArgs.builder()
                    .destinations(ReplicationConfigurationReplicationConfigurationRuleDestinationArgs.builder()
                        .region(exampleRegions.applyValue(getRegionsResult -> getRegionsResult.names()[0]))
                        .registryId(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleReplicationConfiguration:
    type: aws:ecr:ReplicationConfiguration
    properties:
      replicationConfiguration:
        rules:
          - destinations:
              - region: ${exampleRegions.names[0]}
                registryId: ${current.accountId}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  exampleRegions:
    fn::invoke:
      Function: aws:getRegions
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}
## Multiple Region Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const exampleRegions = aws.getRegions({});
const exampleReplicationConfiguration = new aws.ecr.ReplicationConfiguration("exampleReplicationConfiguration", {replicationConfiguration: {
    rules: [{
        destinations: [
            {
                region: exampleRegions.then(exampleRegions => exampleRegions.names?.[0]),
                registryId: current.then(current => current.accountId),
            },
            {
                region: exampleRegions.then(exampleRegions => exampleRegions.names?.[1]),
                registryId: current.then(current => current.accountId),
            },
        ],
    }],
}});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
example_regions = aws.get_regions()
example_replication_configuration = aws.ecr.ReplicationConfiguration("exampleReplicationConfiguration", replication_configuration=aws.ecr.ReplicationConfigurationReplicationConfigurationArgs(
    rules=[aws.ecr.ReplicationConfigurationReplicationConfigurationRuleArgs(
        destinations=[
            aws.ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs(
                region=example_regions.names[0],
                registry_id=current.account_id,
            ),
            aws.ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs(
                region=example_regions.names[1],
                registry_id=current.account_id,
            ),
        ],
    )],
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var exampleRegions = Aws.GetRegions.Invoke();

    var exampleReplicationConfiguration = new Aws.Ecr.ReplicationConfiguration("exampleReplicationConfiguration", new()
    {
        ReplicationConfigurationDetails = new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationArgs
        {
            Rules = new[]
            {
                new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationRuleArgs
                {
                    Destinations = new[]
                    {
                        new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs
                        {
                            Region = exampleRegions.Apply(getRegionsResult => getRegionsResult.Names[0]),
                            RegistryId = current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
                        },
                        new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs
                        {
                            Region = exampleRegions.Apply(getRegionsResult => getRegionsResult.Names[1]),
                            RegistryId = current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleRegions, err := aws.GetRegions(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = ecr.NewReplicationConfiguration(ctx, "exampleReplicationConfiguration", &ecr.ReplicationConfigurationArgs{
			ReplicationConfiguration: &ecr.ReplicationConfigurationReplicationConfigurationArgs{
				Rules: ecr.ReplicationConfigurationReplicationConfigurationRuleArray{
					&ecr.ReplicationConfigurationReplicationConfigurationRuleArgs{
						Destinations: ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArray{
							&ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs{
								Region:     *pulumi.String(exampleRegions.Names[0]),
								RegistryId: *pulumi.String(current.AccountId),
							},
							&ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs{
								Region:     *pulumi.String(exampleRegions.Names[1]),
								RegistryId: *pulumi.String(current.AccountId),
							},
						},
					},
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetRegionsArgs;
import com.pulumi.aws.ecr.ReplicationConfiguration;
import com.pulumi.aws.ecr.ReplicationConfigurationArgs;
import com.pulumi.aws.ecr.inputs.ReplicationConfigurationReplicationConfigurationArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        final var exampleRegions = AwsFunctions.getRegions();

        var exampleReplicationConfiguration = new ReplicationConfiguration("exampleReplicationConfiguration", ReplicationConfigurationArgs.builder()        
            .replicationConfiguration(ReplicationConfigurationReplicationConfigurationArgs.builder()
                .rules(ReplicationConfigurationReplicationConfigurationRuleArgs.builder()
                    .destinations(                    
                        ReplicationConfigurationReplicationConfigurationRuleDestinationArgs.builder()
                            .region(exampleRegions.applyValue(getRegionsResult -> getRegionsResult.names()[0]))
                            .registryId(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                            .build(),
                        ReplicationConfigurationReplicationConfigurationRuleDestinationArgs.builder()
                            .region(exampleRegions.applyValue(getRegionsResult -> getRegionsResult.names()[1]))
                            .registryId(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                            .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleReplicationConfiguration:
    type: aws:ecr:ReplicationConfiguration
    properties:
      replicationConfiguration:
        rules:
          - destinations:
              - region: ${exampleRegions.names[0]}
                registryId: ${current.accountId}
              - region: ${exampleRegions.names[1]}
                registryId: ${current.accountId}
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  exampleRegions:
    fn::invoke:
      Function: aws:getRegions
      Arguments: {}
```

## Repository Filter Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const current = aws.getCallerIdentity({});
const exampleRegions = aws.getRegions({});
const exampleReplicationConfiguration = new aws.ecr.ReplicationConfiguration("exampleReplicationConfiguration", {replicationConfiguration: {
    rules: [{
        destinations: [{
            region: exampleRegions.then(exampleRegions => exampleRegions.names?.[0]),
            registryId: current.then(current => current.accountId),
        }],
        repositoryFilters: [{
            filter: "prod-microservice",
            filterType: "PREFIX_MATCH",
        }],
    }],
}});
```
```python
import pulumi
import pulumi_aws as aws

current = aws.get_caller_identity()
example_regions = aws.get_regions()
example_replication_configuration = aws.ecr.ReplicationConfiguration("exampleReplicationConfiguration", replication_configuration=aws.ecr.ReplicationConfigurationReplicationConfigurationArgs(
    rules=[aws.ecr.ReplicationConfigurationReplicationConfigurationRuleArgs(
        destinations=[aws.ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs(
            region=example_regions.names[0],
            registry_id=current.account_id,
        )],
        repository_filters=[aws.ecr.ReplicationConfigurationReplicationConfigurationRuleRepositoryFilterArgs(
            filter="prod-microservice",
            filter_type="PREFIX_MATCH",
        )],
    )],
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var current = Aws.GetCallerIdentity.Invoke();

    var exampleRegions = Aws.GetRegions.Invoke();

    var exampleReplicationConfiguration = new Aws.Ecr.ReplicationConfiguration("exampleReplicationConfiguration", new()
    {
        ReplicationConfigurationDetails = new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationArgs
        {
            Rules = new[]
            {
                new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationRuleArgs
                {
                    Destinations = new[]
                    {
                        new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs
                        {
                            Region = exampleRegions.Apply(getRegionsResult => getRegionsResult.Names[0]),
                            RegistryId = current.Apply(getCallerIdentityResult => getCallerIdentityResult.AccountId),
                        },
                    },
                    RepositoryFilters = new[]
                    {
                        new Aws.Ecr.Inputs.ReplicationConfigurationReplicationConfigurationRuleRepositoryFilterArgs
                        {
                            Filter = "prod-microservice",
                            FilterType = "PREFIX_MATCH",
                        },
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		current, err := aws.GetCallerIdentity(ctx, nil, nil)
		if err != nil {
			return err
		}
		exampleRegions, err := aws.GetRegions(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = ecr.NewReplicationConfiguration(ctx, "exampleReplicationConfiguration", &ecr.ReplicationConfigurationArgs{
			ReplicationConfiguration: &ecr.ReplicationConfigurationReplicationConfigurationArgs{
				Rules: ecr.ReplicationConfigurationReplicationConfigurationRuleArray{
					&ecr.ReplicationConfigurationReplicationConfigurationRuleArgs{
						Destinations: ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArray{
							&ecr.ReplicationConfigurationReplicationConfigurationRuleDestinationArgs{
								Region:     *pulumi.String(exampleRegions.Names[0]),
								RegistryId: *pulumi.String(current.AccountId),
							},
						},
						RepositoryFilters: ecr.ReplicationConfigurationReplicationConfigurationRuleRepositoryFilterArray{
							&ecr.ReplicationConfigurationReplicationConfigurationRuleRepositoryFilterArgs{
								Filter:     pulumi.String("prod-microservice"),
								FilterType: pulumi.String("PREFIX_MATCH"),
							},
						},
					},
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
import com.pulumi.aws.AwsFunctions;
import com.pulumi.aws.inputs.GetCallerIdentityArgs;
import com.pulumi.aws.inputs.GetRegionsArgs;
import com.pulumi.aws.ecr.ReplicationConfiguration;
import com.pulumi.aws.ecr.ReplicationConfigurationArgs;
import com.pulumi.aws.ecr.inputs.ReplicationConfigurationReplicationConfigurationArgs;
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
        final var current = AwsFunctions.getCallerIdentity();

        final var exampleRegions = AwsFunctions.getRegions();

        var exampleReplicationConfiguration = new ReplicationConfiguration("exampleReplicationConfiguration", ReplicationConfigurationArgs.builder()        
            .replicationConfiguration(ReplicationConfigurationReplicationConfigurationArgs.builder()
                .rules(ReplicationConfigurationReplicationConfigurationRuleArgs.builder()
                    .destinations(ReplicationConfigurationReplicationConfigurationRuleDestinationArgs.builder()
                        .region(exampleRegions.applyValue(getRegionsResult -> getRegionsResult.names()[0]))
                        .registryId(current.applyValue(getCallerIdentityResult -> getCallerIdentityResult.accountId()))
                        .build())
                    .repositoryFilters(ReplicationConfigurationReplicationConfigurationRuleRepositoryFilterArgs.builder()
                        .filter("prod-microservice")
                        .filterType("PREFIX_MATCH")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleReplicationConfiguration:
    type: aws:ecr:ReplicationConfiguration
    properties:
      replicationConfiguration:
        rules:
          - destinations:
              - region: ${exampleRegions.names[0]}
                registryId: ${current.accountId}
            repositoryFilters:
              - filter: prod-microservice
                filterType: PREFIX_MATCH
variables:
  current:
    fn::invoke:
      Function: aws:getCallerIdentity
      Arguments: {}
  exampleRegions:
    fn::invoke:
      Function: aws:getRegions
      Arguments: {}
```


## Import

Using `pulumi import`, import ECR Replication Configuration using the `registry_id`. For example:

```sh
 $ pulumi import aws:ecr/replicationConfiguration:ReplicationConfiguration service 012345678912
```
 