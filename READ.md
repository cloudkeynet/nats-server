# Nats-project
This is a project with a arquitecture in AWS based in micro-services, will be handled by ECS1 containers and tasks, using Amazon elastic container register,
the focus of the project is to get a good data manage and a solid communication with the user, being versatile and useful for anyone business regardless the
area that its cover. Various tools were used for its operation, for its security, error handling, permission, inter alia. For security and permission it will 
be used two tools which are AWS Certificate Manager to generate a certificate SSL and don´t have any issue in development or deploy, the other tool is AWS IAM 
who are in charge of role manage and permissions, AWS CloudWatch are used for the error handling and the way to send messages is through AWS SNS(all the messages
will be sent to the email that was registered).


## Starting 🚀
For this project you must have some requirement for the optimal performance.

## Prerequisite 📋
First of all, you have to clone the repository to work on it, you can made it with github desktop, sing in the account, click the buttom "Clone the repository", 
select the repository, select a folder and you´re ready. Other way to clone the repository is with visual studio code, you have to sing in your github account, 
after you select the source control, click the buttom "Clone Repository", select the repository, select a folder and you have a copy to work.

Also you must to have with an AWS account, exist so many ways to get an AWS account like have an empresarial account, have a university agreement or buy the licence.
The way to get the account doesn´t matter, all of that can use the tools that are used in the project.

## Construction Process on AWS 🔧
-You go to Route 53 and click the number below "DNS administration", this is the section of the host routes, you have to identify the route that 
you will use and copy the name of the route.
-After enter to AWS Certificate Manger and click the buttom "Request certificate", select the option "request a public certificate", in the section "Domain name" 
you have paste the name of the route and click "request"
-Enter to CloudFormation, click "create stack" and select "With new resources", in "Prerequisite" click "Choose an existing template", next in the section of 
"Specify template" choose the option "Upload a template file" and select the respective file ".yml", name the stack, click the section of "Notification options" 
and choose "Create new SNS Topic", name the SNS topic and sign in your email and finally click submit.
-Go to VPC and choose the section of "your VPC´s" and create vpc, select the option "VPC only", named the vpc, in IPv4 CIDR wrote "10.0.0.0/24" and create VPC.

## Built with 🛠️
The tools that are used for the project are from the aws service

Route53 - Identify the host route
Certificate Manager - Create certificate SSL
CloudFormation - Create the stack that will be used
VPC - The virtual network that will be used in the project

## Author
Cloudkeynet

Daniel morales - Director of project - Developer
Sergio Parra - Documentation - Developer

## LICENSE 📄
The license of the proyect is [License](https://github.com/cloudkeynet/nats-server.git/LICENSE).
