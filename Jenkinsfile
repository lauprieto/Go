pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/lauprieto/Go.git'
            }
        }

        stage('Build') {
            steps {
                sh 'go build ./...'
            }
        }

        stage('Run Tests') {
            steps {
                sh 'go test ./...'
            }
        }
    }
}
