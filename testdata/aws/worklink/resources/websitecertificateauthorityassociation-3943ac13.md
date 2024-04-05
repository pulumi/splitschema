{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const example = new aws.worklink.Fleet("example", {});
const test = new aws.worklink.WebsiteCertificateAuthorityAssociation("test", {
    fleetArn: aws_worklink_fleet.test.arn,
    certificate: fs.readFileSync("certificate.pem", "utf8"),
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.worklink.Fleet("example")
test = aws.worklink.WebsiteCertificateAuthorityAssociation("test",
    fleet_arn=aws_worklink_fleet["test"]["arn"],
    certificate=(lambda path: open(path).read())("certificate.pem"))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.WorkLink.Fleet("example");

    var test = new Aws.WorkLink.WebsiteCertificateAuthorityAssociation("test", new()
    {
        FleetArn = aws_worklink_fleet.Test.Arn,
        Certificate = File.ReadAllText("certificate.pem"),
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/worklink"
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
		_, err := worklink.NewFleet(ctx, "example", nil)
		if err != nil {
			return err
		}
		_, err = worklink.NewWebsiteCertificateAuthorityAssociation(ctx, "test", &worklink.WebsiteCertificateAuthorityAssociationArgs{
			FleetArn:    pulumi.Any(aws_worklink_fleet.Test.Arn),
			Certificate: readFileOrPanic("certificate.pem"),
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
import com.pulumi.aws.worklink.Fleet;
import com.pulumi.aws.worklink.WebsiteCertificateAuthorityAssociation;
import com.pulumi.aws.worklink.WebsiteCertificateAuthorityAssociationArgs;
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
        var example = new Fleet("example");

        var test = new WebsiteCertificateAuthorityAssociation("test", WebsiteCertificateAuthorityAssociationArgs.builder()        
            .fleetArn(aws_worklink_fleet.test().arn())
            .certificate(Files.readString(Paths.get("certificate.pem")))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:worklink:Fleet
  test:
    type: aws:worklink:WebsiteCertificateAuthorityAssociation
    properties:
      fleetArn: ${aws_worklink_fleet.test.arn}
      certificate:
        fn::readFile: certificate.pem
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import WorkLink Website Certificate Authority using `FLEET-ARN,WEBSITE-CA-ID`. For example:

```sh
 $ pulumi import aws:worklink/websiteCertificateAuthorityAssociation:WebsiteCertificateAuthorityAssociation example arn:aws:worklink::123456789012:fleet/example,abcdefghijk
```
 