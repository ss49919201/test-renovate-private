#!/bin/bash

echo "# test-renovate-private" >> README.md && git init && git add README.md && git commit -m "first commit" && git branch -M main && git remote add origin https://github.com/ss49919201/test-renovate-private.git && git push -u origin main
