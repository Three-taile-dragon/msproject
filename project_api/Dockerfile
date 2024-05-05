FROM alpine
WORKDIR /Initial
RUN mkdir -p /msProject/project-api/logs/debug
RUN mkdir -p /msProject/project-api/logs/info
RUN mkdir -p /msProject/project-api/logs/error
COPY ./target/project_api .
COPY ./config/config_docker.yaml .
COPY ./config/bootstrap_docker.yaml .
RUN  mkdir config && mv config_docker.yaml config/config.yaml && mv bootstrap_docker.yaml config/bootstrap.yaml
EXPOSE 80 2379 8848
ENTRYPOINT ["./project_api"]