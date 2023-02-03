FROM alpine
WORKDIR /Initial
COPY ./target/project_user .
COPY ./config/config_docker.yaml .
RUN  mkdir config && mv config_docker.yaml config/config.yaml
EXPOSE 8080 8881
ENTRYPOINT ["./project_user"]