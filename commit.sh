#!/bin/zsh

DIR="${HOME}/Users/shimozawakenta/project/workManager/"
path="app/"
new="commit.go"

cd ~/project/workManager/app/
echo "" >> ~/project/workManager/app/commit.go

cd ~/project/workManager
git add .
git commit -m 'fix'
git push origin HEAD