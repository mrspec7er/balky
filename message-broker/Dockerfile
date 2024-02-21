FROM rabbitmq:3-management

# Set the hostname
ENV HOSTNAME my-rabbit

# Expose ports
EXPOSE 5672 15672

# Define a volume for RabbitMQ data
VOLUME /var/lib/rabbitmq

# Run RabbitMQ in the background
CMD ["rabbitmq-server"]