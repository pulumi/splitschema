Provides a resource to manage AWS EMR Security Configurations

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foo = new aws.emr.SecurityConfiguration("foo", {configuration: `{
  "EncryptionConfiguration": {
    "AtRestEncryptionConfiguration": {
      "S3EncryptionConfiguration": {
        "EncryptionMode": "SSE-S3"
      },
      "LocalDiskEncryptionConfiguration": {
        "EncryptionKeyProviderType": "AwsKms",
        "AwsKmsKey": "arn:aws:kms:us-west-2:187416307283:alias/my_emr_test_key"
      }
    },
    "EnableInTransitEncryption": false,
    "EnableAtRestEncryption": true
  }
}

`});
```
```python
import pulumi
import pulumi_aws as aws

foo = aws.emr.SecurityConfiguration("foo", configuration="""{
  "EncryptionConfiguration": {
    "AtRestEncryptionConfiguration": {
      "S3EncryptionConfiguration": {
        "EncryptionMode": "SSE-S3"
      },
      "LocalDiskEncryptionConfiguration": {
        "EncryptionKeyProviderType": "AwsKms",
        "AwsKmsKey": "arn:aws:kms:us-west-2:187416307283:alias/my_emr_test_key"
      }
    },
    "EnableInTransitEncryption": false,
    "EnableAtRestEncryption": true
  }
}

""")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foo = new Aws.Emr.SecurityConfiguration("foo", new()
    {
        Configuration = @"{
  ""EncryptionConfiguration"": {
    ""AtRestEncryptionConfiguration"": {
      ""S3EncryptionConfiguration"": {
        ""EncryptionMode"": ""SSE-S3""
      },
      ""LocalDiskEncryptionConfiguration"": {
        ""EncryptionKeyProviderType"": ""AwsKms"",
        ""AwsKmsKey"": ""arn:aws:kms:us-west-2:187416307283:alias/my_emr_test_key""
      }
    },
    ""EnableInTransitEncryption"": false,
    ""EnableAtRestEncryption"": true
  }
}

",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/emr"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := emr.NewSecurityConfiguration(ctx, "foo", &emr.SecurityConfigurationArgs{
			Configuration: pulumi.String(`{
  "EncryptionConfiguration": {
    "AtRestEncryptionConfiguration": {
      "S3EncryptionConfiguration": {
        "EncryptionMode": "SSE-S3"
      },
      "LocalDiskEncryptionConfiguration": {
        "EncryptionKeyProviderType": "AwsKms",
        "AwsKmsKey": "arn:aws:kms:us-west-2:187416307283:alias/my_emr_test_key"
      }
    },
    "EnableInTransitEncryption": false,
    "EnableAtRestEncryption": true
  }
}

`),
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
import com.pulumi.aws.emr.SecurityConfiguration;
import com.pulumi.aws.emr.SecurityConfigurationArgs;
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
        var foo = new SecurityConfiguration("foo", SecurityConfigurationArgs.builder()        
            .configuration("""
{
  "EncryptionConfiguration": {
    "AtRestEncryptionConfiguration": {
      "S3EncryptionConfiguration": {
        "EncryptionMode": "SSE-S3"
      },
      "LocalDiskEncryptionConfiguration": {
        "EncryptionKeyProviderType": "AwsKms",
        "AwsKmsKey": "arn:aws:kms:us-west-2:187416307283:alias/my_emr_test_key"
      }
    },
    "EnableInTransitEncryption": false,
    "EnableAtRestEncryption": true
  }
}

            """)
            .build());

    }
}
```
```yaml
resources:
  foo:
    type: aws:emr:SecurityConfiguration
    properties:
      configuration: |+
        {
          "EncryptionConfiguration": {
            "AtRestEncryptionConfiguration": {
              "S3EncryptionConfiguration": {
                "EncryptionMode": "SSE-S3"
              },
              "LocalDiskEncryptionConfiguration": {
                "EncryptionKeyProviderType": "AwsKms",
                "AwsKmsKey": "arn:aws:kms:us-west-2:187416307283:alias/my_emr_test_key"
              }
            },
            "EnableInTransitEncryption": false,
            "EnableAtRestEncryption": true
          }
        }
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EMR Security Configurations using the `name`. For example:

```sh
 $ pulumi import aws:emr/securityConfiguration:SecurityConfiguration sc example-sc-name
```
 