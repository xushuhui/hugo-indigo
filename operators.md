
# Java 运算符

我们已经学会了如何声明变量和赋值变量。那么你可能想知道如何对它们进行操作。本小节我们学习的运算符就可以对它们进行运算。

运算符是一些特殊的符号，它们对一个，两个或多个**操作数**执行特定的运算，然后返回一个结果。这里的操作数指的就是运算符操作的实体。

`Java`提供了一组丰富的运算符来操作变量。 我们可以将所有`Java`运算符分为以下几类：

* 算术运算符
* 关系运算符
* 位运算符
* 逻辑运算符
* 赋值运算符
* 其他运算符

本小节我们将按照以上分类介绍 Java 中的运算符，并在最后介绍运算符的优先级。

## 1. 算术运算符

### 1.1 概述

Java 语言提供了执行加减乘除四则运算的运算符。 算数运算符被用在数学表达式中，可以使用任意嵌套的小括号，其作用与数学中相同。下表列出了算术运算符：

（在例子中，初始化两个整型变量a、b：`int a = 2;` `int b = 4;`）
<table><thead><tr><th align="center">运算符</th><th align="center">描述</th><th align="center">例子</th></tr></thead><tbody><tr><td align="center">`+`</td><td align="center">加法运算符 （也用于字符串连接）</td><td align="center">a + b 等于 6</td></tr><tr><td align="center">`-`</td><td align="center">减法运算符</td><td align="center">a - b 等于 -2</td></tr><tr><td align="center">`*`</td><td align="center">乘法运算符</td><td align="center">a * b 等于 8</td></tr><tr><td align="center">`/`</td><td align="center">除法运算符</td><td align="center">b / a 等于 2</td></tr><tr><td align="center">`%`</td><td align="center">取余运算符</td><td align="center">b % a 等于 0</td></tr><tr><td align="center">`++`</td><td align="center">自增运算符</td><td align="center">a ++ 等于 3</td></tr><tr><td align="center">`--`</td><td align="center">自减运算符</td><td align="center">b – 等于 3</td></tr></tbody></table>
### 1.2 实例

以下是算数运算符的实例程序。

加法运算符使用实例：
```
public class ArithmeticOperators1 {
    public static void main(String[] args) {
        // 声明两个整型变量 num1,num2 分别赋值为 2, 3
        int num1 = 2, num2 = 3;
        // 使用加法运算符，对num1和num2执行相加操作，并将返回结果赋值给result
        int result = num1 + num2;
        System.out.println( num1 + " + " + num2 + " = " + result);
    }
}
```
运行结果：

```
2 + 3 = 5
```

减法运算符使用实例：
```
public class ArithmeticOperators2 {
    public static void main(String[] args) {
        // 声明两个整型变量 num1, num2 分别赋值为 5, 3
        int num1 = 5, num2 = 3;
        // 使用减法运算符，让 num1 加上 num2，并将返回结果赋值给result
        int result = num1 + num2;
        System.out.println( num1 + " + " + num2 + " = " + result);
        // 计算 num2 - num1 的结果
        int result1 = num2 - num1;
        System.out.println( num2 + " - " + num1 + " = " + result1);
    }
}
```
运行结果：

```
5 + 3 = 8
3 - 5 = -2
```

乘法运算符使用实例：
```
public class ArithmeticOperators3 {
    public static void main(String[] args) {
        // 声明两个整型变量 num1, num2 分别赋值为 2, 5
        int num1 = 2, num2 = 5;
        // 使用乘法运算符，让 num1 与 num2 相乘，并将返回结果赋值给result
        int result = num1 * num2;
        System.out.println( num1 + " * " + num2 + " = " + result);
    }
}
```
运行结果：

```
2 x 5 = 10
```

除法运算符使用实例：
```
public class ArithmeticOperators4 {
    public static void main(String[] args) {
        // 声明两个整型变量 num1, num2 分别赋值为 10, 2
        int num1 = 10, num2 = 2;
        // 使用除法运算符，让 num1 除以 num2，并将返回结果赋值给result
        int result = num1 / num2;
        System.out.println( num1 + " / " + num2 + " = " + result);
    }
}
```
运行结果：

```
10 / 2 = 5
```

取余运算符使用实例：
```
public class ArithmeticOperators5 {
    public static void main(String[] args) {
        // 声明两个整型变量 num1, num2 分别赋值为 5, 2
        int num1 = 5, num2 = 2;
        // 使用取余运算符，让 num1 对 num2 取余，并将返回结果赋值给result
        int result = num1 % num2;
        System.out.println( num1 + " % " + num2 + " = " + result);
    }
}
```
编译执行代码，屏幕上将会打印：

