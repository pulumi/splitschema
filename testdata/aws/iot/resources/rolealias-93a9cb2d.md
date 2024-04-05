Provides an IoT role alias.

{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.iam.IamFunctions;
import com.pulumi.aws.iam.inputs.GetPolicyDocumentArgs;
import com.pulumi.aws.iam.Role;
import com.pulumi.aws.iam.RoleArgs;
import com.pulumi.aws.iot.RoleAlias;
import com.pulumi.aws.iot.RoleAliasArgs;
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
        final var assumeRole = IamFunctions.getPolicyDocument(GetPolicyDocumentArgs.builder()
            .effect("Allow")
            .principals(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
            .actions("sts:AssumeRole")
            .build());

        var role = new Role("role", RoleArgs.builder()        
            .assumeRolePolicy(assumeRole.applyValue(getPolicyDocumentResult -> getPolicyDocumentResult.json()))
            .build());

        var alias = new RoleAlias("alias", RoleAliasArgs.builder()        
            .alias("Thermostat-dynamodb-access-role-alias")
            .roleArn(role.arn())
            .build());

    }
}
```
```yaml
resources:
  role:
    type: aws:iam:Role
    properties:
      assumeRolePolicy: ${assumeRole.json}
  alias:
    type: aws:iot:RoleAlias
    properties:
      alias: Thermostat-dynamodb-access-role-alias
      roleArn: ${role.arn}
variables:
  assumeRole:
    fn::invoke:
      Function: aws:iam:getPolicyDocument
      Arguments:
        effect: Allow
        principals:
          - type: Service
            identifiers:
              - credentials.iot.amazonaws.com
        actions:
          - sts:AssumeRole
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import IOT Role Alias using the alias. For example:

```sh
 $ pulumi import aws:iot/roleAlias:RoleAlias example myalias
```
 