Resource for managing an AWS OpenSearch Serverless Security Config.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic Usage

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.opensearch.ServerlessSecurityConfig;
import com.pulumi.aws.opensearch.ServerlessSecurityConfigArgs;
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
        var example = new ServerlessSecurityConfig("example", ServerlessSecurityConfigArgs.builder()        
            .type("saml")
            .samlOptions(ServerlessSecurityConfigSamlOptionsArgs.builder()
                .metadata(Files.readString(Paths.get(String.format("%s/idp-metadata.xml", path.module()))))
                .build())
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:opensearch:ServerlessSecurityConfig
    properties:
      type: saml
      samlOptions:
        - metadata:
            fn::readFile: ${path.module}/idp-metadata.xml
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import OpenSearchServerless Access Policy using the `name` argument prefixed with the string `saml/account_id/`. For example:

```sh
 $ pulumi import aws:opensearch/serverlessSecurityConfig:ServerlessSecurityConfig example saml/123456789012/example
```
 