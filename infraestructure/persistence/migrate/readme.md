every database will have your own folder.

every folder will have an **makefile** with some scripts:

> **migrate-up** - for create database
>
> **migrate-down** - for drop database
>


and for use the script just use:


``$ USER=yourdbuser PASS=pwdbuser make migrate-up``

or 


``$ USER=yourdbuser PASS=pwdbuser make migrate-down``

there is another way to use that, and also avaiable at the db folder that is a main.go file that can be used