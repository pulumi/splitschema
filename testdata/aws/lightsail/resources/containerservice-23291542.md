An Amazon Lightsail container service is a highly scalable compute and networking resource on which you can deploy, run,
and manage containers. For more information, see
[Container services in Amazon Lightsail](https://lightsail.aws.amazon.com/ls/docs/en_us/articles/amazon-lightsail-container-services).

> **Note:** For more information about the AWS Regions in which you can create Amazon Lightsail container services,
see ["Regions and Availability Zones in Amazon Lightsail"](https://lightsail.aws.amazon.com/ls/docs/overview/article/understanding-regions-and-availability-zones-in-amazon-lightsail).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myContainerService = new aws.lightsail.ContainerService("myContainerService", {
    isDisabled: false,
    power: "nano",
    scale: 1,
    tags: {
        foo1: "bar1",
        foo2: "",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

my_container_service = aws.lightsail.ContainerService("myContainerService",
    is_disabled=False,
    power="nano",
    scale=1,
    tags={
        "foo1": "bar1",
        "foo2": "",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myContainerService = new Aws.LightSail.ContainerService("myContainerService", new()
    {
        IsDisabled = false,
        Power = "nano",
        Scale = 1,
        Tags = 
        {
            { "foo1", "bar1" },
            { "foo2", "" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewContainerService(ctx, "myContainerService", &lightsail.ContainerServiceArgs{
			IsDisabled: pulumi.Bool(false),
			Power:      pulumi.String("nano"),
			Scale:      pulumi.Int(1),
			Tags: pulumi.StringMap{
				"foo1": pulumi.String("bar1"),
				"foo2": pulumi.String(""),
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
import com.pulumi.aws.lightsail.ContainerService;
import com.pulumi.aws.lightsail.ContainerServiceArgs;
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
        var myContainerService = new ContainerService("myContainerService", ContainerServiceArgs.builder()        
            .isDisabled(false)
            .power("nano")
            .scale(1)
            .tags(Map.ofEntries(
                Map.entry("foo1", "bar1"),
                Map.entry("foo2", "")
            ))
            .build());

    }
}
```
```yaml
resources:
  myContainerService:
    type: aws:lightsail:ContainerService
    properties:
      isDisabled: false
      power: nano
      scale: 1
      tags:
        foo1: bar1
        foo2:
```
{{% /example %}}
{{% example %}}
### Public Domain Names

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myContainerService = new aws.lightsail.ContainerService("myContainerService", {publicDomainNames: {
    certificates: [{
        certificateName: "example-certificate",
        domainNames: ["www.example.com"],
    }],
}});
```
```python
import pulumi
import pulumi_aws as aws

my_container_service = aws.lightsail.ContainerService("myContainerService", public_domain_names=aws.lightsail.ContainerServicePublicDomainNamesArgs(
    certificates=[aws.lightsail.ContainerServicePublicDomainNamesCertificateArgs(
        certificate_name="example-certificate",
        domain_names=["www.example.com"],
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
    var myContainerService = new Aws.LightSail.ContainerService("myContainerService", new()
    {
        PublicDomainNames = new Aws.LightSail.Inputs.ContainerServicePublicDomainNamesArgs
        {
            Certificates = new[]
            {
                new Aws.LightSail.Inputs.ContainerServicePublicDomainNamesCertificateArgs
                {
                    CertificateName = "example-certificate",
                    DomainNames = new[]
                    {
                        "www.example.com",
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewContainerService(ctx, "myContainerService", &lightsail.ContainerServiceArgs{
			PublicDomainNames: &lightsail.ContainerServicePublicDomainNamesArgs{
				Certificates: lightsail.ContainerServicePublicDomainNamesCertificateArray{
					&lightsail.ContainerServicePublicDomainNamesCertificateArgs{
						CertificateName: pulumi.String("example-certificate"),
						DomainNames: pulumi.StringArray{
							pulumi.String("www.example.com"),
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
import com.pulumi.aws.lightsail.ContainerService;
import com.pulumi.aws.lightsail.ContainerServiceArgs;
import com.pulumi.aws.lightsail.inputs.ContainerServicePublicDomainNamesArgs;
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
        var myContainerService = new ContainerService("myContainerService", ContainerServiceArgs.builder()        
            .publicDomainNames(ContainerServicePublicDomainNamesArgs.builder()
                .certificates(ContainerServicePublicDomainNamesCertificateArgs.builder()
                    .certificateName("example-certificate")
                    .domainNames("www.example.com")
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  myContainerService:
    type: aws:lightsail:ContainerService
    properties:
      publicDomainNames:
        certificates:
          - certificateName: example-certificate
            domainNames:
              - www.example.com
```
{{% /example %}}
{{% example %}}
### Private Registry Access

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ... other configuration ...
const defaultContainerService = new aws.lightsail.ContainerService("defaultContainerService", {privateRegistryAccess: {
    ecrImagePullerRole: {
        isActive: true,
    },
}});
const defaultPolicyDocument = defaultContainerService.privateRegistryAccess.apply(privateRegistryAccess => aws.iam.getPolicyDocumentOutput({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "AWS",
            identifiers: [privateRegistryAccess.ecrImagePullerRole?.principalArn],
        }],
        actions: [
            "ecr:BatchGetImage",
            "ecr:GetDownloadUrlForLayer",
        ],
    }],
}));
const defaultRepositoryPolicy = new aws.ecr.RepositoryPolicy("defaultRepositoryPolicy", {
    repository: aws_ecr_repository["default"].name,
    policy: defaultPolicyDocument.apply(defaultPolicyDocument => defaultPolicyDocument.json),
});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configuration ...
default_container_service = aws.lightsail.ContainerService("defaultContainerService", private_registry_access=aws.lightsail.ContainerServicePrivateRegistryAccessArgs(
    ecr_image_puller_role=aws.lightsail.ContainerServicePrivateRegistryAccessEcrImagePullerRoleArgs(
        is_active=True,
    ),
))
default_policy_document = default_container_service.private_registry_access.apply(lambda private_registry_access: aws.iam.get_policy_document_output(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="AWS",
        identifiers=[private_registry_access.ecr_image_puller_role.principal_arn],
    )],
    actions=[
        "ecr:BatchGetImage",
        "ecr:GetDownloadUrlForLayer",
    ],
)]))
default_repository_policy = aws.ecr.RepositoryPolicy("defaultRepositoryPolicy",
    repository=aws_ecr_repository["default"]["name"],
    policy=default_policy_document.json)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ... other configuration ...
    var defaultContainerService = new Aws.LightSail.ContainerService("defaultContainerService", new()
    {
        PrivateRegistryAccess = new Aws.LightSail.Inputs.ContainerServicePrivateRegistryAccessArgs
        {
            EcrImagePullerRole = new Aws.LightSail.Inputs.ContainerServicePrivateRegistryAccessEcrImagePullerRoleArgs
            {
                IsActive = true,
            },
        },
    });

    var defaultPolicyDocument = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Effect = "Allow",
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "AWS",
                        Identifiers = new[]
                        {
                            defaultContainerService.PrivateRegistryAccess.EcrImagePullerRole?.PrincipalArn,
                        },
                    },
                },
                Actions = new[]
                {
                    "ecr:BatchGetImage",
                    "ecr:GetDownloadUrlForLayer",
                },
            },
        },
    });

    var defaultRepositoryPolicy = new Aws.Ecr.RepositoryPolicy("defaultRepositoryPolicy", new()
    {
        Repository = aws_ecr_repository.Default.Name,
        Policy = defaultPolicyDocument.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecr"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
defaultContainerService, err := lightsail.NewContainerService(ctx, "defaultContainerService", &lightsail.ContainerServiceArgs{
PrivateRegistryAccess: &lightsail.ContainerServicePrivateRegistryAccessArgs{
EcrImagePullerRole: &lightsail.ContainerServicePrivateRegistryAccessEcrImagePullerRoleArgs{
IsActive: pulumi.Bool(true),
},
},
})
if err != nil {
return err
}
defaultPolicyDocument := defaultContainerService.PrivateRegistryAccess.ApplyT(func(privateRegistryAccess lightsail.ContainerServicePrivateRegistryAccess) (iam.GetPolicyDocumentResult, error) {
return iam.GetPolicyDocumentOutput(ctx, iam.GetPolicyDocumentOutputArgs{
Statements: []iam.GetPolicyDocumentStatement{
{
Effect: "Allow",
Principals: []iam.GetPolicyDocumentStatementPrincipal{
{
Type: "AWS",
Identifiers: interface{}{
privateRegistryAccess.EcrImagePullerRole.PrincipalArn,
},
},
},
Actions: []string{
"ecr:BatchGetImage",
"ecr:GetDownloadUrlForLayer",
},
},
},
}, nil), nil
}).(iam.GetPolicyDocumentResultOutput)
_, err = ecr.NewRepositoryPolicy(ctx, "defaultRepositoryPolicy", &ecr.RepositoryPolicyArgs{
Repository: pulumi.Any(aws_ecr_repository.Default.Name),
Policy: defaultPolicyDocument.ApplyT(func(defaultPolicyDocument iam.GetPolicyDocumentResult) (*string, error) {
return &defaultPolicyDocument.Json, nil
}).(pulumi.StringPtrOutput),
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
import com.pulumi.aws.lightsail.ContainerService;
import com.pulumi.aws.lightsail.ContainerServiceArgs;
import com.pulumi.aws.lightsail.inputs.ContainerServicePrivateRegistryAccessArgs;
import com.pulumi.aws.lightsail.inputs.ContainerServicePrivateRegistryAccessEcrImagePullerRoleArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.ecr.RepositoryPolicy;
import com.pulumi.aws.ecr.RepositoryPolicyArgs;
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
        var defaultContainerService = new ContainerService("defaultContainerService", ContainerServiceArgs.builder()        
            .privateRegistryAccess(ContainerServicePrivateRegistryAccessArgs.builder()
                .ecrImagePullerRole(ContainerServicePrivateRegistryAccessEcrImagePullerRoleArgs.builder()
                    .isActive(true)
                    .build())
                .build())
            .build());

        final var defaultPolicyDocument = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("AWS")
                    .identifiers(defaultContainerService.privateRegistryAccess().applyValue(privateRegistryAccess -> privateRegistryAccess.ecrImagePullerRole().principalArn()))
                    .build())
                .actions(                
                    "ecr:BatchGetImage",
                    "ecr:GetDownloadUrlForLayer")
                .build())
            .build());

        var defaultRepositoryPolicy = new RepositoryPolicy("defaultRepositoryPolicy", RepositoryPolicyArgs.builder()        
            .repository(aws_ecr_repository.default().name())
            .policy(defaultPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult).applyValue(defaultPolicyDocument -> defaultPolicyDocument.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json())))
            .build());

    }
}
```
```yaml
resources:
  defaultContainerService:
    type: aws:lightsail:ContainerService
    properties:
      privateRegistryAccess:
        ecrImagePullerRole:
          isActive: true
  defaultRepositoryPolicy:
    type: aws:ecr:RepositoryPolicy
    properties:
      repository: ${aws_ecr_repository.default.name}
      policy: ${defaultPolicyDocument.json}
variables:
  defaultPolicyDocument:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: AWS
                identifiers:
                  - ${defaultContainerService.privateRegistryAccess.ecrImagePullerRole.principalArn}
            actions:
              - ecr:BatchGetImage
              - ecr:GetDownloadUrlForLayer
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lightsail Container Service using the `name`. For example:

```sh
 $ pulumi import aws:lightsail/containerService:ContainerService my_container_service container-service-1
```
 