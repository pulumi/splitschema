Provides an Elastic Container Registry Scanning Configuration. Can't be completely deleted, instead reverts to the default `BASIC` scanning configuration without rules.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const configuration = new aws.ecr.RegistryScanningConfiguration("configuration", {
    rules: [{
        repositoryFilters: [{
            filter: "example",
            filterType: "WILDCARD",
        }],
        scanFrequency: "CONTINUOUS_SCAN",
    }],
    scanType: "ENHANCED",
});
```
```python
import pulumi
import pulumi_aws as aws

configuration = aws.ecr.RegistryScanningConfiguration("configuration",
    rules=[aws.ecr.RegistryScanningConfigurationRuleArgs(
        repository_filters=[aws.ecr.RegistryScanningConfigurationRuleRepositoryFilterArgs(
            filter="example",
            filter_type="WILDCARD",
        )],
        scan_frequency="CONTINUOUS_SCAN",
    )],
    scan_type="ENHANCED")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var configuration = new Aws.Ecr.RegistryScanningConfiguration("configuration", new()
    {
        Rules = new[]
        {
            new Aws.Ecr.Inputs.RegistryScanningConfigurationRuleArgs
            {
                RepositoryFilters = new[]
                {
                    new Aws.Ecr.Inputs.RegistryScanningConfigurationRuleRepositoryFilterArgs
                    {
                        Filter = "example",
                        FilterType = "WILDCARD",
                    },
                },
                ScanFrequency = "CONTINUOUS_SCAN",
            },
        },
        ScanType = "ENHANCED",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecr.NewRegistryScanningConfiguration(ctx, "configuration", &ecr.RegistryScanningConfigurationArgs{
			Rules: ecr.RegistryScanningConfigurationRuleArray{
				&ecr.RegistryScanningConfigurationRuleArgs{
					RepositoryFilters: ecr.RegistryScanningConfigurationRuleRepositoryFilterArray{
						&ecr.RegistryScanningConfigurationRuleRepositoryFilterArgs{
							Filter:     pulumi.String("example"),
							FilterType: pulumi.String("WILDCARD"),
						},
					},
					ScanFrequency: pulumi.String("CONTINUOUS_SCAN"),
				},
			},
			ScanType: pulumi.String("ENHANCED"),
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
import com.pulumi.aws.ecr.RegistryScanningConfiguration;
import com.pulumi.aws.ecr.RegistryScanningConfigurationArgs;
import com.pulumi.aws.ecr.inputs.RegistryScanningConfigurationRuleArgs;
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
        var configuration = new RegistryScanningConfiguration("configuration", RegistryScanningConfigurationArgs.builder()        
            .rules(RegistryScanningConfigurationRuleArgs.builder()
                .repositoryFilters(RegistryScanningConfigurationRuleRepositoryFilterArgs.builder()
                    .filter("example")
                    .filterType("WILDCARD")
                    .build())
                .scanFrequency("CONTINUOUS_SCAN")
                .build())
            .scanType("ENHANCED")
            .build());

    }
}
```
```yaml
resources:
  configuration:
    type: aws:ecr:RegistryScanningConfiguration
    properties:
      rules:
        - repositoryFilters:
            - filter: example
              filterType: WILDCARD
          scanFrequency: CONTINUOUS_SCAN
      scanType: ENHANCED
```
{{% /example %}}
{{% example %}}
### Multiple rules

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ecr.RegistryScanningConfiguration("test", {
    rules: [
        {
            repositoryFilters: [{
                filter: "*",
                filterType: "WILDCARD",
            }],
            scanFrequency: "SCAN_ON_PUSH",
        },
        {
            repositoryFilters: [{
                filter: "example",
                filterType: "WILDCARD",
            }],
            scanFrequency: "CONTINUOUS_SCAN",
        },
    ],
    scanType: "ENHANCED",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ecr.RegistryScanningConfiguration("test",
    rules=[
        aws.ecr.RegistryScanningConfigurationRuleArgs(
            repository_filters=[aws.ecr.RegistryScanningConfigurationRuleRepositoryFilterArgs(
                filter="*",
                filter_type="WILDCARD",
            )],
            scan_frequency="SCAN_ON_PUSH",
        ),
        aws.ecr.RegistryScanningConfigurationRuleArgs(
            repository_filters=[aws.ecr.RegistryScanningConfigurationRuleRepositoryFilterArgs(
                filter="example",
                filter_type="WILDCARD",
            )],
            scan_frequency="CONTINUOUS_SCAN",
        ),
    ],
    scan_type="ENHANCED")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ecr.RegistryScanningConfiguration("test", new()
    {
        Rules = new[]
        {
            new Aws.Ecr.Inputs.RegistryScanningConfigurationRuleArgs
            {
                RepositoryFilters = new[]
                {
                    new Aws.Ecr.Inputs.RegistryScanningConfigurationRuleRepositoryFilterArgs
                    {
                        Filter = "*",
                        FilterType = "WILDCARD",
                    },
                },
                ScanFrequency = "SCAN_ON_PUSH",
            },
            new Aws.Ecr.Inputs.RegistryScanningConfigurationRuleArgs
            {
                RepositoryFilters = new[]
                {
                    new Aws.Ecr.Inputs.RegistryScanningConfigurationRuleRepositoryFilterArgs
                    {
                        Filter = "example",
                        FilterType = "WILDCARD",
                    },
                },
                ScanFrequency = "CONTINUOUS_SCAN",
            },
        },
        ScanType = "ENHANCED",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecr.NewRegistryScanningConfiguration(ctx, "test", &ecr.RegistryScanningConfigurationArgs{
			Rules: ecr.RegistryScanningConfigurationRuleArray{
				&ecr.RegistryScanningConfigurationRuleArgs{
					RepositoryFilters: ecr.RegistryScanningConfigurationRuleRepositoryFilterArray{
						&ecr.RegistryScanningConfigurationRuleRepositoryFilterArgs{
							Filter:     pulumi.String("*"),
							FilterType: pulumi.String("WILDCARD"),
						},
					},
					ScanFrequency: pulumi.String("SCAN_ON_PUSH"),
				},
				&ecr.RegistryScanningConfigurationRuleArgs{
					RepositoryFilters: ecr.RegistryScanningConfigurationRuleRepositoryFilterArray{
						&ecr.RegistryScanningConfigurationRuleRepositoryFilterArgs{
							Filter:     pulumi.String("example"),
							FilterType: pulumi.String("WILDCARD"),
						},
					},
					ScanFrequency: pulumi.String("CONTINUOUS_SCAN"),
				},
			},
			ScanType: pulumi.String("ENHANCED"),
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
import com.pulumi.aws.ecr.RegistryScanningConfiguration;
import com.pulumi.aws.ecr.RegistryScanningConfigurationArgs;
import com.pulumi.aws.ecr.inputs.RegistryScanningConfigurationRuleArgs;
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
        var test = new RegistryScanningConfiguration("test", RegistryScanningConfigurationArgs.builder()        
            .rules(            
                RegistryScanningConfigurationRuleArgs.builder()
                    .repositoryFilters(RegistryScanningConfigurationRuleRepositoryFilterArgs.builder()
                        .filter("*")
                        .filterType("WILDCARD")
                        .build())
                    .scanFrequency("SCAN_ON_PUSH")
                    .build(),
                RegistryScanningConfigurationRuleArgs.builder()
                    .repositoryFilters(RegistryScanningConfigurationRuleRepositoryFilterArgs.builder()
                        .filter("example")
                        .filterType("WILDCARD")
                        .build())
                    .scanFrequency("CONTINUOUS_SCAN")
                    .build())
            .scanType("ENHANCED")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ecr:RegistryScanningConfiguration
    properties:
      rules:
        - repositoryFilters:
            - filter: '*'
              filterType: WILDCARD
          scanFrequency: SCAN_ON_PUSH
        - repositoryFilters:
            - filter: example
              filterType: WILDCARD
          scanFrequency: CONTINUOUS_SCAN
      scanType: ENHANCED
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECR Scanning Configurations using the `registry_id`. For example:

```sh
 $ pulumi import aws:ecr/registryScanningConfiguration:RegistryScanningConfiguration example 012345678901
```
 