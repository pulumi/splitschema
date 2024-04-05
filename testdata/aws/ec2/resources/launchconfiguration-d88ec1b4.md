Provides a resource to create a new launch configuration, used for autoscaling groups.

!> **WARNING:** The use of launch configurations is discouraged in favour of launch templates. Read more in the [AWS EC2 Documentation](https://docs.aws.amazon.com/autoscaling/ec2/userguide/launch-configurations.html).

> **Note** When using `aws.ec2.LaunchConfiguration` with `aws.autoscaling.Group`, it is recommended to use the `name_prefix` (Optional) instead of the `name` (Optional) attribute.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ubuntu = aws.ec2.getAmi({
    mostRecent: true,
    filters: [
        {
            name: "name",
            values: ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
        },
        {
            name: "virtualization-type",
            values: ["hvm"],
        },
    ],
    owners: ["099720109477"],
});
const asConf = new aws.ec2.LaunchConfiguration("asConf", {
    imageId: ubuntu.then(ubuntu => ubuntu.id),
    instanceType: "t2.micro",
});
```
```python
import pulumi
import pulumi_aws as aws

ubuntu = aws.ec2.get_ami(most_recent=True,
    filters=[
        aws.ec2.GetAmiFilterArgs(
            name="name",
            values=["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
        ),
        aws.ec2.GetAmiFilterArgs(
            name="virtualization-type",
            values=["hvm"],
        ),
    ],
    owners=["099720109477"])
as_conf = aws.ec2.LaunchConfiguration("asConf",
    image_id=ubuntu.id,
    instance_type="t2.micro")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ubuntu = Aws.Ec2.GetAmi.Invoke(new()
    {
        MostRecent = true,
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
                },
            },
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "virtualization-type",
                Values = new[]
                {
                    "hvm",
                },
            },
        },
        Owners = new[]
        {
            "099720109477",
        },
    });

    var asConf = new Aws.Ec2.LaunchConfiguration("asConf", new()
    {
        ImageId = ubuntu.Apply(getAmiResult => getAmiResult.Id),
        InstanceType = "t2.micro",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ubuntu, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
					},
				},
				{
					Name: "virtualization-type",
					Values: []string{
						"hvm",
					},
				},
			},
			Owners: []string{
				"099720109477",
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewLaunchConfiguration(ctx, "asConf", &ec2.LaunchConfigurationArgs{
			ImageId:      *pulumi.String(ubuntu.Id),
			InstanceType: pulumi.String("t2.micro"),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
import com.pulumi.aws.ec2.LaunchConfiguration;
import com.pulumi.aws.ec2.LaunchConfigurationArgs;
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
        final var ubuntu = Ec2Functions.getAmi(GetAmiArgs.builder()
            .mostRecent(true)
            .filters(            
                GetAmiFilterArgs.builder()
                    .name("name")
                    .values("ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*")
                    .build(),
                GetAmiFilterArgs.builder()
                    .name("virtualization-type")
                    .values("hvm")
                    .build())
            .owners("099720109477")
            .build());

        var asConf = new LaunchConfiguration("asConf", LaunchConfigurationArgs.builder()        
            .imageId(ubuntu.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceType("t2.micro")
            .build());

    }
}
```
```yaml
resources:
  asConf:
    type: aws:ec2:LaunchConfiguration
    properties:
      imageId: ${ubuntu.id}
      instanceType: t2.micro
variables:
  ubuntu:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        mostRecent: true
        filters:
          - name: name
            values:
              - ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*
          - name: virtualization-type
            values:
              - hvm
        owners:
          - '099720109477'
```
{{% /example %}}
{{% /examples %}}
## Using with AutoScaling Groups

Launch Configurations cannot be updated after creation with the Amazon
Web Service API. In order to update a Launch Configuration, this provider will
destroy the existing resource and create a replacement. In order to effectively
use a Launch Configuration resource with an AutoScaling Group resource,
it's recommended to specify `create_before_destroy` in a lifecycle block.
Either omit the Launch Configuration `name` attribute, or specify a partial name
with `name_prefix`.  Example:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ubuntu = aws.ec2.getAmi({
    mostRecent: true,
    filters: [
        {
            name: "name",
            values: ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
        },
        {
            name: "virtualization-type",
            values: ["hvm"],
        },
    ],
    owners: ["099720109477"],
});
const asConf = new aws.ec2.LaunchConfiguration("asConf", {
    namePrefix: "lc-example-",
    imageId: ubuntu.then(ubuntu => ubuntu.id),
    instanceType: "t2.micro",
});
const bar = new aws.autoscaling.Group("bar", {
    launchConfiguration: asConf.name,
    minSize: 1,
    maxSize: 2,
});
```
```python
import pulumi
import pulumi_aws as aws

ubuntu = aws.ec2.get_ami(most_recent=True,
    filters=[
        aws.ec2.GetAmiFilterArgs(
            name="name",
            values=["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
        ),
        aws.ec2.GetAmiFilterArgs(
            name="virtualization-type",
            values=["hvm"],
        ),
    ],
    owners=["099720109477"])
as_conf = aws.ec2.LaunchConfiguration("asConf",
    name_prefix="lc-example-",
    image_id=ubuntu.id,
    instance_type="t2.micro")
bar = aws.autoscaling.Group("bar",
    launch_configuration=as_conf.name,
    min_size=1,
    max_size=2)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ubuntu = Aws.Ec2.GetAmi.Invoke(new()
    {
        MostRecent = true,
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
                },
            },
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "virtualization-type",
                Values = new[]
                {
                    "hvm",
                },
            },
        },
        Owners = new[]
        {
            "099720109477",
        },
    });

    var asConf = new Aws.Ec2.LaunchConfiguration("asConf", new()
    {
        NamePrefix = "lc-example-",
        ImageId = ubuntu.Apply(getAmiResult => getAmiResult.Id),
        InstanceType = "t2.micro",
    });

    var bar = new Aws.AutoScaling.Group("bar", new()
    {
        LaunchConfiguration = asConf.Name,
        MinSize = 1,
        MaxSize = 2,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ubuntu, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
					},
				},
				{
					Name: "virtualization-type",
					Values: []string{
						"hvm",
					},
				},
			},
			Owners: []string{
				"099720109477",
			},
		}, nil)
		if err != nil {
			return err
		}
		asConf, err := ec2.NewLaunchConfiguration(ctx, "asConf", &ec2.LaunchConfigurationArgs{
			NamePrefix:   pulumi.String("lc-example-"),
			ImageId:      *pulumi.String(ubuntu.Id),
			InstanceType: pulumi.String("t2.micro"),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "bar", &autoscaling.GroupArgs{
			LaunchConfiguration: asConf.Name,
			MinSize:             pulumi.Int(1),
			MaxSize:             pulumi.Int(2),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
import com.pulumi.aws.ec2.LaunchConfiguration;
import com.pulumi.aws.ec2.LaunchConfigurationArgs;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
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
        final var ubuntu = Ec2Functions.getAmi(GetAmiArgs.builder()
            .mostRecent(true)
            .filters(            
                GetAmiFilterArgs.builder()
                    .name("name")
                    .values("ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*")
                    .build(),
                GetAmiFilterArgs.builder()
                    .name("virtualization-type")
                    .values("hvm")
                    .build())
            .owners("099720109477")
            .build());

        var asConf = new LaunchConfiguration("asConf", LaunchConfigurationArgs.builder()        
            .namePrefix("lc-example-")
            .imageId(ubuntu.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceType("t2.micro")
            .build());

        var bar = new Group("bar", GroupArgs.builder()        
            .launchConfiguration(asConf.name())
            .minSize(1)
            .maxSize(2)
            .build());

    }
}
```
```yaml
resources:
  asConf:
    type: aws:ec2:LaunchConfiguration
    properties:
      namePrefix: lc-example-
      imageId: ${ubuntu.id}
      instanceType: t2.micro
  bar:
    type: aws:autoscaling:Group
    properties:
      launchConfiguration: ${asConf.name}
      minSize: 1
      maxSize: 2
variables:
  ubuntu:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        mostRecent: true
        filters:
          - name: name
            values:
              - ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*
          - name: virtualization-type
            values:
              - hvm
        owners:
          - '099720109477'
```

With this setup this provider generates a unique name for your Launch
Configuration and can then update the AutoScaling Group without conflict before
destroying the previous Launch Configuration.

## Using with Spot Instances

Launch configurations can set the spot instance pricing to be used for the
Auto Scaling Group to reserve instances. Simply specifying the `spot_price`
parameter will set the price on the Launch Configuration which will attempt to
reserve your instances at this price.  See the [AWS Spot Instance
documentation](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-spot-instances.html)
for more information or how to launch [Spot Instances][3] with this provider.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ubuntu = aws.ec2.getAmi({
    mostRecent: true,
    filters: [
        {
            name: "name",
            values: ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
        },
        {
            name: "virtualization-type",
            values: ["hvm"],
        },
    ],
    owners: ["099720109477"],
});
const asConf = new aws.ec2.LaunchConfiguration("asConf", {
    imageId: ubuntu.then(ubuntu => ubuntu.id),
    instanceType: "m4.large",
    spotPrice: "0.001",
});
const bar = new aws.autoscaling.Group("bar", {launchConfiguration: asConf.name});
```
```python
import pulumi
import pulumi_aws as aws

ubuntu = aws.ec2.get_ami(most_recent=True,
    filters=[
        aws.ec2.GetAmiFilterArgs(
            name="name",
            values=["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"],
        ),
        aws.ec2.GetAmiFilterArgs(
            name="virtualization-type",
            values=["hvm"],
        ),
    ],
    owners=["099720109477"])
as_conf = aws.ec2.LaunchConfiguration("asConf",
    image_id=ubuntu.id,
    instance_type="m4.large",
    spot_price="0.001")
bar = aws.autoscaling.Group("bar", launch_configuration=as_conf.name)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var ubuntu = Aws.Ec2.GetAmi.Invoke(new()
    {
        MostRecent = true,
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
                },
            },
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "virtualization-type",
                Values = new[]
                {
                    "hvm",
                },
            },
        },
        Owners = new[]
        {
            "099720109477",
        },
    });

    var asConf = new Aws.Ec2.LaunchConfiguration("asConf", new()
    {
        ImageId = ubuntu.Apply(getAmiResult => getAmiResult.Id),
        InstanceType = "m4.large",
        SpotPrice = "0.001",
    });

    var bar = new Aws.AutoScaling.Group("bar", new()
    {
        LaunchConfiguration = asConf.Name,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/autoscaling"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		ubuntu, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*",
					},
				},
				{
					Name: "virtualization-type",
					Values: []string{
						"hvm",
					},
				},
			},
			Owners: []string{
				"099720109477",
			},
		}, nil)
		if err != nil {
			return err
		}
		asConf, err := ec2.NewLaunchConfiguration(ctx, "asConf", &ec2.LaunchConfigurationArgs{
			ImageId:      *pulumi.String(ubuntu.Id),
			InstanceType: pulumi.String("m4.large"),
			SpotPrice:    pulumi.String("0.001"),
		})
		if err != nil {
			return err
		}
		_, err = autoscaling.NewGroup(ctx, "bar", &autoscaling.GroupArgs{
			LaunchConfiguration: asConf.Name,
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
import com.pulumi.aws.ec2.LaunchConfiguration;
import com.pulumi.aws.ec2.LaunchConfigurationArgs;
import com.pulumi.aws.autoscaling.Group;
import com.pulumi.aws.autoscaling.GroupArgs;
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
        final var ubuntu = Ec2Functions.getAmi(GetAmiArgs.builder()
            .mostRecent(true)
            .filters(            
                GetAmiFilterArgs.builder()
                    .name("name")
                    .values("ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*")
                    .build(),
                GetAmiFilterArgs.builder()
                    .name("virtualization-type")
                    .values("hvm")
                    .build())
            .owners("099720109477")
            .build());

        var asConf = new LaunchConfiguration("asConf", LaunchConfigurationArgs.builder()        
            .imageId(ubuntu.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceType("m4.large")
            .spotPrice("0.001")
            .build());

        var bar = new Group("bar", GroupArgs.builder()        
            .launchConfiguration(asConf.name())
            .build());

    }
}
```
```yaml
resources:
  asConf:
    type: aws:ec2:LaunchConfiguration
    properties:
      imageId: ${ubuntu.id}
      instanceType: m4.large
      spotPrice: '0.001'
  bar:
    type: aws:autoscaling:Group
    properties:
      launchConfiguration: ${asConf.name}
variables:
  ubuntu:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        mostRecent: true
        filters:
          - name: name
            values:
              - ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*
          - name: virtualization-type
            values:
              - hvm
        owners:
          - '099720109477'
```

## Block devices

Each of the `*_block_device` attributes controls a portion of the AWS
Launch Configuration's "Block Device Mapping". It's a good idea to familiarize yourself with [AWS's Block Device
Mapping docs](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/block-device-mapping-concepts.html)
to understand the implications of using these attributes.

Each AWS Instance type has a different set of Instance Store block devices
available for attachment. AWS [publishes a
list](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#StorageOnInstanceTypes)
of which ephemeral devices are available on each type. The devices are always
identified by the `virtual_name` in the format `ephemeral{0..N}`.

> **NOTE:** Changes to `*_block_device` configuration of _existing_ resources
cannot currently be detected by this provider. After updating to block device
configuration, resource recreation can be manually triggered by using the
[`up` command with the --replace argument](https://www.pulumi.com/docs/reference/cli/pulumi_up/).
  
### ebs_block_device

Modifying any of the `ebs_block_device` settings requires resource replacement.

* `device_name` - (Required) The name of the device to mount.
* `snapshot_id` - (Optional) The Snapshot ID to mount.
* `volume_type` - (Optional) The type of volume. Can be `standard`, `gp2`, `gp3`, `st1`, `sc1` or `io1`.
* `volume_size` - (Optional) The size of the volume in gigabytes.
* `iops` - (Optional) The amount of provisioned
  [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html).
  This must be set with a `volume_type` of `"io1"`.
* `throughput` - (Optional) The throughput (MiBps) to provision for a `gp3` volume.
* `delete_on_termination` - (Optional) Whether the volume should be destroyed
  on instance termination (Default: `true`).
* `encrypted` - (Optional) Whether the volume should be encrypted or not. Defaults to `false`.
* `no_device` - (Optional) Whether the device in the block device mapping of the AMI is suppressed.

### ephemeral_block_device

* `device_name` - (Required) The name of the block device to mount on the instance.
* `no_device` - (Optional) Whether the device in the block device mapping of the AMI is suppressed.
* `virtual_name` - (Optional) The [Instance Store Device Name](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/InstanceStorage.html#InstanceStoreDeviceNames).

### root_block_device

> Modifying any of the `root_block_device` settings requires resource replacement.

* `delete_on_termination` - (Optional) Whether the volume should be destroyed on instance termination. Defaults to `true`.
* `encrypted` - (Optional) Whether the volume should be encrypted or not. Defaults to `false`.
* `iops` - (Optional) The amount of provisioned [IOPS](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ebs-io-characteristics.html). This must be set with a `volume_type` of `io1`.
* `throughput` - (Optional) The throughput (MiBps) to provision for a `gp3` volume.
* `volume_size` - (Optional) The size of the volume in gigabytes.
* `volume_type` - (Optional) The type of volume. Can be `standard`, `gp2`, `gp3`, `st1`, `sc1` or `io1`.


## Import

Using `pulumi import`, import launch configurations using the `name`. For example:

```sh
 $ pulumi import aws:ec2/launchConfiguration:LaunchConfiguration as_conf pulumi-lg-123456
```
 