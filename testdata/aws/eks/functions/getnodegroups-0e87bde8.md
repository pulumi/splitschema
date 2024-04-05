Retrieve the EKS Node Groups associated with a named EKS cluster. This will allow you to pass a list of Node Group names to other resources.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleNodeGroups = aws.eks.getNodeGroups({
    clusterName: "example",
});
const exampleNodeGroup = exampleNodeGroups.then(exampleNodeGroups => .map(([, ]) => (aws.eks.getNodeGroup({
    clusterName: "example",
    nodeGroupName: __value,
}))));
```
```python
import pulumi
import pulumi_aws as aws

example_node_groups = aws.eks.get_node_groups(cluster_name="example")
example_node_group = [aws.eks.get_node_group(cluster_name="example",
    node_group_name=__value) for __key, __value in example_node_groups.names]
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleNodeGroups = Aws.Eks.GetNodeGroups.Invoke(new()
    {
        ClusterName = "example",
    });

    var exampleNodeGroup = .Select(__value => 
    {
        return Aws.Eks.GetNodeGroup.Invoke(new()
        {
            ClusterName = "example",
            NodeGroupName = __value,
        });
    }).ToList();

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleNodeGroups, err := eks.GetNodeGroups(ctx, &eks.GetNodeGroupsArgs{
			ClusterName: "example",
		}, nil)
		if err != nil {
			return err
		}
		_ := "TODO: For expression"
		return nil
	})
}
```
{{% /example %}}
{{% /examples %}}