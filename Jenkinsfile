pipeline {
    agent { label 'agent3' }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/lauprieto/Go.git'
            }
        }

        stage('Build') {
            steps {
                sh 'go mod tidy' 
                sh 'go build -o app ./cmd/app'
            }
        }
    }
}
