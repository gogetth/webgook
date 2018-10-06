#!/bin/sh
docker pull aorjoa/geekybase-api
docker stop geekybase-api-container
docker run -d --rm --name geekybase-api-container aorjoa/geekybase-api
