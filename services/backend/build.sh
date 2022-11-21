set -e

IMG_NAME="cupcake"
BUILD_CONTEXT_PATH="services/backend"
IMG_VERSION="001"
REGISTRY="harduim"

echo "Building"
docker build . -f $BUILD_CONTEXT_PATH/Dockerfile \
    --compress \
    --tag="${REGISTRY}/${IMG_NAME}:${IMG_VERSION}" \
    --tag="${REGISTRY}/${IMG_NAME}:latest" \
    --tag="${IMG_NAME}:${IMG_VERSION}" \
    --tag="${IMG_NAME}:latest"
echo ""

echo "Pushing: ${REGISTRY}/${IMG_NAME}:${IMG_VERSION}"
docker push "${REGISTRY}/${IMG_NAME}:${IMG_VERSION}"
docker push "${REGISTRY}/${IMG_NAME}:latest"
