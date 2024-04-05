Provides an EC2 Spot Fleet Request resource. This allows a fleet of Spot
instances to be requested on the Spot market.

> **NOTE [AWS strongly discourages](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-best-practices.html#which-spot-request-method-to-use) the use of the legacy APIs called by this resource.
We recommend using the EC2 Fleet or Auto Scaling Group resources instead.

{{% examples %}}
## Example Usage
{{% example %}}
### Using launch specifications

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// Request a Spot fleet
const cheapCompute = new aws.ec2.SpotFleetRequest("cheapCompute", {
    iamFleetRole: "arn:aws:iam::12345678:role/spot-fleet",
    spotPrice: "0.03",
    allocationStrategy: "diversified",
    targetCapacity: 6,
    validUntil: "2019-11-04T20:44:20Z",
    launchSpecifications: [
        {
            instanceType: "m4.10xlarge",
            ami: "ami-1234",
            spotPrice: "2.793",
            placementTenancy: "dedicated",
            iamInstanceProfileArn: aws_iam_instance_profile.example.arn,
        },
        {
            instanceType: "m4.4xlarge",
            ami: "ami-5678",
            keyName: "my-key",
            spotPrice: "1.117",
            iamInstanceProfileArn: aws_iam_instance_profile.example.arn,
            availabilityZone: "us-west-1a",
            subnetId: "subnet-1234",
            weightedCapacity: "35",
            rootBlockDevices: [{
                volumeSize: 300,
                volumeType: "gp2",
            }],
            tags: {
                Name: "spot-fleet-example",
            },
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

# Request a Spot fleet
cheap_compute = aws.ec2.SpotFleetRequest("cheapCompute",
    iam_fleet_role="arn:aws:iam::12345678:role/spot-fleet",
    spot_price="0.03",
    allocation_strategy="diversified",
    target_capacity=6,
    valid_until="2019-11-04T20:44:20Z",
    launch_specifications=[
        aws.ec2.SpotFleetRequestLaunchSpecificationArgs(
            instance_type="m4.10xlarge",
            ami="ami-1234",
            spot_price="2.793",
            placement_tenancy="dedicated",
            iam_instance_profile_arn=aws_iam_instance_profile["example"]["arn"],
        ),
        aws.ec2.SpotFleetRequestLaunchSpecificationArgs(
            instance_type="m4.4xlarge",
            ami="ami-5678",
            key_name="my-key",
            spot_price="1.117",
            iam_instance_profile_arn=aws_iam_instance_profile["example"]["arn"],
            availability_zone="us-west-1a",
            subnet_id="subnet-1234",
            weighted_capacity="35",
            root_block_devices=[aws.ec2.SpotFleetRequestLaunchSpecificationRootBlockDeviceArgs(
                volume_size=300,
                volume_type="gp2",
            )],
            tags={
                "Name": "spot-fleet-example",
            },
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // Request a Spot fleet
    var cheapCompute = new Aws.Ec2.SpotFleetRequest("cheapCompute", new()
    {
        IamFleetRole = "arn:aws:iam::12345678:role/spot-fleet",
        SpotPrice = "0.03",
        AllocationStrategy = "diversified",
        TargetCapacity = 6,
        ValidUntil = "2019-11-04T20:44:20Z",
        LaunchSpecifications = new[]
        {
            new Aws.Ec2.Inputs.SpotFleetRequestLaunchSpecificationArgs
            {
                InstanceType = "m4.10xlarge",
                Ami = "ami-1234",
                SpotPrice = "2.793",
                PlacementTenancy = "dedicated",
                IamInstanceProfileArn = aws_iam_instance_profile.Example.Arn,
            },
            new Aws.Ec2.Inputs.SpotFleetRequestLaunchSpecificationArgs
            {
                InstanceType = "m4.4xlarge",
                Ami = "ami-5678",
                KeyName = "my-key",
                SpotPrice = "1.117",
                IamInstanceProfileArn = aws_iam_instance_profile.Example.Arn,
                AvailabilityZone = "us-west-1a",
                SubnetId = "subnet-1234",
                WeightedCapacity = "35",
                RootBlockDevices = new[]
                {
                    new Aws.Ec2.Inputs.SpotFleetRequestLaunchSpecificationRootBlockDeviceArgs
                    {
                        VolumeSize = 300,
                        VolumeType = "gp2",
                    },
                },
                Tags = 
                {
                    { "Name", "spot-fleet-example" },
                },
            },
        },
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
		_, err := ec2.NewSpotFleetRequest(ctx, "cheapCompute", &ec2.SpotFleetRequestArgs{
			IamFleetRole:       pulumi.String("arn:aws:iam::12345678:role/spot-fleet"),
			SpotPrice:          pulumi.String("0.03"),
			AllocationStrategy: pulumi.String("diversified"),
			TargetCapacity:     pulumi.Int(6),
			ValidUntil:         pulumi.String("2019-11-04T20:44:20Z"),
			LaunchSpecifications: ec2.SpotFleetRequestLaunchSpecificationArray{
				&ec2.SpotFleetRequestLaunchSpecificationArgs{
					InstanceType:          pulumi.String("m4.10xlarge"),
					Ami:                   pulumi.String("ami-1234"),
					SpotPrice:             pulumi.String("2.793"),
					PlacementTenancy:      pulumi.String("dedicated"),
					IamInstanceProfileArn: pulumi.Any(aws_iam_instance_profile.Example.Arn),
				},
				&ec2.SpotFleetRequestLaunchSpecificationArgs{
					InstanceType:          pulumi.String("m4.4xlarge"),
					Ami:                   pulumi.String("ami-5678"),
					KeyName:               pulumi.String("my-key"),
					SpotPrice:             pulumi.String("1.117"),
					IamInstanceProfileArn: pulumi.Any(aws_iam_instance_profile.Example.Arn),
					AvailabilityZone:      pulumi.String("us-west-1a"),
					SubnetId:              pulumi.String("subnet-1234"),
					WeightedCapacity:      pulumi.String("35"),
					RootBlockDevices: ec2.SpotFleetRequestLaunchSpecificationRootBlockDeviceArray{
						&ec2.SpotFleetRequestLaunchSpecificationRootBlockDeviceArgs{
							VolumeSize: pulumi.Int(300),
							VolumeType: pulumi.String("gp2"),
						},
					},
					Tags: pulumi.StringMap{
						"Name": pulumi.String("spot-fleet-example"),
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
import com.pulumi.aws.ec2.SpotFleetRequest;
import com.pulumi.aws.ec2.SpotFleetRequestArgs;
import com.pulumi.aws.ec2.inputs.SpotFleetRequestLaunchSpecificationArgs;
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
        var cheapCompute = new SpotFleetRequest("cheapCompute", SpotFleetRequestArgs.builder()        
            .iamFleetRole("arn:aws:iam::12345678:role/spot-fleet")
            .spotPrice("0.03")
            .allocationStrategy("diversified")
            .targetCapacity(6)
            .validUntil("2019-11-04T20:44:20Z")
            .launchSpecifications(            
                SpotFleetRequestLaunchSpecificationArgs.builder()
                    .instanceType("m4.10xlarge")
                    .ami("ami-1234")
                    .spotPrice("2.793")
                    .placementTenancy("dedicated")
                    .iamInstanceProfileArn(aws_iam_instance_profile.example().arn())
                    .build(),
                SpotFleetRequestLaunchSpecificationArgs.builder()
                    .instanceType("m4.4xlarge")
                    .ami("ami-5678")
                    .keyName("my-key")
                    .spotPrice("1.117")
                    .iamInstanceProfileArn(aws_iam_instance_profile.example().arn())
                    .availabilityZone("us-west-1a")
                    .subnetId("subnet-1234")
                    .weightedCapacity(35)
                    .rootBlockDevices(SpotFleetRequestLaunchSpecificationRootBlockDeviceArgs.builder()
                        .volumeSize("300")
                        .volumeType("gp2")
                        .build())
                    .tags(Map.of("Name", "spot-fleet-example"))
                    .build())
            .build());

    }
}
```
```yaml
resources:
  # Request a Spot fleet
  cheapCompute:
    type: aws:ec2:SpotFleetRequest
    properties:
      iamFleetRole: arn:aws:iam::12345678:role/spot-fleet
      spotPrice: '0.03'
      allocationStrategy: diversified
      targetCapacity: 6
      validUntil: 2019-11-04T20:44:20Z
      launchSpecifications:
        - instanceType: m4.10xlarge
          ami: ami-1234
          spotPrice: '2.793'
          placementTenancy: dedicated
          iamInstanceProfileArn: ${aws_iam_instance_profile.example.arn}
        - instanceType: m4.4xlarge
          ami: ami-5678
          keyName: my-key
          spotPrice: '1.117'
          iamInstanceProfileArn: ${aws_iam_instance_profile.example.arn}
          availabilityZone: us-west-1a
          subnetId: subnet-1234
          weightedCapacity: 35
          rootBlockDevices:
            - volumeSize: '300'
              volumeType: gp2
          tags:
            Name: spot-fleet-example
```
{{% /example %}}
{{% example %}}
### Using launch templates

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const fooLaunchTemplate = new aws.ec2.LaunchTemplate("fooLaunchTemplate", {
    imageId: "ami-516b9131",
    instanceType: "m1.small",
    keyName: "some-key",
});
const fooSpotFleetRequest = new aws.ec2.SpotFleetRequest("fooSpotFleetRequest", {
    iamFleetRole: "arn:aws:iam::12345678:role/spot-fleet",
    spotPrice: "0.005",
    targetCapacity: 2,
    validUntil: "2019-11-04T20:44:20Z",
    launchTemplateConfigs: [{
        launchTemplateSpecification: {
            id: fooLaunchTemplate.id,
            version: fooLaunchTemplate.latestVersion,
        },
    }],
}, {
    dependsOn: [aws_iam_policy_attachment["test-attach"]],
});
```
```python
import pulumi
import pulumi_aws as aws

foo_launch_template = aws.ec2.LaunchTemplate("fooLaunchTemplate",
    image_id="ami-516b9131",
    instance_type="m1.small",
    key_name="some-key")
foo_spot_fleet_request = aws.ec2.SpotFleetRequest("fooSpotFleetRequest",
    iam_fleet_role="arn:aws:iam::12345678:role/spot-fleet",
    spot_price="0.005",
    target_capacity=2,
    valid_until="2019-11-04T20:44:20Z",
    launch_template_configs=[aws.ec2.SpotFleetRequestLaunchTemplateConfigArgs(
        launch_template_specification=aws.ec2.SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs(
            id=foo_launch_template.id,
            version=foo_launch_template.latest_version,
        ),
    )],
    opts=pulumi.ResourceOptions(depends_on=[aws_iam_policy_attachment["test-attach"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var fooLaunchTemplate = new Aws.Ec2.LaunchTemplate("fooLaunchTemplate", new()
    {
        ImageId = "ami-516b9131",
        InstanceType = "m1.small",
        KeyName = "some-key",
    });

    var fooSpotFleetRequest = new Aws.Ec2.SpotFleetRequest("fooSpotFleetRequest", new()
    {
        IamFleetRole = "arn:aws:iam::12345678:role/spot-fleet",
        SpotPrice = "0.005",
        TargetCapacity = 2,
        ValidUntil = "2019-11-04T20:44:20Z",
        LaunchTemplateConfigs = new[]
        {
            new Aws.Ec2.Inputs.SpotFleetRequestLaunchTemplateConfigArgs
            {
                LaunchTemplateSpecification = new Aws.Ec2.Inputs.SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs
                {
                    Id = fooLaunchTemplate.Id,
                    Version = fooLaunchTemplate.LatestVersion,
                },
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_policy_attachment.Test_attach,
        },
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
		fooLaunchTemplate, err := ec2.NewLaunchTemplate(ctx, "fooLaunchTemplate", &ec2.LaunchTemplateArgs{
			ImageId:      pulumi.String("ami-516b9131"),
			InstanceType: pulumi.String("m1.small"),
			KeyName:      pulumi.String("some-key"),
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewSpotFleetRequest(ctx, "fooSpotFleetRequest", &ec2.SpotFleetRequestArgs{
			IamFleetRole:   pulumi.String("arn:aws:iam::12345678:role/spot-fleet"),
			SpotPrice:      pulumi.String("0.005"),
			TargetCapacity: pulumi.Int(2),
			ValidUntil:     pulumi.String("2019-11-04T20:44:20Z"),
			LaunchTemplateConfigs: ec2.SpotFleetRequestLaunchTemplateConfigArray{
				&ec2.SpotFleetRequestLaunchTemplateConfigArgs{
					LaunchTemplateSpecification: &ec2.SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs{
						Id:      fooLaunchTemplate.ID(),
						Version: fooLaunchTemplate.LatestVersion,
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			aws_iam_policy_attachment.TestAttach,
		}))
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
import com.pulumi.aws.ec2.LaunchTemplate;
import com.pulumi.aws.ec2.LaunchTemplateArgs;
import com.pulumi.aws.ec2.SpotFleetRequest;
import com.pulumi.aws.ec2.SpotFleetRequestArgs;
import com.pulumi.aws.ec2.inputs.SpotFleetRequestLaunchTemplateConfigArgs;
import com.pulumi.aws.ec2.inputs.SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs;
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
        var fooLaunchTemplate = new LaunchTemplate("fooLaunchTemplate", LaunchTemplateArgs.builder()        
            .imageId("ami-516b9131")
            .instanceType("m1.small")
            .keyName("some-key")
            .build());

        var fooSpotFleetRequest = new SpotFleetRequest("fooSpotFleetRequest", SpotFleetRequestArgs.builder()        
            .iamFleetRole("arn:aws:iam::12345678:role/spot-fleet")
            .spotPrice("0.005")
            .targetCapacity(2)
            .validUntil("2019-11-04T20:44:20Z")
            .launchTemplateConfigs(SpotFleetRequestLaunchTemplateConfigArgs.builder()
                .launchTemplateSpecification(SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs.builder()
                    .id(fooLaunchTemplate.id())
                    .version(fooLaunchTemplate.latestVersion())
                    .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_iam_policy_attachment.test-attach())
                .build());

    }
}
```
```yaml
resources:
  fooLaunchTemplate:
    type: aws:ec2:LaunchTemplate
    properties:
      imageId: ami-516b9131
      instanceType: m1.small
      keyName: some-key
  fooSpotFleetRequest:
    type: aws:ec2:SpotFleetRequest
    properties:
      iamFleetRole: arn:aws:iam::12345678:role/spot-fleet
      spotPrice: '0.005'
      targetCapacity: 2
      validUntil: 2019-11-04T20:44:20Z
      launchTemplateConfigs:
        - launchTemplateSpecification:
            id: ${fooLaunchTemplate.id}
            version: ${fooLaunchTemplate.latestVersion}
    options:
      dependson:
        - ${aws_iam_policy_attachment"test-attach"[%!s(MISSING)]}
```

> **NOTE:** This provider does not support the functionality where multiple `subnet_id` or `availability_zone` parameters can be specified in the same
launch configuration block. If you want to specify multiple values, then separate launch configuration blocks should be used or launch template overrides should be configured, one per subnet:
{{% /example %}}
{{% example %}}
### Using multiple launch specifications

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ec2.SpotFleetRequest("foo", {
    iamFleetRole: "arn:aws:iam::12345678:role/spot-fleet",
    launchSpecifications: [
        {
            ami: "ami-d06a90b0",
            availabilityZone: "us-west-2a",
            instanceType: "m1.small",
            keyName: "my-key",
        },
        {
            ami: "ami-d06a90b0",
            availabilityZone: "us-west-2a",
            instanceType: "m5.large",
            keyName: "my-key",
        },
    ],
    spotPrice: "0.005",
    targetCapacity: 2,
    validUntil: "2019-11-04T20:44:20Z",
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ec2.SpotFleetRequest("foo",
    iam_fleet_role="arn:aws:iam::12345678:role/spot-fleet",
    launch_specifications=[
        aws.ec2.SpotFleetRequestLaunchSpecificationArgs(
            ami="ami-d06a90b0",
            availability_zone="us-west-2a",
            instance_type="m1.small",
            key_name="my-key",
        ),
        aws.ec2.SpotFleetRequestLaunchSpecificationArgs(
            ami="ami-d06a90b0",
            availability_zone="us-west-2a",
            instance_type="m5.large",
            key_name="my-key",
        ),
    ],
    spot_price="0.005",
    target_capacity=2,
    valid_until="2019-11-04T20:44:20Z")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ec2.SpotFleetRequest("foo", new()
    {
        IamFleetRole = "arn:aws:iam::12345678:role/spot-fleet",
        LaunchSpecifications = new[]
        {
            new Aws.Ec2.Inputs.SpotFleetRequestLaunchSpecificationArgs
            {
                Ami = "ami-d06a90b0",
                AvailabilityZone = "us-west-2a",
                InstanceType = "m1.small",
                KeyName = "my-key",
            },
            new Aws.Ec2.Inputs.SpotFleetRequestLaunchSpecificationArgs
            {
                Ami = "ami-d06a90b0",
                AvailabilityZone = "us-west-2a",
                InstanceType = "m5.large",
                KeyName = "my-key",
            },
        },
        SpotPrice = "0.005",
        TargetCapacity = 2,
        ValidUntil = "2019-11-04T20:44:20Z",
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
		_, err := ec2.NewSpotFleetRequest(ctx, "foo", &ec2.SpotFleetRequestArgs{
			IamFleetRole: pulumi.String("arn:aws:iam::12345678:role/spot-fleet"),
			LaunchSpecifications: ec2.SpotFleetRequestLaunchSpecificationArray{
				&ec2.SpotFleetRequestLaunchSpecificationArgs{
					Ami:              pulumi.String("ami-d06a90b0"),
					AvailabilityZone: pulumi.String("us-west-2a"),
					InstanceType:     pulumi.String("m1.small"),
					KeyName:          pulumi.String("my-key"),
				},
				&ec2.SpotFleetRequestLaunchSpecificationArgs{
					Ami:              pulumi.String("ami-d06a90b0"),
					AvailabilityZone: pulumi.String("us-west-2a"),
					InstanceType:     pulumi.String("m5.large"),
					KeyName:          pulumi.String("my-key"),
				},
			},
			SpotPrice:      pulumi.String("0.005"),
			TargetCapacity: pulumi.Int(2),
			ValidUntil:     pulumi.String("2019-11-04T20:44:20Z"),
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
import com.pulumi.aws.ec2.SpotFleetRequest;
import com.pulumi.aws.ec2.SpotFleetRequestArgs;
import com.pulumi.aws.ec2.inputs.SpotFleetRequestLaunchSpecificationArgs;
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
        var foo = new SpotFleetRequest("foo", SpotFleetRequestArgs.builder()        
            .iamFleetRole("arn:aws:iam::12345678:role/spot-fleet")
            .launchSpecifications(            
                SpotFleetRequestLaunchSpecificationArgs.builder()
                    .ami("ami-d06a90b0")
                    .availabilityZone("us-west-2a")
                    .instanceType("m1.small")
                    .keyName("my-key")
                    .build(),
                SpotFleetRequestLaunchSpecificationArgs.builder()
                    .ami("ami-d06a90b0")
                    .availabilityZone("us-west-2a")
                    .instanceType("m5.large")
                    .keyName("my-key")
                    .build())
            .spotPrice("0.005")
            .targetCapacity(2)
            .validUntil("2019-11-04T20:44:20Z")
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ec2:SpotFleetRequest
    properties:
      iamFleetRole: arn:aws:iam::12345678:role/spot-fleet
      launchSpecifications:
        - ami: ami-d06a90b0
          availabilityZone: us-west-2a
          instanceType: m1.small
          keyName: my-key
        - ami: ami-d06a90b0
          availabilityZone: us-west-2a
          instanceType: m5.large
          keyName: my-key
      spotPrice: '0.005'
      targetCapacity: 2
      validUntil: 2019-11-04T20:44:20Z
```

> In this example, we use a `dynamic` block to define zero or more `launch_specification` blocks, producing one for each element in the list of subnet ids.

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.ec2.SpotFleetRequest;
import com.pulumi.aws.ec2.SpotFleetRequestArgs;
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
        final var config = ctx.config();
        final var subnets = config.get("subnets");
        var example = new SpotFleetRequest("example", SpotFleetRequestArgs.builder()        
            .iamFleetRole("arn:aws:iam::12345678:role/spot-fleet")
            .targetCapacity(3)
            .validUntil("2019-11-04T20:44:20Z")
            .allocationStrategy("lowestPrice")
            .fleetType("request")
            .waitForFulfillment("true")
            .terminateInstancesWithExpiration("true")
            .dynamic(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Using multiple launch configurations

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.ec2.getSubnets({
    filters: [{
        name: "vpc-id",
        values: [_var.vpc_id],
    }],
});
const fooLaunchTemplate = new aws.ec2.LaunchTemplate("fooLaunchTemplate", {
    imageId: "ami-516b9131",
    instanceType: "m1.small",
    keyName: "some-key",
});
const fooSpotFleetRequest = new aws.ec2.SpotFleetRequest("fooSpotFleetRequest", {
    iamFleetRole: "arn:aws:iam::12345678:role/spot-fleet",
    spotPrice: "0.005",
    targetCapacity: 2,
    validUntil: "2019-11-04T20:44:20Z",
    launchTemplateConfigs: [{
        launchTemplateSpecification: {
            id: fooLaunchTemplate.id,
            version: fooLaunchTemplate.latestVersion,
        },
        overrides: [
            {
                subnetId: example.then(example => example.ids?.[0]),
            },
            {
                subnetId: example.then(example => example.ids?.[1]),
            },
            {
                subnetId: example.then(example => example.ids?.[2]),
            },
        ],
    }],
}, {
    dependsOn: [aws_iam_policy_attachment["test-attach"]],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.ec2.get_subnets(filters=[aws.ec2.GetSubnetsFilterArgs(
    name="vpc-id",
    values=[var["vpc_id"]],
)])
foo_launch_template = aws.ec2.LaunchTemplate("fooLaunchTemplate",
    image_id="ami-516b9131",
    instance_type="m1.small",
    key_name="some-key")
foo_spot_fleet_request = aws.ec2.SpotFleetRequest("fooSpotFleetRequest",
    iam_fleet_role="arn:aws:iam::12345678:role/spot-fleet",
    spot_price="0.005",
    target_capacity=2,
    valid_until="2019-11-04T20:44:20Z",
    launch_template_configs=[aws.ec2.SpotFleetRequestLaunchTemplateConfigArgs(
        launch_template_specification=aws.ec2.SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs(
            id=foo_launch_template.id,
            version=foo_launch_template.latest_version,
        ),
        overrides=[
            aws.ec2.SpotFleetRequestLaunchTemplateConfigOverrideArgs(
                subnet_id=example.ids[0],
            ),
            aws.ec2.SpotFleetRequestLaunchTemplateConfigOverrideArgs(
                subnet_id=example.ids[1],
            ),
            aws.ec2.SpotFleetRequestLaunchTemplateConfigOverrideArgs(
                subnet_id=example.ids[2],
            ),
        ],
    )],
    opts=pulumi.ResourceOptions(depends_on=[aws_iam_policy_attachment["test-attach"]]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Ec2.GetSubnets.Invoke(new()
    {
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetSubnetsFilterInputArgs
            {
                Name = "vpc-id",
                Values = new[]
                {
                    @var.Vpc_id,
                },
            },
        },
    });

    var fooLaunchTemplate = new Aws.Ec2.LaunchTemplate("fooLaunchTemplate", new()
    {
        ImageId = "ami-516b9131",
        InstanceType = "m1.small",
        KeyName = "some-key",
    });

    var fooSpotFleetRequest = new Aws.Ec2.SpotFleetRequest("fooSpotFleetRequest", new()
    {
        IamFleetRole = "arn:aws:iam::12345678:role/spot-fleet",
        SpotPrice = "0.005",
        TargetCapacity = 2,
        ValidUntil = "2019-11-04T20:44:20Z",
        LaunchTemplateConfigs = new[]
        {
            new Aws.Ec2.Inputs.SpotFleetRequestLaunchTemplateConfigArgs
            {
                LaunchTemplateSpecification = new Aws.Ec2.Inputs.SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs
                {
                    Id = fooLaunchTemplate.Id,
                    Version = fooLaunchTemplate.LatestVersion,
                },
                Overrides = new[]
                {
                    new Aws.Ec2.Inputs.SpotFleetRequestLaunchTemplateConfigOverrideArgs
                    {
                        SubnetId = example.Apply(getSubnetsResult => getSubnetsResult.Ids[0]),
                    },
                    new Aws.Ec2.Inputs.SpotFleetRequestLaunchTemplateConfigOverrideArgs
                    {
                        SubnetId = example.Apply(getSubnetsResult => getSubnetsResult.Ids[1]),
                    },
                    new Aws.Ec2.Inputs.SpotFleetRequestLaunchTemplateConfigOverrideArgs
                    {
                        SubnetId = example.Apply(getSubnetsResult => getSubnetsResult.Ids[2]),
                    },
                },
            },
        },
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            aws_iam_policy_attachment.Test_attach,
        },
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
example, err := ec2.GetSubnets(ctx, &ec2.GetSubnetsArgs{
Filters: []ec2.GetSubnetsFilter{
{
Name: "vpc-id",
Values: interface{}{
_var.Vpc_id,
},
},
},
}, nil);
if err != nil {
return err
}
fooLaunchTemplate, err := ec2.NewLaunchTemplate(ctx, "fooLaunchTemplate", &ec2.LaunchTemplateArgs{
ImageId: pulumi.String("ami-516b9131"),
InstanceType: pulumi.String("m1.small"),
KeyName: pulumi.String("some-key"),
})
if err != nil {
return err
}
_, err = ec2.NewSpotFleetRequest(ctx, "fooSpotFleetRequest", &ec2.SpotFleetRequestArgs{
IamFleetRole: pulumi.String("arn:aws:iam::12345678:role/spot-fleet"),
SpotPrice: pulumi.String("0.005"),
TargetCapacity: pulumi.Int(2),
ValidUntil: pulumi.String("2019-11-04T20:44:20Z"),
LaunchTemplateConfigs: ec2.SpotFleetRequestLaunchTemplateConfigArray{
&ec2.SpotFleetRequestLaunchTemplateConfigArgs{
LaunchTemplateSpecification: &ec2.SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs{
Id: fooLaunchTemplate.ID(),
Version: fooLaunchTemplate.LatestVersion,
},
Overrides: ec2.SpotFleetRequestLaunchTemplateConfigOverrideArray{
&ec2.SpotFleetRequestLaunchTemplateConfigOverrideArgs{
SubnetId: *pulumi.String(example.Ids[0]),
},
&ec2.SpotFleetRequestLaunchTemplateConfigOverrideArgs{
SubnetId: *pulumi.String(example.Ids[1]),
},
&ec2.SpotFleetRequestLaunchTemplateConfigOverrideArgs{
SubnetId: *pulumi.String(example.Ids[2]),
},
},
},
},
}, pulumi.DependsOn([]pulumi.Resource{
aws_iam_policy_attachment.TestAttach,
}))
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
import com.pulumi.aws.ec2.inputs.GetSubnetsArgs;
import com.pulumi.aws.ec2.LaunchTemplate;
import com.pulumi.aws.ec2.LaunchTemplateArgs;
import com.pulumi.aws.ec2.SpotFleetRequest;
import com.pulumi.aws.ec2.SpotFleetRequestArgs;
import com.pulumi.aws.ec2.inputs.SpotFleetRequestLaunchTemplateConfigArgs;
import com.pulumi.aws.ec2.inputs.SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs;
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
        final var example = Ec2Functions.getSubnets(GetSubnetsArgs.builder()
            .filters(GetSubnetsFilterArgs.builder()
                .name("vpc-id")
                .values(var_.vpc_id())
                .build())
            .build());

        var fooLaunchTemplate = new LaunchTemplate("fooLaunchTemplate", LaunchTemplateArgs.builder()        
            .imageId("ami-516b9131")
            .instanceType("m1.small")
            .keyName("some-key")
            .build());

        var fooSpotFleetRequest = new SpotFleetRequest("fooSpotFleetRequest", SpotFleetRequestArgs.builder()        
            .iamFleetRole("arn:aws:iam::12345678:role/spot-fleet")
            .spotPrice("0.005")
            .targetCapacity(2)
            .validUntil("2019-11-04T20:44:20Z")
            .launchTemplateConfigs(SpotFleetRequestLaunchTemplateConfigArgs.builder()
                .launchTemplateSpecification(SpotFleetRequestLaunchTemplateConfigLaunchTemplateSpecificationArgs.builder()
                    .id(fooLaunchTemplate.id())
                    .version(fooLaunchTemplate.latestVersion())
                    .build())
                .overrides(                
                    SpotFleetRequestLaunchTemplateConfigOverrideArgs.builder()
                        .subnetId(example.applyValue(getSubnetsResult -> getSubnetsResult.ids()[0]))
                        .build(),
                    SpotFleetRequestLaunchTemplateConfigOverrideArgs.builder()
                        .subnetId(example.applyValue(getSubnetsResult -> getSubnetsResult.ids()[1]))
                        .build(),
                    SpotFleetRequestLaunchTemplateConfigOverrideArgs.builder()
                        .subnetId(example.applyValue(getSubnetsResult -> getSubnetsResult.ids()[2]))
                        .build())
                .build())
            .build(), CustomResourceOptions.builder()
                .dependsOn(aws_iam_policy_attachment.test-attach())
                .build());

    }
}
```
```yaml
resources:
  fooLaunchTemplate:
    type: aws:ec2:LaunchTemplate
    properties:
      imageId: ami-516b9131
      instanceType: m1.small
      keyName: some-key
  fooSpotFleetRequest:
    type: aws:ec2:SpotFleetRequest
    properties:
      iamFleetRole: arn:aws:iam::12345678:role/spot-fleet
      spotPrice: '0.005'
      targetCapacity: 2
      validUntil: 2019-11-04T20:44:20Z
      launchTemplateConfigs:
        - launchTemplateSpecification:
            id: ${fooLaunchTemplate.id}
            version: ${fooLaunchTemplate.latestVersion}
          overrides:
            - subnetId: ${example.ids[0]}
            - subnetId: ${example.ids[1]}
            - subnetId: ${example.ids[2]}
    options:
      dependson:
        - ${aws_iam_policy_attachment"test-attach"[%!s(MISSING)]}
variables:
  example:
    fn::invoke:
      Function: aws:ec2:getSubnets
      Arguments:
        filters:
          - name: vpc-id
            values:
              - ${var.vpc_id}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Spot Fleet Requests using `id`. For example:

```sh
 $ pulumi import aws:ec2/spotFleetRequest:SpotFleetRequest fleet sfr-005e9ec8-5546-4c31-b317-31a62325411e
```
 