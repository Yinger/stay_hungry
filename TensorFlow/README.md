# TensorFlowWorkplace
## 环境准备 - Python3
* Mac自带 Python2.7, 自安装 Python3.7 (TensorFlow2.2 要求 Python3.5-3.7) 安装 virtualenv 后  “command not found”
  * 原因：Python3.7没有设定为默认 (设定默认版本：)
  > alias python=/usr/local/bin/python3.7
  ``` 
  which python3.7
  python -V
  pip3 list
  pip3 uninstall virtualenv
  pip3 install virtualenv
  virtualenv --version
  ```
## 环境准备 - TensorFlow 2.2
* Mac 下 为 TensorFlow 创建虚拟环境
  ```
  # 创建一个新的虚拟环境，方法是选择 Python 解释器并创建一个 ./venv 目录来存放它
  virtualenv --system-site-packages -p python3.7 ./venv

  # 使用特定于 shell 的命令激活该虚拟环境
  source ./venv/bin/activate  # sh, bash, ksh, or zsh
  #当 virtualenv 处于有效状态时，shell 提示符带有 (venv) 前缀。

  # 在不影响主机系统设置的情况下，在虚拟环境中安装软件包。首先升级 pip：
  (venv)$ pip install --upgrade pip
  (venv)$ pip list  # show packages installed within the virtual environment

  # 之后可以使用以下命令退出 virtualenv：
  (venv)$ deactivate  # don't exit until you're done using TensorFlow
  ```
* 安装 tensorflow2.2
    ```
    $ source ./venv/bin/activate
    (venv)$ pip install tensorflow==2.2.0

    # 验证
    (venv)$ python -c "import tensorflow as tf;print(tf.__version__)"
    ```
## 环境准备 - jupyter lab
  ```
    alias python=/usr/local/bin/python3.7
    source ./venv/bin/activate
    (venv)$ pip install --upgrade pip
    (venv)$ pip install jupyter
    (venv)$ python -m ipykernel install --user --name=venv
    (venv)$ jupyter lab 
  ```
## 环境准备 - VS Code
  * settings.json
  ```
  {
    "python.venvPath": "C:\\Work\\1_Study\\venv",
    "python.pythonPath": "C:\\Work\\1_Study\\venv\\Scripts\\python.exe"
  }
  ```

  ### Jupyter 快捷键
  #### 两种模式
* Command mode 和 Edit mode
  * __Edit:__ 在一个cell中 按下 __Enter__ 进入Edit模式
  * __Command:__ 在一个cell中 按下 __Esc__ 进入Command 模式
* __以下命令全部在 command 模式下__
#### 运行当前cell，并移动到下一个Cell
*  __Shift + Enter__
#### 创建Cell
* 按下 __a__ ,即可在这个cell之 __前__ 创建一个新的cell
* 按下 __b__ ,即可在这个cell之 __后__ 创建一个新的cell
#### Cell中 Code 和 Markdown的切换
* 按下 __y__, 进入 __Code__ 
* 按下 __m__, 进入 __Markdown__ 
#### 显示Cell中的行数
* 按下 __l__
#### 删除Cell
* 在一个cell中, 按 __两__ 次 __d__
#### 保存Notebook
* 在一个cell中, 按下 __s__

#### 启动命令面板
* 在一个cell中, 按下 __Ctrl + Shift + P__

#### 其他
* __1__ 将单元格设定一级标题
* __2__ 将单元格设定二级标题
* __3__ 将单元格设定三级标题
* __4__ 将单元格设定四级标题
* __5__ 将单元格设定五级标题
* __6__ 将单元格设定六级标题


## 常用的内置评估指标
* MeanSquaredError（平方差误差，用于回归，可以简写为MSE，函数形式为mse）
* MeanAbsoluteError (绝对值误差，用于回归，可以简写为MAE，函数形式为mae)
* MeanAbsolutePercentageError (平均百分比误差，用于回归，可以简写为MAPE，函数形式为mape)
* RootMeanSquaredError (均方根误差，用于回归)
* Accuracy (准确率，用于分类，可以用字符串"Accuracy"表示，Accuracy=(TP+TN)/(TP+TN+FP+FN)，要求y_true和y_pred都为类别序号编码)
* Precision (精确率，用于二分类，Precision = TP/(TP+FP))
* Recall (召回率，用于二分类，Recall = TP/(TP+FN))
* TruePositives (真正例，用于二分类)
* TrueNegatives (真负例，用于二分类)
* FalsePositives (假正例，用于二分类)
* FalseNegatives (假负例，用于二分类)
* AUC(ROC曲线(TPR vs FPR)下的面积，用于二分类，直观解释为随机抽取一个正样本和一个负样本，正样本的预测值大于负样本的概率)
* CategoricalAccuracy（分类准确率，与Accuracy含义相同，要求y_true(label)为onehot编码形式）
* SparseCategoricalAccuracy (稀疏分类准确率，与Accuracy含义相同，要求y_true(label)为序号编码形式)
* MeanIoU (Intersection-Over-Union，常用于图像分割)
* TopKCategoricalAccuracy (多分类TopK准确率，要求y_true(label)为onehot编码形式)
* SparseTopKCategoricalAccuracy (稀疏多分类TopK准确率，要求y_true(label)为序号编码形式)
* Mean (平均值)
* Sum (求和)
