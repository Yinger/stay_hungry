using System;
namespace MathematicalTraining.Chapter_1_基础思想
{
    public static class Lesson_04_数学归纳法
    {
        /// <summary>
        /// 使用函数的递归（嵌套）调用，进行数学归纳法证明
        /// </summary>
        /// <param name="k">放到第几格</param>
        /// <param name="result">保存当前格子的麦粒数和麦粒总数</param>
        /// <returns>放到第k格时是否成立</returns>
        public static bool Prove(int k, Result result)
        {
            // 证明n = 1时，命题是否成立
            if (k == 1)
            {
                if ((Math.Pow(2, 1) - 1) == 1)
                {
                    result.WheatNum = 1;
                    result.WheatTotalNum = 1;
                    return true;
                }
                else return false;
            }
            // 如果n = (k-1)时命题成立，证明n = k时命题是否成立
            else
            {
                bool proveOfPreviousOne = Prove(k - 1, result);
                result.WheatNum *= 2;
                result.WheatTotalNum += result.WheatNum;
                bool proveOfCurrentOne = false;
                if (result.WheatTotalNum == (Math.Pow(2, k) - 1))
                    proveOfCurrentOne = true;
                if (proveOfPreviousOne && proveOfCurrentOne)
                    return true;
                else return false;
            }
        }
    }

    public class Result
    {
        public long WheatNum = 0; // 当前格的麦粒数
        public long WheatTotalNum = 0; // 目前为止麦粒的总数
    }
}
