Information about RDS orderable DB instances and valid parameter combinations.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.rds.getOrderableDbInstance({
    engine: "mysql",
    engineVersion: "5.7.22",
    licenseModel: "general-public-license",
    preferredInstanceClasses: [
        "db.r6.xlarge",
        "db.m4.large",
        "db.t3.small",
    ],
    storageType: "standard",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.rds.get_orderable_db_instance(engine="mysql",
    engine_version="5.7.22",
    license_model="general-public-license",
    preferred_instance_classes=[
        "db.r6.xlarge",
        "db.m4.large",
        "db.t3.small",
    ],
    storage_type="standard")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Rds.GetOrderableDbInstance.Invoke(new()
    {
        Engine = "mysql",
        EngineVersion = "5.7.22",
        LicenseModel = "general-public-license",
        PreferredInstanceClasses = new[]
        {
            "db.r6.xlarge",
            "db.m4.large",
            "db.t3.small",
        },
        StorageType = "standard",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.GetOrderableDbInstance(ctx, &rds.GetOrderableDbInstanceArgs{
			Engine:        "mysql",
			EngineVersion: pulumi.StringRef("5.7.22"),
			LicenseModel:  pulumi.StringRef("general-public-license"),
			PreferredInstanceClasses: []string{
				"db.r6.xlarge",
				"db.m4.large",
				"db.t3.small",
			},
			StorageType: pulumi.StringRef("standard"),
		}, nil)
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
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetOrderableDbInstanceArgs;
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
        final var test = RdsFunctions.getOrderableDbInstance(GetOrderableDbInstanceArgs.builder()
            .engine("mysql")
            .engineVersion("5.7.22")
            .licenseModel("general-public-license")
            .preferredInstanceClasses(            
                "db.r6.xlarge",
                "db.m4.large",
                "db.t3.small")
            .storageType("standard")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:rds:getOrderableDbInstance
      Arguments:
        engine: mysql
        engineVersion: 5.7.22
        licenseModel: general-public-license
        preferredInstanceClasses:
          - db.r6.xlarge
          - db.m4.large
          - db.t3.small
        storageType: standard
```

Valid parameter combinations can also be found with `preferred_engine_versions` and/or `preferred_instance_classes`.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = aws.rds.getOrderableDbInstance({
    engine: "mysql",
    licenseModel: "general-public-license",
    preferredEngineVersions: [
        "5.6.35",
        "5.6.41",
        "5.6.44",
    ],
    preferredInstanceClasses: [
        "db.t2.small",
        "db.t3.medium",
        "db.t3.large",
    ],
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.rds.get_orderable_db_instance(engine="mysql",
    license_model="general-public-license",
    preferred_engine_versions=[
        "5.6.35",
        "5.6.41",
        "5.6.44",
    ],
    preferred_instance_classes=[
        "db.t2.small",
        "db.t3.medium",
        "db.t3.large",
    ])
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = Aws.Rds.GetOrderableDbInstance.Invoke(new()
    {
        Engine = "mysql",
        LicenseModel = "general-public-license",
        PreferredEngineVersions = new[]
        {
            "5.6.35",
            "5.6.41",
            "5.6.44",
        },
        PreferredInstanceClasses = new[]
        {
            "db.t2.small",
            "db.t3.medium",
            "db.t3.large",
        },
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := rds.GetOrderableDbInstance(ctx, &rds.GetOrderableDbInstanceArgs{
			Engine:       "mysql",
			LicenseModel: pulumi.StringRef("general-public-license"),
			PreferredEngineVersions: []string{
				"5.6.35",
				"5.6.41",
				"5.6.44",
			},
			PreferredInstanceClasses: []string{
				"db.t2.small",
				"db.t3.medium",
				"db.t3.large",
			},
		}, nil)
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
import com.pulumi.aws.rds.RdsFunctions;
import com.pulumi.aws.rds.inputs.GetOrderableDbInstanceArgs;
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
        final var test = RdsFunctions.getOrderableDbInstance(GetOrderableDbInstanceArgs.builder()
            .engine("mysql")
            .licenseModel("general-public-license")
            .preferredEngineVersions(            
                "5.6.35",
                "5.6.41",
                "5.6.44")
            .preferredInstanceClasses(            
                "db.t2.small",
                "db.t3.medium",
                "db.t3.large")
            .build());

    }
}
```
```yaml
variables:
  test:
    fn::invoke:
      Function: aws:rds:getOrderableDbInstance
      Arguments:
        engine: mysql
        licenseModel: general-public-license
        preferredEngineVersions:
          - 5.6.35
          - 5.6.41
          - 5.6.44
        preferredInstanceClasses:
          - db.t2.small
          - db.t3.medium
          - db.t3.large
```
{{% /example %}}
{{% /examples %}}