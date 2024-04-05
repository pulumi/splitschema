Provides a Security Hub custom insight resource. See the [Managing custom insights section](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-custom-insights.html) of the AWS User Guide for more information.

{{% examples %}}
## Example Usage
{{% example %}}
### Filter by AWS account ID

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleInsight = new aws.securityhub.Insight("exampleInsight", {
    filters: {
        awsAccountIds: [
            {
                comparison: "EQUALS",
                value: "1234567890",
            },
            {
                comparison: "EQUALS",
                value: "09876543210",
            },
        ],
    },
    groupByAttribute: "AwsAccountId",
}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_insight = aws.securityhub.Insight("exampleInsight",
    filters=aws.securityhub.InsightFiltersArgs(
        aws_account_ids=[
            aws.securityhub.InsightFiltersAwsAccountIdArgs(
                comparison="EQUALS",
                value="1234567890",
            ),
            aws.securityhub.InsightFiltersAwsAccountIdArgs(
                comparison="EQUALS",
                value="09876543210",
            ),
        ],
    ),
    group_by_attribute="AwsAccountId",
    opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleInsight = new Aws.SecurityHub.Insight("exampleInsight", new()
    {
        Filters = new Aws.SecurityHub.Inputs.InsightFiltersArgs
        {
            AwsAccountIds = new[]
            {
                new Aws.SecurityHub.Inputs.InsightFiltersAwsAccountIdArgs
                {
                    Comparison = "EQUALS",
                    Value = "1234567890",
                },
                new Aws.SecurityHub.Inputs.InsightFiltersAwsAccountIdArgs
                {
                    Comparison = "EQUALS",
                    Value = "09876543210",
                },
            },
        },
        GroupByAttribute = "AwsAccountId",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewInsight(ctx, "exampleInsight", &securityhub.InsightArgs{
			Filters: &securityhub.InsightFiltersArgs{
				AwsAccountIds: securityhub.InsightFiltersAwsAccountIdArray{
					&securityhub.InsightFiltersAwsAccountIdArgs{
						Comparison: pulumi.String("EQUALS"),
						Value:      pulumi.String("1234567890"),
					},
					&securityhub.InsightFiltersAwsAccountIdArgs{
						Comparison: pulumi.String("EQUALS"),
						Value:      pulumi.String("09876543210"),
					},
				},
			},
			GroupByAttribute: pulumi.String("AwsAccountId"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
		}))
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.Insight;
import com.pulumi.aws.securityhub.InsightArgs;
import com.pulumi.aws.securityhub.inputs.InsightFiltersArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleInsight = new Insight("exampleInsight", InsightArgs.builder()        
            .filters(InsightFiltersArgs.builder()
                .awsAccountIds(                
                    InsightFiltersAwsAccountIdArgs.builder()
                        .comparison("EQUALS")
                        .value("1234567890")
                        .build(),
                    InsightFiltersAwsAccountIdArgs.builder()
                        .comparison("EQUALS")
                        .value("09876543210")
                        .build())
                .build())
            .groupByAttribute("AwsAccountId")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:securityhub:Account
  exampleInsight:
    type: aws:securityhub:Insight
    properties:
      filters:
        awsAccountIds:
          - comparison: EQUALS
            value: '1234567890'
          - comparison: EQUALS
            value: '09876543210'
      groupByAttribute: AwsAccountId
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% example %}}
### Filter by date range

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleInsight = new aws.securityhub.Insight("exampleInsight", {
    filters: {
        createdAts: [{
            dateRange: {
                unit: "DAYS",
                value: 5,
            },
        }],
    },
    groupByAttribute: "CreatedAt",
}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_insight = aws.securityhub.Insight("exampleInsight",
    filters=aws.securityhub.InsightFiltersArgs(
        created_ats=[aws.securityhub.InsightFiltersCreatedAtArgs(
            date_range=aws.securityhub.InsightFiltersCreatedAtDateRangeArgs(
                unit="DAYS",
                value=5,
            ),
        )],
    ),
    group_by_attribute="CreatedAt",
    opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleInsight = new Aws.SecurityHub.Insight("exampleInsight", new()
    {
        Filters = new Aws.SecurityHub.Inputs.InsightFiltersArgs
        {
            CreatedAts = new[]
            {
                new Aws.SecurityHub.Inputs.InsightFiltersCreatedAtArgs
                {
                    DateRange = new Aws.SecurityHub.Inputs.InsightFiltersCreatedAtDateRangeArgs
                    {
                        Unit = "DAYS",
                        Value = 5,
                    },
                },
            },
        },
        GroupByAttribute = "CreatedAt",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewInsight(ctx, "exampleInsight", &securityhub.InsightArgs{
			Filters: &securityhub.InsightFiltersArgs{
				CreatedAts: securityhub.InsightFiltersCreatedAtArray{
					&securityhub.InsightFiltersCreatedAtArgs{
						DateRange: &securityhub.InsightFiltersCreatedAtDateRangeArgs{
							Unit:  pulumi.String("DAYS"),
							Value: pulumi.Int(5),
						},
					},
				},
			},
			GroupByAttribute: pulumi.String("CreatedAt"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
		}))
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.Insight;
import com.pulumi.aws.securityhub.InsightArgs;
import com.pulumi.aws.securityhub.inputs.InsightFiltersArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleInsight = new Insight("exampleInsight", InsightArgs.builder()        
            .filters(InsightFiltersArgs.builder()
                .createdAts(InsightFiltersCreatedAtArgs.builder()
                    .dateRange(InsightFiltersCreatedAtDateRangeArgs.builder()
                        .unit("DAYS")
                        .value(5)
                        .build())
                    .build())
                .build())
            .groupByAttribute("CreatedAt")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:securityhub:Account
  exampleInsight:
    type: aws:securityhub:Insight
    properties:
      filters:
        createdAts:
          - dateRange:
              unit: DAYS
              value: 5
      groupByAttribute: CreatedAt
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% example %}}
### Filter by destination IPv4 address

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleInsight = new aws.securityhub.Insight("exampleInsight", {
    filters: {
        networkDestinationIpv4s: [{
            cidr: "10.0.0.0/16",
        }],
    },
    groupByAttribute: "NetworkDestinationIpV4",
}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_insight = aws.securityhub.Insight("exampleInsight",
    filters=aws.securityhub.InsightFiltersArgs(
        network_destination_ipv4s=[aws.securityhub.InsightFiltersNetworkDestinationIpv4Args(
            cidr="10.0.0.0/16",
        )],
    ),
    group_by_attribute="NetworkDestinationIpV4",
    opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleInsight = new Aws.SecurityHub.Insight("exampleInsight", new()
    {
        Filters = new Aws.SecurityHub.Inputs.InsightFiltersArgs
        {
            NetworkDestinationIpv4s = new[]
            {
                new Aws.SecurityHub.Inputs.InsightFiltersNetworkDestinationIpv4Args
                {
                    Cidr = "10.0.0.0/16",
                },
            },
        },
        GroupByAttribute = "NetworkDestinationIpV4",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewInsight(ctx, "exampleInsight", &securityhub.InsightArgs{
			Filters: &securityhub.InsightFiltersArgs{
				NetworkDestinationIpv4s: securityhub.InsightFiltersNetworkDestinationIpv4Array{
					&securityhub.InsightFiltersNetworkDestinationIpv4Args{
						Cidr: pulumi.String("10.0.0.0/16"),
					},
				},
			},
			GroupByAttribute: pulumi.String("NetworkDestinationIpV4"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
		}))
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.Insight;
import com.pulumi.aws.securityhub.InsightArgs;
import com.pulumi.aws.securityhub.inputs.InsightFiltersArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleInsight = new Insight("exampleInsight", InsightArgs.builder()        
            .filters(InsightFiltersArgs.builder()
                .networkDestinationIpv4s(InsightFiltersNetworkDestinationIpv4Args.builder()
                    .cidr("10.0.0.0/16")
                    .build())
                .build())
            .groupByAttribute("NetworkDestinationIpV4")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:securityhub:Account
  exampleInsight:
    type: aws:securityhub:Insight
    properties:
      filters:
        networkDestinationIpv4s:
          - cidr: 10.0.0.0/16
      groupByAttribute: NetworkDestinationIpV4
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% example %}}
### Filter by finding's confidence

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleInsight = new aws.securityhub.Insight("exampleInsight", {
    filters: {
        confidences: [{
            gte: "80",
        }],
    },
    groupByAttribute: "Confidence",
}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_insight = aws.securityhub.Insight("exampleInsight",
    filters=aws.securityhub.InsightFiltersArgs(
        confidences=[aws.securityhub.InsightFiltersConfidenceArgs(
            gte="80",
        )],
    ),
    group_by_attribute="Confidence",
    opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleInsight = new Aws.SecurityHub.Insight("exampleInsight", new()
    {
        Filters = new Aws.SecurityHub.Inputs.InsightFiltersArgs
        {
            Confidences = new[]
            {
                new Aws.SecurityHub.Inputs.InsightFiltersConfidenceArgs
                {
                    Gte = "80",
                },
            },
        },
        GroupByAttribute = "Confidence",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewInsight(ctx, "exampleInsight", &securityhub.InsightArgs{
			Filters: &securityhub.InsightFiltersArgs{
				Confidences: securityhub.InsightFiltersConfidenceArray{
					&securityhub.InsightFiltersConfidenceArgs{
						Gte: pulumi.String("80"),
					},
				},
			},
			GroupByAttribute: pulumi.String("Confidence"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
		}))
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.Insight;
import com.pulumi.aws.securityhub.InsightArgs;
import com.pulumi.aws.securityhub.inputs.InsightFiltersArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleInsight = new Insight("exampleInsight", InsightArgs.builder()        
            .filters(InsightFiltersArgs.builder()
                .confidences(InsightFiltersConfidenceArgs.builder()
                    .gte("80")
                    .build())
                .build())
            .groupByAttribute("Confidence")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:securityhub:Account
  exampleInsight:
    type: aws:securityhub:Insight
    properties:
      filters:
        confidences:
          - gte: '80'
      groupByAttribute: Confidence
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% example %}}
### Filter by resource tags

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleAccount = new aws.securityhub.Account("exampleAccount", {});
const exampleInsight = new aws.securityhub.Insight("exampleInsight", {
    filters: {
        resourceTags: [{
            comparison: "EQUALS",
            key: "Environment",
            value: "Production",
        }],
    },
    groupByAttribute: "ResourceTags",
}, {
    dependsOn: [exampleAccount],
});
```
```python
import pulumi
import pulumi_aws as aws

example_account = aws.securityhub.Account("exampleAccount")
example_insight = aws.securityhub.Insight("exampleInsight",
    filters=aws.securityhub.InsightFiltersArgs(
        resource_tags=[aws.securityhub.InsightFiltersResourceTagArgs(
            comparison="EQUALS",
            key="Environment",
            value="Production",
        )],
    ),
    group_by_attribute="ResourceTags",
    opts=pulumi.ResourceOptions(depends_on=[example_account]))
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleAccount = new Aws.SecurityHub.Account("exampleAccount");

    var exampleInsight = new Aws.SecurityHub.Insight("exampleInsight", new()
    {
        Filters = new Aws.SecurityHub.Inputs.InsightFiltersArgs
        {
            ResourceTags = new[]
            {
                new Aws.SecurityHub.Inputs.InsightFiltersResourceTagArgs
                {
                    Comparison = "EQUALS",
                    Key = "Environment",
                    Value = "Production",
                },
            },
        },
        GroupByAttribute = "ResourceTags",
    }, new CustomResourceOptions
    {
        DependsOn = new[]
        {
            exampleAccount,
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/securityhub"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAccount, err := securityhub.NewAccount(ctx, "exampleAccount", nil)
		if err != nil {
			return err
		}
		_, err = securityhub.NewInsight(ctx, "exampleInsight", &securityhub.InsightArgs{
			Filters: &securityhub.InsightFiltersArgs{
				ResourceTags: securityhub.InsightFiltersResourceTagArray{
					&securityhub.InsightFiltersResourceTagArgs{
						Comparison: pulumi.String("EQUALS"),
						Key:        pulumi.String("Environment"),
						Value:      pulumi.String("Production"),
					},
				},
			},
			GroupByAttribute: pulumi.String("ResourceTags"),
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAccount,
		}))
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
import com.pulumi.aws.securityhub.Account;
import com.pulumi.aws.securityhub.Insight;
import com.pulumi.aws.securityhub.InsightArgs;
import com.pulumi.aws.securityhub.inputs.InsightFiltersArgs;
import com.pulumi.resources.CustomResourceOptions;
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
        var exampleAccount = new Account("exampleAccount");

        var exampleInsight = new Insight("exampleInsight", InsightArgs.builder()        
            .filters(InsightFiltersArgs.builder()
                .resourceTags(InsightFiltersResourceTagArgs.builder()
                    .comparison("EQUALS")
                    .key("Environment")
                    .value("Production")
                    .build())
                .build())
            .groupByAttribute("ResourceTags")
            .build(), CustomResourceOptions.builder()
                .dependsOn(exampleAccount)
                .build());

    }
}
```
```yaml
resources:
  exampleAccount:
    type: aws:securityhub:Account
  exampleInsight:
    type: aws:securityhub:Insight
    properties:
      filters:
        resourceTags:
          - comparison: EQUALS
            key: Environment
            value: Production
      groupByAttribute: ResourceTags
    options:
      dependson:
        - ${exampleAccount}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Security Hub insights using the ARN. For example:

```sh
 $ pulumi import aws:securityhub/insight:Insight example arn:aws:securityhub:us-west-2:1234567890:insight/1234567890/custom/91299ed7-abd0-4e44-a858-d0b15e37141a
```
 