Creates a Signer Signing Job.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const testSp = new aws.signer.SigningProfile("testSp", {platformId: "AWSLambda-SHA384-ECDSA"});
const buildSigningJob = new aws.signer.SigningJob("buildSigningJob", {
    profileName: testSp.name,
    source: {
        s3: {
            bucket: "s3-bucket-name",
            key: "object-to-be-signed.zip",
            version: "jADjFYYYEXAMPLETszPjOmCMFDzd9dN1",
        },
    },
    destination: {
        s3: {
            bucket: "s3-bucket-name",
            prefix: "signed/",
        },
    },
    ignoreSigningJobFailure: true,
});
```
```python
import pulumi
import pulumi_aws as aws

test_sp = aws.signer.SigningProfile("testSp", platform_id="AWSLambda-SHA384-ECDSA")
build_signing_job = aws.signer.SigningJob("buildSigningJob",
    profile_name=test_sp.name,
    source=aws.signer.SigningJobSourceArgs(
        s3=aws.signer.SigningJobSourceS3Args(
            bucket="s3-bucket-name",
            key="object-to-be-signed.zip",
            version="jADjFYYYEXAMPLETszPjOmCMFDzd9dN1",
        ),
    ),
    destination=aws.signer.SigningJobDestinationArgs(
        s3=aws.signer.SigningJobDestinationS3Args(
            bucket="s3-bucket-name",
            prefix="signed/",
        ),
    ),
    ignore_signing_job_failure=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var testSp = new Aws.Signer.SigningProfile("testSp", new()
    {
        PlatformId = "AWSLambda-SHA384-ECDSA",
    });

    var buildSigningJob = new Aws.Signer.SigningJob("buildSigningJob", new()
    {
        ProfileName = testSp.Name,
        Source = new Aws.Signer.Inputs.SigningJobSourceArgs
        {
            S3 = new Aws.Signer.Inputs.SigningJobSourceS3Args
            {
                Bucket = "s3-bucket-name",
                Key = "object-to-be-signed.zip",
                Version = "jADjFYYYEXAMPLETszPjOmCMFDzd9dN1",
            },
        },
        Destination = new Aws.Signer.Inputs.SigningJobDestinationArgs
        {
            S3 = new Aws.Signer.Inputs.SigningJobDestinationS3Args
            {
                Bucket = "s3-bucket-name",
                Prefix = "signed/",
            },
        },
        IgnoreSigningJobFailure = true,
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
		testSp, err := signer.NewSigningProfile(ctx, "testSp", &signer.SigningProfileArgs{
			PlatformId: pulumi.String("AWSLambda-SHA384-ECDSA"),
		})
		if err != nil {
			return err
		}
		_, err = signer.NewSigningJob(ctx, "buildSigningJob", &signer.SigningJobArgs{
			ProfileName: testSp.Name,
			Source: &signer.SigningJobSourceArgs{
				S3: &signer.SigningJobSourceS3Args{
					Bucket:  pulumi.String("s3-bucket-name"),
					Key:     pulumi.String("object-to-be-signed.zip"),
					Version: pulumi.String("jADjFYYYEXAMPLETszPjOmCMFDzd9dN1"),
				},
			},
			Destination: &signer.SigningJobDestinationArgs{
				S3: &signer.SigningJobDestinationS3Args{
					Bucket: pulumi.String("s3-bucket-name"),
					Prefix: pulumi.String("signed/"),
				},
			},
			IgnoreSigningJobFailure: pulumi.Bool(true),
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
import com.pulumi.aws.signer.SigningProfile;
import com.pulumi.aws.signer.SigningProfileArgs;
import com.pulumi.aws.signer.SigningJob;
import com.pulumi.aws.signer.SigningJobArgs;
import com.pulumi.aws.signer.inputs.SigningJobSourceArgs;
import com.pulumi.aws.signer.inputs.SigningJobSourceS3Args;
import com.pulumi.aws.signer.inputs.SigningJobDestinationArgs;
import com.pulumi.aws.signer.inputs.SigningJobDestinationS3Args;
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
        var testSp = new SigningProfile("testSp", SigningProfileArgs.builder()        
            .platformId("AWSLambda-SHA384-ECDSA")
            .build());

        var buildSigningJob = new SigningJob("buildSigningJob", SigningJobArgs.builder()        
            .profileName(testSp.name())
            .source(SigningJobSourceArgs.builder()
                .s3(SigningJobSourceS3Args.builder()
                    .bucket("s3-bucket-name")
                    .key("object-to-be-signed.zip")
                    .version("jADjFYYYEXAMPLETszPjOmCMFDzd9dN1")
                    .build())
                .build())
            .destination(SigningJobDestinationArgs.builder()
                .s3(SigningJobDestinationS3Args.builder()
                    .bucket("s3-bucket-name")
                    .prefix("signed/")
                    .build())
                .build())
            .ignoreSigningJobFailure(true)
            .build());

    }
}
```
```yaml
resources:
  testSp:
    type: aws:signer:SigningProfile
    properties:
      platformId: AWSLambda-SHA384-ECDSA
  buildSigningJob:
    type: aws:signer:SigningJob
    properties:
      profileName: ${testSp.name}
      source:
        s3:
          bucket: s3-bucket-name
          key: object-to-be-signed.zip
          version: jADjFYYYEXAMPLETszPjOmCMFDzd9dN1
      destination:
        s3:
          bucket: s3-bucket-name
          prefix: signed/
      ignoreSigningJobFailure: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Signer signing jobs using the `job_id`. For example:

```sh
 $ pulumi import aws:signer/signingJob:SigningJob test_signer_signing_job 9ed7e5c3-b8d4-4da0-8459-44e0b068f7ee
```
 