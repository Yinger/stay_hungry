// 交叉类型与联合类型
interface DogInterface {
    run(): void
}

interface CatInterface {
    jump(): void
}

let pat: DogInterface & CatInterface

class MyDog implements DogInterface {
    run() { }
    eat() { }
}

class MyCat implements CatInterface {
    jump() { }
    eat() { }
}

let a: number | string = 'a'
let b: 'a' | 'b' | 'c'
let c: 1 | 2 | 3

enum Master { Boy, Girl }
function getPat(master: Master) {
    let pat = master == Master.Boy ? new MyDog() : new MyCat()
    pat.eat()
    return pat
}

interface Square {
    kind: "square"
    size: number
}

interface Rectangle {
    kind: "rectangle"
    width: number
    height: number
}

interface Circle {
    kind: "circle"
    r: number
}

type Shape = Square | Rectangle | Circle

function area(s: Shape) {
    switch (s.kind) {
        case "square":
            return s.size * s.size
        case "rectangle":
            return s.width * s.height
        case "circle":
            return Math.PI * s.r
        default:
            return ((e: never) => { throw new Error(e) })(s)
    }
}

console.log(area({ kind: 'circle', r: 1 }))