Provides a AWS Transfer AS2 Profile resource.

{{% examples %}}
## Example Usage
{{% example %}}
### Basic

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.transfer.Profile;
import com.pulumi.aws.transfer.ProfileArgs;
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
        var example = new Profile("example", ProfileArgs.builder()        
            .as2Id("example")
            .certificateIds(aws_transfer_certificate.example().certificate_id())
            .usage("LOCAL")
            .build());

    }
}
```
```yaml
resources:
  example:
    type: aws:transfer:Profile
    properties:
      as2Id: example
      certificateIds:
        - ${aws_transfer_certificate.example.certificate_id}
      usage: LOCAL
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import Transfer AS2 Profile using the `profile_id`. For example:

```sh
 $ pulumi import aws:transfer/profile:Profile example p-4221a88afd5f4362a
```
 