# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

## [0.1.0]

### Added

- Initial release. `Decoder` type that parses raw email into
  `EmailContent` (HTML, PlainText, Headers) with attachment
  callbacks ([aeed94b])
- Inline attachment support ([5c7c0c4])
- Support for emails without a media type ([91aeedf])
- Attachment reader as `io.Reader` ([916674a])

### Changed

- Go directive bumped from 1.20 to 1.26 ([8fe189e])
- Magic strings extracted into named constants ([cd45dcd])
- Added Makefile, .drone.yml, .golangci.yml, AGENTS.md

<!-- Commit links -->
[aeed94b]: https://github.com/LeakIX/go-emaildecoder/commit/aeed94b
[5c7c0c4]: https://github.com/LeakIX/go-emaildecoder/commit/5c7c0c4
[91aeedf]: https://github.com/LeakIX/go-emaildecoder/commit/91aeedf
[916674a]: https://github.com/LeakIX/go-emaildecoder/commit/916674a
[8fe189e]: https://github.com/LeakIX/go-emaildecoder/commit/8fe189e
[cd45dcd]: https://github.com/LeakIX/go-emaildecoder/commit/cd45dcd