```
5 % 2 = 1
```

自增、自减运算符使用实例：
```
public class ArithmeticOperators6 {
    public static void main(String[] args) {
        // 声明两个整型变量 num1, num2 分别赋值为 5, 2
        int num1 = 5, num2 = 2;
        // 打印 num1 和 num2
        System.out.println("num1=" + num1);
        System.out.println("num2=" + num2);
        num1 ++;
        num2 --;
        System.out.println("num1自增后：" + num1);
        System.out.println("num2自减后：" + num2);
    }
}
```
运行结果：

```
num1=5
num2=2
num1自增后：6
num2自减后：1
```

另外，整型之间的除法运算是整除，这也就意味着`2 / 4`的结果为`0`，如果我们想像数学一样得到一个小数，可以使用浮点类型的操作数。例如：
```
public class OperatorDemo1 {
    public static void main(String[] args) {
        // 分组初始化两个整型变量i1、i2，值分别为2、4
        int i1 = 2, i2 = 4;
        // 使i1除以i2，并将返回结果赋值给i3
        int i3 = i1 / i2;
        System.out.println("整型2除以整型4的结果为：" + i3);
        // 分组初始化两个浮点型变量f1、f2，值分别为2、4
        float f1 = 2f, f2 = 4f;
        // 使f1除以f2，并将返回结果赋值给f3
        float f3 = f1 / f2;
        System.out.println("浮点型2除以浮点型4的结果为：" + f3);
    }
}
```
运行结果：

```
整型2除以整型4的结果为：0
浮点型2除以浮点型4的结果为：0.5
```

要特别注意，在 Java 语言中，算数运算符不能作用于不同类型的操作数。我们来看一个反例：

```
public class OperatorDemo1 {
    public static void main(String[] args) {
        // 初始化布尔类型的变量b，值为true
        boolean b = true;
        // 初始化整型变量i，值为20
        int i = 20;
        // 使用加法运算符，让i与b相加
        int result = i + b;
    }
}
```

编译代码，将会报错：

```
javac OperatorDemo1.java
OperatorDemo1.java:7: 错误: 二元运算符 '+' 的操作数类型错误
        int result = i + b;
                       ^
  第一个类型:  int
  第二个类型: boolean
1 个错误
```

编译器给出明确提示：加法运算符的操作数类型错误。这是因为 Java 是强类型的语言，不同类型的操作数的算数运算是违规的，这个原理同样适用于其他算数运算符，此处不再一一列举。

还有一点，自增自减运算符是有前后之分的，`++i`表示先加1再引用`i`，`i++`表示先引用`i`再加1。将在下一小节举例介绍。

## 2. 关系运算符

### 2.1 概述

关系运算符又称为**比较运算符**，比较的结果是一个布尔类型的值（`true`或`false`）。

Java 语言有几个可用于比较变量的运算符，如下表所示：

（在例子中，初始化两个整型变量a、b：`int a = 2;` `int b = 4;`）
<table><thead><tr><th align="center">运算符</th><th align="center">描述</th><th align="center">例子</th></tr></thead><tbody><tr><td align="center">`==`</td><td align="center">检查如果两个操作数的值是否相等，如果相等则条件为真。</td><td align="center">(a == b) 为假</td></tr><tr><td align="center">`!=`</td><td align="center">检查如果两个操作数的值是否相等，如果值不相等则条件为真。</td><td align="center">(a != b) 为真</td></tr><tr><td align="center">`>`</td><td align="center">检查左操作数的值是否大于右操作数的值，如果是那么条件为真。</td><td align="center">(a > b) 为假</td></tr><tr><td align="center">`<`</td><td align="center">检查左操作数的值是否小于右操作数的值，如果是那么条件为真。</td><td align="center">(a < b)为真</td></tr><tr><td align="center">`>=`</td><td align="center">检查左操作数的值是否大于或等于右操作数的值，如果是那么条件为真。</td><td align="center">(a >= b)为假</td></tr><tr><td align="center">`<=`</td><td align="center">检查左操作数的值是否小于或等于右操作数的值，如果是那么条件为真。</td><td align="center">(a <= b)为真</td></tr></tbody></table>
> > 

> **Tips**：在比较两个操作数是否相等时，必须使用“`==`”而不是“`=`”。



