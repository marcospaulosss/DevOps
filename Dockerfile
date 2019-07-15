# For deployment purpose only
FROM gustavohenrique/golang:1.12-stretch

ARG app_name

COPY . /backend

RUN cd /backend/apps/$app_name \
 && make compile \
 && mv main /usr/local/bin

CMD ["/usr/local/bin/main"]

