@echo off
echo Init Assemblia

cd D:\HOGIT\

git init
git remote add repo https://login:password@git.assembla.com/hdfserver.git

git config --global user.name  "login" 
git config --global user.email "mail@gmail.com"
git config credentinal.helper cash

git clone https://login:password$@git.assembla.com/hdf-server.git

pause
