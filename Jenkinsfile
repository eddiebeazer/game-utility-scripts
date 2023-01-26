pipeline {
    agent { label 'Windows' }
    tools {
        go 'go-1.16.6'
    }
    environment {
        UNREAL_ENGINE_STORAGE_PASSWORD = credentials('unreal_storage_password')
    }
    stages {
        stage('Building') {
            steps {
                bat "powershell.exe -file .\\build.ps1"
            }
        }
        stage('Deploying') {
            when {
                expression {
                    env.BRANCH_NAME == 'main'
                }
            }
            steps {
                echo 'Mounting Tools Share'
                bat 'net use * /delete /y && net use X: \\\\unrealenginestorage.file.core.windows.net\\tools /user:localhost\\unrealenginestorage %UNREAL_ENGINE_STORAGE_PASSWORD%'

                echo 'Copying Executables to Tools Share'
                bat "rmdir X:\\bin\\PlayfabSDK\\ /S /Q && md X:\\bin\\PlayfabSDK\\"
                bat "xcopy $WORKSPACE\\bin X:\\bin\\PlayfabSDK\\ /E /H"
            }
        }
    }
}
