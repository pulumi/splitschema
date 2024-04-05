Provides an EC2 launch template resource. Can be used to create instances or auto scaling groups.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const foo = new aws.ec2.LaunchTemplate("foo", {
    blockDeviceMappings: [{
        deviceName: "/dev/sdf",
        ebs: {
            volumeSize: 20,
        },
    }],
    capacityReservationSpecification: {
        capacityReservationPreference: "open",
    },
    cpuOptions: {
        coreCount: 4,
        threadsPerCore: 2,
    },
    creditSpecification: {
        cpuCredits: "standard",
    },
    disableApiStop: true,
    disableApiTermination: true,
    ebsOptimized: "true",
    elasticGpuSpecifications: [{
        type: "test",
    }],
    elasticInferenceAccelerator: {
        type: "eia1.medium",
    },
    iamInstanceProfile: {
        name: "test",
    },
    imageId: "ami-test",
    instanceInitiatedShutdownBehavior: "terminate",
    instanceMarketOptions: {
        marketType: "spot",
    },
    instanceType: "t2.micro",
    kernelId: "test",
    keyName: "test",
    licenseSpecifications: [{
        licenseConfigurationArn: "arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef",
    }],
    metadataOptions: {
        httpEndpoint: "enabled",
        httpTokens: "required",
        httpPutResponseHopLimit: 1,
        instanceMetadataTags: "enabled",
    },
    monitoring: {
        enabled: true,
    },
    networkInterfaces: [{
        associatePublicIpAddress: "true",
    }],
    placement: {
        availabilityZone: "us-west-2a",
    },
    ramDiskId: "test",
    vpcSecurityGroupIds: ["sg-12345678"],
    tagSpecifications: [{
        resourceType: "instance",
        tags: {
            Name: "test",
        },
    }],
    userData: Buffer.from(fs.readFileSync(`${path.module}/example.sh`, 'binary')).toString('base64'),
});
```
```python
import pulumi
import base64
import pulumi_aws as aws

foo = aws.ec2.LaunchTemplate("foo",
    block_device_mappings=[aws.ec2.LaunchTemplateBlockDeviceMappingArgs(
        device_name="/dev/sdf",
        ebs=aws.ec2.LaunchTemplateBlockDeviceMappingEbsArgs(
            volume_size=20,
        ),
    )],
    capacity_reservation_specification=aws.ec2.LaunchTemplateCapacityReservationSpecificationArgs(
        capacity_reservation_preference="open",
    ),
    cpu_options=aws.ec2.LaunchTemplateCpuOptionsArgs(
        core_count=4,
        threads_per_core=2,
    ),
    credit_specification=aws.ec2.LaunchTemplateCreditSpecificationArgs(
        cpu_credits="standard",
    ),
    disable_api_stop=True,
    disable_api_termination=True,
    ebs_optimized="true",
    elastic_gpu_specifications=[aws.ec2.LaunchTemplateElasticGpuSpecificationArgs(
        type="test",
    )],
    elastic_inference_accelerator=aws.ec2.LaunchTemplateElasticInferenceAcceleratorArgs(
        type="eia1.medium",
    ),
    iam_instance_profile=aws.ec2.LaunchTemplateIamInstanceProfileArgs(
        name="test",
    ),
    image_id="ami-test",
    instance_initiated_shutdown_behavior="terminate",
    instance_market_options=aws.ec2.LaunchTemplateInstanceMarketOptionsArgs(
        market_type="spot",
    ),
    instance_type="t2.micro",
    kernel_id="test",
    key_name="test",
    license_specifications=[aws.ec2.LaunchTemplateLicenseSpecificationArgs(
        license_configuration_arn="arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef",
    )],
    metadata_options=aws.ec2.LaunchTemplateMetadataOptionsArgs(
        http_endpoint="enabled",
        http_tokens="required",
        http_put_response_hop_limit=1,
        instance_metadata_tags="enabled",
    ),
    monitoring=aws.ec2.LaunchTemplateMonitoringArgs(
        enabled=True,
    ),
    network_interfaces=[aws.ec2.LaunchTemplateNetworkInterfaceArgs(
        associate_public_ip_address="true",
    )],
    placement=aws.ec2.LaunchTemplatePlacementArgs(
        availability_zone="us-west-2a",
    ),
    ram_disk_id="test",
    vpc_security_group_ids=["sg-12345678"],
    tag_specifications=[aws.ec2.LaunchTemplateTagSpecificationArgs(
        resource_type="instance",
        tags={
            "Name": "test",
        },
    )],
    user_data=(lambda path: base64.b64encode(open(path).read().encode()).decode())(f"{path['module']}/example.sh"))
