Manages an AWS Opensearch Package Association.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myDomain = new aws.opensearch.Domain("myDomain", {
    engineVersion: "Elasticsearch_7.10",
    clusterConfig: {
        instanceType: "r4.large.search",
    },
});
const examplePackage = new aws.opensearch.Package("examplePackage", {
    packageName: "example-txt",
    packageSource: {
        s3BucketName: aws_s3_bucket.my_opensearch_packages.bucket,
        s3Key: aws_s3_object.example.key,
    },
    packageType: "TXT-DICTIONARY",
});
const examplePackageAssociation = new aws.opensearch.PackageAssociation("examplePackageAssociation", {
    packageId: examplePackage.id,
    domainName: myDomain.domainName,
});
```
```python
import pulumi
import pulumi_aws as aws

my_domain = aws.opensearch.Domain("myDomain",
    engine_version="Elasticsearch_7.10",
    cluster_config=aws.opensearch.DomainClusterConfigArgs(
        instance_type="r4.large.search",
    ))
example_package = aws.opensearch.Package("examplePackage",
    package_name="example-txt",
    package_source=aws.opensearch.PackagePackageSourceArgs(
        s3_bucket_name=aws_s3_bucket["my_opensearch_packages"]["bucket"],
        s3_key=aws_s3_object["example"]["key"],
    ),
    package_type="TXT-DICTIONARY")
example_package_association = aws.opensearch.PackageAssociation("examplePackageAssociation",
    package_id=example_package.id,
    domain_name=my_domain.domain_name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myDomain = new Aws.OpenSearch.Domain("myDomain", new()
    {
        EngineVersion = "Elasticsearch_7.10",
        ClusterConfig = new Aws.OpenSearch.Inputs.DomainClusterConfigArgs
        {
            InstanceType = "r4.large.search",
        },
    });

    var examplePackage = new Aws.OpenSearch.Package("examplePackage", new()
    {
        PackageName = "example-txt",
        PackageSource = new Aws.OpenSearch.Inputs.PackagePackageSourceArgs
        {
            S3BucketName = aws_s3_bucket.My_opensearch_packages.Bucket,
            S3Key = aws_s3_object.Example.Key,
        },
        PackageType = "TXT-DICTIONARY",
    });

    var examplePackageAssociation = new Aws.OpenSearch.PackageAssociation("examplePackageAssociation", new()
    {
        PackageId = examplePackage.Id,
        DomainName = myDomain.DomainName,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/opensearch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myDomain, err := opensearch.NewDomain(ctx, "myDomain", &opensearch.DomainArgs{
			EngineVersion: pulumi.String("Elasticsearch_7.10"),
			ClusterConfig: &opensearch.DomainClusterConfigArgs{
				InstanceType: pulumi.String("r4.large.search"),
			},
		})
		if err != nil {
			return err
		}
		examplePackage, err := opensearch.NewPackage(ctx, "examplePackage", &opensearch.PackageArgs{
			PackageName: pulumi.String("example-txt"),
			PackageSource: &opensearch.PackagePackageSourceArgs{
				S3BucketName: pulumi.Any(aws_s3_bucket.My_opensearch_packages.Bucket),
				S3Key:        pulumi.Any(aws_s3_object.Example.Key),
			},
			PackageType: pulumi.String("TXT-DICTIONARY"),
		})
		if err != nil {
			return err
		}
		_, err = opensearch.NewPackageAssociation(ctx, "examplePackageAssociation", &opensearch.PackageAssociationArgs{
			PackageId:  examplePackage.ID(),
			DomainName: myDomain.DomainName,
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
import com.pulumi.aws.opensearch.Package;
import com.pulumi.aws.opensearch.PackageArgs;
import com.pulumi.aws.opensearch.inputs.PackagePackageSourceArgs;
import com.pulumi.aws.opensearch.PackageAssociation;
import com.pulumi.aws.opensearch.PackageAssociationArgs;
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
        var myDomain = new Domain("myDomain", DomainArgs.builder()        
            .engineVersion("Elasticsearch_7.10")
            .clusterConfig(DomainClusterConfigArgs.builder()
                .instanceType("r4.large.search")
                .build())
            .build());

        var examplePackage = new Package("examplePackage", PackageArgs.builder()        
            .packageName("example-txt")
            .packageSource(PackagePackageSourceArgs.builder()
                .s3BucketName(aws_s3_bucket.my_opensearch_packages().bucket())
                .s3Key(aws_s3_object.example().key())
                .build())
            .packageType("TXT-DICTIONARY")
            .build());

        var examplePackageAssociation = new PackageAssociation("examplePackageAssociation", PackageAssociationArgs.builder()        
            .packageId(examplePackage.id())
            .domainName(myDomain.domainName())
            .build());

    }
}
```
```yaml
resources:
  myDomain:
    type: aws:opensearch:Domain
    properties:
      engineVersion: Elasticsearch_7.10
      clusterConfig:
        instanceType: r4.large.search
  examplePackage:
    type: aws:opensearch:Package
    properties:
      packageName: example-txt
      packageSource:
        s3BucketName: ${aws_s3_bucket.my_opensearch_packages.bucket}
        s3Key: ${aws_s3_object.example.key}
      packageType: TXT-DICTIONARY
  examplePackageAssociation:
    type: aws:opensearch:PackageAssociation
    properties:
      packageId: ${examplePackage.id}
      domainName: ${myDomain.domainName}
```
{{% /example %}}
{{% /examples %}}