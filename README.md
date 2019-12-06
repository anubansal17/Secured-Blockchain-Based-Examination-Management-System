# Secured-Blockchain-based-Examination-management-System

## INTRODUCTION

### Problem Statement
In a traditional examination management system, request for setting question papers and the actual question paper from Controller Of Examination (COE) to various teachers and vice-versa , is done physically. Further, the work of sending the now finalized paper from COE to various superintendents of examination is also done physically. </br>
This system generates a lot of security, economic and financial concerns , like - </br>
- After setting the question paper the confidentiality and integrity is totally in the hands of a third person
- The financial cost of sending a physical request to teacher and the teacher sending it back physically is too much 
- Moreover the question papers sent to different examination centers to Superintendent adds greatly to this cost
- The confidentiality of teacher , i.e. the teacher who designed a particular question paper , is not maintained</br>
There is no way we could solve these concerns without spending more money on this. So there is a need of automation in this field also , an automated system which covers all these concerns in a cost effective way. 



### Proposed Solution

The traditional examination management is a dependent system .In today's automated world , we need automation in this field also.
To design an automated system we need -
- An environment/portal where we can send and receive requests 
- A platform where we can store our question papers. It should be compatible enough so that our superintendent can retrieve the corresponding question paper only
- A discreet way of finalizing a question paper without knowing the teacher
- As for security concerns for our automation, we should apply symmetric and asymmetric encryption wherever necessary

So, we designed a system which fulfills all of the above criteria with many different technologies working together as one system to give the desired functionality. We used a combination of block chain and IPFS to achieve decentralization and secure peer-to-peer communication between teacher and superintendent. Encrypting question papers and our packets to ensure more safety is also used in python framework, for compatibility. The web development is done with the help of Django. 



## Technology Used
- Ubuntu 16.04
- Hyperledger Composer and Fabric for Blockchain Development:-
- *Hyperledger Fabric* is a platform for distributed ledger solutions underpinned by a modular architecture delivering high degrees of confidentiality, resiliency, flexibility and scalability. It is designed to support pluggable implementations of different components and accommodate the complexity and intricacies that exist across the economic ecosystem.
- *Hyperledger Composer* is a set of collaboration tools for building blockchain business networks that make it simple and fast for business owners and developers to create smart contracts and blockchain applications to solve business problems.

- Inter Planetary File System (IPFS) for Data Storage
IPFS is a peer-to-peer distributed file system that seeks to connect all computing devices with the same system of files. In few ways IPFS is similar to the World Wide Web, but IPFS could be seen as a single BitTorrent swarm exchanging objects within one Git repository. IPFS has no single point of failure and nodes do not need to trust each other except for every node they are connected to. Distributed Content Delivery saves bandwidth and prevents DDos attacks, which HTTP struggles with.
- Javascript:- for implementing logical capabilities in Blockchain network.
- HTML/CSS for Front-end purposes
- Django for Back-end purposes
- SSl for Certificate Creation
- RSA for Encryption/Decryption
- Web Services like Azure or AWS:- These services might be used for deploying our Blockchain based system on cloud.

## Flow Chart
![Image of Flow Chart](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Flow%20Diagram.png)

## A quick walkthrough of the project
- The Controller of Examinations (COE) from the datesheet in his local storage,  sends requests to the designated teachers for that subject to set the question paper
- As the COE send the request the teachers, they accept it and set the paper on their local machine in a pdf format
- Now this pdf file when prepared and encrypted with symmetric key, is sent to the distributed file system which is IPFS
- IPFS divides the file into various chunks. These chunks are stored on a decentralized server and each chunk's hash is generated. The hashes are combined to give one hash to retrieve the file
- Further, this hash along with other particulars is sent to be stored in block chain ledger
- Submitted papers start showing on COE portal without any details of paper or teacher. So COE finalizes one paper randomly and sends to superintendent
- The superintendent of center receives these details. Superintendent gets details of question paper retrieval from block chain ledger
- After getting the details automatic request to retrieve paper from IPFS is sent and then superintendent gets the original paper 

## Testing
### Unit Testing:-
Unit Testing is a level of software testing where individual units/ components of a software are tested. The purpose is to validate that each unit of the software performs as designed. Various modules used in the system are tested individually through this testing.
We had 3 modules in our project which are described as follows-</br>

1) Blockchain Module: Testing was performed only on the working of the blockchain system to check if all the functionalities provided by blockchain system are running fine or not. Initially, we used Hyperledger Fabric for the blockchain network implementation and tested the system if we were able to do the transactions in the Blockchain Ledger.</br>
**Status after testing - Success** 

