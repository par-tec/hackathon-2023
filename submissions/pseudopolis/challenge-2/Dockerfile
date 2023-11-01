FROM golang AS build-stage
WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux GOBIN=/ go install github.com/albertobregliano/hackathon-2023/submissions/pseudopolis/challenge-2/robodoc@latest

FROM gcr.io/distroless/base-debian11 AS build-release-stage 
WORKDIR /
COPY --from=build-stage /robodoc /robodoc
ENV OPENAI_KEY=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
ENTRYPOINT [ "robodoc" ]
CMD ["dottore sto male"]