# 设计模式

关于设计模式的 Golang 实现

## 面向对象开发的原则

### 单一职责(Single responsability)
> 职责单一并且只有一个改变的理由

### 开闭原则(Open-Close principle)
> 对扩展开放，对修改关闭

### 里氏替换原则(Liskov substitution principle)
> 一个对象或者接口在其出现的任何地方，都可以用子类实例做替换，并且不会导致程序的错误

多态: 父类(接口)对象的指针指向子类的实例

### 依赖倒置(Dependency inversion)
> 高级模块不应该依赖于低级模块。两者都应该依赖于抽象。抽象不应该依赖于细节。细节应依赖于抽象

底层模块像上层暴露的方法应该使用接口。上层模块只依赖接口。底层模块只需要保证对外提供的接口不变。

### 接口隔离(Interface segregation)
> 客户端不应该被要求实现一些不需要的方法

接口的设计应该尽量简单，避免实现接口时还需要实现一些不相关的方法

### 合成复用原则
> 组合优于继承

使用组合可以有效的减少子类对夫类的依赖。继承会继承父类所有的方法，  甚至是父类的父类的方法


## 创建型模式

### 单例模式
### 简单工厂模式
### 工厂方法模式
### 抽象工厂模式

## 结构型模式

### 适配器模式
### 代理模式
### 装饰器模式
### 外观模式

## 行为型模式

### 命令模式
### 模板方法模式
### 观察者模式