# Stage 1: Build the application
FROM maven:3.9.6-eclipse-temurin-17 AS build

WORKDIR /app

# Copy pom.xml and download dependencies (cached unless pom.xml changes)
COPY pom.xml .
RUN mvn dependency:go-offline

# Copy source code
COPY src ./src

# Package the app
RUN mvn clean package -DskipTests

# Stage 2: Run the app
FROM openjdk:17-jdk-slim

WORKDIR /app

# Copy built JAR from the build stage
COPY --from=build /app/target/*.jar app.jar

# Expose the app port
EXPOSE 3002

# Start the app
ENTRYPOINT ["java", "-jar", "app.jar"]
