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


        stage('Deploy') {
            steps {
                retry(3) {
                    sh 'echo "Deploying....."'
                }

                timeout(time: 3, unit: 'MINUTES') {
                    sh 'echo "Timeout Processing....."'
                }
            }
        }


        stage('Test') {
            steps {
                sh 'echo "Fail!"; exit 1'
            }
        }
    }

    post {
        always {
            echo 'This will always run'
        }
        success {
            echo 'This will run only if successful'
        }
        failure {
            echo 'This will run only if failed'
        }
        unstable {
            echo 'This will run only if the run was marked as unstable'
        }
        changed {
            echo 'This will run only if the state of the Pipeline has changed'
            echo 'For example, if the Pipeline was previously failing but is now successful'
        }
    }

}