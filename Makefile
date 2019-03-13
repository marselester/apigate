build_docker_auth:
	docker build --tag=marselester/travel-auth:v1.0.0 --file=docker/auth.Dockerfile .

build_docker_ratelimit:
	docker build --tag=marselester/travel-ratelimit:v1.0.0 --file=docker/ratelimit.Dockerfile .

build_docker_hotel:
	docker build --tag=marselester/travel-hotel:v1.0.0 --file=docker/hotel.Dockerfile .

build_docker_car:
	docker build --tag=marselester/travel-car:v1.0.0 --file=docker/car.Dockerfile .
