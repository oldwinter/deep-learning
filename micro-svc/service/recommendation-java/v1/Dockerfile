FROM openjdk:8-jdk-slim
COPY . /app
WORKDIR /app
RUN javac Main.java
EXPOSE 3005
CMD ["java", "Main"]
