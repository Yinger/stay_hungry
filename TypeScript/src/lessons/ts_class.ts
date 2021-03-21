abstract class Animal {
  eat() {
    console.log("eat");
  }

  abstract sleep(): void;
}
class Dog extends Animal {
  constructor(name: string) {
    super();
    this.name = name;
  }
  name: string;
  run() {}
  static food: string = "bones";
  sleep() {
    console.log("dog sleep");
  }
}

console.log(Dog.prototype);
console.log(Dog.food);

let dog = new Dog("WangWang");
console.log(dog);
dog.eat();
// console.log(dog.food)

class Husky extends Dog {
  constructor(name: string, color: string) {
    super(name);
    this.color = color;
  }

  color: string;
}

console.log(Husky.food);

class Cat extends Animal {
  constructor() {
    super();
  }
  sleep() {
    console.log("cat sleep");
  }
}

let cat = new Cat();

let animals: Animal[] = [dog, cat];
animals.forEach((i) => {
  i.sleep();
});

class WorkFlow {
  step1() {
    return this;
  }
  step2() {
    return this;
  }
}

new WorkFlow().step1().step2();

class MyWorkFlow extends WorkFlow {
  next() {
    return this;
  }
}

new MyWorkFlow().step1().next().step2().next();
