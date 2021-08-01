# golang-web-project-framework
Golang Web Project Framework is standards and specifications of web project, which defines project framework, directory structure and so on.
It makes web project clearer and easier to understand, and ensures the unity of style when developers code.



## Project Framework
![image](golang-web-project-framework.png)


### Modules
Split by module or function first, then layer the module.

Tips:

1.layer the project first, it will be large and redundancy in the late stages.

2.As the project gets bigger and bigger, maybe some modules will be independent.

#### api
It contains router, and handlers which are responsible for parsing paramters.

#### service
It is responsible for validate paramters and business logic processing, then invokes model layer.
Attention: Never allow use model or dao B through Service A, but Service A to Service B, then Service B invokes model B.

#### model
It indicates concrete instantiated object, which matches with database.

#### dao
Encapsulating access to the database: adding, deleting, modifying and querying.Do not involve business logic.


### Config
Public Config, which can be used by any others.

### Utils 
Public function part, which can be used by any others.



## To be Continue...
