FROM mongo:7.0

COPY conf/mongod.conf /etc/mongo/mongod.conf
COPY docker-entrypoint-initdb/ /docker-entrypoint-initdb.d/

CMD ["mongod", "--config", "/etc/mongo/mongod.conf"]
