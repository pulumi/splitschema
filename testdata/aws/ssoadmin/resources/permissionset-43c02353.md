Provides a Single Sign-On (SSO) Permission Set resource

> **NOTE:** Updating this resource will automatically [Provision the Permission Set](https://docs.aws.amazon.com/singlesignon/latest/APIReference/API_ProvisionPermissionSet.html) to apply the corresponding updates to all assigned accounts.


## Import

Using `pulumi import`, import SSO Permission Sets using the `arn` and `instance_arn` separated by a comma (`,`). For example:

```sh
 $ pulumi import aws:ssoadmin/permissionSet:PermissionSet example arn:aws:sso:::permissionSet/ssoins-2938j0x8920sbj72/ps-80383020jr9302rk,arn:aws:sso:::instance/ssoins-2938j0x8920sbj72
```
 