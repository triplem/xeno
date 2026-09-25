# For one run only. Trivy's misconfiguration scanner reports a container that runs
# as root, which is a finding it produces from a file rather than from a database,
# so the failing path can be provoked without a vulnerable dependency.
FROM scratch
COPY xeno /xeno
ENTRYPOINT ["/xeno"]
