Provides an SNS platform application resource

{{% examples %}}
## Example Usage
{{% example %}}
### Apple Push Notification Service (APNS) using certificate-based authentication

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const apnsApplication = new aws.sns.PlatformApplication("apnsApplication", {
    platform: "APNS",
    platformCredential: "<APNS PRIVATE KEY>",
    platformPrincipal: "<APNS CERTIFICATE>",
});
```
```python
import pulumi
import pulumi_aws as aws

apns_application = aws.sns.PlatformApplication("apnsApplication",
    platform="APNS",
    platform_credential="<APNS PRIVATE KEY>",
    platform_principal="<APNS CERTIFICATE>")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var apnsApplication = new Aws.Sns.PlatformApplication("apnsApplication", new()
    {
        Platform = "APNS",
        PlatformCredential = "<APNS PRIVATE KEY>",
        PlatformPrincipal = "<APNS CERTIFICATE>",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sns.NewPlatformApplication(ctx, "apnsApplication", &sns.PlatformApplicationArgs{
			Platform:           pulumi.String("APNS"),
			PlatformCredential: pulumi.String("<APNS PRIVATE KEY>"),
			PlatformPrincipal:  pulumi.String("<APNS CERTIFICATE>"),
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
import com.pulumi.aws.sns.PlatformApplication;
import com.pulumi.aws.sns.PlatformApplicationArgs;
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
        var apnsApplication = new PlatformApplication("apnsApplication", PlatformApplicationArgs.builder()        
            .platform("APNS")
            .platformCredential("<APNS PRIVATE KEY>")
            .platformPrincipal("<APNS CERTIFICATE>")
            .build());

    }
}
```
```yaml
resources:
  apnsApplication:
    type: aws:sns:PlatformApplication
    properties:
      platform: APNS
      platformCredential: <APNS PRIVATE KEY>
      platformPrincipal: <APNS CERTIFICATE>
```
{{% /example %}}
{{% example %}}
### Apple Push Notification Service (APNS) using token-based authentication

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const apnsApplication = new aws.sns.PlatformApplication("apnsApplication", {
    applePlatformBundleId: "<APPLE BUNDLE ID>",
    applePlatformTeamId: "<APPLE TEAM ID>",
    platform: "APNS",
    platformCredential: "<APNS SIGNING KEY>",
    platformPrincipal: "<APNS SIGNING KEY ID>",
});
```
```python
import pulumi
import pulumi_aws as aws

apns_application = aws.sns.PlatformApplication("apnsApplication",
    apple_platform_bundle_id="<APPLE BUNDLE ID>",
    apple_platform_team_id="<APPLE TEAM ID>",
    platform="APNS",
    platform_credential="<APNS SIGNING KEY>",
    platform_principal="<APNS SIGNING KEY ID>")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var apnsApplication = new Aws.Sns.PlatformApplication("apnsApplication", new()
    {
        ApplePlatformBundleId = "<APPLE BUNDLE ID>",
        ApplePlatformTeamId = "<APPLE TEAM ID>",
        Platform = "APNS",
        PlatformCredential = "<APNS SIGNING KEY>",
        PlatformPrincipal = "<APNS SIGNING KEY ID>",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sns.NewPlatformApplication(ctx, "apnsApplication", &sns.PlatformApplicationArgs{
			ApplePlatformBundleId: pulumi.String("<APPLE BUNDLE ID>"),
			ApplePlatformTeamId:   pulumi.String("<APPLE TEAM ID>"),
			Platform:              pulumi.String("APNS"),
			PlatformCredential:    pulumi.String("<APNS SIGNING KEY>"),
			PlatformPrincipal:     pulumi.String("<APNS SIGNING KEY ID>"),
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
import com.pulumi.aws.sns.PlatformApplication;
import com.pulumi.aws.sns.PlatformApplicationArgs;
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
        var apnsApplication = new PlatformApplication("apnsApplication", PlatformApplicationArgs.builder()        
            .applePlatformBundleId("<APPLE BUNDLE ID>")
            .applePlatformTeamId("<APPLE TEAM ID>")
            .platform("APNS")
            .platformCredential("<APNS SIGNING KEY>")
            .platformPrincipal("<APNS SIGNING KEY ID>")
            .build());

    }
}
```
```yaml
resources:
  apnsApplication:
    type: aws:sns:PlatformApplication
    properties:
      applePlatformBundleId: <APPLE BUNDLE ID>
      applePlatformTeamId: <APPLE TEAM ID>
      platform: APNS
      platformCredential: <APNS SIGNING KEY>
      platformPrincipal: <APNS SIGNING KEY ID>
```
{{% /example %}}
{{% example %}}
### Google Cloud Messaging (GCM)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const gcmApplication = new aws.sns.PlatformApplication("gcmApplication", {
    platform: "GCM",
    platformCredential: "<GCM API KEY>",
});
```
```python
import pulumi
import pulumi_aws as aws

gcm_application = aws.sns.PlatformApplication("gcmApplication",
    platform="GCM",
    platform_credential="<GCM API KEY>")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var gcmApplication = new Aws.Sns.PlatformApplication("gcmApplication", new()
    {
        Platform = "GCM",
        PlatformCredential = "<GCM API KEY>",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/sns"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sns.NewPlatformApplication(ctx, "gcmApplication", &sns.PlatformApplicationArgs{
			Platform:           pulumi.String("GCM"),
			PlatformCredential: pulumi.String("<GCM API KEY>"),
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
import com.pulumi.aws.sns.PlatformApplication;
import com.pulumi.aws.sns.PlatformApplicationArgs;
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
        var gcmApplication = new PlatformApplication("gcmApplication", PlatformApplicationArgs.builder()        
            .platform("GCM")
            .platformCredential("<GCM API KEY>")
            .build());

    }
}
```
```yaml
resources:
  gcmApplication:
    type: aws:sns:PlatformApplication
    properties:
      platform: GCM
      platformCredential: <GCM API KEY>
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import SNS platform applications using the ARN. For example:

```sh
 $ pulumi import aws:sns/platformApplication:PlatformApplication gcm_application arn:aws:sns:us-west-2:0123456789012:app/GCM/gcm_application
```
 