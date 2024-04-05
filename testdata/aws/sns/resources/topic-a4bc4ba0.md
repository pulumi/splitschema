Provides an SNS topic resource

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const userUpdates = new aws.sns.Topic("userUpdates", {});
```
```python
import pulumi
import pulumi_aws as aws

user_updates = aws.sns.Topic("userUpdates")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var userUpdates = new Aws.Sns.Topic("userUpdates");

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
		_, err := sns.NewTopic(ctx, "userUpdates", nil)
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
import com.pulumi.aws.sns.Topic;
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
        var userUpdates = new Topic("userUpdates");

    }
}
```
```yaml
resources:
  userUpdates:
    type: aws:sns:Topic
```
{{% /example %}}
{{% /examples %}}
## Example with Delivery Policy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const userUpdates = new aws.sns.Topic("userUpdates", {deliveryPolicy: `{
  "http": {
    "defaultHealthyRetryPolicy": {
      "minDelayTarget": 20,
      "maxDelayTarget": 20,
      "numRetries": 3,
      "numMaxDelayRetries": 0,
      "numNoDelayRetries": 0,
      "numMinDelayRetries": 0,
      "backoffFunction": "linear"
    },
    "disableSubscriptionOverrides": false,
    "defaultThrottlePolicy": {
      "maxReceivesPerSecond": 1
    }
  }
}

`});
```
```python
import pulumi
import pulumi_aws as aws

user_updates = aws.sns.Topic("userUpdates", delivery_policy="""{
  "http": {
    "defaultHealthyRetryPolicy": {
      "minDelayTarget": 20,
      "maxDelayTarget": 20,
      "numRetries": 3,
      "numMaxDelayRetries": 0,
      "numNoDelayRetries": 0,
      "numMinDelayRetries": 0,
      "backoffFunction": "linear"
    },
    "disableSubscriptionOverrides": false,
    "defaultThrottlePolicy": {
      "maxReceivesPerSecond": 1
    }
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
    var userUpdates = new Aws.Sns.Topic("userUpdates", new()
    {
        DeliveryPolicy = @"{
  ""http"": {
    ""defaultHealthyRetryPolicy"": {
      ""minDelayTarget"": 20,
      ""maxDelayTarget"": 20,
      ""numRetries"": 3,
      ""numMaxDelayRetries"": 0,
      ""numNoDelayRetries"": 0,
      ""numMinDelayRetries"": 0,
      ""backoffFunction"": ""linear""
    },
    ""disableSubscriptionOverrides"": false,
    ""defaultThrottlePolicy"": {
      ""maxReceivesPerSecond"": 1
    }
  }
}

",
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
		_, err := sns.NewTopic(ctx, "userUpdates", &sns.TopicArgs{
			DeliveryPolicy: pulumi.String(`{
  "http": {
    "defaultHealthyRetryPolicy": {
      "minDelayTarget": 20,
      "maxDelayTarget": 20,
      "numRetries": 3,
      "numMaxDelayRetries": 0,
      "numNoDelayRetries": 0,
      "numMinDelayRetries": 0,
      "backoffFunction": "linear"
    },
    "disableSubscriptionOverrides": false,
    "defaultThrottlePolicy": {
      "maxReceivesPerSecond": 1
    }
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
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.sns.TopicArgs;
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
        var userUpdates = new Topic("userUpdates", TopicArgs.builder()        
            .deliveryPolicy("""
{
  "http": {
    "defaultHealthyRetryPolicy": {
      "minDelayTarget": 20,
      "maxDelayTarget": 20,
      "numRetries": 3,
      "numMaxDelayRetries": 0,
      "numNoDelayRetries": 0,
      "numMinDelayRetries": 0,
      "backoffFunction": "linear"
    },
    "disableSubscriptionOverrides": false,
    "defaultThrottlePolicy": {
      "maxReceivesPerSecond": 1
    }
  }
}

            """)
            .build());

    }
}
```
```yaml
resources:
  userUpdates:
    type: aws:sns:Topic
    properties:
      deliveryPolicy: |+
        {
          "http": {
            "defaultHealthyRetryPolicy": {
              "minDelayTarget": 20,
              "maxDelayTarget": 20,
              "numRetries": 3,
              "numMaxDelayRetries": 0,
              "numNoDelayRetries": 0,
              "numMinDelayRetries": 0,
              "backoffFunction": "linear"
            },
            "disableSubscriptionOverrides": false,
            "defaultThrottlePolicy": {
              "maxReceivesPerSecond": 1
            }
          }
        }
```

## Example with Server-side encryption (SSE)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const userUpdates = new aws.sns.Topic("userUpdates", {kmsMasterKeyId: "alias/aws/sns"});
```
```python
import pulumi
import pulumi_aws as aws

user_updates = aws.sns.Topic("userUpdates", kms_master_key_id="alias/aws/sns")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var userUpdates = new Aws.Sns.Topic("userUpdates", new()
    {
        KmsMasterKeyId = "alias/aws/sns",
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
		_, err := sns.NewTopic(ctx, "userUpdates", &sns.TopicArgs{
			KmsMasterKeyId: pulumi.String("alias/aws/sns"),
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
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.sns.TopicArgs;
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
        var userUpdates = new Topic("userUpdates", TopicArgs.builder()        
            .kmsMasterKeyId("alias/aws/sns")
            .build());

    }
}
```
```yaml
resources:
  userUpdates:
    type: aws:sns:Topic
    properties:
      kmsMasterKeyId: alias/aws/sns
```

## Example with First-In-First-Out (FIFO)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const userUpdates = new aws.sns.Topic("userUpdates", {
    contentBasedDeduplication: true,
    fifoTopic: true,
});
```
```python
import pulumi
import pulumi_aws as aws

user_updates = aws.sns.Topic("userUpdates",
    content_based_deduplication=True,
    fifo_topic=True)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var userUpdates = new Aws.Sns.Topic("userUpdates", new()
    {
        ContentBasedDeduplication = true,
        FifoTopic = true,
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
		_, err := sns.NewTopic(ctx, "userUpdates", &sns.TopicArgs{
			ContentBasedDeduplication: pulumi.Bool(true),
			FifoTopic:                 pulumi.Bool(true),
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
import com.pulumi.aws.sns.Topic;
import com.pulumi.aws.sns.TopicArgs;
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
        var userUpdates = new Topic("userUpdates", TopicArgs.builder()        
            .contentBasedDeduplication(true)
            .fifoTopic(true)
            .build());

    }
}
```
```yaml
resources:
  userUpdates:
    type: aws:sns:Topic
    properties:
      contentBasedDeduplication: true
      fifoTopic: true
```

## Message Delivery Status Arguments

The `<endpoint>_success_feedback_role_arn` and `<endpoint>_failure_feedback_role_arn` arguments are used to give Amazon SNS write access to use CloudWatch Logs on your behalf. The `<endpoint>_success_feedback_sample_rate` argument is for specifying the sample rate percentage (0-100) of successfully delivered messages. After you configure the  `<endpoint>_failure_feedback_role_arn` argument, then all failed message deliveries generate CloudWatch Logs.


## Import

Using `pulumi import`, import SNS Topics using the topic `arn`. For example:

```sh
 $ pulumi import aws:sns/topic:Topic user_updates arn:aws:sns:us-west-2:0123456789012:my-topic
```
 