Provides an CloudSearch domain resource.

The provider waits for the domain to become `Active` when applying a configuration.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudsearch.Domain("example", {
    indexFields: [
        {
            analysisScheme: "_en_default_",
            highlight: false,
            name: "headline",
            "return": true,
            search: true,
            sort: true,
            type: "text",
        },
        {
            facet: true,
            name: "price",
            "return": true,
            search: true,
            sort: true,
            sourceFields: "headline",
            type: "double",
        },
    ],
    scalingParameters: {
        desiredInstanceType: "search.medium",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudsearch.Domain("example",
    index_fields=[
        aws.cloudsearch.DomainIndexFieldArgs(
            analysis_scheme="_en_default_",
            highlight=False,
            name="headline",
            return_=True,
            search=True,
            sort=True,
            type="text",
        ),
        aws.cloudsearch.DomainIndexFieldArgs(
            facet=True,
            name="price",
            return_=True,
            search=True,
            sort=True,
            source_fields="headline",
            type="double",
        ),
    ],
    scaling_parameters=aws.cloudsearch.DomainScalingParametersArgs(
        desired_instance_type="search.medium",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudSearch.Domain("example", new()
    {
        IndexFields = new[]
        {
            new Aws.CloudSearch.Inputs.DomainIndexFieldArgs
            {
                AnalysisScheme = "_en_default_",
                Highlight = false,
                Name = "headline",
                Return = true,
                Search = true,
                Sort = true,
                Type = "text",
            },
            new Aws.CloudSearch.Inputs.DomainIndexFieldArgs
            {
                Facet = true,
                Name = "price",
                Return = true,
                Search = true,
                Sort = true,
                SourceFields = "headline",
                Type = "double",
            },
        },
        ScalingParameters = new Aws.CloudSearch.Inputs.DomainScalingParametersArgs
        {
            DesiredInstanceType = "search.medium",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudsearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudsearch.NewDomain(ctx, "example", &cloudsearch.DomainArgs{
			IndexFields: cloudsearch.DomainIndexFieldArray{
				&cloudsearch.DomainIndexFieldArgs{
					AnalysisScheme: pulumi.String("_en_default_"),
					Highlight:      pulumi.Bool(false),
					Name:           pulumi.String("headline"),
					Return:         pulumi.Bool(true),
					Search:         pulumi.Bool(true),
					Sort:           pulumi.Bool(true),
					Type:           pulumi.String("text"),
				},
				&cloudsearch.DomainIndexFieldArgs{
					Facet:        pulumi.Bool(true),
					Name:         pulumi.String("price"),
					Return:       pulumi.Bool(true),
					Search:       pulumi.Bool(true),
					Sort:         pulumi.Bool(true),
					SourceFields: pulumi.String("headline"),
					Type:         pulumi.String("double"),
				},
			},
			ScalingParameters: &cloudsearch.DomainScalingParametersArgs{
				DesiredInstanceType: pulumi.String("search.medium"),
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
import com.pulumi.aws.cloudsearch.Domain;
import com.pulumi.aws.cloudsearch.DomainArgs;
import com.pulumi.aws.cloudsearch.inputs.DomainIndexFieldArgs;
import com.pulumi.aws.cloudsearch.inputs.DomainScalingParametersArgs;
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
        var example = new Domain("example", DomainArgs.builder()        
            .indexFields(            
                DomainIndexFieldArgs.builder()
                    .analysisScheme("_en_default_")
                    .highlight(false)
                    .name("headline")
                    .return_(true)
                    .search(true)
                    .sort(true)
                    .type("text")
                    .build(),
                DomainIndexFieldArgs.builder()
                    .facet(true)
                    .name("price")
                    .return_(true)
                    .search(true)
                    .sort(true)
                    .sourceFields("headline")
                    .type("double")
                    .build())
            .scalingParameters(DomainScalingParametersArgs.builder()
                .desiredInstanceType("search.medium")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudsearch:Domain
    properties:
      indexFields:
        - analysisScheme: _en_default_
          highlight: false
          name: headline
          return: true
          search: true
          sort: true
          type: text
        - facet: true
          name: price
          return: true
          search: true
          sort: true
          sourceFields: headline
          type: double
      scalingParameters:
        desiredInstanceType: search.medium
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import CloudSearch Domains using the `name`. For example:

```sh
 $ pulumi import aws:cloudsearch/domain:Domain example example-domain
```
 