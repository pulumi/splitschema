Resource for managing an AWS CodeGuru Reviewer Repository Association.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleKey = new aws.kms.Key("exampleKey", {});
const exampleRepository = new aws.codecommit.Repository("exampleRepository", {repositoryName: "example-repo"});
const exampleRepositoryAssociation = new aws.codegurureviewer.RepositoryAssociation("exampleRepositoryAssociation", {
    repository: {
        codecommit: {
            name: exampleRepository.repositoryName,
        },
    },
    kmsKeyDetails: {
        encryptionOption: "CUSTOMER_MANAGED_CMK",
        kmsKeyId: exampleKey.keyId,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_key = aws.kms.Key("exampleKey")
example_repository = aws.codecommit.Repository("exampleRepository", repository_name="example-repo")
example_repository_association = aws.codegurureviewer.RepositoryAssociation("exampleRepositoryAssociation",
    repository=aws.codegurureviewer.RepositoryAssociationRepositoryArgs(
        codecommit=aws.codegurureviewer.RepositoryAssociationRepositoryCodecommitArgs(
            name=example_repository.repository_name,
        ),
    ),
    kms_key_details=aws.codegurureviewer.RepositoryAssociationKmsKeyDetailsArgs(
        encryption_option="CUSTOMER_MANAGED_CMK",
        kms_key_id=example_key.key_id,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleKey = new Aws.Kms.Key("exampleKey");

    var exampleRepository = new Aws.CodeCommit.Repository("exampleRepository", new()
    {
        RepositoryName = "example-repo",
    });

    var exampleRepositoryAssociation = new Aws.CodeGuruReviewer.RepositoryAssociation("exampleRepositoryAssociation", new()
    {
        Repository = new Aws.CodeGuruReviewer.Inputs.RepositoryAssociationRepositoryArgs
        {
            Codecommit = new Aws.CodeGuruReviewer.Inputs.RepositoryAssociationRepositoryCodecommitArgs
            {
                Name = exampleRepository.RepositoryName,
            },
        },
        KmsKeyDetails = new Aws.CodeGuruReviewer.Inputs.RepositoryAssociationKmsKeyDetailsArgs
        {
            EncryptionOption = "CUSTOMER_MANAGED_CMK",
            KmsKeyId = exampleKey.KeyId,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codecommit"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/codegurureviewer"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/kms"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleKey, err := kms.NewKey(ctx, "exampleKey", nil)
		if err != nil {
			return err
		}
		exampleRepository, err := codecommit.NewRepository(ctx, "exampleRepository", &codecommit.RepositoryArgs{
			RepositoryName: pulumi.String("example-repo"),
		})
		if err != nil {
			return err
		}
		_, err = codegurureviewer.NewRepositoryAssociation(ctx, "exampleRepositoryAssociation", &codegurureviewer.RepositoryAssociationArgs{
			Repository: &codegurureviewer.RepositoryAssociationRepositoryArgs{
				Codecommit: &codegurureviewer.RepositoryAssociationRepositoryCodecommitArgs{
					Name: exampleRepository.RepositoryName,
				},
			},
			KmsKeyDetails: &codegurureviewer.RepositoryAssociationKmsKeyDetailsArgs{
				EncryptionOption: pulumi.String("CUSTOMER_MANAGED_CMK"),
				KmsKeyId:         exampleKey.KeyId,
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
import com.pulumi.aws.codecommit.Repository;
import com.pulumi.aws.codecommit.RepositoryArgs;
import com.pulumi.aws.codegurureviewer.RepositoryAssociation;
import com.pulumi.aws.codegurureviewer.RepositoryAssociationArgs;
import com.pulumi.aws.codegurureviewer.inputs.RepositoryAssociationRepositoryArgs;
import com.pulumi.aws.codegurureviewer.inputs.RepositoryAssociationRepositoryCodecommitArgs;
import com.pulumi.aws.codegurureviewer.inputs.RepositoryAssociationKmsKeyDetailsArgs;
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

        var exampleRepository = new Repository("exampleRepository", RepositoryArgs.builder()        
            .repositoryName("example-repo")
            .build());

        var exampleRepositoryAssociation = new RepositoryAssociation("exampleRepositoryAssociation", RepositoryAssociationArgs.builder()        
            .repository(RepositoryAssociationRepositoryArgs.builder()
                .codecommit(RepositoryAssociationRepositoryCodecommitArgs.builder()
                    .name(exampleRepository.repositoryName())
                    .build())
                .build())
            .kmsKeyDetails(RepositoryAssociationKmsKeyDetailsArgs.builder()
                .encryptionOption("CUSTOMER_MANAGED_CMK")
                .kmsKeyId(exampleKey.keyId())
                .build())
            .build());

    }
}
```
```yaml
resources:
  exampleKey:
    type: aws:kms:Key
  exampleRepository:
    type: aws:codecommit:Repository
    properties:
      repositoryName: example-repo
  exampleRepositoryAssociation:
    type: aws:codegurureviewer:RepositoryAssociation
    properties:
      repository:
        codecommit:
          name: ${exampleRepository.repositoryName}
      kmsKeyDetails:
        encryptionOption: CUSTOMER_MANAGED_CMK
        kmsKeyId: ${exampleKey.keyId}
```
{{% /example %}}
{{% /examples %}}