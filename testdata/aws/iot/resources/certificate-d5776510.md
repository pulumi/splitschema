Creates and manages an AWS IoT certificate.

{{% examples %}}
## Example Usage
{{% example %}}
### With CSR

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const cert = new aws.iot.Certificate("cert", {
    csr: fs.readFileSync("/my/csr.pem", "utf8"),
    active: true,
});
```
```python
import pulumi
import pulumi_aws as aws

cert = aws.iot.Certificate("cert",
    csr=(lambda path: open(path).read())("/my/csr.pem"),
    active=True)
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var cert = new Aws.Iot.Certificate("cert", new()
    {
        Csr = File.ReadAllText("/my/csr.pem"),
        Active = true,
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
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
		_, err := iot.NewCertificate(ctx, "cert", &iot.CertificateArgs{
			Csr:    readFileOrPanic("/my/csr.pem"),
			Active: pulumi.Bool(true),
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
import com.pulumi.aws.iot.Certificate;
import com.pulumi.aws.iot.CertificateArgs;
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
        var cert = new Certificate("cert", CertificateArgs.builder()        
            .csr(Files.readString(Paths.get("/my/csr.pem")))
            .active(true)
            .build());

    }
}
```
```yaml
resources:
  cert:
    type: aws:iot:Certificate
    properties:
      csr:
        fn::readFile: /my/csr.pem
      active: true
```
{{% /example %}}
{{% example %}}
### Without CSR

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const cert = new aws.iot.Certificate("cert", {active: true});
```
```python
import pulumi
import pulumi_aws as aws

cert = aws.iot.Certificate("cert", active=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var cert = new Aws.Iot.Certificate("cert", new()
    {
        Active = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := iot.NewCertificate(ctx, "cert", &iot.CertificateArgs{
			Active: pulumi.Bool(true),
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
import com.pulumi.aws.iot.Certificate;
import com.pulumi.aws.iot.CertificateArgs;
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
        var cert = new Certificate("cert", CertificateArgs.builder()        
            .active(true)
            .build());

    }
}
```
```yaml
resources:
  cert:
    type: aws:iot:Certificate
    properties:
      active: true
```
{{% /example %}}
{{% example %}}
### From existing certificate without a CA

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const cert = new aws.iot.Certificate("cert", {
    certificatePem: fs.readFileSync("/my/cert.pem", "utf8"),
    active: true,
});
```
```python
import pulumi
import pulumi_aws as aws

cert = aws.iot.Certificate("cert",
    certificate_pem=(lambda path: open(path).read())("/my/cert.pem"),
    active=True)
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var cert = new Aws.Iot.Certificate("cert", new()
    {
        CertificatePem = File.ReadAllText("/my/cert.pem"),
        Active = true,
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
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
		_, err := iot.NewCertificate(ctx, "cert", &iot.CertificateArgs{
			CertificatePem: readFileOrPanic("/my/cert.pem"),
			Active:         pulumi.Bool(true),
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
import com.pulumi.aws.iot.Certificate;
import com.pulumi.aws.iot.CertificateArgs;
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
        var cert = new Certificate("cert", CertificateArgs.builder()        
            .certificatePem(Files.readString(Paths.get("/my/cert.pem")))
            .active(true)
            .build());

    }
}
```
```yaml
resources:
  cert:
    type: aws:iot:Certificate
    properties:
      certificatePem:
        fn::readFile: /my/cert.pem
      active: true
```
{{% /example %}}
{{% /examples %}}