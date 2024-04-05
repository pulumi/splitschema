Manages a revision of an ECS task definition to be used in `aws.ecs.Service`.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const service = new aws.ecs.TaskDefinition("service", {
    family: "service",
    containerDefinitions: JSON.stringify([
        {
            name: "first",
            image: "service-first",
            cpu: 10,
            memory: 512,
            essential: true,
            portMappings: [{
                containerPort: 80,
                hostPort: 80,
            }],
        },
        {
            name: "second",
            image: "service-second",
            cpu: 10,
            memory: 256,
            essential: true,
            portMappings: [{
                containerPort: 443,
                hostPort: 443,
            }],
        },
    ]),
    volumes: [{
        name: "service-storage",
        hostPath: "/ecs/service-storage",
    }],
    placementConstraints: [{
        type: "memberOf",
        expression: "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]",
    }],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

service = aws.ecs.TaskDefinition("service",
    family="service",
    container_definitions=json.dumps([
        {
            "name": "first",
            "image": "service-first",
            "cpu": 10,
            "memory": 512,
            "essential": True,
            "portMappings": [{
                "containerPort": 80,
                "hostPort": 80,
            }],
        },
        {
            "name": "second",
            "image": "service-second",
            "cpu": 10,
            "memory": 256,
            "essential": True,
            "portMappings": [{
                "containerPort": 443,
                "hostPort": 443,
            }],
        },
    ]),
    volumes=[aws.ecs.TaskDefinitionVolumeArgs(
        name="service-storage",
        host_path="/ecs/service-storage",
    )],
    placement_constraints=[aws.ecs.TaskDefinitionPlacementConstraintArgs(
        type="memberOf",
        expression="attribute:ecs.availability-zone in [us-west-2a, us-west-2b]",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var service = new Aws.Ecs.TaskDefinition("service", new()
    {
        Family = "service",
        ContainerDefinitions = JsonSerializer.Serialize(new[]
        {
            new Dictionary<string, object?>
            {
                ["name"] = "first",
                ["image"] = "service-first",
                ["cpu"] = 10,
                ["memory"] = 512,
                ["essential"] = true,
                ["portMappings"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["containerPort"] = 80,
                        ["hostPort"] = 80,
                    },
                },
            },
            new Dictionary<string, object?>
            {
                ["name"] = "second",
                ["image"] = "service-second",
                ["cpu"] = 10,
                ["memory"] = 256,
                ["essential"] = true,
                ["portMappings"] = new[]
                {
                    new Dictionary<string, object?>
                    {
                        ["containerPort"] = 443,
                        ["hostPort"] = 443,
                    },
                },
            },
        }),
        Volumes = new[]
        {
            new Aws.Ecs.Inputs.TaskDefinitionVolumeArgs
            {
                Name = "service-storage",
                HostPath = "/ecs/service-storage",
            },
        },
        PlacementConstraints = new[]
        {
            new Aws.Ecs.Inputs.TaskDefinitionPlacementConstraintArgs
            {
                Type = "memberOf",
                Expression = "attribute:ecs.availability-zone in [us-west-2a, us-west-2b]",
            },
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal([]interface{}{
			map[string]interface{}{
				"name":      "first",
				"image":     "service-first",
				"cpu":       10,
				"memory":    512,
				"essential": true,
				"portMappings": []map[string]interface{}{
					map[string]interface{}{
						"containerPort": 80,
						"hostPort":      80,
					},
				},
			},
			map[string]interface{}{
				"name":      "second",
				"image":     "service-second",
				"cpu":       10,
				"memory":    256,
				"essential": true,
				"portMappings": []map[string]interface{}{
					map[string]interface{}{
						"containerPort": 443,
						"hostPort":      443,
					},
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = ecs.NewTaskDefinition(ctx, "service", &ecs.TaskDefinitionArgs{
			Family:               pulumi.String("service"),
			ContainerDefinitions: pulumi.String(json0),
			Volumes: ecs.TaskDefinitionVolumeArray{
				&ecs.TaskDefinitionVolumeArgs{
					Name:     pulumi.String("service-storage"),
					HostPath: pulumi.String("/ecs/service-storage"),
				},
			},
			PlacementConstraints: ecs.TaskDefinitionPlacementConstraintArray{
				&ecs.TaskDefinitionPlacementConstraintArgs{
					Type:       pulumi.String("memberOf"),
					Expression: pulumi.String("attribute:ecs.availability-zone in [us-west-2a, us-west-2b]"),
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
import com.pulumi.aws.ecs.TaskDefinition;
import com.pulumi.aws.ecs.TaskDefinitionArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionVolumeArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionPlacementConstraintArgs;
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
        var service = new TaskDefinition("service", TaskDefinitionArgs.builder()        
            .family("service")
            .containerDefinitions(serializeJson(
                jsonArray(
                    jsonObject(
                        jsonProperty("name", "first"),
                        jsonProperty("image", "service-first"),
                        jsonProperty("cpu", 10),
                        jsonProperty("memory", 512),
                        jsonProperty("essential", true),
                        jsonProperty("portMappings", jsonArray(jsonObject(
                            jsonProperty("containerPort", 80),
                            jsonProperty("hostPort", 80)
                        )))
                    ), 
                    jsonObject(
                        jsonProperty("name", "second"),
                        jsonProperty("image", "service-second"),
                        jsonProperty("cpu", 10),
                        jsonProperty("memory", 256),
                        jsonProperty("essential", true),
                        jsonProperty("portMappings", jsonArray(jsonObject(
                            jsonProperty("containerPort", 443),
                            jsonProperty("hostPort", 443)
                        )))
                    )
                )))
            .volumes(TaskDefinitionVolumeArgs.builder()
                .name("service-storage")
                .hostPath("/ecs/service-storage")
                .build())
            .placementConstraints(TaskDefinitionPlacementConstraintArgs.builder()
                .type("memberOf")
                .expression("attribute:ecs.availability-zone in [us-west-2a, us-west-2b]")
                .build())
            .build());

    }
}
```
```yaml
resources:
  service:
    type: aws:ecs:TaskDefinition
    properties:
      family: service
      containerDefinitions:
        fn::toJSON:
          - name: first
            image: service-first
            cpu: 10
            memory: 512
            essential: true
            portMappings:
              - containerPort: 80
                hostPort: 80
          - name: second
            image: service-second
            cpu: 10
            memory: 256
            essential: true
            portMappings:
              - containerPort: 443
                hostPort: 443
      volumes:
        - name: service-storage
          hostPath: /ecs/service-storage
      placementConstraints:
        - type: memberOf
          expression: attribute:ecs.availability-zone in [us-west-2a, us-west-2b]
```
{{% /example %}}
{{% example %}}
### With AppMesh Proxy

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const service = new aws.ecs.TaskDefinition("service", {
    family: "service",
    containerDefinitions: fs.readFileSync("task-definitions/service.json", "utf8"),
    proxyConfiguration: {
        type: "APPMESH",
        containerName: "applicationContainerName",
        properties: {
            AppPorts: "8080",
            EgressIgnoredIPs: "169.254.170.2,169.254.169.254",
            IgnoredUID: "1337",
            ProxyEgressPort: "15001",
            ProxyIngressPort: "15000",
        },
    },
});
```
```python
import pulumi
import pulumi_aws as aws

service = aws.ecs.TaskDefinition("service",
    family="service",
    container_definitions=(lambda path: open(path).read())("task-definitions/service.json"),
    proxy_configuration=aws.ecs.TaskDefinitionProxyConfigurationArgs(
        type="APPMESH",
        container_name="applicationContainerName",
        properties={
            "AppPorts": "8080",
            "EgressIgnoredIPs": "169.254.170.2,169.254.169.254",
            "IgnoredUID": "1337",
            "ProxyEgressPort": "15001",
            "ProxyIngressPort": "15000",
        },
    ))
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var service = new Aws.Ecs.TaskDefinition("service", new()
    {
        Family = "service",
        ContainerDefinitions = File.ReadAllText("task-definitions/service.json"),
        ProxyConfiguration = new Aws.Ecs.Inputs.TaskDefinitionProxyConfigurationArgs
        {
            Type = "APPMESH",
            ContainerName = "applicationContainerName",
            Properties = 
            {
                { "AppPorts", "8080" },
                { "EgressIgnoredIPs", "169.254.170.2,169.254.169.254" },
                { "IgnoredUID", "1337" },
                { "ProxyEgressPort", "15001" },
                { "ProxyIngressPort", "15000" },
            },
        },
    });

});
```
```go
package main

import (
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.NewTaskDefinition(ctx, "service", &ecs.TaskDefinitionArgs{
			Family:               pulumi.String("service"),
			ContainerDefinitions: readFileOrPanic("task-definitions/service.json"),
			ProxyConfiguration: &ecs.TaskDefinitionProxyConfigurationArgs{
				Type:          pulumi.String("APPMESH"),
				ContainerName: pulumi.String("applicationContainerName"),
				Properties: pulumi.StringMap{
					"AppPorts":         pulumi.String("8080"),
					"EgressIgnoredIPs": pulumi.String("169.254.170.2,169.254.169.254"),
					"IgnoredUID":       pulumi.String("1337"),
					"ProxyEgressPort":  pulumi.String("15001"),
					"ProxyIngressPort": pulumi.String("15000"),
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
import com.pulumi.aws.ecs.TaskDefinition;
import com.pulumi.aws.ecs.TaskDefinitionArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionProxyConfigurationArgs;
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
        var service = new TaskDefinition("service", TaskDefinitionArgs.builder()        
            .family("service")
            .containerDefinitions(Files.readString(Paths.get("task-definitions/service.json")))
            .proxyConfiguration(TaskDefinitionProxyConfigurationArgs.builder()
                .type("APPMESH")
                .containerName("applicationContainerName")
                .properties(Map.ofEntries(
                    Map.entry("AppPorts", "8080"),
                    Map.entry("EgressIgnoredIPs", "169.254.170.2,169.254.169.254"),
                    Map.entry("IgnoredUID", "1337"),
                    Map.entry("ProxyEgressPort", 15001),
                    Map.entry("ProxyIngressPort", 15000)
                ))
                .build())
            .build());

    }
}
```
```yaml
resources:
  service:
    type: aws:ecs:TaskDefinition
    properties:
      family: service
      containerDefinitions:
        fn::readFile: task-definitions/service.json
      proxyConfiguration:
        type: APPMESH
        containerName: applicationContainerName
        properties:
          AppPorts: '8080'
          EgressIgnoredIPs: 169.254.170.2,169.254.169.254
          IgnoredUID: '1337'
          ProxyEgressPort: 15001
          ProxyIngressPort: 15000
```
{{% /example %}}
{{% example %}}
### Example Using `docker_volume_configuration`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const service = new aws.ecs.TaskDefinition("service", {
    family: "service",
    containerDefinitions: fs.readFileSync("task-definitions/service.json", "utf8"),
    volumes: [{
        name: "service-storage",
        dockerVolumeConfiguration: {
            scope: "shared",
            autoprovision: true,
            driver: "local",
            driverOpts: {
                type: "nfs",
                device: `${aws_efs_file_system.fs.dns_name}:/`,
                o: `addr=${aws_efs_file_system.fs.dns_name},rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport`,
            },
        },
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

service = aws.ecs.TaskDefinition("service",
    family="service",
    container_definitions=(lambda path: open(path).read())("task-definitions/service.json"),
    volumes=[aws.ecs.TaskDefinitionVolumeArgs(
        name="service-storage",
        docker_volume_configuration=aws.ecs.TaskDefinitionVolumeDockerVolumeConfigurationArgs(
            scope="shared",
            autoprovision=True,
            driver="local",
            driver_opts={
                "type": "nfs",
                "device": f"{aws_efs_file_system['fs']['dns_name']}:/",
                "o": f"addr={aws_efs_file_system['fs']['dns_name']},rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport",
            },
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var service = new Aws.Ecs.TaskDefinition("service", new()
    {
        Family = "service",
        ContainerDefinitions = File.ReadAllText("task-definitions/service.json"),
        Volumes = new[]
        {
            new Aws.Ecs.Inputs.TaskDefinitionVolumeArgs
            {
                Name = "service-storage",
                DockerVolumeConfiguration = new Aws.Ecs.Inputs.TaskDefinitionVolumeDockerVolumeConfigurationArgs
                {
                    Scope = "shared",
                    Autoprovision = true,
                    Driver = "local",
                    DriverOpts = 
                    {
                        { "type", "nfs" },
                        { "device", $"{aws_efs_file_system.Fs.Dns_name}:/" },
                        { "o", $"addr={aws_efs_file_system.Fs.Dns_name},rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport" },
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
	"fmt"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.NewTaskDefinition(ctx, "service", &ecs.TaskDefinitionArgs{
			Family:               pulumi.String("service"),
			ContainerDefinitions: readFileOrPanic("task-definitions/service.json"),
			Volumes: ecs.TaskDefinitionVolumeArray{
				&ecs.TaskDefinitionVolumeArgs{
					Name: pulumi.String("service-storage"),
					DockerVolumeConfiguration: &ecs.TaskDefinitionVolumeDockerVolumeConfigurationArgs{
						Scope:         pulumi.String("shared"),
						Autoprovision: pulumi.Bool(true),
						Driver:        pulumi.String("local"),
						DriverOpts: pulumi.StringMap{
							"type":   pulumi.String("nfs"),
							"device": pulumi.String(fmt.Sprintf("%v:/", aws_efs_file_system.Fs.Dns_name)),
							"o":      pulumi.String(fmt.Sprintf("addr=%v,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport", aws_efs_file_system.Fs.Dns_name)),
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
import com.pulumi.aws.ecs.TaskDefinition;
import com.pulumi.aws.ecs.TaskDefinitionArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionVolumeArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionVolumeDockerVolumeConfigurationArgs;
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
        var service = new TaskDefinition("service", TaskDefinitionArgs.builder()        
            .family("service")
            .containerDefinitions(Files.readString(Paths.get("task-definitions/service.json")))
            .volumes(TaskDefinitionVolumeArgs.builder()
                .name("service-storage")
                .dockerVolumeConfiguration(TaskDefinitionVolumeDockerVolumeConfigurationArgs.builder()
                    .scope("shared")
                    .autoprovision(true)
                    .driver("local")
                    .driverOpts(Map.ofEntries(
                        Map.entry("type", "nfs"),
                        Map.entry("device", String.format("%s:/", aws_efs_file_system.fs().dns_name())),
                        Map.entry("o", String.format("addr=%s,rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport", aws_efs_file_system.fs().dns_name()))
                    ))
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  service:
    type: aws:ecs:TaskDefinition
    properties:
      family: service
      containerDefinitions:
        fn::readFile: task-definitions/service.json
      volumes:
        - name: service-storage
          dockerVolumeConfiguration:
            scope: shared
            autoprovision: true
            driver: local
            driverOpts:
              type: nfs
              device: ${aws_efs_file_system.fs.dns_name}:/
              o: addr=${aws_efs_file_system.fs.dns_name},rsize=1048576,wsize=1048576,hard,timeo=600,retrans=2,noresvport
```
{{% /example %}}
{{% example %}}
### Example Using `efs_volume_configuration`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const service = new aws.ecs.TaskDefinition("service", {
    family: "service",
    containerDefinitions: fs.readFileSync("task-definitions/service.json", "utf8"),
    volumes: [{
        name: "service-storage",
        efsVolumeConfiguration: {
            fileSystemId: aws_efs_file_system.fs.id,
            rootDirectory: "/opt/data",
            transitEncryption: "ENABLED",
            transitEncryptionPort: 2999,
            authorizationConfig: {
                accessPointId: aws_efs_access_point.test.id,
                iam: "ENABLED",
            },
        },
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

service = aws.ecs.TaskDefinition("service",
    family="service",
    container_definitions=(lambda path: open(path).read())("task-definitions/service.json"),
    volumes=[aws.ecs.TaskDefinitionVolumeArgs(
        name="service-storage",
        efs_volume_configuration=aws.ecs.TaskDefinitionVolumeEfsVolumeConfigurationArgs(
            file_system_id=aws_efs_file_system["fs"]["id"],
            root_directory="/opt/data",
            transit_encryption="ENABLED",
            transit_encryption_port=2999,
            authorization_config=aws.ecs.TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigArgs(
                access_point_id=aws_efs_access_point["test"]["id"],
                iam="ENABLED",
            ),
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var service = new Aws.Ecs.TaskDefinition("service", new()
    {
        Family = "service",
        ContainerDefinitions = File.ReadAllText("task-definitions/service.json"),
        Volumes = new[]
        {
            new Aws.Ecs.Inputs.TaskDefinitionVolumeArgs
            {
                Name = "service-storage",
                EfsVolumeConfiguration = new Aws.Ecs.Inputs.TaskDefinitionVolumeEfsVolumeConfigurationArgs
                {
                    FileSystemId = aws_efs_file_system.Fs.Id,
                    RootDirectory = "/opt/data",
                    TransitEncryption = "ENABLED",
                    TransitEncryptionPort = 2999,
                    AuthorizationConfig = new Aws.Ecs.Inputs.TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigArgs
                    {
                        AccessPointId = aws_efs_access_point.Test.Id,
                        Iam = "ENABLED",
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
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.NewTaskDefinition(ctx, "service", &ecs.TaskDefinitionArgs{
			Family:               pulumi.String("service"),
			ContainerDefinitions: readFileOrPanic("task-definitions/service.json"),
			Volumes: ecs.TaskDefinitionVolumeArray{
				&ecs.TaskDefinitionVolumeArgs{
					Name: pulumi.String("service-storage"),
					EfsVolumeConfiguration: &ecs.TaskDefinitionVolumeEfsVolumeConfigurationArgs{
						FileSystemId:          pulumi.Any(aws_efs_file_system.Fs.Id),
						RootDirectory:         pulumi.String("/opt/data"),
						TransitEncryption:     pulumi.String("ENABLED"),
						TransitEncryptionPort: pulumi.Int(2999),
						AuthorizationConfig: &ecs.TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigArgs{
							AccessPointId: pulumi.Any(aws_efs_access_point.Test.Id),
							Iam:           pulumi.String("ENABLED"),
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
import com.pulumi.aws.ecs.TaskDefinition;
import com.pulumi.aws.ecs.TaskDefinitionArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionVolumeArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionVolumeEfsVolumeConfigurationArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigArgs;
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
        var service = new TaskDefinition("service", TaskDefinitionArgs.builder()        
            .family("service")
            .containerDefinitions(Files.readString(Paths.get("task-definitions/service.json")))
            .volumes(TaskDefinitionVolumeArgs.builder()
                .name("service-storage")
                .efsVolumeConfiguration(TaskDefinitionVolumeEfsVolumeConfigurationArgs.builder()
                    .fileSystemId(aws_efs_file_system.fs().id())
                    .rootDirectory("/opt/data")
                    .transitEncryption("ENABLED")
                    .transitEncryptionPort(2999)
                    .authorizationConfig(TaskDefinitionVolumeEfsVolumeConfigurationAuthorizationConfigArgs.builder()
                        .accessPointId(aws_efs_access_point.test().id())
                        .iam("ENABLED")
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  service:
    type: aws:ecs:TaskDefinition
    properties:
      family: service
      containerDefinitions:
        fn::readFile: task-definitions/service.json
      volumes:
        - name: service-storage
          efsVolumeConfiguration:
            fileSystemId: ${aws_efs_file_system.fs.id}
            rootDirectory: /opt/data
            transitEncryption: ENABLED
            transitEncryptionPort: 2999
            authorizationConfig:
              accessPointId: ${aws_efs_access_point.test.id}
              iam: ENABLED
```
{{% /example %}}
{{% example %}}
### Example Using `fsx_windows_file_server_volume_configuration`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as fs from "fs";

const test = new aws.secretsmanager.SecretVersion("test", {
    secretId: aws_secretsmanager_secret.test.id,
    secretString: JSON.stringify({
        username: "admin",
        password: aws_directory_service_directory.test.password,
    }),
});
const service = new aws.ecs.TaskDefinition("service", {
    family: "service",
    containerDefinitions: fs.readFileSync("task-definitions/service.json", "utf8"),
    volumes: [{
        name: "service-storage",
        fsxWindowsFileServerVolumeConfiguration: {
            fileSystemId: aws_fsx_windows_file_system.test.id,
            rootDirectory: "\\data",
            authorizationConfig: {
                credentialsParameter: test.arn,
                domain: aws_directory_service_directory.test.name,
            },
        },
    }],
});
```
```python
import pulumi
import json
import pulumi_aws as aws

test = aws.secretsmanager.SecretVersion("test",
    secret_id=aws_secretsmanager_secret["test"]["id"],
    secret_string=json.dumps({
        "username": "admin",
        "password": aws_directory_service_directory["test"]["password"],
    }))
service = aws.ecs.TaskDefinition("service",
    family="service",
    container_definitions=(lambda path: open(path).read())("task-definitions/service.json"),
    volumes=[aws.ecs.TaskDefinitionVolumeArgs(
        name="service-storage",
        fsx_windows_file_server_volume_configuration=aws.ecs.TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationArgs(
            file_system_id=aws_fsx_windows_file_system["test"]["id"],
            root_directory="\\data",
            authorization_config=aws.ecs.TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigArgs(
                credentials_parameter=test.arn,
                domain=aws_directory_service_directory["test"]["name"],
            ),
        ),
    )])
```
```csharp
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.SecretsManager.SecretVersion("test", new()
    {
        SecretId = aws_secretsmanager_secret.Test.Id,
        SecretString = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["username"] = "admin",
            ["password"] = aws_directory_service_directory.Test.Password,
        }),
    });

    var service = new Aws.Ecs.TaskDefinition("service", new()
    {
        Family = "service",
        ContainerDefinitions = File.ReadAllText("task-definitions/service.json"),
        Volumes = new[]
        {
            new Aws.Ecs.Inputs.TaskDefinitionVolumeArgs
            {
                Name = "service-storage",
                FsxWindowsFileServerVolumeConfiguration = new Aws.Ecs.Inputs.TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationArgs
                {
                    FileSystemId = aws_fsx_windows_file_system.Test.Id,
                    RootDirectory = "\\data",
                    AuthorizationConfig = new Aws.Ecs.Inputs.TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigArgs
                    {
                        CredentialsParameter = test.Arn,
                        Domain = aws_directory_service_directory.Test.Name,
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
	"encoding/json"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/secretsmanager"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func readFileOrPanic(path string) pulumi.StringPtrInput {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return pulumi.String(string(data))
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"username": "admin",
			"password": aws_directory_service_directory.Test.Password,
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		test, err := secretsmanager.NewSecretVersion(ctx, "test", &secretsmanager.SecretVersionArgs{
			SecretId:     pulumi.Any(aws_secretsmanager_secret.Test.Id),
			SecretString: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = ecs.NewTaskDefinition(ctx, "service", &ecs.TaskDefinitionArgs{
			Family:               pulumi.String("service"),
			ContainerDefinitions: readFileOrPanic("task-definitions/service.json"),
			Volumes: ecs.TaskDefinitionVolumeArray{
				&ecs.TaskDefinitionVolumeArgs{
					Name: pulumi.String("service-storage"),
					FsxWindowsFileServerVolumeConfiguration: &ecs.TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationArgs{
						FileSystemId:  pulumi.Any(aws_fsx_windows_file_system.Test.Id),
						RootDirectory: pulumi.String("\\data"),
						AuthorizationConfig: &ecs.TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigArgs{
							CredentialsParameter: test.Arn,
							Domain:               pulumi.Any(aws_directory_service_directory.Test.Name),
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
import com.pulumi.aws.secretsmanager.SecretVersion;
import com.pulumi.aws.secretsmanager.SecretVersionArgs;
import com.pulumi.aws.ecs.TaskDefinition;
import com.pulumi.aws.ecs.TaskDefinitionArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionVolumeArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigArgs;
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
        var test = new SecretVersion("test", SecretVersionArgs.builder()        
            .secretId(aws_secretsmanager_secret.test().id())
            .secretString(serializeJson(
                jsonObject(
                    jsonProperty("username", "admin"),
                    jsonProperty("password", aws_directory_service_directory.test().password())
                )))
            .build());

        var service = new TaskDefinition("service", TaskDefinitionArgs.builder()        
            .family("service")
            .containerDefinitions(Files.readString(Paths.get("task-definitions/service.json")))
            .volumes(TaskDefinitionVolumeArgs.builder()
                .name("service-storage")
                .fsxWindowsFileServerVolumeConfiguration(TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationArgs.builder()
                    .fileSystemId(aws_fsx_windows_file_system.test().id())
                    .rootDirectory("\\data")
                    .authorizationConfig(TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfigArgs.builder()
                        .credentialsParameter(test.arn())
                        .domain(aws_directory_service_directory.test().name())
                        .build())
                    .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  service:
    type: aws:ecs:TaskDefinition
    properties:
      family: service
      containerDefinitions:
        fn::readFile: task-definitions/service.json
      volumes:
        - name: service-storage
          fsxWindowsFileServerVolumeConfiguration:
            fileSystemId: ${aws_fsx_windows_file_system.test.id}
            rootDirectory: \data
            authorizationConfig:
              credentialsParameter: ${test.arn}
              domain: ${aws_directory_service_directory.test.name}
  test:
    type: aws:secretsmanager:SecretVersion
    properties:
      secretId: ${aws_secretsmanager_secret.test.id}
      secretString:
        fn::toJSON:
          username: admin
          password: ${aws_directory_service_directory.test.password}
```
{{% /example %}}
{{% example %}}
### Example Using `container_definitions` and `inference_accelerator`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ecs.TaskDefinition("test", {
    containerDefinitions: `[
  {
    "cpu": 10,
    "command": ["sleep", "10"],
    "entryPoint": ["/"],
    "environment": [
      {"name": "VARNAME", "value": "VARVAL"}
    ],
    "essential": true,
    "image": "jenkins",
    "memory": 128,
    "name": "jenkins",
    "portMappings": [
      {
        "containerPort": 80,
        "hostPort": 8080
      }
    ],
        "resourceRequirements":[
            {
                "type":"InferenceAccelerator",
                "value":"device_1"
            }
        ]
  }
]

`,
    family: "test",
    inferenceAccelerators: [{
        deviceName: "device_1",
        deviceType: "eia1.medium",
    }],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ecs.TaskDefinition("test",
    container_definitions="""[
  {
    "cpu": 10,
    "command": ["sleep", "10"],
    "entryPoint": ["/"],
    "environment": [
      {"name": "VARNAME", "value": "VARVAL"}
    ],
    "essential": true,
    "image": "jenkins",
    "memory": 128,
    "name": "jenkins",
    "portMappings": [
      {
        "containerPort": 80,
        "hostPort": 8080
      }
    ],
        "resourceRequirements":[
            {
                "type":"InferenceAccelerator",
                "value":"device_1"
            }
        ]
  }
]

""",
    family="test",
    inference_accelerators=[aws.ecs.TaskDefinitionInferenceAcceleratorArgs(
        device_name="device_1",
        device_type="eia1.medium",
    )])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ecs.TaskDefinition("test", new()
    {
        ContainerDefinitions = @"[
  {
    ""cpu"": 10,
    ""command"": [""sleep"", ""10""],
    ""entryPoint"": [""/""],
    ""environment"": [
      {""name"": ""VARNAME"", ""value"": ""VARVAL""}
    ],
    ""essential"": true,
    ""image"": ""jenkins"",
    ""memory"": 128,
    ""name"": ""jenkins"",
    ""portMappings"": [
      {
        ""containerPort"": 80,
        ""hostPort"": 8080
      }
    ],
        ""resourceRequirements"":[
            {
                ""type"":""InferenceAccelerator"",
                ""value"":""device_1""
            }
        ]
  }
]

",
        Family = "test",
        InferenceAccelerators = new[]
        {
            new Aws.Ecs.Inputs.TaskDefinitionInferenceAcceleratorArgs
            {
                DeviceName = "device_1",
                DeviceType = "eia1.medium",
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.NewTaskDefinition(ctx, "test", &ecs.TaskDefinitionArgs{
			ContainerDefinitions: pulumi.String(`[
  {
    "cpu": 10,
    "command": ["sleep", "10"],
    "entryPoint": ["/"],
    "environment": [
      {"name": "VARNAME", "value": "VARVAL"}
    ],
    "essential": true,
    "image": "jenkins",
    "memory": 128,
    "name": "jenkins",
    "portMappings": [
      {
        "containerPort": 80,
        "hostPort": 8080
      }
    ],
        "resourceRequirements":[
            {
                "type":"InferenceAccelerator",
                "value":"device_1"
            }
        ]
  }
]

`),
			Family: pulumi.String("test"),
			InferenceAccelerators: ecs.TaskDefinitionInferenceAcceleratorArray{
				&ecs.TaskDefinitionInferenceAcceleratorArgs{
					DeviceName: pulumi.String("device_1"),
					DeviceType: pulumi.String("eia1.medium"),
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
import com.pulumi.aws.ecs.TaskDefinition;
import com.pulumi.aws.ecs.TaskDefinitionArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionInferenceAcceleratorArgs;
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
        var test = new TaskDefinition("test", TaskDefinitionArgs.builder()        
            .containerDefinitions("""
[
  {
    "cpu": 10,
    "command": ["sleep", "10"],
    "entryPoint": ["/"],
    "environment": [
      {"name": "VARNAME", "value": "VARVAL"}
    ],
    "essential": true,
    "image": "jenkins",
    "memory": 128,
    "name": "jenkins",
    "portMappings": [
      {
        "containerPort": 80,
        "hostPort": 8080
      }
    ],
        "resourceRequirements":[
            {
                "type":"InferenceAccelerator",
                "value":"device_1"
            }
        ]
  }
]

            """)
            .family("test")
            .inferenceAccelerators(TaskDefinitionInferenceAcceleratorArgs.builder()
                .deviceName("device_1")
                .deviceType("eia1.medium")
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ecs:TaskDefinition
    properties:
      containerDefinitions: |+
        [
          {
            "cpu": 10,
            "command": ["sleep", "10"],
            "entryPoint": ["/"],
            "environment": [
              {"name": "VARNAME", "value": "VARVAL"}
            ],
            "essential": true,
            "image": "jenkins",
            "memory": 128,
            "name": "jenkins",
            "portMappings": [
              {
                "containerPort": 80,
                "hostPort": 8080
              }
            ],
                "resourceRequirements":[
                    {
                        "type":"InferenceAccelerator",
                        "value":"device_1"
                    }
                ]
          }
        ]

      family: test
      inferenceAccelerators:
        - deviceName: device_1
          deviceType: eia1.medium
```
{{% /example %}}
{{% example %}}
### Example Using `runtime_platform` and `fargate`

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.ecs.TaskDefinition("test", {
    containerDefinitions: `[
  {
    "name": "iis",
    "image": "mcr.microsoft.com/windows/servercore/iis",
    "cpu": 1024,
    "memory": 2048,
    "essential": true
  }
]

`,
    cpu: "1024",
    family: "test",
    memory: "2048",
    networkMode: "awsvpc",
    requiresCompatibilities: ["FARGATE"],
    runtimePlatform: {
        cpuArchitecture: "X86_64",
        operatingSystemFamily: "WINDOWS_SERVER_2019_CORE",
    },
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.ecs.TaskDefinition("test",
    container_definitions="""[
  {
    "name": "iis",
    "image": "mcr.microsoft.com/windows/servercore/iis",
    "cpu": 1024,
    "memory": 2048,
    "essential": true
  }
]

""",
    cpu="1024",
    family="test",
    memory="2048",
    network_mode="awsvpc",
    requires_compatibilities=["FARGATE"],
    runtime_platform=aws.ecs.TaskDefinitionRuntimePlatformArgs(
        cpu_architecture="X86_64",
        operating_system_family="WINDOWS_SERVER_2019_CORE",
    ))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.Ecs.TaskDefinition("test", new()
    {
        ContainerDefinitions = @"[
  {
    ""name"": ""iis"",
    ""image"": ""mcr.microsoft.com/windows/servercore/iis"",
    ""cpu"": 1024,
    ""memory"": 2048,
    ""essential"": true
  }
]

",
        Cpu = "1024",
        Family = "test",
        Memory = "2048",
        NetworkMode = "awsvpc",
        RequiresCompatibilities = new[]
        {
            "FARGATE",
        },
        RuntimePlatform = new Aws.Ecs.Inputs.TaskDefinitionRuntimePlatformArgs
        {
            CpuArchitecture = "X86_64",
            OperatingSystemFamily = "WINDOWS_SERVER_2019_CORE",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := ecs.NewTaskDefinition(ctx, "test", &ecs.TaskDefinitionArgs{
			ContainerDefinitions: pulumi.String(`[
  {
    "name": "iis",
    "image": "mcr.microsoft.com/windows/servercore/iis",
    "cpu": 1024,
    "memory": 2048,
    "essential": true
  }
]

`),
			Cpu:         pulumi.String("1024"),
			Family:      pulumi.String("test"),
			Memory:      pulumi.String("2048"),
			NetworkMode: pulumi.String("awsvpc"),
			RequiresCompatibilities: pulumi.StringArray{
				pulumi.String("FARGATE"),
			},
			RuntimePlatform: &ecs.TaskDefinitionRuntimePlatformArgs{
				CpuArchitecture:       pulumi.String("X86_64"),
				OperatingSystemFamily: pulumi.String("WINDOWS_SERVER_2019_CORE"),
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
import com.pulumi.aws.ecs.TaskDefinition;
import com.pulumi.aws.ecs.TaskDefinitionArgs;
import com.pulumi.aws.ecs.inputs.TaskDefinitionRuntimePlatformArgs;
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
        var test = new TaskDefinition("test", TaskDefinitionArgs.builder()        
            .containerDefinitions("""
[
  {
    "name": "iis",
    "image": "mcr.microsoft.com/windows/servercore/iis",
    "cpu": 1024,
    "memory": 2048,
    "essential": true
  }
]

            """)
            .cpu(1024)
            .family("test")
            .memory(2048)
            .networkMode("awsvpc")
            .requiresCompatibilities("FARGATE")
            .runtimePlatform(TaskDefinitionRuntimePlatformArgs.builder()
                .cpuArchitecture("X86_64")
                .operatingSystemFamily("WINDOWS_SERVER_2019_CORE")
                .build())
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:ecs:TaskDefinition
    properties:
      containerDefinitions: |+
        [
          {
            "name": "iis",
            "image": "mcr.microsoft.com/windows/servercore/iis",
            "cpu": 1024,
            "memory": 2048,
            "essential": true
          }
        ]

      cpu: 1024
      family: test
      memory: 2048
      networkMode: awsvpc
      requiresCompatibilities:
        - FARGATE
      runtimePlatform:
        cpuArchitecture: X86_64
        operatingSystemFamily: WINDOWS_SERVER_2019_CORE
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import ECS Task Definitions using their ARNs. For example:

```sh
 $ pulumi import aws:ecs/taskDefinition:TaskDefinition example arn:aws:ecs:us-east-1:012345678910:task-definition/mytaskfamily:123
```
 