FROM alpine:3.2
ARG APP=${APP}
ADD ./data /app/data
ADD ./build/${APP} /app
WORKDIR /app
RUN echo -e "#! /bin/sh\n echo cmd $* \n/app/$APP $*" > start.sh
RUN chmod +x /app/start.sh
ENTRYPOINT [ "/app/start.sh" ]