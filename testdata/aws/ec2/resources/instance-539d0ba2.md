Provides an EC2 instance resource. This allows instances to be created, updated, and deleted.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic example using AMI lookup

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const ubuntu = aws.ec2.getAmi({
    mostRecent: true,
    filters: [
        {
            name: "name",
            values: ["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"],
        },
        {
            name: "virtualization-type",
            values: ["hvm"],
        },
    ],
    owners: ["099720109477"],
});
const web = new aws.ec2.Instance("web", {
    ami: ubuntu.then(ubuntu => ubuntu.id),
    instanceType: "t3.micro",
    tags: {
        Name: "HelloWorld",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

ubuntu = aws.ec2.get_ami(most_recent=True,
    filters=[
        aws.ec2.GetAmiFilterArgs(
            name="name",
            values=["ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"],
        ),
        aws.ec2.GetAmiFilterArgs(
            name="virtualization-type",
            values=["hvm"],
        ),
    ],
    owners=["099720109477"])
web = aws.ec2.Instance("web",
    ami=ubuntu.id,
    instance_type="t3.micro",
    tags={
        "Name": "HelloWorld",
    })
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
                    "ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*",
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

    var web = new Aws.Ec2.Instance("web", new()
    {
        Ami = ubuntu.Apply(getAmiResult => getAmiResult.Id),
        InstanceType = "t3.micro",
        Tags = 
        {
            { "Name", "HelloWorld" },
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
		ubuntu, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*",
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
		_, err = ec2.NewInstance(ctx, "web", &ec2.InstanceArgs{
			Ami:          *pulumi.String(ubuntu.Id),
			InstanceType: pulumi.String("t3.micro"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("HelloWorld"),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
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
                    .values("ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*")
                    .build(),
                GetAmiFilterArgs.builder()
                    .name("virtualization-type")
                    .values("hvm")
                    .build())
            .owners("099720109477")
            .build());

        var web = new Instance("web", InstanceArgs.builder()        
            .ami(ubuntu.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceType("t3.micro")
            .tags(Map.of("Name", "HelloWorld"))
            .build());

    }
}
```
```yaml
resources:
  web:
    type: aws:ec2:Instance
    properties:
      ami: ${ubuntu.id}
      instanceType: t3.micro
      tags:
        Name: HelloWorld
variables:
  ubuntu:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        mostRecent: true
        filters:
          - name: name
            values:
              - ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*
          - name: virtualization-type
            values:
              - hvm
        owners:
          - '099720109477'
```
{{% /example %}}
{{% example %}}
### Spot instance example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const thisAmi = aws.ec2.getAmi({
    mostRecent: true,
    owners: ["amazon"],
    filters: [
        {
            name: "architecture",
            values: ["arm64"],
        },
        {
            name: "name",
            values: ["al2023-ami-2023*"],
        },
    ],
});
const thisInstance = new aws.ec2.Instance("thisInstance", {
    ami: thisAmi.then(thisAmi => thisAmi.id),
    instanceMarketOptions: {
        spotOptions: {
            maxPrice: "0.0031",
        },
    },
    instanceType: "t4g.nano",
    tags: {
        Name: "test-spot",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

this_ami = aws.ec2.get_ami(most_recent=True,
    owners=["amazon"],
    filters=[
        aws.ec2.GetAmiFilterArgs(
            name="architecture",
            values=["arm64"],
        ),
        aws.ec2.GetAmiFilterArgs(
            name="name",
            values=["al2023-ami-2023*"],
        ),
    ])
this_instance = aws.ec2.Instance("thisInstance",
    ami=this_ami.id,
    instance_market_options=aws.ec2.InstanceInstanceMarketOptionsArgs(
        spot_options=aws.ec2.InstanceInstanceMarketOptionsSpotOptionsArgs(
            max_price="0.0031",
        ),
    ),
    instance_type="t4g.nano",
    tags={
        "Name": "test-spot",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var thisAmi = Aws.Ec2.GetAmi.Invoke(new()
    {
        MostRecent = true,
        Owners = new[]
        {
            "amazon",
        },
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "architecture",
                Values = new[]
                {
                    "arm64",
                },
            },
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "al2023-ami-2023*",
                },
            },
        },
    });

    var thisInstance = new Aws.Ec2.Instance("thisInstance", new()
    {
        Ami = thisAmi.Apply(getAmiResult => getAmiResult.Id),
        InstanceMarketOptions = new Aws.Ec2.Inputs.InstanceInstanceMarketOptionsArgs
        {
            SpotOptions = new Aws.Ec2.Inputs.InstanceInstanceMarketOptionsSpotOptionsArgs
            {
                MaxPrice = "0.0031",
            },
        },
        InstanceType = "t4g.nano",
        Tags = 
        {
            { "Name", "test-spot" },
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
		thisAmi, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Owners: []string{
				"amazon",
			},
			Filters: []ec2.GetAmiFilter{
				{
					Name: "architecture",
					Values: []string{
						"arm64",
					},
				},
				{
					Name: "name",
					Values: []string{
						"al2023-ami-2023*",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewInstance(ctx, "thisInstance", &ec2.InstanceArgs{
			Ami: *pulumi.String(thisAmi.Id),
			InstanceMarketOptions: &ec2.InstanceInstanceMarketOptionsArgs{
				SpotOptions: &ec2.InstanceInstanceMarketOptionsSpotOptionsArgs{
					MaxPrice: pulumi.String("0.0031"),
				},
			},
			InstanceType: pulumi.String("t4g.nano"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("test-spot"),
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
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import com.pulumi.aws.ec2.inputs.InstanceInstanceMarketOptionsArgs;
import com.pulumi.aws.ec2.inputs.InstanceInstanceMarketOptionsSpotOptionsArgs;
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
        final var thisAmi = Ec2Functions.getAmi(GetAmiArgs.builder()
            .mostRecent(true)
            .owners("amazon")
            .filters(            
                GetAmiFilterArgs.builder()
                    .name("architecture")
                    .values("arm64")
                    .build(),
                GetAmiFilterArgs.builder()
                    .name("name")
                    .values("al2023-ami-2023*")
                    .build())
            .build());

        var thisInstance = new Instance("thisInstance", InstanceArgs.builder()        
            .ami(thisAmi.applyValue(getAmiResult -> getAmiResult.id()))
            .instanceMarketOptions(InstanceInstanceMarketOptionsArgs.builder()
                .spotOptions(InstanceInstanceMarketOptionsSpotOptionsArgs.builder()
                    .maxPrice(0.0031)
                    .build())
                .build())
            .instanceType("t4g.nano")
            .tags(Map.of("Name", "test-spot"))
            .build());

    }
}
```
```yaml
resources:
  thisInstance:
    type: aws:ec2:Instance
    properties:
      ami: ${thisAmi.id}
      instanceMarketOptions:
        spotOptions:
          maxPrice: 0.0031
      instanceType: t4g.nano
      tags:
        Name: test-spot
variables:
  thisAmi:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        mostRecent: true
        owners:
          - amazon
        filters:
          - name: architecture
            values:
              - arm64
          - name: name
            values:
              - al2023-ami-2023*
```
{{% /example %}}
{{% example %}}
### Network and credit specification example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const myVpc = new aws.ec2.Vpc("myVpc", {
    cidrBlock: "172.16.0.0/16",
    tags: {
        Name: "tf-example",
    },
});
const mySubnet = new aws.ec2.Subnet("mySubnet", {
    vpcId: myVpc.id,
    cidrBlock: "172.16.10.0/24",
    availabilityZone: "us-west-2a",
    tags: {
        Name: "tf-example",
    },
});
const fooNetworkInterface = new aws.ec2.NetworkInterface("fooNetworkInterface", {
    subnetId: mySubnet.id,
    privateIps: ["172.16.10.100"],
    tags: {
        Name: "primary_network_interface",
    },
});
const fooInstance = new aws.ec2.Instance("fooInstance", {
    ami: "ami-005e54dee72cc1d00",
    instanceType: "t2.micro",
    networkInterfaces: [{
        networkInterfaceId: fooNetworkInterface.id,
        deviceIndex: 0,
    }],
    creditSpecification: {
        cpuCredits: "unlimited",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

my_vpc = aws.ec2.Vpc("myVpc",
    cidr_block="172.16.0.0/16",
    tags={
        "Name": "tf-example",
    })
my_subnet = aws.ec2.Subnet("mySubnet",
    vpc_id=my_vpc.id,
    cidr_block="172.16.10.0/24",
    availability_zone="us-west-2a",
    tags={
        "Name": "tf-example",
    })
foo_network_interface = aws.ec2.NetworkInterface("fooNetworkInterface",
    subnet_id=my_subnet.id,
    private_ips=["172.16.10.100"],
    tags={
        "Name": "primary_network_interface",
    })
foo_instance = aws.ec2.Instance("fooInstance",
    ami="ami-005e54dee72cc1d00",
    instance_type="t2.micro",
    network_interfaces=[aws.ec2.InstanceNetworkInterfaceArgs(
        network_interface_id=foo_network_interface.id,
        device_index=0,
    )],
    credit_specification=aws.ec2.InstanceCreditSpecificationArgs(
        cpu_credits="unlimited",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var myVpc = new Aws.Ec2.Vpc("myVpc", new()
    {
        CidrBlock = "172.16.0.0/16",
        Tags = 
        {
            { "Name", "tf-example" },
        },
    });

    var mySubnet = new Aws.Ec2.Subnet("mySubnet", new()
    {
        VpcId = myVpc.Id,
        CidrBlock = "172.16.10.0/24",
        AvailabilityZone = "us-west-2a",
        Tags = 
        {
            { "Name", "tf-example" },
        },
    });

    var fooNetworkInterface = new Aws.Ec2.NetworkInterface("fooNetworkInterface", new()
    {
        SubnetId = mySubnet.Id,
        PrivateIps = new[]
        {
            "172.16.10.100",
        },
        Tags = 
        {
            { "Name", "primary_network_interface" },
        },
    });

    var fooInstance = new Aws.Ec2.Instance("fooInstance", new()
    {
        Ami = "ami-005e54dee72cc1d00",
        InstanceType = "t2.micro",
        NetworkInterfaces = new[]
        {
            new Aws.Ec2.Inputs.InstanceNetworkInterfaceArgs
            {
                NetworkInterfaceId = fooNetworkInterface.Id,
                DeviceIndex = 0,
            },
        },
        CreditSpecification = new Aws.Ec2.Inputs.InstanceCreditSpecificationArgs
        {
            CpuCredits = "unlimited",
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
		myVpc, err := ec2.NewVpc(ctx, "myVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("172.16.0.0/16"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-example"),
			},
		})
		if err != nil {
			return err
		}
		mySubnet, err := ec2.NewSubnet(ctx, "mySubnet", &ec2.SubnetArgs{
			VpcId:            myVpc.ID(),
			CidrBlock:        pulumi.String("172.16.10.0/24"),
			AvailabilityZone: pulumi.String("us-west-2a"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-example"),
			},
		})
		if err != nil {
			return err
		}
		fooNetworkInterface, err := ec2.NewNetworkInterface(ctx, "fooNetworkInterface", &ec2.NetworkInterfaceArgs{
			SubnetId: mySubnet.ID(),
			PrivateIps: pulumi.StringArray{
				pulumi.String("172.16.10.100"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("primary_network_interface"),
			},
		})
		if err != nil {
			return err
		}
		_, err = ec2.NewInstance(ctx, "fooInstance", &ec2.InstanceArgs{
			Ami:          pulumi.String("ami-005e54dee72cc1d00"),
			InstanceType: pulumi.String("t2.micro"),
			NetworkInterfaces: ec2.InstanceNetworkInterfaceArray{
				&ec2.InstanceNetworkInterfaceArgs{
					NetworkInterfaceId: fooNetworkInterface.ID(),
					DeviceIndex:        pulumi.Int(0),
				},
			},
			CreditSpecification: &ec2.InstanceCreditSpecificationArgs{
				CpuCredits: pulumi.String("unlimited"),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.ec2.NetworkInterface;
import com.pulumi.aws.ec2.NetworkInterfaceArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import com.pulumi.aws.ec2.inputs.InstanceNetworkInterfaceArgs;
import com.pulumi.aws.ec2.inputs.InstanceCreditSpecificationArgs;
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
        var myVpc = new Vpc("myVpc", VpcArgs.builder()        
            .cidrBlock("172.16.0.0/16")
            .tags(Map.of("Name", "tf-example"))
            .build());

        var mySubnet = new Subnet("mySubnet", SubnetArgs.builder()        
            .vpcId(myVpc.id())
            .cidrBlock("172.16.10.0/24")
            .availabilityZone("us-west-2a")
            .tags(Map.of("Name", "tf-example"))
            .build());

        var fooNetworkInterface = new NetworkInterface("fooNetworkInterface", NetworkInterfaceArgs.builder()        
            .subnetId(mySubnet.id())
            .privateIps("172.16.10.100")
            .tags(Map.of("Name", "primary_network_interface"))
            .build());

        var fooInstance = new Instance("fooInstance", InstanceArgs.builder()        
            .ami("ami-005e54dee72cc1d00")
            .instanceType("t2.micro")
            .networkInterfaces(InstanceNetworkInterfaceArgs.builder()
                .networkInterfaceId(fooNetworkInterface.id())
                .deviceIndex(0)
                .build())
            .creditSpecification(InstanceCreditSpecificationArgs.builder()
                .cpuCredits("unlimited")
                .build())
            .build());

    }
}
```
```yaml
resources:
  myVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 172.16.0.0/16
      tags:
        Name: tf-example
  mySubnet:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${myVpc.id}
      cidrBlock: 172.16.10.0/24
      availabilityZone: us-west-2a
      tags:
        Name: tf-example
  fooNetworkInterface:
    type: aws:ec2:NetworkInterface
    properties:
      subnetId: ${mySubnet.id}
      privateIps:
        - 172.16.10.100
      tags:
        Name: primary_network_interface
  fooInstance:
    type: aws:ec2:Instance
    properties:
      ami: ami-005e54dee72cc1d00
      # us-west-2
      instanceType: t2.micro
      networkInterfaces:
        - networkInterfaceId: ${fooNetworkInterface.id}
          deviceIndex: 0
      creditSpecification:
        cpuCredits: unlimited
```
{{% /example %}}
{{% example %}}
### CPU options example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleVpc = new aws.ec2.Vpc("exampleVpc", {
    cidrBlock: "172.16.0.0/16",
    tags: {
        Name: "tf-example",
    },
});
const exampleSubnet = new aws.ec2.Subnet("exampleSubnet", {
    vpcId: exampleVpc.id,
    cidrBlock: "172.16.10.0/24",
    availabilityZone: "us-east-2a",
    tags: {
        Name: "tf-example",
    },
});
const amzn-linux-2023-ami = aws.ec2.getAmi({
    mostRecent: true,
    owners: ["amazon"],
    filters: [{
        name: "name",
        values: ["al2023-ami-2023.*-x86_64"],
    }],
});
const exampleInstance = new aws.ec2.Instance("exampleInstance", {
    ami: amzn_linux_2023_ami.then(amzn_linux_2023_ami => amzn_linux_2023_ami.id),
    instanceType: "c6a.2xlarge",
    subnetId: exampleSubnet.id,
    cpuOptions: {
        coreCount: 2,
        threadsPerCore: 2,
    },
    tags: {
        Name: "tf-example",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example_vpc = aws.ec2.Vpc("exampleVpc",
    cidr_block="172.16.0.0/16",
    tags={
        "Name": "tf-example",
    })
example_subnet = aws.ec2.Subnet("exampleSubnet",
    vpc_id=example_vpc.id,
    cidr_block="172.16.10.0/24",
    availability_zone="us-east-2a",
    tags={
        "Name": "tf-example",
    })
amzn_linux_2023_ami = aws.ec2.get_ami(most_recent=True,
    owners=["amazon"],
    filters=[aws.ec2.GetAmiFilterArgs(
        name="name",
        values=["al2023-ami-2023.*-x86_64"],
    )])
example_instance = aws.ec2.Instance("exampleInstance",
    ami=amzn_linux_2023_ami.id,
    instance_type="c6a.2xlarge",
    subnet_id=example_subnet.id,
    cpu_options=aws.ec2.InstanceCpuOptionsArgs(
        core_count=2,
        threads_per_core=2,
    ),
    tags={
        "Name": "tf-example",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleVpc = new Aws.Ec2.Vpc("exampleVpc", new()
    {
        CidrBlock = "172.16.0.0/16",
        Tags = 
        {
            { "Name", "tf-example" },
        },
    });

    var exampleSubnet = new Aws.Ec2.Subnet("exampleSubnet", new()
    {
        VpcId = exampleVpc.Id,
        CidrBlock = "172.16.10.0/24",
        AvailabilityZone = "us-east-2a",
        Tags = 
        {
            { "Name", "tf-example" },
        },
    });

    var amzn_linux_2023_ami = Aws.Ec2.GetAmi.Invoke(new()
    {
        MostRecent = true,
        Owners = new[]
        {
            "amazon",
        },
        Filters = new[]
        {
            new Aws.Ec2.Inputs.GetAmiFilterInputArgs
            {
                Name = "name",
                Values = new[]
                {
                    "al2023-ami-2023.*-x86_64",
                },
            },
        },
    });

    var exampleInstance = new Aws.Ec2.Instance("exampleInstance", new()
    {
        Ami = amzn_linux_2023_ami.Apply(amzn_linux_2023_ami => amzn_linux_2023_ami.Apply(getAmiResult => getAmiResult.Id)),
        InstanceType = "c6a.2xlarge",
        SubnetId = exampleSubnet.Id,
        CpuOptions = new Aws.Ec2.Inputs.InstanceCpuOptionsArgs
        {
            CoreCount = 2,
            ThreadsPerCore = 2,
        },
        Tags = 
        {
            { "Name", "tf-example" },
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
		exampleVpc, err := ec2.NewVpc(ctx, "exampleVpc", &ec2.VpcArgs{
			CidrBlock: pulumi.String("172.16.0.0/16"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-example"),
			},
		})
		if err != nil {
			return err
		}
		exampleSubnet, err := ec2.NewSubnet(ctx, "exampleSubnet", &ec2.SubnetArgs{
			VpcId:            exampleVpc.ID(),
			CidrBlock:        pulumi.String("172.16.10.0/24"),
			AvailabilityZone: pulumi.String("us-east-2a"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-example"),
			},
		})
		if err != nil {
			return err
		}
		amzn_linux_2023_ami, err := ec2.LookupAmi(ctx, &ec2.LookupAmiArgs{
			MostRecent: pulumi.BoolRef(true),
			Owners: []string{
				"amazon",
			},
			Filters: []ec2.GetAmiFilter{
				{
					Name: "name",
					Values: []string{
						"al2023-ami-2023.*-x86_64",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		_, err = ec2.NewInstance(ctx, "exampleInstance", &ec2.InstanceArgs{
			Ami:          *pulumi.String(amzn_linux_2023_ami.Id),
			InstanceType: pulumi.String("c6a.2xlarge"),
			SubnetId:     exampleSubnet.ID(),
			CpuOptions: &ec2.InstanceCpuOptionsArgs{
				CoreCount:      pulumi.Int(2),
				ThreadsPerCore: pulumi.Int(2),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-example"),
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
import com.pulumi.aws.ec2.Vpc;
import com.pulumi.aws.ec2.VpcArgs;
import com.pulumi.aws.ec2.Subnet;
import com.pulumi.aws.ec2.SubnetArgs;
import com.pulumi.aws.ec2.Ec2Functions;
import com.pulumi.aws.ec2.inputs.GetAmiArgs;
import com.pulumi.aws.ec2.Instance;
import com.pulumi.aws.ec2.InstanceArgs;
import com.pulumi.aws.ec2.inputs.InstanceCpuOptionsArgs;
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
        var exampleVpc = new Vpc("exampleVpc", VpcArgs.builder()        
            .cidrBlock("172.16.0.0/16")
            .tags(Map.of("Name", "tf-example"))
            .build());

        var exampleSubnet = new Subnet("exampleSubnet", SubnetArgs.builder()        
            .vpcId(exampleVpc.id())
            .cidrBlock("172.16.10.0/24")
            .availabilityZone("us-east-2a")
            .tags(Map.of("Name", "tf-example"))
            .build());

        final var amzn-linux-2023-ami = Ec2Functions.getAmi(GetAmiArgs.builder()
            .mostRecent(true)
            .owners("amazon")
            .filters(GetAmiFilterArgs.builder()
                .name("name")
                .values("al2023-ami-2023.*-x86_64")
                .build())
            .build());

        var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()        
            .ami(amzn_linux_2023_ami.id())
            .instanceType("c6a.2xlarge")
            .subnetId(exampleSubnet.id())
            .cpuOptions(InstanceCpuOptionsArgs.builder()
                .coreCount(2)
                .threadsPerCore(2)
                .build())
            .tags(Map.of("Name", "tf-example"))
            .build());

    }
}
```
```yaml
resources:
  exampleVpc:
    type: aws:ec2:Vpc
    properties:
      cidrBlock: 172.16.0.0/16
      tags:
        Name: tf-example
  exampleSubnet:
    type: aws:ec2:Subnet
    properties:
      vpcId: ${exampleVpc.id}
      cidrBlock: 172.16.10.0/24
      availabilityZone: us-east-2a
      tags:
        Name: tf-example
  exampleInstance:
    type: aws:ec2:Instance
    properties:
      ami: ${["amzn-linux-2023-ami"].id}
      instanceType: c6a.2xlarge
      subnetId: ${exampleSubnet.id}
      cpuOptions:
        coreCount: 2
        threadsPerCore: 2
      tags:
        Name: tf-example
variables:
  amzn-linux-2023-ami:
    fn::invoke:
      Function: aws:ec2:getAmi
      Arguments:
        mostRecent: true
        owners:
          - amazon
        filters:
          - name: name
            values:
              - al2023-ami-2023.*-x86_64
```
{{% /example %}}
{{% example %}}
### Host resource group or Licence Manager registered AMI example

A host resource group is a collection of Dedicated Hosts that you can manage as a single entity. As you launch instances, License Manager allocates the hosts and launches instances on them based on the settings that you configured. You can add existing Dedicated Hosts to a host resource group and take advantage of automated host management through License Manager.

> **NOTE:** A dedicated host is automatically associated with a License Manager host resource group if **Allocate hosts automatically** is enabled. Otherwise, use the `host_resource_group_arn` argument to explicitly associate the instance with the host resource group.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const _this = new aws.ec2.Instance("this", {
    ami: "ami-0dcc1e21636832c5d",
    hostResourceGroupArn: "arn:aws:resource-groups:us-west-2:012345678901:group/win-testhost",
    instanceType: "m5.large",
    tenancy: "host",
});
```
```python
import pulumi
import pulumi_aws as aws

this = aws.ec2.Instance("this",
    ami="ami-0dcc1e21636832c5d",
    host_resource_group_arn="arn:aws:resource-groups:us-west-2:012345678901:group/win-testhost",
    instance_type="m5.large",
    tenancy="host")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var @this = new Aws.Ec2.Instance("this", new()
    {
        Ami = "ami-0dcc1e21636832c5d",
        HostResourceGroupArn = "arn:aws:resource-groups:us-west-2:012345678901:group/win-testhost",
        InstanceType = "m5.large",
        Tenancy = "host",
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
		_, err := ec2.NewInstance(ctx, "this", &ec2.InstanceArgs{
			Ami:                  pulumi.String("ami-0dcc1e21636832c5d"),
			HostResourceGroupArn: pulumi.String("arn:aws:resource-groups:us-west-2:012345678901:group/win-testhost"),
			InstanceType:         pulumi.String("m5.large"),
			Tenancy:              pulumi.String("host"),
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
        var this_ = new Instance("this", InstanceArgs.builder()        
            .ami("ami-0dcc1e21636832c5d")
            .hostResourceGroupArn("arn:aws:resource-groups:us-west-2:012345678901:group/win-testhost")
            .instanceType("m5.large")
            .tenancy("host")
            .build());

    }
}
```
```yaml
resources:
  this:
    type: aws:ec2:Instance
    properties:
      ami: ami-0dcc1e21636832c5d
      hostResourceGroupArn: arn:aws:resource-groups:us-west-2:012345678901:group/win-testhost
      instanceType: m5.large
      tenancy: host
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import instances using the `id`. For example:

```sh
 $ pulumi import aws:ec2/instance:Instance web i-12345678
```
 