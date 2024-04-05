Provides a Lightsail Database. Amazon Lightsail is a service to provide easy virtual private servers
with custom software already setup. See [What is Amazon Lightsail?](https://lightsail.aws.amazon.com/ls/docs/getting-started/article/what-is-amazon-lightsail)
for more information.

> **Note:** Lightsail is currently only supported in a limited number of AWS Regions, please see ["Regions and Availability Zones"](https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/) for more details

{{% examples %}}
## Example Usage
{{% example %}}
### Basic mysql blueprint

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lightsail.Database("test", {
    availabilityZone: "us-east-1a",
    blueprintId: "mysql_8_0",
    bundleId: "micro_1_0",
    masterDatabaseName: "testdatabasename",
    masterPassword: "testdatabasepassword",
    masterUsername: "test",
    relationalDatabaseName: "test",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lightsail.Database("test",
    availability_zone="us-east-1a",
    blueprint_id="mysql_8_0",
    bundle_id="micro_1_0",
    master_database_name="testdatabasename",
    master_password="testdatabasepassword",
    master_username="test",
    relational_database_name="test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LightSail.Database("test", new()
    {
        AvailabilityZone = "us-east-1a",
        BlueprintId = "mysql_8_0",
        BundleId = "micro_1_0",
        MasterDatabaseName = "testdatabasename",
        MasterPassword = "testdatabasepassword",
        MasterUsername = "test",
        RelationalDatabaseName = "test",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewDatabase(ctx, "test", &lightsail.DatabaseArgs{
			AvailabilityZone:       pulumi.String("us-east-1a"),
			BlueprintId:            pulumi.String("mysql_8_0"),
			BundleId:               pulumi.String("micro_1_0"),
			MasterDatabaseName:     pulumi.String("testdatabasename"),
			MasterPassword:         pulumi.String("testdatabasepassword"),
			MasterUsername:         pulumi.String("test"),
			RelationalDatabaseName: pulumi.String("test"),
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
import com.pulumi.aws.lightsail.Database;
import com.pulumi.aws.lightsail.DatabaseArgs;
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
        var test = new Database("test", DatabaseArgs.builder()        
            .availabilityZone("us-east-1a")
            .blueprintId("mysql_8_0")
            .bundleId("micro_1_0")
            .masterDatabaseName("testdatabasename")
            .masterPassword("testdatabasepassword")
            .masterUsername("test")
            .relationalDatabaseName("test")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lightsail:Database
    properties:
      availabilityZone: us-east-1a
      blueprintId: mysql_8_0
      bundleId: micro_1_0
      masterDatabaseName: testdatabasename
      masterPassword: testdatabasepassword
      masterUsername: test
      relationalDatabaseName: test
```
{{% /example %}}
{{% example %}}
### Basic postrgres blueprint

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lightsail.Database("test", {
    availabilityZone: "us-east-1a",
    blueprintId: "postgres_12",
    bundleId: "micro_1_0",
    masterDatabaseName: "testdatabasename",
    masterPassword: "testdatabasepassword",
    masterUsername: "test",
    relationalDatabaseName: "test",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lightsail.Database("test",
    availability_zone="us-east-1a",
    blueprint_id="postgres_12",
    bundle_id="micro_1_0",
    master_database_name="testdatabasename",
    master_password="testdatabasepassword",
    master_username="test",
    relational_database_name="test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LightSail.Database("test", new()
    {
        AvailabilityZone = "us-east-1a",
        BlueprintId = "postgres_12",
        BundleId = "micro_1_0",
        MasterDatabaseName = "testdatabasename",
        MasterPassword = "testdatabasepassword",
        MasterUsername = "test",
        RelationalDatabaseName = "test",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewDatabase(ctx, "test", &lightsail.DatabaseArgs{
			AvailabilityZone:       pulumi.String("us-east-1a"),
			BlueprintId:            pulumi.String("postgres_12"),
			BundleId:               pulumi.String("micro_1_0"),
			MasterDatabaseName:     pulumi.String("testdatabasename"),
			MasterPassword:         pulumi.String("testdatabasepassword"),
			MasterUsername:         pulumi.String("test"),
			RelationalDatabaseName: pulumi.String("test"),
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
import com.pulumi.aws.lightsail.Database;
import com.pulumi.aws.lightsail.DatabaseArgs;
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
        var test = new Database("test", DatabaseArgs.builder()        
            .availabilityZone("us-east-1a")
            .blueprintId("postgres_12")
            .bundleId("micro_1_0")
            .masterDatabaseName("testdatabasename")
            .masterPassword("testdatabasepassword")
            .masterUsername("test")
            .relationalDatabaseName("test")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lightsail:Database
    properties:
      availabilityZone: us-east-1a
      blueprintId: postgres_12
      bundleId: micro_1_0
      masterDatabaseName: testdatabasename
      masterPassword: testdatabasepassword
      masterUsername: test
      relationalDatabaseName: test
```
{{% /example %}}
{{% example %}}
### Custom backup and maintenance windows

Below is an example that sets a custom backup and maintenance window. Times are specified in UTC. This example will allow daily backups to take place between 16:00 and 16:30 each day. This example also requires any maintiance tasks (anything that would cause an outage, including changing some attributes) to take place on Tuesdays between 17:00 and 17:30. An action taken against this database that would cause an outage will wait until this time window to make the requested changes.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lightsail.Database("test", {
    availabilityZone: "us-east-1a",
    blueprintId: "postgres_12",
    bundleId: "micro_1_0",
    masterDatabaseName: "testdatabasename",
    masterPassword: "testdatabasepassword",
    masterUsername: "test",
    preferredBackupWindow: "16:00-16:30",
    preferredMaintenanceWindow: "Tue:17:00-Tue:17:30",
    relationalDatabaseName: "test",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lightsail.Database("test",
    availability_zone="us-east-1a",
    blueprint_id="postgres_12",
    bundle_id="micro_1_0",
    master_database_name="testdatabasename",
    master_password="testdatabasepassword",
    master_username="test",
    preferred_backup_window="16:00-16:30",
    preferred_maintenance_window="Tue:17:00-Tue:17:30",
    relational_database_name="test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LightSail.Database("test", new()
    {
        AvailabilityZone = "us-east-1a",
        BlueprintId = "postgres_12",
        BundleId = "micro_1_0",
        MasterDatabaseName = "testdatabasename",
        MasterPassword = "testdatabasepassword",
        MasterUsername = "test",
        PreferredBackupWindow = "16:00-16:30",
        PreferredMaintenanceWindow = "Tue:17:00-Tue:17:30",
        RelationalDatabaseName = "test",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewDatabase(ctx, "test", &lightsail.DatabaseArgs{
			AvailabilityZone:           pulumi.String("us-east-1a"),
			BlueprintId:                pulumi.String("postgres_12"),
			BundleId:                   pulumi.String("micro_1_0"),
			MasterDatabaseName:         pulumi.String("testdatabasename"),
			MasterPassword:             pulumi.String("testdatabasepassword"),
			MasterUsername:             pulumi.String("test"),
			PreferredBackupWindow:      pulumi.String("16:00-16:30"),
			PreferredMaintenanceWindow: pulumi.String("Tue:17:00-Tue:17:30"),
			RelationalDatabaseName:     pulumi.String("test"),
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
import com.pulumi.aws.lightsail.Database;
import com.pulumi.aws.lightsail.DatabaseArgs;
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
        var test = new Database("test", DatabaseArgs.builder()        
            .availabilityZone("us-east-1a")
            .blueprintId("postgres_12")
            .bundleId("micro_1_0")
            .masterDatabaseName("testdatabasename")
            .masterPassword("testdatabasepassword")
            .masterUsername("test")
            .preferredBackupWindow("16:00-16:30")
            .preferredMaintenanceWindow("Tue:17:00-Tue:17:30")
            .relationalDatabaseName("test")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lightsail:Database
    properties:
      availabilityZone: us-east-1a
      blueprintId: postgres_12
      bundleId: micro_1_0
      masterDatabaseName: testdatabasename
      masterPassword: testdatabasepassword
      masterUsername: test
      preferredBackupWindow: 16:00-16:30
      preferredMaintenanceWindow: Tue:17:00-Tue:17:30
      relationalDatabaseName: test
```
{{% /example %}}
{{% example %}}
### Final Snapshots

To enable creating a final snapshot of your database on deletion, use the `final_snapshot_name` argument to provide a name to be used for the snapshot.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lightsail.Database("test", {
    availabilityZone: "us-east-1a",
    blueprintId: "postgres_12",
    bundleId: "micro_1_0",
    finalSnapshotName: "MyFinalSnapshot",
    masterDatabaseName: "testdatabasename",
    masterPassword: "testdatabasepassword",
    masterUsername: "test",
    preferredBackupWindow: "16:00-16:30",
    preferredMaintenanceWindow: "Tue:17:00-Tue:17:30",
    relationalDatabaseName: "test",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lightsail.Database("test",
    availability_zone="us-east-1a",
    blueprint_id="postgres_12",
    bundle_id="micro_1_0",
    final_snapshot_name="MyFinalSnapshot",
    master_database_name="testdatabasename",
    master_password="testdatabasepassword",
    master_username="test",
    preferred_backup_window="16:00-16:30",
    preferred_maintenance_window="Tue:17:00-Tue:17:30",
    relational_database_name="test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LightSail.Database("test", new()
    {
        AvailabilityZone = "us-east-1a",
        BlueprintId = "postgres_12",
        BundleId = "micro_1_0",
        FinalSnapshotName = "MyFinalSnapshot",
        MasterDatabaseName = "testdatabasename",
        MasterPassword = "testdatabasepassword",
        MasterUsername = "test",
        PreferredBackupWindow = "16:00-16:30",
        PreferredMaintenanceWindow = "Tue:17:00-Tue:17:30",
        RelationalDatabaseName = "test",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewDatabase(ctx, "test", &lightsail.DatabaseArgs{
			AvailabilityZone:           pulumi.String("us-east-1a"),
			BlueprintId:                pulumi.String("postgres_12"),
			BundleId:                   pulumi.String("micro_1_0"),
			FinalSnapshotName:          pulumi.String("MyFinalSnapshot"),
			MasterDatabaseName:         pulumi.String("testdatabasename"),
			MasterPassword:             pulumi.String("testdatabasepassword"),
			MasterUsername:             pulumi.String("test"),
			PreferredBackupWindow:      pulumi.String("16:00-16:30"),
			PreferredMaintenanceWindow: pulumi.String("Tue:17:00-Tue:17:30"),
			RelationalDatabaseName:     pulumi.String("test"),
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
import com.pulumi.aws.lightsail.Database;
import com.pulumi.aws.lightsail.DatabaseArgs;
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
        var test = new Database("test", DatabaseArgs.builder()        
            .availabilityZone("us-east-1a")
            .blueprintId("postgres_12")
            .bundleId("micro_1_0")
            .finalSnapshotName("MyFinalSnapshot")
            .masterDatabaseName("testdatabasename")
            .masterPassword("testdatabasepassword")
            .masterUsername("test")
            .preferredBackupWindow("16:00-16:30")
            .preferredMaintenanceWindow("Tue:17:00-Tue:17:30")
            .relationalDatabaseName("test")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lightsail:Database
    properties:
      availabilityZone: us-east-1a
      blueprintId: postgres_12
      bundleId: micro_1_0
      finalSnapshotName: MyFinalSnapshot
      masterDatabaseName: testdatabasename
      masterPassword: testdatabasepassword
      masterUsername: test
      preferredBackupWindow: 16:00-16:30
      preferredMaintenanceWindow: Tue:17:00-Tue:17:30
      relationalDatabaseName: test
```
{{% /example %}}
{{% example %}}
### Apply Immediately

To enable applying changes immediately instead of waiting for a maintiance window, use the `apply_immediately` argument.

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.lightsail.Database("test", {
    applyImmediately: true,
    availabilityZone: "us-east-1a",
    blueprintId: "postgres_12",
    bundleId: "micro_1_0",
    masterDatabaseName: "testdatabasename",
    masterPassword: "testdatabasepassword",
    masterUsername: "test",
    relationalDatabaseName: "test",
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.lightsail.Database("test",
    apply_immediately=True,
    availability_zone="us-east-1a",
    blueprint_id="postgres_12",
    bundle_id="micro_1_0",
    master_database_name="testdatabasename",
    master_password="testdatabasepassword",
    master_username="test",
    relational_database_name="test")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.LightSail.Database("test", new()
    {
        ApplyImmediately = true,
        AvailabilityZone = "us-east-1a",
        BlueprintId = "postgres_12",
        BundleId = "micro_1_0",
        MasterDatabaseName = "testdatabasename",
        MasterPassword = "testdatabasepassword",
        MasterUsername = "test",
        RelationalDatabaseName = "test",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lightsail"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := lightsail.NewDatabase(ctx, "test", &lightsail.DatabaseArgs{
			ApplyImmediately:       pulumi.Bool(true),
			AvailabilityZone:       pulumi.String("us-east-1a"),
			BlueprintId:            pulumi.String("postgres_12"),
			BundleId:               pulumi.String("micro_1_0"),
			MasterDatabaseName:     pulumi.String("testdatabasename"),
			MasterPassword:         pulumi.String("testdatabasepassword"),
			MasterUsername:         pulumi.String("test"),
			RelationalDatabaseName: pulumi.String("test"),
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
import com.pulumi.aws.lightsail.Database;
import com.pulumi.aws.lightsail.DatabaseArgs;
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
        var test = new Database("test", DatabaseArgs.builder()        
            .applyImmediately(true)
            .availabilityZone("us-east-1a")
            .blueprintId("postgres_12")
            .bundleId("micro_1_0")
            .masterDatabaseName("testdatabasename")
            .masterPassword("testdatabasepassword")
            .masterUsername("test")
            .relationalDatabaseName("test")
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:lightsail:Database
    properties:
      applyImmediately: true
      availabilityZone: us-east-1a
      blueprintId: postgres_12
      bundleId: micro_1_0
      masterDatabaseName: testdatabasename
      masterPassword: testdatabasepassword
      masterUsername: test
      relationalDatabaseName: test
```
{{% /example %}}
{{% /examples %}}
## Blueprint Ids

A list of all available Lightsail Blueprints for Relational Databases the [aws lightsail get-relational-database-blueprints](https://docs.aws.amazon.com/cli/latest/reference/lightsail/get-relational-database-blueprints.html) aws cli command.

### Examples

- `mysql_8_0`
- `postgres_12`

### Prefix

A Blueprint ID starts with a prefix of the engine type.

### Suffix

A Blueprint ID has a sufix of the engine version.

## Bundles

A list of all available Lightsail Bundles for Relational Databases the [aws lightsail get-relational-database-bundles](https://docs.aws.amazon.com/cli/latest/reference/lightsail/get-relational-database-bundles.html) aws cli command.

### Examples

- `small_1_0`
- `small_ha_1_0`
- `large_1_0`
- `large_ha_1_0`

### Prefix

A Bundle ID starts with one of the below size prefixes:

- `micro_`
- `small_`
- `medium_`
- `large_`

### Infixes (Optional for HA Database)

A Bundle Id can have the following infix added in order to use the HA option of the selected bundle.

- `ha_`

### Suffix

A Bundle ID ends with one of the following suffix: `1_0`


## Import

Using `pulumi import`, import Lightsail Databases using their name. For example:

```sh
 $ pulumi import aws:lightsail/database:Database foo 'bar'
```
 