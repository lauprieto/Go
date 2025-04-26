pipeline {
    agent { label 'agent3' }

    stages {
        stage('Descargar código') {
            steps {
                git branch: 'master', url: 'https://github.com/lauprieto/Go.git'
            }
        }
        
        stage('Build') {
            steps {
                sh 'go mod tidy' // Asegura dependencias
                sh 'go build -o app main.go'
            }
        }
    }
}
