pipeline {
    agent any
    environment {
        GO111MODULE = 'on'
    }
    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/lauprieto/Go.git'
            }
        }
        stage('Build') {
            steps {
                sh 'go build -v ./...'
            }
        }
        stage('Test') {
            steps {
                sh 'go test -v ./...'
            }
        }
    }
}

