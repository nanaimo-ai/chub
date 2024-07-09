move-stuff:
	scp -r . v@192.168.1.14:/home/v/compute-hub

build:
	sudo docker build --no-cache -t chubbir/test .

run:
	sudo docker run --gpus all -e CHUB_CONTAINER_REGISTRY=compute-hub \
	-e CHUB_KEY=clyc33zea00008fsfr9h4blvu -e CHUB_ID=adsadadasd \
	-e CHUB_HTTP_TK=adaddasdasdad -e CHUB_ENV=local \
	-e CHUB_TB_TK=cC5leUoxSWpvZ0lqSTNOMkpqTm1KakxUSTFNbUl0TkRnNU5TMWhZVGd4TFdZeVpXTXdNamcxTkRBMU5DSXNJQ0pwWkNJNklDSmxNMkZpWlRabE1TMDJaamxtTFRSak5qRXRZV1JoTnkwMVkyRmhZMkV5TVRkak9UZ2lMQ0FpYUc5emRDSTZJQ0oxY3kxbFlYTjBMV0YzY3lKOS5CNXk2eWlkZ3VielRnV0FxWlB6Z1M0QVF6Q2tISV9fdnVlUHhPRjhGa0Z3Cg== \
	chubbir/test