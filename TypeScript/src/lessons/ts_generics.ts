//汎型

//改造log函數
function log<T>(value: T): T {
  console.log(value);
  return value;
}

//調用
//1 指明T的類型
log<string[]>(["a", "b"]);
//2 利用TS的類型推斷省略類型參數
log(["a", "b"]);

//使用汎型來定義一個函數類型
type Log = <T>(value: T) => T;
let myLog: Log = log;

//使用汎型來定義一個interface
interface Logg<T = string> {
  (value: T): T;
}
let myLogg: Logg = log;
myLogg("1");

//汎型Class
class Logc<T> {
  run(value: T) {
    console.log(value);
    return value;
  }
}

let log1 = new Logc<number>();
log1.run(2);

let log2 = new Logc();
log2.run("3");

interface Length {
  length: number;
}

function logLength<T extends Length>(value: T): T {
  console.log(value, value.length);
  return value;
}

logLength([123]);
log("12345");
log({ length: 1 });

// 泛型的好处
// 1.函数和类可以支持多种类型，增加的程序的可扩展性
// 2.不必写多条函数重载，联合类型声明，增强代码的可读性
// 3.灵活控制类型之间的约束
