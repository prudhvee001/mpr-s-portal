# Stage 1: Build
FROM maven:3.9.6-eclipse-temurin-17 AS build
WORKDIR /app
COPY pom.xml .
RUN mvn dependency:go-offline
COPY src ./src
RUN mvn clean package -DskipTests

# Stage 2: Run
FROM openjdk:17-jdk-slim
WORKDIR /app
COPY --from=build /app/target/results-service-0.0.1-SNAPSHOT.jar app.jar
EXPOSE 3002
ENTRYPOINT ["java", "-jar", "app.jar"]
