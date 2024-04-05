Provides a GameLift Alias resource.

{{% examples %}}
## Example Usage
{{% example %}}

```typescript
import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const exampleGameSessionQueue = new aws.gamelift.GameSessionQueue("exampleGameSessionQueue", {
    destinations: [],
    playerLatencyPolicies: [
        {
            maximumIndividualPlayerLatencyMilliseconds: 3,
            policyDurationSeconds: 7,
        },
        {
            maximumIndividualPlayerLatencyMilliseconds: 10,
        },
    ],
    timeoutInSeconds: 25,
});
const exampleMatchmakingRuleSet = new aws.gamelift.MatchmakingRuleSet("exampleMatchmakingRuleSet", {ruleSetBody: JSON.stringify({
    name: "test",
    ruleLanguageVersion: "1.0",
    teams: [{
        name: "alpha",
        minPlayers: 1,
        maxPlayers: 5,
    }],
})});
const exampleMatchmakingConfiguration = new aws.gamelift.MatchmakingConfiguration("exampleMatchmakingConfiguration", {
    acceptanceRequired: false,
    customEventData: "pvp",
    gameSessionData: "game_session_data",
    backfillMode: "MANUAL",
    requestTimeoutSeconds: 30,
    ruleSetName: aws_gamelift_matchmaking_rule_set.test.name,
    gameSessionQueueArns: [aws_gamelift_game_session_queue.test.arn],
    tags: {
        key1: "value1",
    },
});
```
```python
import pulumi
import json
import pulumi_aws as aws

example_game_session_queue = aws.gamelift.GameSessionQueue("exampleGameSessionQueue",
    destinations=[],
    player_latency_policies=[
        aws.gamelift.GameSessionQueuePlayerLatencyPolicyArgs(
            maximum_individual_player_latency_milliseconds=3,
            policy_duration_seconds=7,
        ),
        aws.gamelift.GameSessionQueuePlayerLatencyPolicyArgs(
            maximum_individual_player_latency_milliseconds=10,
        ),
    ],
    timeout_in_seconds=25)
example_matchmaking_rule_set = aws.gamelift.MatchmakingRuleSet("exampleMatchmakingRuleSet", rule_set_body=json.dumps({
    "name": "test",
    "ruleLanguageVersion": "1.0",
    "teams": [{
        "name": "alpha",
        "minPlayers": 1,
        "maxPlayers": 5,
    }],
}))
example_matchmaking_configuration = aws.gamelift.MatchmakingConfiguration("exampleMatchmakingConfiguration",
    acceptance_required=False,
    custom_event_data="pvp",
    game_session_data="game_session_data",
    backfill_mode="MANUAL",
    request_timeout_seconds=30,
    rule_set_name=aws_gamelift_matchmaking_rule_set["test"]["name"],
    game_session_queue_arns=[aws_gamelift_game_session_queue["test"]["arn"]],
    tags={
        "key1": "value1",
    })
```
```csharp
using System.Collections.Generic;
using System.Linq;
using System.Text.Json;
using Pulumi;
using Aws = Pulumi.Aws;

return await Deployment.RunAsync(() => 
{
    var exampleGameSessionQueue = new Aws.GameLift.GameSessionQueue("exampleGameSessionQueue", new()
    {
        Destinations = new[] {},
        PlayerLatencyPolicies = new[]
        {
            new Aws.GameLift.Inputs.GameSessionQueuePlayerLatencyPolicyArgs
            {
                MaximumIndividualPlayerLatencyMilliseconds = 3,
                PolicyDurationSeconds = 7,
            },
            new Aws.GameLift.Inputs.GameSessionQueuePlayerLatencyPolicyArgs
            {
                MaximumIndividualPlayerLatencyMilliseconds = 10,
            },
        },
        TimeoutInSeconds = 25,
    });

    var exampleMatchmakingRuleSet = new Aws.GameLift.MatchmakingRuleSet("exampleMatchmakingRuleSet", new()
    {
        RuleSetBody = JsonSerializer.Serialize(new Dictionary<string, object?>
        {
            ["name"] = "test",
            ["ruleLanguageVersion"] = "1.0",
            ["teams"] = new[]
            {
                new Dictionary<string, object?>
                {
                    ["name"] = "alpha",
                    ["minPlayers"] = 1,
                    ["maxPlayers"] = 5,
                },
            },
        }),
    });

    var exampleMatchmakingConfiguration = new Aws.GameLift.MatchmakingConfiguration("exampleMatchmakingConfiguration", new()
    {
        AcceptanceRequired = false,
        CustomEventData = "pvp",
        GameSessionData = "game_session_data",
        BackfillMode = "MANUAL",
        RequestTimeoutSeconds = 30,
        RuleSetName = aws_gamelift_matchmaking_rule_set.Test.Name,
        GameSessionQueueArns = new[]
        {
            aws_gamelift_game_session_queue.Test.Arn,
        },
        Tags = 
        {
            { "key1", "value1" },
        },
    });

});
```
```go
package main

import (
	"encoding/json"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/gamelift"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := gamelift.NewGameSessionQueue(ctx, "exampleGameSessionQueue", &gamelift.GameSessionQueueArgs{
			Destinations: pulumi.StringArray{},
			PlayerLatencyPolicies: gamelift.GameSessionQueuePlayerLatencyPolicyArray{
				&gamelift.GameSessionQueuePlayerLatencyPolicyArgs{
					MaximumIndividualPlayerLatencyMilliseconds: pulumi.Int(3),
					PolicyDurationSeconds:                      pulumi.Int(7),
				},
				&gamelift.GameSessionQueuePlayerLatencyPolicyArgs{
					MaximumIndividualPlayerLatencyMilliseconds: pulumi.Int(10),
				},
			},
			TimeoutInSeconds: pulumi.Int(25),
		})
		if err != nil {
			return err
		}
		tmpJSON0, err := json.Marshal(map[string]interface{}{
			"name":                "test",
			"ruleLanguageVersion": "1.0",
			"teams": []map[string]interface{}{
				map[string]interface{}{
					"name":       "alpha",
					"minPlayers": 1,
					"maxPlayers": 5,
				},
			},
		})
		if err != nil {
			return err
		}
		json0 := string(tmpJSON0)
		_, err = gamelift.NewMatchmakingRuleSet(ctx, "exampleMatchmakingRuleSet", &gamelift.MatchmakingRuleSetArgs{
			RuleSetBody: pulumi.String(json0),
		})
		if err != nil {
			return err
		}
		_, err = gamelift.NewMatchmakingConfiguration(ctx, "exampleMatchmakingConfiguration", &gamelift.MatchmakingConfigurationArgs{
			AcceptanceRequired:    pulumi.Bool(false),
			CustomEventData:       pulumi.String("pvp"),
			GameSessionData:       pulumi.String("game_session_data"),
			BackfillMode:          pulumi.String("MANUAL"),
			RequestTimeoutSeconds: pulumi.Int(30),
			RuleSetName:           pulumi.Any(aws_gamelift_matchmaking_rule_set.Test.Name),
			GameSessionQueueArns: pulumi.StringArray{
				aws_gamelift_game_session_queue.Test.Arn,
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
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
import com.pulumi.aws.gamelift.MatchmakingRuleSet;
import com.pulumi.aws.gamelift.MatchmakingRuleSetArgs;
import com.pulumi.aws.gamelift.MatchmakingConfiguration;
import com.pulumi.aws.gamelift.MatchmakingConfigurationArgs;
import static com.pulumi.codegen.internal.Serialization.*;
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
        var exampleGameSessionQueue = new GameSessionQueue("exampleGameSessionQueue", GameSessionQueueArgs.builder()        
            .destinations()
            .playerLatencyPolicies(            
                GameSessionQueuePlayerLatencyPolicyArgs.builder()
                    .maximumIndividualPlayerLatencyMilliseconds(3)
                    .policyDurationSeconds(7)
                    .build(),
                GameSessionQueuePlayerLatencyPolicyArgs.builder()
                    .maximumIndividualPlayerLatencyMilliseconds(10)
                    .build())
            .timeoutInSeconds(25)
            .build());

        var exampleMatchmakingRuleSet = new MatchmakingRuleSet("exampleMatchmakingRuleSet", MatchmakingRuleSetArgs.builder()        
            .ruleSetBody(serializeJson(
                jsonObject(
                    jsonProperty("name", "test"),
                    jsonProperty("ruleLanguageVersion", "1.0"),
                    jsonProperty("teams", jsonArray(jsonObject(
                        jsonProperty("name", "alpha"),
                        jsonProperty("minPlayers", 1),
                        jsonProperty("maxPlayers", 5)
                    )))
                )))
            .build());

        var exampleMatchmakingConfiguration = new MatchmakingConfiguration("exampleMatchmakingConfiguration", MatchmakingConfigurationArgs.builder()        
            .acceptanceRequired(false)
            .customEventData("pvp")
            .gameSessionData("game_session_data")
            .backfillMode("MANUAL")
            .requestTimeoutSeconds(30)
            .ruleSetName(aws_gamelift_matchmaking_rule_set.test().name())
            .gameSessionQueueArns(aws_gamelift_game_session_queue.test().arn())
            .tags(Map.of("key1", "value1"))
            .build());

    }
}
```
```yaml
resources:
  exampleGameSessionQueue:
    type: aws:gamelift:GameSessionQueue
    properties:
      destinations: []
      playerLatencyPolicies:
        - maximumIndividualPlayerLatencyMilliseconds: 3
          policyDurationSeconds: 7
        - maximumIndividualPlayerLatencyMilliseconds: 10
      timeoutInSeconds: 25
  exampleMatchmakingRuleSet:
    type: aws:gamelift:MatchmakingRuleSet
    properties:
      ruleSetBody:
        fn::toJSON:
          name: test
          ruleLanguageVersion: '1.0'
          teams:
            - name: alpha
              minPlayers: 1
              maxPlayers: 5
  exampleMatchmakingConfiguration:
    type: aws:gamelift:MatchmakingConfiguration
    properties:
      acceptanceRequired: false
      customEventData: pvp
      gameSessionData: game_session_data
      backfillMode: MANUAL
      requestTimeoutSeconds: 30
      ruleSetName: ${aws_gamelift_matchmaking_rule_set.test.name}
      gameSessionQueueArns:
        - ${aws_gamelift_game_session_queue.test.arn}
      tags:
        key1: value1
```
{{% /example %}}
{{% /examples %}}

## Import

GameLift Matchmaking Configurations can be imported using the ID, e.g.,

```sh
 $ pulumi import aws:gamelift/matchmakingConfiguration:MatchmakingConfiguration example <matchmakingconfiguration-id>
```
 