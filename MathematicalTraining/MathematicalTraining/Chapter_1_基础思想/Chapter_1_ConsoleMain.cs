using System;

namespace MathematicalTraining.Chapter_1_基础思想
{
    public static class Chapter_1_ConsoleMain
    {
        public static void Lesson01_二进制()
        {
            int a = 53;
            string strb = "110101";
            Console.WriteLine("{0} →  binary = {1}", a, Lesson_01_二进制.IntToBinary(a));
            Console.WriteLine("{0} → decimal = {1}", strb, Lesson_01_二进制.BinaryToInt(strb));

            int num = 53;
            int m = 1;
            Console.WriteLine("{0} left  shift {1} = {2}", num, m, Lesson_01_二进制.LeftShift(num, m));
            Console.WriteLine("{0} right shift {1} = {2}", num, m, Lesson_01_二进制.RightShift(num, m));

            a = 53;
            int b = 35;
            Console.WriteLine("{0} | {1} = {2}", a, b, Lesson_01_二进制.Or(a, b));
            Console.WriteLine("{0} & {1} = {2}", a, b, Lesson_01_二进制.And(a, b));
            Console.WriteLine("{0} ^ {1} = {2}", a, a, Lesson_01_二进制.Xor(a, a));

            int grid = 63;
            Console.WriteLine("到第 {0} 格时，总共需要 {1} 麦粒", grid, Lesson_03_迭代法.GetNumberOfWheat(grid));

            num = 10;
            double squareRoot = Lesson_03_迭代法.GetSqureRoot(num, 0.000001, 10000);
            if (squareRoot == -1.0)
                Console.WriteLine("请输入大于1的整数");
            else if (squareRoot == -2.0)
                Console.WriteLine("未能找到解");
            else
                Console.WriteLine("{0}的平方根是{1}", num, squareRoot);

            string[] dictionary = { "talk", "is", "cheap", "show", "me", "the", "code" };
            Array.Sort(dictionary);
            string wordToFind = "me";
            bool found = Lesson_03_迭代法.SearchByIterative(dictionary, wordToFind);
            if (found)
                Console.WriteLine("找到了单词 {0}", wordToFind);
            else
                Console.WriteLine("未能找到单词 {0}", wordToFind);

            Result result = new Result();
            Console.WriteLine(Lesson_04_数学归纳法.Prove(grid, result));
        }
    }
}