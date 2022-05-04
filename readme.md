Studie repo for mariaDB and DDD aproach

Folder struct:

> **cmd/** = every app that communicate with core and need some infra it's here 
>
>> **cmd/web** = web app is here
>>
>> **cmd/web/main.go** = main app
>>
>>> **cmd/web/handlers** = all the handlers are here
>>>
>>> **cmd/web/handlers/beer.go** = all endpoints definitions are here
>>>
>>> **cmd/web/handlers/error.go** = all errors for handlers are defined here
>>
> **core** = business logic find here
>
>> **core/beer** = all logic for beer model
>>
>>> **core/beer/contract.go** = all interfaces for service and repository (infra)
>>>
>>> **core/beer/entity.go** = business model
>>>
>>> **core/beer/service.go** = services with business logic
>>>
> **infraestructure** = only definitions, access and static assets
> 
>> **infraestructure/migrate** = especifics migrations files
>>
>>>  **infraestructure/migrate/mariadb** = migrations for MariaDB
>>>
>>>> **infraestructure/migrate/mariadb/main.go** = function for up migrations find at /migration folder
>>>>
>>>> **infraestructure/migrate/mariadb/MakeFile** = script for execute migrations (see README.md)
>>>>
>>>>> **infraestructure/migrate/mariadb/migrations** = migrations sql files
>>>>>
>>>>> **infraestructure/migrate/mariadb/migrations/000001_create_beer_table.down.sql** = down sql cmd
>>>>>
>>>>> **infraestructure/migrate/mariadb/migrations/000001_create_beer_table.up.sql** = up sql cmd
>>>>>
>>
>> **infraestructure/repository** = all repository specs and files
>> 
>>> **infraestructure/repository/driver.go** = definition of the driver to access database (MariaDB)
>>>
>>> **infraestructure/repository/beer.go** == beer repository

Be happy :D