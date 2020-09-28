
FROM maven:3.6-jdk-8-slim as builder
COPY . /app
WORKDIR /app
RUN mvn install 
FROM openjdk:8-jdk-slim
WORKDIR /app
COPY --from=builder /app/target/RecommendationSystem-1.0-SNAPSHOT.jar .
EXPOSE 3005
CMD ["java", "-jar", "RecommendationSystem-1.0-SNAPSHOT.jar"]
