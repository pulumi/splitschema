Returns a unique endpoint specific to the AWS account making the call.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as kubernetes from "@pulumi/kubernetes";

const example = aws.iot.getEndpoint({});
const agent = new kubernetes.index.Kubernetes_pod("agent", {
    metadata: [{
        name: "my-device",
    }],
    spec: [{
        container: [{
            image: "gcr.io/my-project/image-name",
            name: "image-name",
            env: [{
                name: "IOT_ENDPOINT",
                value: example.endpointAddress,
            }],
        }],
    }],
});
```
```python
import pulumi
import pulumi_aws as aws
import pulumi_kubernetes as kubernetes

example = aws.iot.get_endpoint()
agent = kubernetes.index.Kubernetes_pod("agent",
    metadata=[{
        name: my-device,
    }],
    spec=[{
        container: [{
            image: gcr.io/my-project/image-name,
            name: image-name,
            env: [{
                name: IOT_ENDPOINT,
                value: example.endpoint_address,
            }],
        }],
    }])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;
using Kubernetes = Pulumi.Kubernetes;

return await Deployment.RunAsync(() => 
{
    var example = Aws.Iot.GetEndpoint.Invoke();

    var agent = new Kubernetes.Index.Kubernetes_pod("agent", new()
    {
        Metadata = new[]
        {
            
            {
                { "name", "my-device" },
            },
        },
        Spec = new[]
        {
            
            {
                { "container", new[]
                {
                    
                    {
                        { "image", "gcr.io/my-project/image-name" },
                        { "name", "image-name" },
                        { "env", new[]
                        {
                            
                            {
                                { "name", "IOT_ENDPOINT" },
                                { "value", example.Apply(getEndpointResult => getEndpointResult.EndpointAddress) },
                            },
                        } },
                    },
                } },
            },
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iot"
	"github.com/pulumi/pulumi-kubernetes/sdk/v1/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		example, err := iot.GetEndpoint(ctx, nil, nil)
		if err != nil {
			return err
		}
		_, err = index.NewKubernetes_pod(ctx, "agent", &index.Kubernetes_podArgs{
			Metadata: []map[string]interface{}{
				map[string]interface{}{
					"name": "my-device",
				},
			},
			Spec: []map[string]interface{}{
				map[string]interface{}{
					"container": []map[string]interface{}{
						map[string]interface{}{
							"image": "gcr.io/my-project/image-name",
							"name":  "image-name",
							"env": []map[string]interface{}{
								map[string]interface{}{
									"name":  "IOT_ENDPOINT",
									"value": example.EndpointAddress,
								},
							},
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
import com.pulumi.aws.iot.IotFunctions;
import com.pulumi.aws.iot.inputs.GetEndpointArgs;
import com.pulumi.kubernetes.kubernetes_pod;
import com.pulumi.kubernetes.Kubernetes_podArgs;
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
        final var example = IotFunctions.getEndpoint();

        var agent = new Kubernetes_pod("agent", Kubernetes_podArgs.builder()        
            .metadata(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .spec(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .build());

    }
}
```
```yaml
resources:
  agent:
    type: kubernetes:kubernetes_pod
    properties:
      metadata:
        - name: my-device
      spec:
        - container:
            - image: gcr.io/my-project/image-name
              name: image-name
              env:
                - name: IOT_ENDPOINT
                  value: ${example.endpointAddress}
variables:
  example:
    fn::invoke:
      Function: aws:iot:getEndpoint
      Arguments: {}
```
{{% /example %}}
{{% /examples %}}