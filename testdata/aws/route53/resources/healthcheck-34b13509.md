Provides a Route53 health check.

{{% examples %}}
## Example Usage
{{% example %}}
### Connectivity and HTTP Status Code Check

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53.HealthCheck("example", {
    failureThreshold: 5,
    fqdn: "example.com",
    port: 80,
    requestInterval: 30,
    resourcePath: "/",
    tags: {
        Name: "tf-test-health-check",
    },
    type: "HTTP",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.HealthCheck("example",
    failure_threshold=5,
    fqdn="example.com",
    port=80,
    request_interval=30,
    resource_path="/",
    tags={
        "Name": "tf-test-health-check",
    },
    type="HTTP")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53.HealthCheck("example", new()
    {
        FailureThreshold = 5,
        Fqdn = "example.com",
        Port = 80,
        RequestInterval = 30,
        ResourcePath = "/",
        Tags = 
        {
            { "Name", "tf-test-health-check" },
        },
        Type = "HTTP",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewHealthCheck(ctx, "example", &route53.HealthCheckArgs{
			FailureThreshold: pulumi.Int(5),
			Fqdn:             pulumi.String("example.com"),
			Port:             pulumi.Int(80),
			RequestInterval:  pulumi.Int(30),
			ResourcePath:     pulumi.String("/"),
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-test-health-check"),
			},
			Type: pulumi.String("HTTP"),
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
import com.pulumi.aws.route53.HealthCheck;
import com.pulumi.aws.route53.HealthCheckArgs;
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
        var example = new HealthCheck("example", HealthCheckArgs.builder()        
            .failureThreshold("5")
            .fqdn("example.com")
            .port(80)
            .requestInterval("30")
            .resourcePath("/")
            .tags(Map.of("Name", "tf-test-health-check"))
            .type("HTTP")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53:HealthCheck
    properties:
      failureThreshold: '5'
      fqdn: example.com
      port: 80
      requestInterval: '30'
      resourcePath: /
      tags:
        Name: tf-test-health-check
      type: HTTP
```
{{% /example %}}
{{% example %}}
### Connectivity and String Matching Check

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.route53.HealthCheck("example", {
    failureThreshold: 5,
    fqdn: "example.com",
    port: 443,
    requestInterval: 30,
    resourcePath: "/",
    searchString: "example",
    type: "HTTPS_STR_MATCH",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.route53.HealthCheck("example",
    failure_threshold=5,
    fqdn="example.com",
    port=443,
    request_interval=30,
    resource_path="/",
    search_string="example",
    type="HTTPS_STR_MATCH")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Route53.HealthCheck("example", new()
    {
        FailureThreshold = 5,
        Fqdn = "example.com",
        Port = 443,
        RequestInterval = 30,
        ResourcePath = "/",
        SearchString = "example",
        Type = "HTTPS_STR_MATCH",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewHealthCheck(ctx, "example", &route53.HealthCheckArgs{
			FailureThreshold: pulumi.Int(5),
			Fqdn:             pulumi.String("example.com"),
			Port:             pulumi.Int(443),
			RequestInterval:  pulumi.Int(30),
			ResourcePath:     pulumi.String("/"),
			SearchString:     pulumi.String("example"),
			Type:             pulumi.String("HTTPS_STR_MATCH"),
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
import com.pulumi.aws.route53.HealthCheck;
import com.pulumi.aws.route53.HealthCheckArgs;
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
        var example = new HealthCheck("example", HealthCheckArgs.builder()        
            .failureThreshold("5")
            .fqdn("example.com")
            .port(443)
            .requestInterval("30")
            .resourcePath("/")
            .searchString("example")
            .type("HTTPS_STR_MATCH")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:route53:HealthCheck
    properties:
      failureThreshold: '5'
      fqdn: example.com
      port: 443
      requestInterval: '30'
      resourcePath: /
      searchString: example
      type: HTTPS_STR_MATCH
```
{{% /example %}}
{{% example %}}
### Aggregate Check

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const parent = new aws.route53.HealthCheck("parent", {
    type: "CALCULATED",
    childHealthThreshold: 1,
    childHealthchecks: [aws_route53_health_check.child.id],
    tags: {
        Name: "tf-test-calculated-health-check",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

parent = aws.route53.HealthCheck("parent",
    type="CALCULATED",
    child_health_threshold=1,
    child_healthchecks=[aws_route53_health_check["child"]["id"]],
    tags={
        "Name": "tf-test-calculated-health-check",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var parent = new Aws.Route53.HealthCheck("parent", new()
    {
        Type = "CALCULATED",
        ChildHealthThreshold = 1,
        ChildHealthchecks = new[]
        {
            aws_route53_health_check.Child.Id,
        },
        Tags = 
        {
            { "Name", "tf-test-calculated-health-check" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := route53.NewHealthCheck(ctx, "parent", &route53.HealthCheckArgs{
			Type:                 pulumi.String("CALCULATED"),
			ChildHealthThreshold: pulumi.Int(1),
			ChildHealthchecks: pulumi.StringArray{
				aws_route53_health_check.Child.Id,
			},
			Tags: pulumi.StringMap{
				"Name": pulumi.String("tf-test-calculated-health-check"),
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
import com.pulumi.aws.route53.HealthCheck;
import com.pulumi.aws.route53.HealthCheckArgs;
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
        var parent = new HealthCheck("parent", HealthCheckArgs.builder()        
            .type("CALCULATED")
            .childHealthThreshold(1)
            .childHealthchecks(aws_route53_health_check.child().id())
            .tags(Map.of("Name", "tf-test-calculated-health-check"))
            .build());

    }
}
```
```yaml
resources:
  parent:
    type: aws:route53:HealthCheck
    properties:
      type: CALCULATED
      childHealthThreshold: 1
      childHealthchecks:
        - ${aws_route53_health_check.child.id}
      tags:
        Name: tf-test-calculated-health-check
```
{{% /example %}}
{{% example %}}
### CloudWatch Alarm Check

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const foobar = new aws.cloudwatch.MetricAlarm("foobar", {
    comparisonOperator: "GreaterThanOrEqualToThreshold",
    evaluationPeriods: 2,
    metricName: "CPUUtilization",
    namespace: "AWS/EC2",
    period: 120,
    statistic: "Average",
    threshold: 80,
    alarmDescription: "This metric monitors ec2 cpu utilization",
});
const foo = new aws.route53.HealthCheck("foo", {
    type: "CLOUDWATCH_METRIC",
    cloudwatchAlarmName: foobar.name,
    cloudwatchAlarmRegion: "us-west-2",
    insufficientDataHealthStatus: "Healthy",
});
```
```python
import pulumi
import pulumi_aws as aws

foobar = aws.cloudwatch.MetricAlarm("foobar",
    comparison_operator="GreaterThanOrEqualToThreshold",
    evaluation_periods=2,
    metric_name="CPUUtilization",
    namespace="AWS/EC2",
    period=120,
    statistic="Average",
    threshold=80,
    alarm_description="This metric monitors ec2 cpu utilization")
foo = aws.route53.HealthCheck("foo",
    type="CLOUDWATCH_METRIC",
    cloudwatch_alarm_name=foobar.name,
    cloudwatch_alarm_region="us-west-2",
    insufficient_data_health_status="Healthy")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var foobar = new Aws.CloudWatch.MetricAlarm("foobar", new()
    {
        ComparisonOperator = "GreaterThanOrEqualToThreshold",
        EvaluationPeriods = 2,
        MetricName = "CPUUtilization",
        Namespace = "AWS/EC2",
        Period = 120,
        Statistic = "Average",
        Threshold = 80,
        AlarmDescription = "This metric monitors ec2 cpu utilization",
    });

    var foo = new Aws.Route53.HealthCheck("foo", new()
    {
        Type = "CLOUDWATCH_METRIC",
        CloudwatchAlarmName = foobar.Name,
        CloudwatchAlarmRegion = "us-west-2",
        InsufficientDataHealthStatus = "Healthy",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		foobar, err := cloudwatch.NewMetricAlarm(ctx, "foobar", &cloudwatch.MetricAlarmArgs{
			ComparisonOperator: pulumi.String("GreaterThanOrEqualToThreshold"),
			EvaluationPeriods:  pulumi.Int(2),
			MetricName:         pulumi.String("CPUUtilization"),
			Namespace:          pulumi.String("AWS/EC2"),
			Period:             pulumi.Int(120),
			Statistic:          pulumi.String("Average"),
			Threshold:          pulumi.Float64(80),
			AlarmDescription:   pulumi.String("This metric monitors ec2 cpu utilization"),
		})
		if err != nil {
			return err
		}
		_, err = route53.NewHealthCheck(ctx, "foo", &route53.HealthCheckArgs{
			Type:                         pulumi.String("CLOUDWATCH_METRIC"),
			CloudwatchAlarmName:          foobar.Name,
			CloudwatchAlarmRegion:        pulumi.String("us-west-2"),
			InsufficientDataHealthStatus: pulumi.String("Healthy"),
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
import com.pulumi.aws.cloudwatch.MetricAlarm;
import com.pulumi.aws.cloudwatch.MetricAlarmArgs;
import com.pulumi.aws.route53.HealthCheck;
import com.pulumi.aws.route53.HealthCheckArgs;
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
        var foobar = new MetricAlarm("foobar", MetricAlarmArgs.builder()        
            .comparisonOperator("GreaterThanOrEqualToThreshold")
            .evaluationPeriods("2")
            .metricName("CPUUtilization")
            .namespace("AWS/EC2")
            .period("120")
            .statistic("Average")
            .threshold("80")
            .alarmDescription("This metric monitors ec2 cpu utilization")
            .build());

        var foo = new HealthCheck("foo", HealthCheckArgs.builder()        
            .type("CLOUDWATCH_METRIC")
            .cloudwatchAlarmName(foobar.name())
            .cloudwatchAlarmRegion("us-west-2")
            .insufficientDataHealthStatus("Healthy")
            .build());

    }
}
```
```yaml
resources:
  foobar:
    type: aws:cloudwatch:MetricAlarm
    properties:
      comparisonOperator: GreaterThanOrEqualToThreshold
      evaluationPeriods: '2'
      metricName: CPUUtilization
      namespace: AWS/EC2
      period: '120'
      statistic: Average
      threshold: '80'
      alarmDescription: This metric monitors ec2 cpu utilization
  foo:
    type: aws:route53:HealthCheck
    properties:
      type: CLOUDWATCH_METRIC
      cloudwatchAlarmName: ${foobar.name}
      cloudwatchAlarmRegion: us-west-2
      insufficientDataHealthStatus: Healthy
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Route53 Health Checks using the health check `id`. For example:

```sh
 $ pulumi import aws:route53/healthCheck:HealthCheck http_check abcdef11-2222-3333-4444-555555fedcba
```
 