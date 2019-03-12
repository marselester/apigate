# API gateway demo

[API gateway](https://docs.microsoft.com/en-us/azure/architecture/microservices/design/gateway)
acts as a reverse proxy, routing API requests from clients to services.
Usually it also performs authentication and rate limiting, so the services behind the gate don't have to.

This demo is based on a dummy [Traveling project](https://traveling.docs.apiary.io/)
where we have services to rent a car and book a hotel.
Check out short tutorials:

- [Traefik as API Gateway](http://marselester.com/apigate-traefik.html)
