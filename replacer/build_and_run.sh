DIR="$(dirname "$(realpath "$0")")"
cd $DIR

docker build -t replacer .
docker run --rm replacer