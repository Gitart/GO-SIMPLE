
@echo off
echo Push to assembla

cd d:\HOGIT\hdf-server\removeservice\
rem https://git.assembla.com/hdf-server.git


rem 
rem fetch

rem stash        -- свои изменения в промежуточный буфер
rem stash applay -- перенос состояния из буфера к себе в дерево

rem git pull --rebase  --

rem Установка на последний коммит 
rem git pull
git add -A
git commit -m "Added new commit"
git push -u origin master

pause
