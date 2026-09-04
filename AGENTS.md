# AGENTS.md - Guidance for AI Agents

## Critical Rules

- **Follow ASD-STE100 writing guidelines** for your prose and documentation.
- **NEVER push directly to `master`** - Any change must be done through a
  patch with a new branch. Create a feature branch, commit your changes,
  and open a PR.
- **NEVER use `/` in branch names** - Use `-` or `_` as separators instead.
- **Keep PRs small and reviewable.** One branch and one PR per logical
  unit of work.
- **Add a CHANGELOG entry for every PR, in its own commit.** The changelog
  entry MUST always be a separate commit from the code it describes.
- **When cutting a tag/release, the CHANGELOG MUST be updated in the same
  PR.** Move the entry from `[Unreleased]` to the version heading.
- **Record the model in the commit message and the PR.** A trailing line
  such as `Model: <provider>/<model-id>`.
- **Do not reference plan or task numbers in commits, PRs, or the changelog.**
- **Verify the full CI pipeline is green before considering a branch done.**

## Security

- **NEVER add any new dependency or update a dependency** without the
  explicit consent of the prompt engineer.

## Key Architectural Patterns

### Decoder

The `Decoder` parses a raw email (`io.Reader`) into an `EmailContent`
(HTML, plain text, headers). It walks MIME multipart trees recursively,
decodes transfer encodings (base64, quoted-printable) and charsets
(via `golang.org/x/text`), and hands attachments to a callback.

### Attachment

`Attachment` wraps an `io.Reader` with a filename and content type.
The decoder assigns a synthetic filename (`attachement-N.file`) when the
MIME part has no usable filename.

## Dependencies

- `golang.org/x/text` - charset decoding (ISO-8859, Windows, KOI8, etc.)

## Development Commands

```bash
make help          # Show available targets
make vet           # Run go vet
make test          # Run tests
make lint          # Run golangci-lint
make vuln          # Run govulncheck
make check-format  # Check Go formatting
make check-changelog # Check changelog commit references
```