# Build stage - use glibc-based image so protoc binaries work
FROM maven:3.9-eclipse-temurin-25 AS builder
WORKDIR /build
COPY pom.xml .
RUN mvn dependency:go-offline
COPY src ./src
RUN mvn clean package -DskipTests

# Runtime stage - alpine is fine here, no protoc needed
FROM eclipse-temurin:25-jre-alpine
WORKDIR /app
COPY --from=builder /build/target/mailing-service.jar /app/app.jar
EXPOSE 8200
CMD ["java","-jar","/app/app.jar"]