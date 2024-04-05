Provides an AWS EBS Volume Attachment as a top level resource, to attach and
detach volumes from AWS Instances.

> **NOTE on EBS block devices:** If you use `ebs_block_device` on an `aws.ec2.Instance`, this provider will assume management over the full set of non-root EBS block devices for the instance, and treats additional block devices as drift. For this reason, `ebs_block_device` cannot be mixed with external `aws.ebs.Volume` + `aws.ec2.VolumeAttachment` resources for a given instance.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const web = new aws.ec2.Instance("web", {
    ami: "ami-21f78e11",
    availabilityZone: "us-west-2a",
    instanceType: "t2.micro",
    tags: {
        Name: "HelloWorld",
    },
});
const example = new aws.ebs.Volume("example", {
    availabilityZone: "us-west-2a",
    size: 1,
});
const ebsAtt = new aws.ec2.VolumeAttachment("ebsAtt", {
    deviceName: "/dev/sdh",
    volumeId: example.id,
    instanceId: web.id,
});
```
```python
import pulumi
import pulumi_aws as aws

web = aws.ec2.Instance("web",
    ami="ami-21f78e11",
    availability_zone="us-west-2a",
    instance_type="t2.micro",
    tags={
        "Name": "HelloWorld",
    })
example = aws.ebs.Volume("example",
    availability_zone="us-west-2a",
    size=1)
ebs_att = aws.ec2.VolumeAttachment("ebsAtt",
    device_name="/dev/sdh",
    volume_id=example.id,
    instance_id=web.id)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var web = new Aws.Ec2.Instance("web", new()
    {
        Ami = "ami-21f78e11",
        AvailabilityZone = "us-west-2a",
        InstanceType = "t2.micro",
        Tags = 
        {
            { "Name", "HelloWorld" },
        },
    });

    var example = new Aws.Ebs.Volume("example", new()
    {
        AvailabilityZone = "us-west-2a",
        Size = 1,
    });

    var ebsAtt = new Aws.Ec2.VolumeAttachment("ebsAtt", new()
    {
        DeviceName = "/dev/sdh",
        VolumeId = example.Id,
        InstanceId = web.Id,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ebs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		web, err := ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
			Ami:              pulumi.String("ami-21f78e11"),
			AvailabilityZone: pulumi.String("us-west-2a"),
			InstanceType:     pulumi.String("t2.micro"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
			},
		})
		if err != nil {
			return err
		}
		example, err := ebs.NewVolume(ctx, "example", &ebs.VolumeArgs{
			AvailabilityZone: pulumi.String("us-west-2a"),
			Size:             pulumi.Int(1),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewVolumeAttachment(ctx, "ebsAtt", &ec2.VolumeAttachmentArgs{
			DeviceName: pulumi.String("/dev/sdh"),
			VolumeId:   example.ID(),
			InstanceId: web.ID(),
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
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import com.pulumi.aws.ebs.Volume;
import com.pulumi.aws.ebs.VolumeArgs;
import com.pulumi.aws.ec2.VolumeAttachment;
import com.pulumi.aws.ec2.VolumeAttachmentArgs;
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
        var web = new Instance("web", InstanceArgs.builder()        
            .ami("ami-21f78e11")
            .availabilityZone("us-west-2a")
            .instanceType("t2.micro")
            .tags(Map.of("Name", "HelloWorld"))
            .build());

        var example = new Volume("example", VolumeArgs.builder()        
            .availabilityZone("us-west-2a")
            .size(1)
            .build());

        var ebsAtt = new VolumeAttachment("ebsAtt", VolumeAttachmentArgs.builder()        
            .deviceName("/dev/sdh")
            .volumeId(example.id())
            .instanceId(web.id())
            .build());

    }
}
```
```yaml
resources:
  ebsAtt:
    type: aws:ec2:VolumeAttachment
    properties:
      deviceName: /dev/sdh
      volumeId: ${example.id}
      instanceId: ${web.id}
  web:
    type: aws:ec2:Instance
    properties:
      ami: ami-21f78e11
      availabilityZone: us-west-2a
      instanceType: t2.micro
      tags:
        Name: HelloWorld
  example:
    type: aws:ebs:Volume
    properties:
      availabilityZone: us-west-2a
      size: 1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EBS Volume Attachments using `DEVICE_NAME:VOLUME_ID:INSTANCE_ID`. For example:

```sh
 $ pulumi import aws:ec2/volumeAttachment:VolumeAttachment example /dev/sdh:vol-049df61146c4d7901:i-12345678
```
 