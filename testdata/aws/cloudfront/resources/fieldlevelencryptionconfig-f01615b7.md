Provides a CloudFront Field-level Encryption Config resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.cloudfront.FieldLevelEncryptionConfig("test", {
    comment: "test comment",
    contentTypeProfileConfig: {
        forwardWhenContentTypeIsUnknown: true,
        contentTypeProfiles: {
            items: [{
                contentType: "application/x-www-form-urlencoded",
                format: "URLEncoded",
            }],
        },
    },
    queryArgProfileConfig: {
        forwardWhenQueryArgProfileIsUnknown: true,
        queryArgProfiles: {
            items: [{
                profileId: aws_cloudfront_field_level_encryption_profile.test.id,
                queryArg: "Arg1",
            }],
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.cloudfront.FieldLevelEncryptionConfig("test",
    comment="test comment",
    content_type_profile_config=aws.cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigArgs(
        forward_when_content_type_is_unknown=True,
        content_type_profiles=aws.cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesArgs(
            items=[aws.cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesItemArgs(
                content_type="application/x-www-form-urlencoded",
                format="URLEncoded",
            )],
        ),
    ),
    query_arg_profile_config=aws.cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigArgs(
        forward_when_query_arg_profile_is_unknown=True,
        query_arg_profiles=aws.cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesArgs(
            items=[aws.cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesItemArgs(
                profile_id=aws_cloudfront_field_level_encryption_profile["test"]["id"],
                query_arg="Arg1",
            )],
        ),
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.CloudFront.FieldLevelEncryptionConfig("test", new()
    {
        Comment = "test comment",
        ContentTypeProfileConfig = new Aws.CloudFront.Inputs.FieldLevelEncryptionConfigContentTypeProfileConfigArgs
        {
            ForwardWhenContentTypeIsUnknown = true,
            ContentTypeProfiles = new Aws.CloudFront.Inputs.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesArgs
            {
                Items = new[]
                {
                    new Aws.CloudFront.Inputs.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesItemArgs
                    {
                        ContentType = "application/x-www-form-urlencoded",
                        Format = "URLEncoded",
                    },
                },
            },
        },
        QueryArgProfileConfig = new Aws.CloudFront.Inputs.FieldLevelEncryptionConfigQueryArgProfileConfigArgs
        {
            ForwardWhenQueryArgProfileIsUnknown = true,
            QueryArgProfiles = new Aws.CloudFront.Inputs.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesArgs
            {
                Items = new[]
                {
                    new Aws.CloudFront.Inputs.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesItemArgs
                    {
                        ProfileId = aws_cloudfront_field_level_encryption_profile.Test.Id,
                        QueryArg = "Arg1",
                    },
                },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudfront"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudfront.NewFieldLevelEncryptionConfig(ctx, "test", &cloudfront.FieldLevelEncryptionConfigArgs{
			Comment: pulumi.String("test comment"),
			ContentTypeProfileConfig: &cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigArgs{
				ForwardWhenContentTypeIsUnknown: pulumi.Bool(true),
				ContentTypeProfiles: &cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesArgs{
					Items: cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesItemArray{
						&cloudfront.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesItemArgs{
							ContentType: pulumi.String("application/x-www-form-urlencoded"),
							Format:      pulumi.String("URLEncoded"),
						},
					},
				},
			},
			QueryArgProfileConfig: &cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigArgs{
				ForwardWhenQueryArgProfileIsUnknown: pulumi.Bool(true),
				QueryArgProfiles: &cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesArgs{
					Items: cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesItemArray{
						&cloudfront.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesItemArgs{
							ProfileId: pulumi.Any(aws_cloudfront_field_level_encryption_profile.Test.Id),
							QueryArg:  pulumi.String("Arg1"),
						},
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
import com.pulumi.aws.cloudfront.FieldLevelEncryptionConfig;
import com.pulumi.aws.cloudfront.FieldLevelEncryptionConfigArgs;
import com.pulumi.aws.cloudfront.inputs.FieldLevelEncryptionConfigContentTypeProfileConfigArgs;
import com.pulumi.aws.cloudfront.inputs.FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesArgs;
import com.pulumi.aws.cloudfront.inputs.FieldLevelEncryptionConfigQueryArgProfileConfigArgs;
import com.pulumi.aws.cloudfront.inputs.FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesArgs;
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
        var test = new FieldLevelEncryptionConfig("test", FieldLevelEncryptionConfigArgs.builder()        
            .comment("test comment")
            .contentTypeProfileConfig(FieldLevelEncryptionConfigContentTypeProfileConfigArgs.builder()
                .forwardWhenContentTypeIsUnknown(true)
                .contentTypeProfiles(FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesArgs.builder()
                    .items(FieldLevelEncryptionConfigContentTypeProfileConfigContentTypeProfilesItemArgs.builder()
                        .contentType("application/x-www-form-urlencoded")
                        .format("URLEncoded")
                        .build())
                    .build())
                .build())
            .queryArgProfileConfig(FieldLevelEncryptionConfigQueryArgProfileConfigArgs.builder()
                .forwardWhenQueryArgProfileIsUnknown(true)
                .queryArgProfiles(FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesArgs.builder()
                    .items(FieldLevelEncryptionConfigQueryArgProfileConfigQueryArgProfilesItemArgs.builder()
                        .profileId(aws_cloudfront_field_level_encryption_profile.test().id())
                        .queryArg("Arg1")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:cloudfront:FieldLevelEncryptionConfig
    properties:
      comment: test comment
      contentTypeProfileConfig:
        forwardWhenContentTypeIsUnknown: true
        contentTypeProfiles:
          items:
            - contentType: application/x-www-form-urlencoded
              format: URLEncoded
      queryArgProfileConfig:
        forwardWhenQueryArgProfileIsUnknown: true
        queryArgProfiles:
          items:
            - profileId: ${aws_cloudfront_field_level_encryption_profile.test.id}
              queryArg: Arg1
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Cloudfront Field Level Encryption Config using the `id`. For example:

```sh
 $ pulumi import aws:cloudfront/fieldLevelEncryptionConfig:FieldLevelEncryptionConfig config E74FTE3AEXAMPLE
```
 