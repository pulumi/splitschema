Provides a AWS Transfer Server resource.

> **NOTE on AWS IAM permissions:** If the `endpoint_type` is set to `VPC`, the `ec2:DescribeVpcEndpoints` and `ec2:ModifyVpcEndpoint` [actions](https://docs.aws.amazon.com/service-authorization/latest/reference/list_amazonec2.html#amazonec2-actions-as-permissions) are used.

> **NOTE:** Use the `aws.transfer.Tag` resource to manage the system tags used for [custom hostnames](https://docs.aws.amazon.com/transfer/latest/userguide/requirements-dns.html#tag-custom-hostname-cdk).

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Server("example", {tags: {
    Name: "Example",
}});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Server("example", tags={
    "Name": "Example",
})
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Server("example", new()
    {
        Tags = 
        {
            { "Name", "Example" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewServer(ctx, "example", &transfer.ServerArgs{
			Tags: pulumi.StringMap{
				"Name": pulumi.String("Example"),
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
import com.pulumi.aws.transfer.Server;
import com.pulumi.aws.transfer.ServerArgs;
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
        var example = new Server("example", ServerArgs.builder()        
            .tags(Map.of("Name", "Example"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Server
    properties:
      tags:
        Name: Example
```
{{% /example %}}
{{% example %}}
### Security Policy Name

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Server("example", {securityPolicyName: "TransferSecurityPolicy-2020-06"});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Server("example", security_policy_name="TransferSecurityPolicy-2020-06")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Server("example", new()
    {
        SecurityPolicyName = "TransferSecurityPolicy-2020-06",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewServer(ctx, "example", &transfer.ServerArgs{
			SecurityPolicyName: pulumi.String("TransferSecurityPolicy-2020-06"),
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
import com.pulumi.aws.transfer.Server;
import com.pulumi.aws.transfer.ServerArgs;
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
        var example = new Server("example", ServerArgs.builder()        
            .securityPolicyName("TransferSecurityPolicy-2020-06")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Server
    properties:
      securityPolicyName: TransferSecurityPolicy-2020-06
```
{{% /example %}}
{{% example %}}
### VPC Endpoint

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Server("example", {
    endpointType: "VPC",
    endpointDetails: {
        addressAllocationIds: [aws_eip.example.id],
        subnetIds: [aws_subnet.example.id],
        vpcId: aws_vpc.example.id,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Server("example",
    endpoint_type="VPC",
    endpoint_details=aws.transfer.ServerEndpointDetailsArgs(
        address_allocation_ids=[aws_eip["example"]["id"]],
        subnet_ids=[aws_subnet["example"]["id"]],
        vpc_id=aws_vpc["example"]["id"],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Server("example", new()
    {
        EndpointType = "VPC",
        EndpointDetails = new Aws.Transfer.Inputs.ServerEndpointDetailsArgs
        {
            AddressAllocationIds = new[]
            {
                aws_eip.Example.Id,
            },
            SubnetIds = new[]
            {
                aws_subnet.Example.Id,
            },
            VpcId = aws_vpc.Example.Id,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewServer(ctx, "example", &transfer.ServerArgs{
			EndpointType: pulumi.String("VPC"),
			EndpointDetails: &transfer.ServerEndpointDetailsArgs{
				AddressAllocationIds: pulumi.StringArray{
					aws_eip.Example.Id,
				},
				SubnetIds: pulumi.StringArray{
					aws_subnet.Example.Id,
				},
				VpcId: pulumi.Any(aws_vpc.Example.Id),
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
import com.pulumi.aws.transfer.Server;
import com.pulumi.aws.transfer.ServerArgs;
import com.pulumi.aws.transfer.inputs.ServerEndpointDetailsArgs;
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
        var example = new Server("example", ServerArgs.builder()        
            .endpointType("VPC")
            .endpointDetails(ServerEndpointDetailsArgs.builder()
                .addressAllocationIds(aws_eip.example().id())
                .subnetIds(aws_subnet.example().id())
                .vpcId(aws_vpc.example().id())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Server
    properties:
      endpointType: VPC
      endpointDetails:
        addressAllocationIds:
          - ${aws_eip.example.id}
        subnetIds:
          - ${aws_subnet.example.id}
        vpcId: ${aws_vpc.example.id}
```
{{% /example %}}
{{% example %}}
### AWS Directory authentication

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Server("example", {
    identityProviderType: "AWS_DIRECTORY_SERVICE",
    directoryId: aws_directory_service_directory.example.id,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Server("example",
    identity_provider_type="AWS_DIRECTORY_SERVICE",
    directory_id=aws_directory_service_directory["example"]["id"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Server("example", new()
    {
        IdentityProviderType = "AWS_DIRECTORY_SERVICE",
        DirectoryId = aws_directory_service_directory.Example.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewServer(ctx, "example", &transfer.ServerArgs{
			IdentityProviderType: pulumi.String("AWS_DIRECTORY_SERVICE"),
			DirectoryId:          pulumi.Any(aws_directory_service_directory.Example.Id),
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
import com.pulumi.aws.transfer.Server;
import com.pulumi.aws.transfer.ServerArgs;
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
        var example = new Server("example", ServerArgs.builder()        
            .identityProviderType("AWS_DIRECTORY_SERVICE")
            .directoryId(aws_directory_service_directory.example().id())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Server
    properties:
      identityProviderType: AWS_DIRECTORY_SERVICE
      directoryId: ${aws_directory_service_directory.example.id}
```
{{% /example %}}
{{% example %}}
### AWS Lambda authentication

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Server("example", {
    identityProviderType: "AWS_LAMBDA",
    "function": aws_lambda_identity_provider.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Server("example",
    identity_provider_type="AWS_LAMBDA",
    function=aws_lambda_identity_provider["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Server("example", new()
    {
        IdentityProviderType = "AWS_LAMBDA",
        Function = aws_lambda_identity_provider.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewServer(ctx, "example", &transfer.ServerArgs{
			IdentityProviderType: pulumi.String("AWS_LAMBDA"),
			Function:             pulumi.Any(aws_lambda_identity_provider.Example.Arn),
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
import com.pulumi.aws.transfer.Server;
import com.pulumi.aws.transfer.ServerArgs;
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
        var example = new Server("example", ServerArgs.builder()        
            .identityProviderType("AWS_LAMBDA")
            .function(aws_lambda_identity_provider.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Server
    properties:
      identityProviderType: AWS_LAMBDA
      function: ${aws_lambda_identity_provider.example.arn}
```
{{% /example %}}
{{% example %}}
### Protocols

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.transfer.Server("example", {
    endpointType: "VPC",
    endpointDetails: {
        subnetIds: [aws_subnet.example.id],
        vpcId: aws_vpc.example.id,
    },
    protocols: [
        "FTP",
        "FTPS",
    ],
    certificate: aws_acm_certificate.example.arn,
    identityProviderType: "API_GATEWAY",
    url: `${aws_api_gateway_deployment.example.invoke_url}${aws_api_gateway_resource.example.path}`,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.transfer.Server("example",
    endpoint_type="VPC",
    endpoint_details=aws.transfer.ServerEndpointDetailsArgs(
        subnet_ids=[aws_subnet["example"]["id"]],
        vpc_id=aws_vpc["example"]["id"],
    ),
    protocols=[
        "FTP",
        "FTPS",
    ],
    certificate=aws_acm_certificate["example"]["arn"],
    identity_provider_type="API_GATEWAY",
    url=f"{aws_api_gateway_deployment['example']['invoke_url']}{aws_api_gateway_resource['example']['path']}")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Transfer.Server("example", new()
    {
        EndpointType = "VPC",
        EndpointDetails = new Aws.Transfer.Inputs.ServerEndpointDetailsArgs
        {
            SubnetIds = new[]
            {
                aws_subnet.Example.Id,
            },
            VpcId = aws_vpc.Example.Id,
        },
        Protocols = new[]
        {
            "FTP",
            "FTPS",
        },
        Certificate = aws_acm_certificate.Example.Arn,
        IdentityProviderType = "API_GATEWAY",
        Url = $"{aws_api_gateway_deployment.Example.Invoke_url}{aws_api_gateway_resource.Example.Path}",
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := transfer.NewServer(ctx, "example", &transfer.ServerArgs{
			EndpointType: pulumi.String("VPC"),
			EndpointDetails: &transfer.ServerEndpointDetailsArgs{
				SubnetIds: pulumi.StringArray{
					aws_subnet.Example.Id,
				},
				VpcId: pulumi.Any(aws_vpc.Example.Id),
			},
			Protocols: pulumi.StringArray{
				pulumi.String("FTP"),
				pulumi.String("FTPS"),
			},
			Certificate:          pulumi.Any(aws_acm_certificate.Example.Arn),
			IdentityProviderType: pulumi.String("API_GATEWAY"),
			Url:                  pulumi.String(fmt.Sprintf("%v%v", aws_api_gateway_deployment.Example.Invoke_url, aws_api_gateway_resource.Example.Path)),
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
import com.pulumi.aws.transfer.Server;
import com.pulumi.aws.transfer.ServerArgs;
import com.pulumi.aws.transfer.inputs.ServerEndpointDetailsArgs;
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
        var example = new Server("example", ServerArgs.builder()        
            .endpointType("VPC")
            .endpointDetails(ServerEndpointDetailsArgs.builder()
                .subnetIds(aws_subnet.example().id())
                .vpcId(aws_vpc.example().id())
                .build())
            .protocols(            
                "FTP",
                "FTPS")
            .certificate(aws_acm_certificate.example().arn())
            .identityProviderType("API_GATEWAY")
            .url(String.format("%s%s", aws_api_gateway_deployment.example().invoke_url(),aws_api_gateway_resource.example().path()))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Server
    properties:
      endpointType: VPC
      endpointDetails:
        subnetIds:
          - ${aws_subnet.example.id}
        vpcId: ${aws_vpc.example.id}
      protocols:
        - FTP
        - FTPS
      certificate: ${aws_acm_certificate.example.arn}
      identityProviderType: API_GATEWAY
      url: ${aws_api_gateway_deployment.example.invoke_url}${aws_api_gateway_resource.example.path}
```
{{% /example %}}
{{% example %}}
### Using Structured Logging Destinations

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const transferLogGroup = new aws.cloudwatch.LogGroup("transferLogGroup", {namePrefix: "transfer_test_"});
const transferAssumeRole = aws.iam.getPolicyDocument({
    statements: [{
        effect: "Allow",
        principals: [{
            type: "Service",
            identifiers: ["transfer.amazonaws.com"],
        }],
        actions: ["sts:AssumeRole"],
    }],
});
const iamForTransfer = new aws.iam.Role("iamForTransfer", {
    namePrefix: "iam_for_transfer_",
    assumeRolePolicy: transferAssumeRole.then(transferAssumeRole => transferAssumeRole.json),
    managedPolicyArns: ["arn:aws:iam::aws:policy/service-role/AWSTransferLoggingAccess"],
});
const transferServer = new aws.transfer.Server("transferServer", {
    endpointType: "PUBLIC",
    loggingRole: iamForTransfer.arn,
    protocols: ["SFTP"],
    structuredLogDestinations: [pulumi.interpolate`${transferLogGroup.arn}:*`],
});
```
```python
import pulumi
import pulumi_aws as aws

transfer_log_group = aws.cloudwatch.LogGroup("transferLogGroup", name_prefix="transfer_test_")
transfer_assume_role = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    effect="Allow",
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["transfer.amazonaws.com"],
    )],
    actions=["sts:AssumeRole"],
)])
iam_for_transfer = aws.iam.Role("iamForTransfer",
    name_prefix="iam_for_transfer_",
    assume_role_policy=transfer_assume_role.json,
    managed_policy_arns=["arn:aws:iam::aws:policy/service-role/AWSTransferLoggingAccess"])
transfer_server = aws.transfer.Server("transferServer",
    endpoint_type="PUBLIC",
    logging_role=iam_for_transfer.arn,
    protocols=["SFTP"],
    structured_log_destinations=[transfer_log_group.arn.apply(lambda arn: f"{arn}:*")])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var transferLogGroup = new Aws.CloudWatch.LogGroup("transferLogGroup", new()
    {
        NamePrefix = "transfer_test_",
    });

    var transferAssumeRole = Aws.Iam.GetPolicyDocument.Invoke(new()
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
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "transfer.amazonaws.com",
                        },
                    },
                },
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
            },
        },
    });

    var iamForTransfer = new Aws.Iam.Role("iamForTransfer", new()
    {
        NamePrefix = "iam_for_transfer_",
        AssumeRolePolicy = transferAssumeRole.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
        ManagedPolicyArns = new[]
        {
            "arn:aws:iam::aws:policy/service-role/AWSTransferLoggingAccess",
        },
    });

    var transferServer = new Aws.Transfer.Server("transferServer", new()
    {
        EndpointType = "PUBLIC",
        LoggingRole = iamForTransfer.Arn,
        Protocols = new[]
        {
            "SFTP",
        },
        StructuredLogDestinations = new[]
        {
            transferLogGroup.Arn.Apply(arn => $"{arn}:*"),
        },
    });

});
```
```go
package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/transfer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		transferLogGroup, err := cloudwatch.NewLogGroup(ctx, "transferLogGroup", &cloudwatch.LogGroupArgs{
			NamePrefix: pulumi.String("transfer_test_"),
		})
		if err != nil {
			return err
		}
		transferAssumeRole, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Effect: pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"transfer.amazonaws.com",
							},
						},
					},
					Actions: []string{
						"sts:AssumeRole",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		iamForTransfer, err := iam.NewRole(ctx, "iamForTransfer", &iam.RoleArgs{
			NamePrefix:       pulumi.String("iam_for_transfer_"),
			AssumeRolePolicy: *pulumi.String(transferAssumeRole.Json),
			ManagedPolicyArns: pulumi.StringArray{
				pulumi.String("arn:aws:iam::aws:policy/service-role/AWSTransferLoggingAccess"),
			},
		})
		if err != nil {
			return err
		}
		_, err = transfer.NewServer(ctx, "transferServer", &transfer.ServerArgs{
			EndpointType: pulumi.String("PUBLIC"),
			LoggingRole:  iamForTransfer.Arn,
			Protocols: pulumi.StringArray{
				pulumi.String("SFTP"),
			},
			StructuredLogDestinations: pulumi.StringArray{
				transferLogGroup.Arn.ApplyT(func(arn string) (string, error) {
					return fmt.Sprintf("%v:*", arn), nil
				}).(pulumi.StringOutput),
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
import com.pulumi.aws.cloudwatch.LogGroup;
import com.pulumi.aws.cloudwatch.LogGroupArgs;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.transfer.Server;
import com.pulumi.aws.transfer.ServerArgs;
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
        var transferLogGroup = new LogGroup("transferLogGroup", LogGroupArgs.builder()        
            .namePrefix("transfer_test_")
            .build());

        final var transferAssumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .effect("Allow")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("transfer.amazonaws.com")
                    .build())
                .actions("sts:AssumeRole")
                .build())
            .build());

        var iamForTransfer = new Role("iamForTransfer", RoleArgs.builder()        
            .namePrefix("iam_for_transfer_")
            .assumeRolePolicy(transferAssumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .managedPolicyArns("arn:aws:iam::aws:policy/service-role/AWSTransferLoggingAccess")
            .build());

        var transferServer = new Server("transferServer", ServerArgs.builder()        
            .endpointType("PUBLIC")
            .loggingRole(iamForTransfer.arn())
            .protocols("SFTP")
            .structuredLogDestinations(transferLogGroup.arn().applyValue(arn -> String.format("%s:*", arn)))
            .build());

    }
}
```
```yaml
resources:
  transferLogGroup:
    type: aws:cloudwatch:LogGroup
    properties:
      namePrefix: transfer_test_
  iamForTransfer:
    type: aws:iam:Role
    properties:
      namePrefix: iam_for_transfer_
      assumeRolePolicy: ${transferAssumeRole.json}
      managedPolicyArns:
        - arn:aws:iam::aws:policy/service-role/AWSTransferLoggingAccess
  transferServer:
    type: aws:transfer:Server
    properties:
      endpointType: PUBLIC
      loggingRole: ${iamForTransfer.arn}
      protocols:
        - SFTP
      structuredLogDestinations:
        - ${transferLogGroup.arn}:*
variables:
  transferAssumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - effect: Allow
            principals:
              - type: Service
                identifiers:
                  - transfer.amazonaws.com
            actions:
              - sts:AssumeRole
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transfer Servers using the server `id`. For example:

```sh
 $ pulumi import aws:transfer/server:Server example s-12345678
```
 Certain resource arguments, such as `host_key`, cannot be read via the API and imported into the provider. This provider will display a difference for these arguments the first run after import if declared in the provider configuration for an imported resource.

