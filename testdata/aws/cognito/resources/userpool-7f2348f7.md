Provides a Cognito User Pool resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic configuration

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const pool = new aws.cognito.UserPool("pool", {});
```
```python
import pulumi
import pulumi_aws as aws

pool = aws.cognito.UserPool("pool")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var pool = new Aws.Cognito.UserPool("pool");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognito.NewUserPool(ctx, "pool", nil)
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
import com.pulumi.aws.cognito.UserPool;
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
        var pool = new UserPool("pool");

    }
}
```
```yaml
resources:
  pool:
    type: aws:cognito:UserPool
```
{{% /example %}}
{{% example %}}
### Enabling SMS and Software Token Multi-Factor Authentication

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

// ... other configuration ...
const example = new aws.cognito.UserPool("example", {
    mfaConfiguration: "ON",
    smsAuthenticationMessage: "Your code is {####}",
    smsConfiguration: {
        externalId: "example",
        snsCallerArn: aws_iam_role.example.arn,
        snsRegion: "us-east-1",
    },
    softwareTokenMfaConfiguration: {
        enabled: true,
    },
});
```
```python
import pulumi
import pulumi_aws as aws

# ... other configuration ...
example = aws.cognito.UserPool("example",
    mfa_configuration="ON",
    sms_authentication_message="Your code is {####}",
    sms_configuration=aws.cognito.UserPoolSmsConfigurationArgs(
        external_id="example",
        sns_caller_arn=aws_iam_role["example"]["arn"],
        sns_region="us-east-1",
    ),
    software_token_mfa_configuration=aws.cognito.UserPoolSoftwareTokenMfaConfigurationArgs(
        enabled=True,
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    // ... other configuration ...
    var example = new Aws.Cognito.UserPool("example", new()
    {
        MfaConfiguration = "ON",
        SmsAuthenticationMessage = "Your code is {####}",
        SmsConfiguration = new Aws.Cognito.Inputs.UserPoolSmsConfigurationArgs
        {
            ExternalId = "example",
            SnsCallerArn = aws_iam_role.Example.Arn,
            SnsRegion = "us-east-1",
        },
        SoftwareTokenMfaConfiguration = new Aws.Cognito.Inputs.UserPoolSoftwareTokenMfaConfigurationArgs
        {
            Enabled = true,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognito.NewUserPool(ctx, "example", &cognito.UserPoolArgs{
			MfaConfiguration:         pulumi.String("ON"),
			SmsAuthenticationMessage: pulumi.String("Your code is {####}"),
			SmsConfiguration: &cognito.UserPoolSmsConfigurationArgs{
				ExternalId:   pulumi.String("example"),
				SnsCallerArn: pulumi.Any(aws_iam_role.Example.Arn),
				SnsRegion:    pulumi.String("us-east-1"),
			},
			SoftwareTokenMfaConfiguration: &cognito.UserPoolSoftwareTokenMfaConfigurationArgs{
				Enabled: pulumi.Bool(true),
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolArgs;
import com.pulumi.aws.cognito.inputs.UserPoolSmsConfigurationArgs;
import com.pulumi.aws.cognito.inputs.UserPoolSoftwareTokenMfaConfigurationArgs;
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
        var example = new UserPool("example", UserPoolArgs.builder()        
            .mfaConfiguration("ON")
            .smsAuthenticationMessage("Your code is {####}")
            .smsConfiguration(UserPoolSmsConfigurationArgs.builder()
                .externalId("example")
                .snsCallerArn(aws_iam_role.example().arn())
                .snsRegion("us-east-1")
                .build())
            .softwareTokenMfaConfiguration(UserPoolSoftwareTokenMfaConfigurationArgs.builder()
                .enabled(true)
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cognito:UserPool
    properties:
      mfaConfiguration: ON
      smsAuthenticationMessage: Your code is {####}
      smsConfiguration:
        externalId: example
        snsCallerArn: ${aws_iam_role.example.arn}
        snsRegion: us-east-1
      softwareTokenMfaConfiguration:
        enabled: true
```
{{% /example %}}
{{% example %}}
### Using Account Recovery Setting

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.cognito.UserPool("test", {accountRecoverySetting: {
    recoveryMechanisms: [
        {
            name: "verified_email",
            priority: 1,
        },
        {
            name: "verified_phone_number",
            priority: 2,
        },
    ],
}});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.cognito.UserPool("test", account_recovery_setting=aws.cognito.UserPoolAccountRecoverySettingArgs(
    recovery_mechanisms=[
        aws.cognito.UserPoolAccountRecoverySettingRecoveryMechanismArgs(
            name="verified_email",
            priority=1,
        ),
        aws.cognito.UserPoolAccountRecoverySettingRecoveryMechanismArgs(
            name="verified_phone_number",
            priority=2,
        ),
    ],
))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Cognito.UserPool("test", new()
    {
        AccountRecoverySetting = new Aws.Cognito.Inputs.UserPoolAccountRecoverySettingArgs
        {
            RecoveryMechanisms = new[]
            {
                new Aws.Cognito.Inputs.UserPoolAccountRecoverySettingRecoveryMechanismArgs
                {
                    Name = "verified_email",
                    Priority = 1,
                },
                new Aws.Cognito.Inputs.UserPoolAccountRecoverySettingRecoveryMechanismArgs
                {
                    Name = "verified_phone_number",
                    Priority = 2,
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cognito.NewUserPool(ctx, "test", &cognito.UserPoolArgs{
			AccountRecoverySetting: &cognito.UserPoolAccountRecoverySettingArgs{
				RecoveryMechanisms: cognito.UserPoolAccountRecoverySettingRecoveryMechanismArray{
					&cognito.UserPoolAccountRecoverySettingRecoveryMechanismArgs{
						Name:     pulumi.String("verified_email"),
						Priority: pulumi.Int(1),
					},
					&cognito.UserPoolAccountRecoverySettingRecoveryMechanismArgs{
						Name:     pulumi.String("verified_phone_number"),
						Priority: pulumi.Int(2),
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
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolArgs;
import com.pulumi.aws.cognito.inputs.UserPoolAccountRecoverySettingArgs;
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
        var test = new UserPool("test", UserPoolArgs.builder()        
            .accountRecoverySetting(UserPoolAccountRecoverySettingArgs.builder()
                .recoveryMechanisms(                
                    UserPoolAccountRecoverySettingRecoveryMechanismArgs.builder()
                        .name("verified_email")
                        .priority(1)
                        .build(),
                    UserPoolAccountRecoverySettingRecoveryMechanismArgs.builder()
                        .name("verified_phone_number")
                        .priority(2)
                        .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:cognito:UserPool
    properties:
      accountRecoverySetting:
        recoveryMechanisms:
          - name: verified_email
            priority: 1
          - name: verified_phone_number
            priority: 2
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cognito User Pools using the `id`. For example:

```sh
 $ pulumi import aws:cognito/userPool:UserPool pool us-west-2_abc123
```
 