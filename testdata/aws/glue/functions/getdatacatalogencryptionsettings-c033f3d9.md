This data source can be used to fetch information about AWS Glue Data Catalog Encryption Settings.

{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.glue.GlueFunctions;
import com.pulumi.aws.glue.inputs.GetDataCatalogEncryptionSettingsArgs;
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
        final var example = GlueFunctions.getDataCatalogEncryptionSettings(GetDataCatalogEncryptionSettingsArgs.builder()
            .id("123456789123")
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:glue:getDataCatalogEncryptionSettings
      Arguments:
        id: '123456789123'
```
{{% /example %}}
{{% /examples %}}