pipeline {
    agent any

    stages {
        stage('Hello') {
            steps {
                echo 'Hello World, from Jenkins! :D'
                sh 'echo running on $(hostname)'
            }
        }
    }
}