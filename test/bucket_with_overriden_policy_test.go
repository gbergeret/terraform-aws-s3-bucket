package test

import (
   "testing"
   "github.com/gruntwork-io/terratest/modules/terraform"
)

func TestBucketWithOverridenPolicy(t *testing.T) {
   t.Parallel()

    // Construct the terraform options with default retryable errors to handle the most common
    // retryable errors in terraform testing.
    terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
        // The path to where our Terraform code is located
        TerraformDir: "../examples/bucket_with_overriden_policy",
    })

    // At the end of the test, run `terraform destroy` to clean up any resources that were created.
    defer terraform.Destroy(t, terraformOptions)

    // Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
    terraform.InitAndApply(t, terraformOptions)
}
