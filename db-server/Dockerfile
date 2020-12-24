FROM mariadb/server AS db-stage

LABEL name="database-server"
LABEL Production="true"

ARG DB_PWD=aldu
ARG RW_PWD=rw_aldu
ARG RO_PWD=ro_aldu
ARG SRV_NAME=dbserver

#ENV MARIADB_DATABASE="booksdb"
ENV MARIADB_ROOT_PASSWORD=$DB_PWD

#VOLUME "/var/db"

WORKDIR "/app-database"
RUN chown mysql "/app-database"

#RUN echo "#!/bin/bash" > /docker-entrypoint-initdb.d/dbconfig.sh
#RUN echo "sudo sed -i -r  's/^#bind-address += 127.0.0.1/bind-address = 0.0.0.0/' /etc/mysql/mariadb.conf.d/50-server.cnf" >> /docker-entrypoint-initdb.d/dbconfig.sh
COPY dbconfig.sh /docker-entrypoint-initdb.d/
COPY books_data.sql /docker-entrypoint-initdb.d/

#CMD ["/docker-entrypoint-initdb.d/dbconfig.sh"] 