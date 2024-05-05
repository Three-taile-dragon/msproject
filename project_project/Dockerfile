FROM alpine
WORKDIR /Initial
RUN mkdir -p /msProject/project-project/logs/debug
RUN mkdir -p /msProject/project-project/logs/info
RUN mkdir -p /msProject/project-project/logs/error
COPY ./target/project_project .
COPY ./config/config_docker.yaml .
COPY ./config/bootstrap.yaml .
RUN  mkdir config && mv config_docker.yaml config/config.yaml && mv bootstrap.yaml config/bootstrap.yaml
EXPOSE 3458 8882 2379 6376 3309 8848
ENTRYPOINT ["./project_project"]