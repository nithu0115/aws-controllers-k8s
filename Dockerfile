# Build the manager binary
FROM golang:1.14.1 as builder

ARG service_alias
ARG work_dir=/github.com/aws/aws-controllers-k8s
WORKDIR $work_dir
# For building Go Module required
ENV GOPROXY=https://proxy.golang.org,direct
ENV VERSION_PKG=github.com/aws/aws-controllers-k8s/pkg/version
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN  go mod download
# Copy the go source
COPY . $work_dir/
# Build
RUN GIT_VERSION=$(git describe --tags --dirty --always) && \
    GIT_COMMIT=$(git rev-parse HEAD) && \
    BUILD_DATE=$(date +%Y-%m-%dT%H:%M:%S%z) && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on \
    go build -ldflags="-X ${VERSION_PKG}.GitVersion=${GIT_VERSION} -X ${VERSION_PKG}.GitCommit=${GIT_COMMIT} -X ${VERSION_PKG}.BuildDate=${BUILD_DATE}" -a -o $work_dir/bin/controller $work_dir/services/$service_alias/cmd/controller/main.go

FROM amazonlinux:2
ARG work_dir=/github.com/aws/aws-controllers-k8s
WORKDIR /
COPY --from=builder $work_dir/bin/controller $work_dir/LICENSE $work_dir/ATTRIBUTION.md /bin/
ENTRYPOINT ["/bin/controller"]
