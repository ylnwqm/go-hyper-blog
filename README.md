# go-hyper-blog
Super speed blog developed using Go


## 初衷
刚刚由PHP转为Go，建立此git仓库也是为了能够快速上手Go相关知识

## 项目定义
此项目定义为博客，经过多方选型，选用iris框架作为骨架包进行学习开发

## 为什么选用iris
**从图中可以看出，iris对于各个模块的支持相对来说比较全面，虽然其他框架也有相对基础的相关模块，但是为了不过多的重新造轮子，尽量选用全而美**

![img.png](https://pic4.zhimg.com/v2-49443c9154b18ab93979f8867a481e87_r.jpg)

其实，选用iris心里也有一些忐忑，毕竟作用国产框架，Beego的文档显然比iris支持的更好

## 开发

### 版本
选用较新的1.18

>why: 因为从1.18开始，go支持了泛型操作

### 页面渲染
**不会前端**，更别提vue react这些框架
也没有过多的精力去仔细研究css和js
那么只能拿前端框架来改改

作为phper，还是有几个混开blog比较合适
这里选用了白俊遥的博客template
> 开发中将选用这套模板进行数据嵌套渲染，其中所有的css和js在原则上不予变动

### 数据库

这里也使用了白大大的数据表进行开发

### 第三方扩展

laravel框架集成了相当便利的扩展包，而go在这方面可能还会有些许缺失，在原则上，博客会先简单的实现简单的功能