### 3.2 实例

下面是一个比较运算符的实例程序：
```
public class OperateDemo2 {
    public static void main(String[] args) {
        // 初始化一个双精度浮点型变量d1，值为10
        double d1 = 10;
        // 初始化一个整型变量i1，值也为10
        int i1 = 10;
        System.out.println("id == i1的结果为：" + (d1 == i1));
        System.out.println("id != i1的结果为：" + (d1 != i1));
        System.out.println("id > i1的结果为：" + (d1 > i1));
        System.out.println("id < i1的结果为：" + (d1 < i1));
        System.out.println("id >= i1的结果为：" + (d1 >= i1));
        System.out.println("id <= i1的结果为：" + (d1 <= i1));
    }
}
```
运行结果：

```
id == i1的结果为：true
id != i1的结果为：false
id > i1的结果为：false
id < i1的结果为：false
id >= i1的结果为：true
id <= i1的结果为：true
```

> > 

> **Tips**：`>`、`<`、`>=`、`<=` 这几个运算符左右两边的操作数必须是`byte`，`short`，`int`，`long`，`double`，`float`，`char`这几种数据类型；而`==`和`!=`运算符的操作数既可以是基本数据类型，又可以是引用数据类型。



## 3. 位运算符

Java 语言还提供了对整数类型执行按位和移位操作的运算符，称作**位运算符**。

它在实际的编码中并不常用，这部分内容了解即可。

假设`a = 60，b = 13`;它们的二进制格式表示将如下：

```
a = 0011 1100
b = 0000 1101
-----------------
a & b = 0000 1100
a | b = 0011 1101
a ^ b = 0011 0001
~a = 1100 0011
```

下表列出了位运算符的基本运算，假设整数变量 a 的值为 60 和变量 b 的值为 13：
<table><thead><tr><th align="center">运算符</th><th align="center">描述</th><th align="center">例子</th></tr></thead><tbody><tr><td align="center">＆（按位与）</td><td align="center">如果相对应位都是1，则结果为1，否则为0</td><td align="center">（a＆b），得到12，即0000 1100</td></tr><tr><td align="center">|（按位或）</td><td align="center">如果相对应位都是 0，则结果为 0，否则为 1</td><td align="center">（a | b）得到61，即 0011 1101</td></tr><tr><td align="center">^（按位异或）</td><td align="center">如果相对应位值相同，则结果为0，否则为1</td><td align="center">（a ^ b）得到49，即 0011 0001</td></tr><tr><td align="center">〜（按位取反）</td><td align="center">按位取反运算符翻转操作数的每一位，即0变成1，1变成0。</td><td align="center">（〜a）得到-61，即1100 0011</td></tr><tr><td align="center"><< （左位移）</td><td align="center">按位左移运算符。左操作数按位左移右操作数指定的位数。</td><td align="center">a << 2得到240，即 1111 0000</td></tr><tr><td align="center">>> （右位移）</td><td align="center">按位右移运算符。左操作数按位右移右操作数指定的位数。</td><td align="center">a >> 2得到15即 1111</td></tr><tr><td align="center">>>> （零填充右移）</td><td align="center">按位右移补零操作符。左操作数的值按右操作数指定的位数右移，移动得到的空位以零填充。</td><td align="center">a>>>2得到15即0000 1111</td></tr></tbody></table>

## 4. 逻辑运算符

### 4.1 概述

逻辑运算符可以在表达式中生成组合条件，例如在执行特定语句块之前必须满足的两个或多个条件。使用逻辑运算符，可以描述这些组合条件。逻辑运算的返回结果只能为真或假。

Java 语言中的逻辑运算符，如下表所示：

（在例子中，初始化两个整型变量a、b：`int a = 0;` `int b = 1;`）
<table><thead><tr><th align="center">运算符</th><th align="center">描述</th><th align="center">例子</th></tr></thead><tbody><tr><td align="center">&&（逻辑与）</td><td align="center">当且仅当两个操作数都为真，条件才为真。</td><td align="center">(a && b)为假</td></tr><tr><td align="center">|| （逻辑或）</td><td align="center">如果任何两个操作数任何一个为真，条件为真。</td><td align="center">(a || b)为真</td></tr><tr><td align="center">！（逻辑非）</td><td align="center">用来反转操作数的逻辑状态。如果条件为真，则逻辑非运算符将得到假。</td><td align="center">!(a && b)为假</td></tr></tbody></table>
### 4.2 短路运算