```
```csharp
using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

	
string ReadFileBase64(string path) 
{
    return Convert.ToBase64String(Encoding.UTF8.GetBytes(File.ReadAllText(path)));
}

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ec2.LaunchTemplate("foo", new()
    {
        BlockDeviceMappings = new[]
        {
            new Aws.Ec2.Inputs.LaunchTemplateBlockDeviceMappingArgs
            {
                DeviceName = "/dev/sdf",
                Ebs = new Aws.Ec2.Inputs.LaunchTemplateBlockDeviceMappingEbsArgs
                {
                    VolumeSize = 20,
                },
            },
        },
        CapacityReservationSpecification = new Aws.Ec2.Inputs.LaunchTemplateCapacityReservationSpecificationArgs
        {
            CapacityReservationPreference = "open",
        },
        CpuOptions = new Aws.Ec2.Inputs.LaunchTemplateCpuOptionsArgs
        {
            CoreCount = 4,
            ThreadsPerCore = 2,
        },
        CreditSpecification = new Aws.Ec2.Inputs.LaunchTemplateCreditSpecificationArgs
        {
            CpuCredits = "standard",
        },
        DisableApiStop = true,
        DisableApiTermination = true,
        EbsOptimized = "true",
        ElasticGpuSpecifications = new[]
        {
            new Aws.Ec2.Inputs.LaunchTemplateElasticGpuSpecificationArgs
            {
                Type = "test",
            },
        },
        ElasticInferenceAccelerator = new Aws.Ec2.Inputs.LaunchTemplateElasticInferenceAcceleratorArgs
        {
            Type = "eia1.medium",
        },
        IamInstanceProfile = new Aws.Ec2.Inputs.LaunchTemplateIamInstanceProfileArgs
        {
            Name = "test",
        },
        ImageId = "ami-test",
        InstanceInitiatedShutdownBehavior = "terminate",
        InstanceMarketOptions = new Aws.Ec2.Inputs.LaunchTemplateInstanceMarketOptionsArgs
        {
            MarketType = "spot",
        },
        InstanceType = "t2.micro",
        KernelId = "test",
        KeyName = "test",
        LicenseSpecifications = new[]
        {
            new Aws.Ec2.Inputs.LaunchTemplateLicenseSpecificationArgs
            {
                LicenseConfigurationArn = "arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef",
            },
        },
        MetadataOptions = new Aws.Ec2.Inputs.LaunchTemplateMetadataOptionsArgs
        {
            HttpEndpoint = "enabled",
            HttpTokens = "required",
            HttpPutResponseHopLimit = 1,
            InstanceMetadataTags = "enabled",
        },
        Monitoring = new Aws.Ec2.Inputs.LaunchTemplateMonitoringArgs
        {
            Enabled = true,
        },
        NetworkInterfaces = new[]
        {
            new Aws.Ec2.Inputs.LaunchTemplateNetworkInterfaceArgs
            {
                AssociatePublicIpAddress = "true",
            },
        },
        Placement = new Aws.Ec2.Inputs.LaunchTemplatePlacementArgs
        {
            AvailabilityZone = "us-west-2a",
        },
        RamDiskId = "test",
        VpcSecurityGroupIds = new[]
        {
            "sg-12345678",
        },
        TagSpecifications = new[]
        {
            new Aws.Ec2.Inputs.LaunchTemplateTagSpecificationArgs
            {
                ResourceType = "instance",
                Tags = 
                {
                    { "Name", "test" },
                },
            },
        },
        UserData = ReadFileBase64($"{path.Module}/example.sh"),
    });

});
```
```go
package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func filebase64OrPanic(path string) string {
	if fileData, err := os.ReadFile(path); err == nil {
		return base64.StdEncoding.EncodeToString(fileData[:])
	} else {
		panic(err.Error())
	}
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ec2.NewLaunchTemplate(ctx, "foo", &ec2.LaunchTemplateArgs{
			BlockDeviceMappings: ec2.LaunchTemplateBlockDeviceMappingArray{
				&ec2.LaunchTemplateBlockDeviceMappingArgs{
					DeviceName: pulumi.String("/dev/sdf"),
					Ebs: &ec2.LaunchTemplateBlockDeviceMappingEbsArgs{
						VolumeSize: pulumi.Int(20),
					},
				},
			},
			CapacityReservationSpecification: &ec2.LaunchTemplateCapacityReservationSpecificationArgs{
				CapacityReservationPreference: pulumi.String("open"),
			},
			CpuOptions: &ec2.LaunchTemplateCpuOptionsArgs{
				CoreCount:      pulumi.Int(4),
				ThreadsPerCore: pulumi.Int(2),
			},
			CreditSpecification: &ec2.LaunchTemplateCreditSpecificationArgs{
				CpuCredits: pulumi.String("standard"),
			},
			DisableApiStop:        pulumi.Bool(true),
			DisableApiTermination: pulumi.Bool(true),
			EbsOptimized:          pulumi.String("true"),
			ElasticGpuSpecifications: ec2.LaunchTemplateElasticGpuSpecificationArray{
				&ec2.LaunchTemplateElasticGpuSpecificationArgs{
					Type: pulumi.String("test"),
				},
			},
			ElasticInferenceAccelerator: &ec2.LaunchTemplateElasticInferenceAcceleratorArgs{
				Type: pulumi.String("eia1.medium"),
			},
			IamInstanceProfile: &ec2.LaunchTemplateIamInstanceProfileArgs{
				Name: pulumi.String("test"),
			},
			ImageId:                           pulumi.String("ami-test"),
			InstanceInitiatedShutdownBehavior: pulumi.String("terminate"),
			InstanceMarketOptions: &ec2.LaunchTemplateInstanceMarketOptionsArgs{
				MarketType: pulumi.String("spot"),
			},
			InstanceType: pulumi.String("t2.micro"),
			KernelId:     pulumi.String("test"),
			KeyName:      pulumi.String("test"),
			LicenseSpecifications: ec2.LaunchTemplateLicenseSpecificationArray{
				&ec2.LaunchTemplateLicenseSpecificationArgs{
					LicenseConfigurationArn: pulumi.String("arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef"),
				},
			},
			MetadataOptions: &ec2.LaunchTemplateMetadataOptionsArgs{
				HttpEndpoint:            pulumi.String("enabled"),
				HttpTokens:              pulumi.String("required"),
				HttpPutResponseHopLimit: pulumi.Int(1),
				InstanceMetadataTags:    pulumi.String("enabled"),
			},
			Monitoring: &ec2.LaunchTemplateMonitoringArgs{
				Enabled: pulumi.Bool(true),
			},
			NetworkInterfaces: ec2.LaunchTemplateNetworkInterfaceArray{
				&ec2.LaunchTemplateNetworkInterfaceArgs{
					AssociatePublicIpAddress: pulumi.String("true"),
				},
			},
			Placement: &ec2.LaunchTemplatePlacementArgs{
				AvailabilityZone: pulumi.String("us-west-2a"),
			},
			RamDiskId: pulumi.String("test"),
			VpcSecurityGroupIds: pulumi.StringArray{
				pulumi.String("sg-12345678"),
			},
			TagSpecifications: ec2.LaunchTemplateTagSpecificationArray{
				&ec2.LaunchTemplateTagSpecificationArgs{
					ResourceType: pulumi.String("instance"),
					Tags: pulumi.StringMap{
						"Name": pulumi.String("test"),
					},
				},
			},
			UserData: filebase64OrPanic(fmt.Sprintf("%v/example.sh", path.Module)),
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
import com.pulumi.aws.ec2.LaunchTemplate;
import com.pulumi.aws.ec2.LaunchTemplateArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateBlockDeviceMappingArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateBlockDeviceMappingEbsArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateCapacityReservationSpecificationArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateCpuOptionsArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateCreditSpecificationArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateElasticGpuSpecificationArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateElasticInferenceAcceleratorArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateIamInstanceProfileArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateInstanceMarketOptionsArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateLicenseSpecificationArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateMetadataOptionsArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateMonitoringArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateNetworkInterfaceArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplatePlacementArgs;
import com.pulumi.aws.ec2.inputs.LaunchTemplateTagSpecificationArgs;
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
        var foo = new LaunchTemplate("foo", LaunchTemplateArgs.builder()        
            .blockDeviceMappings(LaunchTemplateBlockDeviceMappingArgs.builder()
                .deviceName("/dev/sdf")
                .ebs(LaunchTemplateBlockDeviceMappingEbsArgs.builder()
                    .volumeSize(20)
                    .build())
                .build())
            .capacityReservationSpecification(LaunchTemplateCapacityReservationSpecificationArgs.builder()
                .capacityReservationPreference("open")
                .build())
            .cpuOptions(LaunchTemplateCpuOptionsArgs.builder()
                .coreCount(4)
                .threadsPerCore(2)
                .build())
            .creditSpecification(LaunchTemplateCreditSpecificationArgs.builder()
                .cpuCredits("standard")
                .build())
            .disableApiStop(true)
            .disableApiTermination(true)
            .ebsOptimized(true)
            .elasticGpuSpecifications(LaunchTemplateElasticGpuSpecificationArgs.builder()
                .type("test")
                .build())
            .elasticInferenceAccelerator(LaunchTemplateElasticInferenceAcceleratorArgs.builder()
                .type("eia1.medium")
                .build())
            .iamInstanceProfile(LaunchTemplateIamInstanceProfileArgs.builder()
                .name("test")
                .build())
            .imageId("ami-test")
            .instanceInitiatedShutdownBehavior("terminate")
            .instanceMarketOptions(LaunchTemplateInstanceMarketOptionsArgs.builder()
                .marketType("spot")
                .build())
            .instanceType("t2.micro")
            .kernelId("test")
            .keyName("test")
            .licenseSpecifications(LaunchTemplateLicenseSpecificationArgs.builder()
                .licenseConfigurationArn("arn:aws:license-manager:eu-west-1:123456789012:license-configuration:lic-0123456789abcdef0123456789abcdef")
                .build())
            .metadataOptions(LaunchTemplateMetadataOptionsArgs.builder()
                .httpEndpoint("enabled")
                .httpTokens("required")
                .httpPutResponseHopLimit(1)
                .instanceMetadataTags("enabled")
                .build())
            .monitoring(LaunchTemplateMonitoringArgs.builder()
                .enabled(true)
                .build())
            .networkInterfaces(LaunchTemplateNetworkInterfaceArgs.builder()
                .associatePublicIpAddress(true)
                .build())
            .placement(LaunchTemplatePlacementArgs.builder()
                .availabilityZone("us-west-2a")
                .build())
            .ramDiskId("test")
            .vpcSecurityGroupIds("sg-12345678")
            .tagSpecifications(LaunchTemplateTagSpecificationArgs.builder()
                .resourceType("instance")
                .tags(Map.of("Name", "test"))
                .build())
            .userData(Base64.getEncoder().encodeToString(Files.readAllBytes(Paths.get(String.format("%s/example.sh", path.module())))))
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Launch Templates using the `id`. For example:

```sh
 $ pulumi import aws:ec2/launchTemplate:LaunchTemplate web lt-12345678
```
 