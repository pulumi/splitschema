Provides a Redshift Cluster Resource.


> **NOTE:** A Redshift cluster's default IAM role can be managed both by this resource's `default_iam_role_arn` argument and the `aws.redshift.ClusterIamRoles` resource's `default_iam_role_arn` argument. Do not configure different values for both arguments. Doing so will cause a conflict of default IAM roles.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.Cluster("example", {
    clusterIdentifier: "tf-redshift-cluster",
    clusterType: "single-node",
    databaseName: "mydb",
    masterPassword: "Mustbe8characters",
    masterUsername: "exampleuser",
    nodeType: "dc1.large",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.Cluster("example",
    cluster_identifier="tf-redshift-cluster",
    cluster_type="single-node",
    database_name="mydb",
    master_password="Mustbe8characters",
    master_username="exampleuser",
    node_type="dc1.large")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedShift.Cluster("example", new()
    {
        ClusterIdentifier = "tf-redshift-cluster",
        ClusterType = "single-node",
        DatabaseName = "mydb",
        MasterPassword = "Mustbe8characters",
        MasterUsername = "exampleuser",
        NodeType = "dc1.large",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewCluster(ctx, "example", &redshift.ClusterArgs{
			ClusterIdentifier: pulumi.String("tf-redshift-cluster"),
			ClusterType:       pulumi.String("single-node"),
			DatabaseName:      pulumi.String("mydb"),
			MasterPassword:    pulumi.String("Mustbe8characters"),
			MasterUsername:    pulumi.String("exampleuser"),
			NodeType:          pulumi.String("dc1.large"),
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
import com.pulumi.aws.redshift.Cluster;
import com.pulumi.aws.redshift.ClusterArgs;
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
        var example = new Cluster("example", ClusterArgs.builder()        
            .clusterIdentifier("tf-redshift-cluster")
            .clusterType("single-node")
            .databaseName("mydb")
            .masterPassword("Mustbe8characters")
            .masterUsername("exampleuser")
            .nodeType("dc1.large")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:Cluster
    properties:
      clusterIdentifier: tf-redshift-cluster
      clusterType: single-node
      databaseName: mydb
      masterPassword: Mustbe8characters
      masterUsername: exampleuser
      nodeType: dc1.large
```
{{% /example %}}
{{% example %}}
### With Managed Credentials

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.redshift.Cluster("example", {
    clusterIdentifier: "tf-redshift-cluster",
    clusterType: "single-node",
    databaseName: "mydb",
    manageMasterPassword: true,
    masterUsername: "exampleuser",
    nodeType: "dc1.large",
});
```
```python
import pulumi
import pulumi_aws as aws

example = aws.redshift.Cluster("example",
    cluster_identifier="tf-redshift-cluster",
    cluster_type="single-node",
    database_name="mydb",
    manage_master_password=True,
    master_username="exampleuser",
    node_type="dc1.large")
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var example = new Aws.RedShift.Cluster("example", new()
    {
        ClusterIdentifier = "tf-redshift-cluster",
        ClusterType = "single-node",
        DatabaseName = "mydb",
        ManageMasterPassword = true,
        MasterUsername = "exampleuser",
        NodeType = "dc1.large",
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/redshift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := redshift.NewCluster(ctx, "example", &redshift.ClusterArgs{
			ClusterIdentifier:    pulumi.String("tf-redshift-cluster"),
			ClusterType:          pulumi.String("single-node"),
			DatabaseName:         pulumi.String("mydb"),
			ManageMasterPassword: pulumi.Bool(true),
			MasterUsername:       pulumi.String("exampleuser"),
			NodeType:             pulumi.String("dc1.large"),
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
import com.pulumi.aws.redshift.Cluster;
import com.pulumi.aws.redshift.ClusterArgs;
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
        var example = new Cluster("example", ClusterArgs.builder()        
            .clusterIdentifier("tf-redshift-cluster")
            .clusterType("single-node")
            .databaseName("mydb")
            .manageMasterPassword(true)
            .masterUsername("exampleuser")
            .nodeType("dc1.large")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:redshift:Cluster
    properties:
      clusterIdentifier: tf-redshift-cluster
      clusterType: single-node
      databaseName: mydb
      manageMasterPassword: true
      masterUsername: exampleuser
      nodeType: dc1.large
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Redshift Clusters using the `cluster_identifier`. For example:

```sh
 $ pulumi import aws:redshift/cluster:Cluster myprodcluster tf-redshift-cluster-12345
```
 