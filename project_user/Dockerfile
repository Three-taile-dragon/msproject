FROM alpine
WORKDIR /Initial
RUN mkdir -p /msProject/project-user/logs/debug
RUN mkdir -p /msProject/project-user/logs/info
RUN mkdir -p /msProject/project-user/logs/error
COPY ./target/project_user .
COPY ./config/config_docker.yaml .
COPY ./config/bootstrap.yaml .
RUN  mkdir config && mv config_docker.yaml config/config.yaml && mv bootstrap.yaml config/bootstrap.yaml
EXPOSE 3457 8881 8848 2379 3309
ENTRYPOINT ["./project_user"]