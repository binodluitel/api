# --build-arg to be carried over to other Docker stages
ARG APP_NAME=${APP_NAME}
ARG BUILD_TIME=${BUILD_TIME}
ARG GIT_REF_NAME=${GIT_REF_NAME}
ARG GIT_REF_SHA=${GIT_REF_SHA}
ARG VERSION=${VERSION}

# build stage
FROM golang:1.23rc1-bullseye AS build-stage
ARG GITHUB_USER
ARG GITHUB_TOKEN

ARG APP_NAME
ARG BUILD_TIME
ARG GIT_REF_NAME
ARG GIT_REF_SHA
ARG VERSION

WORKDIR /go/src/github.com/binodluitel/api
COPY . .

# Inject the GitHub user/token into all Git requests so we can run go mod download. This is typically
# only used when GitHub actions builds the Docker image
RUN if [ -n $GITHUB_USER ]; then \
    git config --global url."https://${GITHUB_USER}:${GITHUB_TOKEN}@github.com/binodluitel/".insteadOf https://github.com/binodluitel/; \
    fi

# If there's a .gocache directory present, we need to point (symlink) the Go mod/build cache to
# the .gocache directories so that Go uses it
RUN if [ -d .gocache ]; then \
    mkdir -p /root/.cache/ /go/pkg/; \
    ln -s $(pwd)/.gocache/root/.cache/go-build /root/.cache/go-build; \
    ln -s $(pwd)/.gocache/go/pkg/mod /go/pkg/mod; \
    fi

# Run go mod download prior to the build and pass in the SSH agent (if provided)
RUN --mount=type=ssh go mod download
RUN go build \
    -ldflags "-X main.Version=${VERSION} -X main.Commit=${GIT_REF_SHA}" \
    -o ./build/_output/bin/api \
    ./main.go

# Final stage
FROM registry.access.redhat.com/ubi8/ubi:8.10
RUN yum -y update && yum install ca-certificates -y && yum clean all -y
RUN groupadd apiuser
RUN useradd -r -g apiuser apiuser
USER apiuser

LABEL maintainer="https://github.ibm.com/binodluitel/api"
LABEL description="This is an example API service Docker Image"

WORKDIR /app/
COPY --from=build-stage /go/src/github.com/binodluitel/api/build/_output/bin/api /bin/api

# Set environment variables from build time
ARG APP_NAME
ARG BUILD_TIME
ARG GIT_REF_NAME
ARG GIT_REF_SHA
ARG VERSION

ENV APP_NAME=${APP_NAME}
ENV APP_BUILD_TIME=${BUILD_TIME}
ENV APP_VERSION=${VERSION}
ENV APP_GIT_REF_NAME=${GIT_REF_NAME}
ENV APP_GIT_REF_SHA=${GIT_REF_SHA}

ENTRYPOINT [ "/bin/api" ]
