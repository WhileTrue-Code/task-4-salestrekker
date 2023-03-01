# task-4-salestrekker

I used Golang,Docker and MongoDB to finish this task. I know there is one "mistake". I set .env variables DB_USERNAME and DB_PASSWORD,
because I wanted to imitate getting that data from configuring microservice which we can use in system
to get important parameters for regular service working.

/upload route accepts Input json as specification determined. if contact with specified information already exist in database but logically deleted it will be undeleted.


/get/{id} route accepts query param id and returns entity with specified id


/remove/{id} route also accepts query param id and logically deleting entity with specified id


/list route only returns all (undeleted) entities in system
