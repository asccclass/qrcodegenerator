#!/bin/sh

echo "Clean <none> images..."
docker rmi $(docker images | grep "none" | awk '{print $3}')
docker images
echo "done."
