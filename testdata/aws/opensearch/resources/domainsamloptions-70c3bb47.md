Manages SAML authentication options for an AWS OpenSearch Domain.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const exampleDomain = new aws.opensearch.Domain("exampleDomain", {
    engineVersion: "OpenSearch_1.1",
    clusterConfig: {
        instanceType: "r4.large.search",
    },
    snapshotOptions: {
        automatedSnapshotStartHour: 23,
    },
    tags: {
        Domain: "TestDomain",
    },
});
const exampleDomainSamlOptions = new aws.opensearch.DomainSamlOptions("exampleDomainSamlOptions", {
    domainName: exampleDomain.domainName,
    samlOptions: {
        enabled: true,
        idp: {
            entityId: "https://example.com",
            metadataContent: fs.readFileSync("./saml-metadata.xml", "utf8"),
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_domain = aws.opensearch.Domain("exampleDomain",
    engine_version="OpenSearch_1.1",
    cluster_config=aws.opensearch.DomainClusterConfigArgs(
        instance_type="r4.large.search",
    ),
    snapshot_options=aws.opensearch.DomainSnapshotOptionsArgs(
        automated_snapshot_start_hour=23,
    ),
    tags={
        "Domain": "TestDomain",
    })
example_domain_saml_options = aws.opensearch.DomainSamlOptions("exampleDomainSamlOptions",
    domain_name=example_domain.domain_name,
    saml_options=aws.opensearch.DomainSamlOptionsSamlOptionsArgs(
        enabled=True,
        idp=aws.opensearch.DomainSamlOptionsSamlOptionsIdpArgs(
            entity_id="https://example.com",
            metadata_content=(lambda path: open(path).read())("./saml-metadata.xml"),
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleDomain = new Aws.OpenSearch.Domain("exampleDomain", new()
    {
        EngineVersion = "OpenSearch_1.1",
        ClusterConfig = new Aws.OpenSearch.Inputs.DomainClusterConfigArgs
        {
            InstanceType = "r4.large.search",
        },
        SnapshotOptions = new Aws.OpenSearch.Inputs.DomainSnapshotOptionsArgs
        {
            AutomatedSnapshotStartHour = 23,
        },
        Tags = 
        {
            { "Domain", "TestDomain" },
        },
    });

    var exampleDomainSamlOptions = new Aws.OpenSearch.DomainSamlOptions("exampleDomainSamlOptions", new()
    {
        DomainName = exampleDomain.DomainName,
        SamlOptions = new Aws.OpenSearch.Inputs.DomainSamlOptionsSamlOptionsArgs
        {
            Enabled = true,
            Idp = new Aws.OpenSearch.Inputs.DomainSamlOptionsSamlOptionsIdpArgs
            {
                EntityId = "https://example.com",
                MetadataContent = File.ReadAllText("./saml-metadata.xml"),
            },
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleDomain, err := opensearch.NewDomain(ctx, "exampleDomain", &opensearch.DomainArgs{
			EngineVersion: pulumi.String("OpenSearch_1.1"),
			ClusterConfig: &opensearch.DomainClusterConfigArgs{
				InstanceType: pulumi.String("r4.large.search"),
			},
			SnapshotOptions: &opensearch.DomainSnapshotOptionsArgs{
				AutomatedSnapshotStartHour: pulumi.Int(23),
			},
			Tags: pulumi.StringMap{
				"Domain": pulumi.String("TestDomain"),
			},
		})
		if err != nil {
			return err
		}
		_, err = opensearch.NewDomainSamlOptions(ctx, "exampleDomainSamlOptions", &opensearch.DomainSamlOptionsArgs{
			DomainName: exampleDomain.DomainName,
			SamlOptions: &opensearch.DomainSamlOptionsSamlOptionsArgs{
				Enabled: pulumi.Bool(true),
				Idp: &opensearch.DomainSamlOptionsSamlOptionsIdpArgs{
					EntityId:        pulumi.String("https://example.com"),
					MetadataContent: readFileOrPanic("./saml-metadata.xml"),
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
import com.pulumi.aws.opensearch.Domain;
import com.pulumi.aws.opensearch.DomainArgs;
import com.pulumi.aws.opensearch.inputs.DomainClusterConfigArgs;
import com.pulumi.aws.opensearch.inputs.DomainSnapshotOptionsArgs;
import com.pulumi.aws.opensearch.DomainSamlOptions;
import com.pulumi.aws.opensearch.DomainSamlOptionsArgs;
import com.pulumi.aws.opensearch.inputs.DomainSamlOptionsSamlOptionsArgs;
import com.pulumi.aws.opensearch.inputs.DomainSamlOptionsSamlOptionsIdpArgs;
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
        var exampleDomain = new Domain("exampleDomain", DomainArgs.builder()        
            .engineVersion("OpenSearch_1.1")
            .clusterConfig(DomainClusterConfigArgs.builder()
                .instanceType("r4.large.search")
                .build())
            .snapshotOptions(DomainSnapshotOptionsArgs.builder()
                .automatedSnapshotStartHour(23)
                .build())
            .tags(Map.of("Domain", "TestDomain"))
            .build());

        var exampleDomainSamlOptions = new DomainSamlOptions("exampleDomainSamlOptions", DomainSamlOptionsArgs.builder()        
            .domainName(exampleDomain.domainName())
            .samlOptions(DomainSamlOptionsSamlOptionsArgs.builder()
                .enabled(true)
                .idp(DomainSamlOptionsSamlOptionsIdpArgs.builder()
                    .entityId("https://example.com")
                    .metadataContent(Files.readString(Paths.get("./saml-metadata.xml")))
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleDomain:
    type: aws:opensearch:Domain
    properties:
      engineVersion: OpenSearch_1.1
      clusterConfig:
        instanceType: r4.large.search
      snapshotOptions:
        automatedSnapshotStartHour: 23
      tags:
        Domain: TestDomain
  exampleDomainSamlOptions:
    type: aws:opensearch:DomainSamlOptions
    properties:
      domainName: ${exampleDomain.domainName}
      samlOptions:
        enabled: true
        idp:
          entityId: https://example.com
          metadataContent:
            fn::readFile: ./saml-metadata.xml
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import OpenSearch domains using the `domain_name`. For example:

```sh
 $ pulumi import aws:opensearch/domainSamlOptions:DomainSamlOptions example domain_name
```
 