#!/bin/bash

#DIR="~/project/workManager/"

DIR="${HOME}/Users/shimozawakenta/project/workManager/"
new="app/commit.go"

echo "" >> $DIR$new

cd $DIR
git add .
git commit -m 'fix'
git push origin HEAD