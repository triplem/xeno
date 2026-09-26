# Security

## Reporting a vulnerability

Do not open a public issue. Use GitHub's private vulnerability reporting, on this
repository's Security tab, with a description and, where possible, a way to reproduce
it. It opens a channel only the maintainers can read, and it keeps the report with the
repository rather than in somebody's mailbox, which is why no address is named here.

## Scope

The runner evaluates gates without network access and runs no model. Findings about
the gate path, the hashing in Appendix B of the process definition, the handling of
credentials read from the environment and the CI workflow are in scope.

## Releases

Releases carry checksums. Signatures arrive with publication, as the implementation
plan sets out under WP0.
