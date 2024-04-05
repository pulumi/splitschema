Creates a MWAA Environment resource.

{{% examples %}}
## Example Usage

A MWAA Environment requires an IAM role (`aws.iam.Role`), two subnets in the private zone (`aws.ec2.Subnet`) and a versioned S3 bucket (`aws.s3.BucketV2`).
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.mwaa.Environment("example", {
    dagS3Path: "dags/",
    executionRoleArn: aws_iam_role.example.arn,
    networkConfiguration: {
        securityGroupIds: [aws_security_group.example.id],
        subnetIds: aws_subnet["private"].map(__item => __item.id),
    },
    sourceBucketArn: aws_s3_bucket.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.mwaa.Environment("example",
    dag_s3_path="dags/",
    execution_role_arn=aws_iam_role["example"]["arn"],
    network_configuration=aws.mwaa.EnvironmentNetworkConfigurationArgs(
        security_group_ids=[aws_security_group["example"]["id"]],
        subnet_ids=[__item["id"] for __item in aws_subnet["private"]],
    ),
    source_bucket_arn=aws_s3_bucket["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Mwaa.Environment("example", new()
    {
        DagS3Path = "dags/",
        ExecutionRoleArn = aws_iam_role.Example.Arn,
        NetworkConfiguration = new Aws.Mwaa.Inputs.EnvironmentNetworkConfigurationArgs
        {
            SecurityGroupIds = new[]
            {
                aws_security_group.Example.Id,
            },
            SubnetIds = aws_subnet.Private.Select(__item => __item.Id).ToList(),
        },
        SourceBucketArn = aws_s3_bucket.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mwaa"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := mwaa.NewEnvironment(ctx, "example", &mwaa.EnvironmentArgs{
DagS3Path: pulumi.String("dags/"),
ExecutionRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
NetworkConfiguration: &mwaa.EnvironmentNetworkConfigurationArgs{
SecurityGroupIds: pulumi.StringArray{
aws_security_group.Example.Id,
},
SubnetIds: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-resources-aws:mwaa-environment:Environment.pp:5,25-49),
},
SourceBucketArn: pulumi.Any(aws_s3_bucket.Example.Arn),
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
import com.pulumi.aws.mwaa.Environment;
import com.pulumi.aws.mwaa.EnvironmentArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentNetworkConfigurationArgs;
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
        var example = new Environment("example", EnvironmentArgs.builder()        
            .dagS3Path("dags/")
            .executionRoleArn(aws_iam_role.example().arn())
            .networkConfiguration(EnvironmentNetworkConfigurationArgs.builder()
                .securityGroupIds(aws_security_group.example().id())
                .subnetIds(aws_subnet.private().stream().map(element -> element.id()).collect(toList()))
                .build())
            .sourceBucketArn(aws_s3_bucket.example().arn())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Example with Airflow configuration options

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.mwaa.Environment("example", {
    airflowConfigurationOptions: {
        "core.default_task_retries": "16",
        "core.parallelism": "1",
    },
    dagS3Path: "dags/",
    executionRoleArn: aws_iam_role.example.arn,
    networkConfiguration: {
        securityGroupIds: [aws_security_group.example.id],
        subnetIds: aws_subnet["private"].map(__item => __item.id),
    },
    sourceBucketArn: aws_s3_bucket.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.mwaa.Environment("example",
    airflow_configuration_options={
        "core.default_task_retries": "16",
        "core.parallelism": "1",
    },
    dag_s3_path="dags/",
    execution_role_arn=aws_iam_role["example"]["arn"],
    network_configuration=aws.mwaa.EnvironmentNetworkConfigurationArgs(
        security_group_ids=[aws_security_group["example"]["id"]],
        subnet_ids=[__item["id"] for __item in aws_subnet["private"]],
    ),
    source_bucket_arn=aws_s3_bucket["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Mwaa.Environment("example", new()
    {
        AirflowConfigurationOptions = 
        {
            { "core.default_task_retries", "16" },
            { "core.parallelism", "1" },
        },
        DagS3Path = "dags/",
        ExecutionRoleArn = aws_iam_role.Example.Arn,
        NetworkConfiguration = new Aws.Mwaa.Inputs.EnvironmentNetworkConfigurationArgs
        {
            SecurityGroupIds = new[]
            {
                aws_security_group.Example.Id,
            },
            SubnetIds = aws_subnet.Private.Select(__item => __item.Id).ToList(),
        },
        SourceBucketArn = aws_s3_bucket.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mwaa"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := mwaa.NewEnvironment(ctx, "example", &mwaa.EnvironmentArgs{
AirflowConfigurationOptions: pulumi.StringMap{
"core.default_task_retries": pulumi.String("16"),
"core.parallelism": pulumi.String("1"),
},
DagS3Path: pulumi.String("dags/"),
ExecutionRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
NetworkConfiguration: &mwaa.EnvironmentNetworkConfigurationArgs{
SecurityGroupIds: pulumi.StringArray{
aws_security_group.Example.Id,
},
SubnetIds: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-resources-aws:mwaa-environment:Environment.pp:9,25-49),
},
SourceBucketArn: pulumi.Any(aws_s3_bucket.Example.Arn),
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
import com.pulumi.aws.mwaa.Environment;
import com.pulumi.aws.mwaa.EnvironmentArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentNetworkConfigurationArgs;
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
        var example = new Environment("example", EnvironmentArgs.builder()        
            .airflowConfigurationOptions(Map.ofEntries(
                Map.entry("core.default_task_retries", 16),
                Map.entry("core.parallelism", 1)
            ))
            .dagS3Path("dags/")
            .executionRoleArn(aws_iam_role.example().arn())
            .networkConfiguration(EnvironmentNetworkConfigurationArgs.builder()
                .securityGroupIds(aws_security_group.example().id())
                .subnetIds(aws_subnet.private().stream().map(element -> element.id()).collect(toList()))
                .build())
            .sourceBucketArn(aws_s3_bucket.example().arn())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Example with logging configurations

Note that Airflow task logs are enabled by default with the `INFO` log level.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.mwaa.Environment("example", {
    dagS3Path: "dags/",
    executionRoleArn: aws_iam_role.example.arn,
    loggingConfiguration: {
        dagProcessingLogs: {
            enabled: true,
            logLevel: "DEBUG",
        },
        schedulerLogs: {
            enabled: true,
            logLevel: "INFO",
        },
        taskLogs: {
            enabled: true,
            logLevel: "WARNING",
        },
        webserverLogs: {
            enabled: true,
            logLevel: "ERROR",
        },
        workerLogs: {
            enabled: true,
            logLevel: "CRITICAL",
        },
    },
    networkConfiguration: {
        securityGroupIds: [aws_security_group.example.id],
        subnetIds: aws_subnet["private"].map(__item => __item.id),
    },
    sourceBucketArn: aws_s3_bucket.example.arn,
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.mwaa.Environment("example",
    dag_s3_path="dags/",
    execution_role_arn=aws_iam_role["example"]["arn"],
    logging_configuration=aws.mwaa.EnvironmentLoggingConfigurationArgs(
        dag_processing_logs=aws.mwaa.EnvironmentLoggingConfigurationDagProcessingLogsArgs(
            enabled=True,
            log_level="DEBUG",
        ),
        scheduler_logs=aws.mwaa.EnvironmentLoggingConfigurationSchedulerLogsArgs(
            enabled=True,
            log_level="INFO",
        ),
        task_logs=aws.mwaa.EnvironmentLoggingConfigurationTaskLogsArgs(
            enabled=True,
            log_level="WARNING",
        ),
        webserver_logs=aws.mwaa.EnvironmentLoggingConfigurationWebserverLogsArgs(
            enabled=True,
            log_level="ERROR",
        ),
        worker_logs=aws.mwaa.EnvironmentLoggingConfigurationWorkerLogsArgs(
            enabled=True,
            log_level="CRITICAL",
        ),
    ),
    network_configuration=aws.mwaa.EnvironmentNetworkConfigurationArgs(
        security_group_ids=[aws_security_group["example"]["id"]],
        subnet_ids=[__item["id"] for __item in aws_subnet["private"]],
    ),
    source_bucket_arn=aws_s3_bucket["example"]["arn"])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Mwaa.Environment("example", new()
    {
        DagS3Path = "dags/",
        ExecutionRoleArn = aws_iam_role.Example.Arn,
        LoggingConfiguration = new Aws.Mwaa.Inputs.EnvironmentLoggingConfigurationArgs
        {
            DagProcessingLogs = new Aws.Mwaa.Inputs.EnvironmentLoggingConfigurationDagProcessingLogsArgs
            {
                Enabled = true,
                LogLevel = "DEBUG",
            },
            SchedulerLogs = new Aws.Mwaa.Inputs.EnvironmentLoggingConfigurationSchedulerLogsArgs
            {
                Enabled = true,
                LogLevel = "INFO",
            },
            TaskLogs = new Aws.Mwaa.Inputs.EnvironmentLoggingConfigurationTaskLogsArgs
            {
                Enabled = true,
                LogLevel = "WARNING",
            },
            WebserverLogs = new Aws.Mwaa.Inputs.EnvironmentLoggingConfigurationWebserverLogsArgs
            {
                Enabled = true,
                LogLevel = "ERROR",
            },
            WorkerLogs = new Aws.Mwaa.Inputs.EnvironmentLoggingConfigurationWorkerLogsArgs
            {
                Enabled = true,
                LogLevel = "CRITICAL",
            },
        },
        NetworkConfiguration = new Aws.Mwaa.Inputs.EnvironmentNetworkConfigurationArgs
        {
            SecurityGroupIds = new[]
            {
                aws_security_group.Example.Id,
            },
            SubnetIds = aws_subnet.Private.Select(__item => __item.Id).ToList(),
        },
        SourceBucketArn = aws_s3_bucket.Example.Arn,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mwaa"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := mwaa.NewEnvironment(ctx, "example", &mwaa.EnvironmentArgs{
DagS3Path: pulumi.String("dags/"),
ExecutionRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
LoggingConfiguration: &mwaa.EnvironmentLoggingConfigurationArgs{
DagProcessingLogs: &mwaa.EnvironmentLoggingConfigurationDagProcessingLogsArgs{
Enabled: pulumi.Bool(true),
LogLevel: pulumi.String("DEBUG"),
},
SchedulerLogs: &mwaa.EnvironmentLoggingConfigurationSchedulerLogsArgs{
Enabled: pulumi.Bool(true),
LogLevel: pulumi.String("INFO"),
},
TaskLogs: &mwaa.EnvironmentLoggingConfigurationTaskLogsArgs{
Enabled: pulumi.Bool(true),
LogLevel: pulumi.String("WARNING"),
},
WebserverLogs: &mwaa.EnvironmentLoggingConfigurationWebserverLogsArgs{
Enabled: pulumi.Bool(true),
LogLevel: pulumi.String("ERROR"),
},
WorkerLogs: &mwaa.EnvironmentLoggingConfigurationWorkerLogsArgs{
Enabled: pulumi.Bool(true),
LogLevel: pulumi.String("CRITICAL"),
},
},
NetworkConfiguration: &mwaa.EnvironmentNetworkConfigurationArgs{
SecurityGroupIds: pulumi.StringArray{
aws_security_group.Example.Id,
},
SubnetIds: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-resources-aws:mwaa-environment:Environment.pp:27,25-49),
},
SourceBucketArn: pulumi.Any(aws_s3_bucket.Example.Arn),
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
import com.pulumi.aws.mwaa.Environment;
import com.pulumi.aws.mwaa.EnvironmentArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationDagProcessingLogsArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationSchedulerLogsArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationTaskLogsArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationWebserverLogsArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentLoggingConfigurationWorkerLogsArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentNetworkConfigurationArgs;
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
        var example = new Environment("example", EnvironmentArgs.builder()        
            .dagS3Path("dags/")
            .executionRoleArn(aws_iam_role.example().arn())
            .loggingConfiguration(EnvironmentLoggingConfigurationArgs.builder()
                .dagProcessingLogs(EnvironmentLoggingConfigurationDagProcessingLogsArgs.builder()
                    .enabled(true)
                    .logLevel("DEBUG")
                    .build())
                .schedulerLogs(EnvironmentLoggingConfigurationSchedulerLogsArgs.builder()
                    .enabled(true)
                    .logLevel("INFO")
                    .build())
                .taskLogs(EnvironmentLoggingConfigurationTaskLogsArgs.builder()
                    .enabled(true)
                    .logLevel("WARNING")
                    .build())
                .webserverLogs(EnvironmentLoggingConfigurationWebserverLogsArgs.builder()
                    .enabled(true)
                    .logLevel("ERROR")
                    .build())
                .workerLogs(EnvironmentLoggingConfigurationWorkerLogsArgs.builder()
                    .enabled(true)
                    .logLevel("CRITICAL")
                    .build())
                .build())
            .networkConfiguration(EnvironmentNetworkConfigurationArgs.builder()
                .securityGroupIds(aws_security_group.example().id())
                .subnetIds(aws_subnet.private().stream().map(element -> element.id()).collect(toList()))
                .build())
            .sourceBucketArn(aws_s3_bucket.example().arn())
            .build());

    }
}
```
{{% /example %}}
{{% example %}}
### Example with tags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.mwaa.Environment("example", {
    dagS3Path: "dags/",
    executionRoleArn: aws_iam_role.example.arn,
    networkConfiguration: {
        securityGroupIds: [aws_security_group.example.id],
        subnetIds: aws_subnet["private"].map(__item => __item.id),
    },
    sourceBucketArn: aws_s3_bucket.example.arn,
    tags: {
        Name: "example",
        Environment: "production",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.mwaa.Environment("example",
    dag_s3_path="dags/",
    execution_role_arn=aws_iam_role["example"]["arn"],
    network_configuration=aws.mwaa.EnvironmentNetworkConfigurationArgs(
        security_group_ids=[aws_security_group["example"]["id"]],
        subnet_ids=[__item["id"] for __item in aws_subnet["private"]],
    ),
    source_bucket_arn=aws_s3_bucket["example"]["arn"],
    tags={
        "Name": "example",
        "Environment": "production",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.Mwaa.Environment("example", new()
    {
        DagS3Path = "dags/",
        ExecutionRoleArn = aws_iam_role.Example.Arn,
        NetworkConfiguration = new Aws.Mwaa.Inputs.EnvironmentNetworkConfigurationArgs
        {
            SecurityGroupIds = new[]
            {
                aws_security_group.Example.Id,
            },
            SubnetIds = aws_subnet.Private.Select(__item => __item.Id).ToList(),
        },
        SourceBucketArn = aws_s3_bucket.Example.Arn,
        Tags = 
        {
            { "Name", "example" },
            { "Environment", "production" },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/mwaa"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)
func main() {
pulumi.Run(func(ctx *pulumi.Context) error {
_, err := mwaa.NewEnvironment(ctx, "example", &mwaa.EnvironmentArgs{
DagS3Path: pulumi.String("dags/"),
ExecutionRoleArn: pulumi.Any(aws_iam_role.Example.Arn),
NetworkConfiguration: &mwaa.EnvironmentNetworkConfigurationArgs{
SecurityGroupIds: pulumi.StringArray{
aws_security_group.Example.Id,
},
SubnetIds: %!v(PANIC=Format method: fatal: A failure has occurred: unlowered splat expression @ #-resources-aws:mwaa-environment:Environment.pp:5,25-49),
},
SourceBucketArn: pulumi.Any(aws_s3_bucket.Example.Arn),
Tags: pulumi.StringMap{
"Name": pulumi.String("example"),
"Environment": pulumi.String("production"),
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
import com.pulumi.aws.mwaa.Environment;
import com.pulumi.aws.mwaa.EnvironmentArgs;
import com.pulumi.aws.mwaa.inputs.EnvironmentNetworkConfigurationArgs;
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
        var example = new Environment("example", EnvironmentArgs.builder()        
            .dagS3Path("dags/")
            .executionRoleArn(aws_iam_role.example().arn())
            .networkConfiguration(EnvironmentNetworkConfigurationArgs.builder()
                .securityGroupIds(aws_security_group.example().id())
                .subnetIds(aws_subnet.private().stream().map(element -> element.id()).collect(toList()))
                .build())
            .sourceBucketArn(aws_s3_bucket.example().arn())
            .tags(Map.ofEntries(
                Map.entry("Name", "example"),
                Map.entry("Environment", "production")
            ))
            .build());

    }
}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import MWAA Environment using `Name`. For example:

```sh
 $ pulumi import aws:mwaa/environment:Environment example MyAirflowEnvironment
```
 