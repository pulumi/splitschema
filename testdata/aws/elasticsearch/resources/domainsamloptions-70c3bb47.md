Manages SAML authentication options for an AWS Elasticsearch Domain.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const exampleDomain = new aws.elasticsearch.Domain("exampleDomain", {
    elasticsearchVersion: "1.5",
    clusterConfig: {
        instanceType: "r4.large.elasticsearch",
    },
    snapshotOptions: {
        automatedSnapshotStartHour: 23,
    },
    tags: {
        Domain: "TestDomain",
    },
});
const exampleDomainSamlOptions = new aws.elasticsearch.DomainSamlOptions("exampleDomainSamlOptions", {
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

example_domain = aws.elasticsearch.Domain("exampleDomain",
    elasticsearch_version="1.5",
    cluster_config=aws.elasticsearch.DomainClusterConfigArgs(
        instance_type="r4.large.elasticsearch",
    ),
    snapshot_options=aws.elasticsearch.DomainSnapshotOptionsArgs(
        automated_snapshot_start_hour=23,
    ),
    tags={
        "Domain": "TestDomain",
    })
example_domain_saml_options = aws.elasticsearch.DomainSamlOptions("exampleDomainSamlOptions",
    domain_name=example_domain.domain_name,
    saml_options=aws.elasticsearch.DomainSamlOptionsSamlOptionsArgs(
        enabled=True,
        idp=aws.elasticsearch.DomainSamlOptionsSamlOptionsIdpArgs(
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
    var exampleDomain = new Aws.ElasticSearch.Domain("exampleDomain", new()
    {
        ElasticsearchVersion = "1.5",
        ClusterConfig = new Aws.ElasticSearch.Inputs.DomainClusterConfigArgs
        {
            InstanceType = "r4.large.elasticsearch",
        },
        SnapshotOptions = new Aws.ElasticSearch.Inputs.DomainSnapshotOptionsArgs
        {
            AutomatedSnapshotStartHour = 23,
        },
        Tags = 
        {
            { "Domain", "TestDomain" },
        },
    });

    var exampleDomainSamlOptions = new Aws.ElasticSearch.DomainSamlOptions("exampleDomainSamlOptions", new()
    {
        DomainName = exampleDomain.DomainName,
        SamlOptions = new Aws.ElasticSearch.Inputs.DomainSamlOptionsSamlOptionsArgs
        {
            Enabled = true,
            Idp = new Aws.ElasticSearch.Inputs.DomainSamlOptionsSamlOptionsIdpArgs
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elasticsearch"
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
		exampleDomain, err := elasticsearch.NewDomain(ctx, "exampleDomain", &elasticsearch.DomainArgs{
			ElasticsearchVersion: pulumi.String("1.5"),
			ClusterConfig: &elasticsearch.DomainClusterConfigArgs{
				InstanceType: pulumi.String("r4.large.elasticsearch"),
			},
			SnapshotOptions: &elasticsearch.DomainSnapshotOptionsArgs{
				AutomatedSnapshotStartHour: pulumi.Int(23),
			},
			Tags: pulumi.StringMap{
				"Domain": pulumi.String("TestDomain"),
			},
		})
		if err != nil {
			return err
		}
		_, err = elasticsearch.NewDomainSamlOptions(ctx, "exampleDomainSamlOptions", &elasticsearch.DomainSamlOptionsArgs{
			DomainName: exampleDomain.DomainName,
			SamlOptions: &elasticsearch.DomainSamlOptionsSamlOptionsArgs{
				Enabled: pulumi.Bool(true),
				Idp: &elasticsearch.DomainSamlOptionsSamlOptionsIdpArgs{
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
import com.pulumi.aws.elasticsearch.Domain;
import com.pulumi.aws.elasticsearch.DomainArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainClusterConfigArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainSnapshotOptionsArgs;
import com.pulumi.aws.elasticsearch.DomainSamlOptions;
import com.pulumi.aws.elasticsearch.DomainSamlOptionsArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainSamlOptionsSamlOptionsArgs;
import com.pulumi.aws.elasticsearch.inputs.DomainSamlOptionsSamlOptionsIdpArgs;
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
            .elasticsearchVersion("1.5")
            .clusterConfig(DomainClusterConfigArgs.builder()
                .instanceType("r4.large.elasticsearch")
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
    type: aws:elasticsearch:Domain
    properties:
      elasticsearchVersion: '1.5'
      clusterConfig:
        instanceType: r4.large.elasticsearch
      snapshotOptions:
        automatedSnapshotStartHour: 23
      tags:
        Domain: TestDomain
  exampleDomainSamlOptions:
    type: aws:elasticsearch:DomainSamlOptions
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

Using `pulumi import`, import Elasticsearch domains using the `domain_name`. For example:

```sh
 $ pulumi import aws:elasticsearch/domainSamlOptions:DomainSamlOptions example domain_name
```
 