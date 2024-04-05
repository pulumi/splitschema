Provides information about a Signer Signing Job.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const buildSigningJob = aws.signer.getSigningJob({
    jobId: "9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee",
});
```
```python
import pulumi
import pulumi_aws as aws

build_signing_job = aws.signer.get_signing_job(job_id="9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var buildSigningJob = Aws.Signer.GetSigningJob.Invoke(new()
    {
        JobId = "9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/signer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := signer.LookupSigningJob(ctx, &signer.LookupSigningJobArgs{
			JobId: "9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee",
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
import com.pulumi.aws.signer.SignerFunctions;
import com.pulumi.aws.signer.inputs.GetSigningJobArgs;
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
        final var buildSigningJob = SignerFunctions.getSigningJob(GetSigningJobArgs.builder()
            .jobId("9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee")
            .build());

    }
}
```
```yaml
variables:
  buildSigningJob:
    fn::invoke:
      Function: aws:signer:getSigningJob
      Arguments:
        jobId: 9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee
```
{{% /example %}}
{{% /examples %}}