2) Web + Data Storage Module: Testing was performed to check if all the functionalities are working fine on the front-end part of the website and if all the data involved in the system is storing on the local database or online distributed database accordingly.</br> 
**Status after testing - Success**

3) Security Module: Testing was performed to check if the techniques like symmetric and asymmetric encryption used to enhance the security of the system are working fine and to check vulnerabilities in these techniques, if any.</br>
**Status after testing - Success**

### Integration Testing:-
Integration Testing is a level of software testing where individual units are combined and tested as a group. The purpose of this level of testing is to expose faults in the interaction between integrated units. 
Above mentioned three modules are integrated together to check the compatibility of all three modules with each other.
To do this, first of all security module was integrated with web + data storage module and then all the functionalities provided by these modules are tested.</br>
**Status after testing - Success**
Secondly, we tried to integrate Blockchain module with already integrated modules(web and security) to check the compatibility of these modules with each other. After testing we came to know that interaction between Django web server and Hyperledger Fabric Blockchain server is not possible directly,i.e. These can not be integrated directly.</br>
**Status after testing - Failed**

### System Testing:-
System Testing is a type of software testing that is performed on a complete, integrated system to evaluate the compliance of the system with the corresponding requirements.
After integration testing, we searched for some other work around which can be used to integrate the Blockchain module with (web+security) module and we got to know that Hyperledger Composer can be used to integrate both the modules. Hyperledger Composer made the interaction between the web and Blockchain server smooth. All the functionalities provided by the system are working as per the requirements.</br>
**Status after testing - Success**

## Project Snapshots
### 1) Login Page
![Login page](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Login.PNG)
### 2) COE Dashboard
![COE Dashboard](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/COE%20Dashboard.png)
### 3)Request Details
![Request Details](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Request%20Details.png)
### 4)Request Status
![Request Status](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Request%20Status.png)
### 5)Teacher Dashboard
![Teacher Dashboard](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Teacher%20Dashboard.PNG)
### 6)Upload Paper
![Upload Paper](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Upload%20Paper.PNG)
### 7)Composer Rest Server GUI
![Composer Rest Server GUI](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Composer%20Rest%20Sever%20GUI.PNG)
### 8)Request Status- Uploaded
![Request Status Uploaded](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Request%20Status-Uploaded.PNG)
### 9)COE Finalization of paper
![COE Finalization of paper](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/COE%20Finalization%20of%20paper.PNG)
### 10)Request Status- Finalized
![Request Status Finalized](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Request%20Status%20-%20Finalized.PNG)
### 11)Superintendent Dashboard
![Superintendent Dashboard](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Superintedent%20Dashboard.PNG)
### 12)Paper ready for Download
![Paper Ready for Download](https://github.com/anubansal17/Secured-Blockchain-Based-Examination-Management-System/blob/master/images/Paper%20ready%20for%20Download.PNG)

## Future Scope 
This project has a very vivid future scope since decentralised safe storage of data is today’s biggest need.</br>

### Project Upgradation 
We have the project on our local host as of now, it is the first iteration of a working prototype. There is a lot of scope for improvement and upgradation.
- Project can be properly deployed on web
- Openssl certificates could be deployed on the web server to make every transaction secure
- Kerberos can be used for high end authentication
- Time stamp could be implemented on Blockchain composure for ensuring no access to question paper before allowed time i.e 30 minutes before the exam

### Future scope of application
This project can be sold as a full fledged enterprise web based tool to conduct safe and secure examinations across various colleges and Universities across India, and overseas later.
This project can find relevance and use in many other sectors such as-
- Healthcare sector to store confidential information of patients on a decentralised database
- In human resource departments to store confidential information about employees
- Police Department can use it to store very crucial information about the Criminals and Crime scenes confidentially
Similarly almost every sector can use this product to store their sensitive data/information which needs high security.

## Installation Guide</br>
Follow the Hyperledger Fabric installation guide from here - https://hyperledger-fabric.readthedocs.io/en/release-1.4/prereqs.html
Follow the Hyperledger Composer installation guide from here - https://hyperledger.github.io/composer/latest/installing/development-tools.html</br>
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

You might get errors while installing Hyperledger Fabric, some of them are listed below with solutions-</br>
Docker-got permission denied :- https://www.digitalocean.com/community/questions/how-to-fix-docker-got-permission-denied-while-trying-to-connect-to-the-docker-daemon-socket </br>

You might get errors while installing web server, some of them are listed below with solutions- </br>
1)Error while installing psycopg2 - https://stackoverflow.com/questions/5420789/how-to-install-psycopg2-with-pip-on-python</br>
2)Error while collecting static files - https://stackoverflow.com/questions/27323350/django-collecstatic-with-crontab
