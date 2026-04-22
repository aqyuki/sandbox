FROM postgres:18.3

COPY ./setup-replica.sh /usr/local/bin/setup-replica.sh
RUN chmod +x /usr/local/bin/setup-replica.sh

USER postgres

ENTRYPOINT [ "setup-replica.sh" ]
CMD [ "postgres" ]
