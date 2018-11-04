package test

import (
	"testing"
	"time"
        "fmt"

	"github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/gruntwork-io/terratest/modules/packer"
)

func TestWebServer(t *testing.T) {
  packerOptions := &packer.Options {
    // The path to where the Packer template is located
    Template: "../web-server/web-server.json",
  }
  // Build the AMI
  amiId := packer.BuildAmi(t, packerOptions)
  terraformOptions := &terraform.Options {
    // The path to where your Terraform code is located
    TerraformDir: "../web-server",
    // Variables to pass to our Terraform code using -var options
    Vars: map[string]interface{} {
      "ami_id": amiId,
    },
  }

  fmt.Println(*terraformOptions)
  // At the end of the test, run `terraform destroy`
  defer terraform.Destroy(t, terraformOptions)
  // Run `terraform init` and `terraform apply`
  terraform.InitAndApply(t, terraformOptions)
  // Run `terraform output` to get the value of an output variable
  url := terraform.Output(t, terraformOptions, "url")
  // Verify that we get back a 200 OK with the expected text. It
  // takes ~1 min for the Instance to boot, so retry a few times.
  status := 200
  text := "Hello, World"
  retries := 15
  sleep := 5 * time.Second
  http_helper.HttpGetWithRetry(t, url, status, text, retries, sleep)
}
