<<<<<<< HEAD
=======
library 'LEAD'
>>>>>>> c54e117a9cfe155fe12497f758b7a46f80d915ee
pipeline {
  agent {
    label "lead-toolchain-skaffold"
  }
  stages {
<<<<<<< HEAD
    ### [build]
    stage('Build') {
      steps {
=======
    /// [build]
    stage('Build') {
      steps {
        notifyPipelineStart()
        notifyStageStart()
>>>>>>> c54e117a9cfe155fe12497f758b7a46f80d915ee
        container('skaffold') {
          sh "skaffold build --quiet > image.json"
        }
      }
<<<<<<< HEAD
    }
    ### [build]
    
    ### [stage]
=======
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
>>>>>>> c54e117a9cfe155fe12497f758b7a46f80d915ee
    stage ('Deploy to Staging') {
      environment {
        TILLER_NAMESPACE = "${env.stagingNamespace}"
        DOMAIN   = "${env.stagingDomain}"
      }
      steps {
<<<<<<< HEAD
=======
        notifyStageStart()
>>>>>>> c54e117a9cfe155fe12497f758b7a46f80d915ee
        container('skaffold') {
          sh "skaffold deploy -a image.json -n ${TILLER_NAMESPACE}"
        }
      }
<<<<<<< HEAD
    }
    ### [stage]
    
    ### [gate]
=======
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
>>>>>>> c54e117a9cfe155fe12497f758b7a46f80d915ee
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
<<<<<<< HEAD
    ### [gate]

    ### [prod]
=======
    /// [gate]

    /// [prod]
>>>>>>> c54e117a9cfe155fe12497f758b7a46f80d915ee
    stage ('Deploy to Production') {
      when {
          branch 'master'
      }
      environment {
        TILLER_NAMESPACE = "${env.productionNamespace}"
        DOMAIN   = "${env.productionDomain}"
      }
      steps {
<<<<<<< HEAD
=======
        notifyStageStart()
>>>>>>> c54e117a9cfe155fe12497f758b7a46f80d915ee
        container('skaffold') {
          sh "skaffold deploy -a image.json -n ${TILLER_NAMESPACE}"
        }
      }
<<<<<<< HEAD
    }
    ### [prod]
=======
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
>>>>>>> c54e117a9cfe155fe12497f758b7a46f80d915ee
  }
}
