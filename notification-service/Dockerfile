FROM openjdk:25-jdk-alpine
WORKDIR /app

# Install the application dependencies
COPY target/mailing-service.jar /app/mailing-service.jar

EXPOSE 8200

CMD ["java","-jar","/app/mailing-service.jar"]