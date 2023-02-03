FROM alpine
WORKDIR /Initial
COPY ./target/project_api .
COPY ./config/config_docker.yaml .
RUN  mkdir config && mv config_docker.yaml config/config.yaml
EXPOSE 80 8881
ENTRYPOINT ["./project_api"]