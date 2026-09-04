# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Added

- Dependabot configuration for Go module and GitHub Actions updates ([3ec2843])
- zizmor linting of GitHub Actions workflows in CI ([3ec2843])
- Makefile, golangci-lint config, changelog script, AGENTS.md ([140bcf4])

### Changed

- Replaced the internal CI configuration with a GitHub Actions workflow ([3ec2843])
- Pinned all GitHub Actions to commit hashes and scoped workflow
  permissions ([3ec2843])
- Go directive bumped from 1.20 to 1.26, and tightened example file
  permissions ([baff0f6])

## [0.1.0]

### Added

- Initial release. Email decoder with MIME multipart parsing, charset
  decoding, and attachment callbacks ([aeed94b])

<!-- Commit links -->
[aeed94b]: https://github.com/LeakIX/go-emaildecoder/commit/aeed94b
[baff0f6]: https://github.com/LeakIX/go-emaildecoder/commit/baff0f6
[140bcf4]: https://github.com/LeakIX/go-emaildecoder/commit/140bcf4
[3ec2843]: https://github.com/LeakIX/go-emaildecoder/commit/3ec2843