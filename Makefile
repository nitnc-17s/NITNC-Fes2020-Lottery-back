DEVNAME=lottery-back-dev
NAME=lottery-back
VERSION=1.1

dev-build:
	docker build -f ./build/Dockerfile --target develop -t $(DEVNAME):$(VERSION) .

dev-run:
	docker run --rm -it -p 8080:8080 --name $(DEVNAME) $(DEVNAME):$(VERSION)

dev-stop:
	docker stop $(DEVNAME)

dev-remove:
	docker rm $(DEVNAME)

dev-logs:
	docker logs $(DEVNAME)

main-build:
	docker build -f ./build/Dockerfile --target app -t $(NAME):$(VERSION) .

main-start:
	docker run -itd -p 8080:8080 --name $(NAME) $(NAME):$(VERSION)

main-start!:
	docker run -it --rm -p 8080:8080 --name $(NAME) $(NAME):$(VERSION)

main-up:
	docker build -f ./build/Dockerfile --target app -t $(NAME):$(VERSION) .
	docker run -itd -p 3000:3000 --name $(NAME) $(NAME):$(VERSION)

main-stop:
	docker stop $(NAME)
	docker rm $(NAME)

main-logs:
	docker logs $(NAME)
