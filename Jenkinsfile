pipeline {
    agent {
            docker {
                image 'golang:1.15-alpine'
                args '-v /data/my-app-cache:/go/.cache'
            }
        }

        options {
            timeout(time: 20, unit: 'MINUTES')
            disableConcurrentBuilds()
        }

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