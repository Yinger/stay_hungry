# TypeScriptWorkplace
1. 下載並安裝nodejs https://nodejs.org/
2. 配置VSCode
3. 新建文件夾
    * npm init -y
    * npm i typescript -g
    * tsc --init
    * tsc ./src/index.ts
4. 配置工具
    * npm i webpack webpack-cli webpack-dev-server -D
    * npm i ts-loader typescript -D
    * npm i html-webpack-plugin -D
    * npm i clean-webpack-plugin -D
    * npm i webpack-merge -D
5. create build folder
    * webpack.config.js
    * webpack.base.config.js
    * webpack.dev.config.js
    * webpack.pro.config.js
6. change package.json
7. npm start / npm run build

## PS : VSCode 配置 Github
1. git init
    ```
    git init
    git config --global user.name "yourname"
    git config --global user.email "youremail"
    ``` 
2. create github repository (ProjectDemo)
3. create ssh key
    ```
    ssh-keygen -t rsa -C "your email"
    ```
4. set github deploy keys
    * copy id_rsa.pub to deploy key
5. binding local folder to github repository
    ```
    git remote add ProjectDemo git@github.com:yourname/ProjectDemo.git
    ```
6. test ssh
    ```
    ssh -T git@github.com
    ```
7. pull
    ```
    git pull ProjectDemo master
    ```
8. push
    ```
    git push --force ProjectDemo master
    ```
