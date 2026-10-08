# The runner as an OCI image, which is the channel section 13's distribution table gives
# CI. A pipeline that declares an image declares its runner version with it, where one
# that downloads a release binary resolves whichever version the download happens to
# return, and a verdict is reproducible only against a named runner.
#
# The binary is not built here. It is copied in from dist/, where the release has already
# built it with the version and the plugin digest stamped through ldflags, so the image
# carries the bytes SHA256SUMS covers rather than a second build of the same source that
# nothing compares against. Building this file by hand therefore means building the
# binary first, under the name the COPY below expects.
#
# Nothing in the image reaches the network at gate time. What it holds besides the runner
# is git and a certificate bundle, and each is there for a reason:
#
#   git, because the gate path reads the commit range through it. internal/git starts
#   `git log` and `git diff` over the two refs the run was given, which four of section
#   9's predicate types and `xeno intent verify` depend on. An image without git would
#   run `xeno gate verify` and fail every commit predicate, which is half of what an
#   image exists to carry into CI.
#
#   the certificate bundle, because `xeno enforcement check` is the first job of the
#   wrapper WP10 generates and the one command that talks to the host. It is not on the
#   gate path, so it does not weaken the sentence above; a TLS client with no roots would
#   simply make that job impossible to run from this image.

# Pinned by the digest of the manifest index rather than by a tag, because a tag is
# repointed by whoever owns it; docs/supply-chain.md makes that the rule for every action
# and tool in this pipeline and an image is no different.
#
# Debian slim and not a smaller musl base, which is a choice about the generated wrapper
# rather than about size. The GitHub wrapper names this image as the job's container, and
# the first thing that runs in it is actions/checkout, a JavaScript action the host
# executes with a node binary it mounts in. Whether the host mounts a musl build of node
# is a property of the runner image and not of this repository, so a glibc base removes a
# question that would otherwise be answered somewhere nobody here can see it.
#
# The cost is measured rather than guessed, because it is the larger half: the image is
# 195 MB unpacked against the 32 MB the same two packages come to on a musl base, and
# 57 MB of the difference is perl, which git depends on here and does not there. That is
# one pull per CI runner against a question nobody in this repository can answer, and the
# day somebody can run a container job on Alpine and watch it work, this paragraph is
# what the change deletes.
ARG BASE=docker.io/library/debian@sha256:a29215f6a35e51e22adffa17f89e9d2ef06214e64a2bad10d765c46aea49f11f
FROM $BASE

# --no-install-recommends, because what git recommends here is an ssh client, a pager and
# a patch tool, and the gate reads a local clone rather than fetching one. The package
# lists are removed in the same layer: an index in the image is a cache of a network the
# image is not supposed to need, and in a later layer it would still be in this one.
RUN set -eu; \
    export DEBIAN_FRONTEND=noninteractive; \
    apt-get update; \
    apt-get install -y --no-install-recommends ca-certificates git; \
    rm -rf /var/lib/apt/lists/*

# The gate runs as a user with an account rather than as root, and the uid is the whole of
# the difference between an image that works in a container job and one that only parses.
# Both hosts hand the image a clone they made themselves, so what matters is less that the
# user is unprivileged than that it is the uid of the tree it is given.
#
# 1001, because that is the uid GitHub's workspace belongs to. A container job bind mounts
# the runner's own `_work` directory at /__w and passes no --user, so the container runs as
# whatever this file declares, against files owned by the host's `runner` account; the
# runner's published image creates that account with `--uid 1001` and the hosted Ubuntu
# image uses the same number. The first thing in such a job is actions/checkout, which runs
# inside the container and writes the clone, so a user that cannot write /__w fails before
# any gate is reached. Matching the number is how this file gets a writable workspace
# without asking the host for a --user option that only the calling workflow can pass.
#
# safe.directory, because matching the uid is not something this file can guarantee and
# because one of the two hosts does not offer it at all. git refuses to read a repository
# owned by another uid — "detected dubious ownership" — and internal/git reaches git by
# starting `git log` and `git diff`, so a mismatch arrives as a failed verdict rather than
# as a permission error, which is the worse of the two. GitLab's docker executor clones in
# its helper container and leaves the build directory world writable with `umask 0000`, so
# there the writes succeed and the ownership mismatch stays;
# FF_DISABLE_UMASK_FOR_DOCKER_EXECUTOR changes the ownership instead and is off by default.
# This setting is therefore what makes the GitLab wrapper work at all, and what keeps the
# GitHub one working on a runner whose uid is not 1001.
#
# What it widens is said rather than implied. `*` turns the ownership check off for every
# repository, and that check is what stops git parsing the config of a repository it was
# pointed at by somebody else — a config that can name a pager, a hooks path or an
# fsmonitor, each of which is a command. On a shared machine that would be the wrong trade.
# Here the only repository is the clone the host made for this job: its `.git/config` is
# written by the host's checkout and not by the branch under review, which cannot reach it,
# and the uid that owns it is the host's own and not a third party's. Trusting that clone's
# config is the same trust as running the gate over it at all. It goes in the system config,
# which is where git requires this setting to be and which the user below cannot write.
RUN set -eu; \
    groupadd --gid 1001 xeno; \
    useradd --uid 1001 --gid 1001 --create-home --shell /bin/sh xeno; \
    install -d -o xeno -g xeno /work; \
    git config --system --add safe.directory '*'

# TARGETARCH is set by a builder that was told a platform and unset by one that was not,
# so the default names the architecture the release publishes. It is written as an
# argument rather than fixed into the path because the only thing a second architecture
# then needs is a builder that can produce it, and nothing in this file.
ARG TARGETARCH=amd64
# The version the release carries, which is also what names the binary to copy. The
# default is the unstamped version a local build produces, so that
# `CGO_ENABLED=0 GOOS=linux go build -o dist/xeno-0.0.0-dev-linux-amd64 ./cmd/xeno`
# is enough to build this image by hand; the release passes --build-arg VERSION.
ARG VERSION=0.0.0-dev
COPY dist/xeno-${VERSION}-linux-${TARGETARCH} /usr/local/bin/xeno

# No ENTRYPOINT. Both hosts run a list of commands in the image they are given, and an
# entrypoint of the runner itself would turn every one of those commands into an argument
# to `xeno`. The runner is on PATH and the shell is the default, which is what a GitLab
# `script:` block and a GitHub `run:` step both expect.
#
# The user is selected after the copy above rather than before it, so the runner stays owned
# by root and world executable: the account the gate runs as can start the binary and
# cannot write over it.
USER xeno
WORKDIR /work
