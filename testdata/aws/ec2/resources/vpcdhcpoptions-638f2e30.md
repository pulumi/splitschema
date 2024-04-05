Provides a VPC DHCP Options resource.

{{% examples %}}
## Example Usage
{{% example %}}

Basic usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const dnsResolver = new aws.ec2.VpcDhcpOptions("dnsResolver", {domainNameServers: [
    "8.8.8.8",
    "8.8.4.4",
]});
```
```python
import pulumi
import pulumi_aws as aws

dns_resolver = aws.ec2.VpcDhcpOptions("dnsResolver", domain_name_servers=[
    "8.8.8.8",
    "8.8.4.4",
])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var dnsResolver = new Aws.Ec2.VpcDhcpOptions("dnsResolver", new()
    {
        DomainNameServers = new[]
        {
            "8.8.8.8",
            "8.8.4.4",
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
		_, err := ec2.NewVpcDhcpOptions(ctx, "dnsResolver", &ec2.VpcDhcpOptionsArgs{
			DomainNameServers: pulumi.StringArray{
				pulumi.String("8.8.8.8"),
				pulumi.String("8.8.4.4"),
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
import com.pulumi.aws.ec2.VpcDhcpOptions;
import com.pulumi.aws.ec2.VpcDhcpOptionsArgs;
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
        var dnsResolver = new VpcDhcpOptions("dnsResolver", VpcDhcpOptionsArgs.builder()        
            .domainNameServers(            
                "8.8.8.8",
                "8.8.4.4")
            .build());

    }
}
```
```yaml
resources:
  dnsResolver:
    type: aws:ec2:VpcDhcpOptions
    properties:
      domainNameServers:
        - 8.8.8.8
        - 8.8.4.4
```

Full usage:

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.ec2.VpcDhcpOptions("foo", {
    domainName: "service.consul",
    domainNameServers: [
        "127.0.0.1",
        "10.0.0.2",
    ],
    netbiosNameServers: ["127.0.0.1"],
    netbiosNodeType: "2",
    ntpServers: ["127.0.0.1"],
    tags: {
        Name: "foo-name",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.ec2.VpcDhcpOptions("foo",
    domain_name="service.consul",
    domain_name_servers=[
        "127.0.0.1",
        "10.0.0.2",
    ],
    netbios_name_servers=["127.0.0.1"],
    netbios_node_type="2",
    ntp_servers=["127.0.0.1"],
    tags={
        "Name": "foo-name",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Ec2.VpcDhcpOptions("foo", new()
    {
        DomainName = "service.consul",
        DomainNameServers = new[]
        {
            "127.0.0.1",
            "10.0.0.2",
        },
        NetbiosNameServers = new[]
        {
            "127.0.0.1",
        },
        NetbiosNodeType = "2",
        NtpServers = new[]
        {
            "127.0.0.1",
        },
        Tags = 
        {
            { "Name", "foo-name" },
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
		_, err := ec2.NewVpcDhcpOptions(ctx, "foo", &ec2.VpcDhcpOptionsArgs{
			DomainName: pulumi.String("service.consul"),
			DomainNameServers: pulumi.StringArray{
				pulumi.String("127.0.0.1"),
				pulumi.String("10.0.0.2"),
			},
			NetbiosNameServers: pulumi.StringArray{
				pulumi.String("127.0.0.1"),
			},
			NetbiosNodeType: pulumi.String("2"),
			NtpServers: pulumi.StringArray{
				pulumi.String("127.0.0.1"),
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("foo-name"),
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
import com.pulumi.aws.ec2.VpcDhcpOptions;
import com.pulumi.aws.ec2.VpcDhcpOptionsArgs;
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
        var foo = new VpcDhcpOptions("foo", VpcDhcpOptionsArgs.builder()        
            .domainName("service.consul")
            .domainNameServers(            
                "127.0.0.1",
                "10.0.0.2")
            .netbiosNameServers("127.0.0.1")
            .netbiosNodeType(2)
            .ntpServers("127.0.0.1")
            .tags(Map.of("Name", "foo-name"))
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:ec2:VpcDhcpOptions
    properties:
      domainName: service.consul
      domainNameServers:
        - 127.0.0.1
        - 10.0.0.2
      netbiosNameServers:
        - 127.0.0.1
      netbiosNodeType: 2
      ntpServers:
        - 127.0.0.1
      tags:
        Name: foo-name
```
{{% /example %}}
{{% /examples %}}
## Remarks

* Notice that all arguments are optional but you have to specify at least one argument.
* `domain_name_servers`, `netbios_name_servers`, `ntp_servers` are limited by AWS to maximum four servers only.
* To actually use the DHCP Options Set you need to associate it to a VPC using `aws.ec2.VpcDhcpOptionsAssociation`.
* If you delete a DHCP Options Set, all VPCs using it will be associated to AWS's `default` DHCP Option Set.
* In most cases unless you're configuring your own DNS you'll want to set `domain_name_servers` to `AmazonProvidedDNS`.


## Import

Using `pulumi import`, import VPC DHCP Options using the DHCP Options `id`. For example:

```sh
 $ pulumi import aws:ec2/vpcDhcpOptions:VpcDhcpOptions my_options dopt-d9070ebb
```
 