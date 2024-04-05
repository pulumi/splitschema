Provides a CloudTrail Event Data Store.

More information about event data stores can be found in the [Event Data Store User Guide](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/query-event-data-store.html).

> **Tip:** For an organization event data store you must create this resource in the management account.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

The most simple event data store configuration requires us to only set the `name` attribute. The event data store will automatically capture all management events. To capture management events from all the regions, `multi_region_enabled` must be `true`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.cloudtrail.EventDataStore("example", {});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.cloudtrail.EventDataStore("example")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.CloudTrail.EventDataStore("example");

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudtrail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cloudtrail.NewEventDataStore(ctx, "example", nil)
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
import com.pulumi.aws.cloudtrail.EventDataStore;
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
        var example = new EventDataStore("example");

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudtrail:EventDataStore
```
{{% /example %}}
### Data Event Logging

CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 bucket objects and Lambda function invocations. Additional information about data event configuration can be found in the following links:

- [CloudTrail API AdvancedFieldSelector documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedFieldSelector.html)
{{% example %}}
### Log all DynamoDB PutEvent actions for a specific DynamoDB table

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const table = aws.dynamodb.getTable({
    name: "not-important-dynamodb-table",
});
// ... other configuration ...
const example = new aws.cloudtrail.EventDataStore("example", {advancedEventSelectors: [{
    name: "Log all DynamoDB PutEvent actions for a specific DynamoDB table",
    fieldSelectors: [
        {
            field: "eventCategory",
            equals: ["Data"],
        },
        {
            field: "resources.type",
            equals: ["AWS::DynamoDB::Table"],
        },
        {
            field: "eventName",
            equals: ["PutItem"],
        },
        {
            field: "resources.ARN",
            equals: [table.then(table => table.arn)],
        },
    ],
}]});
```
```python
import pulumi
import pulumi_aws as aws

table = aws.dynamodb.get_table(name="not-important-dynamodb-table")
# ... other configuration ...
example = aws.cloudtrail.EventDataStore("example", advanced_event_selectors=[aws.cloudtrail.EventDataStoreAdvancedEventSelectorArgs(
    name="Log all DynamoDB PutEvent actions for a specific DynamoDB table",
    field_selectors=[
        aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
            field="eventCategory",
            equals=["Data"],
        ),
        aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
            field="resources.type",
            equals=["AWS::DynamoDB::Table"],
        ),
        aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
            field="eventName",
            equals=["PutItem"],
        ),
        aws.cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs(
            field="resources.ARN",
            equals=[table.arn],
        ),
    ],
)])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var table = Aws.DynamoDB.GetTable.Invoke(new()
    {
        Name = "not-important-dynamodb-table",
    });

    // ... other configuration ...
    var example = new Aws.CloudTrail.EventDataStore("example", new()
    {
        AdvancedEventSelectors = new[]
        {
            new Aws.CloudTrail.Inputs.EventDataStoreAdvancedEventSelectorArgs
            {
                Name = "Log all DynamoDB PutEvent actions for a specific DynamoDB table",
                FieldSelectors = new[]
                {
                    new Aws.CloudTrail.Inputs.EventDataStoreAdvancedEventSelectorFieldSelectorArgs
                    {
                        Field = "eventCategory",
                        Equals = new[]
                        {
                            "Data",
                        },
                    },
                    new Aws.CloudTrail.Inputs.EventDataStoreAdvancedEventSelectorFieldSelectorArgs
                    {
                        Field = "resources.type",
                        Equals = new[]
                        {
                            "AWS::DynamoDB::Table",
                        },
                    },
                    new Aws.CloudTrail.Inputs.EventDataStoreAdvancedEventSelectorFieldSelectorArgs
                    {
                        Field = "eventName",
                        Equals = new[]
                        {
                            "PutItem",
                        },
                    },
                    new Aws.CloudTrail.Inputs.EventDataStoreAdvancedEventSelectorFieldSelectorArgs
                    {
                        Field = "resources.ARN",
                        Equals = new[]
                        {
                            table.Apply(getTableResult => getTableResult.Arn),
                        },
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
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudtrail"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		table, err := dynamodb.LookupTable(ctx, &dynamodb.LookupTableArgs{
			Name: "not-important-dynamodb-table",
		}, nil)
		if err != nil {
			return err
		}
		_, err = cloudtrail.NewEventDataStore(ctx, "example", &cloudtrail.EventDataStoreArgs{
			AdvancedEventSelectors: cloudtrail.EventDataStoreAdvancedEventSelectorArray{
				&cloudtrail.EventDataStoreAdvancedEventSelectorArgs{
					Name: pulumi.String("Log all DynamoDB PutEvent actions for a specific DynamoDB table"),
					FieldSelectors: cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArray{
						&cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs{
							Field: pulumi.String("eventCategory"),
							Equals: pulumi.StringArray{
								pulumi.String("Data"),
							},
						},
						&cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs{
							Field: pulumi.String("resources.type"),
							Equals: pulumi.StringArray{
								pulumi.String("AWS::DynamoDB::Table"),
							},
						},
						&cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs{
							Field: pulumi.String("eventName"),
							Equals: pulumi.StringArray{
								pulumi.String("PutItem"),
							},
						},
						&cloudtrail.EventDataStoreAdvancedEventSelectorFieldSelectorArgs{
							Field: pulumi.String("resources.ARN"),
							Equals: pulumi.StringArray{
								*pulumi.String(table.Arn),
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
import com.pulumi.aws.dynamodb.DynamodbFunctions;
import com.pulumi.aws.dynamodb.inputs.GetTableArgs;
import com.pulumi.aws.cloudtrail.EventDataStore;
import com.pulumi.aws.cloudtrail.EventDataStoreArgs;
import com.pulumi.aws.cloudtrail.inputs.EventDataStoreAdvancedEventSelectorArgs;
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
        final var table = DynamodbFunctions.getTable(GetTableArgs.builder()
            .name("not-important-dynamodb-table")
            .build());

        var example = new EventDataStore("example", EventDataStoreArgs.builder()        
            .advancedEventSelectors(EventDataStoreAdvancedEventSelectorArgs.builder()
                .name("Log all DynamoDB PutEvent actions for a specific DynamoDB table")
                .fieldSelectors(                
                    EventDataStoreAdvancedEventSelectorFieldSelectorArgs.builder()
                        .field("eventCategory")
                        .equals("Data")
                        .build(),
                    EventDataStoreAdvancedEventSelectorFieldSelectorArgs.builder()
                        .field("resources.type")
                        .equals("AWS::DynamoDB::Table")
                        .build(),
                    EventDataStoreAdvancedEventSelectorFieldSelectorArgs.builder()
                        .field("eventName")
                        .equals("PutItem")
                        .build(),
                    EventDataStoreAdvancedEventSelectorFieldSelectorArgs.builder()
                        .field("resources.ARN")
                        .equals(table.applyValue(getTableResult -> getTableResult.arn()))
                        .build())
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:cloudtrail:EventDataStore
    properties:
      advancedEventSelectors:
        - name: Log all DynamoDB PutEvent actions for a specific DynamoDB table
          fieldSelectors:
            - field: eventCategory
              equals:
                - Data
            - field: resources.type
              equals:
                - AWS::DynamoDB::Table
            - field: eventName
              equals:
                - PutItem
            - field: resources.ARN
              equals:
                - ${table.arn}
variables:
  table:
    fn::invoke:
      Function: aws:dynamodb:getTable
      Arguments:
        name: not-important-dynamodb-table
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import event data stores using their `arn`. For example:

```sh
 $ pulumi import aws:cloudtrail/eventDataStore:EventDataStore example arn:aws:cloudtrail:us-east-1:123456789123:eventdatastore/22333815-4414-412c-b155-dd254033gfhf
```
 