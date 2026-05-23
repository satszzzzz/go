pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Docker build') {
            steps {
                sh 'DOCKER_CONFIG=/tmp/empty-docker-config docker build -t go-app:latest .'
            }
        }

        stage('Docker run') {
            steps {
                sh 'docker rm -f go-app || true'
                sh 'docker run -d --name go-app go-app:latest'
            }
        }
    }

    post {
        failure {
            echo 'Pipeline failed!'
        }
        success {
            echo 'Pipeline completed!'
            sh 'docker logs go-app'
            sh 'docker rm -f go-app || true'
            sh 'docker rmi go-app:latest || true'
        }
    }
}