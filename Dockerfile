######
## Stage 2: Build application with Go's build tools
######
FROM golang as build
ENV GO111MODULE=on
# https://stackoverflow.com/questions/36279253/go-compiled-binary-wont-run-in-an-alpine-docker-container-on-ubuntu-host
#ENV CGO_ENABLED=0 # cannot be set as otherwise plugins don't run
WORKDIR /app
COPY . /app
RUN go version
# RUN	go test ./...
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o raa.so raa/raa/raa.go
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o dummy.so raa/dummy/dummy.go
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o demo-rule.so risks/custom/demo/demo-rule.go
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -o threagile
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o secure-communication.so risks/custom/secure-communication/secure-communication.go
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o use-of-weak-cryptography.so risks/custom/use-of-weak-cryptography/use-of-weak-crypto.go
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o missing-monitoring-rule.so risks/custom/missing-monitoring/missing-monitoring-rule.go
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o accidental-logging-of-sensitive-data-rule.so risks/custom/accidental-logging-of-sensitive-data/accidental-logging-of-sensitive-data-rule.go
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o missing-audit-of-sensitive-asset-rule.so risks/custom/missing-audit-of-sensitive-asset/missing-audit-of-sensitive-asset-rule.go
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o credential-stored-outside-of-vault-rule.so risks/custom/credential-stored-outside-of-vault/credential-stored-outside-of-vault.go
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o insecure-handling-of-sensitive-data-rule.so risks/custom/insecure-handling-of-sensitive-data/insecure-handling-of-sensitive-data.go
RUN GOOS=linux go build -a -trimpath -ldflags="-s -w -X main.buildTimestamp=$(date '+%Y%m%d%H%M%S')" -gcflags="all=-trimpath=/src" -asmflags="all=-trimpath=/src" -buildmode=plugin -o running-as-privileged-user.so risks/custom/running-as-privileged-user/running-as-privileged-user.go

# add the -race parameter to go build call in order to instrument with race condition detector: https://blog.golang.org/race-detector




######
## Stage 3: Make final small image
######
FROM alpine

# label used in other scripts to filter
LABEL type="threagile"

# add certificates
RUN apk add ca-certificates
# add graphviz, fonts
RUN apk add --update --no-cache graphviz ttf-freefont
# https://stackoverflow.com/questions/66963068/docker-alpine-executable-binary-not-found-even-if-in-path
RUN apk add libc6-compat
# https://stackoverflow.com/questions/34729748/installed-go-binary-not-found-in-path-on-alpine-linux-docker
# RUN mkdir -p /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# clean apk cache
RUN rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build /app/threagile /app/threagile
COPY --from=build /app/raa.so /app/raa.so
COPY --from=build /app/dummy.so /app/dummy.so
COPY --from=build /app/demo-rule.so /app/demo-rule.so
COPY --from=build /app/missing-monitoring-rule.so /app/missing-monitoring-rule.so
COPY --from=build /app/accidental-logging-of-sensitive-data-rule.so /app/accidental-logging-of-sensitive-data-rule.so
COPY --from=build /app/missing-audit-of-sensitive-asset-rule.so /app/missing-audit-of-sensitive-asset-rule.so
COPY --from=build /app/credential-stored-outside-of-vault-rule.so /app/credential-stored-outside-of-vault-rule.so
COPY --from=build /app/insecure-handling-of-sensitive-data-rule.so /app/insecure-handling-of-sensitive-data-rule.so
COPY --from=build /app/running-as-privileged-user.so /app/running-as-privileged-user.so
COPY --from=build /app/use-of-weak-cryptography.so /app/use-of-weak-cryptography.so
COPY --from=build /app/secure-communication.so /app/secure-communication.so
COPY --from=build /app/LICENSE.txt /app/LICENSE.txt
COPY --from=build /app/report/template/background.pdf /app/background.pdf
COPY --from=build /app/support/openapi.yaml /app/openapi.yaml
COPY --from=build /app/support/schema.json /app/schema.json
COPY --from=build /app/support/live-templates.txt /app/live-templates.txt
COPY --from=build /app/support/render-data-asset-diagram.sh /app/render-data-asset-diagram.sh
COPY --from=build /app/support/render-data-flow-diagram.sh /app/render-data-flow-diagram.sh
COPY --from=build /app/server /app/server
COPY --from=build /app/demo/example/threagile.yaml /app/threagile-example-model.yaml
COPY --from=build /app/demo/stub/threagile.yaml /app/threagile-stub-model.yaml

RUN mkdir /data

RUN chown -R 1000:1000 /app /data
USER 1000:1000

ENV PATH=/app:$PATH
ENV GIN_MODE=release

ENTRYPOINT ["/app/threagile"]
CMD ["-help"]