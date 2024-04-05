Manages an AWS Auto Scaling scaling plan.
More information can be found in the [AWS Auto Scaling User Guide](https://docs.aws.amazon.com/autoscaling/plans/userguide/what-is-aws-auto-scaling.html).

> **NOTE:** The AWS Auto Scaling service uses an AWS IAM service-linked role to manage predictive scaling of Amazon EC2 Auto Scaling groups. The service attempts to automatically create this role the first time a scaling plan with predictive scaling enabled is created.
An `aws.iam.ServiceLinkedRole` resource can be used to manually manage this role.
See the [AWS documentation](https://docs.aws.amazon.com/autoscaling/plans/userguide/aws-auto-scaling-service-linked-roles.html#create-service-linked-role-manual) for more details.

{{% examples %}}
## Example Usage
{{% /examples %}}

## Import

Using `pulumi import`, import Auto Scaling scaling plans using the `name`. For example:

```sh
 $ pulumi import aws:autoscalingplans/scalingPlan:ScalingPlan example MyScale1
```
 