`&&`和`||`运算符存在**短路**行为。短路的意思是：只有在需要的时候才会判断第二个操作数的真假。例如：
```
class LogicOperators {
    public static void main(String[] args){
        int a = 1, b = 2;
        if((a == 2) && (b == 2)) {
            System.out.println("a和b都等于2");
        }
        if((a == 1) || (b == 1)) {
            System.out.println("a等于1或b等于1");
        }
    }
}
```
运行结果：

```
a等于1或b等于1
```

程序解析：有两个整型变量`a`和`b`，值分别为`1`和`2`。第一个`if`语句的条件为逻辑与运算，其第一个操作数`a == 2`为假，所以无论第二个操作数是真是假，都不去判断，条件直接被判定为假；第二个`if`语句的条件为逻辑或运算， 其第一个操作数`a == 1`为真，所以无论第二个操作数是真是假，都不会去判断，条件直接被判定为真。这就是所谓的短路。

## 5. 赋值运算符

### 5.1 概述

赋值运算符是为指定变量分配值的符号。下标列出了常用 Java 中常用的赋值运算符：
<table><thead><tr><th align="center">运算符</th><th align="center">描述</th><th align="center">例子</th></tr></thead><tbody><tr><td align="center">=</td><td align="center">简单的赋值运算符。将值从右侧操作数分配给左侧操作数。</td><td align="center">c = a + b将a + b的值赋给c</td></tr><tr><td align="center">+=</td><td align="center">加和赋值运算符。它将右操作数添加到左操作数，并将结果分配给左操作数。</td><td align="center">c + = a等于c = c + a</td></tr><tr><td align="center">-=</td><td align="center">减和赋值运算符。它从左侧操作数中减去右侧操作数，并将结果分配给左侧操作数。</td><td align="center">c -= a等效于c = c – a</td></tr><tr><td align="center">*=</td><td align="center">乘和赋值运算符。它将右操作数与左操作数相乘，并将结果分配给左操作数。</td><td align="center">c *= a等效于c = c * a</td></tr><tr><td align="center">/ =</td><td align="center">除和赋值运算符。它将左操作数除以右操作数，并将结果分配给左操作数。</td><td align="center">c /= a等于c = c / a</td></tr></tbody></table>
### 5.2 实例

我们来看一个赋值运算符的实例：
```
public class OperateDemo5 {
    public static void main(String[] args) {
        // 分组初始化三个变量 num1、num2、result，值分别为 20、10、0
        int num1 = 20, num2 = 10,  result = 0;
        System.out.println("初始值：");
        System.out.print("num1=" + num1 + '\t');
        System.out.print("num2=" + num2 + '\t');
        System.out.print("result=" + result + "\t\n");
        System.out.println("开始赋值运算：");
        result += num1;
        System.out.println("result += num1 结果为：" + result);
        result -= num2;
        System.out.println("result -= num2 结果为：" + result);
        result *= num1;
        System.out.println("result *= num1 结果为：" + result);
        result /= num2;
        System.out.println("result /= num2 结果为：" + result);
        result %= 15;
        System.out.println("result %= 15 结果为：" + result);
    }
}
```
运行结果：

```
初始值：
num1=20	num2=10	result=0	
开始赋值运算：
result += num1 结果为：20
result -= num2 结果为：10
result *= num1 结果为：200
result /= num2 结果为：20
result %= 15 结果为：5
```

## 6. 其他运算符

### 6.1 条件运算符（? :）

条件运算符也称为三元运算符。我们会在条件语句小节中再次对其介绍。

该运算符由三个操作数组成，用于判断**布尔表达式**。它的目的是确定应将哪个值分配给变量。条件运算符的语法为：

```
变量 = (布尔表达式) ? 值1 : 值2
```

如果布尔表达式为真，就将`值1`分配变量，否则将`值2`分配给变量。

下面是一个实例程序：
```
public class ConditionalOperators {
    public static void main(String[] args) {
        int age = 15;
        System.out.println(age >= 18 ? "在中国你已经成年" :  "在中国你还未成年"); 
    }
}
```
由于`age`变量值为15，小于18，`age >= 18`返回结果为假，因此编译执行后，屏幕将打印：

```
在中国你还未成年
```

### 6.2 instanceof 运算符

> > 

> **Tips**：了解 instanceof 运算符需要一些面向对象的前置知识。目前你可以选择性学习。



`instanceof`运算符将对象与指定类型进行比较，检查对象是否是一个特定类型（类类型或接口类型）。

