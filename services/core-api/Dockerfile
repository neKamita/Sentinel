# Use a multi-stage build to create a lean final image.

# Stage 1: Build the application
FROM maven:3.9-eclipse-temurin-21-alpine AS builder

WORKDIR /app

# Copy the pom.xml and download dependencies
COPY pom.xml .
RUN mvn dependency:go-offline

# Copy the rest of the source code
COPY src ./src

# Package the application
RUN mvn package -DskipTests

# Stage 2: Create the final image
FROM eclipse-temurin:21-jre-alpine

WORKDIR /app

# Copy the executable JAR from the builder stage
COPY --from=builder /app/target/*.jar /app/app.jar

# Expose the port the application runs on
EXPOSE 8081

# Run the JAR file
ENTRYPOINT ["java", "-jar", "/app/app.jar"]
