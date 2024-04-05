Provides a Batch Job Definition resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Job definition of type container

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.batch.JobDefinition("test", {
    type: "container",
    containerProperties: JSON.stringify({
        command: [
            "ls",
            "-la",
        ],
        image: "busybox",
        resourceRequirements: [
            {
                type: "VCPU",
                value: "0.25",
            },
            {
                type: "MEMORY",
                value: "512",
            },
        ],
        volumes: [{
            host: {
                sourcePath: "/tmp",
            },
            name: "tmp",
        }],
        environment: [{
            name: "VARNAME",
            value: "VARVAL",
        }],
        mountPoints: [{
            sourceVolume: "tmp",
            containerPath: "/tmp",
            readOnly: false,
        }],
        ulimits: [{
            hardLimit: 1024,
            name: "nofile",
            softLimit: 1024,
        }],
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

test = aws.batch.JobDefinition("test",
    type="container",
    container_properties=json.dumps({
        "command": [
            "ls",
            "-la",
        ],
        "image": "busybox",
        "resourceRequirements": [
            {
                "type": "VCPU",
                "value": "0.25",
            },
            {
                "type": "MEMORY",
                "value": "512",
            },
        ],
        "volumes": [{
            "host": {
                "sourcePath": "/tmp",
            },
            "name": "tmp",
        }],
        "environment": [{
            "name": "VARNAME",
            "value": "VARVAL",
        }],
        "mountPoints": [{
            "sourceVolume": "tmp",
            "containerPath": "/tmp",
            "readOnly": False,
        }],
        "ulimits": [{
            "hardLimit": 1024,
            "name": "nofile",
            "softLimit": 1024,
        }],
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Batch.JobDefinition("test", new()
    {
        Type = "container",
        ContainerProperties = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["command"] = new[]
            {
                "ls",
                "-la",
            },
            ["image"] = "busybox",
            ["resourceRequirements"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["type"] = "VCPU",
                    ["value"] = "0.25",
                },
                new Dictionary<string, object?>
                {
                    ["type"] = "MEMORY",
                    ["value"] = "512",
                },
            },
            ["volumes"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["host"] = new Dictionary<string, object?>
                    {
                        ["sourcePath"] = "/tmp",
                    },
                    ["name"] = "tmp",
                },
            },
            ["environment"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["name"] = "VARNAME",
                    ["value"] = "VARVAL",
                },
            },
            ["mountPoints"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["sourceVolume"] = "tmp",
                    ["containerPath"] = "/tmp",
                    ["readOnly"] = false,
                },
            },
            ["ulimits"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["hardLimit"] = 1024,
                    ["name"] = "nofile",
                    ["softLimit"] = 1024,
                },
            },
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/batch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"command": []string{
				"ls",
				"-la",
			},
			"image": "busybox",
			"resourceRequirements": []map[string]interface{}{
				map[string]interface{}{
					"type":  "VCPU",
					"value": "0.25",
				},
				map[string]interface{}{
					"type":  "MEMORY",
					"value": "512",
				},
			},
			"volumes": []map[string]interface{}{
				map[string]interface{}{
					"host": map[string]interface{}{
						"sourcePath": "/tmp",
					},
					"name": "tmp",
				},
			},
			"environment": []map[string]interface{}{
				map[string]interface{}{
					"name":  "VARNAME",
					"value": "VARVAL",
				},
			},
			"mountPoints": []map[string]interface{}{
				map[string]interface{}{
					"sourceVolume":  "tmp",
					"containerPath": "/tmp",
					"readOnly":      false,
				},
			},
			"ulimits": []map[string]interface{}{
				map[string]interface{}{
					"hardLimit": 1024,
					"name":      "nofile",
					"softLimit": 1024,
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = batch.NewJobDefinition(ctx, "test", &batch.JobDefinitionArgs{
			Type:                pulumi.String("container"),
			ContainerProperties: pulumi.String(json0),
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
import com.pulumi.aws.batch.JobDefinition;
import com.pulumi.aws.batch.JobDefinitionArgs;
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
        var test = new JobDefinition("test", JobDefinitionArgs.builder()        
            .type("container")
            .containerProperties(serializeJson(
                jsonObject(
                    jsonProperty("command", jsonArray(
                        "ls", 
                        "-la"
                    )),
                    jsonProperty("image", "busybox"),
                    jsonProperty("resourceRequirements", jsonArray(
                        jsonObject(
                            jsonProperty("type", "VCPU"),
                            jsonProperty("value", "0.25")
                        ), 
                        jsonObject(
                            jsonProperty("type", "MEMORY"),
                            jsonProperty("value", "512")
                        )
                    )),
                    jsonProperty("volumes", jsonArray(jsonObject(
                        jsonProperty("host", jsonObject(
                            jsonProperty("sourcePath", "/tmp")
                        )),
                        jsonProperty("name", "tmp")
                    ))),
                    jsonProperty("environment", jsonArray(jsonObject(
                        jsonProperty("name", "VARNAME"),
                        jsonProperty("value", "VARVAL")
                    ))),
                    jsonProperty("mountPoints", jsonArray(jsonObject(
                        jsonProperty("sourceVolume", "tmp"),
                        jsonProperty("containerPath", "/tmp"),
                        jsonProperty("readOnly", false)
                    ))),
                    jsonProperty("ulimits", jsonArray(jsonObject(
                        jsonProperty("hardLimit", 1024),
                        jsonProperty("name", "nofile"),
                        jsonProperty("softLimit", 1024)
                    )))
                )))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:batch:JobDefinition
    properties:
      type: container
      containerProperties:
        fn::toJSON:
          command:
            - ls
            - -la
          image: busybox
          resourceRequirements:
            - type: VCPU
              value: '0.25'
            - type: MEMORY
              value: '512'
          volumes:
            - host:
                sourcePath: /tmp
              name: tmp
          environment:
            - name: VARNAME
              value: VARVAL
          mountPoints:
            - sourceVolume: tmp
              containerPath: /tmp
              readOnly: false
          ulimits:
            - hardLimit: 1024
              name: nofile
              softLimit: 1024
```
{{% /example %}}
{{% example %}}
### Job definition of type multinode

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.batch.JobDefinition("test", {
    type: "multinode",
    nodeProperties: JSON.stringify({
        mainNode: 0,
        nodeRangeProperties: [
            {
                container: {
                    command: [
                        "ls",
                        "-la",
                    ],
                    image: "busybox",
                    memory: 128,
                    vcpus: 1,
                },
                targetNodes: "0:",
            },
            {
                container: {
                    command: [
                        "echo",
                        "test",
                    ],
                    image: "busybox",
                    memory: 128,
                    vcpus: 1,
                },
                targetNodes: "1:",
            },
        ],
        numNodes: 2,
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

test = aws.batch.JobDefinition("test",
    type="multinode",
    node_properties=json.dumps({
        "mainNode": 0,
        "nodeRangeProperties": [
            {
                "container": {
                    "command": [
                        "ls",
                        "-la",
                    ],
                    "image": "busybox",
                    "memory": 128,
                    "vcpus": 1,
                },
                "targetNodes": "0:",
            },
            {
                "container": {
                    "command": [
                        "echo",
                        "test",
                    ],
                    "image": "busybox",
                    "memory": 128,
                    "vcpus": 1,
                },
                "targetNodes": "1:",
            },
        ],
        "numNodes": 2,
    }))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Batch.JobDefinition("test", new()
    {
        Type = "multinode",
        NodeProperties = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["mainNode"] = 0,
            ["nodeRangeProperties"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["container"] = new Dictionary<string, object?>
                    {
                        ["command"] = new[]
                        {
                            "ls",
                            "-la",
                        },
                        ["image"] = "busybox",
                        ["memory"] = 128,
                        ["vcpus"] = 1,
                    },
                    ["targetNodes"] = "0:",
                },
                new Dictionary<string, object?>
                {
                    ["container"] = new Dictionary<string, object?>
                    {
                        ["command"] = new[]
                        {
                            "echo",
                            "test",
                        },
                        ["image"] = "busybox",
                        ["memory"] = 128,
                        ["vcpus"] = 1,
                    },
                    ["targetNodes"] = "1:",
                },
            },
            ["numNodes"] = 2,
        }),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/batch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"mainNode": 0,
			"nodeRangeProperties": []map[string]interface{}{
				map[string]interface{}{
					"container": map[string]interface{}{
						"command": []string{
							"ls",
							"-la",
						},
						"image":  "busybox",
						"memory": 128,
						"vcpus":  1,
					},
					"targetNodes": "0:",
				},
				map[string]interface{}{
					"container": map[string]interface{}{
						"command": []string{
							"echo",
							"test",
						},
						"image":  "busybox",
						"memory": 128,
						"vcpus":  1,
					},
					"targetNodes": "1:",
				},
			},
			"numNodes": 2,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = batch.NewJobDefinition(ctx, "test", &batch.JobDefinitionArgs{
			Type:           pulumi.String("multinode"),
			NodeProperties: pulumi.String(json0),
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
import com.pulumi.aws.batch.JobDefinition;
import com.pulumi.aws.batch.JobDefinitionArgs;
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
        var test = new JobDefinition("test", JobDefinitionArgs.builder()        
            .type("multinode")
            .nodeProperties(serializeJson(
                jsonObject(
                    jsonProperty("mainNode", 0),
                    jsonProperty("nodeRangeProperties", jsonArray(
                        jsonObject(
                            jsonProperty("container", jsonObject(
                                jsonProperty("command", jsonArray(
                                    "ls", 
                                    "-la"
                                )),
                                jsonProperty("image", "busybox"),
                                jsonProperty("memory", 128),
                                jsonProperty("vcpus", 1)
                            )),
                            jsonProperty("targetNodes", "0:")
                        ), 
                        jsonObject(
                            jsonProperty("container", jsonObject(
                                jsonProperty("command", jsonArray(
                                    "echo", 
                                    "test"
                                )),
                                jsonProperty("image", "busybox"),
                                jsonProperty("memory", 128),
                                jsonProperty("vcpus", 1)
                            )),
                            jsonProperty("targetNodes", "1:")
                        )
                    )),
                    jsonProperty("numNodes", 2)
                )))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:batch:JobDefinition
    properties:
      type: multinode
      nodeProperties:
        fn::toJSON:
          mainNode: 0
          nodeRangeProperties:
            - container:
                command:
                  - ls
                  - -la
                image: busybox
                memory: 128
                vcpus: 1
              targetNodes: '0:'
            - container:
                command:
                  - echo
                  - test
                image: busybox
                memory: 128
                vcpus: 1
              targetNodes: '1:'
          numNodes: 2
```
{{% /example %}}
{{% example %}}
### Fargate Platform Capability

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const assumeRolePolicy = aws.iam.getPolicyDocument({
    statements: [{
        actions: ["sts:AssumeRole"],
        principals: [{
            type: "Service",
            identifiers: ["ecs-tasks.amazonaws.com"],
        }],
    }],
});
const ecsTaskExecutionRole = new aws.iam.Role("ecsTaskExecutionRole", {assumeRolePolicy: assumeRolePolicy.then(assumeRolePolicy => assumeRolePolicy.json)});
const ecsTaskExecutionRolePolicy = new aws.iam.RolePolicyAttachment("ecsTaskExecutionRolePolicy", {
    role: ecsTaskExecutionRole.name,
    policyArn: "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
});
const test = new aws.batch.JobDefinition("test", {
    type: "container",
    platformCapabilities: ["FARGATE"],
    containerProperties: ecsTaskExecutionRole.arn.apply(arn => JSON.stringify({
        command: [
            "echo",
            "test",
        ],
        image: "busybox",
        jobRoleArn: "arn:aws:iam::123456789012:role/AWSBatchS3ReadOnly",
        fargatePlatformConfiguration: {
            platformVersion: "LATEST",
        },
        resourceRequirements: [
            {
                type: "VCPU",
                value: "0.25",
            },
            {
                type: "MEMORY",
                value: "512",
            },
        ],
        executionRoleArn: arn,
    })),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

assume_role_policy = aws.iam.get_policy_document(statements=[aws.iam.GetPolicyDocumentStatementArgs(
    actions=["sts:AssumeRole"],
    principals=[aws.iam.GetPolicyDocumentStatementPrincipalArgs(
        type="Service",
        identifiers=["ecs-tasks.amazonaws.com"],
    )],
)])
ecs_task_execution_role = aws.iam.Role("ecsTaskExecutionRole", assume_role_policy=assume_role_policy.json)
ecs_task_execution_role_policy = aws.iam.RolePolicyAttachment("ecsTaskExecutionRolePolicy",
    role=ecs_task_execution_role.name,
    policy_arn="arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy")
test = aws.batch.JobDefinition("test",
    type="container",
    platform_capabilities=["FARGATE"],
    container_properties=ecs_task_execution_role.arn.apply(lambda arn: json.dumps({
        "command": [
            "echo",
            "test",
        ],
        "image": "busybox",
        "jobRoleArn": "arn:aws:iam::123456789012:role/AWSBatchS3ReadOnly",
        "fargatePlatformConfiguration": {
            "platformVersion": "LATEST",
        },
        "resourceRequirements": [
            {
                "type": "VCPU",
                "value": "0.25",
            },
            {
                "type": "MEMORY",
                "value": "512",
            },
        ],
        "executionRoleArn": arn,
    })))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var assumeRolePolicy = Aws.Iam.GetPolicyDocument.Invoke(new()
    {
        Statements = new[]
        {
            new Aws.Iam.Inputs.GetPolicyDocumentStatementInputArgs
            {
                Actions = new[]
                {
                    "sts:AssumeRole",
                },
                Principals = new[]
                {
                    new Aws.Iam.Inputs.GetPolicyDocumentStatementPrincipalInputArgs
                    {
                        Type = "Service",
                        Identifiers = new[]
                        {
                            "ecs-tasks.amazonaws.com",
                        },
                    },
                },
            },
        },
    });

    var ecsTaskExecutionRole = new Aws.Iam.Role("ecsTaskExecutionRole", new()
    {
        AssumeRolePolicy = assumeRolePolicy.Apply(getPolicyDocumentResult => getPolicyDocumentResult.Json),
    });

    var ecsTaskExecutionRolePolicy = new Aws.Iam.RolePolicyAttachment("ecsTaskExecutionRolePolicy", new()
    {
        Role = ecsTaskExecutionRole.Name,
        PolicyArn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy",
    });

    var test = new Aws.Batch.JobDefinition("test", new()
    {
        Type = "container",
        PlatformCapabilities = new[]
        {
            "FARGATE",
        },
        ContainerProperties = ecsTaskExecutionRole.Arn.Apply(arn => JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["command"] = new[]
            {
                "echo",
                "test",
            },
            ["image"] = "busybox",
            ["jobRoleArn"] = "arn:aws:iam::123456789012:role/AWSBatchS3ReadOnly",
            ["fargatePlatformConfiguration"] = new Dictionary<string, object?>
            {
                ["platformVersion"] = "LATEST",
            },
            ["resourceRequirements"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["type"] = "VCPU",
                    ["value"] = "0.25",
                },
                new Dictionary<string, object?>
                {
                    ["type"] = "MEMORY",
                    ["value"] = "512",
                },
            },
            ["executionRoleArn"] = arn,
        })),
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/batch"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		assumeRolePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{
						"sts:AssumeRole",
					},
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type: "Service",
							Identifiers: []string{
								"ecs-tasks.amazonaws.com",
							},
						},
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}
		ecsTaskExecutionRole, err := iam.NewRole(ctx, "ecsTaskExecutionRole", &iam.RoleArgs{
			AssumeRolePolicy: *pulumi.String(assumeRolePolicy.Json),
		})
		if err != nil {
			return err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "ecsTaskExecutionRolePolicy", &iam.RolePolicyAttachmentArgs{
			Role:      ecsTaskExecutionRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"),
		})
		if err != nil {
			return err
		}
		_, err = batch.NewJobDefinition(ctx, "test", &batch.JobDefinitionArgs{
			Type: pulumi.String("container"),
			PlatformCapabilities: pulumi.StringArray{
				pulumi.String("FARGATE"),
			},
			ContainerProperties: ecsTaskExecutionRole.Arn.ApplyT(func(arn string) (pulumi.String, error) {
				var _zero pulumi.String
				tmpJSON0, err := json.Marshal(map[string]interface{}{
					"command": []string{
						"echo",
						"test",
					},
					"image":      "busybox",
					"jobRoleArn": "arn:aws:iam::123456789012:role/AWSBatchS3ReadOnly",
					"fargatePlatformConfiguration": map[string]interface{}{
						"platformVersion": "LATEST",
					},
					"resourceRequirements": []map[string]interface{}{
						map[string]interface{}{
							"type":  "VCPU",
							"value": "0.25",
						},
						map[string]interface{}{
							"type":  "MEMORY",
							"value": "512",
						},
					},
					"executionRoleArn": arn,
				})
				if err != nil {
					return _zero, err
				}
				json0 := string(tmpJSON0)
				return pulumi.String(json0), nil
			}).(pulumi.StringOutput),
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
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iam.RolePolicyAttachment;
import com.pulumi.aws.iam.RolePolicyAttachmentArgs;
import com.pulumi.aws.batch.JobDefinition;
import com.pulumi.aws.batch.JobDefinitionArgs;
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
        final var assumeRolePolicy = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .statements(GetPolicyDocumentStatementArgs.builder()
                .actions("sts:AssumeRole")
                .principals(GetPolicyDocumentStatementPrincipalArgs.builder()
                    .type("Service")
                    .identifiers("ecs-tasks.amazonaws.com")
                    .build())
                .build())
            .build());

        var ecsTaskExecutionRole = new Role("ecsTaskExecutionRole", RoleArgs.builder()        
            .assumeRolePolicy(assumeRolePolicy.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var ecsTaskExecutionRolePolicy = new RolePolicyAttachment("ecsTaskExecutionRolePolicy", RolePolicyAttachmentArgs.builder()        
            .role(ecsTaskExecutionRole.name())
            .policyArn("arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy")
            .build());

        var test = new JobDefinition("test", JobDefinitionArgs.builder()        
            .type("container")
            .platformCapabilities("FARGATE")
            .containerProperties(ecsTaskExecutionRole.arn().applyValue(arn -> serializeJson(
                jsonObject(
                    jsonProperty("command", jsonArray(
                        "echo", 
                        "test"
                    )),
                    jsonProperty("image", "busybox"),
                    jsonProperty("jobRoleArn", "arn:aws:iam::123456789012:role/AWSBatchS3ReadOnly"),
                    jsonProperty("fargatePlatformConfiguration", jsonObject(
                        jsonProperty("platformVersion", "LATEST")
                    )),
                    jsonProperty("resourceRequirements", jsonArray(
                        jsonObject(
                            jsonProperty("type", "VCPU"),
                            jsonProperty("value", "0.25")
                        ), 
                        jsonObject(
                            jsonProperty("type", "MEMORY"),
                            jsonProperty("value", "512")
                        )
                    )),
                    jsonProperty("executionRoleArn", arn)
                ))))
            .build());

    }
}
```
```yaml
resources:
  ecsTaskExecutionRole:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRolePolicy.json}
  ecsTaskExecutionRolePolicy:
    type: aws:iam:RolePolicyAttachment
    properties:
      role: ${ecsTaskExecutionRole.name}
      policyArn: arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy
  test:
    type: aws:batch:JobDefinition
    properties:
      type: container
      platformCapabilities:
        - FARGATE
      containerProperties:
        fn::toJSON:
          command:
            - echo
            - test
          image: busybox
          jobRoleArn: arn:aws:iam::123456789012:role/AWSBatchS3ReadOnly
          fargatePlatformConfiguration:
            platformVersion: LATEST
          resourceRequirements:
            - type: VCPU
              value: '0.25'
            - type: MEMORY
              value: '512'
          executionRoleArn: ${ecsTaskExecutionRole.arn}
variables:
  assumeRolePolicy:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        statements:
          - actions:
              - sts:AssumeRole
            principals:
              - type: Service
                identifiers:
                  - ecs-tasks.amazonaws.com
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Batch Job Definition using the `arn`. For example:

```sh
 $ pulumi import aws:batch/jobDefinition:JobDefinition test arn:aws:batch:us-east-1:123456789012:job-definition/sample
```
 