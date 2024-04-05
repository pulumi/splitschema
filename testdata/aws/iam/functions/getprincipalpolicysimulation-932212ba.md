Runs a simulation of the IAM policies of a particular principal against a given hypothetical request.

You can use this data source in conjunction with
Preconditions and Postconditions so that your configuration can test either whether it should have sufficient access to do its own work, or whether policies your configuration declares itself are sufficient for their intended use elsewhere.

> **Note:** Correctly using this data source requires familiarity with various details of AWS Identity and Access Management, and how various AWS services integrate with it. For general information on the AWS IAM policy simulator, see [Testing IAM policies with the IAM policy simulator](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_testing-policies.html). This data source wraps the `iam:SimulatePrincipalPolicy` API action described on that page.

{{% examples %}}
## Example Usage
{{% /examples %}}