# Bookstore
Monolithic Backend server system.

# How to run
docker run -it -p 8000:8080 --mount type=bind,source="$(pwd)"/,target=/go/src/go-docker --name altsch-go-docker  altsch/go-docker

# Note
Due to varying versions of dependencies in the vendor file,I had to remove it.(ps. I'm still looking for solution to this )
