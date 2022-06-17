echo "build jetbrains start"

docker buildx build -t toodo/jetbrains:$1 --platform=linux/arm64,linux/amd64 -f release/docker/Dockerfile . --push

docker buildx build -t toodo/jetbrains:latest --platform=linux/arm64,linux/amd64 -f release/docker/Dockerfile . --push

echo "build jetbrains finish"
