Manages an AWS Storage Gateway file, tape, or volume gateway in the provider region.

> **NOTE:** The Storage Gateway API requires the gateway to be connected to properly return information after activation. If you are receiving `The specified gateway is not connected` errors during resource creation (gateway activation), ensure your gateway instance meets the [Storage Gateway requirements](https://docs.aws.amazon.com/storagegateway/latest/userguide/Requirements.html).

{{% examples %}}
## Example Usage
{{% example %}}
### Local Cache

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testVolumeAttachment = new aws.ec2.VolumeAttachment("testVolumeAttachment", {
    deviceName: "/dev/xvdb",
    volumeId: aws_ebs_volume.test.id,
    instanceId: aws_instance.test.id,
});
const testLocalDisk = testVolumeAttachment.deviceName.apply(deviceName => aws.storagegateway.getLocalDiskOutput({
    diskNode: deviceName,
    gatewayArn: aws_storagegateway_gateway.test.arn,
}));
const testCache = new aws.storagegateway.Cache("testCache", {
    diskId: testLocalDisk.apply(testLocalDisk => testLocalDisk.diskId),
    gatewayArn: aws_storagegateway_gateway.test.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

test_volume_attachment = aws.ec2.VolumeAttachment("testVolumeAttachment",
    device_name="/dev/xvdb",
    volume_id=aws_ebs_volume["test"]["id"],
    instance_id=aws_instance["test"]["id"])
test_local_disk = test_volume_attachment.device_name.apply(lambda device_name: aws.storagegateway.get_local_disk_output(disk_node=device_name,
    gateway_arn=aws_storagegateway_gateway["test"]["arn"]))
test_cache = aws.storagegateway.Cache("testCache",
    disk_id=test_local_disk.disk_id,
    gateway_arn=aws_storagegateway_gateway["test"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testVolumeAttachment = new Aws.Ec2.VolumeAttachment("testVolumeAttachment", new()
    {
        DeviceName = "/dev/xvdb",
        VolumeId = aws_ebs_volume.Test.Id,
        InstanceId = aws_instance.Test.Id,
    });

    var testLocalDisk = Aws.StorageGateway.GetLocalDisk.Invoke(new()
    {
        DiskNode = testVolumeAttachment.DeviceName,
        GatewayArn = aws_storagegateway_gateway.Test.Arn,
    });

    var testCache = new Aws.StorageGateway.Cache("testCache", new()
    {
        DiskId = testLocalDisk.Apply(getLocalDiskResult => getLocalDiskResult.DiskId),
        GatewayArn = aws_storagegateway_gateway.Test.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/storagegateway"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testVolumeAttachment, err := ec2.NewVolumeAttachment(ctx, "testVolumeAttachment", &ec2.VolumeAttachmentArgs{
			DeviceName: pulumi.String("/dev/xvdb"),
			VolumeId:   pulumi.Any(aws_ebs_volume.Test.Id),
			InstanceId: pulumi.Any(aws_instance.Test.Id),
		})
		if err != nil {
			return err
		}
		testLocalDisk := testVolumeAttachment.DeviceName.ApplyT(func(deviceName string) (storagegateway.GetLocalDiskResult, error) {
			return storagegateway.GetLocalDiskOutput(ctx, storagegateway.GetLocalDiskOutputArgs{
				DiskNode:   deviceName,
				GatewayArn: aws_storagegateway_gateway.Test.Arn,
			}, nil), nil
		}).(storagegateway.GetLocalDiskResultOutput)
		_, err = storagegateway.NewCache(ctx, "testCache", &storagegateway.CacheArgs{
			DiskId: testLocalDisk.ApplyT(func(testLocalDisk storagegateway.GetLocalDiskResult) (*string, error) {
				return &testLocalDisk.DiskId, nil
			}).(pulumi.StringPtrOutput),
			GatewayArn: pulumi.Any(aws_storagegateway_gateway.Test.Arn),
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
import com.pulumi.aws.ec2.VolumeAttachment;
import com.pulumi.aws.ec2.VolumeAttachmentArgs;
import com.pulumi.aws.storagegateway.StoragegatewayFunctions;
import com.pulumi.aws.storagegateway.inputs.GetLocalDiskArgs;
import com.pulumi.aws.storagegateway.Cache;
import com.pulumi.aws.storagegateway.CacheArgs;
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
        var testVolumeAttachment = new VolumeAttachment("testVolumeAttachment", VolumeAttachmentArgs.builder()        
            .deviceName("/dev/xvdb")
            .volumeId(aws_ebs_volume.test().id())
            .instanceId(aws_instance.test().id())
            .build());

        final var testLocalDisk = StoragegatewayFunctions.getLocalDisk(GetLocalDiskArgs.builder()
            .diskNode(testVolumeAttachment.deviceName())
            .gatewayArn(aws_storagegateway_gateway.test().arn())
            .build());

        var testCache = new Cache("testCache", CacheArgs.builder()        
            .diskId(testLocalDisk.applyValue(getLocalDiskResult -> getLocalDiskResult).applyValue(testLocalDisk -> testLocalDisk.applyValue(getLocalDiskResult -> getLocalDiskResult.diskId())))
            .gatewayArn(aws_storagegateway_gateway.test().arn())
            .build());

    }
}
```
```yaml
resources:
  testVolumeAttachment:
    type: aws:ec2:VolumeAttachment
    properties:
      deviceName: /dev/xvdb
      volumeId: ${aws_ebs_volume.test.id}
      instanceId: ${aws_instance.test.id}
  testCache:
    type: aws:storagegateway:Cache
    properties:
      diskId: ${testLocalDisk.diskId}
      gatewayArn: ${aws_storagegateway_gateway.test.arn}
variables:
  testLocalDisk:
    fn::invoke:
      Function: aws:storagegateway:getLocalDisk
      Arguments:
        diskNode: ${testVolumeAttachment.deviceName}
        gatewayArn: ${aws_storagegateway_gateway.test.arn}
```
{{% /example %}}
{{% example %}}
### FSx File Gateway

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.Gateway("example", {
    gatewayIpAddress: "1.2.3.4",
    gatewayName: "example",
    gatewayTimezone: "GMT",
    gatewayType: "FILE_FSX_SMB",
    smbActiveDirectorySettings: {
        domainName: "corp.example.com",
        password: "avoid-plaintext-passwords",
        username: "Admin",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.Gateway("example",
    gateway_ip_address="1.2.3.4",
    gateway_name="example",
    gateway_timezone="GMT",
    gateway_type="FILE_FSX_SMB",
    smb_active_directory_settings=aws.storagegateway.GatewaySmbActiveDirectorySettingsArgs(
        domain_name="corp.example.com",
        password="avoid-plaintext-passwords",
        username="Admin",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.Gateway("example", new()
    {
        GatewayIpAddress = "1.2.3.4",
        GatewayName = "example",
        GatewayTimezone = "GMT",
        GatewayType = "FILE_FSX_SMB",
        SmbActiveDirectorySettings = new Aws.StorageGateway.Inputs.GatewaySmbActiveDirectorySettingsArgs
        {
            DomainName = "corp.example.com",
            Password = "avoid-plaintext-passwords",
            Username = "Admin",
        },
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
		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
			GatewayIpAddress: pulumi.String("1.2.3.4"),
			GatewayName:      pulumi.String("example"),
			GatewayTimezone:  pulumi.String("GMT"),
			GatewayType:      pulumi.String("FILE_FSX_SMB"),
			SmbActiveDirectorySettings: &storagegateway.GatewaySmbActiveDirectorySettingsArgs{
				DomainName: pulumi.String("corp.example.com"),
				Password:   pulumi.String("avoid-plaintext-passwords"),
				Username:   pulumi.String("Admin"),
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
import com.pulumi.aws.storagegateway.Gateway;
import com.pulumi.aws.storagegateway.GatewayArgs;
import com.pulumi.aws.storagegateway.inputs.GatewaySmbActiveDirectorySettingsArgs;
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
        var example = new Gateway("example", GatewayArgs.builder()        
            .gatewayIpAddress("1.2.3.4")
            .gatewayName("example")
            .gatewayTimezone("GMT")
            .gatewayType("FILE_FSX_SMB")
            .smbActiveDirectorySettings(GatewaySmbActiveDirectorySettingsArgs.builder()
                .domainName("corp.example.com")
                .password("avoid-plaintext-passwords")
                .username("Admin")
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:Gateway
    properties:
      gatewayIpAddress: 1.2.3.4
      gatewayName: example
      gatewayTimezone: GMT
      gatewayType: FILE_FSX_SMB
      smbActiveDirectorySettings:
        domainName: corp.example.com
        password: avoid-plaintext-passwords
        username: Admin
```
{{% /example %}}
{{% example %}}
### S3 File Gateway

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.Gateway("example", {
    gatewayIpAddress: "1.2.3.4",
    gatewayName: "example",
    gatewayTimezone: "GMT",
    gatewayType: "FILE_S3",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.Gateway("example",
    gateway_ip_address="1.2.3.4",
    gateway_name="example",
    gateway_timezone="GMT",
    gateway_type="FILE_S3")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.Gateway("example", new()
    {
        GatewayIpAddress = "1.2.3.4",
        GatewayName = "example",
        GatewayTimezone = "GMT",
        GatewayType = "FILE_S3",
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
		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
			GatewayIpAddress: pulumi.String("1.2.3.4"),
			GatewayName:      pulumi.String("example"),
			GatewayTimezone:  pulumi.String("GMT"),
			GatewayType:      pulumi.String("FILE_S3"),
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
import com.pulumi.aws.storagegateway.Gateway;
import com.pulumi.aws.storagegateway.GatewayArgs;
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
        var example = new Gateway("example", GatewayArgs.builder()        
            .gatewayIpAddress("1.2.3.4")
            .gatewayName("example")
            .gatewayTimezone("GMT")
            .gatewayType("FILE_S3")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:Gateway
    properties:
      gatewayIpAddress: 1.2.3.4
      gatewayName: example
      gatewayTimezone: GMT
      gatewayType: FILE_S3
```
{{% /example %}}
{{% example %}}
### Tape Gateway

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.Gateway("example", {
    gatewayIpAddress: "1.2.3.4",
    gatewayName: "example",
    gatewayTimezone: "GMT",
    gatewayType: "VTL",
    mediumChangerType: "AWS-Gateway-VTL",
    tapeDriveType: "IBM-ULT3580-TD5",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.Gateway("example",
    gateway_ip_address="1.2.3.4",
    gateway_name="example",
    gateway_timezone="GMT",
    gateway_type="VTL",
    medium_changer_type="AWS-Gateway-VTL",
    tape_drive_type="IBM-ULT3580-TD5")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.Gateway("example", new()
    {
        GatewayIpAddress = "1.2.3.4",
        GatewayName = "example",
        GatewayTimezone = "GMT",
        GatewayType = "VTL",
        MediumChangerType = "AWS-Gateway-VTL",
        TapeDriveType = "IBM-ULT3580-TD5",
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
		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
			GatewayIpAddress:  pulumi.String("1.2.3.4"),
			GatewayName:       pulumi.String("example"),
			GatewayTimezone:   pulumi.String("GMT"),
			GatewayType:       pulumi.String("VTL"),
			MediumChangerType: pulumi.String("AWS-Gateway-VTL"),
			TapeDriveType:     pulumi.String("IBM-ULT3580-TD5"),
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
import com.pulumi.aws.storagegateway.Gateway;
import com.pulumi.aws.storagegateway.GatewayArgs;
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
        var example = new Gateway("example", GatewayArgs.builder()        
            .gatewayIpAddress("1.2.3.4")
            .gatewayName("example")
            .gatewayTimezone("GMT")
            .gatewayType("VTL")
            .mediumChangerType("AWS-Gateway-VTL")
            .tapeDriveType("IBM-ULT3580-TD5")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:Gateway
    properties:
      gatewayIpAddress: 1.2.3.4
      gatewayName: example
      gatewayTimezone: GMT
      gatewayType: VTL
      mediumChangerType: AWS-Gateway-VTL
      tapeDriveType: IBM-ULT3580-TD5
```
{{% /example %}}
{{% example %}}
### Volume Gateway (Cached)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.Gateway("example", {
    gatewayIpAddress: "1.2.3.4",
    gatewayName: "example",
    gatewayTimezone: "GMT",
    gatewayType: "CACHED",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.Gateway("example",
    gateway_ip_address="1.2.3.4",
    gateway_name="example",
    gateway_timezone="GMT",
    gateway_type="CACHED")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.Gateway("example", new()
    {
        GatewayIpAddress = "1.2.3.4",
        GatewayName = "example",
        GatewayTimezone = "GMT",
        GatewayType = "CACHED",
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
		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
			GatewayIpAddress: pulumi.String("1.2.3.4"),
			GatewayName:      pulumi.String("example"),
			GatewayTimezone:  pulumi.String("GMT"),
			GatewayType:      pulumi.String("CACHED"),
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
import com.pulumi.aws.storagegateway.Gateway;
import com.pulumi.aws.storagegateway.GatewayArgs;
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
        var example = new Gateway("example", GatewayArgs.builder()        
            .gatewayIpAddress("1.2.3.4")
            .gatewayName("example")
            .gatewayTimezone("GMT")
            .gatewayType("CACHED")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:Gateway
    properties:
      gatewayIpAddress: 1.2.3.4
      gatewayName: example
      gatewayTimezone: GMT
      gatewayType: CACHED
```
{{% /example %}}
{{% example %}}
### Volume Gateway (Stored)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.storagegateway.Gateway("example", {
    gatewayIpAddress: "1.2.3.4",
    gatewayName: "example",
    gatewayTimezone: "GMT",
    gatewayType: "STORED",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.storagegateway.Gateway("example",
    gateway_ip_address="1.2.3.4",
    gateway_name="example",
    gateway_timezone="GMT",
    gateway_type="STORED")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.StorageGateway.Gateway("example", new()
    {
        GatewayIpAddress = "1.2.3.4",
        GatewayName = "example",
        GatewayTimezone = "GMT",
        GatewayType = "STORED",
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
		_, err := storagegateway.NewGateway(ctx, "example", &storagegateway.GatewayArgs{
			GatewayIpAddress: pulumi.String("1.2.3.4"),
			GatewayName:      pulumi.String("example"),
			GatewayTimezone:  pulumi.String("GMT"),
			GatewayType:      pulumi.String("STORED"),
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
import com.pulumi.aws.storagegateway.Gateway;
import com.pulumi.aws.storagegateway.GatewayArgs;
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
        var example = new Gateway("example", GatewayArgs.builder()        
            .gatewayIpAddress("1.2.3.4")
            .gatewayName("example")
            .gatewayTimezone("GMT")
            .gatewayType("STORED")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:storagegateway:Gateway
    properties:
      gatewayIpAddress: 1.2.3.4
      gatewayName: example
      gatewayTimezone: GMT
      gatewayType: STORED
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_storagegateway_gateway` using the gateway Amazon Resource Name (ARN). For example:

```sh
 $ pulumi import aws:storagegateway/gateway:Gateway example arn:aws:storagegateway:us-east-1:123456789012:gateway/sgw-12345678
```
 Certain resource arguments, like `gateway_ip_address` do not have a Storage Gateway API method for reading the information after creation, either omit the argument from the Pulumi program or use `ignore_changes` to hide the difference. For example:

