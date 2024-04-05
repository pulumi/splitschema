Provides a Load Balancer Listener Rule resource.

> **Note:** `aws.alb.ListenerRule` is known as `aws.lb.ListenerRule`. The functionality is identical.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const frontEndLoadBalancer = new aws.lb.LoadBalancer("frontEndLoadBalancer", {});
// ...
const frontEndListener = new aws.lb.Listener("frontEndListener", {});
// Other parameters
const static = new aws.lb.ListenerRule("static", {
    listenerArn: frontEndListener.arn,
    priority: 100,
    actions: [{
        type: "forward",
        targetGroupArn: aws_lb_target_group.static.arn,
    }],
    conditions: [
        {
            pathPattern: {
                values: ["/static/*"],
            },
        },
        {
            hostHeader: {
                values: ["example.com"],
            },
        },
    ],
});
// Forward action
const hostBasedWeightedRouting = new aws.lb.ListenerRule("hostBasedWeightedRouting", {
    listenerArn: frontEndListener.arn,
    priority: 99,
    actions: [{
        type: "forward",
        targetGroupArn: aws_lb_target_group.static.arn,
    }],
    conditions: [{
        hostHeader: {
            values: ["my-service.*.mycompany.io"],
        },
    }],
});
// Weighted Forward action
const hostBasedRouting = new aws.lb.ListenerRule("hostBasedRouting", {
    listenerArn: frontEndListener.arn,
    priority: 99,
    actions: [{
        type: "forward",
        forward: {
            targetGroups: [
                {
                    arn: aws_lb_target_group.main.arn,
                    weight: 80,
                },
                {
                    arn: aws_lb_target_group.canary.arn,
                    weight: 20,
                },
            ],
            stickiness: {
                enabled: true,
                duration: 600,
            },
        },
    }],
    conditions: [{
        hostHeader: {
            values: ["my-service.*.mycompany.io"],
        },
    }],
});
// Redirect action
const redirectHttpToHttps = new aws.lb.ListenerRule("redirectHttpToHttps", {
    listenerArn: frontEndListener.arn,
    actions: [{
        type: "redirect",
        redirect: {
            port: "443",
            protocol: "HTTPS",
            statusCode: "HTTP_301",
        },
    }],
    conditions: [{
        httpHeader: {
            httpHeaderName: "X-Forwarded-For",
            values: ["192.168.1.*"],
        },
    }],
});
// Fixed-response action
const healthCheck = new aws.lb.ListenerRule("healthCheck", {
    listenerArn: frontEndListener.arn,
    actions: [{
        type: "fixed-response",
        fixedResponse: {
            contentType: "text/plain",
            messageBody: "HEALTHY",
            statusCode: "200",
        },
    }],
    conditions: [{
        queryStrings: [
            {
                key: "health",
                value: "check",
            },
            {
                value: "bar",
            },
        ],
    }],
});
// Authenticate-cognito Action
const pool = new aws.cognito.UserPool("pool", {});
// ...
const client = new aws.cognito.UserPoolClient("client", {});
// ...
const domain = new aws.cognito.UserPoolDomain("domain", {});
// ...
const admin = new aws.lb.ListenerRule("admin", {
    listenerArn: frontEndListener.arn,
    actions: [
        {
            type: "authenticate-cognito",
            authenticateCognito: {
                userPoolArn: pool.arn,
                userPoolClientId: client.id,
                userPoolDomain: domain.domain,
            },
        },
        {
            type: "forward",
            targetGroupArn: aws_lb_target_group.static.arn,
        },
    ],
});
// Authenticate-oidc Action
const oidc = new aws.lb.ListenerRule("oidc", {
    listenerArn: frontEndListener.arn,
    actions: [
        {
            type: "authenticate-oidc",
            authenticateOidc: {
                authorizationEndpoint: "https://example.com/authorization_endpoint",
                clientId: "client_id",
                clientSecret: "client_secret",
                issuer: "https://example.com",
                tokenEndpoint: "https://example.com/token_endpoint",
                userInfoEndpoint: "https://example.com/user_info_endpoint",
            },
        },
        {
            type: "forward",
            targetGroupArn: aws_lb_target_group.static.arn,
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

front_end_load_balancer = aws.lb.LoadBalancer("frontEndLoadBalancer")
# ...
front_end_listener = aws.lb.Listener("frontEndListener")
# Other parameters
static = aws.lb.ListenerRule("static",
    listener_arn=front_end_listener.arn,
    priority=100,
    actions=[aws.lb.ListenerRuleActionArgs(
        type="forward",
        target_group_arn=aws_lb_target_group["static"]["arn"],
    )],
    conditions=[
        aws.lb.ListenerRuleConditionArgs(
            path_pattern=aws.lb.ListenerRuleConditionPathPatternArgs(
                values=["/static/*"],
            ),
        ),
        aws.lb.ListenerRuleConditionArgs(
            host_header=aws.lb.ListenerRuleConditionHostHeaderArgs(
                values=["example.com"],
            ),
        ),
    ])
# Forward action
host_based_weighted_routing = aws.lb.ListenerRule("hostBasedWeightedRouting",
    listener_arn=front_end_listener.arn,
    priority=99,
    actions=[aws.lb.ListenerRuleActionArgs(
        type="forward",
        target_group_arn=aws_lb_target_group["static"]["arn"],
    )],
    conditions=[aws.lb.ListenerRuleConditionArgs(
        host_header=aws.lb.ListenerRuleConditionHostHeaderArgs(
            values=["my-service.*.mycompany.io"],
        ),
    )])
# Weighted Forward action
host_based_routing = aws.lb.ListenerRule("hostBasedRouting",
    listener_arn=front_end_listener.arn,
    priority=99,
    actions=[aws.lb.ListenerRuleActionArgs(
        type="forward",
        forward=aws.lb.ListenerRuleActionForwardArgs(
            target_groups=[
                aws.lb.ListenerRuleActionForwardTargetGroupArgs(
                    arn=aws_lb_target_group["main"]["arn"],
                    weight=80,
                ),
                aws.lb.ListenerRuleActionForwardTargetGroupArgs(
                    arn=aws_lb_target_group["canary"]["arn"],
                    weight=20,
                ),
            ],
            stickiness=aws.lb.ListenerRuleActionForwardStickinessArgs(
                enabled=True,
                duration=600,
            ),
        ),
    )],
    conditions=[aws.lb.ListenerRuleConditionArgs(
        host_header=aws.lb.ListenerRuleConditionHostHeaderArgs(
            values=["my-service.*.mycompany.io"],
        ),
    )])
# Redirect action
redirect_http_to_https = aws.lb.ListenerRule("redirectHttpToHttps",
    listener_arn=front_end_listener.arn,
    actions=[aws.lb.ListenerRuleActionArgs(
        type="redirect",
        redirect=aws.lb.ListenerRuleActionRedirectArgs(
            port="443",
            protocol="HTTPS",
            status_code="HTTP_301",
        ),
    )],
    conditions=[aws.lb.ListenerRuleConditionArgs(
        http_header=aws.lb.ListenerRuleConditionHttpHeaderArgs(
            http_header_name="X-Forwarded-For",
            values=["192.168.1.*"],
        ),
    )])
# Fixed-response action
health_check = aws.lb.ListenerRule("healthCheck",
    listener_arn=front_end_listener.arn,
    actions=[aws.lb.ListenerRuleActionArgs(
        type="fixed-response",
        fixed_response=aws.lb.ListenerRuleActionFixedResponseArgs(
            content_type="text/plain",
            message_body="HEALTHY",
            status_code="200",
        ),
    )],
    conditions=[aws.lb.ListenerRuleConditionArgs(
        query_strings=[
            aws.lb.ListenerRuleConditionQueryStringArgs(
                key="health",
                value="check",
            ),
            aws.lb.ListenerRuleConditionQueryStringArgs(
                value="bar",
            ),
        ],
    )])
# Authenticate-cognito Action
pool = aws.cognito.UserPool("pool")
# ...
client = aws.cognito.UserPoolClient("client")
# ...
domain = aws.cognito.UserPoolDomain("domain")
# ...
admin = aws.lb.ListenerRule("admin",
    listener_arn=front_end_listener.arn,
    actions=[
        aws.lb.ListenerRuleActionArgs(
            type="authenticate-cognito",
            authenticate_cognito=aws.lb.ListenerRuleActionAuthenticateCognitoArgs(
                user_pool_arn=pool.arn,
                user_pool_client_id=client.id,
                user_pool_domain=domain.domain,
            ),
        ),
        aws.lb.ListenerRuleActionArgs(
            type="forward",
            target_group_arn=aws_lb_target_group["static"]["arn"],
        ),
    ])
# Authenticate-oidc Action
oidc = aws.lb.ListenerRule("oidc",
    listener_arn=front_end_listener.arn,
    actions=[
        aws.lb.ListenerRuleActionArgs(
            type="authenticate-oidc",
            authenticate_oidc=aws.lb.ListenerRuleActionAuthenticateOidcArgs(
                authorization_endpoint="https://example.com/authorization_endpoint",
                client_id="client_id",
                client_secret="client_secret",
                issuer="https://example.com",
                token_endpoint="https://example.com/token_endpoint",
                user_info_endpoint="https://example.com/user_info_endpoint",
            ),
        ),
        aws.lb.ListenerRuleActionArgs(
            type="forward",
            target_group_arn=aws_lb_target_group["static"]["arn"],
        ),
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var frontEndLoadBalancer = new Aws.LB.LoadBalancer("frontEndLoadBalancer");

    // ...
    var frontEndListener = new Aws.LB.Listener("frontEndListener");

    // Other parameters
    var @static = new Aws.LB.ListenerRule("static", new()
    {
        ListenerArn = frontEndListener.Arn,
        Priority = 100,
        Actions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleActionArgs
            {
                Type = "forward",
                TargetGroupArn = aws_lb_target_group.Static.Arn,
            },
        },
        Conditions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleConditionArgs
            {
                PathPattern = new Aws.LB.Inputs.ListenerRuleConditionPathPatternArgs
                {
                    Values = new[]
                    {
                        "/static/*",
                    },
                },
            },
            new Aws.LB.Inputs.ListenerRuleConditionArgs
            {
                HostHeader = new Aws.LB.Inputs.ListenerRuleConditionHostHeaderArgs
                {
                    Values = new[]
                    {
                        "example.com",
                    },
                },
            },
        },
    });

    // Forward action
    var hostBasedWeightedRouting = new Aws.LB.ListenerRule("hostBasedWeightedRouting", new()
    {
        ListenerArn = frontEndListener.Arn,
        Priority = 99,
        Actions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleActionArgs
            {
                Type = "forward",
                TargetGroupArn = aws_lb_target_group.Static.Arn,
            },
        },
        Conditions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleConditionArgs
            {
                HostHeader = new Aws.LB.Inputs.ListenerRuleConditionHostHeaderArgs
                {
                    Values = new[]
                    {
                        "my-service.*.mycompany.io",
                    },
                },
            },
        },
    });

    // Weighted Forward action
    var hostBasedRouting = new Aws.LB.ListenerRule("hostBasedRouting", new()
    {
        ListenerArn = frontEndListener.Arn,
        Priority = 99,
        Actions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleActionArgs
            {
                Type = "forward",
                Forward = new Aws.LB.Inputs.ListenerRuleActionForwardArgs
                {
                    TargetGroups = new[]
                    {
                        new Aws.LB.Inputs.ListenerRuleActionForwardTargetGroupArgs
                        {
                            Arn = aws_lb_target_group.Main.Arn,
                            Weight = 80,
                        },
                        new Aws.LB.Inputs.ListenerRuleActionForwardTargetGroupArgs
                        {
                            Arn = aws_lb_target_group.Canary.Arn,
                            Weight = 20,
                        },
                    },
                    Stickiness = new Aws.LB.Inputs.ListenerRuleActionForwardStickinessArgs
                    {
                        Enabled = true,
                        Duration = 600,
                    },
                },
            },
        },
        Conditions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleConditionArgs
            {
                HostHeader = new Aws.LB.Inputs.ListenerRuleConditionHostHeaderArgs
                {
                    Values = new[]
                    {
                        "my-service.*.mycompany.io",
                    },
                },
            },
        },
    });

    // Redirect action
    var redirectHttpToHttps = new Aws.LB.ListenerRule("redirectHttpToHttps", new()
    {
        ListenerArn = frontEndListener.Arn,
        Actions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleActionArgs
            {
                Type = "redirect",
                Redirect = new Aws.LB.Inputs.ListenerRuleActionRedirectArgs
                {
                    Port = "443",
                    Protocol = "HTTPS",
                    StatusCode = "HTTP_301",
                },
            },
        },
        Conditions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleConditionArgs
            {
                HttpHeader = new Aws.LB.Inputs.ListenerRuleConditionHttpHeaderArgs
                {
                    HttpHeaderName = "X-Forwarded-For",
                    Values = new[]
                    {
                        "192.168.1.*",
                    },
                },
            },
        },
    });

    // Fixed-response action
    var healthCheck = new Aws.LB.ListenerRule("healthCheck", new()
    {
        ListenerArn = frontEndListener.Arn,
        Actions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleActionArgs
            {
                Type = "fixed-response",
                FixedResponse = new Aws.LB.Inputs.ListenerRuleActionFixedResponseArgs
                {
                    ContentType = "text/plain",
                    MessageBody = "HEALTHY",
                    StatusCode = "200",
                },
            },
        },
        Conditions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleConditionArgs
            {
                QueryStrings = new[]
                {
                    new Aws.LB.Inputs.ListenerRuleConditionQueryStringArgs
                    {
                        Key = "health",
                        Value = "check",
                    },
                    new Aws.LB.Inputs.ListenerRuleConditionQueryStringArgs
                    {
                        Value = "bar",
                    },
                },
            },
        },
    });

    // Authenticate-cognito Action
    var pool = new Aws.Cognito.UserPool("pool");

    // ...
    var client = new Aws.Cognito.UserPoolClient("client");

    // ...
    var domain = new Aws.Cognito.UserPoolDomain("domain");

    // ...
    var admin = new Aws.LB.ListenerRule("admin", new()
    {
        ListenerArn = frontEndListener.Arn,
        Actions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleActionArgs
            {
                Type = "authenticate-cognito",
                AuthenticateCognito = new Aws.LB.Inputs.ListenerRuleActionAuthenticateCognitoArgs
                {
                    UserPoolArn = pool.Arn,
                    UserPoolClientId = client.Id,
                    UserPoolDomain = domain.Domain,
                },
            },
            new Aws.LB.Inputs.ListenerRuleActionArgs
            {
                Type = "forward",
                TargetGroupArn = aws_lb_target_group.Static.Arn,
            },
        },
    });

    // Authenticate-oidc Action
    var oidc = new Aws.LB.ListenerRule("oidc", new()
    {
        ListenerArn = frontEndListener.Arn,
        Actions = new[]
        {
            new Aws.LB.Inputs.ListenerRuleActionArgs
            {
                Type = "authenticate-oidc",
                AuthenticateOidc = new Aws.LB.Inputs.ListenerRuleActionAuthenticateOidcArgs
                {
                    AuthorizationEndpoint = "https://example.com/authorization_endpoint",
                    ClientId = "client_id",
                    ClientSecret = "client_secret",
                    Issuer = "https://example.com",
                    TokenEndpoint = "https://example.com/token_endpoint",
                    UserInfoEndpoint = "https://example.com/user_info_endpoint",
                },
            },
            new Aws.LB.Inputs.ListenerRuleActionArgs
            {
                Type = "forward",
                TargetGroupArn = aws_lb_target_group.Static.Arn,
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cognito"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lb.NewLoadBalancer(ctx, "frontEndLoadBalancer", nil)
		if err != nil {
			return err
		}
		frontEndListener, err := lb.NewListener(ctx, "frontEndListener", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "static", &lb.ListenerRuleArgs{
			ListenerArn: frontEndListener.Arn,
			Priority:    pulumi.Int(100),
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					Type:           pulumi.String("forward"),
					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
				},
			},
			Conditions: lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					PathPattern: &lb.ListenerRuleConditionPathPatternArgs{
						Values: pulumi.StringArray{
							pulumi.String("/static/*"),
						},
					},
				},
				&lb.ListenerRuleConditionArgs{
					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
						Values: pulumi.StringArray{
							pulumi.String("example.com"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "hostBasedWeightedRouting", &lb.ListenerRuleArgs{
			ListenerArn: frontEndListener.Arn,
			Priority:    pulumi.Int(99),
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					Type:           pulumi.String("forward"),
					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
				},
			},
			Conditions: lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
						Values: pulumi.StringArray{
							pulumi.String("my-service.*.mycompany.io"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "hostBasedRouting", &lb.ListenerRuleArgs{
			ListenerArn: frontEndListener.Arn,
			Priority:    pulumi.Int(99),
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					Type: pulumi.String("forward"),
					Forward: &lb.ListenerRuleActionForwardArgs{
						TargetGroups: lb.ListenerRuleActionForwardTargetGroupArray{
							&lb.ListenerRuleActionForwardTargetGroupArgs{
								Arn:    pulumi.Any(aws_lb_target_group.Main.Arn),
								Weight: pulumi.Int(80),
							},
							&lb.ListenerRuleActionForwardTargetGroupArgs{
								Arn:    pulumi.Any(aws_lb_target_group.Canary.Arn),
								Weight: pulumi.Int(20),
							},
						},
						Stickiness: &lb.ListenerRuleActionForwardStickinessArgs{
							Enabled:  pulumi.Bool(true),
							Duration: pulumi.Int(600),
						},
					},
				},
			},
			Conditions: lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
						Values: pulumi.StringArray{
							pulumi.String("my-service.*.mycompany.io"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "redirectHttpToHttps", &lb.ListenerRuleArgs{
			ListenerArn: frontEndListener.Arn,
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					Type: pulumi.String("redirect"),
					Redirect: &lb.ListenerRuleActionRedirectArgs{
						Port:       pulumi.String("443"),
						Protocol:   pulumi.String("HTTPS"),
						StatusCode: pulumi.String("HTTP_301"),
					},
				},
			},
			Conditions: lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					HttpHeader: &lb.ListenerRuleConditionHttpHeaderArgs{
						HttpHeaderName: pulumi.String("X-Forwarded-For"),
						Values: pulumi.StringArray{
							pulumi.String("192.168.1.*"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "healthCheck", &lb.ListenerRuleArgs{
			ListenerArn: frontEndListener.Arn,
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					Type: pulumi.String("fixed-response"),
					FixedResponse: &lb.ListenerRuleActionFixedResponseArgs{
						ContentType: pulumi.String("text/plain"),
						MessageBody: pulumi.String("HEALTHY"),
						StatusCode:  pulumi.String("200"),
					},
				},
			},
			Conditions: lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					QueryStrings: lb.ListenerRuleConditionQueryStringArray{
						&lb.ListenerRuleConditionQueryStringArgs{
							Key:   pulumi.String("health"),
							Value: pulumi.String("check"),
						},
						&lb.ListenerRuleConditionQueryStringArgs{
							Value: pulumi.String("bar"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		pool, err := cognito.NewUserPool(ctx, "pool", nil)
		if err != nil {
			return err
		}
		client, err := cognito.NewUserPoolClient(ctx, "client", nil)
		if err != nil {
			return err
		}
		domain, err := cognito.NewUserPoolDomain(ctx, "domain", nil)
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "admin", &lb.ListenerRuleArgs{
			ListenerArn: frontEndListener.Arn,
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					Type: pulumi.String("authenticate-cognito"),
					AuthenticateCognito: &lb.ListenerRuleActionAuthenticateCognitoArgs{
						UserPoolArn:      pool.Arn,
						UserPoolClientId: client.ID(),
						UserPoolDomain:   domain.Domain,
					},
				},
				&lb.ListenerRuleActionArgs{
					Type:           pulumi.String("forward"),
					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
				},
			},
		})
		if err != nil {
			return err
		}
		_, err = lb.NewListenerRule(ctx, "oidc", &lb.ListenerRuleArgs{
			ListenerArn: frontEndListener.Arn,
			Actions: lb.ListenerRuleActionArray{
				&lb.ListenerRuleActionArgs{
					Type: pulumi.String("authenticate-oidc"),
					AuthenticateOidc: &lb.ListenerRuleActionAuthenticateOidcArgs{
						AuthorizationEndpoint: pulumi.String("https://example.com/authorization_endpoint"),
						ClientId:              pulumi.String("client_id"),
						ClientSecret:          pulumi.String("client_secret"),
						Issuer:                pulumi.String("https://example.com"),
						TokenEndpoint:         pulumi.String("https://example.com/token_endpoint"),
						UserInfoEndpoint:      pulumi.String("https://example.com/user_info_endpoint"),
					},
				},
				&lb.ListenerRuleActionArgs{
					Type:           pulumi.String("forward"),
					TargetGroupArn: pulumi.Any(aws_lb_target_group.Static.Arn),
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
import com.pulumi.aws.lb.LoadBalancer;
import com.pulumi.aws.lb.Listener;
import com.pulumi.aws.lb.ListenerRule;
import com.pulumi.aws.lb.ListenerRuleArgs;
import com.pulumi.aws.lb.inputs.ListenerRuleActionArgs;
import com.pulumi.aws.lb.inputs.ListenerRuleConditionArgs;
import com.pulumi.aws.lb.inputs.ListenerRuleConditionPathPatternArgs;
import com.pulumi.aws.lb.inputs.ListenerRuleConditionHostHeaderArgs;
import com.pulumi.aws.lb.inputs.ListenerRuleActionForwardArgs;
import com.pulumi.aws.lb.inputs.ListenerRuleActionForwardStickinessArgs;
import com.pulumi.aws.lb.inputs.ListenerRuleActionRedirectArgs;
import com.pulumi.aws.lb.inputs.ListenerRuleConditionHttpHeaderArgs;
import com.pulumi.aws.lb.inputs.ListenerRuleActionFixedResponseArgs;
import com.pulumi.aws.cognito.UserPool;
import com.pulumi.aws.cognito.UserPoolClient;
import com.pulumi.aws.cognito.UserPoolDomain;
import com.pulumi.aws.lb.inputs.ListenerRuleActionAuthenticateCognitoArgs;
import com.pulumi.aws.lb.inputs.ListenerRuleActionAuthenticateOidcArgs;
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
        var frontEndLoadBalancer = new LoadBalancer("frontEndLoadBalancer");

        var frontEndListener = new Listener("frontEndListener");

        var static_ = new ListenerRule("static", ListenerRuleArgs.builder()        
            .listenerArn(frontEndListener.arn())
            .priority(100)
            .actions(ListenerRuleActionArgs.builder()
                .type("forward")
                .targetGroupArn(aws_lb_target_group.static().arn())
                .build())
            .conditions(            
                ListenerRuleConditionArgs.builder()
                    .pathPattern(ListenerRuleConditionPathPatternArgs.builder()
                        .values("/static/*")
                        .build())
                    .build(),
                ListenerRuleConditionArgs.builder()
                    .hostHeader(ListenerRuleConditionHostHeaderArgs.builder()
                        .values("example.com")
                        .build())
                    .build())
            .build());

        var hostBasedWeightedRouting = new ListenerRule("hostBasedWeightedRouting", ListenerRuleArgs.builder()        
            .listenerArn(frontEndListener.arn())
            .priority(99)
            .actions(ListenerRuleActionArgs.builder()
                .type("forward")
                .targetGroupArn(aws_lb_target_group.static().arn())
                .build())
            .conditions(ListenerRuleConditionArgs.builder()
                .hostHeader(ListenerRuleConditionHostHeaderArgs.builder()
                    .values("my-service.*.mycompany.io")
                    .build())
                .build())
            .build());

        var hostBasedRouting = new ListenerRule("hostBasedRouting", ListenerRuleArgs.builder()        
            .listenerArn(frontEndListener.arn())
            .priority(99)
            .actions(ListenerRuleActionArgs.builder()
                .type("forward")
                .forward(ListenerRuleActionForwardArgs.builder()
                    .targetGroups(                    
                        ListenerRuleActionForwardTargetGroupArgs.builder()
                            .arn(aws_lb_target_group.main().arn())
                            .weight(80)
                            .build(),
                        ListenerRuleActionForwardTargetGroupArgs.builder()
                            .arn(aws_lb_target_group.canary().arn())
                            .weight(20)
                            .build())
                    .stickiness(ListenerRuleActionForwardStickinessArgs.builder()
                        .enabled(true)
                        .duration(600)
                        .build())
                    .build())
                .build())
            .conditions(ListenerRuleConditionArgs.builder()
                .hostHeader(ListenerRuleConditionHostHeaderArgs.builder()
                    .values("my-service.*.mycompany.io")
                    .build())
                .build())
            .build());

        var redirectHttpToHttps = new ListenerRule("redirectHttpToHttps", ListenerRuleArgs.builder()        
            .listenerArn(frontEndListener.arn())
            .actions(ListenerRuleActionArgs.builder()
                .type("redirect")
                .redirect(ListenerRuleActionRedirectArgs.builder()
                    .port("443")
                    .protocol("HTTPS")
                    .statusCode("HTTP_301")
                    .build())
                .build())
            .conditions(ListenerRuleConditionArgs.builder()
                .httpHeader(ListenerRuleConditionHttpHeaderArgs.builder()
                    .httpHeaderName("X-Forwarded-For")
                    .values("192.168.1.*")
                    .build())
                .build())
            .build());

        var healthCheck = new ListenerRule("healthCheck", ListenerRuleArgs.builder()        
            .listenerArn(frontEndListener.arn())
            .actions(ListenerRuleActionArgs.builder()
                .type("fixed-response")
                .fixedResponse(ListenerRuleActionFixedResponseArgs.builder()
                    .contentType("text/plain")
                    .messageBody("HEALTHY")
                    .statusCode("200")
                    .build())
                .build())
            .conditions(ListenerRuleConditionArgs.builder()
                .queryStrings(                
                    ListenerRuleConditionQueryStringArgs.builder()
                        .key("health")
                        .value("check")
                        .build(),
                    ListenerRuleConditionQueryStringArgs.builder()
                        .value("bar")
                        .build())
                .build())
            .build());

        var pool = new UserPool("pool");

        var client = new UserPoolClient("client");

        var domain = new UserPoolDomain("domain");

        var admin = new ListenerRule("admin", ListenerRuleArgs.builder()        
            .listenerArn(frontEndListener.arn())
            .actions(            
                ListenerRuleActionArgs.builder()
                    .type("authenticate-cognito")
                    .authenticateCognito(ListenerRuleActionAuthenticateCognitoArgs.builder()
                        .userPoolArn(pool.arn())
                        .userPoolClientId(client.id())
                        .userPoolDomain(domain.domain())
                        .build())
                    .build(),
                ListenerRuleActionArgs.builder()
                    .type("forward")
                    .targetGroupArn(aws_lb_target_group.static().arn())
                    .build())
            .build());

        var oidc = new ListenerRule("oidc", ListenerRuleArgs.builder()        
            .listenerArn(frontEndListener.arn())
            .actions(            
                ListenerRuleActionArgs.builder()
                    .type("authenticate-oidc")
                    .authenticateOidc(ListenerRuleActionAuthenticateOidcArgs.builder()
                        .authorizationEndpoint("https://example.com/authorization_endpoint")
                        .clientId("client_id")
                        .clientSecret("client_secret")
                        .issuer("https://example.com")
                        .tokenEndpoint("https://example.com/token_endpoint")
                        .userInfoEndpoint("https://example.com/user_info_endpoint")
                        .build())
                    .build(),
                ListenerRuleActionArgs.builder()
                    .type("forward")
                    .targetGroupArn(aws_lb_target_group.static().arn())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  frontEndLoadBalancer:
    type: aws:lb:LoadBalancer
  frontEndListener:
    type: aws:lb:Listener
  static: # Forward action
    type: aws:lb:ListenerRule
    properties:
      listenerArn: ${frontEndListener.arn}
      priority: 100
      actions:
        - type: forward
          targetGroupArn: ${aws_lb_target_group.static.arn}
      conditions:
        - pathPattern:
            values:
              - /static/*
        - hostHeader:
            values:
              - example.com
  hostBasedWeightedRouting: # Weighted Forward action
    type: aws:lb:ListenerRule
    properties:
      listenerArn: ${frontEndListener.arn}
      priority: 99
      actions:
        - type: forward
          targetGroupArn: ${aws_lb_target_group.static.arn}
      conditions:
        - hostHeader:
            values:
              - my-service.*.mycompany.io
  hostBasedRouting: # Redirect action
    type: aws:lb:ListenerRule
    properties:
      listenerArn: ${frontEndListener.arn}
      priority: 99
      actions:
        - type: forward
          forward:
            targetGroups:
              - arn: ${aws_lb_target_group.main.arn}
                weight: 80
              - arn: ${aws_lb_target_group.canary.arn}
                weight: 20
            stickiness:
              enabled: true
              duration: 600
      conditions:
        - hostHeader:
            values:
              - my-service.*.mycompany.io
  redirectHttpToHttps: # Fixed-response action
    type: aws:lb:ListenerRule
    properties:
      listenerArn: ${frontEndListener.arn}
      actions:
        - type: redirect
          redirect:
            port: '443'
            protocol: HTTPS
            statusCode: HTTP_301
      conditions:
        - httpHeader:
            httpHeaderName: X-Forwarded-For
            values:
              - 192.168.1.*
  healthCheck: # Authenticate-cognito Action
    type: aws:lb:ListenerRule
    properties:
      listenerArn: ${frontEndListener.arn}
      actions:
        - type: fixed-response
          fixedResponse:
            contentType: text/plain
            messageBody: HEALTHY
            statusCode: '200'
      conditions:
        - queryStrings:
            - key: health
              value: check
            - value: bar
  pool:
    type: aws:cognito:UserPool
  client:
    type: aws:cognito:UserPoolClient
  domain:
    type: aws:cognito:UserPoolDomain
  admin: # Authenticate-oidc Action
    type: aws:lb:ListenerRule
    properties:
      listenerArn: ${frontEndListener.arn}
      actions:
        - type: authenticate-cognito
          authenticateCognito:
            userPoolArn: ${pool.arn}
            userPoolClientId: ${client.id}
            userPoolDomain: ${domain.domain}
        - type: forward
          targetGroupArn: ${aws_lb_target_group.static.arn}
  oidc:
    type: aws:lb:ListenerRule
    properties:
      listenerArn: ${frontEndListener.arn}
      actions:
        - type: authenticate-oidc
          authenticateOidc:
            authorizationEndpoint: https://example.com/authorization_endpoint
            clientId: client_id
            clientSecret: client_secret
            issuer: https://example.com
            tokenEndpoint: https://example.com/token_endpoint
            userInfoEndpoint: https://example.com/user_info_endpoint
        - type: forward
          targetGroupArn: ${aws_lb_target_group.static.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import rules using their ARN. For example:

```sh
 $ pulumi import aws:lb/listenerRule:ListenerRule front_end arn:aws:elasticloadbalancing:us-west-2:187416307283:listener-rule/app/test/8e4497da625e2d8a/9ab28ade35828f96/67b3d2d36dd7c26b
```
 