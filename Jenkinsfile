library 'LEAD'
pipeline {
  agent {
    label "lead-toolchain-skaffold"
  }
  stages {
    /// [build]
    stage('Build') {
      steps {
        notifyPipelineStart()
        notifyStageStart()
        container('skaffold') {
          sh "skaffold build --quiet > image.json"
        }
      }
      post {
        success {
          notifyStageEnd()
        }
        failure {
          notifyStageEnd([result: "build failed"])
        }
      }
    }
    /// [build]

    /// [stage]
    stage ('Deploy to Staging') {
      when {
          branch 'master'
      }
      environment {
        TILLER_NAMESPACE = "${env.stagingNamespace}"
        ISTIO_DOMAIN   = "${env.stagingDomain}"
      }
      steps {
        notifyStageStart()
        container('skaffold') {
          sh "skaffold deploy -a image.json -n ${TILLER_NAMESPACE}"
        }
      }
      post {
        success {
          notifyStageEnd([status: "Successfully deployed to Staging: ${env.stagingDomain}"])
        }
        failure {
          notifyStageEnd([result: "Failed to deploy to Staging"])
        }
      }
    }
    /// [stage]

    /// [gate]
    stage ('Manual Ready Check') {
      when {
        branch 'master'
      }
      input {
        message 'Deploy to Production?'
      }
      steps {
        echo "Deploying"
      }
    }
    /// [gate]

    /// [prod]
    stage ('Deploy to Production') {
      when {
          branch 'master'
      }
      environment {
        TILLER_NAMESPACE = "${env.productionNamespace}"
        ISTIO_DOMAIN   = "${env.productionDomain}"
      }
      steps {
        notifyStageStart()
        container('skaffold') {
          sh "skaffold deploy -a image.json -n ${TILLER_NAMESPACE}"
        }
      }
      post {
        success {
          notifyStageEnd([status: "Successfully deployed to Production: ${env.productionNamespace}"])
        }
        failure {
          notifyStageEnd([result: "Failed to deploy to Production"])
        }
      }
    }
    /// [prod]
  }
}
