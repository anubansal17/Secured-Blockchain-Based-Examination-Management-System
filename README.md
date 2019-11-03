# Secured-Blockchain-based-Examination-management-System

## Installation Guide</br>
Follow the Hyperledger Fabric installation guide from here - https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html

You might get errors while installing Hyperledger Fabric, some of them are listed below with solutions-</br>
1)Docker-got permission denied :- https://www.digitalocean.com/community/questions/how-to-fix-docker-got-permission-denied-while-trying-to-connect-to-the-docker-daemon-socket

Follow the instructions mentioned below for the installation of web server-</br>
https://www.digitalocean.com/community/tutorials/how-to-use-postgresql-with-your-django-application-on-ubuntu-14-04

Changes to be made in the commands- </br>
  myproject to ems</br>
  myprojectuser to admin_ems</br>
  password to iamadmin@123</br>
  UTC to Asia/Kolkata</br>

Now open another terminal window and create a folder EMS wherever you want.(Home directory is preferable)</br>
'mkdir EMS' is the instruction to create a folder

Write 'cd EMS' to enter that folder</br>

Now create a virtual environment using instruction 'virtualenv -p python3 env'</br>
Write 'cd env' to enter the env</br>

Write 'source bin/activate' to activate the env </br>

Create new folder 'project' using 'mkdir project'</br>

Enter project folder using 'cd project'</br>

Clone the project using 'git clone https://github.com/deveshd2k/clgproject.git'</br>

Enter the project using 'cd clgproject'</br>

Install all the requirements using 'pip install -r requirements.txt'</br>

Now run migrations using 'python manage.py migrate'</br>

Collect static files using 'python manage.py collectstatic'</br>

Now we cannot run the local host without ipfs daemon running.</br>
For that follow the below instructions :</br>

Install ipfs according to your system from 'https://dist.ipfs.io/#go-ipfs'</br>
Then extract the files somewhere in your system. Don't put it in the same directory as our project</br>
Open new terminal and enter the folder named 'go-ipfs' using cd 'go-ipfs'</br>
Write 'sudo ./install.sh' to move the file</br>
Check whether ipfs has been installed using 'ipfs version'</br>
Initialize Node using 'ipfs init'. Error will come but ignore</br>
Write 'ipfs daemon' to run the daemon'</br>

Now go back to the previous terminal where our main project was and run 'python manage.py runserver'. Do not close any terminal except the postgresql one.</br>

If the localhost is up and running, then kudos ;)</br>
Else discuss the errors

You might get errors while installing web server, some of them are listed below with solutions- </br>
1)Error while installing psycopg2 - https://stackoverflow.com/questions/5420789/how-to-install-psycopg2-with-pip-on-python</br>
2)Error while collecting static files - https://stackoverflow.com/questions/27323350/django-collecstatic-with-crontab
