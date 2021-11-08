# Verification using Redis database
In this program, I used the Redis database to store users' information temporarily. 
The process of the program is following: 

1) Firstly, the new users will sign up using the /signup endpoint. When the new user enters his username and email, 
the program will send 6 digits random number for his email. At this time, his username, email and password will also be inserted 
into the Redis database. *NOTE: you have to verify your account in an hour, because after one hour, the random number will expire.

2) Secondly, after the process of signing up, the user verifies his information using the /verify endpoint.
To do this, the user will enter his username and password, the program checks his password with the password in the
Redis database, if it matches correctly, then his username, email, and password will be inserted into the 
PostgreSQL database.
