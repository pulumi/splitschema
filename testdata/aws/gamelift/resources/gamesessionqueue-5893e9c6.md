Provides an GameLift Game Session Queue resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const test = new aws.gamelift.GameSessionQueue("test", {
    destinations: [
        aws_gamelift_fleet.us_west_2_fleet.arn,
        aws_gamelift_fleet.eu_central_1_fleet.arn,
    ],
    notificationTarget: aws_sns_topic.game_session_queue_notifications.arn,
    playerLatencyPolicies: [
        {
            maximumIndividualPlayerLatencyMilliseconds: 100,
            policyDurationSeconds: 5,
        },
        {
            maximumIndividualPlayerLatencyMilliseconds: 200,
        },
    ],
    timeoutInSeconds: 60,
});
```
```python
import pulumi
import pulumi_aws as aws

test = aws.gamelift.GameSessionQueue("test",
    destinations=[
        aws_gamelift_fleet["us_west_2_fleet"]["arn"],
        aws_gamelift_fleet["eu_central_1_fleet"]["arn"],
    ],
    notification_target=aws_sns_topic["game_session_queue_notifications"]["arn"],
    player_latency_policies=[
        aws.gamelift.GameSessionQueuePlayerLatencyPolicyArgs(
            maximum_individual_player_latency_milliseconds=100,
            policy_duration_seconds=5,
        ),
        aws.gamelift.GameSessionQueuePlayerLatencyPolicyArgs(
            maximum_individual_player_latency_milliseconds=200,
        ),
    ],
    timeout_in_seconds=60)
```
```csharp
using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var test = new Aws.GameLift.GameSessionQueue("test", new()
    {
        Destinations = new[]
        {
            aws_gamelift_fleet.Us_west_2_fleet.Arn,
            aws_gamelift_fleet.Eu_central_1_fleet.Arn,
        },
        NotificationTarget = aws_sns_topic.Game_session_queue_notifications.Arn,
        PlayerLatencyPolicies = new[]
        {
            new Aws.GameLift.Inputs.GameSessionQueuePlayerLatencyPolicyArgs
            {
                MaximumIndividualPlayerLatencyMilliseconds = 100,
                PolicyDurationSeconds = 5,
            },
            new Aws.GameLift.Inputs.GameSessionQueuePlayerLatencyPolicyArgs
            {
                MaximumIndividualPlayerLatencyMilliseconds = 200,
            },
        },
        TimeoutInSeconds = 60,
    });

});
```
```go
package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/gamelift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := gamelift.NewGameSessionQueue(ctx, "test", &gamelift.GameSessionQueueArgs{
			Destinations: pulumi.StringArray{
				aws_gamelift_fleet.Us_west_2_fleet.Arn,
				aws_gamelift_fleet.Eu_central_1_fleet.Arn,
			},
			NotificationTarget: pulumi.Any(aws_sns_topic.Game_session_queue_notifications.Arn),
			PlayerLatencyPolicies: gamelift.GameSessionQueuePlayerLatencyPolicyArray{
				&gamelift.GameSessionQueuePlayerLatencyPolicyArgs{
					MaximumIndividualPlayerLatencyMilliseconds: pulumi.Int(100),
					PolicyDurationSeconds:                      pulumi.Int(5),
				},
				&gamelift.GameSessionQueuePlayerLatencyPolicyArgs{
					MaximumIndividualPlayerLatencyMilliseconds: pulumi.Int(200),
				},
			},
			TimeoutInSeconds: pulumi.Int(60),
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
import com.pulumi.aws.gamelift.GameSessionQueue;
import com.pulumi.aws.gamelift.GameSessionQueueArgs;
import com.pulumi.aws.gamelift.inputs.GameSessionQueuePlayerLatencyPolicyArgs;
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
        var test = new GameSessionQueue("test", GameSessionQueueArgs.builder()        
            .destinations(            
                aws_gamelift_fleet.us_west_2_fleet().arn(),
                aws_gamelift_fleet.eu_central_1_fleet().arn())
            .notificationTarget(aws_sns_topic.game_session_queue_notifications().arn())
            .playerLatencyPolicies(            
                GameSessionQueuePlayerLatencyPolicyArgs.builder()
                    .maximumIndividualPlayerLatencyMilliseconds(100)
                    .policyDurationSeconds(5)
                    .build(),
                GameSessionQueuePlayerLatencyPolicyArgs.builder()
                    .maximumIndividualPlayerLatencyMilliseconds(200)
                    .build())
            .timeoutInSeconds(60)
            .build());

    }
}
```
```yaml
resources:
  test:
    type: aws:gamelift:GameSessionQueue
    properties:
      destinations:
        - ${aws_gamelift_fleet.us_west_2_fleet.arn}
        - ${aws_gamelift_fleet.eu_central_1_fleet.arn}
      notificationTarget: ${aws_sns_topic.game_session_queue_notifications.arn}
      playerLatencyPolicies:
        - maximumIndividualPlayerLatencyMilliseconds: 100
          policyDurationSeconds: 5
        - maximumIndividualPlayerLatencyMilliseconds: 200
      timeoutInSeconds: 60
```
{{% /example %}}
{{% /examples %}}

## Import

Using `pulumi import`, import GameLift Game Session Queues using their `name`. For example:

```sh
 $ pulumi import aws:gamelift/gameSessionQueue:GameSessionQueue example example
```
 