using System;
namespace MathematicalTraining.Chapter_1_基础思想
{
    /// <summary>
    /// Iterative Method
    /// (二分法中的迭代式逼近)
    /// </summary>
    public static class Lesson_03_迭代法
    {
        /// <summary>
        /// 算算舍罕王给了多少粒麦子
        ///
        /// 在棋盘的第一个小格内放入一粒麦子，
        /// 在第二个小格内放入两粒，
        /// 第三小格内放入给四粒，
        /// 以此类推，每一小格内都比前一小格加一倍的麦子，
        /// 直至放满 64 个格子
        /// 
        /// </summary>
        /// <param name="grid">放到第几格</param>
        /// <returns>麦粒的总数</returns>
        public static long GetNumberOfWheat(int grid)
        {
            long sum = 0; // 麦粒总数
            long numberOfWheatInGrid = 0; // 当前格子里麦粒的数量

            numberOfWheatInGrid = 1; // 第一个格子里麦粒的数量
            sum += numberOfWheatInGrid;

            for (int i = 2; i <= grid; i++)
            {
                numberOfWheatInGrid *= 2;   // 当前格子里麦粒的数量是前一格的2倍
                sum += numberOfWheatInGrid;   // 累计麦粒总数
            }

            return sum;
        }

        /// <summary>
        /// 计算大于1的正整数之平方根
        /// </summary>
        /// <param name="n">待求的数</param>
        /// <param name="deltaThreshold">误差的阈值</param>
        /// <param name="maxTry">二分查找的最大次数</param>
        /// <returns>平方根的解</returns>
        public static double GetSqureRoot(int n, double deltaThreshold, int maxTry)
        {
            if (n <= 1)
                return -1.0;

            double min = 1.0, max = (double)n;

            for (int i = 0; i < maxTry; i++)
            {
                double middle = (min + max) / 2;
                double square = middle * middle;
                double delta = Math.Abs((square / n) - 1);//如果两数完全一样那么他们相除结果为1，两数相除减一取绝对值，表示两数误差大小

                if (delta <= deltaThreshold)
                    return middle;
                else
                {
                    if (square > n)
                        max = middle;
                    else
                        min = middle;
                }
            }
            return -2.0;
        }

        /// <summary>
        /// 查找某个单词是否在字典里出现
        /// </summary>
        /// <param name="dictionary">排序后的字典</param>
        /// <param name="wordToFind">待查的单词</param>
        /// <returns>是否发现待查的单词</returns>
        public static bool SearchByIterative(string[] dictionary, string wordToFind)
        {
            if (dictionary == null)
                return false;

            if (dictionary.Length == 0)
                return false;

            int left = 0, right = dictionary.Length - 1;

            while (left <= right)
            {
                //int middle = (left + right) / 2;      //如果left和right都是接近系统设定的最大值，那么两者相加会溢出。
                int middle = left + (right - left) / 2; //如果只加两者差值的一半，那么不会超过两者中较大的值，自然也不会溢出
                if (dictionary[middle].Equals(wordToFind))
                    return true;
                else
                {
                    if (dictionary[middle].CompareTo(wordToFind) > 0)
                        right = middle - 1;
                    else
                        left = middle + 1;
                }
            }
            return false;
        }
    }
}
