Provides a CE Anomaly Monitor.

{{% examples %}}
## Example Usage

There are two main types of a Cost Anomaly Monitor: `DIMENSIONAL` and `CUSTOM`.
{{% example %}}
### Dimensional Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const serviceMonitor = new aws.costexplorer.AnomalyMonitor("serviceMonitor", {
    monitorDimension: "SERVICE",
    monitorType: "DIMENSIONAL",
});
```
```python
import pulumi
import pulumi_aws as aws

service_monitor = aws.costexplorer.AnomalyMonitor("serviceMonitor",
    monitor_dimension="SERVICE",
    monitor_type="DIMENSIONAL")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var serviceMonitor = new Aws.CostExplorer.AnomalyMonitor("serviceMonitor", new()
    {
        MonitorDimension = "SERVICE",
        MonitorType = "DIMENSIONAL",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/costexplorer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costexplorer.NewAnomalyMonitor(ctx, "serviceMonitor", &costexplorer.AnomalyMonitorArgs{
			MonitorDimension: pulumi.String("SERVICE"),
			MonitorType:      pulumi.String("DIMENSIONAL"),
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
import com.pulumi.aws.costexplorer.AnomalyMonitor;
import com.pulumi.aws.costexplorer.AnomalyMonitorArgs;
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
        var serviceMonitor = new AnomalyMonitor("serviceMonitor", AnomalyMonitorArgs.builder()        
            .monitorDimension("SERVICE")
            .monitorType("DIMENSIONAL")
            .build());

    }
}
```
```yaml
resources:
  serviceMonitor:
    type: aws:costexplorer:AnomalyMonitor
    properties:
      monitorDimension: SERVICE
      monitorType: DIMENSIONAL
```
{{% /example %}}
{{% example %}}
### Custom Example

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.costexplorer.AnomalyMonitor("test", {
    monitorType: "CUSTOM",
    monitorSpecification: JSON.stringify({
        And: undefined,
        CostCategories: undefined,
        Dimensions: undefined,
        Not: undefined,
        Or: undefined,
        Tags: {
            Key: "CostCenter",
            MatchOptions: undefined,
            Values: ["10000"],
        },
    }),
});
```
```python
import pulumi
import json
import pulumi_aws as aws

test = aws.costexplorer.AnomalyMonitor("test",
    monitor_type="CUSTOM",
    monitor_specification=json.dumps({
        "And": None,
        "CostCategories": None,
        "Dimensions": None,
        "Not": None,
        "Or": None,
        "Tags": {
            "Key": "CostCenter",
            "MatchOptions": None,
            "Values": ["10000"],
        },
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
    var test = new Aws.CostExplorer.AnomalyMonitor("test", new()
    {
        MonitorType = "CUSTOM",
        MonitorSpecification = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["And"] = null,
            ["CostCategories"] = null,
            ["Dimensions"] = null,
            ["Not"] = null,
            ["Or"] = null,
            ["Tags"] = new Dictionary<string, object?>
            {
                ["Key"] = "CostCenter",
                ["MatchOptions"] = null,
                ["Values"] = new[]
                {
                    "10000",
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

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/costexplorer"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"And":            nil,
			"CostCategories": nil,
			"Dimensions":     nil,
			"Not":            nil,
			"Or":             nil,
			"Tags": map[string]interface{}{
				"Key":          "CostCenter",
				"MatchOptions": nil,
				"Values": []string{
					"10000",
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = costexplorer.NewAnomalyMonitor(ctx, "test", &costexplorer.AnomalyMonitorArgs{
			MonitorType:          pulumi.String("CUSTOM"),
			MonitorSpecification: pulumi.String(json0),
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
import com.pulumi.aws.costexplorer.AnomalyMonitor;
import com.pulumi.aws.costexplorer.AnomalyMonitorArgs;
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
        var test = new AnomalyMonitor("test", AnomalyMonitorArgs.builder()        
            .monitorType("CUSTOM")
            .monitorSpecification(serializeJson(
                jsonObject(
                    jsonProperty("And", null),
                    jsonProperty("CostCategories", null),
                    jsonProperty("Dimensions", null),
                    jsonProperty("Not", null),
                    jsonProperty("Or", null),
                    jsonProperty("Tags", jsonObject(
                        jsonProperty("Key", "CostCenter"),
                        jsonProperty("MatchOptions", null),
                        jsonProperty("Values", jsonArray("10000"))
                    ))
                )))
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:costexplorer:AnomalyMonitor
    properties:
      monitorType: CUSTOM
      monitorSpecification:
        fn::toJSON:
          And: null
          CostCategories: null
          Dimensions: null
          Not: null
          Or: null
          Tags:
            Key: CostCenter
            MatchOptions: null
            Values:
              - '10000'
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import `aws_ce_anomaly_monitor` using the `id`. For example:

```sh
 $ pulumi import aws:costexplorer/anomalyMonitor:AnomalyMonitor example costAnomalyMonitorARN
```
 