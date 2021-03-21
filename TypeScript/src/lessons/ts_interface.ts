interface List {
  readonly id: number; //只读属性
  name: string;
  age?: number; //可选属性
}

interface Result {
  data: List[];
}

function render(result: Result) {
  result.data.forEach((value) => {
    console.log(value.id, value.name);

    if (value.age) {
      console.log(value.age);
    }
  });
}

let result = {
  data: [
    { id: 1, name: "A", sex: "male" },
    { id: 2, name: "B", age: 10 },
  ],
};

render(result);

//当不确定接口中属性个数时需要使用 索引签名
//索引签名包括字符串索引签名和数字索引签名
interface StringArray {
  [index: number]: string;
}

let chars: StringArray = ["A", "B"];

//当接口中定义了一个索引后，例如设置了 【x:string】= string，就不能设置y：number了
//可以同时使用两种类型的索引，但是数字索引的返回值必须是字符串索引返回值类型的子类型。
interface Names {
  [x: string]: string;
  // y:number
  [z: number]: string;
}

let names: Names = {
  name1: "a",
  name2: "b",
};

// let add:(x:number, y:number) => number

//interface : 创建新的类型，接口之间还可以继承、声明合并。
// interface Add{
//     (x: number, y: number) :number
// }

//type: 不是创建新的类型，只是为一个给定的类型起一个名字。type还可以进行联合、交叉等操作，引用起来更简洁。
type Add = (x: number, y: number) => number;

let add: Add = (a, b) => a + b;

interface Lib {
  (): void;
  version: string;
  doSomething(): void;
}

function getLib() {
  let lib: Lib = (() => {}) as Lib;
  lib.version = "1.0";
  lib.doSomething = () => {};
  return lib;
}

let lib1 = getLib();
lib1();
lib1.doSomething();
let lib2 = getLib();
