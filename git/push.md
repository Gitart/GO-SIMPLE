
1.Create a local Git and .git folder in the root of the project folder:  
**$ git init**

2.Add all of the files to the repository and start tracking them:  
**$ git add .**

3.Make the first commit:  
**$ git commit -am "initial commit"**

4.Add the GitHub remote destination:   
**$ git remote add your-github-repo-ssh-url**

It might look something like this:   
**$ git remote add origin git@github.com:azat-co/simple-message-board.gi**

5.Now everything should be set to push your local Git repositoryto the remote destination on GitHub with the following command:
**$ git push origin master**  

6.You should be able to see your files at github.com under your account and repository.
Later, when you make changes to the file, there is no need to repeat all of these steps.

## Just execute:   
$ git add .   
$ git commit -am "some message"   
$ git push origin master   

## If there are no new untracked files you want to start tracking, use this:
$ git commit -am "some message"   
$ git push origin master  

To include changes from individual files, run:
$ git commit filename -m "some message"   
$ git push origin master   

## To remove a file from the Git repository, use:
$ git rm filename   

For more Git commands, see:
$ git --help   

Deploying applications with Windows Azure or Heroku is as simple as pushing code
and files to GitHub. The last three steps (4–6) would be substituted with a different remote
destination (URL) and a different alias.
