pipeline {
    agent any
    stages {
        stage('Build') {
            steps {
                sh '''
                go mod tidy
                go build -o bin/my-app main.go
                '''
            }
        }
        stage('Test'){
            steps {
               echo "test"
            }
        }
        stage('Deploy') {
            steps {
                echo "deploy"
            }
        }
    }
}