Provides a Lambda event source mapping. This allows Lambda functions to get events from Kinesis, DynamoDB, SQS, Amazon MQ and Managed Streaming for Apache Kafka (MSK).

For information about Lambda and how to use it, see [What is AWS Lambda?](http://docs.aws.amazon.com/lambda/latest/dg/welcome.html).
For information about event source mappings, see [CreateEventSourceMapping](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateEventSourceMapping.html) in the API docs.

{{% examples %}}
## Example Usage
{{% example %}}
### DynamoDB

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.EventSourceMapping("example", {
    eventSourceArn: aws_dynamodb_table.example.stream_arn,
    functionName: aws_lambda_function.example.arn,
    startingPosition: "LATEST",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lambda_.EventSourceMapping("example",
    event_source_arn=aws_dynamodb_table["example"]["stream_arn"],
    function_name=aws_lambda_function["example"]["arn"],
    starting_position="LATEST")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.EventSourceMapping("example", new()
    {
        EventSourceArn = aws_dynamodb_table.Example.Stream_arn,
        FunctionName = aws_lambda_function.Example.Arn,
        StartingPosition = "LATEST",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
			EventSourceArn:   pulumi.Any(aws_dynamodb_table.Example.Stream_arn),
			FunctionName:     pulumi.Any(aws_lambda_function.Example.Arn),
			StartingPosition: pulumi.String("LATEST"),
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
import com.pulumi.aws.lambda.EventSourceMapping;
import com.pulumi.aws.lambda.EventSourceMappingArgs;
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
        var example = new EventSourceMapping("example", EventSourceMappingArgs.builder()        
            .eventSourceArn(aws_dynamodb_table.example().stream_arn())
            .functionName(aws_lambda_function.example().arn())
            .startingPosition("LATEST")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:EventSourceMapping
    properties:
      eventSourceArn: ${aws_dynamodb_table.example.stream_arn}
      functionName: ${aws_lambda_function.example.arn}
      startingPosition: LATEST
```
{{% /example %}}
{{% example %}}
### Kinesis

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.EventSourceMapping("example", {
    eventSourceArn: aws_kinesis_stream.example.arn,
    functionName: aws_lambda_function.example.arn,
    startingPosition: "LATEST",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lambda_.EventSourceMapping("example",
    event_source_arn=aws_kinesis_stream["example"]["arn"],
    function_name=aws_lambda_function["example"]["arn"],
    starting_position="LATEST")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.EventSourceMapping("example", new()
    {
        EventSourceArn = aws_kinesis_stream.Example.Arn,
        FunctionName = aws_lambda_function.Example.Arn,
        StartingPosition = "LATEST",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
			EventSourceArn:   pulumi.Any(aws_kinesis_stream.Example.Arn),
			FunctionName:     pulumi.Any(aws_lambda_function.Example.Arn),
			StartingPosition: pulumi.String("LATEST"),
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
import com.pulumi.aws.lambda.EventSourceMapping;
import com.pulumi.aws.lambda.EventSourceMappingArgs;
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
        var example = new EventSourceMapping("example", EventSourceMappingArgs.builder()        
            .eventSourceArn(aws_kinesis_stream.example().arn())
            .functionName(aws_lambda_function.example().arn())
            .startingPosition("LATEST")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:EventSourceMapping
    properties:
      eventSourceArn: ${aws_kinesis_stream.example.arn}
      functionName: ${aws_lambda_function.example.arn}
      startingPosition: LATEST
```
{{% /example %}}
{{% example %}}
### Managed Streaming for Apache Kafka (MSK)

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.EventSourceMapping("example", {
    eventSourceArn: aws_msk_cluster.example.arn,
    functionName: aws_lambda_function.example.arn,
    topics: ["Example"],
    startingPosition: "TRIM_HORIZON",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lambda_.EventSourceMapping("example",
    event_source_arn=aws_msk_cluster["example"]["arn"],
    function_name=aws_lambda_function["example"]["arn"],
    topics=["Example"],
    starting_position="TRIM_HORIZON")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.EventSourceMapping("example", new()
    {
        EventSourceArn = aws_msk_cluster.Example.Arn,
        FunctionName = aws_lambda_function.Example.Arn,
        Topics = new[]
        {
            "Example",
        },
        StartingPosition = "TRIM_HORIZON",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
			EventSourceArn: pulumi.Any(aws_msk_cluster.Example.Arn),
			FunctionName:   pulumi.Any(aws_lambda_function.Example.Arn),
			Topics: pulumi.StringArray{
				pulumi.String("Example"),
			},
			StartingPosition: pulumi.String("TRIM_HORIZON"),
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
import com.pulumi.aws.lambda.EventSourceMapping;
import com.pulumi.aws.lambda.EventSourceMappingArgs;
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
        var example = new EventSourceMapping("example", EventSourceMappingArgs.builder()        
            .eventSourceArn(aws_msk_cluster.example().arn())
            .functionName(aws_lambda_function.example().arn())
            .topics("Example")
            .startingPosition("TRIM_HORIZON")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:EventSourceMapping
    properties:
      eventSourceArn: ${aws_msk_cluster.example.arn}
      functionName: ${aws_lambda_function.example.arn}
      topics:
        - Example
      startingPosition: TRIM_HORIZON
```
{{% /example %}}
{{% example %}}
### Self Managed Apache Kafka

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.EventSourceMapping("example", {
    functionName: aws_lambda_function.example.arn,
    topics: ["Example"],
    startingPosition: "TRIM_HORIZON",
    selfManagedEventSource: {
        endpoints: {
            KAFKA_BOOTSTRAP_SERVERS: "kafka1.example.com:9092,kafka2.example.com:9092",
        },
    },
    sourceAccessConfigurations: [
        {
            type: "VPC_SUBNET",
            uri: "subnet:subnet-example1",
        },
        {
            type: "VPC_SUBNET",
            uri: "subnet:subnet-example2",
        },
        {
            type: "VPC_SECURITY_GROUP",
            uri: "security_group:sg-example",
        },
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lambda_.EventSourceMapping("example",
    function_name=aws_lambda_function["example"]["arn"],
    topics=["Example"],
    starting_position="TRIM_HORIZON",
    self_managed_event_source=aws.lambda_.EventSourceMappingSelfManagedEventSourceArgs(
        endpoints={
            "KAFKA_BOOTSTRAP_SERVERS": "kafka1.example.com:9092,kafka2.example.com:9092",
        },
    ),
    source_access_configurations=[
        aws.lambda_.EventSourceMappingSourceAccessConfigurationArgs(
            type="VPC_SUBNET",
            uri="subnet:subnet-example1",
        ),
        aws.lambda_.EventSourceMappingSourceAccessConfigurationArgs(
            type="VPC_SUBNET",
            uri="subnet:subnet-example2",
        ),
        aws.lambda_.EventSourceMappingSourceAccessConfigurationArgs(
            type="VPC_SECURITY_GROUP",
            uri="security_group:sg-example",
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
    var example = new Aws.Lambda.EventSourceMapping("example", new()
    {
        FunctionName = aws_lambda_function.Example.Arn,
        Topics = new[]
        {
            "Example",
        },
        StartingPosition = "TRIM_HORIZON",
        SelfManagedEventSource = new Aws.Lambda.Inputs.EventSourceMappingSelfManagedEventSourceArgs
        {
            Endpoints = 
            {
                { "KAFKA_BOOTSTRAP_SERVERS", "kafka1.example.com:9092,kafka2.example.com:9092" },
            },
        },
        SourceAccessConfigurations = new[]
        {
            new Aws.Lambda.Inputs.EventSourceMappingSourceAccessConfigurationArgs
            {
                Type = "VPC_SUBNET",
                Uri = "subnet:subnet-example1",
            },
            new Aws.Lambda.Inputs.EventSourceMappingSourceAccessConfigurationArgs
            {
                Type = "VPC_SUBNET",
                Uri = "subnet:subnet-example2",
            },
            new Aws.Lambda.Inputs.EventSourceMappingSourceAccessConfigurationArgs
            {
                Type = "VPC_SECURITY_GROUP",
                Uri = "security_group:sg-example",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
			FunctionName: pulumi.Any(aws_lambda_function.Example.Arn),
			Topics: pulumi.StringArray{
				pulumi.String("Example"),
			},
			StartingPosition: pulumi.String("TRIM_HORIZON"),
			SelfManagedEventSource: &lambda.EventSourceMappingSelfManagedEventSourceArgs{
				Endpoints: pulumi.StringMap{
					"KAFKA_BOOTSTRAP_SERVERS": pulumi.String("kafka1.example.com:9092,kafka2.example.com:9092"),
				},
			},
			SourceAccessConfigurations: lambda.EventSourceMappingSourceAccessConfigurationArray{
				&lambda.EventSourceMappingSourceAccessConfigurationArgs{
					Type: pulumi.String("VPC_SUBNET"),
					Uri:  pulumi.String("subnet:subnet-example1"),
				},
				&lambda.EventSourceMappingSourceAccessConfigurationArgs{
					Type: pulumi.String("VPC_SUBNET"),
					Uri:  pulumi.String("subnet:subnet-example2"),
				},
				&lambda.EventSourceMappingSourceAccessConfigurationArgs{
					Type: pulumi.String("VPC_SECURITY_GROUP"),
					Uri:  pulumi.String("security_group:sg-example"),
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
import com.pulumi.aws.lambda.EventSourceMapping;
import com.pulumi.aws.lambda.EventSourceMappingArgs;
import com.pulumi.aws.lambda.inputs.EventSourceMappingSelfManagedEventSourceArgs;
import com.pulumi.aws.lambda.inputs.EventSourceMappingSourceAccessConfigurationArgs;
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
        var example = new EventSourceMapping("example", EventSourceMappingArgs.builder()        
            .functionName(aws_lambda_function.example().arn())
            .topics("Example")
            .startingPosition("TRIM_HORIZON")
            .selfManagedEventSource(EventSourceMappingSelfManagedEventSourceArgs.builder()
                .endpoints(Map.of("KAFKA_BOOTSTRAP_SERVERS", "kafka1.example.com:9092,kafka2.example.com:9092"))
                .build())
            .sourceAccessConfigurations(            
                EventSourceMappingSourceAccessConfigurationArgs.builder()
                    .type("VPC_SUBNET")
                    .uri("subnet:subnet-example1")
                    .build(),
                EventSourceMappingSourceAccessConfigurationArgs.builder()
                    .type("VPC_SUBNET")
                    .uri("subnet:subnet-example2")
                    .build(),
                EventSourceMappingSourceAccessConfigurationArgs.builder()
                    .type("VPC_SECURITY_GROUP")
                    .uri("security_group:sg-example")
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:EventSourceMapping
    properties:
      functionName: ${aws_lambda_function.example.arn}
      topics:
        - Example
      startingPosition: TRIM_HORIZON
      selfManagedEventSource:
        endpoints:
          KAFKA_BOOTSTRAP_SERVERS: kafka1.example.com:9092,kafka2.example.com:9092
      sourceAccessConfigurations:
        - type: VPC_SUBNET
          uri: subnet:subnet-example1
        - type: VPC_SUBNET
          uri: subnet:subnet-example2
        - type: VPC_SECURITY_GROUP
          uri: security_group:sg-example
```
{{% /example %}}
{{% example %}}
### SQS

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.EventSourceMapping("example", {
    eventSourceArn: aws_sqs_queue.sqs_queue_test.arn,
    functionName: aws_lambda_function.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.lambda_.EventSourceMapping("example",
    event_source_arn=aws_sqs_queue["sqs_queue_test"]["arn"],
    function_name=aws_lambda_function["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.EventSourceMapping("example", new()
    {
        EventSourceArn = aws_sqs_queue.Sqs_queue_test.Arn,
        FunctionName = aws_lambda_function.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
			EventSourceArn: pulumi.Any(aws_sqs_queue.Sqs_queue_test.Arn),
			FunctionName:   pulumi.Any(aws_lambda_function.Example.Arn),
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
import com.pulumi.aws.lambda.EventSourceMapping;
import com.pulumi.aws.lambda.EventSourceMappingArgs;
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
        var example = new EventSourceMapping("example", EventSourceMappingArgs.builder()        
            .eventSourceArn(aws_sqs_queue.sqs_queue_test().arn())
            .functionName(aws_lambda_function.example().arn())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:EventSourceMapping
    properties:
      eventSourceArn: ${aws_sqs_queue.sqs_queue_test.arn}
      functionName: ${aws_lambda_function.example.arn}
```
{{% /example %}}
{{% example %}}
### SQS with event filter

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.lambda.EventSourceMapping("example", {
    eventSourceArn: aws_sqs_queue.sqs_queue_test.arn,
    functionName: aws_lambda_function.example.arn,
    filterCriteria: {
        filters: [{
            pattern: JSON.stringify({
                body: {
                    Temperature: [{
                        numeric: [
                            ">",
                            0,
                            "<=",
                            100,
                        ],
                    }],
                    Location: ["New York"],
                },
            }),
        }],
    },
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example = aws.lambda_.EventSourceMapping("example",
    event_source_arn=aws_sqs_queue["sqs_queue_test"]["arn"],
    function_name=aws_lambda_function["example"]["arn"],
    filter_criteria=aws.lambda_.EventSourceMappingFilterCriteriaArgs(
        filters=[aws.lambda_.EventSourceMappingFilterCriteriaFilterArgs(
            pattern=json.dumps({
                "body": {
                    "Temperature": [{
                        "numeric": [
                            ">",
                            0,
                            "<=",
                            100,
                        ],
                    }],
                    "Location": ["New York"],
                },
            }),
        )],
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Lambda.EventSourceMapping("example", new()
    {
        EventSourceArn = aws_sqs_queue.Sqs_queue_test.Arn,
        FunctionName = aws_lambda_function.Example.Arn,
        FilterCriteria = new Aws.Lambda.Inputs.EventSourceMappingFilterCriteriaArgs
        {
            Filters = new[]
            {
                new Aws.Lambda.Inputs.EventSourceMappingFilterCriteriaFilterArgs
                {
                    Pattern = JsonSerializer.Serialize(new Dictionary<string, object?>
                    {
                        ["body"] = new Dictionary<string, object?>
                        {
                            ["Temperature"] = new[]
                            {
                                new Dictionary<string, object?>
                                {
                                    ["numeric"] = new[]
                                    {
                                        ">",
                                        0,
                                        "<=",
                                        100,
                                    },
                                },
                            },
                            ["Location"] = new[]
                            {
                                "New York",
                            },
                        },
                    }),
                },
            },
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"body": map[string]interface{}{
				"Temperature": []map[string]interface{}{
					map[string]interface{}{
						"numeric": []interface{}{
							">",
							0,
							"<=",
							100,
						},
					},
				},
				"Location": []string{
					"New York",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = lambda.NewEventSourceMapping(ctx, "example", &lambda.EventSourceMappingArgs{
			EventSourceArn: pulumi.Any(aws_sqs_queue.Sqs_queue_test.Arn),
			FunctionName:   pulumi.Any(aws_lambda_function.Example.Arn),
			FilterCriteria: &lambda.EventSourceMappingFilterCriteriaArgs{
				Filters: lambda.EventSourceMappingFilterCriteriaFilterArray{
					&lambda.EventSourceMappingFilterCriteriaFilterArgs{
						Pattern: pulumi.String(json0),
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
import com.pulumi.aws.lambda.EventSourceMapping;
import com.pulumi.aws.lambda.EventSourceMappingArgs;
import com.pulumi.aws.lambda.inputs.EventSourceMappingFilterCriteriaArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var example = new EventSourceMapping("example", EventSourceMappingArgs.builder()        
            .eventSourceArn(aws_sqs_queue.sqs_queue_test().arn())
            .functionName(aws_lambda_function.example().arn())
            .filterCriteria(EventSourceMappingFilterCriteriaArgs.builder()
                .filters(EventSourceMappingFilterCriteriaFilterArgs.builder()
                    .pattern(serializeJson(
                        jsonObject(
                            jsonProperty("body", jsonObject(
                                jsonProperty("Temperature", jsonArray(jsonObject(
                                    jsonProperty("numeric", jsonArray(
                                        ">", 
                                        0, 
                                        "<=", 
                                        100
                                    ))
                                ))),
                                jsonProperty("Location", jsonArray("New York"))
                            ))
                        )))
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:EventSourceMapping
    properties:
      eventSourceArn: ${aws_sqs_queue.sqs_queue_test.arn}
      functionName: ${aws_lambda_function.example.arn}
      filterCriteria:
        filters:
          - pattern:
              fn::toJSON:
                body:
                  Temperature:
                    - numeric:
                        - '>'
                        - 0
                        - <=
                        - 100
                  Location:
                    - New York
```
{{% /example %}}
{{% example %}}
### Amazon MQ (ActiveMQ)

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.lambda.EventSourceMapping;
import com.pulumi.aws.lambda.EventSourceMappingArgs;
import com.pulumi.aws.lambda.inputs.EventSourceMappingSourceAccessConfigurationArgs;
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
        var example = new EventSourceMapping("example", EventSourceMappingArgs.builder()        
            .batchSize(10)
            .eventSourceArn(aws_mq_broker.example().arn())
            .enabled(true)
            .functionName(aws_lambda_function.example().arn())
            .queues("example")
            .sourceAccessConfigurations(EventSourceMappingSourceAccessConfigurationArgs.builder()
                .type("BASIC_AUTH")
                .uri(aws_secretsmanager_secret_version.example().arn())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:EventSourceMapping
    properties:
      batchSize: 10
      eventSourceArn: ${aws_mq_broker.example.arn}
      enabled: true
      functionName: ${aws_lambda_function.example.arn}
      queues:
        - example
      sourceAccessConfigurations:
        - type: BASIC_AUTH
          uri: ${aws_secretsmanager_secret_version.example.arn}
```
{{% /example %}}
{{% example %}}
### Amazon MQ (RabbitMQ)

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.lambda.EventSourceMapping;
import com.pulumi.aws.lambda.EventSourceMappingArgs;
import com.pulumi.aws.lambda.inputs.EventSourceMappingSourceAccessConfigurationArgs;
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
        var example = new EventSourceMapping("example", EventSourceMappingArgs.builder()        
            .batchSize(1)
            .eventSourceArn(aws_mq_broker.example().arn())
            .enabled(true)
            .functionName(aws_lambda_function.example().arn())
            .queues("example")
            .sourceAccessConfigurations(            
                EventSourceMappingSourceAccessConfigurationArgs.builder()
                    .type("VIRTUAL_HOST")
                    .uri("/example")
                    .build(),
                EventSourceMappingSourceAccessConfigurationArgs.builder()
                    .type("BASIC_AUTH")
                    .uri(aws_secretsmanager_secret_version.example().arn())
                    .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:lambda:EventSourceMapping
    properties:
      batchSize: 1
      eventSourceArn: ${aws_mq_broker.example.arn}
      enabled: true
      functionName: ${aws_lambda_function.example.arn}
      queues:
        - example
      sourceAccessConfigurations:
        - type: VIRTUAL_HOST
          uri: /example
        - type: BASIC_AUTH
          uri: ${aws_secretsmanager_secret_version.example.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Lambda event source mappings using the `UUID` (event source mapping identifier). For example:

```sh
 $ pulumi import aws:lambda/eventSourceMapping:EventSourceMapping event_source_mapping 12345kxodurf3443
```
 