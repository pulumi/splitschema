Provides an API Gateway VPC Link.

> **Note:** Amazon API Gateway Version 1 VPC Links enable private integrations that connect REST APIs to private resources in a VPC.
To enable private integration for HTTP APIs, use the Amazon API Gateway Version 2 VPC Link resource.

{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.lb.LoadBalancer;
import com.pulumi.aws.lb.LoadBalancerArgs;
import com.pulumi.aws.lb.inputs.LoadBalancerSubnetMappingArgs;
import com.pulumi.aws.apigateway.VpcLink;
import com.pulumi.aws.apigateway.VpcLinkArgs;
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
        var exampleLoadBalancer = new LoadBalancer("exampleLoadBalancer", LoadBalancerArgs.builder()        
            .internal(true)
            .loadBalancerType("network")
            .subnetMappings(LoadBalancerSubnetMappingArgs.builder()
                .subnetId("12345")
                .build())
            .build());

        var exampleVpcLink = new VpcLink("exampleVpcLink", VpcLinkArgs.builder()        
            .description("example description")
            .targetArn(exampleLoadBalancer.arn())
            .build());

    }
}
```
```yaml
resources:
  exampleLoadBalancer:
    type: aws:lb:LoadBalancer
    properties:
      internal: true
      loadBalancerType: network
      subnetMappings:
        - subnetId: '12345'
  exampleVpcLink:
    type: aws:apigateway:VpcLink
    properties:
      description: example description
      targetArn:
        - ${exampleLoadBalancer.arn}
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import API Gateway VPC Link using the `id`. For example:

```sh
 $ pulumi import aws:apigateway/vpcLink:VpcLink example 12345abcde
```
 