Manages an Access Analyzer Analyzer. More information can be found in the [Access Analyzer User Guide](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Account Analyzer

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.accessanalyzer.Analyzer("example", {analyzerName: "example"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.accessanalyzer.Analyzer("example", analyzer_name="example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.AccessAnalyzer.Analyzer("example", new()
    {
        AnalyzerName = "example",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/accessanalyzer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := accessanalyzer.NewAnalyzer(ctx, "example", &accessanalyzer.AnalyzerArgs{
			AnalyzerName: pulumi.String("example"),
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
import com.pulumi.aws.accessanalyzer.Analyzer;
import com.pulumi.aws.accessanalyzer.AnalyzerArgs;
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
        var example = new Analyzer("example", AnalyzerArgs.builder()        
            .analyzerName("example")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:accessanalyzer:Analyzer
    properties:
      analyzerName: example
```
{{% /example %}}
{{% example %}}
### Organization Analyzer

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleOrganization = new aws.organizations.Organization("exampleOrganization", {awsServiceAccessPrincipals: ["access-analyzer.amazonaws.com"]});
const exampleAnalyzer = new aws.accessanalyzer.Analyzer("exampleAnalyzer", {
    analyzerName: "example",
    type: "ORGANIZATION",
}, {
    dependsOn: [exampleOrganization],
});
```
```python
import pulumi
import pulumi_aws as aws

example_organization = aws.organizations.Organization("exampleOrganization", aws_service_access_principals=["access-analyzer.amazonaws.com"])
example_analyzer = aws.accessanalyzer.Analyzer("exampleAnalyzer",
    analyzer_name="example",
    type="ORGANIZATION",
    opts=pulumi.ResourceOptions(depends_on=[example_organization]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleOrganization = new Aws.Organizations.Organization("exampleOrganization", new()
    {
        AwsServiceAccessPrincipals = new[]
        {
            "access-analyzer.amazonaws.com",
        },
    });

    var exampleAnalyzer = new Aws.AccessAnalyzer.Analyzer("exampleAnalyzer", new()
    {
        AnalyzerName = "example",
        Type = "ORGANIZATION",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleOrganization,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/accessanalyzer"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/organizations"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleOrganization, err := organizations.NewOrganization(ctx, "exampleOrganization", &organizations.OrganizationArgs{
			AwsServiceAccessPrincipals: pulumi.StringArray{
				pulumi.String("access-analyzer.amazonaws.com"),
			},
		})
		if err != nil {
			return err
		}
		_, err = accessanalyzer.NewAnalyzer(ctx, "exampleAnalyzer", &accessanalyzer.AnalyzerArgs{
			AnalyzerName: pulumi.String("example"),
			Type:         pulumi.String("ORGANIZATION"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleOrganization,
		}))
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
import com.pulumi.aws.organizations.Organization;
import com.pulumi.aws.organizations.OrganizationArgs;
import com.pulumi.aws.accessanalyzer.Analyzer;
import com.pulumi.aws.accessanalyzer.AnalyzerArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleOrganization = new Organization("exampleOrganization", OrganizationArgs.builder()        
            .awsServiceAccessPrincipals("access-analyzer.amazonaws.com")
            .build());

        var exampleAnalyzer = new Analyzer("exampleAnalyzer", AnalyzerArgs.builder()        
            .analyzerName("example")
            .type("ORGANIZATION")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleOrganization)
                .build());

    }
}
```
```yaml
resources:
  exampleOrganization:
    type: aws:organizations:Organization
    properties:
      awsServiceAccessPrincipals:
        - access-analyzer.amazonaws.com
  exampleAnalyzer:
    type: aws:accessanalyzer:Analyzer
    properties:
      analyzerName: example
      type: ORGANIZATION
    options:
      dependson:
        - ${exampleOrganization}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Access Analyzer Analyzers using the `analyzer_name`. For example:

```sh
 $ pulumi import aws:accessanalyzer/analyzer:Analyzer example example
```
 