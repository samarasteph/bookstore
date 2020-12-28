FROM mariadb/server AS db-stage

LABEL name="database-server"
LABEL Production="true"

ARG RW_USER
ARG RW_PWD
ARG RO_USER
ARG RO_PWD
ARG SRV_NAME=dbserver
ARG DB_PWD=aldu

#ENV MARIADB_DATABASE="booksdb"
ENV MARIADB_ROOT_PASSWORD=$DB_PWD

RUN addgroup dbusers
#create read write user: no home dir, dbusers group
RUN useradd -M -g dbusers $RW_USER
RUN echo $RW_USER:$RW_PWD |chpasswd
#Same with read only user
RUN useradd -M -g dbusers $RO_USER
RUN echo $RO_USER:$RO_PWD |chpasswd

WORKDIR "/app-database"
RUN chown mysql "/app-database"

COPY mariadb /etc/pam.d/
COPY dbconfig.sh /docker-entrypoint-initdb.d/
COPY books_data.sql /docker-entrypoint-initdb.d/
COPY setup-users.sql /docker-entrypoint-initdb.d/

#RUN echo [mariadb] >> /etc/mysql/mariadb.conf.d/docker.cnf
#RUN echo enable-cleartext-plugin >> /etc/mysql/mariadb.conf.d/docker.cnf
#RUN echo default_authentication_plugin=sha256_password
#RUN ln -s /usr/lib/x86_64-linux-gnu/libmariadb3/plugin/sha256_password.so /usr/lib/mysql/plugin/sha256_password.so
