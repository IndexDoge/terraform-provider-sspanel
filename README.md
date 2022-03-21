<a href="https://terraform.io">
    <img src="https://raw.githubusercontent.com/hashicorp/terraform-website/master/public/img/logo-text.svg" alt="Terraform logo" title="Terraform" align="right" height="50" />
</a>


# SSPanel Terraform Provider

## Using the provider


Full, comprehensive documentation is available on the [Terraform Registry](https://registry.terraform.io/providers/IndexDoge/sspanel/latest/docs).

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

## Feedback

If you would like to provide feedback (not a bug or feature request) on the Terraform provider, you're welcome to via issues.

