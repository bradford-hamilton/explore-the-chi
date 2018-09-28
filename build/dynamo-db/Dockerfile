# FROM openjdk:7-jre-alpine
FROM openjdk:7

# Some metadata.
LABEL maintainer="brad.lamson@gmail.com"

# Create our main application folder.
RUN mkdir -p opt/dynamodb
WORKDIR /opt/dynamodb

# Download and unpack dynamodb.
# Links are from: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.html
RUN wget https://s3-us-west-2.amazonaws.com/dynamodb-local/dynamodb_local_latest.tar.gz -q -O - | tar -xz
  
# The entrypoint is the dynamodb jar. Default port is 8000. SharedDb flag so that all connections use
# the same database files (this application and the shell at localhost:8000/shell)
EXPOSE 8000
ENTRYPOINT ["java", "-jar", "DynamoDBLocal.jar", "-sharedDb"]
