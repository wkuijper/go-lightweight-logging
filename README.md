# go-lightweight-logging

Lightweight logging library for go.

Piggybacks on module system and semver pre-release tags for the most
efficient loglevel trimming possible at compile time (absent actual
code-generation or build tags).

In contrast to build tags this approach does not require any change to
source files or meta-conventions in the form of an agreed upon set of
build tags.

To be used in conjunction with the go-lightweight-log<level> modules.

By includeing the versions with pre-release tag `enabled` (i.e.:
`v0.2.0-enabled`) the loglevel (error, warning, info, etc.)  will be
enabled and code will be generated at each call site.

By including the versions with pre-release tag `disabled` (i.e.:
`v0.2.0-disabled`) the loglevel will be disabled, all the functions
will have an empty function body which means the compiler (at least
for recent versions of go) will not generate any code at the call site
thanks to function inlining and dead code elimination.

This common package contains some further utility code for routing log
messages.
