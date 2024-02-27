FROM rabbitmq:3-management

WORKDIR /message-hub

ENV HOSTNAME my-rabbit

COPY api-definitions.json .

COPY start.sh .

EXPOSE 5672 15672

RUN chmod +x ./start.sh

CMD ["./start.sh"]