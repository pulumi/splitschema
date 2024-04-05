Provides an Amazon Connect Phone Number resource. For more information see
[Amazon Connect: Getting Started](https://docs.aws.amazon.com/connect/latest/adminguide/amazon-connect-get-started.html)

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.PhoneNumber("example", {
    targetArn: aws_connect_instance.example.arn,
    countryCode: "US",
    type: "DID",
    tags: {
        hello: "world",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.PhoneNumber("example",
    target_arn=aws_connect_instance["example"]["arn"],
    country_code="US",
    type="DID",
    tags={
        "hello": "world",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.PhoneNumber("example", new()
    {
        TargetArn = aws_connect_instance.Example.Arn,
        CountryCode = "US",
        Type = "DID",
        Tags = 
        {
            { "hello", "world" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewPhoneNumber(ctx, "example", &connect.PhoneNumberArgs{
			TargetArn:   pulumi.Any(aws_connect_instance.Example.Arn),
			CountryCode: pulumi.String("US"),
			Type:        pulumi.String("DID"),
			Tags: pulumi.StringMap{
				"hello": pulumi.String("world"),
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
import com.pulumi.aws.connect.PhoneNumber;
import com.pulumi.aws.connect.PhoneNumberArgs;
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
        var example = new PhoneNumber("example", PhoneNumberArgs.builder()        
            .targetArn(aws_connect_instance.example().arn())
            .countryCode("US")
            .type("DID")
            .tags(Map.of("hello", "world"))
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:PhoneNumber
    properties:
      targetArn: ${aws_connect_instance.example.arn}
      countryCode: US
      type: DID
      tags:
        hello: world
```
{{% /example %}}
{{% example %}}
### Description

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.PhoneNumber("example", {
    targetArn: aws_connect_instance.example.arn,
    countryCode: "US",
    type: "DID",
    description: "example description",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.PhoneNumber("example",
    target_arn=aws_connect_instance["example"]["arn"],
    country_code="US",
    type="DID",
    description="example description")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.PhoneNumber("example", new()
    {
        TargetArn = aws_connect_instance.Example.Arn,
        CountryCode = "US",
        Type = "DID",
        Description = "example description",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewPhoneNumber(ctx, "example", &connect.PhoneNumberArgs{
			TargetArn:   pulumi.Any(aws_connect_instance.Example.Arn),
			CountryCode: pulumi.String("US"),
			Type:        pulumi.String("DID"),
			Description: pulumi.String("example description"),
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
import com.pulumi.aws.connect.PhoneNumber;
import com.pulumi.aws.connect.PhoneNumberArgs;
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
        var example = new PhoneNumber("example", PhoneNumberArgs.builder()        
            .targetArn(aws_connect_instance.example().arn())
            .countryCode("US")
            .type("DID")
            .description("example description")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:PhoneNumber
    properties:
      targetArn: ${aws_connect_instance.example.arn}
      countryCode: US
      type: DID
      description: example description
```
{{% /example %}}
{{% example %}}
### Prefix to filter phone numbers

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.connect.PhoneNumber("example", {
    targetArn: aws_connect_instance.example.arn,
    countryCode: "US",
    type: "DID",
    prefix: "+18005",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.connect.PhoneNumber("example",
    target_arn=aws_connect_instance["example"]["arn"],
    country_code="US",
    type="DID",
    prefix="+18005")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Connect.PhoneNumber("example", new()
    {
        TargetArn = aws_connect_instance.Example.Arn,
        CountryCode = "US",
        Type = "DID",
        Prefix = "+18005",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/connect"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := connect.NewPhoneNumber(ctx, "example", &connect.PhoneNumberArgs{
			TargetArn:   pulumi.Any(aws_connect_instance.Example.Arn),
			CountryCode: pulumi.String("US"),
			Type:        pulumi.String("DID"),
			Prefix:      pulumi.String("+18005"),
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
import com.pulumi.aws.connect.PhoneNumber;
import com.pulumi.aws.connect.PhoneNumberArgs;
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
        var example = new PhoneNumber("example", PhoneNumberArgs.builder()        
            .targetArn(aws_connect_instance.example().arn())
            .countryCode("US")
            .type("DID")
            .prefix("+18005")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:connect:PhoneNumber
    properties:
      targetArn: ${aws_connect_instance.example.arn}
      countryCode: US
      type: DID
      prefix: '+18005'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Amazon Connect Phone Numbers using its `id`. For example:

```sh
 $ pulumi import aws:connect/phoneNumber:PhoneNumber example 12345678-abcd-1234-efgh-9876543210ab
```
 