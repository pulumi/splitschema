Provides an EventBridge connection resource.

> **Note:** EventBridge was formerly known as CloudWatch Events. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.cloudwatch.EventConnection("test", {
    authParameters: {
        apiKey: {
            key: "x-signature",
            value: "1234",
        },
    },
    authorizationType: "API_KEY",
    description: "A connection description",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.cloudwatch.EventConnection("test",
    auth_parameters=aws.cloudwatch.EventConnectionAuthParametersArgs(
        api_key=aws.cloudwatch.EventConnectionAuthParametersApiKeyArgs(
            key="x-signature",
            value="1234",
        ),
    ),
    authorization_type="API_KEY",
    description="A connection description")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.CloudWatch.EventConnection("test", new()
    {
        AuthParameters = new Aws.CloudWatch.Inputs.EventConnectionAuthParametersArgs
        {
            ApiKey = new Aws.CloudWatch.Inputs.EventConnectionAuthParametersApiKeyArgs
            {
                Key = "x-signature",
                Value = "1234",
            },
        },
        AuthorizationType = "API_KEY",
        Description = "A connection description",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewEventConnection(ctx, "test", &cloudwatch.EventConnectionArgs{
			AuthParameters: &cloudwatch.EventConnectionAuthParametersArgs{
				ApiKey: &cloudwatch.EventConnectionAuthParametersApiKeyArgs{
					Key:   pulumi.String("x-signature"),
					Value: pulumi.String("1234"),
				},
			},
			AuthorizationType: pulumi.String("API_KEY"),
			Description:       pulumi.String("A connection description"),
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
import com.pulumi.aws.cloudwatch.EventConnection;
import com.pulumi.aws.cloudwatch.EventConnectionArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersApiKeyArgs;
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
        var test = new EventConnection("test", EventConnectionArgs.builder()        
            .authParameters(EventConnectionAuthParametersArgs.builder()
                .apiKey(EventConnectionAuthParametersApiKeyArgs.builder()
                    .key("x-signature")
                    .value("1234")
                    .build())
                .build())
            .authorizationType("API_KEY")
            .description("A connection description")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:cloudwatch:EventConnection
    properties:
      authParameters:
        apiKey:
          key: x-signature
          value: '1234'
      authorizationType: API_KEY
      description: A connection description
```

{{% /example %}}
{{% example %}}
### Basic Authorization

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.cloudwatch.EventConnection("test", {
    authParameters: {
        basic: {
            password: "Pass1234!",
            username: "user",
        },
    },
    authorizationType: "BASIC",
    description: "A connection description",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.cloudwatch.EventConnection("test",
    auth_parameters=aws.cloudwatch.EventConnectionAuthParametersArgs(
        basic=aws.cloudwatch.EventConnectionAuthParametersBasicArgs(
            password="Pass1234!",
            username="user",
        ),
    ),
    authorization_type="BASIC",
    description="A connection description")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.CloudWatch.EventConnection("test", new()
    {
        AuthParameters = new Aws.CloudWatch.Inputs.EventConnectionAuthParametersArgs
        {
            Basic = new Aws.CloudWatch.Inputs.EventConnectionAuthParametersBasicArgs
            {
                Password = "Pass1234!",
                Username = "user",
            },
        },
        AuthorizationType = "BASIC",
        Description = "A connection description",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudwatch.NewEventConnection(ctx, "test", &cloudwatch.EventConnectionArgs{
			AuthParameters: &cloudwatch.EventConnectionAuthParametersArgs{
				Basic: &cloudwatch.EventConnectionAuthParametersBasicArgs{
					Password: pulumi.String("Pass1234!"),
					Username: pulumi.String("user"),
				},
			},
			AuthorizationType: pulumi.String("BASIC"),
			Description:       pulumi.String("A connection description"),
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
import com.pulumi.aws.cloudwatch.EventConnection;
import com.pulumi.aws.cloudwatch.EventConnectionArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersBasicArgs;
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
        var test = new EventConnection("test", EventConnectionArgs.builder()        
            .authParameters(EventConnectionAuthParametersArgs.builder()
                .basic(EventConnectionAuthParametersBasicArgs.builder()
                    .password("Pass1234!")
                    .username("user")
                    .build())
                .build())
            .authorizationType("BASIC")
            .description("A connection description")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:cloudwatch:EventConnection
    properties:
      authParameters:
        basic:
          password: Pass1234!
          username: user
      authorizationType: BASIC
      description: A connection description
```

{{% /example %}}
{{% example %}}
### OAuth Authorization

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.cloudwatch.EventConnection;
import com.pulumi.aws.cloudwatch.EventConnectionArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersOauthArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersOauthClientParametersArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersOauthOauthHttpParametersArgs;
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
        var test = new EventConnection("test", EventConnectionArgs.builder()        
            .authParameters(EventConnectionAuthParametersArgs.builder()
                .oauth(EventConnectionAuthParametersOauthArgs.builder()
                    .authorizationEndpoint("https://auth.url.com/endpoint")
                    .clientParameters(EventConnectionAuthParametersOauthClientParametersArgs.builder()
                        .clientId("1234567890")
                        .clientSecret("Pass1234!")
                        .build())
                    .httpMethod("GET")
                    .oauthHttpParameters(EventConnectionAuthParametersOauthOauthHttpParametersArgs.builder()
                        .body(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                        .header(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                        .queryString(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                        .build())
                    .build())
                .build())
            .authorizationType("OAUTH_CLIENT_CREDENTIALS")
            .description("A connection description")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:cloudwatch:EventConnection
    properties:
      authParameters:
        oauth:
          authorizationEndpoint: https://auth.url.com/endpoint
          clientParameters:
            clientId: '1234567890'
            clientSecret: Pass1234!
          httpMethod: GET
          oauthHttpParameters:
            body:
              - isValueSecret: false
                key: body-parameter-key
                value: body-parameter-value
            header:
              - isValueSecret: false
                key: header-parameter-key
                value: header-parameter-value
            queryString:
              - isValueSecret: false
                key: query-string-parameter-key
                value: query-string-parameter-value
      authorizationType: OAUTH_CLIENT_CREDENTIALS
      description: A connection description
```

{{% /example %}}
{{% example %}}
### Invocation Http Parameters

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.cloudwatch.EventConnection;
import com.pulumi.aws.cloudwatch.EventConnectionArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersBasicArgs;
import com.pulumi.aws.cloudwatch.inputs.EventConnectionAuthParametersInvocationHttpParametersArgs;
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
        var test = new EventConnection("test", EventConnectionArgs.builder()        
            .authParameters(EventConnectionAuthParametersArgs.builder()
                .basic(EventConnectionAuthParametersBasicArgs.builder()
                    .password("Pass1234!")
                    .username("user")
                    .build())
                .invocationHttpParameters(EventConnectionAuthParametersInvocationHttpParametersArgs.builder()
                    .body(                    
                        %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference),
                        %!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                    .header(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                    .queryString(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
                    .build())
                .build())
            .authorizationType("BASIC")
            .description("A connection description")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:cloudwatch:EventConnection
    properties:
      authParameters:
        basic:
          password: Pass1234!
          username: user
        invocationHttpParameters:
          body:
            - isValueSecret: false
              key: body-parameter-key
              value: body-parameter-value
            - isValueSecret: true
              key: body-parameter-key2
              value: body-parameter-value2
          header:
            - isValueSecret: false
              key: header-parameter-key
              value: header-parameter-value
          queryString:
            - isValueSecret: false
              key: query-string-parameter-key
              value: query-string-parameter-value
      authorizationType: BASIC
      description: A connection description
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import EventBridge EventBridge connection using the `name`. For example:

```sh
 $ pulumi import aws:cloudwatch/eventConnection:EventConnection test ngrok-connection
```
 