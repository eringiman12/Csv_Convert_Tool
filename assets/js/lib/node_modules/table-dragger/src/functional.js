// 容器===范畴
const a = {
  value: 1,
  addOne(x) {
    return x + 5
  }
}

// 本质上，函数式编程只是范畴论的运算方法，跟数理逻辑、微积分、行列式是同一类东西，都是数学方法，只是碰巧它能用来写程序。
// 任何具有map方法的数据结构，都可以当作函子的实现。

// 一般约定，函子的标志就是容器具有map方法。该方法将容器里面的每一个值，映射到另一个容器。
class Functor {
  constructor(val) {
    this.val = val;
  }

  map(f) {
    return new Functor(f(this.val));
  }
}

Functor.of = function(val) {
  return new Functor(val);
};

Functor.of(2).map(function (two) {
  return two + 2;
});
