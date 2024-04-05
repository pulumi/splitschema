Associate an Amazon FSx file system with the FSx File Gateway. After the association process is complete, the file shares on the Amazon FSx file system are available for access through the gateway. This operation only supports the FSx File Gateway type.

[FSx File Gateway requirements](https://docs.aws.amazon.com/filegateway/latest/filefsxw/Requirements.html).

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.FileSystemAssociation("example", {
    gatewayArn: aws_storagegateway_gateway.example.arn,
    locationArn: aws_fsx_windows_file_system.example.arn,
    username: "Admin",
    password: "avoid-plaintext-passwords",
    auditDestinationArn: aws_s3_bucket.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.FileSystemAssociation("example",
    gateway_arn=aws_storagegateway_gateway["example"]["arn"],
    location_arn=aws_fsx_windows_file_system["example"]["arn"],
    username="Admin",
    password="avoid-plaintext-passwords",
    audit_destination_arn=aws_s3_bucket["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.FileSystemAssociation("example", new()
    {
        GatewayArn = aws_storagegateway_gateway.Example.Arn,
        LocationArn = aws_fsx_windows_file_system.Example.Arn,
        Username = "Admin",
        Password = "avoid-plaintext-passwords",
        AuditDestinationArn = aws_s3_bucket.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagegateway.NewFileSystemAssociation(ctx, "example", &storagegateway.FileSystemAssociationArgs{
			GatewayArn:          pulumi.Any(aws_storagegateway_gateway.Example.Arn),
			LocationArn:         pulumi.Any(aws_fsx_windows_file_system.Example.Arn),
			Username:            pulumi.String("Admin"),
			Password:            pulumi.String("avoid-plaintext-passwords"),
			AuditDestinationArn: pulumi.Any(aws_s3_bucket.Example.Arn),
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
import com.pulumi.aws.storagegateway.FileSystemAssociation;
import com.pulumi.aws.storagegateway.FileSystemAssociationArgs;
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
        var example = new FileSystemAssociation("example", FileSystemAssociationArgs.builder()        
            .gatewayArn(aws_storagegateway_gateway.example().arn())
            .locationArn(aws_fsx_windows_file_system.example().arn())
            .username("Admin")
            .password("avoid-plaintext-passwords")
            .auditDestinationArn(aws_s3_bucket.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:FileSystemAssociation
    properties:
      gatewayArn: ${aws_storagegateway_gateway.example.arn}
      locationArn: ${aws_fsx_windows_file_system.example.arn}
      username: Admin
      password: avoid-plaintext-passwords
      auditDestinationArn: ${aws_s3_bucket.example.arn}
```
{{% /example %}}
{{% /examples %}}
## Required Services Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const awsServiceStoragegatewayAmiFILES3Latest = aws.ssm.getParameter({
    name: "/aws/service/storagegateway/ami/FILE_S3/latest",
});
const testInstance = new aws.ec2.Instance("testInstance", {
    ami: awsServiceStoragegatewayAmiFILES3Latest.then(awsServiceStoragegatewayAmiFILES3Latest => awsServiceStoragegatewayAmiFILES3Latest.value),
    associatePublicIpAddress: true,
    instanceType: aws.ec2.instancetype.InstanceType[data.aws_ec2_instance_type_offering.available.instance_type],
    vpcSecurityGroupIds: [aws_security_group.test.id],
    subnetId: aws_subnet.test[0].id,
}, {
    dependsOn: [
        aws_route.test,
        aws_vpc_dhcp_options_association.test,
    ],
});
const testGateway = new aws.storagegateway.Gateway("testGateway", {
    gatewayIpAddress: testInstance.publicIp,
    gatewayName: "test-sgw",
    gatewayTimezone: "GMT",
    gatewayType: "FILE_FSX_SMB",
    smbActiveDirectorySettings: {
        domainName: aws_directory_service_directory.test.name,
        password: aws_directory_service_directory.test.password,
        username: "Admin",
    },
});
const testWindowsFileSystem = new aws.fsx.WindowsFileSystem("testWindowsFileSystem", {
    activeDirectoryId: aws_directory_service_directory.test.id,
    securityGroupIds: [aws_security_group.test.id],
    skipFinalBackup: true,
    storageCapacity: 32,
    subnetIds: [aws_subnet.test[0].id],
    throughputCapacity: 8,
});
const fsx = new aws.storagegateway.FileSystemAssociation("fsx", {
    gatewayArn: testGateway.arn,
    locationArn: testWindowsFileSystem.arn,
    username: "Admin",
    password: aws_directory_service_directory.test.password,
    cacheAttributes: {
        cacheStaleTimeoutInSeconds: 400,
    },
    auditDestinationArn: aws_cloudwatch_log_group.test.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

aws_service_storagegateway_ami_files3_latest = aws.ssm.get_parameter(name="/aws/service/storagegateway/ami/FILE_S3/latest")
test_instance = aws.ec2.Instance("testInstance",
    ami=aws_service_storagegateway_ami_files3_latest.value,
    associate_public_ip_address=True,
    instance_type=aws.ec2/instancetype.InstanceType(data["aws_ec2_instance_type_offering"]["available"]["instance_type"]),
    vpc_security_group_ids=[aws_security_group["test"]["id"]],
    subnet_id=aws_subnet["test"][0]["id"],
    opts=pulumi.ResourceOptions(depends_on=[
            aws_route["test"],
            aws_vpc_dhcp_options_association["test"],
        ]))
test_gateway = aws.storagegateway.Gateway("testGateway",
    gateway_ip_address=test_instance.public_ip,
    gateway_name="test-sgw",
    gateway_timezone="GMT",
    gateway_type="FILE_FSX_SMB",
    smb_active_directory_settings=aws.storagegateway.GatewaySmbActiveDirectorySettingsArgs(
        domain_name=aws_directory_service_directory["test"]["name"],
        password=aws_directory_service_directory["test"]["password"],
        username="Admin",
    ))
test_windows_file_system = aws.fsx.WindowsFileSystem("testWindowsFileSystem",
    active_directory_id=aws_directory_service_directory["test"]["id"],
    security_group_ids=[aws_security_group["test"]["id"]],
    skip_final_backup=True,
    storage_capacity=32,
    subnet_ids=[aws_subnet["test"][0]["id"]],
    throughput_capacity=8)
fsx = aws.storagegateway.FileSystemAssociation("fsx",
    gateway_arn=test_gateway.arn,
    location_arn=test_windows_file_system.arn,
    username="Admin",
    password=aws_directory_service_directory["test"]["password"],
    cache_attributes=aws.storagegateway.FileSystemAssociationCacheAttributesArgs(
        cache_stale_timeout_in_seconds=400,
    ),
    audit_destination_arn=aws_cloudwatch_log_group["test"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var awsServiceStoragegatewayAmiFILES3Latest = Aws.Ssm.GetParameter.Invoke(new()
    {
        Name = "/aws/service/storagegateway/ami/FILE_S3/latest",
    });

    var testInstance = new Aws.Ec2.Instance("testInstance", new()
    {
        Ami = awsServiceStoragegatewayAmiFILES3Latest.Apply(getParameterResult => getParameterResult.Value),
        AssociatePublicIpAddress = true,
        InstanceType = System.Enum.Parse<Aws.Ec2.InstanceType.InstanceType>(data.Aws_ec2_instance_type_offering.Available.Instance_type),
        VpcSecurityGroupIds = new[]
        {
            aws_security_group.Test.Id,
        },
        SubnetId = aws_subnet.Test[0].Id,
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_route.Test,
            aws_vpc_dhcp_options_association.Test,
        },
    });

    var testGateway = new Aws.StorageGateway.Gateway("testGateway", new()
    {
        GatewayIpAddress = testInstance.PublicIp,
        GatewayName = "test-sgw",
        GatewayTimezone = "GMT",
        GatewayType = "FILE_FSX_SMB",
        SmbActiveDirectorySettings = new Aws.StorageGateway.Inputs.GatewaySmbActiveDirectorySettingsArgs
        {
            DomainName = aws_directory_service_directory.Test.Name,
            Password = aws_directory_service_directory.Test.Password,
            Username = "Admin",
        },
    });

    var testWindowsFileSystem = new Aws.Fsx.WindowsFileSystem("testWindowsFileSystem", new()
    {
        ActiveDirectoryId = aws_directory_service_directory.Test.Id,
        SecurityGroupIds = new[]
        {
            aws_security_group.Test.Id,
        },
        SkipFinalBackup = true,
        StorageCapacity = 32,
        SubnetIds = new[]
        {
            aws_subnet.Test[0].Id,
        },
        ThroughputCapacity = 8,
    });

    var fsx = new Aws.StorageGateway.FileSystemAssociation("fsx", new()
    {
        GatewayArn = testGateway.Arn,
        LocationArn = testWindowsFileSystem.Arn,
        Username = "Admin",
        Password = aws_directory_service_directory.Test.Password,
        CacheAttributes = new Aws.StorageGateway.Inputs.FileSystemAssociationCacheAttributesArgs
        {
            CacheStaleTimeoutInSeconds = 400,
        },
        AuditDestinationArn = aws_cloudwatch_log_group.Test.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/fsx"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ssm"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		awsServiceStoragegatewayAmiFILES3Latest, err := ssm.LookupParameter(ctx, &ssm.LookupParameterArgs{
			Name: "/aws/service/storagegateway/ami/FILE_S3/latest",
		}, nil)
		if err != nil {
			return err
		}
		testInstance, err := ec2.NewInstance(ctx, "testInstance", &ec2.InstanceArgs{
			Ami:                      *pulumi.String(awsServiceStoragegatewayAmiFILES3Latest.Value),
			AssociatePublicIpAddress: pulumi.Bool(true),
			InstanceType:             ec2.InstanceType(data.Aws_ec2_instance_type_offering.Available.Instance_type),
			VpcSecurityGroupIds: pulumi.StringArray{
				aws_security_group.Test.Id,
			},
			SubnetId: pulumi.Any(aws_subnet.Test[0].Id),
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_route.Test,
			aws_vpc_dhcp_options_association.Test,
		}))
		if err != nil {
			return err
		}
		testGateway, err := storagegateway.NewGateway(ctx, "testGateway", &storagegateway.GatewayArgs{
			GatewayIpAddress: testInstance.PublicIp,
			GatewayName:      pulumi.String("test-sgw"),
			GatewayTimezone:  pulumi.String("GMT"),
			GatewayType:      pulumi.String("FILE_FSX_SMB"),
			SmbActiveDirectorySettings: &storagegateway.GatewaySmbActiveDirectorySettingsArgs{
				DomainName: pulumi.Any(aws_directory_service_directory.Test.Name),
				Password:   pulumi.Any(aws_directory_service_directory.Test.Password),
				Username:   pulumi.String("Admin"),
			},
		})
		if err != nil {
			return err
		}
		testWindowsFileSystem, err := fsx.NewWindowsFileSystem(ctx, "testWindowsFileSystem", &fsx.WindowsFileSystemArgs{
			ActiveDirectoryId: pulumi.Any(aws_directory_service_directory.Test.Id),
			SecurityGroupIds: pulumi.StringArray{
				aws_security_group.Test.Id,
			},
			SkipFinalBackup: pulumi.Bool(true),
			StorageCapacity: pulumi.Int(32),
			SubnetIds: pulumi.StringArray{
				aws_subnet.Test[0].Id,
			},
			ThroughputCapacity: pulumi.Int(8),
		})
		if err != nil {
			return err
		}
		_, err = storagegateway.NewFileSystemAssociation(ctx, "fsx", &storagegateway.FileSystemAssociationArgs{
			GatewayArn:  testGateway.Arn,
			LocationArn: testWindowsFileSystem.Arn,
			Username:    pulumi.String("Admin"),
			Password:    pulumi.Any(aws_directory_service_directory.Test.Password),
			CacheAttributes: &storagegateway.FileSystemAssociationCacheAttributesArgs{
				CacheStaleTimeoutInSeconds: pulumi.Int(400),
			},
			AuditDestinationArn: pulumi.Any(aws_cloudwatch_log_group.Test.Arn),
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
import com.pulumi.aws.ssm.SsmFunctions;
import com.pulumi.aws.ssm.inputs.GetParameterArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import com.pulumi.aws.storagegateway.Gateway;
import com.pulumi.aws.storagegateway.GatewayArgs;
import com.pulumi.aws.storagegateway.inputs.GatewaySmbActiveDirectorySettingsArgs;
import com.pulumi.aws.fsx.WindowsFileSystem;
import com.pulumi.aws.fsx.WindowsFileSystemArgs;
import com.pulumi.aws.storagegateway.FileSystemAssociation;
import com.pulumi.aws.storagegateway.FileSystemAssociationArgs;
import com.pulumi.aws.storagegateway.inputs.FileSystemAssociationCacheAttributesArgs;
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
        final var awsServiceStoragegatewayAmiFILES3Latest = SsmFunctions.getParameter(GetParameterArgs.builder()
            .name("/aws/service/storagegateway/ami/FILE_S3/latest")
            .build());

        var testInstance = new Instance("testInstance", InstanceArgs.builder()        
            .ami(awsServiceStoragegatewayAmiFILES3Latest.applyValue(getParameterResult -> getParameterResult.value()))
            .associatePublicIpAddress(true)
            .instanceType(data.aws_ec2_instance_type_offering().available().instance_type())
            .vpcSecurityGroupIds(aws_security_group.test().id())
            .subnetId(aws_subnet.test()[0].id())
            .build(), CustomResourceOptions.builder()
                .dependsOn(                
                    aws_route.test(),
                    aws_vpc_dhcp_options_association.test())
                .build());

        var testGateway = new Gateway("testGateway", GatewayArgs.builder()        
            .gatewayIpAddress(testInstance.publicIp())
            .gatewayName("test-sgw")
            .gatewayTimezone("GMT")
            .gatewayType("FILE_FSX_SMB")
            .smbActiveDirectorySettings(GatewaySmbActiveDirectorySettingsArgs.builder()
                .domainName(aws_directory_service_directory.test().name())
                .password(aws_directory_service_directory.test().password())
                .username("Admin")
                .build())
            .build());

        var testWindowsFileSystem = new WindowsFileSystem("testWindowsFileSystem", WindowsFileSystemArgs.builder()        
            .activeDirectoryId(aws_directory_service_directory.test().id())
            .securityGroupIds(aws_security_group.test().id())
            .skipFinalBackup(true)
            .storageCapacity(32)
            .subnetIds(aws_subnet.test()[0].id())
            .throughputCapacity(8)
            .build());

        var fsx = new FileSystemAssociation("fsx", FileSystemAssociationArgs.builder()        
            .gatewayArn(testGateway.arn())
            .locationArn(testWindowsFileSystem.arn())
            .username("Admin")
            .password(aws_directory_service_directory.test().password())
            .cacheAttributes(FileSystemAssociationCacheAttributesArgs.builder()
                .cacheStaleTimeoutInSeconds(400)
                .build())
            .auditDestinationArn(aws_cloudwatch_log_group.test().arn())
            .build());

    }
}
```
```yaml
resources:
  testInstance:
    type: aws:ec2:Instance
    properties:
      ami: ${awsServiceStoragegatewayAmiFILES3Latest.value}
      associatePublicIpAddress: true
      instanceType: ${data.aws_ec2_instance_type_offering.available.instance_type}
      vpcSecurityGroupIds:
        - ${aws_security_group.test.id}
      subnetId: ${aws_subnet.test[0].id}
    options:
      dependson:
        - ${aws_route.test}
        - ${aws_vpc_dhcp_options_association.test}
  testGateway:
    type: aws:storagegateway:Gateway
    properties:
      gatewayIpAddress: ${testInstance.publicIp}
      gatewayName: test-sgw
      gatewayTimezone: GMT
      gatewayType: FILE_FSX_SMB
      smbActiveDirectorySettings:
        domainName: ${aws_directory_service_directory.test.name}
        password: ${aws_directory_service_directory.test.password}
        username: Admin
  testWindowsFileSystem:
    type: aws:fsx:WindowsFileSystem
    properties:
      activeDirectoryId: ${aws_directory_service_directory.test.id}
      securityGroupIds:
        - ${aws_security_group.test.id}
      skipFinalBackup: true
      storageCapacity: 32
      subnetIds:
        - ${aws_subnet.test[0].id}
      throughputCapacity: 8
  fsx:
    type: aws:storagegateway:FileSystemAssociation
    properties:
      gatewayArn: ${testGateway.arn}
      locationArn: ${testWindowsFileSystem.arn}
      username: Admin
      password: ${aws_directory_service_directory.test.password}
      cacheAttributes:
        cacheStaleTimeoutInSeconds: 400
      auditDestinationArn: ${aws_cloudwatch_log_group.test.arn}
variables:
  awsServiceStoragegatewayAmiFILES3Latest:
    fn::invoke:
      Function: aws:ssm:getParameter
      Arguments:
        name: /aws/service/storagegateway/ami/FILE_S3/latest
```


## Import

Using `pulumi import`, import `aws_storagegateway_file_system_association` using the FSx file system association Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:storagegateway/fileSystemAssociation:FileSystemAssociation example arn:aws:storagegateway:us-east-1:123456789012:fs-association/fsa-0DA347732FDB40125
```
 