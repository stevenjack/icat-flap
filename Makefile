
.PHONY: run
.SILENT: run
.DEFAULT: run

run:
	echo "=> Running app"
	docker run --privileged -v $(pwd):/app -v /dev/gpiomem:/dev/gpiomem --rm=true -it smaj/icatflap
