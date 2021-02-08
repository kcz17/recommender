NAME = kcz17/recommender
DBNAME = kcz17/recommender-db

INSTANCE = recommender

.PHONY: default copy test

default: test

release:
	docker build -t $(NAME) -f ./docker/recommender/Dockerfile .

#dockertravisbuild: build
#	docker build -t $(NAME):$(TAG) -f docker/recommender/Dockerfile-release docker/recommender/
#	docker build -t $(DBNAME):$(TAG) -f docker/recommender-db/Dockerfile docker/recommender-db/
#	docker login -u $(DOCKER_USER) -p $(DOCKER_PASS)
#	scripts/push.sh
