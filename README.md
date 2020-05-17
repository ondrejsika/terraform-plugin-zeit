> This repository is in **Work in Progress** state. If you need something, create an [issue](https://github.com/ondrejsika/terraform-provider-vercel/issues/new)

# terraform-provider-vercel

    2019 Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/terraform-provider-vercel

![Build](https://github.com/ondrejsika/terraform-provider-zeit/workflows/Build/badge.svg)

## My Related Projects

- [ondrejsika/zeit-go](https://github.com/ondrejsika/zeit-go) - Go client for Zeit API
- [ondrejsika/zeit-api-mock](https://github.com/ondrejsika/zeit-api-mock) - Zeit API Mock

## Buy Domain on Zeit.co using Terraform

![Buy Domain on Zeit.co using Terraform](buy-domain-on-zeit-using-terraform.png)

## Example usage

```terraform
provider "zeit" {
  token = "secret-token"
  // Optional
  // api_origin = "https://zeit-api-mock.sikademo.com"
}

resource "zeit_domain" "sikademozeit_com" {
  domain = "sikademozeit.com"
  expected_price = 12
}

resource "zeit_dns" "sikademozeit_com" {
  domain = zeit_domain.sikademozeit_com.domain
  name = ""
  value = "1.2.3.4"
  type = "A"
}

resource "zeit_dns" "www_sikademozeit_com" {
  domain = zeit_domain.sikademozeit_com.domain
  name = "www"
  value = "sikademozeit.com."
  type = "CNAME"
}

resource "zeit_dns" "mail_sikademozeit_com" {
  domain = zeit_domain.sikademozeit_com.domain
  name = "mail"
  value = "5.6.7.8"
  type = "A"
}

resource "zeit_dns" "mx_sikademozeit_com" {
  domain = zeit_domain.sikademozeit_com.domain
  name = ""
  value = "99 mail.sikademozeit.com."
  type = "MX"
}

resource "zeit_project" "demo" {
  name = "sika-demo-zeit"
}
```

## Change Log

### v1.3.2

- Fix error handing of errors from `ondrejsika/zeit-go` API client
- Handle buy of unavailable domains

### v1.3.1

- Update `ondrejsika/zeit-go` for `/v4/domain/buy` API

### v1.3.0

- Add parameter `remove_domain_on_destroy` with default `false` to `zeit_domain`. When you call `terraform destroy` domain will be kept on Zeit if you not set `remove_domain_on_destroy=true`
- Rewrite for [ondrejsika/zeit-go](https://github.com/ondrejsika/zeit-go)

### v1.2.0

- Add resource `zeit_domain` for buy domains on Zeit

### v1.1.0

- Add `api_origin` configuration for provider

### v1.0.0

- Create provider `zeit`
- Add resource `zeit_dns` with minimum configuration
