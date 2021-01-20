# 第三节exper表达式

前提：注册模型：RegisterModel

## 一、使用QueryTable

```
1.
	o := orm.NewOrm\(\)

	qs := o.QueryTable\("user"\) 表名作为参数
2.
	qs := o.QueryTable\(new\(User\)\) 也可以直接使用对象作为表名

返回的是QuerySeter
```

## 二、exper表达式的使用：两个下换线

1.exact / iexact:等于，默认值，大小写敏感 / 不敏感

```
qs.Filter("name__exact","Zhiliao").One(&stu)
```

2.contains / icontains:包含，大小写敏感 / 不敏感

```
qs.Filter("name__contains","Zhil").One(&stu)
```

3.gt / gte：大于/大于等于

```
qs.Filter("age__gt",18).One(&stu)
```

4.lt / lte：小于/小于等于

```
qs.Filter("age__lt",18).One(&stu)
```

5.startswith / istartswith：以…起始，大小写敏感 / 不敏感

```
qs.Filter("name__startswith","Zh").One(&stu)
qs.Filter("name__istartswith","Zh").One(&stu)
```

6.endswith / iendswith：以…结束，大小写敏感 / 不敏感

```
qs.Filter("name__endswith","Liao").One(&stu)
qs.Filter("name__iendswith","Liao").One(&stu)
```

7.in：在某个范围中,值为不定长参数

```
qs.Filter("age__in",12,13,16,19).One(&stu)
```

8.isnull：为空，值为 true / false

```
qs.Filter("gender__isnull",true).One(&stu)
```

#### 查看orm执行的sql语句：

- 只作用于当前的模型：当前模型的init中orm.Debug = true
- 作用于所有模型:beego.Run()前面设置orm.Debug = true