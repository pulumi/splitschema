Resource for managing SES Identity Notification Topics

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ses.IdentityNotificationTopic("test", {
    topicArn: aws_sns_topic.example.arn,
    notificationType: "Bounce",
    identity: aws_ses_domain_identity.example.domain,
    includeOriginalHeaders: true,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ses.IdentityNotificationTopic("test",
    topic_arn=aws_sns_topic["example"]["arn"],
    notification_type="Bounce",
    identity=aws_ses_domain_identity["example"]["domain"],
    include_original_headers=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ses.IdentityNotificationTopic("test", new()
    {
        TopicArn = aws_sns_topic.Example.Arn,
        NotificationType = "Bounce",
        Identity = aws_ses_domain_identity.Example.Domain,
        IncludeOriginalHeaders = true,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ses"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ses.NewIdentityNotificationTopic(ctx, "test", &ses.IdentityNotificationTopicArgs{
			TopicArn:               pulumi.Any(aws_sns_topic.Example.Arn),
			NotificationType:       pulumi.String("Bounce"),
			Identity:               pulumi.Any(aws_ses_domain_identity.Example.Domain),
			IncludeOriginalHeaders: pulumi.Bool(true),
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
import com.pulumi.aws.ses.IdentityNotificationTopic;
import com.pulumi.aws.ses.IdentityNotificationTopicArgs;
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
        var test = new IdentityNotificationTopic("test", IdentityNotificationTopicArgs.builder()        
            .topicArn(aws_sns_topic.example().arn())
            .notificationType("Bounce")
            .identity(aws_ses_domain_identity.example().domain())
            .includeOriginalHeaders(true)
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ses:IdentityNotificationTopic
    properties:
      topicArn: ${aws_sns_topic.example.arn}
      notificationType: Bounce
      identity: ${aws_ses_domain_identity.example.domain}
      includeOriginalHeaders: true
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Identity Notification Topics using the ID of the record. The ID is made up as `IDENTITY|TYPE` where `IDENTITY` is the SES Identity and `TYPE` is the Notification Type. For example:

```sh
 $ pulumi import aws:ses/identityNotificationTopic:IdentityNotificationTopic test 'example.com|Bounce'
```
 