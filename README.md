# indexing

# Using SQS Locally with LocalStack

Here's a basic outline of how you might use SQS locally with LocalStack:

1. **Install LocalStack**: First, you need to install LocalStack. You can usually install it via pip:
    
    ```bash
    pip install localstack
    ```
    
2. **Start LocalStack**: Once installed, you can start LocalStack using the following command:
    
    ```bash
    localstack start
    ```
    
    This command starts up the LocalStack environment, including the emulated AWS services.
    
3. **Create SQS Queue**: After LocalStack is running, you can create an SQS queue using the AWS CLI
    
    ```bash
    aws --endpoint-url=http://localhost:4566 sqs create-queue --queue-name my-queue
    ```
    
4. **Send and Receive Messages**: With the queue created, you can send messages to it and receive messages from it using the AWS CLI or SDK.
    
    For example, to send a message:
    
    ```bash
    aws --endpoint-url=http://localhost:4566 sqs send-message --queue-url=http://localhost:4566/000000000000/my-queue --message-body "Hello, world!"
    ```
    
    And to receive a message:
    
    ```bash
    aws --endpoint-url=http://localhost:4566 sqs receive-message --queue-url=http://localhost:4566/000000000000/my-queue
    ```
    
5. **Stop LocalStack**: After you're done using LocalStack, you can stop it using:
    
    ```bash
    localstack stop
    ```

```bash
export AWS_ACCESS_KEY_ID=YOUR_ACCESS_KEY
export AWS_SECRET_ACCESS_KEY=YOUR_SECRET_KEY
export SQS_QUEUE_URL=http://localhost:4566/000000000000/my-queue
export SQS_ENDPOINT=http://localhost:4566
```

```
go run main.go
```


## Sample
![Screenshot 2024-03-26 at 03 46 11](https://github.com/uwezukwechibuzor/indexing/assets/66339097/61936b4d-c873-45be-9abd-fa1030fe4c30)

