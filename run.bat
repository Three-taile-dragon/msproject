chcp 65001
cd project_user
docker build -t project_user:latest .
cd ..
docker-compose up -d