Use this data source to get information about a SSH public key associated with the specified IAM user.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = aws.iam.getUserSshKey({
    encoding: "SSH",
    sshPublicKeyId: "APKARUZ32GUTKIGARLXE",
    username: "test-user",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.iam.get_user_ssh_key(encoding="SSH",
    ssh_public_key_id="APKARUZ32GUTKIGARLXE",
    username="test-user")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Iam.GetUserSshKey.Invoke(new()
    {
        Encoding = "SSH",
        SshPublicKeyId = "APKARUZ32GUTKIGARLXE",
        Username = "test-user",
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
		_, err := iam.GetUserSshKey(ctx, &iam.GetUserSshKeyArgs{
			Encoding:       "SSH",
			SshPublicKeyId: "APKARUZ32GUTKIGARLXE",
			Username:       "test-user",
		}, nil)
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetUserSshKeyArgs;
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
        final var example = IamFunctions.getUserSshKey(GetUserSshKeyArgs.builder()
            .encoding("SSH")
            .sshPublicKeyId("APKARUZ32GUTKIGARLXE")
            .username("test-user")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:iam:getUserSshKey
      Arguments:
        encoding: SSH
        sshPublicKeyId: APKARUZ32GUTKIGARLXE
        username: test-user
```
{{% /example %}}
{{% /examples %}}