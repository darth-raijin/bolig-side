pipeline {
    agent none

    environment {
        DOCKER_USERNAME = 
        DOCKER_PASSWORD = 
    }

    stages {
        stage("setup") {
            env.GIT_REVISION = sh(returnStdout: true, script: "git rev-parse HEAD").trim()
            env.BRANCH_TYPE = sh(returnStdout: true, script: "git rev-parse --abbrev-ref HEAD").trim()
        }

        stage("deploy image to docker hub") {
            sh(returnStdout: true, script: "docker push h0tsauce/loop-api:${GIT_REVISION} .").trim()
        }
    }
}