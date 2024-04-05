Manages an individual Autoscaling Group (ASG) tag. This resource should only be used in cases where ASGs are created outside the provider (e.g., ASGs implicitly created by EKS Node Groups).

> **NOTE:** This tagging resource should not be combined with the resource for managing the parent resource. For example, using `aws.autoscaling.Group` and `aws.autoscaling.Tag` to manage tags of the same ASG will cause a perpetual difference where the `aws.autoscaling.Group` resource will try to remove the tag being added by the `aws.autoscaling.Tag` resource.

> **NOTE:** This tagging resource does not use the provider `ignore_tags` configuration.


## Import

Using `pulumi import`, import `aws_autoscaling_group_tag` using the ASG name and key, separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:autoscaling/tag:Tag example asg-example,k8s.io/cluster-autoscaler/node-template/label/eks.amazonaws.com/capacityType
```
 