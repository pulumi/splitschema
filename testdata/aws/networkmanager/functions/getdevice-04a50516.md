Retrieve information about a device.

{{% examples %}}
## Example Usage
{{% example %}}

```java
package generated_program;

import com.pulumi.Context;
import com.pulumi.Pulumi;
import com.pulumi.core.Output;
import com.pulumi.aws.networkmanager.NetworkmanagerFunctions;
import com.pulumi.aws.networkmanager.inputs.GetDeviceArgs;
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
        final var example = NetworkmanagerFunctions.getDevice(GetDeviceArgs.builder()
            .globalNetworkIdId(var_.global_network_id())
            .deviceId(var_.device_id())
            .build());

    }
}
```
```yaml
variables:
  example:
    fn::invoke:
      Function: aws:networkmanager:getDevice
      Arguments:
        globalNetworkIdId: ${var.global_network_id}
        deviceId: ${var.device_id}
```
{{% /example %}}
{{% /examples %}}