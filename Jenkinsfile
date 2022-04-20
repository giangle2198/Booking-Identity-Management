pipeline {
    agent { docker { image 'golang:1.17.5-alpine' } }
    stages {
        stage('Build') {
            steps {
                sh 'echo "Hello World"'
                sh '''
                    echo "Multiple shell steps works too"
                    ls -lah
                '''
            }
        }
    }


}