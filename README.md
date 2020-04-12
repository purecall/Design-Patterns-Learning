# Golang-Design-Patterns
学习设计模式，并使用Golang等语言实现





## Creational Pattern 创建型模式

1. **Builder （建造者模式，生成器模式）**

   - builder_facets
   - builder_parameter
   - functional_builder

2. **Factories**（抽象工厂好像go的interface{}已经是很自然的设计了，不再赘述）

   - **interface factory 工厂方法模式**

     工厂方法是粒度很小的设计模式，因为模式的表现只是一个**抽象的方法**。

     **在工厂和产品中间增加接口**，工厂不再负责产品的创建，由接口**根据不同条件返回具体的类实例**

     完全遵循"**开放-关闭原则**"，**可扩展**

   - **factory generator**：用于**生成特定角色的工厂**，两种方式：functional approach和structural approach

   -  **prototype_factory**，每个预定义类型都用一个常量表示

3. **Prototype 原型模式（克隆，Clone）**

   - **prototype serialization**
   - **prototype factory**

4. **Singleton 单例模式**

   使用场景

   - It makes sense to **have one in the system**: Database repository，Object factory
   - **The construction call is expensive**
     - We only do it once
     - We give everyone the same instance
   - Want to prevent anyone creating additional copies
   - Need to take care of **lazy instantiation**

   

   具体代码 **singleton.go**：

   其中再次提到了**DIP(Dependence Inversion Principle)**，测试单例数据库实例，比如当时两个数据1和2，不应该ok := GetInstance().GetCityPopulation() == (1+2)，因为1和2其实是magic number，任何时候数据库都可能被修改。这样的测试语句，其实是**在测试数据库的读取**，**而非在测试获取人口的函数GetCityPopulation()**

   改进方法：DIP，抽象出一个接口，创建类DummyDatabase，用这个类来测试GetCityPopulation()函数，**无需依赖高层模块**



## Structural Pattern 结构型模式

1. **Adapter**，**适配器模式**、封装器模式：将一个接口转换成客户希望的另一个接口，使接口不兼容的那些类可以一起工作
2. **Bridge**，**桥接模式**：在go中，原本参数传入具体struct，参数换为一个接口，以此来传入不同的实现该接口的struct
3. 



## Behavioral Pattern 行为模式

