Manages an individual DynamoDB resource tag. This resource should only be used in cases where DynamoDB resources are created outside the provider (e.g., Table replicas in other regions).

> **NOTE:** This tagging resource should not be combined with the resource for managing the parent resource. For example, using `aws.dynamodb.Table` and `aws.dynamodb.Tag` to manage tags of the same DynamoDB Table in the same region will cause a perpetual difference where the `aws_dynamodb_cluster` resource will try to remove the tag being added by the `aws.dynamodb.Tag` resource.

> **NOTE:** This tagging resource does not use the provider `ignore_tags` configuration.


## Import

Using `pulumi import`, import `aws_dynamodb_tag` using the DynamoDB resource identifier and key, separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:dynamodb/tag:Tag example arn:aws:dynamodb:us-east-1:123456789012:table/example,Name
```
 