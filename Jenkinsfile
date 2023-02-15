pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                sh 'go build -o /bin/my-app main.go'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}