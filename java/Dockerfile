# syntax=docker/dockerfile:1

# ---- Build stage: compile and package the jar with Maven ----
FROM maven:3.9-eclipse-temurin-21 AS build
WORKDIR /app
# Cache dependencies first
COPY pom.xml ./
RUN mvn dependency:go-offline -B
# Build the application
COPY src ./src
RUN mvn clean package -DskipTests -B

# ---- Runtime stage: small JRE image, non-root ----
FROM eclipse-temurin:21-jre
WORKDIR /app
COPY --from=build /app/target/*.jar app.jar
RUN useradd --create-home appuser
USER appuser
EXPOSE 8080
ENTRYPOINT ["java", "-jar", "app.jar"]
