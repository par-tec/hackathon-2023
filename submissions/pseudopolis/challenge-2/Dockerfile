FROM golang AS build-stage
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux GOBIN=/ go install github.com/albertobregliano/robodoc@v0.2.0

FROM gcr.io/distroless/base-debian11 AS build-release-stage 
LABEL author="Alberto Bregliano"
WORKDIR /
COPY --from=build-stage /robodoc /robodoc
ENTRYPOINT [ "/robodoc" ]
CMD ["dottore sto male"]