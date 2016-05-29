# sss



## Description

Take a screenshot and upload to AWS S3.

## Usage

You have to do some settings.

```~/.config/sss/config.toml:toml
aws_access_key_id = ''
aws_secret_access_id = ''
aws_s3_region = ''
aws_s3_bucket_name = ''
aws_s3_path = ''
```

Save to `~/.config/sss/config.toml`


|Option|Description|
|---|---|
|-help|Show help|
|-baseurl|Base URL (ex. ss.example.com)|
|-browser|if set this option, open screenshot url in browser|
|-clipboard|if set this option, copy screenshot url to clipboard|
|-config|config file (default ~/.config/sss/config.toml)|
|-version|show version|


## Install

To install, use `go get`:

```bash
$ go get -d github.com/upamune/sss
```

## Contribution

1. Fork ([https://github.com/upamune/sss/fork](https://github.com/upamune/sss/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[upamune](https://github.com/upamune)
