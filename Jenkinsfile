pipeline {
    agent any

    stages {
        stage('pull code') {
            steps {
                git credentialsId: '8baf1406-130c-4a7e-a099-babbbfcdd00a', url: 'git@github.com:JW-Zhang001/Test-Actions.git'
            }
        }
    }
}
