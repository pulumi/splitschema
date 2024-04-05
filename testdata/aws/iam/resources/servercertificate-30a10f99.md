Provides an IAM Server Certificate resource to upload Server Certificates.
Certs uploaded to IAM can easily work with other AWS services such as:

- AWS Elastic Beanstalk
- Elastic Load Balancing
- CloudFront
- AWS OpsWorks

For information about server certificates in IAM, see [Managing Server
Certificates][2] in AWS Documentation.

{{% examples %}}
## Example Usage
{{% example %}}

**Using certs on file:**

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const testCert = new aws.iam.ServerCertificate("testCert", {
    certificateBody: fs.readFileSync("self-ca-cert.pem", "utf8"),
    privateKey: fs.readFileSync("test-key.pem", "utf8"),
});
```
```python
import pulumi
import pulumi_aws as aws

test_cert = aws.iam.ServerCertificate("testCert",
    certificate_body=(lambda path: open(path).read())("self-ca-cert.pem"),
    private_key=(lambda path: open(path).read())("test-key.pem"))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCert = new Aws.Iam.ServerCertificate("testCert", new()
    {
        CertificateBody = File.ReadAllText("self-ca-cert.pem"),
        PrivateKey = File.ReadAllText("test-key.pem"),
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewServerCertificate(ctx, "testCert", &iam.ServerCertificateArgs{
			CertificateBody: readFileOrPanic("self-ca-cert.pem"),
			PrivateKey:      readFileOrPanic("test-key.pem"),
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
import com.pulumi.aws.iam.ServerCertificate;
import com.pulumi.aws.iam.ServerCertificateArgs;
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
        var testCert = new ServerCertificate("testCert", ServerCertificateArgs.builder()        
            .certificateBody(Files.readString(Paths.get("self-ca-cert.pem")))
            .privateKey(Files.readString(Paths.get("test-key.pem")))
            .build());

    }
}
```
```yaml
resources:
  testCert:
    type: aws:iam:ServerCertificate
    properties:
      certificateBody:
        fn::readFile: self-ca-cert.pem
      privateKey:
        fn::readFile: test-key.pem
```

**Example with cert in-line:**

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testCertAlt = new aws.iam.ServerCertificate("testCertAlt", {
    certificateBody: `-----BEGIN CERTIFICATE-----
[......] # cert contents
-----END CERTIFICATE-----

`,
    privateKey: `-----BEGIN RSA PRIVATE KEY-----
[......] # cert contents
-----END RSA PRIVATE KEY-----

`,
});
```
```python
import pulumi
import pulumi_aws as aws

test_cert_alt = aws.iam.ServerCertificate("testCertAlt",
    certificate_body="""-----BEGIN CERTIFICATE-----
[......] # cert contents
-----END CERTIFICATE-----

""",
    private_key="""-----BEGIN RSA PRIVATE KEY-----
[......] # cert contents
-----END RSA PRIVATE KEY-----

""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCertAlt = new Aws.Iam.ServerCertificate("testCertAlt", new()
    {
        CertificateBody = @"-----BEGIN CERTIFICATE-----
[......] # cert contents
-----END CERTIFICATE-----

",
        PrivateKey = @"-----BEGIN RSA PRIVATE KEY-----
[......] # cert contents
-----END RSA PRIVATE KEY-----

",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iam.NewServerCertificate(ctx, "testCertAlt", &iam.ServerCertificateArgs{
			CertificateBody: pulumi.String("-----BEGIN CERTIFICATE-----\n[......] # cert contents\n-----END CERTIFICATE-----\n\n"),
			PrivateKey:      pulumi.String("-----BEGIN RSA PRIVATE KEY-----\n[......] # cert contents\n-----END RSA PRIVATE KEY-----\n\n"),
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
import com.pulumi.aws.iam.ServerCertificate;
import com.pulumi.aws.iam.ServerCertificateArgs;
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
        var testCertAlt = new ServerCertificate("testCertAlt", ServerCertificateArgs.builder()        
            .certificateBody("""
-----BEGIN CERTIFICATE-----
[......] # cert contents
-----END CERTIFICATE-----

            """)
            .privateKey("""
-----BEGIN RSA PRIVATE KEY-----
[......] # cert contents
-----END RSA PRIVATE KEY-----

            """)
            .build());

    }
}
```
```yaml
resources:
  testCertAlt:
    type: aws:iam:ServerCertificate
    properties:
      certificateBody: |+
        -----BEGIN CERTIFICATE-----
        [......] # cert contents
        -----END CERTIFICATE-----

      privateKey: |+
        -----BEGIN RSA PRIVATE KEY-----
        [......] # cert contents
        -----END RSA PRIVATE KEY-----
```

**Use in combination with an AWS ELB resource:**

Some properties of an IAM Server Certificates cannot be updated while they are
in use. In order for the provider to effectively manage a Certificate in this situation, it is
recommended you utilize the `name_prefix` attribute and enable the
`create_before_destroy`. This will allow this provider
to create a new, updated `aws.iam.ServerCertificate` resource and replace it in
dependant resources before attempting to destroy the old version.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const testCert = new aws.iam.ServerCertificate("testCert", {
    namePrefix: "example-cert",
    certificateBody: fs.readFileSync("self-ca-cert.pem", "utf8"),
    privateKey: fs.readFileSync("test-key.pem", "utf8"),
});
const ourapp = new aws.elb.LoadBalancer("ourapp", {
    availabilityZones: ["us-west-2a"],
    crossZoneLoadBalancing: true,
    listeners: [{
        instancePort: 8000,
        instanceProtocol: "http",
        lbPort: 443,
        lbProtocol: "https",
        sslCertificateId: testCert.arn,
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

test_cert = aws.iam.ServerCertificate("testCert",
    name_prefix="example-cert",
    certificate_body=(lambda path: open(path).read())("self-ca-cert.pem"),
    private_key=(lambda path: open(path).read())("test-key.pem"))
ourapp = aws.elb.LoadBalancer("ourapp",
    availability_zones=["us-west-2a"],
    cross_zone_load_balancing=True,
    listeners=[aws.elb.LoadBalancerListenerArgs(
        instance_port=8000,
        instance_protocol="http",
        lb_port=443,
        lb_protocol="https",
        ssl_certificate_id=test_cert.arn,
    )])
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testCert = new Aws.Iam.ServerCertificate("testCert", new()
    {
        NamePrefix = "example-cert",
        CertificateBody = File.ReadAllText("self-ca-cert.pem"),
        PrivateKey = File.ReadAllText("test-key.pem"),
    });

    var ourapp = new Aws.Elb.LoadBalancer("ourapp", new()
    {
        AvailabilityZones = new[]
        {
            "us-west-2a",
        },
        CrossZoneLoadBalancing = true,
        Listeners = new[]
        {
            new Aws.Elb.Inputs.LoadBalancerListenerArgs
            {
                InstancePort = 8000,
                InstanceProtocol = "http",
                LbPort = 443,
                LbProtocol = "https",
                SslCertificateId = testCert.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/elb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		testCert, err := iam.NewServerCertificate(ctx, "testCert", &iam.ServerCertificateArgs{
			NamePrefix:      pulumi.String("example-cert"),
			CertificateBody: readFileOrPanic("self-ca-cert.pem"),
			PrivateKey:      readFileOrPanic("test-key.pem"),
		})
		if err != nil {
			return err
		}
		_, err = elb.NewLoadBalancer(ctx, "ourapp", &elb.LoadBalancerArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("us-west-2a"),
			},
			CrossZoneLoadBalancing: pulumi.Bool(true),
			Listeners: elb.LoadBalancerListenerArray{
				&elb.LoadBalancerListenerArgs{
					InstancePort:     pulumi.Int(8000),
					InstanceProtocol: pulumi.String("http"),
					LbPort:           pulumi.Int(443),
					LbProtocol:       pulumi.String("https"),
					SslCertificateId: testCert.Arn,
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
import com.pulumi.aws.iam.ServerCertificate;
import com.pulumi.aws.iam.ServerCertificateArgs;
import com.pulumi.aws.elb.LoadBalancer;
import com.pulumi.aws.elb.LoadBalancerArgs;
import com.pulumi.aws.elb.inputs.LoadBalancerListenerArgs;
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
        var testCert = new ServerCertificate("testCert", ServerCertificateArgs.builder()        
            .namePrefix("example-cert")
            .certificateBody(Files.readString(Paths.get("self-ca-cert.pem")))
            .privateKey(Files.readString(Paths.get("test-key.pem")))
            .build());

        var ourapp = new LoadBalancer("ourapp", LoadBalancerArgs.builder()        
            .availabilityZones("us-west-2a")
            .crossZoneLoadBalancing(true)
            .listeners(LoadBalancerListenerArgs.builder()
                .instancePort(8000)
                .instanceProtocol("http")
                .lbPort(443)
                .lbProtocol("https")
                .sslCertificateId(testCert.arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  testCert:
    type: aws:iam:ServerCertificate
    properties:
      namePrefix: example-cert
      certificateBody:
        fn::readFile: self-ca-cert.pem
      privateKey:
        fn::readFile: test-key.pem
  ourapp:
    type: aws:elb:LoadBalancer
    properties:
      availabilityZones:
        - us-west-2a
      crossZoneLoadBalancing: true
      listeners:
        - instancePort: 8000
          instanceProtocol: http
          lbPort: 443
          lbProtocol: https
          sslCertificateId: ${testCert.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IAM Server Certificates using the `name`. For example:

```sh
 $ pulumi import aws:iam/serverCertificate:ServerCertificate certificate example.com-certificate-until-2018
```
 