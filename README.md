# Terraform Provider for AWS SSO Admin Role

**Archived:** The official AWS provider now includes [`data.aws_iam_roles`](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/data-sources/iam_roles#roles-provisioned-by-aws-sso) which can be used to find one (or many) SSO roles.

<hr>

This provider exposes a [data source](https://github.com/hashicorp/terraform-provider-aws/pull/18048) that to be merged into the upstream [terraform-provider-aws](https://github.com/hashicorp/terraform-provider-aws) project. The data source fetches the IAM role created in the target AWS account by the AWS SSO instance for the AWS organization.

## Configuring the provider

This project uses the same provider setup from [terraform-provider-aws](https://registry.terraform.io/providers/hashicorp/aws/latest/docs)

## Example usage

```hcl
data "aws_ssoadmin_instances" "sso" {}

resource "aws_ssoadmin_permission_set" "readonly" {
  instance_arn = tolist(data.aws_ssoadmin_instances.sso.arns)[0]
  name         = "ReadOnly"
}

data "awssso_role" "readonly" {
  permission_set_name = aws_ssoadmin_permission_set.readonly.name
}

output "role" {
  description = "IAM role ARN for the role created by the AWS SSO instance for the AWS organization."
  value       = data.awssso_role.sso_readonly.arn
}
```
