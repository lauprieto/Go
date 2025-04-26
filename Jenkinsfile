pipeline {
    agent { label 'agent3' }

    environment {
        GO111MODULE = 'on'
    }

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

        stage('Test') {
            steps {
                script {
                    def testResult = sh(script: 'go test -v ./...', returnStatus: true)
                    if (testResult != 0) {
                        error "Tests fallaron"
                    }
                }
            }
        }
    }
}


