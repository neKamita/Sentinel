import pika
import time
import os
import logging

logging.basicConfig(level=logging.INFO, format='%(asctime)s - %(levelname)s - %(message)s')

def connect_to_rabbitmq():
    """Establishes a connection to RabbitMQ, retrying if necessary."""
    rabbitmq_host = os.getenv('RABBITMQ_HOST', 'rabbitmq')
    connection_params = pika.ConnectionParameters(host=rabbitmq_host, credentials=pika.PlainCredentials('sentinel', 'supersecretpassword'))
    
    while True:
        try:
            connection = pika.BlockingConnection(connection_params)
            logging.info("Successfully connected to RabbitMQ.")
            return connection
        except pika.exceptions.AMQPConnectionError as e:
            logging.warning(f"Failed to connect to RabbitMQ: {e}. Retrying in 5 seconds...")
            time.sleep(5)

def main():
    """Main function to set up RabbitMQ consumer."""
    connection = connect_to_rabbitmq()
    channel = connection.channel()

    # Declare the queues
    channel.queue_declare(queue='events.raw', durable=True)
    channel.queue_declare(queue='incidents.detected', durable=True)

    def callback(ch, method, properties, body):
        logging.info(f"Received event: {body.decode()}")
        
        # Placeholder for analysis logic
        # In a real scenario, you would parse the event, apply rules, and run ML models.
        is_anomalous = "error" in body.decode().lower() # Simple rule for demonstration

        if is_anomalous:
            incident_message = f'"incident_type": "Potential Error Detected", "original_event": {body.decode()}'
            logging.warning(f"Anomaly detected. Publishing to incidents.detected: {incident_message}")
            ch.basic_publish(
                exchange='',
                routing_key='incidents.detected',
                body=incident_message,
                properties=pika.BasicProperties(delivery_mode=2) # make message persistent
            )
        
        ch.basic_ack(delivery_tag=method.delivery_tag)

    channel.basic_qos(prefetch_count=1)
    channel.basic_consume(queue='events.raw', on_message_callback=callback)

    logging.info('Waiting for events. To exit press CTRL+C')
    try:
        channel.start_consuming()
    except KeyboardInterrupt:
        logging.info("Shutting down...")
        channel.stop_consuming()
    finally:
        connection.close()
        logging.info("RabbitMQ connection closed.")

if __name__ == '__main__':
    main()