instanceof 运算符的语法为：

```
( Object reference variable ) instanceof  (class/interface type)
```

如果`instanceof`左侧的变量所指向的对象，是`instanceof`右侧类或接口的一个对象，结果为真，否则结果为假。
```
public class InstanceOfOperators1 {
    public static void main(String[] args) {
        String name = "imooc";
        boolean b = name instanceof String; 
     	  System.out.println("结果为：" + b);
    }
}
```
由于字符串变量`name`是`String`类型，所以执行代码，屏幕会打印：

```
结果为：true
```

注意，`instanceof`运算符不能用于操作基本数据类型，如果将字符串类型`name`变量改为一个`char`类型的变量，编译代码将会报错:

```
InstanceOfOperators1.java:4: 错误: 意外的类型
        boolean b = name instanceof String;
                    ^
  需要: 引用
  找到:    char
1 个错误
```

## 7. 运算符的优先级

当多种运算符在一同使用的时候，会有一个执行先后顺序的问题。

下表中的运算符按优先顺序排序。运算符越靠近表格顶部，其优先级越高。具有较高优先级的运算符将在具有相对较低优先级的运算符之前计算。同一行上的运算符具有相同的优先级。
<table><thead><tr><th align="center">类别</th><th align="center">操作符</th><th align="center">关联性</th></tr></thead><tbody><tr><td align="center">后缀</td><td align="center">() [] . (点操作符)</td><td align="center">左到右</td></tr><tr><td align="center">一元</td><td align="center">+ + - ！〜</td><td align="center">**从右到左**</td></tr><tr><td align="center">乘性</td><td align="center">* /％</td><td align="center">左到右</td></tr><tr><td align="center">加性</td><td align="center">+ -</td><td align="center">左到右</td></tr><tr><td align="center">移位</td><td align="center">>> >>> <<</td><td align="center">左到右</td></tr><tr><td align="center">关系</td><td align="center">>> = << =</td><td align="center">左到右</td></tr><tr><td align="center">相等</td><td align="center">== !=</td><td align="center">左到右</td></tr><tr><td align="center">按位与</td><td align="center">＆</td><td align="center">左到右</td></tr><tr><td align="center">按位异或</td><td align="center">^</td><td align="center">左到右</td></tr><tr><td align="center">按位或</td><td align="center">|</td><td align="center">左到右</td></tr><tr><td align="center">逻辑与</td><td align="center">&&</td><td align="center">左到右</td></tr><tr><td align="center">逻辑或</td><td align="center">| |</td><td align="center">左到右</td></tr><tr><td align="center">条件</td><td align="center">？：</td><td align="center">**从右到左**</td></tr><tr><td align="center">赋值</td><td align="center">= + = - = * = / =％= >> = << =＆= ^ = | =</td><td align="center">**从右到左**</td></tr><tr><td align="center">逗号</td><td align="center">，</td><td align="center">左到右</td></tr></tbody></table>
当**相同优先级**的运算符出现在同一表达式中时，如何控制它们计算的先后呢。我们来看一个实例：
```
public class OperateorsPriority {
    public static void main(String[] args) {
      	int a = 2;
      	int b = 4;
      	int c = 6;
      	int result = a + b - c + a;
     	System.out.println("result = " + result);
    }
}
```
在计算`result`的语句的右侧，`+` 、`-`两个运算符优先级相同，如果我们不加以控制，将按照从左到右顺序计算，打印结果为`result = 2`；但是如果我们想先计算`a + b`和`c + a`的值，再计算两者之差，我们可以使用括号`()`将其顺序进行控制：`(a + b) - (c + a)`，再执行代码将打印我们想要的结果：`result = -2`。

## 8. 小结

本小节我们按照分类介绍了`Java`的运算符，并且在最后给出了运算符的优先级表格。本节内容可能有些繁杂，但细节不必全部记住，你可以在需要的时候回过头来查查表格。当然，最好的方式还是多写多练，对照实例和表格自己敲敲代码，用心分析代码执行的流程，是增强记忆最好的方式。

* 划线
* 写笔记
<span class="imv2-close2 close"></span><textarea placeholder="学习要认真，笔记应当先"></textarea>
<label><span class="icon imv2-checkbox"></span><span>公开笔记</span></label><span class="number">0/1000</span><button class="submit gary">提交</button>
![]()<span class="nick-name"></span><span class="imv2-close2 close"></span>
```

```
<button class="del">删除</button><button class="edit">编辑</button>