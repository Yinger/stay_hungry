using System;

namespace MathematicalTraining.Chapter_1_基础思想
{
    /// <summary>
    /// 二进制和十进制互相转换
    /// 二进制的位操作
    /// </summary>
    public static class Lesson_01_二进制
    {
        /// <summary>
        /// decimal to binary
        /// </summary>
        /// <param name="source"></param>
        /// <returns></returns>
        public static string IntToBinary(int source)
        {
            string binary = Convert.ToString(source, 2);

            return binary;
        }

        /// <summary>
        /// binary to decimal
        /// </summary>
        /// <param name="source"></param>
        /// <returns></returns>
        public static string BinaryToInt(string source)
        {
            string integer = Convert.ToInt32(source, 2).ToString();

            return integer;
        }

        public static int LeftShift(int num,int m)
        {
            return num << m;
        }

        public static int RightShift(int num,int m)
        {
            return num >> m;
        }

        public static int Or(int num1, int num2)
        {
            return num1 | num2;
        }

        public static int And(int num1, int num2)
        {
            return num1 & num2;
        }

        public static int Xor(int num1, int num2)
        {
            return num1 ^ num2;
        }
    }
}